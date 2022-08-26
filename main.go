package main

import "fmt"

func main() {
	nums := []int{2, 5, 6, 9, 10}

	fmt.Println(findGCD(nums))
}
func findGCD(nums []int) int {
	return common(smallest(nums), greatest(nums))

}

func common(small, large int) int {
	var greatestdivisor = 1
	for i := 2; i <= small; i++ {
		if large%i == 0 && small%i == 0 {
			greatestdivisor = i
			return greatestdivisor
		}
	}
	return greatestdivisor
}

func greatest(nums []int) int {
	greatest := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[0] {
			greatest = nums[i]
		}
	}

	return greatest
}

func smallest(nums []int) int {
	smallest := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums[0] {
			smallest = nums[i]
		}
	}
	return smallest
}
