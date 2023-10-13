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

package main

import "fmt"

func main() {
	const premiumPlanName = "Preium Plan"
	const basicPlanName = "Basic Plan"

	// don't edit below this line

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
}
