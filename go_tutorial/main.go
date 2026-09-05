//go run main.go
//quicksave first

//go mod init first_go_project  creates a go.mod file

package main

import "fmt"

func variablesTutorial() {
	name := "John"
	age := 20 //var age int=20

	fmt.Println(name, age)

	name = "Bob"
	age = 21

	fmt.Println(name, age)
}

func operatorsTutorial() {
	number := 10

	number += 5
	fmt.Println(number)

	number -= 3
	fmt.Println(number)

	number *= 2
	fmt.Println(number)

	number /= 4
	fmt.Println(number)
}

func inputTutorial() {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}

func conditionsTutorial() {
	age := 20

	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 13 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Child")
	}
	// && and
	// || or
}

func switchTutorial() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}
	//Doesn't need break
}

func loopsTutorial() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//or

	number := 0

	for number < 5 {
		fmt.Println(number)
		number++
	}
	//break breaks loop
	//continue skips current iteration
}

func arraysTutorial() {
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println(numbers)      //prints whole array
	fmt.Println(numbers[2])   //access an iten
	fmt.Println(len(numbers)) //get array length

	for i := 0; i < len(numbers); i++ { // loops through array
		fmt.Println(numbers[i])
	}
}

func slicesTutorial() {
	//same as arrays, but don't have fixed size

	numbers := []int{10, 20, 30, 40, 50}

	numbers = append(numbers, 60) // adds an item to the back, can add multiple items

	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5, 6}
	numbers1 = append(numbers1, numbers2...) // add another slice

	numbers = append(numbers[:2], numbers[3:]...) // removes 2nd index

	for i, number := range numbers { // loops through slice
		fmt.Println(i, number)
	}

	for _, number := range numbers { // index is not required
		fmt.Println(number)
	}
}

func mapsTutorial() {
	ages := map[string]int{ //first show name of variable[string] and then the data type[int]
		"John":  20,
		"Mike":  25,
		"Sarah": 22,
	}

	fmt.Println(ages)

	fmt.Println(ages["John"]) // access, doesn't guarantee order
	delete(ages, "Mike")      //delete

	age, exists := ages["John"] //check if exists

	if exists {
		fmt.Println("Age:", age)
	}

	fmt.Println(len(ages)) // length
}

func functionsTutorial(a int, b int) int { //return type showed at the end
	return a + b

	//func getPerson() (string, int) { //functions can return multiple variables
	//	return "John", 20
	//}
	//name, age := getPerson()
	//name, _ := getPerson() you can ignore a value
}

func variadicFunctionsTutorial(numbers ...int) int { //infinite values can be added
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
	//result := variadicFunctionsTutorial(10, 20, 30, 40)
	//result := variadicFunctionsTutorial(numbers...) can add whole lists
}

func structTutorial() {
	type Player struct {
		name   string
		health int
		damage int
		level  int
	}

	players := []Player{
		{
			name:   "Knight",
			health: 100,
			damage: 25,
			level:  5,
		},
		{
			name:   "Wizard",
			health: 70,
			damage: 40,
			level:  7,
		},
	}

	fmt.Println(players[0].name)

	for _, player := range players {
		fmt.Println(player.name, player.health)
	}
}

func main() {

	fmt.Println("Hello world")

}
