// Copyright 2021 EricWinn
// Author:   Eric Winn
// Email:    eng.eric.winn@gmail.com
// Time:     2023/6/26 21:03
// File:     main.go
// Software: GoLand

package main

import (
	"encoding/json"
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/itnotebooks/chinese-holiday/utils/http"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	SearchUrl = "http://sousuo.gov.cn/s.htm"
)

var (
	wg = sync.WaitGroup{}
)

type HolidayDate struct {
	Year     int
	Name     string
	Date     time.Time
	IsOffDay bool
}

type Holiday struct {
	Url          string
	Year         int
	HolidayDates []HolidayDate
	Container    []soup.Root
}

func InitHolidayParse(url string, year int) *Holiday {
	return &Holiday{
		Url:  url,
		Year: year,
	}
}

// FetchPage 请求页面并定位到 id = UCAP-CONTENT 的 div 容器，读取所有的 p 标签条目
func (s *Holiday) FetchPage() error {

	r, err := http.Get(s.Url, map[string]interface{}{})
	if err != nil {
		return err
	}
	//定位到 id = UCAP-CONTENT 的 div 容器，读取所有的 p 标签条目
	s.Container = soup.HTMLParse(r).Find("div", "id", "UCAP-CONTENT").FindAll("p")

	if len(s.Container) == 0 {
		return fmt.Errorf("Page parse error ")
	}
	return nil
}

// ParseRules 通过正则分析每个 p 标签的内容，判断是否为大写数字开头的序号，大写数字开头的序号为具体放假安排，分析 休息日还是工作日
func (s *Holiday) ParseRules() {
	for _, p := range s.Container {
		if p.Text() == "" {
			continue
		}

		//判断是否为大写数字开头的序号，大写数字开头的序号为具体放假安排
		mRegex := regexp.MustCompile(`[一二三四五六七八九十]、(.+?)：(.+)`)
		match := mRegex.FindStringSubmatch(p.FullText())
		if len(match) <= 2 {
			continue
		}

		//分段处理，降低匹配复杂度
		for _, str := range regexp.MustCompile("[，。；]").Split(match[2], -1) {
			if str == "" {
				continue
			}

			//获取休息日
			rest := regexp.MustCompile(`(.+)(放假|补休|调休|公休)+(?:\d+天)?$`).FindStringSubmatch(str)
			if len(rest) > 2 {
				//解析具体日期
				s.ExtractDates(match[1], rest[1], true)
				continue
			}

			//获取工作日
			work := regexp.MustCompile(`(.+)上班$`).FindStringSubmatch(str)
			if len(work) > 1 {
				// 解析具体日期
				s.ExtractDates(match[1], work[1], false)
				continue
			}
		}

	}
}

// IsExist 判断是否已存在
func (s *Holiday) IsExist(date string) bool {
	for _, i := range s.HolidayDates {
		if i.Date.Format("2006-1-2") == date {
			return true
		}
	}
	return false
}

// GetDate 日期字符串转成日期对象
func (s *Holiday) GetDate(y, m, d string) time.Time {
	t, _ := time.ParseInLocation("2006-1-2", fmt.Sprintf("%s-%s-%s", y, m, d), time.Local)
	return t
}

// ExtractDates 分析具体放假安排，取对应的年月日关键字
func (s *Holiday) ExtractDates(name, txt string, offDay bool) {
	txt = strings.ReplaceAll(txt, "(", "（")
	txt = strings.ReplaceAll(txt, ")", "）")

	//[xxxx年][x月]x日
	matches := regexp.MustCompile(`(?:(\d+)年)?(?:(\d+)月)?(\d+)日`).FindAllStringSubmatch(txt, -1)
	for _, match := range matches {
		if match[2] == "" {
			continue
		}

		if match[1] == "" {
			match[1] = strconv.Itoa(s.Year)
		}

		if s.IsExist(fmt.Sprintf("%s-%s-%s", match[1], match[2], match[3])) {
			continue
		}
		s.HolidayDates = append(s.HolidayDates, HolidayDate{
			s.Year,
			name,
			s.GetDate(match[1], match[2], match[3]),
			offDay,
		})
	}

	//[xxxx年]x月x日至[xxxx年][x月]x日
	ext2txt := regexp.MustCompile(`（.+?）`).ReplaceAllString(txt, "")
	matches = regexp.MustCompile(`(?:(\d+)年)?(?:(\d+)月)?(\d+)日(?:至|-|—)(?:(\d+)年)?(?:(\d+)月)?(\d+)日`).
		FindAllStringSubmatch(ext2txt, -1)
	for _, match := range matches {
		if len(match) < 6 {
			continue
		}

		if match[1] == "" {
			match[1] = strconv.Itoa(s.Year)
		}
		if match[4] == "" {
			match[4] = strconv.Itoa(s.Year)
		}

		if match[5] == "" {
			match[5] = match[2]
		}

		start := s.GetDate(match[1], match[2], match[3])
		end := s.GetDate(match[4], match[5], match[6])
		//解析日期范围
		for i := 0; i <= int(end.Sub(start).Hours()/24); i++ {
			d := s.GetDate(match[1], match[2], match[3]).AddDate(0, 0, i)

			if s.IsExist(d.Format("2006-1-2")) {
				continue
			}
			s.HolidayDates = append(s.HolidayDates, HolidayDate{
				s.Year,
				name,
				d,
				offDay,
			})
		}
	}

	//x月x日(星期x)、x月x日(星期x)
	ext3txt := regexp.MustCompile(`（.+?）`).ReplaceAllString(txt, "")
	matches = regexp.MustCompile(
		`(?:(\d+)年)?(?:(\d+)月)?(\d+)日(?:（[^）]+）)?(?:、(?:(\d+)年)?(?:(\d+)月)?(\d+)日(?:（[^）]+）)?)+`,
	).FindAllStringSubmatch(ext3txt, -1)
	for _, match := range matches {

		if len(match) < 6 {
			continue
		}

		if match[1] == "" {
			match[1] = strconv.Itoa(s.Year)
		}
		if match[4] == "" {
			match[4] = strconv.Itoa(s.Year)
		}

		if match[5] == "" {
			match[5] = match[2]
		}
		d := s.GetDate(match[1], match[2], match[3])
		if !s.IsExist(d.Format("2006-1-2")) {
			s.HolidayDates = append(s.HolidayDates, HolidayDate{
				s.Year,
				name,
				d,
				offDay,
			})
		}

		d = s.GetDate(match[4], match[5], match[6])
		if !s.IsExist(d.Format("2006-1-2")) {
			s.HolidayDates = append(s.HolidayDates, HolidayDate{
				s.Year,
				name,
				d,
				offDay,
			})
		}
	}

}

func Search(y int) {
	defer wg.Done()
	year := strconv.Itoa(y)

	//检索关键字
	queryParams := map[string]interface{}{
		"t":           "paper",
		"advance":     "true",
		"title":       year,
		"q":           "假期",
		"pcodeJiguan": "国办发明电",
		"puborg":      "国务院办公厅",
	}

	ret, err := http.Get(SearchUrl, queryParams)
	if err != nil {
		log.Fatalln(fmt.Sprintf("查询 %d 放假通知异常", y))
	} else {
		// 分析页面，获取对应年份的放假通知条目具体页面的 url
		urls := regexp.MustCompile(`<li class="res-list".*?<a href="(.+?)".*?</li>`).FindAllStringSubmatch(strings.Replace(ret, "\n", "", -1), -1)
		for _, url := range urls {
			if len(url) < 2 {
				continue
			}

			if url[1] == "" {
				continue
			}

			// 请求具体通知页面，并分析放假安排
			log.Printf("[ %d ] ====> %s", y, url[1])
			holiday := InitHolidayParse(url[1], y)
			if err = holiday.FetchPage(); err != nil {
				fmt.Errorf("获取并分析 %d 放假通知页面，异常\n%s\n", y, url[1])
				continue
			}
			holiday.ParseRules()

			// todo: 结果处理
			jsonStr, _ := json.Marshal(holiday.HolidayDates)
			fmt.Println(string(jsonStr))
		}

	}
}

/*
 1. 请求 BaseSearch url 查询指定年份的放假通知条目
 2. 请求第1步查询到的放假通知页面的 URL
 3. 分析页面，定位到 id = UCAP-CONTENT 的 div 容器，读取所有的 p 标签条目
 4. 通过正则分析每个 p 标签的内容，判断是否为大写数字开头的序号，大写数字开头的序号为具体放假安排
 5. 分析具体放假安排，取对应的年月日关键字；通过分析过往几年的通知内容，规律如下：
    5.1 休息日还是工作日，会以以下两种文言描述
    5.1.1 休息日：放假|补休|调休|公休
    5.1.2 工作日：上班
    5.2 具体的日期，会以以下三种文言描述：
    5.2.1 [xxxx年]x月x日至[xxxx年][x月]x日
    5.2.2 x月x日(星期x)、x月x日(星期x)

以下为2023年的放假安排：
----------

	国务院办公厅关于2023年
	部分节假日安排的通知
	国办发明电〔2022〕16号

	各省、自治区、直辖市人民政府，国务院各部委、各直属机构：
	经国务院批准，现将2023年元旦、春节、清明节、劳动节、端午节、中秋节和国庆节放假调休日期的具体安排通知如下。
	一、元旦：2022年12月31日至2023年1月2日放假调休，共3天。
	二、春节：1月21日至27日放假调休，共7天。1月28日（星期六）、1月29日（星期日）上班。
	三、清明节：4月5日放假，共1天。
	四、劳动节：4月29日至5月3日放假调休，共5天。4月23日（星期日）、5月6日（星期六）上班。
	五、端午节：6月22日至24日放假调休，共3天。6月25日（星期日）上班。
	六、中秋节、国庆节：9月29日至10月6日放假调休，共8天。10月7日（星期六）、10月8日（星期日）上班。
	节假日期间，各地区、各部门要妥善安排好值班和安全、保卫、疫情防控等工作，遇有重大突发事件，要按规定及时报告并妥善处置，确保人民群众祥和平安度过节日假期。
	国务院办公厅
	2022年12月8日

----------
*/
func main() {
	// 获取当前及下一个年份
	for y := time.Now().Year(); y <= time.Now().Year()+1; y++ {
		wg.Add(1)
		go Search(y)
	}
	wg.Wait()
}
