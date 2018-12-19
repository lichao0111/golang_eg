package main

import "fmt"

//指针练习
func main ()  {
	b := 255
	var a *int =&b
	fmt.Println(a)    //a 的内存地址
	fmt.Println(*a)   //a 指向的是255
	//给a指向的指+1 也就是b+1
	*a++
	fmt.Println(*a)  // a 是256
	fmt.Println(b)   // b 是256




	}
