package main

import "fmt"

func main() {
	var firstName string = "Aditya"

	var lastName string = "Fajri"

	fmt.Println(firstName, lastName)

	var (
		age     int
		address string
	)

	age = 23
	address = "Kebumen"

	fmt.Printf("Halo, saya %s %s berumur %d dan saya tinggal di %s", firstName, lastName, age, address)
}
