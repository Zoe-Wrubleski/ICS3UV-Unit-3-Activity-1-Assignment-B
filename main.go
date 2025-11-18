/*
 * Author Zoe Wrubleski
 * Version 1.0.0
 * Date 2025-11-18
 * Fileoverview This program calculates the cost of a meal
 */

package main

import "fmt"

func main() {
	// set variables 
	var subtotal float32 = 415.50
	var people int = 8

	// calculate tax of 13%
	var tax float32 = subtotal * 0.13

	// calculate tip of 10%
	var tip float32 = subtotal * 0.1

	// calculate the total
	var total float32 = subtotal + tax + tip

	// calculate the cost per person
	var perPerson float32 = total / float32(people)

	// display answers
	fmt.Println("Subtotal: $", subtotal)
	fmt.Println("Tax (13%): $", tax)
	fmt.Println("Tip (10%): $", tip)
	fmt.Println("Total: $", total)
	fmt.Println("Cost per person: $", perPerson)

	fmt.Println("\nDone.")
}
