package main

import "fmt"

func main() {
	/*Prvi zadatak*/
	var firstNumber int = 30
	var secondNumber int = 20

	firstNumber = firstNumber + secondNumber
	secondNumber = firstNumber - secondNumber
	firstNumber = firstNumber - secondNumber

	fmt.Println("Prvi broj je: ", firstNumber)
	fmt.Println("Drugi broj je: ", secondNumber)

	/*Drugi zadatak*/

	var firstName string = "Karla"
	var lastName string = "Fišić"
	const fullName = "Ime i prezime: "
	var imeprezime = fullName + " " + firstName + " " + " " + lastName

	fmt.Println(imeprezime)

}
