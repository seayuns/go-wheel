package sorts

//冒泡排序
//两轮循环，外循环为 数组冒泡区减少，内存循环为为冒泡比对
//这一步没有出现任何一次交换，那么说明所有的元素已经有序，不需要再进行下面的步骤了,加入flag

func BubbleSort(arrs []int) ([]int, int) {
	n := 0
	aLen := len(arrs)
	for i := aLen - 1; i >= 0; i-- {
		flag := 0
		for j := 0; j < i; j++ {
			if arrs[j] > arrs[j+1] {
				arrs[j], arrs[j+1] = arrs[j+1], arrs[j]
				n++
				flag = 1
			}
		}
		if flag == 1 {
			break
		}
	}

	return arrs, n
}
