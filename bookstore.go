//This ia an book store cli-app
package main

import (
	"fmt"
	"reflect"
)


//////////////// User interface ///////////////////////////////
type user struct {
	userName string
	userContact uint
	userEmail string
	userCart := make(map[string]float64)
)


// Prints user info
func (u *user) getUserInfo() {
	fmt.Printf("User Name: %s\nUser Contact: %d\nUser Email: %s\n",
		userName, userContact, userEmail)
}

// print users items out
func (u *user) getUserCart(cart map) {
	fmt.Printf("Customer %s's Cart\n", userName) 
	fmt.Println("---------------------------------")
	fmt.Println("ITEM\t\t\Price\n")
	if cart != "map[]"  {
		for key, value := range cart {
			fmt.Printf("%s\t\t%.2f", key, value)
		}
	} else {
		fmt.Println("Note!: Nothing in the cart to display")
	}
}

/////////////////////////////////////////////////////////////////
