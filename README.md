# 工作日计算器

一个简单的 GUI 工具，用于计算两个日期之间的工作日数量，支持中国法定节假日和调休安排。

## 功能特点

- 图形用户界面，简单易用
- 支持计算任意两个日期之间的工作日数量
- 自动考虑周末（周六日）
- 包含 2024 年中国法定节假日数据
- 支持调休安排（如节假日期间的补班）

## 使用要求

- Go 1.16 或更高版本
- 安装 Fyne GUI 工具包的依赖：
  - Windows: 需要安装 gcc 和 MinGW-w64
  - Linux: 需要安装 gcc 和相关开发工具
  - macOS: 需要安装 Xcode Command Line Tools

## 导入包并运行
```
go mod init workdays
go get fyne.io/fyne/v2@latest
go mod tidy
go run .
```