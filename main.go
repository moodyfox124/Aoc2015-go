package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("./data/day1,txt")
	check(err)

	floor := 0

	for _, char := range data {
		if char == '(' {
			floor += 1
		} else {
			floor -= 1
		}
	}
	fmt.Println("Floor number: ", floor)

	var index_to_go_to_basement int
	basement_floor := 0

	for i, char := range data {
		if char == '(' {
			basement_floor += 1
		} else {
			basement_floor -= 1
		}

		if basement_floor == -1 {
			index_to_go_to_basement = i + 1
			break
		}

	}

	fmt.Println("Position of first character to go in basemant: ", index_to_go_to_basement)
}
