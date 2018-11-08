package main // package定义  main主体包 每个程序必有main主体包

import (
	"errors"
	"fmt"
	"math"
) //import导入  fmt等各种库 或 其它包

/*
import fm "fmt"  使用包 别名 ，避免命名冲突
多行注释

fmt.Printf("字符串拼接符 +")
*/


var v int = 5 /* var变量类型声明 变量名v 为 int整数5
   在所有init或main 模具前的初始化声明，将是全局通用   */

func init() { /* 如果存在 init 模具，执行优先级 init模具，再 按照 程序主体模具main 中先后调用的顺序来执行相关模具
	   模具命名第一个字符为大写 字母，能被其它包调用，否则为本包私有*/
}

func main() { //func定义模具  每个程序必有  main主体模具
	模具一一数字运算()

}

func 模具一一模板() {
	fmt.Printf(" ") /* 

				*/
	fmt.Printf(" ") /* 

				*/
	fmt.Printf(" ") /* 
	
					*/
}

func 模具一一标准库一一安装与使用() {
	fmt.Println("新建一个项目") /*
		项目目录路径以 src(代码源)目录下 项目名目录 结尾.
		新建项目的目录 注册在 GOPATH环境变量里面，GOPATH可以用; 分好定义多个环境位置。
		所有新建的 自定义 库（包）都 新建同包名目录在 src 下，同包名目录 性能 同包名的.go文件
	*/
	fmt.Println("安装项目与安装库（包）") /*

		库名.go 文件安装前要 确定保存
		DOS界面 输入
		go build 库名   //有 main主体包 时 使用
		go install 库名
		成功之后在 项目的目录下，src(代码源)同级 新建 pkg(已经编译文件)下 ，再新建 本机系统类型的目录 ，创造 库名.a 文件
		到此，安装完成
	*/
	fmt.Println("库（包）的 运用") /*
		库名.go 运用的模具 命名首个字母 为大写
		导入库名 import "./库名" 例如 "./ha6"
		运用库  库名.Q模具一一运算() 例如ha6.Q模具一一运算()
	*/
	fmt.Println("下载 封禁的golang 社区 库") /*
		下载地址从 golang社区 改变为 https://github.com/golang/库名
		在 代码源 src目录 下新建 golang.org目录 ，再建 x目录，然后 DOS 通道此 目录
		输入 并下载 https://github.com/golang/库名
		安装失败的 库 ，在 github 下载安装包，安放在 src/github.com目录下，go install 库名 安装
	*/

}


func 模具一一定义模具示例(传入参数 int, 参数2 int) (返回参数 int, 参数2 int) { //int为类型 简写(传入参数,参数2 传入类型)
	fmt.fmt.Printfln("hello",a)// 打印
	
}

func 模具一一类型一一常量() {
	const Pi = 3.14159     //常量使用关键字 const 定义，用于存储不会改变的数据。
	const b string = "abc" //显式类型定义
	const b = "abc"        //隐式类型定义： 省略类型说明符 string[type类型] 简写
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	) // 声明多个  常量

}

func 模具一一类型一一变量() {
	数字a := 5.0    // 简写 声明变量 数字a 并 附值 首次声明变量名才能用 := ，之后 附值用 =
	数字B := int(a) // 简写 声明变量 数字B 并 类型转换
	
	var a int = 15 //变量使用关键字 var 定义
	var a = 15 // 简写
	var (
	a = 15
	b = false
	str = "Go"
	city string
)//声明多个  变量
	var a, b, c int //同一类型的多个变量可以声明在同一行
	a, b, c := 5, 7, "abc" //多变量可以在同一行进行赋值
	
	
}

func 模具一一类型一一运算符() {
	关系运算符 := 5    // 非 !(!T )、并 &&(T && T)、或 ||(T || T)
	逻辑运算符:= 5    // ==、!= 、<、<=、>、>=
	算术运算符:= 5    /* +、-、* 和 /、%
					/ 对于整数运算而言，结果依旧为整数，例如：9 / 4 -> 2。
					取余运算符只能作用于整数：9 % 4 -> 1。
					*/

}

func 模具一一类型一一变量格式化() {
	
	fmt.Printf("数字 %d",5)// %d 用于格式化整数，%g 用于格式化浮点型 %数字.mg 用于表示格式化 ‘数字’ 并精确到小数点后m 位
	文本字符串 := "字符" /*   
	%s用于格式化 %c字符；当和字符配合使用时，%U 输出格式为 U+hhhh 的字符串（另一个示例见第 5.4.4 节）。
    判断是否为字母：unicode.IsLetter(文本字符串)
    判断是否为数字：unicode.IsDigit(文本字符串)
    判断是否为空白符号：unicode.IsSpace(文本字符串)
*/

}

func 模具一一类型一一字符串() {	
	fmt.Printf("解释字符串") /* 
				\n：换行符
				\r：回车符
				\t：tab 键
				\u 或 \U：Unicode 字符
				\\：反斜杠自身
*/
	fmt.Printf("字节长度"，len(str)) /* 
				//len() 来获取字符串所占的字节长度
			*/
	fmt.Printf("字符串的内容（纯字节）可以通过标准索引法来获取") /*
			    字符串 str 的第 1 个字节：str[0]
			    第 i 个字节：str[i - 1]
			    最后 1 个字节：str[len(str)-1]
				*/
				
	fmt.Printf("字符串拼接符 +") /*
				两个字符串 s1 和 s2 可以通过 s := s1 + s2 拼接在一起。
				*/

}

func 模具一一文本字符串清洗() {// strings 包 判断字符串

	原文本 ：="a判断a"
	规则字符="a"
	strings.Replace(原文本, 规则字符, 新字符, 替换次数) string // 替换次数= -1 则替换所有
	strings.TrimSpace(原文本)/* 剔除字符串开头和结尾的空白（空格）符号
				strings.Trim(s, "剔除字符") 来将开头和结尾的 字符 去除掉
					剔除开头的字符串 TrimLeft 
				剔除结尾的字符串 TrimRight 

*/
	
	文本列表=strings.Split(原文本, 分割符号)// 以 分割符号 分割成列表


	strings.HasPrefix(原文本, 规则字符 string) bool//判断 原文本 是否以 规则字符串 开头 前缀
	strings.HasSuffix(原文本, 规则字符 string) bool//判断 结尾 后缀
	strings.Contains(原文本, 规则字符 string) bool//判断 原文本 是否包含 规则字符串
	
	strings.Index(原文本, 规则字符 string) int /*
								判断 规则字符  在 原文本出现的位置（第一个字符 索引）					
					-1 表示 不包含 规则字符*/
	strings.LastIndex(原文本, 规则字符 string) int // 判断 最后出现位置的索引，-1 表示 不包含 字符串
	strings.IndexRune(原文本 string, 规则字符 rune) int// 查询非 ASCII 编码的字符在 原文本 的位置	
	strings.Count(原文本, 规则字符 string) int//规则字符串 出现的非重叠次数
	重复3次原文本 = strings.Repeat(原文本, 3)//重复 3 次 原文本 并返回一个新的字符串
	strings.ToLower(原文本) string //全部转换为相应的小写字符
	strings.ToUpper(原文本) string//全部转换为相应的大写字符
	文本列表=strings.Fields(原文本)// 以空白符号 分割成列表
	列表转文本 := strings.Join(原列表,"合并字符") // 拼接 列表 为字符串
	
	整数转字符串:=strconv.Itoa(数字 int) string//返回数字 i 所表示的字符串类型的十进制数
	浮点型数字转换为字符串:=strconv.FormatFloat(浮点型数字 float64, fmt byte, 精度位数 int, 浮点类型 int) string
					 /* 其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'），
					prec 表示精度位数，
					浮点类型bitSize 则使用 32 表示 float32，用 64 表示 float64。*/

	字符串转换为整数,可能错误= strconv.Atoi(原文本)/*
		模具strconv.Atoi(原文本 string) (整数 int, err error)
						会返回 2 个值，第 1 个是转换后的结果（如果转换成功），
						第 2 个是可能出现的错误
					将字符串转换为其它类型 tp 并不总是可能的，
					可能会在运行时抛出错误 parsing "…": invalid argument。*/
	将字符串转换为float64型,可能错误:=strconv.ParseFloat(原文本)/*
	strconv.ParseFloat(原文本 string, 浮点类型 int) (f float64, err error)
	浮点类型bitSize 则使用 32 表示 float32，用 64 表示 float64。
	*/



	
	fmt.Printf(" ") /* 
	
					*/
	fmt.Printf(" ") /* 
	
					*/
}

func 模具一一异常判断示例() {
	
	
	原文本 := "666"
	fmt.Println("Hello, %s!\n", 原文本)
	字符串转换为整数, 异常原因 := strconv.Atoi(原文本)
	if 异常原因 != nil { //nil 零 空
		fmt.Println("异常原因 : ", 异常原因)
		return 异常原因//返回
	}
	fmt.Println("%d\n", 字符串转换为整数)
}

func 模具一一时间() {
	时间 := time.Now()//获取当前时间
	fmt.Println("%s\n", 时间)
	fmt.Println("%02d.%02d.%4d\n", 时间.Day(), 时间.Month(), 时间.Year())

}

func 模具一一条件控制() {

	逻辑型条件 := 0
	if 逻辑型条件 == 0 { // 布尔型或逻辑型 的语句，条件成立
				//  break//结束循环、continue//跳过本次循环、goto//跳转到 或者 return//返回 
	   			// 非 !(!T )、并 &&(T && T)、或 ||(T || T)
				// ==、!= 、<、<=、>、>=
		fmt.fmt.fmt.Printfff("相等\n")
	} else if 逻辑型条件 == 2 { //其它条件  !(逻辑型条件 == 2) 相反结果 
		fmt.Println("其它条件相等\n")

	} else { // 否则
		fmt.Println("不相等\n")
	}

	布尔型对否 := true
	if 布尔型对否 { // if !布尔型对否 !(布尔型对否) 相反结果
		fmt.Println("对\n")
	} else {
		fmt.Println("否\n")
	}

}

func 模具一一多项条件判断() {
	比较数值 := 98

	switch 比较数值 { //多项条件
	case 98: //条件一
		fmt.fmt.fmt.Printffln("大于 98")
	case 99, 100: //条件二
		fmt.fmt.Printfln("等于 100")
	default: //否则
		fmt.fmt.Printfln("不存在于 98 或 100")
	}

}
func 模具一一for计数器迭代() {// break//结束循环、continue//跳过本次循环、goto//跳转到 或者 return//返回 
	

	范围 := "256415遍历"
	for 计数器 := 0; 计数器 < len(范围); 计数器 = 计数器 + 1 {

		fmt.fmt.fmt.Printfff("遍历%d 字符串%c\n", 计数器, 范围[计数器])

	}
	for i := 0; i < 5; i++ {
		fmt.Println("遍历数 %d \n", i)
	}
}

func 模具一一模具一命名返回值(运算初数 int) (x2, x3 int) {//x2, x3 := 模具一一运算(运算初数)
	defer 模具一untrace("a")//在return返回 后， 最终 执行的 命令
	x2 = 2 * 运算初数
	x3 = 3 * 运算初数
	// return x2, x3
	return // 返回
	
}


func 模具一一模具一传递指针() {//传递指针给模具不但可以节省内存（因为没有复制变量的值），而且赋予了函数直接修改外部变量的能力，所以被修改的变量不再需要使用 return 返回

	初数 := 0
	初数的指针 := &初数 // 赋予指针
	模具一运算(10, 5, 初数的指针)
	fmt.fmt.Printfln("两数相乘:", *初数的指针) // *指针的变量

	func 模具一运算(值一, 值二 int, 初数的指针 *int) {
	*初数的指针 = 值一 * 值二
}

}

func 模具一一模具一内置函数() {

	fmt.fmt.Printfln("close")             // 用于管道通信
	fmt.fmt.Printfln("len")               //len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）
	fmt.fmt.Printfln("cap")               // cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
	fmt.fmt.Printfln("copy")              // 用于复制
	fmt.fmt.Printfln("append")            // 连接切片 ，列表追加
	fmt.fmt.Printfln("panic、recover")     // 两者均用于错误处理机制
	fmt.fmt.Printfln("fmt.Printf、fmt.Printfln")     // 底层打印函数（详见第 4.2 节），在部署环境中建议使用 fmt 包
	fmt.fmt.Printfln("complex、real imag") // 用于创建和操作复数（详见第 4.5.2.2 节）

}

func 模具一一模具一计算函数执行时间() {
	开始时间 := time.Now()
	//longCalculation()
	结束时间 := time.Now()
	执行时间 := 结束时间.Sub(开始时间)
	fmt.Println("计算模具 执行时间: %s\n", 执行时间)

}

func 模具一一数组循环() {
	fmt.fmt.Printfln("for 循环列表") // range范围 生成方式 for i,_:= range 列表 {
	列表 := [...]string{"a", "b", "c", "d"}
	for 序号 := range 列表 {
		fmt.fmt.Printfln("列表索引", 序号, "列表值为", 列表[序号])
	}

}

func 模具一一数组生成() {

	数组 := [...]string{"a", "b", "c", "d"}
	fmt.fmt.Printfln("数组", 数组) //
	数组二 := [10]int{18, 20, 15, 22, 16}
	fmt.fmt.Printfln("数组二", 列表二) //
	数组三 := [5]string{3: "3Chris", 4: "4Ron"}
	fmt.fmt.Printfln("数组三", 数组三) //

}

func 模具一一传递数组的指针() {

	浮点数数组 := [3]float64{7.0, 8.5, 9.1}
	相加数 := 模具一一数组相加(&浮点数数组) // 传递列表指针

	fmt.Println("列表数相加为: %f \n", 相加数)

	func 模具一一数组相加(a *[3]float64) (相加数 float64) {
		for _, 列表值 := range a { // for 循环 列表 范围
			相加数 = 相加数 + 数组值
		}
		return
}

}

func 模具一一列表生成() {

	列表 := []int{1, 2, 3}
	fmt.fmt.Printfln("列表", 列表)    //
	列表二 := make([]int, 0)    //列表:= make([]type类型, len列表 初始长度)
	fmt.fmt.Printfln("列表二", 列表二)  // make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel（参见第 8 章，第 13 章）。
	列表三 := new([cap]int)[0:5] //cap整数 是可选参数，数组的长度。new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体（参见第 10 章）；它相当于 &T{}
	fmt.fmt.Printfln("列表三", 列表三)  //
	列表 = 列表[0:len(列表)+1]// 列表 初始长度 扩展 1 位

}

func 模具一一列表循环() {

	列表 := []string{"Sp", "Su", "Au", "Win"}
	for 索引, 列表值 := range 列表 {// for循环range范围 列表
		fmt.fmt.fmt.Printfff("索引 %d 值: %s\n", 索引, 列表值)
	}

	//var 列表值 string
	for _, 列表值 := range 列表 { //_ 可以用于忽略索引
		fmt.Println("%s\n", 列表值)
	}
}

func 模具一一列表复制与追加() {

	复制的列表 := []int{1, 2, 3, 5}
	主体列表 := make([]int, 10)

	复制列表个数 := copy(主体列表, 复制的列表) // 合并复制
	fmt.fmt.Printfln("列表二 %c ", 主体列表)
	fmt.Println("复制列表个数 %d \n", 复制列表个数) // n == 3

	列表三 := []int{1, 2, 3}
	列表三 = append(列表三, 4, 5, 6)   //追加  追加的元素必须和原切片的元素同类型
	fmt.Println("列表三 %d \n", 列表三) //
}

func 模具一一for循环遍历2字节位文本() {

	文本 := "追加的元素必须和原切片的元素同类"
	//fmt.Println("%s \n", 文本)
	字节位文本列表 := []int32(文本) // 类似的， := []rune(文本) 每个 int 都会包含对应的 Unicode 代码
	for 索引, 值 := range 字节位文本列表 {
		fmt.Println("%d:%c ", 索引, 值)

	}
	fmt.Println("\n ")
	for _, 值 := range 字节位文本列表 { //  _, 忽略

		fmt.Println("%c", 值) //%s 格式化的文本为1字节位，%c 格式化2字节位
	}
	fmt.Println("\n ")
}

func 模具一一修改字符串中的某个字符() {

	原文本 := "hello"
	文本列表 := []byte(原文本)         //先将字符串转换成字节数组，
	文本列表[0] = 'c'               // 然后再通过修改数组中的元素值来达到修改字符串的目的，
	修改后文本 := string(文本列表)       // 最后将字节数组转换回字符串格式。
	fmt.fmt.Printfln("修改后文本", 修改后文本) // 修改后 ' cello'
}


func 模具一一append列表追加常见操作() {

	列表 = append(列表[:i], 列表[i+1:])               // 删除位于索引 i 的元素：
	列表 = append(列表[:i], 列表[j:])                 //切除切片 a 中从索引 i 至 j 位置的元素：
	列表 = append(列表, make([]T, j))               // 为 列表 扩展 j 个元素长度：
	列表 = append(列表[:i], append([]T{x}, 列表[i:])) //在索引 i 的位置插入元素 x：
	列表 = append(列表[:i], append(b, 列表[i:]))      // 在索引 i 的位置插入切片 b 的所有元素：
	x = 列表[:len(列表)-1]                          // 取出位于 列表 最末尾的元素 x：
	x, 列表 = 列表[len(列表)-1], 列表[:len(列表)-1]       // 取出位于 列表 最末尾的元素 x：
	列表 = append(列表, x)                          //将元素 x 追加到 列表：
	//
}

func 模具一一字典生成() {

	字典 := map[string]int{} //简写 创建字典
	//字典 = map[string]string{}
	字典["to"] = 20

	fmt.Println("字典 \"to\"键 is: %s\n", 字典["to"]) // 无 字典键 附值为 0
	mp1 := make(map[int][]int)                  //列表作为 字典 的值
	mp2 := make(map[int]*[]int)                 //列表作为 字典 的值
	列表 := make([]map[int]int, 5)// 字典 类型的列表
}

func 模具一一判断字典键值对是否存在() {
	var 值 int
	var 对否 bool

	字典 := make(map[string]int)
	字典["键1"] = 55
	字典["键2"] = 20
	字典["键3"] = 25
	值, 对否 = 字典["键2"]//  返回布尔型
	if 对否 {//布尔型
		fmt.fmt.fmt.Printfff("值 \"键2\" 存在于 字典: %d\n", 值)
	} else {
		fmt.Println("字典 不存在 键2")
	}

	值, 对否 = 字典["Paris"]
	fmt.Println(" \"Paris\" 存在于 字典: %t\n", 对否)
	fmt.Println("值 为: %d\n", 值) //空值为0

	// delete an item:
	delete(字典, "键3")//删除 字典的一个键
	值, 对否 = 字典["键3"]
	if 对否 {
		fmt.Println(" 值\"键3\" 存在于 字典: %d\n", 值)
	} else {
		fmt.fmt.Printfln("字典 不存在 键3")
	}
}

func 模具一一for循环遍历字典() {
	点数字典 := make(map[int]float32)
	点数字典[1] = 1.0
	点数字典[2] = 2.0
	点数字典[3] = 3.0
	点数字典[4] = 4.0
	for 字典键, 值 := range 点数字典 {
		fmt.Printf("字典键: %d - 值: %f\n", 字典键, 值)
	}
	文本字典 := map[string]string{"键1": "Pa", "键2": "Ro", "键3": "To"}
	for 键 := range 文本字典 {
		fmt.Println("文本字典: ", 键, "值为", 文本字典[键])
	}
}

func 模具一一标准库概述() {
	fmt.Println("运行") /*完整列表可以在 Go Walker https://gowalker.org/search?q=gorepos

	  unsafe: 包含了一些打破 Go 语言“类型安全”的命令，一般的程序中不会被使用，可用在 C/C++ 程序的调用中。
	  syscall-os-os/exec:
	      os: 提供给我们一个平台无关性的操作系统功能接口，采用类Unix设计，隐藏了不同操作系统间差异，让不同的文件系统和操作系统对象表现一致。
	      os/exec: 提供我们运行外部操作系统命令和程序的方式。
	      syscall: 底层的外部包，提供了操作系统底层调用的基本接口。

	  archive/tar 和 /zip-compress：压缩(解压缩)文件功能。
	  fmt-io-bufio-path/filepath-flag:
	      fmt: 提供了格式化输入输出功能。
	      io: 提供了基本输入输出功能，大多数是围绕系统功能的封装。
	      bufio: 缓冲输入输出功能的封装。
	      path/filepath: 用来操作在当前系统中的目标文件名路径。
	      flag: 对命令行参数的操作。
	  strings-strconv-unicode-regexp-bytes:
	      strings: 提供对字符串的操作。
	      strconv: 提供将字符串转换为基础类型的功能。
	      unicode: 为 unicode 型的字符串提供特殊的功能。
	      regexp: 正则表达式功能。
	      bytes: 提供对字符型分片的操作。
	      index/suffixarray: 子字符串快速查询。
	  math-math/cmath-math/big-math/rand-sort:
	      math: 基本的数学函数。
	      math/cmath: 对复数的操作。
	      math/rand: 伪随机数生成。
	      sort: 为数组排序和自定义集合。
	      math/big: 大数的实现和计算。
	  container-/list-ring-heap: 实现对集合的操作。
	      list: 双链表。
	      ring: 环形链表。

	  time-log:
	      time: 日期和时间的基本操作。
	      log: 记录程序运行时产生的日志,我们将在后面的章节使用它。
	  encoding/json-encoding/xml-text/template:
	      encoding/json: 读取并解码和写入并编码 JSON 数据。
	      encoding/xml:简单的 XML1.0 解析器,有关 JSON 和 XML 的实例请查阅第 12.9/10 章节。
	      text/template:生成像 HTML 一样的数据与文本混合的数据驱动模板（参见第 15.7 节）。
	  net-net/http-html:（参见第 15 章）
	      net: 网络数据的基本操作。
	      http: 提供了一个可扩展的 HTTP 服务器和基础客户端，解析 HTTP 请求和回复。
	      html: HTML5 解析器。
	  runtime: Go 程序运行时的交互操作，例如垃圾回收和协程创建。
	  reflect: 实现通过程序运行时反射，让程序操作任意类型的变量。


	*/
}

func 模具一一标准库一一正则表达式regexp() {

	原文本 := "i LoVe You , xiaorui.cc 13684390231 ,i like tornado lang, golang 正则 . shell zsh golang bash c+ java jsp php"
	//替换 或 删除所有
	re, _ := regexp.Compile(".{0,}?xi") //变量声明
	删除后文本 := re.ReplaceAllString(原文本, "xi")
	fmt.Println("替换删除所有", 删除后文本)
	//转换成小写
	re, _ = regexp.Compile("\\w{0,}") // _忽略错误
	原文本 = re.ReplaceAllStringFunc(原文本, strings.ToLower)
	fmt.Println("转换成小写", 原文本)

	//替换所有数字
	re, _ = regexp.Compile("[0-9]{0,}")
	替换所有数字 := re.ReplaceAllString(原文本, "")
	fmt.Println("替换删除所有数字", 替换所有数字)

	
	
	//后面的提取 太复杂
	
	
	
	//判断是否是手机号码
	//fmt.Println(IsPhone("13684390231"))
	//b := []string{}
	re, _ = regexp.Compile("[0-9]{0,6}")
	//提取文本 := re.FindIndex([]byte(原文本))
	提取文本列表 := re.FindSubmatch([]byte(原文本))

	fmt.Println("判断是否是手机号码", "FindIndex", 提取文本列表)

	one1, _ := regexp.Compile("正则")
	index1 := one1.FindIndex([]byte(原文本))

	fmt.Println("手机号码", "FindIndex", index1)

	//第一个元素是匹配的全部元素，第二个元素是第一个()里面的
	re2, _ := regexp.Compile("zsh\\s(\\w*)\\sbash")
	submatch := re2.FindSubmatch([]byte(原文本))
	for _, v := range submatch {
		fmt.Println("第一个元素是匹配的全部元素", string(v))
	}

	// fmt.Println(strings.TrimSpace(原文本))
}

func 模具一一读取用户的输入() {
	等待输入 := bufio.NewReader(os.Stdin)
	fmt.Println("请等待输入内容 ，回车 确认键结束: ")
	input, 异常原因 := 等待输入.ReadString('\n') // 输入以  \n 回车 确认键结束
	if 异常原因 != nil {
		fmt.Printf("异常原因: %s\n", 异常原因)

	}
	fmt.Printf("输入的内容: %s\n", input)
}

func 模具一一每行读取文件() {
	文件路径 := "F:\\影视发帖推广\\可以注册手机号码.txt"

	文件缓存, 读取文件异常原因 := os.Open(文件路径)
	if 读取文件异常原因 != nil { // != nil 非空
		fmt.Printf("打开输入文件时发生错误\n" +
			"文件是否存在？\n" +
			"你有没有权限访问它?\n")
		return // exit the function on error
	}
	defer 文件缓存.Close()

	文件内容 := bufio.NewReader(文件缓存)

	for {
		每行内容, 读取异常原因 := 文件内容.ReadString('\n') //Windows的行结束符是 \r\n。在使用 ReadString 和 ReadBytes ，直接使用 \n 就可以了
		fmt.Printf("每行内容: %s", 每行内容)
		if 读取异常原因 == io.EOF { // true
			fmt.Printf("\n")
			return
		}
	}

}

func 模具一一读取文件一全部读取() {
	文件路径 := "F:\\影视发帖推广\\可以注册手机号码.txt"
	文件路径二 := "F:\\影视发帖推广\\可以注册手机号码.txt"
	文件内容, 读取文件异常原因 := ioutil.ReadFile(文件路径)
	if 读取文件异常原因 != nil {
		fmt.Fprintf(os.Stderr, "读取文件异常原因: %s\n", 读取文件异常原因)
		// panic(读取文件异常原因.读取文件异常原因or())
	}
	fmt.Printf("文件内容 ：\n%s\n", string(文件内容))
	读取文件异常原因 = ioutil.WriteFile(文件路径二, 文件内容, 0644) // oct, not hex
	if 读取文件异常原因 != nil {
		panic(读取文件异常原因.Error())
	}

}

func 模具一一文件写入() {
	/*导入库
	"bufio"
	"fmt"
	"os"
	*/
	文件路径 := "F:\\影视发帖推广\\临时文本.txt"
	文件缓存, 文件写入异常原因 := os.OpenFile(文件路径, os.O_WRONLY|os.O_CREATE, 0666)
	/*常会用到以下标志：
	      os.O_RDONLY：只读
	      os.O_WRONLY：只写（部分修改）
	      os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
	      os.O_TRUNC：（全部重写）截断：如果指定文件已存在，就将该文件的长度截为0。
	  写文件时，不管是 Unix 还是 Windows，文件的权限都需要使用 0666。
	*/
	if 文件写入异常原因 != nil {
		fmt.Printf("文件打开或创建时出错\n")
		return //返回
	}
	defer 文件缓存.Close() //缓存关闭

	文件写入 := bufio.NewWriter(文件缓存)
	写入内容 := "开或\r\n" // WIN系统 换行为 \r\n

	for i := 0; i < 5; i++ {
		文件写入.WriteString(写入内容)
	}
	文件写入.Flush() // 文件写入完成
	fmt.Printf("文件写入完成\n")

}

func 模具一一文件复制(目标文件, 源文件 string) (written int64, err error) {
	/*导入库
	"fmt"
	"io"
	"os"

目标文件:="F:\\影视发帖推广\\临时文本三.txt"
	源文件:="F:\\影视发帖推广\\临时文本.txt"
	模具一一文件复制(目标文件, 源文件)
	fmt.Println("文件复制 完成!")


	*/
	 文件缓存, 错误数 := os.Open(源文件) //打开文件
    if 错误数 != nil {// 非零
        return
    }
    defer 文件缓存.Close()
    文件写入, 错误数 := os.OpenFile(目标文件, os.O_WRONLY|os.O_CREATE, 0644) //写入文件
    if 错误数 != nil {// 非零
        return
    }
    defer 文件写入.Close()//关闭
    return io.Copy(文件写入, 文件缓存) //核心在这里，io.Copy 一行代码搞定



}

func 模具一一协程例子() {
	
	
	fmt.Println("启动 协程()")
	go 模具一一协程例子()
	go shortWait()
	fmt.Println("等待()")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(5e9)
	fmt.Println("协程结束()")
	
	
	
	/*导入库
	"fmt"
	"time"  */
	fmt.Println("第一例子1()")
	time.Sleep(2e9) // sleep for 5 seconds
	fmt.Println("第一例子2()")
}

func 模具一一协程例子的通道构建() {

	通道 := make(chan string)//简写
	
	var 通道1 chan string
	通道1 = make(chan string)


}