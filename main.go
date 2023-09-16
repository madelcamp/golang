package main

import "fmt"

var pl = fmt.Println
var pf = fmt.Printf
var sf = fmt.Sprintf

func main() {
	// print
	fmt.Println("hello world")

	// strings
	var nameOne string = "Mario"
	var nameTwo = "Luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Peach"
	nameThree = "Bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Yoshi"
	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 21
	println(numOne)

	var numTwo int8 = -128
	println(numTwo)

	var numThree uint = 25
	println(numThree)

	var scoreOne float32 = -1.5
	var scoreTwo float64 = 738434.7
	scoreThree := 1.5

	println(scoreOne, scoreTwo, scoreThree)

	// Print
	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line\n")

	// Println
	fmt.Println("Hello ninjas")
	fmt.Println("Goodbye ninjas")

	age := 25
	name := "Mario"

	pl("my age is", age, "and my name is", name)

	//Printf (formatted strings) %_ = format specifier
	pf("my age is %v and my name is %v \n", age, name)
	pf("my age is %q and my name is %q \n", age, name)
	pf("my age is of type %T \n", age)
	pf("You scored %f points! \n", 225.55)
	pf("You scored %0.1f points! \n", 225.55)

	// Sprintf (save formatted strings)
	var str = sf("my age is %v and my name is %v \n", age, name)

	pl("The saved string is: ", str)

}
