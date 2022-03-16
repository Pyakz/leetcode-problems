package main

import "fmt"
func main() {
	fmt.Println("Leet Code Practice")
}

// func main() {
// 	// 10
// 	// 19
// 	// 27
// 	// 34
// 	// 40
// 	// 45
// 	// 49
// 	// 52
// 	// 54
// 	// 55
// 	numbers := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
// 	target := 54
// 	fmt.Println(AddSum(numbers, target))
// }

// func AddSum(nums []int, target int) []int {
// 	currentValue := 0
// 	startingIndex := 0
// 	matcherIndex := 0
// 	for index, num := range nums {
// 		if currentValue == target {
// 			matcherIndex = index - 1
// 		}
// 		currentValue = currentValue + num
// 	}
// 	return []int{startingIndex, matcherIndex}
// }
