// array.go
package main

import (
	"fmt"
)

/**
经过今天的学习学习总结以下指示
在Go中，变量的声明需要明确类型！赋值也需要声明的类型
非指针类型数组格式:
	第一种形式
	var 变量名 [长度]类型 = [长度]类型{索引:默认值, 索引:默认值}
	第二种形式
	var 变量名 [长度]类型
	变量名 = [长度]类型{索引:默认值, 索引:默认值}
	第三种形式
	变量名 := [长度]类型{索引:默认值, 索引:默认值}
指针类型数组格式:
	第一种形式
		变量名 = [长度]*类型{索引:new(类型), 索引:new(类型)}

数组的循环输出
	for i,v := range 数组变量{
		fmt.Printf("索引是:%d, 数值是:%d\n", i, v)
	}

函数：
	func 函数名(变量名 [长度]类型){
		变量名[下标]
		fmt.Println(变量名)
	}

	func 函数名(变量名 *[长度]类型){
		变量名[下标]
		fmt.Println(*变量名)
	}
*/
func main() {
	//定义数组
	var array [5]int
	// 没有初始化  都是0值
	fmt.Println(array[0])

	array = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array[0])

	//在函数内部 可以简洁定义新的变量
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2[0])

	array3 := [...]int{1, 2, 3}
	fmt.Println(array3[0])

	// 设置默认值  {索引:默认值, 索引:默认值}
	array4 := [...]int{0: 2, 3: 4}
	fmt.Println(array4[0])

	//修改数据
	array4[0] = 6
	fmt.Printf("%d\n", array4[0])

	//使用for循环输出前4个元素  这种方式有问题：如果没有设置下标的话 会出错
	for i := 0; i < 4; i++ {
		fmt.Printf("array4索引是:%d, 数值是:%d\n", i, array4[i])
	}

	//go的特殊循环方式  这种方式太棒了  不用考虑下标的数量
	for i, v := range array4 {
		fmt.Printf("array4索引是:%d, 数值是:%d\n", i, v)
	}

	array5 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("array5数值是:%d\n", array5[0])
	var array6 [5]int = array5
	fmt.Printf("array6数值是:%d\n", array6[0])
	// 不同类型【长度或者值的类型】的数组不能相互赋值
	//var array7 [4]int = array5

	//指针数组  分配内存
	array7 := [5]*int{0: new(int), 3: new(int)}
	*array7[0] = 123
	fmt.Printf("指针数组值:%d\n", *array7[0])
	//分配内存  之后才能赋值
	array7[1] = new(int)
	*array7[1] = 456
	fmt.Printf("指针数组值:%d\n", *array7[1])

	//调用函数
	modify(array6)
	fmt.Println(array6)

	//调用指针函数
	modify2(&array6)
	fmt.Println(array6)

}

//函数间传递数组
func modify(a [5]int) {
	a[1] = 123567
	//fmt.Printf("数组的值是:%d\n", a[1])
	fmt.Println(a)
}

//如果要更改原来的数据  可以通过指针来做
func modify2(a *[5]int) {
	//*a[1] = new(int)
	a[1] = 5765
	fmt.Println(*a)
}
