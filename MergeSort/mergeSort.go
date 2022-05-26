package main

import "fmt"

// main is the entry point for the application
func main() {
	arr := []int{5, 4, 3, 2, 1}
	fmt.Println("Unsorted array:", arr)
	fmt.Println("Sorted array:", mergeSort(arr))
}

// mergeSort sorts the input array in ascending order
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

// merge merges two sorted arrays into one sorted array
func merge(left, right []int) []int {
	var result []int

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}
