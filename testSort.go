package main

import (
	"fmt"
	"math/rand"
)

func main() {

	nums := make([]int, 10, 10)

	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(100)
	}

	fmt.Println(nums)
	Sort_func(nums)
	fmt.Println(nums)
}


func Sort_func(data []int) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				tmp := data[j]
				data[j] = data[j+1]
				data[j+1] = tmp
			}
		}
	}
}