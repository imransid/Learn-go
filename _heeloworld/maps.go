// // package main

// // import "fmt"

// // func main() {
// // 	m := make(map[string]string)

// // 	m["name"] = "golang"

// // 	fmt.Println(m["name"])

// // }

// // package main

// // import "fmt"

// // // veridic function
// // func sum(nums ...int) int {

// // 	total := 0

// // 	for _, num := range nums {

// // 		total = total + num
// // 	}

// // 	return total

// // }

// // func main() {
// // 	result := sum(1, 3, 4)
// // 	fmt.Println(result)
// // }

// // nopw closer

// package main

// import "fmt"

// func counter() func() int {

// 	var count int = 0

// 	return func() int {
// 		count = +1
// 		return count
// 	}

// }

// func main() {

// 	increment := counter()

// 	fmt.Println(increment)
// 	fmt.Println(increment())

// }

package main

import "fmt"

type order struct {
	name   string
	age    int
	status bool
}

func newOrder(name string, age int, status bool) *order {

	myOrder := order{
		name:   name,
		age:    age,
		status: status,
	}

	return &myOrder

}

func (o *order) changeStatus(status bool) {
	o.status = status
}

func (o order) getAge() int {
	return o.age
}

func main() {

	myOrder := newOrder("imran", 30, false)

	fmt.Println(myOrder.name)

	fmt.Println(myOrder)

}
