package main

import (
	"flag"
	"fmt"
)

func main() {
	// 命令行标志 是命令行程序指定选项的常用方式。例如，在 wc -l 中， 这个 -l 就是一个命令行标志。

	// 基本的标记声明仅支持字符串、整数和布尔值选项。
	//这里我们声明一个默认值为 "foo" 的字符串标志 word 并带有一个简短的描述。
	//这里的 flag.String 函数返回一个字符串指针（不是一个字符串值）， 在下面我们会看到是如何使用这个指针的。
	wordPtr := flag.String("word", "foo", "a string")

	// 使用和声明 word 标志相同的方法来声明 numb 和 fork 标志。
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// 用程序中已有的参数来声明一个标志也是可以的。 注意在标志声明函数中需要使用该参数的指针。
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 所有标志都声明完成以后，调用 flag.Parse() 来执行命令行解析。
	flag.Parse()

	// 这里我们将仅输出解析的选项以及后面的位置参数。 注意，我们需要使用类似 *wordPtr 这样的语法来对指针解引用， 从而得到选项真正的值。
	// 使用 -h 或者 --help 标志来得到自动生成的这个命令行程序的帮助文本。
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
