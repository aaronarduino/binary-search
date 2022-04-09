package main

import (
	"fmt"
	"time"
)

var arr = []int{-2, 3, 4, 7, 8, 9, 11, 13}

func main() {
	target := 11

	result := search(arr, target)

	fmt.Println(fmt.Sprintf("Searching for %d:", target))
	if result < 0 {
		fmt.Println("Could not find target in array.")
	} else {
		fmt.Printf("%d found at index %d\n", target, result)
	}
}

func search(arr []int, target int) int {
	defer elapsed("search binary")()

	l := 0
	r := len(arr) - 1

	for {
		i := r - l/2
		if target < arr[i] {
			r = i - 1
		} else if target > arr[i] {
			l = i + 1
		} else {
			return i
		}
	}

	// Old code
	//for i, num := range arr {
	//	if num == target {
	//		return i
	//	}
	//}

	// New Code
	return -1
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
