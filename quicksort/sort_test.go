package quicksort

import (
	"reflect"
	"testing"
)

func TestListFull(t *testing.T) {
	var arrayToSort = []int{55, 21, 3, 108, -1, 10, 0, 8, 4, 4}
	var expected = []int{-1, 0, 3, 4, 4, 8, 10, 21, 55, 108}
	sortedArray := sort(arrayToSort)

	if !reflect.DeepEqual(sortedArray, expected) {
		t.Errorf("Array is incorrect, got: %d, want: %d.", sortedArray, expected)
	}

}

func TestListOneElement(t *testing.T) {
	var arrayToSort = []int{1}
	var expected = []int{1}
	sortedArray := sort(arrayToSort)

	if !reflect.DeepEqual(sortedArray, expected) {
		t.Errorf("Array with one emenet is incorrect, got: %d, want: %d.", sortedArray, expected)
	}
}
