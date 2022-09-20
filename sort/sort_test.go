package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr := []int{1, 6, 7, -2, 8, 5, 10}
	QuickSort(arr)
	assert.Equal(t, []int{-2, 1, 5, 6, 7, 8, 10}, arr, "except get {-2, 1, 5,6,7,8} but got ", arr)
}

func TestSelectSort(t *testing.T) {
	arr := []int{1, 6, 7, -2, 8, 5,10}
	SelectSort(arr)
	assert.Equal(t, []int{-2, 1, 5, 6, 7, 8, 10}, arr, "except get {-2, 1, 5,6,7,8} but got ", arr)
}

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 6, 7, -2, 8, 5,10}
	BubbleSort(arr)
	assert.Equal(t, []int{-2, 1, 5, 6, 7, 8, 10}, arr, "except get {-2, 1, 5,6,7,8, 10} but got ", arr)
}
