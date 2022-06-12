package main

import (
	"fmt"
)

func bubbleSort(arr []int, n int) {
	for i := range arr {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1] = arr[j] + arr[j+1]
				arr[j] = arr[j+1] - arr[j]
				arr[j+1] = arr[j+1] - arr[j]

			}
		}
	}
}

func main() {

	arr := []int{5, 8, 1, 9, 3, 1}
	size := len(arr)
	bubbleSort(arr, size)
	fmt.Print(arr)

}
