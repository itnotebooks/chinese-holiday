// Copyright 2021 EricWinn
// Author:   Eric Winn
// Email:    eng.eric.winn@gmail.com
// Time:     2023/6/26 15:37
// File:     http.go
// Software: GoLand

package http

import (
	"bytes"
	"fmt"
	"github.com/itnotebooks/chinese-holiday/utils/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

var httpClient = &http.Client{}

type HeaderOption struct {
	Name  string
	Value string
}

func responseHandle(resp *http.Response, err error) (string, error) {
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	respBody := string(b)
	//log.Printf("api %s 响应结果: %s", resp.Request.URL, respBody)
	return respBody, nil
}

func GetRequest(url string, headerOptions ...HeaderOption) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	for _, headerOption := range headerOptions {
		req.Header.Set(headerOption.Name, headerOption.Value)
	}
	resp, err := httpClient.Do(req)
	defer func() {
		if resp != nil {
			if e := resp.Body.Close(); e != nil {
				fmt.Println(e)
			}
		}
	}()
	return responseHandle(resp, err)
}
func ConvertToQueryParams(params map[string]interface{}) string {
	paramsJson := json.ToJsonIgnoreError(params)
	params = map[string]interface{}{}
	_ = json.FromJson(paramsJson, &params)

	if &params == nil || len(params) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	buffer.WriteString("?")
	for k, v := range params {
		if v == nil {
			continue
		}
		buffer.WriteString(fmt.Sprintf("%s=%v&", url.QueryEscape(k), url.QueryEscape(v.(string))))
	}
	buffer.Truncate(buffer.Len() - 1)
	return buffer.String()
}

func Get(url string, params map[string]interface{}, headerOptions ...HeaderOption) (string, error) {
	fullUrl := url + ConvertToQueryParams(params)
	return GetRequest(fullUrl, headerOptions...)
}
