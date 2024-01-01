package main

import "fmt"

// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}

// 函数值不仅仅是一串代码，还记录了状态。
// 我们看到变量的生命周期不由它的作用域决定：squares返回后，变量x仍然隐式的存在于f中。
