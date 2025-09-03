package main

import "fmt"

// jwtToken := 30000 // var less wont work outside the function

const LoginToken string = "nklhbnjlkbvhb" // Public // capital "L" here signifies it is a public varible

func main() {
	var username string = "ojass"
	fmt.Println(username)
	fmt.Printf("Varible is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Varible is of type: %T \n", isLoggedIn)

	var smallVal uint = 255
	fmt.Println(smallVal)
	fmt.Printf("Varible is of type: %T \n", smallVal)

	var smallFloat float64 = 255.234567556453423
	fmt.Println(smallFloat)
	fmt.Printf("Varible is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Varible is of type: %T \n", anotherVariable)

	//implicit type
	var website = "chaiaurcode.ojass.in"
	fmt.Println(website)

	//no var style
	numberofUser := 30000
	fmt.Println(numberofUser)

	fmt.Println(LoginToken)
	fmt.Printf("Varible is of type: %T \n", LoginToken)
}

//inside ant\y method you are allowed to use the var less operator
