package main

import "fmt"

// Merge sort
func mergeSort(listItems []int) []int {
	if len(listItems) < 2 {
		return listItems
	}
	firstPart := mergeSort(listItems[:len(listItems)/2])
	secondPart := mergeSort(listItems[len(listItems)/2:])
	return merge(firstPart, secondPart)
}

// Merge algorithm
func merge(items1 []int, items2 []int) []int {
	finalItems := []int{}

	i := 0
	j := 0

	for i < len(items1) && j < len(items2) {
		if items1[i] < items2[j] {
			finalItems = append(finalItems, items1[i])
			i++
		} else {
			finalItems = append(finalItems, items2[j])
			j++
		}
	}
	for ; i < len(items1); i++ {
		finalItems = append(finalItems, items1[i])
	}
	for ; j < len(items2); j++ {
		finalItems = append(finalItems, items2[j])
	}
	return finalItems
}

func main() {
	fmt.Println("Merge sort working...")
	sampleList := []int{1, 25, 33, 99, 6, 15, 36, 5, 10, 22, 74, 66, 41}
	result := mergeSort(sampleList)
	fmt.Println(result)
}
