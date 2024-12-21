package main

import (
    "fmt"
    "time"
)

// Holiday 结构体用于解析节假日API的返回数据
type Holiday struct {
    Name   string
    Date   string
    IsWork bool // true表示调休工作日，false表示放假
}

// 获取指定年份的节假日信息
func getHolidays(year int) ([]Holiday, error) {
    // 2024年节假日数据
    if year == 2024 {
        return []Holiday{
            {Name: "元旦", Date: "2024-01-01", IsWork: false},
            {Name: "春节", Date: "2024-02-10", IsWork: false},
            {Name: "春节", Date: "2024-02-11", IsWork: false},
            {Name: "春节", Date: "2024-02-12", IsWork: false},
            {Name: "春节", Date: "2024-02-13", IsWork: false},
            {Name: "春节", Date: "2024-02-14", IsWork: false},
            {Name: "春节", Date: "2024-02-15", IsWork: false},
            {Name: "春节", Date: "2024-02-16", IsWork: false},
            {Name: "春节", Date: "2024-02-17", IsWork: false},
            {Name: "清明节", Date: "2024-04-04", IsWork: false},
            {Name: "清明节", Date: "2024-04-05", IsWork: false},
            {Name: "清明节", Date: "2024-04-06", IsWork: false},
            {Name: "劳动节", Date: "2024-05-01", IsWork: false},
            {Name: "劳动节", Date: "2024-05-02", IsWork: false},
            {Name: "劳动节", Date: "2024-05-03", IsWork: false},
            {Name: "劳动节", Date: "2024-05-04", IsWork: false},
            {Name: "劳动节", Date: "2024-05-05", IsWork: false},
            {Name: "端午节", Date: "2024-06-10", IsWork: false},
            {Name: "中秋节", Date: "2024-09-15", IsWork: false},
            {Name: "中秋节", Date: "2024-09-16", IsWork: false},
            {Name: "中秋节", Date: "2024-09-17", IsWork: false},
            {Name: "国庆节", Date: "2024-10-01", IsWork: false},
            {Name: "国庆节", Date: "2024-10-02", IsWork: false},
            {Name: "国庆节", Date: "2024-10-03", IsWork: false},
            {Name: "国庆节", Date: "2024-10-04", IsWork: false},
            {Name: "国庆节", Date: "2024-10-05", IsWork: false},
            {Name: "国庆节", Date: "2024-10-06", IsWork: false},
            {Name: "国庆节", Date: "2024-10-07", IsWork: false},
            // 调休工作日
            {Name: "春节调休", Date: "2024-02-04", IsWork: true},
            {Name: "春节调休", Date: "2024-02-18", IsWork: true},
            {Name: "清明节调休", Date: "2024-04-07", IsWork: true},
            {Name: "劳动节调休", Date: "2024-04-28", IsWork: true},
            {Name: "劳动节调休", Date: "2024-05-11", IsWork: true},
            {Name: "中秋节调休", Date: "2024-09-14", IsWork: true},
            {Name: "国庆节调休", Date: "2024-09-29", IsWork: true},
            {Name: "国庆节调休", Date: "2024-10-12", IsWork: true},
        }, nil
    }
    
    // 如果是当前年份，返回空的节假日列表（只考虑周末）
    if year == time.Now().Year() {
        return []Holiday{}, nil
    }
    
    return nil, fmt.Errorf("暂不支持%d年的节假日数据", year)
}

// 判断某一天是否为工作日
func isWorkday(date time.Time, holidays []Holiday) bool {
    // 周末默认为非工作日
    if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
        // 检查是否为调休工作日
        dateStr := date.Format("2006-01-02")
        for _, holiday := range holidays {
            if holiday.Date == dateStr && holiday.IsWork {
                return true
            }
        }
        return false
    }

    // 检查是否为法定节假日
    dateStr := date.Format("2006-01-02")
    for _, holiday := range holidays {
        if holiday.Date == dateStr && !holiday.IsWork {
            return false
        }
    }
    return true
}

// 计算两个日期之间的工作日数量
func countWorkdays(start, end time.Time, holidays []Holiday) int {
    count := 0
    current := start

    for current.Before(end) || current.Equal(end) {
        if isWorkday(current, holidays) {
            count++
        }
        current = current.AddDate(0, 0, 1)
    }
    return count
} 