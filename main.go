package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeRandomSlice(numItems, max int) []int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, numItems)

	for i := 0; i < numItems; i++ {
		slice[i] = random.Intn(max)
	}
	return slice
}

func printSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

func checkSorted(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			fmt.Println("The slice is NOT sorted!")
			return
		}
	}
	fmt.Println("The slice is sorted")
}

func quicksort(slice []int) {
	hi := len(slice)

	if hi <= 1 {
		return
	}

	p := partition(slice, 0, len(slice)-1)
	quicksort(slice[:p])
	quicksort(slice[p+1:])
}

func partition(slice []int, lo, hi int) int {
	pivot := slice[hi]

	i := lo - 1

	for j := lo; j < hi; j++ {
		if slice[j] <= pivot {
			i = i + 1
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	i = i + 1
	slice[i], slice[hi] = slice[hi], slice[i]
	return i
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	quicksort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}
