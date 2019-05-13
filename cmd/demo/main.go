package main

import (
	"fmt"
	"os"

	"github.com/thoas/go-funk"
)

func main() {
	// Print all environment variables.
	for _, e := range os.Environ() {
		fmt.Println(e)
	}

	fmt.Println()
	fmt.Println("----------------------------")

	fmt.Printf("funk.Sum: %f\n", funk.Sum([]int{0, 1, 2, 3, 4}))
}
