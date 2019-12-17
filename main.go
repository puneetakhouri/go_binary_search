package main

import "bufio"
import "fmt"
import "os"
import "strconv"

func main() {
	fmt.Printf("Enter the value to search in the array arr=3,5,7,9,19\n")
	arr := []int{3, 5, 7, 9, 19}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		i, _ := strconv.Atoi(input)
		if search(arr, i) {
			fmt.Println("Found")
		} else {
			fmt.Println("Not Found")
		}
		break
	}
}

func search(arr []int, i int) bool {
	return binarySearch(arr, 0, len(arr), i)
}

func binarySearch(arr []int, start int, end int, search int) bool {

	fmt.Printf("start=%d, end=%d, search=%d\n", start, end, search)
	if start == end {
		return search == arr[start]
	}
	if search > arr[end-1] {
		fmt.Printf("search greater than length\n")
		return false
	}

	if search == arr[end-1] || search == arr[start] {
		fmt.Println("%d found, returning true", search)
		return true
	}

	if search >= arr[end/2] {
		fmt.Println("search > arr.len/2")
		return binarySearch(arr, end/2, end, search)
	} else {
		fmt.Println("search < arr.len/2")
		return binarySearch(arr, start, end/2, search)
	}
}