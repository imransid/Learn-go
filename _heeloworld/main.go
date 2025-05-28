package main

import "fmt"

// "time"

func main() {

	// name := "imran khan rafa" // shorthand

	// fmt.Println("hello world " + name)

	// // loop
	// fmt.Println("LOOP")
	// for i := 1; i < 3; i++ {
	// 	fmt.Println(i)
	// }

	// fmt.Println("Range")
	// // range
	// for i := range 3 {
	// 	fmt.Println(i)
	// }

	// fmt.Println("IF ELSE")

	// age := 116

	// if age <= 16 {
	// 	fmt.Println("KID")
	// } else if age <= 20 {
	// 	fmt.Println("ADULT")
	// } else {
	// 	fmt.Println("ELDER")
	// }

	// // there was no ternary operator in golang

	// if name := "imran"; name == "imran" {
	// 	fmt.Println("WELCOME ", name)
	// } else {
	// 	fmt.Println(name, " You are not Admin")
	// }

	// // switch
	// fmt.Println("SWITCH")

	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println("ONE")
	// case 2:
	// 	fmt.Println("TWO")
	// case 3:
	// 	fmt.Println("THREE")
	// default:
	// 	fmt.Println("FIVE")
	// }

	// // multiple condition switch

	// fmt.Println("MULTIPLE CONDITION SWITCH")

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Friday:
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("workday")

	// }

	fruits := []string{"apple", "banana", "mango"}

	val := "banana"
	for i, v := range fruits {
		if v == val {
			fruits = append(fruits[:i], fruits[i+1:]...)
			break
		}
	}

	fmt.Println("fruits", fruits)

}
