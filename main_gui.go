package main

import (
    "fmt"
    "time"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func main() {
    myApp := app.New()
    window := myApp.NewWindow("工作日计算器")

    // 创建日期输入
    entry := widget.NewEntry()
    entry.SetPlaceHolder("YYYY-MM-DD")

    // 创建结果显示标签
    result := widget.NewLabel("")

    // 创建计算按钮
    calculateBtn := widget.NewButton("计算工作日", func() {
        dateStr := entry.Text
        
        // 解析用户输入的日期
        targetDate, err := time.Parse("2006-01-02", dateStr)
        if err != nil {
            result.SetText("日期格式错误！请使用 YYYY-MM-DD 格式")
            return
        }

        today := time.Now()

        // 获取当年和目标年份的节假日信息
        currentYear := today.Year()
        targetYear := targetDate.Year()
        
        var allHolidays []Holiday
        
        // 获取所需年份的节假日信息
        for year := currentYear; year <= targetYear; year++ {
            holidays, err := getHolidays(year)
            if err != nil {
                result.SetText(fmt.Sprintf("获取%d年节假日信息失败: %v", year, err))
                return
            }
            allHolidays = append(allHolidays, holidays...)
        }

        var workdays int
        var resultText string
        
        if targetDate.After(today) {
            workdays = countWorkdays(today, targetDate, allHolidays)
            resultText = fmt.Sprintf("从今天到%s还有%d个工作日", 
                targetDate.Format("2006-01-02"), 
                workdays)
        } else if targetDate.Before(today) {
            workdays = countWorkdays(targetDate, today, allHolidays)
            resultText = fmt.Sprintf("从%s到今天已经过去了%d个工作日", 
                targetDate.Format("2006-01-02"), 
                workdays)
        } else {
            resultText = "你输入的是今天！"
        }
        
        result.SetText(resultText)
    })

    // 创建布局
    content := container.NewVBox(
        widget.NewLabel("请输入目标日期 (格式: YYYY-MM-DD)："),
        entry,
        calculateBtn,
        result,
    )

    window.SetContent(content)
    window.Resize(fyne.NewSize(300, 200))
    window.ShowAndRun()
} 