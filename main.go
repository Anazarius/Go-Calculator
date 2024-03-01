package main

import f "fmt"

func main() {
	first_number := 0.0
	second_number := 0.0
	operation := "+"
	for {
		f.Println("Enter 2 numbers")
		f.Scan(&first_number, &second_number)
		f.Println("Enter operation (+,-,/,*)")
		f.Scan(&operation)
		switch operation {
		case "+":
			f.Print("Answer: ")
			f.Println(first_number + second_number)
		case "-":
			f.Print("Answer: ")
			f.Println(first_number - second_number)
		case "/":
			if second_number == 0 {
				f.Println("ERROR")
			} else {
				f.Print("Answer: ")
				f.Println(first_number / second_number)
			}
		case "*":
			f.Print("Answer: ")
			f.Println(first_number * second_number)
		default:
			f.Println("ERROR")
		}
	}
}
