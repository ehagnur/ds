package main

import (
	"fmt"
)

// Bubble sort is inplace
// time complexity O(n2)

func swap(a []int, idx1 int) {
	temp := a[idx1]
	a[idx1] = a[idx1+1]
	a[idx1+1] = temp
}

func main() {
	unsorted := [...]int{7, 2, 5, 1, 5, 0, 10, 11, 12}
	var toBubble int
	fmt.Printf("Initial unsorted list: %v\n", unsorted)
	for i := 0; i < len(unsorted)-1; i++ {
		flag := false
		for j := 0; j < len(unsorted)-i-1; j++ {
			if unsorted[j] > unsorted[j+1] {
				toBubble = unsorted[j]
				swap(unsorted[:], j)
				flag = true
			}
		}
		// return if array is already sorted
		if !flag {
			fmt.Printf("Array is already sorted, exiting after the %dth iteration\n", i)
			break
		}
		fmt.Printf("Array after bubbling %d %v\n", toBubble, unsorted)
	}
	fmt.Printf("Final sorted list, %v: \n", unsorted)
}

