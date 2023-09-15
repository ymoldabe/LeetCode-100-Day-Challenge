package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	k := removeElement(&nums, val)
	sort.Ints(nums)
	fmt.Println(k, nums)
}

func removeElement(nums *[]int, val int) int {
	for i := 0; i < len(*nums); i++ {
		if (*nums)[i] == val {
			if i < len(*nums) {
				*nums = append((*nums)[:i], (*nums)[i+1:]...)
			} else if i == len(*nums)-1 {
				*nums = (*nums)[:len(*nums)-1]
			}
			i -= 1
		}
	}
	return len(*nums)
}
