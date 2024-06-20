package main

/*
https://cloudmessage.top/archives/go%E5%AE%9E%E7%8E%B0%E5%B8%B8%E7%94%A8%E7%9A%84%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E5%B7%A5%E5%8E%82%E6%A8%A1%E5%BC%8F
*/
import "fmt"

/*

 */
// Printer 简单工厂要返回的接口类型
type Printer interface {
	Print(name string) string
}

// CnPrinter 是 Printer 接口的实现，它说中文
type CnPrinter struct{}

func (*CnPrinter) Print(name string) string {
	return fmt.Sprintf("你好, %s", name)
}

// EnPrinter 是 Printer 接口的实现，它说中文
type EnPrinter struct{}

func (*EnPrinter) Print(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func (*EnPrinter) Print1(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

// 打印机的构造函数
func NewPrinter(lang string) Printer {
	//实例化
	switch lang {
	case "cn":
		return new(CnPrinter)
	case "en":
		return new(EnPrinter)
	default:
		return new(CnPrinter)
	}
}

func main() {
	printer := NewPrinter("en")
	fmt.Println(printer.Print("Bob"))

}
