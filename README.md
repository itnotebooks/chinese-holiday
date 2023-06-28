# chinese-holiday

抓取国务院公告，分析中国法定节假日/法定工作日

# 说明

抓取后的结果请根据自己的需求自行处理

```text
2023/06/28 16:14:12 [ 2023 ] ====> http://www.gov.cn/zhengce/content/2022-12/08/content_5730844.htm
{"Year":2023,"Name":"元旦","Date":"2022-12-31T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"元旦","Date":"2023-01-02T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"元旦","Date":"2023-01-01T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"春节","Date":"2023-01-21T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"春节","Date":"2023-01-22T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"春节","Date":"2023-01-23T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"春节","Date":"2023-01-24T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"春节","Date":"2023-01-25T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"春节","Date":"2023-01-26T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"春节","Date":"2023-01-27T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"春节","Date":"2023-01-28T00:00:00+08:00","IsOffDay":false}
{"Year":2023,"Name":"春节","Date":"2023-01-29T00:00:00+08:00","IsOffDay":false}
{"Year":2023,"Name":"清明节","Date":"2023-04-05T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"劳动节","Date":"2023-04-29T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"劳动节","Date":"2023-05-03T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"劳动节","Date":"2023-04-30T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"劳动节","Date":"2023-05-01T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"劳动节","Date":"2023-05-02T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"劳动节","Date":"2023-04-23T00:00:00+08:00","IsOffDay":false}
{"Year":2023,"Name":"劳动节","Date":"2023-05-06T00:00:00+08:00","IsOffDay":false}
{"Year":2023,"Name":"端午节","Date":"2023-06-22T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"端午节","Date":"2023-06-23T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"端午节","Date":"2023-06-24T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"端午节","Date":"2023-06-25T00:00:00+08:00","IsOffDay":false}
{"Year":2023,"Name":"中秋节、国庆节","Date":"2023-09-29T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"中秋节、国庆节","Date":"2023-10-06T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"中秋节、国庆节","Date":"2023-09-30T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"中秋节、国庆节","Date":"2023-10-01T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"中秋节、国庆节","Date":"2023-10-02T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"中秋节、国庆节","Date":"2023-10-03T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"中秋节、国庆节","Date":"2023-10-04T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"中秋节、国庆节","Date":"2023-10-05T00:00:00+08:00","IsOffDay":true}
{"Year":2023,"Name":"中秋节、国庆节","Date":"2023-10-07T00:00:00+08:00","IsOffDay":false}
{"Year":2023,"Name":"中秋节、国庆节","Date":"2023-10-08T00:00:00+08:00","IsOffDay":false}
```

# 鸣谢

主体逻辑参考 Python 开源项目[holiday-cn](https://github.com/NateScarlet/holiday-cn)