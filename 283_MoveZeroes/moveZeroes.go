// Program 283 Move Zeroes Given an integer array nums, move all 0's to the end of it
// while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.

// Example 1:

// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Example 2:

// Input: nums = [0]
// Output: [0]

/*
Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1
*/

package main

import "fmt"

func moveZeroes(nums []int) []int {
	left := 0

	for right := left + 1; right < len(nums); right++ {
		if nums[left] != 0 {
			left++
		}

		if nums[right] == 0 {
			continue
		}

		if nums[right] != 0 {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
		}
	}

	return nums
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	res := moveZeroes(nums)
	fmt.Println(res)
}
