//// 2023/10/13 // 21:08 //

// bool

// string

// int int8 int16 int32 int64
// uint uint8 uint16 uint16 uint64 uintptr

// byte // alias for uint8

// rune // alias for int 32
//		// represents a Unicode code point

// float32 float64

// complex64 complex128

// Declaring a variables
// var number int
// var pi float64 = 3.14159

//// 21:14
// package main

// import "fmt"

// func main() {
// 	// initialize variables here
// 	var smsSendingLimit int
// 	var costPerSMS float64
// 	var hasPermission bool
// 	var username string

// 	fmt.Printf(
// 		"%v %f %v %q\n",
// 		smsSendingLimit,
// 		costPerSMS,
// 		hasPermission,
// 		username,
// 	)

// }

//// 21:18
//// Short Variable Declaration
// var empty string
// empty := ""

// numCars := 10 // inferred to be an integer
// temperature := 0.0 // temperature is inferred to be a floating point number
// var isFunny = true // isFunny is inferred to be a boolean

// package main

// import "fmt"

// func main() {
// 	// declare here

// 	congrats := "happy birthday!"

// 	fmt.Println(congrats)

// }

//// 21:21
// var i int
// j := i // j is also an int

// i := 42 // int
// f := 3.14 // float64
// g := 0.867 + 0.5i // complex128

// package main

// import "fmt"

// func main() {
// 	penniesPerText := 2.0
// 	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)
// }

//// 21:27

//// Sme Line Declarations
// mileage, company := 80276, "Tesla"
// is the same as
// mileage := 80276
// company := "Tesla"

// package main

// import "fmt"

// func main() {
// 	// declare here
// 	avearageOpenRate, displayMessage := .23, "is the average open rate of your messages"

// 	fmt.Println(avearageOpenRate, displayMessage)
// }

//// 21:31
//// Conversion
// temperatureInt := 88
// temperatureFloat := float64(temperatureInt)

// package main

// import "fmt"

// func main() {
// 	accountAge := 2.6

// 	// create a new "acoountAgeInt" here
// 	accountAgeInt := int(accountAge)
// 	// it should be the result of casting "accountAge" to an integer

// 	fmt.Println("Your account has existed for", accountAgeInt, "years")
// }

//// 21:34
//// Prefer "DEFAULT" types
// Unless you have a good reason to, stick to the following types:
// bool
// string
// int
// uint32
// byte
// rune
// float64
// complex128

//// 21:36
//// Use Two Separate Constants

// package main

// import "fmt"

// func main() {
// 	const premiumPlanName = "Preium Plan"
// 	const basicPlanName = "Basic Plan"

// 	// don't edit below this line

// 	fmt.Println("plan:", premiumPlanName)
// 	fmt.Println("plan:", basicPlanName)
// }

//// 22:07
// Computed Constatns
// const firstName = "Lane"
// const lastName = "Wagner"
// const fullName = firstName + " " + lastName

// package main

// import "fmt"

// func main() {
// 	const secondsInMinute = 60
// 	const minutesInHour = 60
// 	const secondsInHour = minutesInHour * secondsInMinute

// 	// don't edit below this line
// 	fmt.Println("number of seconds in an hour:", secondsInHour)
// }

//// 22:10
// Formatting Strings in Go

// %v - Interpolate the default representation
// fmt.Printf("I am %v years old", 10)
// I am 10 years old
// fmt.Printf("I am %v years old", "way too many")
// I am way too many years old

// %s - Interpolate a string
// fmt.Printf("I am %s years old", "way too many")
// I am way too many years old

// %d - Interpolate an integer in decimal form (as opposed to a binary form)
// fmt.Printf("I am %d years old", 10)
// I am 10 years old

// %f - Interpolate a decimal
// fmt.Printf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decomal places
// fmt.Printf("I am %.2f years old", 10.523)
// I am 10.53 years old

// package main

// import "fmt"

// func main() {
// 	const name = "Saul Goodman"
// 	const openRate = 30.5

// 	msg := fmt.Sprintf(
// 		"Hi %s, your open rate is %.1f percent",
// 		name,
// 		openRate,
// 	)

// 	// ?

// 	// don't edit below this line

// 	fmt.Println(msg)
// }

// //// 22:20
// //// Conditionals
// // if statements in Go don't use parentheses aroudn the condition
// if height > 4 {
// 	fmt.Println("You are tall enough!")
// }
// // else if and else are supported as you would expect:
// if height > 6 {
// 	fmt.Println("You are super tall!")
// } else if height > 4 {
// 	fmt.Println("You are tall enough!")
// } else {
// 	fmt.Println("You could be taller but hey... At least you're cool enough!")
// }

// Here are some of the comparison operators in Go:
// == equal to
// != not equal to
// < less than
// > greater than
// <= less than or equal to
// >= greater than or equal to

package main

import "fmt"

func main() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen)

	// don't touch above this line

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}

//// 22:30
// Alternate way of writing the if statement
// The initial statement of an if block
// if INITIAL_STATEMENT; CONDITION {
//
// }

// // This is syntactic sugar to shorten your code, for example this:
// length := getLength(email)
// if length < 1 {
// 	fmt.Println("Email is invalid")
// }
// // We can do:
// if length := getLength(email); length < 1 {
// 	fmt.Println("Email is invalid")
// }
// Not only is this code a bit shorter, but it also removes length from the
// parent scope, which is convenient because we don't need it there - we only
// need access to it while checking a condition.
