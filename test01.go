package main

import "fmt"

//切片学习
func main ()  {
	s := make([]string,3)     //初始化一个切片
	t := [] int{1,2,3,4,5,9}  //初始化切片并赋值
	fmt.Println(s)
	s[0]="a"
	s[1]="b"
	s[2]="c"
	//s[3]="d"
	fmt.Println(s)
	fmt.Println(s[1])
	fmt.Println(len(s))
	s = append(s, "e")  //切片增加元素
	fmt.Println(s)
	fmt.Println(len(s))
	c := make([]string,len(s))
	copy(c,s)                   //切片copy
	fmt.Println("c:",c)

	fmt.Println(c[2:4])
	fmt.Println(t)
	t = append(t, 10)
	fmt.Println(t)
}