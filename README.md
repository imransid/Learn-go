# Learn Golang 🚀

Welcome to the **Learn Golang** project! This repository is designed for beginners who want to learn the Go programming language through hands-on practice and simple code examples.

---

## 📌 What is Go?

[Go](https://go.dev/), also known as Golang, is an open-source programming language created by Google. It is known for its simplicity, performance, and strong support for concurrency.

---

## 🛠 Prerequisites

Before you start, make sure you have Go installed on your machine.

### ✅ Install Go

- Download: [https://go.dev/dl](https://go.dev/dl)
- Follow the installation instructions for your operating system.

After installation, verify it:

Go (Golang) Beginner Guide

This document summarizes all the important beginner concepts in Go (Golang) with simple explanations, real-life examples, and code snippets.

1. Variables and Arrays

Arrays:

var names = [3]string{"rafa", "imran", "opu"}

Fixed size

Can’t grow

Use when you know the exact number of elements

Check if array is blank:

empty := [3]string{}
fmt.Println(empty == names) // false

2. Slices

Definition:

Dynamic size

Built on top of arrays

var list = []string{"Go", "Python"}
list = append(list, "Java")

Real Life:

To-do lists, shopping carts, dynamic data.

3. make() Function

tasks := make([]string, 0, 20) // slice with capacity 20

Creates slices, maps, and channels

Pre-allocates memory for efficiency

4. Maps

user := map[string]string{"name": "Rafa", "city": "Barishal"}

Key-value pairs

Like objects/dictionaries in other languages

Real Life:

Storing user profile data

Lookup tables

5. range Keyword

Used in loops:

for i, v := range tasks {
fmt.Println(i, v)
}

6. Variadic Functions

func sum(nums ...int) int {
total := 0
for \_, n := range nums {
total += n
}
return total
}

Accepts unlimited number of arguments

7. Closures

func counter() func() int {
count := 0
return func() int {
count++
return count
}
}

Function inside a function that remembers outer variables

8. Structs

type User struct {
Name string
Age int
}

u := User{"Rafa", 30}

Custom data types

Like classes but simpler

Real Life:

Representing users, orders, products

9. Interfaces

type Speaker interface {
Speak() string
}

Defines behavior without implementation

Allows polymorphism

Real Life:

Multiple types (Dog, Cat) implementing the same behavior (Speak)

10. Enums (Using iota)

type Day int

const (
Sunday Day = iota
Monday
Tuesday
)

Use iota to create enumerated constants

11. Goroutines

go sayHello()

Lightweight threads

Run functions concurrently

Real Life:

Sending emails, background jobs

12. WaitGroups

var wg sync.WaitGroup
wg.Add(1)
go func() {
defer wg.Done()
// work
}()
wg.Wait()

Waits for all goroutines to finish

13. Channels

ch := make(chan string)
go func() { ch <- "Hi" }()
msg := <-ch
fmt.Println(msg)

Used to communicate between goroutines

14. Mutex

var mu sync.Mutex
mu.Lock()
// critical section
mu.Unlock()

Prevents race conditions by locking shared resources

Let me know if you want this as a downloadable PDF or extended with more topics like file handling, JSON, HTTP server, etc. 🚀
