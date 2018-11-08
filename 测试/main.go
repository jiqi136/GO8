// 测试 project main.go  "github.com/chromedp-master"

package main

import (
	"fmt"

	"time"
)

//  :=
func main() { // main模具即使 前面没有 go协程命令，但依然 是协程

	fmt.Println("运r行") //有
	fmt.Println("启动 协程()")

	ch := make(chan string)

	go 模具一一协程例子(ch)
	go getData(ch)

	time.Sleep(1e9)
	fmt.Println("等待1秒")

	fmt.Println("协程结束()")

}

func 模具一一协程例子(ch chan string) {

	/*导入库
	"fmt"
	"time"  */

	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "mysql-masterTokyo"
}

func getData(ch chan string) {
	//var input string

	input := ""
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s \n", input)
	}
}
