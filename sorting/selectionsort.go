package main

import (
	"fmt"
)

// Selection sort is inplace
// time complexity O(n2)

func swap(a []int, idx1 int, idx2 int) {
	temp := a[idx1]
	a[idx1] = a[idx2]
	a[idx2] = temp
}

func main() {
	unsorted := []int{7, 2, 5, 1, 5, 0}
	fmt.Printf("Initial unsorted list: %v\n", unsorted)
	for i := 0; i < len(unsorted)-1; i++ {
		for j := i + 1; j < len(unsorted); j++ {
			if unsorted[i] > unsorted[j] {
				fmt.Printf("%d is greater than %d, ", unsorted[i], unsorted[j])
				swap(unsorted, i, j)
				fmt.Printf("List after swapping %v\n", unsorted)
			}
		}
	}
	fmt.Printf("Final sorted list, %v: \n", unsorted)
}

