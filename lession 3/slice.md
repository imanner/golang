
内部实现：
	切片slice是动态的数组，其底层就是数组！
	数据结构有三个字段:指向数组的指针，切片的长度，切片的容量！

声明和初始化:
	最简洁的make方式:
		变量名 := make([]类型,长度【,容量】)
		例如：
			slice := make([]int, 5)	//声明长度和容量都是为5的切片
			slcie := make([]int, 5, 10) //声明长度是5，容量是10 的切片

		注意：长度必须小于等于容量！！！

	直接初始值的形式:
		方式一:给出所有的初始值
			变量名 := []类型{值列表}
		例如:
			slice := []int{1,2,3,4,5}

		方式二:只初始化部分索引的值
			变量名 := []类型{索引:值}
		例如:
			slice := []int{4:5}

		注意:注意比较切片和数组之间的细微差别
			//数组
			array:=[5]int{4:1}
			//切片
			slice:=[]int{4:1}

		方式三:nil切片和空切片
			//nil切片
			var nilSlice []int
			//空切片
			slice:=[]int{}

	基于现有的切片或者数组创建切片,特别像substr函数
		新切片变量 := 老切片变量[【起始下标】:【终止下标】【:新切片的容量】]
		注意：
			新切片包含起始位置不包含结束位置！
			起始位置可省略，默认为0
			终止位置可省略，默认到结尾！
			新老切片共用一个底层数组！他们是引用关系！

		例子：
			slice := []int{1, 2, 3, 4, 5}
			slice1 := slice[:]
			slice2 := slice[0:]
			slice3 := slice[:5]

			fmt.Println(slice1)		//1 2 3 4 5
			fmt.Println(slice2)		//1 2 3 4 5
			fmt.Println(slice3)		//1 2 3 4 5

		新切片的长度和容量
			方式一:计算，对于底层数组容量是k的切片slice[i:j]来说
				长度：j-i
				容量:k-i
			方式二:go系统函数
				slice := []int{1, 2, 3, 4, 5}
				newSlice := slice[1:3]
				fmt.Printf("newSlice长度:%d,容量:%d",len(newSlice),cap(newSlice))
