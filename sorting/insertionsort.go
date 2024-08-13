package main

import (
	"fmt"
)

func main() {
	unsorted := [...]int{7, 2, 5, 1, 5, 0}
	fmt.Printf("Initial unsorted list: %v\n", unsorted)
	for i := 1; i < len(unsorted); i++ {
		value := unsorted[i]
		hole := i
		for j := i; j > 0; j-- {
			if unsorted[j-1] > value {
				unsorted[j] = unsorted[j-1]
				hole = j - 1
			}
		}
		unsorted[hole] = value
		fmt.Printf("Iteration %d, inserting %d: Array: %v\n", i, value, unsorted)
	}
}

