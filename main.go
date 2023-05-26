// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello Waleed")
// }

// VARIABLE

// func main() {
// 	var student1 string = "John" //type is string
// 	var student2 = "Jane"        //type is inferred
// 	x := 2                       //type is inferred

// 	fmt.Println(student1)
// 	fmt.Println(student2)
// 	fmt.Println(x)
// }

// MULTIPLE VARIABLE
// 1

// package main
// import ("fmt")

// func main() {
//   var a, b, c, d int = 1, 3, 5, 7

//   fmt.Println(a)
//   fmt.Println(b)
//   fmt.Println(c)
//   fmt.Println(d)
// }

// 2
// package main
// import ("fmt")

// func main() {
//   var a, b = 6, "Hello"
//   c, d := 7, "World!"

//   fmt.Println(a)
//   fmt.Println(b)
//   fmt.Println(c)
//   fmt.Println(d)
// }

// 3
// package main

// import "fmt"

// func main() {
// 	var (
// 		a int
// 		b int    = 1
// 		c string = "hello"
// 	)

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// DATA TYPES
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a bool = true    // Boolean
// 	var b int = 5        // Integer
// 	var c float32 = 3.14 // Floating point number
// 	var d string = "Hi!" // String

// 	fmt.Println("Boolean: ", a)
// 	fmt.Println("Integer: ", b)
// 	fmt.Println("Float:   ", c)
// 	fmt.Println("String:  ", d)
// }

// GOLANG ARRAY with The Length defined

// package main
// import "fmt"

// func main() {
// 	var arr1 = [3]int{1, 2, 3}
// 	arr2 := [5]int{4, 5, 6, 7, 8}

// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// }

// This example declares two arrays (arr1 and arr2) with inferred lengths:
// package main

// import "fmt"

// func main() {
// 	var arr1 = [...]int{1, 2, 3}
// 	arr2 := [...]int{4, 5, 6, 7, 8}

// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// }

// Declares an array of strings

// package main
// import "fmt"

// func main() {
// 	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
// 	fmt.Print(cars)
// }

// how to access the elements and find the length of variable

// package main

// import "fmt"

// func main() {
// 	names := [...]string{"waleed", "Hassan", "Abdullah", "Joel"}
// 	fmt.Println(len(names), names[0])
// }

// The array above has 5 elements.

// 1:10 means: assign 10 to array index 1 (second element).
// 2:40 means: assign 40 to array index 2 (third element).

// package main
// import "fmt"

// func main() {
//   arr1 := [5]int{1:10,2:40}

//   fmt.Println(arr1)
// }

// SLICE

// Simple explanation Slice are self growing form of array so there are two main properties.

// Length is total no of elements() the slice is having and can be used for looping through the elements we stored in slice.
// Also when we print the slice all elements till length gets printed.

// Capacity is total no elements in underlying array, when you append more elements the length increases till capacity.
// After that any further append to slice causes the capacity to increase automatically(apprx double) and length by no of elements appended.

// package main
// import "fmt"

// func main() {
// 	// myslice1 := []int{}
// 	// fmt.Println(len(myslice1))
// 	// fmt.Println(cap(myslice1))
// 	// fmt.Println(myslice1)

// 	myslice2 := [...]string{"Go", "Slices", "Are", "Powerful"}
// 	// fmt.Println(len(myslice2))
// 	// fmt.Println(cap(myslice2))
// 	fmt.Println(myslice2)
// }

// example 2
// To append elements to the end of a slice:
// package main
// import "fmt"

// func main() {
//   myslice1 := []int{1, 2, 3, 4, 5, 6}
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))

//   myslice1 = append(myslice1, 20, 21)
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))
// }

// to append one slice to another slice:

// package main

// import "fmt"

// func main() {
// 	myslice1 := []int{1, 2, 3}
// 	myslice2 := []int{4, 5, 6}
// 	myslice3 := append(myslice1, myslice2...)

// 	fmt.Printf("myslice3=%v\n", myslice3)
// 	fmt.Printf("length=%d\n", len(myslice3))
// 	fmt.Printf("capacity=%d\n", cap(myslice3))
// }

// Change the length of slice:
// package main

// import "fmt"

// func main() {
// 	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
// 	myslice1 := arr1[1:5]                 // Slice array
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))

// 	myslice1 = arr1[1:3] // Change length by re-slicing the array
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))

// 	myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))
// }

// GoLANG operators

// package main

// import "fmt"

// func main() {
// 	var (
// 		sum1 = 100 + 50    // 150 (100 + 50)
// 		sum2 = sum1 + 250  // 400 (150 + 250)
// 		sum3 = sum2 + sum2 // 800 (400 + 400)
// 	)
// 	fmt.Println(sum3)
// }

// GOLANG CONDITIONS

// IF ELSE

// package main

// import "fmt"

// func main() {
// 	temperature := 16
// 	if temperature > 15 {
// 		fmt.Println("It is warm out there")
// 	} else {
// 		fmt.Println("It is cold out there")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	time := 2
// 	if time < 10 {
// 		fmt.Println("Good morning.")
// 	} else if time < 20 {
// 		fmt.Println("Good day.")
// 	} else {
// 		fmt.Println("Good evening.")
// 	}
// }

// nested if

// package main

// import "fmt"

// func main() {
// 	num := 20
// 	if num >= 10 {
// 		fmt.Println("Num is more than 10.")
// 		if num > 15 {
// 			fmt.Println("Num is also more than 15.")
// 		}
// 	} else {
// 		fmt.Println("Num is less than 10.")
// 	}
// }

// Switch Cases

// package main

// import "fmt"

// func main() {
// 	day := 2

// 	switch day {
// 	case 1:
// 		fmt.Println("Monday")
// 	case 2:
// 		fmt.Println("Tuesday")
// 	case 3:
// 		fmt.Println("Wednesday")
// 	case 4:
// 		fmt.Println("Thursday")
// 	case 5:
// 		fmt.Println("Friday")
// 	case 6:
// 		fmt.Println("Saturday")
// 	case 7:
// 		fmt.Println("Sunday")
// 	default:
// 		fmt.Println("Not a weekday")
// 	}
// }

// Multi Cases

// package main

// import "fmt"

// func main() {
// 	day := 5

// 	switch day {
// 	case 1, 3, 5:
// 		fmt.Println("Odd weekday")
// 	case 2, 4:
// 		fmt.Println("Even weekday")
// 	case 6, 7:
// 		fmt.Println("Weekend")
// 	default:
// 		fmt.Println("Invalid day of day number")
// 	}
// }

// GOLANG FOR LOOPS

// package main

// import "fmt"

// func main() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}
// }

// package main
// import "fmt"

// func main() {
// 	for i := 0; i <= 100; i += 10 {
// 		fmt.Println(i)
// 	}
// }

// TABLE OF ANY NUMBER IN GOLANG

// package main

// import "fmt"

// func main() {
// 	var num int = 0

// 	fmt.Print("Enter Number: ")
// 	fmt.Scanf("%d", &num)

// 	for i := 1; i <= 10; i++ {
// 		fmt.Printf("%d * %d = %d\n", num, i, num*i)
// 	}
// }

// Nested Loops

// package main

// import "fmt"

// func main() {
// 	adj := [2]string{"big", "tasty"}
// 	fruits := [3]string{"apple", "orange", "banana"}
// 	for i := 0; i < len(adj); i++ {
// 		for j := 0; j < len(fruits); j++ {
// 			fmt.Println(adj[i], fruits[j])
// 		}
// 	}
// }

// Function Returns

// package main

// import "fmt"

// func myFunction(x int, y string) (result int, txt1 string) {
// 	result = x + x
// 	txt1 = y + " World!"
// 	return
// }

// func main() {
// 	fmt.Println(myFunction(5, "Hello"))
// }

// Function Return Value

// package main

// import "fmt"

// func myFunction(x int, y int) (result int) {
// 	result = x + y
// 	return
// }

// func main() {
// 	total := myFunction(1, 2)
// 	fmt.Println(total)
// }

// Function Return Multiple Value
// package main

// import "fmt"

// func myFunction(x int, y string) (result int, txt1 string) {
// 	result = x + x
// 	txt1 = y + " World!"
// 	return
// }

// func main() {
// 	fmt.Println(myFunction(5, "Hello"))
// }

// Recursion Function
package main

import "fmt"

func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

func main() {
	testcount(1)
}

// Factorial Recursion
// package main

// import "fmt"

// func factorial_recursion(x float64) (y float64) {
// 	if x > 0 {
// 		y = x * factorial_recursion(x-1)
// 	} else {
// 		y = 1
// 	}
// 	return
// }

// func main() {
// 	fmt.Println(factorial_recursion(4))
// }
