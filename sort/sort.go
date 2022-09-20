package sort

func QuickSort(array []int) {
	quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, left, right int) {
	if left < right {
		partitionIndex := partition(array, left, right)
		quickSort(array, left, partitionIndex-1)
		quickSort(array, partitionIndex+1, right)
	}
}

func partition(array []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if array[i] < array[pivot] {
			swap(array, i, index)
			index += 1
		}
	}
	swap(array, pivot, index-1)
	return index - 1
}

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func SelectSort(array []int) {
	for i := 0; i < len(array); i++ {
		index := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[index] {
				index = j
			}
		}
		swap(array, i, index)
	}
}

func BubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				swap(array, j, j+1)
			}
		}
	}
}
