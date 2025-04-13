package main

import (
	"fmt"
	"stream-api/streams"
)

func main() {
	// Sample input slice
	a := []int{1, 2, 3, 4, 5}

	// Chaining Filter, Map, and then converting to slice
	b := streams.New(a).
		Filter(func(i int) bool {
			return i%2 == 0 // Keep even numbers
		}).
		Map(func(i int) any {
			return i * i // Square each number
		}).Collect()

	// Output: [4 16]
	fmt.Println("Filtered and mapped:", b)

	// Use Reduce to sum up the elements of the slice
	sum := streams.New(a).
		Reduce(func(acc, v int) int {
			return acc + v // Sum the elements
		}, 0)

	// Output: 15
	fmt.Println("Sum:", sum)

	// Use Find to get the first even number
	found, _ := streams.New(a).Find(func(i int) bool {
		return i%2 == 0
	})

	// Output: 2 true
	fmt.Println("First even number:", found)

	// Use Count to count elements
	count := streams.New(a).Count()

	// Output: 5
	fmt.Println("Count of elements:", count)

	// Reverse the slice
	reversed := streams.New(a).Reverse().Collect()

	// Output: [5 4 3 2 1]
	fmt.Println("Reversed slice:", reversed)

	forEach := streams.New(a).ForEach(func(i int) {
		fmt.Printf("Element: %d, ", i)
	})
	fmt.Println()

	fmt.Println("ForEach result:", forEach)
}
