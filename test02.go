package main

import (
	"fmt"
)

//map 学习 map 是一个key->vlaue 的数据类型
func main()  {
	//初始化一个map并赋值
	m := map[int]string{
		1:"lichao",
		2:"richard",

	}
	//遍历map
	for _ ,v :=range m {

		//fmt.Println("名字：",v)
		//fmt.Printf("学号：%d ---> 名字：%s\n",k,v)
		fmt.Println("名字",v)
	}
	//获取map的索引1对应的指
	name1 := m[1]
	fmt.Println(name1)
	//检查键指是不是存在
	if v ,ok := m[3] ;ok {
		fmt.Println(v)
	}else {
		fmt.Printf("2 no thing\n")
	}
    //增加map元素
    m[3]="wang"
    fmt.Println(m)
    //删除1的vlaue
    delete(m,1)
    fmt.Println(m)

}
