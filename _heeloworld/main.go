package main

import "fmt"

func main() {

	name := "imran khan rafa" // shorthand

	fmt.Println("hello world " + name)

	// loop
	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}

	// range
	for i := range 3 {
		fmt.Println(i)
	}
}
