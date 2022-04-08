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
	defer elapsed("search")()
	for i, num := range arr {
		if num == target {
			return i
		}
	}
	return -1
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
