//Problem 1209. Remove all Adjacent Duplicates in String II
// You are given a string s and an integer k, a k duplicate removal consists of
// choosing k adjacent and equal letters from s and removing them,
// causing the left and the right side of the deleted substring to concatenate together.

/*
Example 1:

Input: s = "abcd", k = 2
Output: "abcd"
Explanation: There's nothing to delete.
Example 2:

Input: s = "deeedbbcccbdaa", k = 3
Output: "aa"
Explanation:
First delete "eee" and "ccc", get "ddbbbdaa"
Then delete "bbb", get "dddaa"
Finally delete "ddd", get "aa"
Example 3:

Input: s = "pbbcggttciiippooaais", k = 2
Output: "ps"
*/

package main

import (
	"fmt"
	"strings"
)

type stackCount struct {
	num   rune
	count int
}

func main() {
	str := "deeedbbcccbdaa"
	m := make(map[string]int)
	res := removeDuplicates(str, m, 3)
	fmt.Println("Result: ", res)
}

func removeDuplicates(s string, m map[string]int, k int) string {
	stack := []stackCount{}
	//stack := []rune{}

	for _, val := range s {
		if len(stack) > 0 && stack[len(stack)-1].num == val {
			stack[len(stack)-1].count++

			if stack[len(stack)-1].count == k {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, stackCount{num: val, count: 1})
		}
	}

	fmt.Println(stack)
	var sb strings.Builder

	for _, val := range stack {
		if val.count > 1 {
			for i := 0; i < val.count; i++ {
				sb.WriteRune(val.num)
			}
		} else {
			sb.WriteRune(val.num)
		}
	}

	return sb.String()
}
