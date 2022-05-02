package main

import "fmt"

func main() {
	var animal = "eagle"
	fmt.Println("My favorite animal is", animal)

	petBorn, petAge := 2022, 1
	fmt.Println("I my", animal, "was born in", petBorn, "It is", petAge, "years old")

	var (
		petfirstIn = "C"
		petlastIn  = "B"
	)

	fmt.Println("My pet name start with an", petfirstIn, "It's last name starts with a", petlastIn)

	var petdaysAlive int

	petdaysAlive = petAge * 365

	fmt.Println("my pet is", petdaysAlive, "day old")

}
