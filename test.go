package main

import "fmt"

func main() {
	//fmt.Println("hello world")
	//slice2 := make([]float32, 3, 5) // [0 0 0] 长度为3容量为5的切片
	////slice2 = append(slice2, 1, 2, 3, 4,5,6,7,8,9) // [0, 0, 0, 1, 2, 3, 4]
	//fmt.Println(len(slice2), cap(slice2),slice2) // 7 12
	str := "Golang"
	var p *string = &str // p 是指向 str 的指针
	*p = "Hello"
	fmt.Println(str) // Hello 修改了 p，str 的值也发生了改变
	//test github
	fmt.Println(str) // Hello 修改了 p，str 的值也发生了改变
}
