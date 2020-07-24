package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	s := make(map[int]int)

	// For every element, search for it's complement
	for i := range nums {
		complement := target - nums[i]
		value, exists := s[complement]
		if exists {
			return []int{value, i}
		}
		s[nums[i]] = i
	}

	return []int{0, 0}

}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func main() {

	// Test1
	test1Arr := []int{2, 7, 11, 15}
	target1 := 9
	expected1 := []int{0, 1}
	var res = twoSum(test1Arr, target1)

	fmt.Printf("Test array 1: %v \ntarget: %d\nexpected value: %v\n", test1Arr, target1, expected1)

	if Equal(res, expected1) {
		fmt.Println("Success!!")
	} else {
		fmt.Println("Failure...")
	}

}
