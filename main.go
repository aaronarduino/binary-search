package main

import (
	"fmt"
	"time"
)

var arrs = [][]int{
	{-2, 3, 4, 7, 8, 9, 11, 13},
	{-20, -19, -18, -17, -16, -15, -14, -13, -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
}

var targets = []int{11, 9}

func main() {
	for i, arr := range arrs {
		target := targets[i]

		fmt.Println(fmt.Sprintf("Searching for %d:", target))

		lsResult := linearSearch(arr, target)
		if lsResult < 0 {
			fmt.Println("linearSearch could not find target in array.")
		} else {
			fmt.Printf("linearSearch found %d at index %d\n", target, lsResult)
		}

		bsResult := binarySearch(arr, target)
		if bsResult < 0 {
			fmt.Println("binarySearch could not find target in array.")
		} else {
			fmt.Printf("binarySearch %d found at index %d\n", target, bsResult)
		}
	}
}

func linearSearch(arr []int, target int) int {
	defer elapsed("linear search")()
	for i, number := range arr {
		if number == target {
			return i
		}
	}
	return -1
}

func binarySearch(arr []int, target int) int {
	defer elapsed("binary search")()

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
