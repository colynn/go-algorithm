package algorithm

// BubbleSort ..
// No.1
// 冒泡排序 （稳定）算法步骤：
// 1. 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
// 2. 对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
// 3. 针对所有的元素重复以上的步骤，除了最后一个。
// 4. 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需求比较。
func BubbleSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
