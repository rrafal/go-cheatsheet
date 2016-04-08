// Welcome! Let's start main package so we'll get executable binary.
package main

// This is how we import code from other packages.
import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"runtime"
	"time"
)

// init is run when the file loads
func init() {
	// do nothing
}

// main is function runs when executable starts
func main() {
	fmt.Println("Here is how we print something to standard output.")
	log.Println("We can also print to standard error.")
}

// whatIsCheatsheetFor is private to package because it starts with lower case.
func whatIsCheatsheetFor() {
	fmt.Println("I need to learn go.")
}

// WhatIsGoFor is public; it can be imported by other packages.
func WhatIsGoFor() {
	fmt.Println("Go is great for server-side programming.")
}

// PI is a untyped numeric constant.
const Pi = 3.14159

// AUTHOR is a untyped string constant.
const Author = "Rafal Radulski"

// letsDefineVariables defines some variables.
func letsDefineVariables() {
	// int
	var a = 1
	var b int = 42
	var c, d int = 9, 27
	e := 3

	// char
	letter := 'l'

	// string
	abc := "text"

	// float
	x := 1.7

	// array of length 3
	var arr [3]int
	arr[0] = 91

	// slide of length 3, capacity 10
	s := make([]int, 3, 10)
	s = append(s, 20)

	// map of string -> int
	var m map[string]int
	m = make(map[string]int)
	m["zero"] = 0

	n := map[string]int{"zero": 0, "one": 1}
	delete(n, "one")

	// Local variables must be used or Go won't compile.
	fmt.Println(a, b, c, d, e, letter, abc, x)
}

// letsDoTypeConversion converts between char, string and int
func letsDoTypeConversion() {
	c := 'R'
	s := string(c)
	i := int(c)
	fmt.Println(c, s, i)
}

// letsDoConditional uses if statement. If can have two parts, init and condition
func letsDoConditional() {
	rand.Seed(time.Now().UTC().UnixNano())

	if r := rand.Intn(100); r > 9 {
		fmt.Println("We generated number greater than 9.")
	} else {
		fmt.Println("We generated a small number.")
	}
}

// letsDoLoops runs "for" statement a couple of ways
func letsDoLoops(start, stop int) int {
	sum := 0

	// we start with traditional for loop
	for i := start; i < stop; i++ {
		sum += i
	}

	// next is "while" loop
	for sum < 20 {
		sum += 1
	}

	// next is infinit loop
	for true {
		// but we are not going to wait forever
		break
	}

	// let's try range, which returns index and value
	list := [6]int{1, 2, 3, 5, 7, 11}
	for _, v := range list {
		sum += v
	}

	// or just index
	for i := range list {
		sum += list[i]
	}
	return sum
}

// letsDoSwitch uses switch statement
func letsDoSwitch() {
	// traditional way
	name := runtime.GOOS
	switch name {
	case "darwin":
		fmt.Println("We are on OS X.")
	default:
		fmt.Printf("We are on %s.", name)
	}

	// Go way
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("We are on OS X.")
	default:
		fmt.Printf("We are on %s.", os)
	}
}

// letsUseDefer runs defered code at the end of a function
func letsUseDefer() {
	defer fmt.Print("world.\n")
	fmt.Print("Hello ")
}

// HowDoYouReturnErrors indicates error with return value
func HowDoYouReturnErrors(wantError bool) (string, err) {
	if wantError {
		return "", errors.New("if you must... here is an error")
	} else {
		return "No error this time!", nil
	}
}

// Vertex is a simple struct
type Vertex struct {
	X int // X cordinate (uppercase == public field)
	Y int // Y cordinate
}

// String is a method which implements Stringer interface and returns string representation of an object
func (v Vertex) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

// abs is a method private to package (it's lower case)
func (v Vertex) abs() float64 {
	return math.Sqrt(float64(v.X)*float64(v.X) + float64(v.Y)*float64(v.Y))
}

// Currency is a custom type.
type Currency float64

// currency is another type, but it's private to pckage.
type currency float64

// GetNamer is interface which requires a single method GetName, which must take no parameters and return string
type GetNamer interface {
	GetName() string
}

// Animal is a new type, a struct
type Animal struct {
	name string // animal's name (lowercase == private field)
}

// GetName is a method on Animal
func (a Animal) GetName() string {
	return a.name
}

// Cat is a struct with anonymouse field. Fields and methods are inherited from Animal.
type Cat struct {
	Animal
	size int
}

// Mouse is another type. It has the same fields as Animal, but it has no methods.
type Mouse Animal

// TryToGetName accepts object of any type, and tries to get a name from object.
func TryToGetName(obj interface{}) string {
	switch o := obj.(type) {
	default:
		return "unnamed"
	case GetNamer:
		return o.GetName()
	case fmt.Stringer:
		return o.String()
	}
}
