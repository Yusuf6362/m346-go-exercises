package main

import "fmt"

var firstName, lastName string = "Yusuf", "Er"
var dayOfBirth, monthOfBirth, yearOfBirth, numberOfSiblings int = 11, 01, 2007, 2
var heightInMeters float32 = 1.80
var zodiacSign rune = '\u2651'

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
