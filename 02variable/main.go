package main

import "fmt"

const LoginToken string = "fkdhhfivh" //public

func main() {
	var username string = "Lavish"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type : %T \n", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type : %T \n", smallVal)

	var smallFloat float64 = 255.324345456778565
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	//default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	var anotherUsername string
	fmt.Println(anotherUsername)
	fmt.Printf("Variable is of type : %T \n", anotherUsername)

	// implicit type or way  of declaring variables

	var website = "Kubesimplify.in"
	fmt.Println(website)

	// no using of var style for declaring the variables

	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)

}
