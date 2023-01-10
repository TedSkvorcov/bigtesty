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
	fmt.Println("Need Sort_func")
}
