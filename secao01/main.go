// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello World")
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {

// 	if len(os.Args) != 2 {
// 		os.Exit(1)
// 	}
// 	fmt.Println("Hello", os.Args[1])

// }

package main

import (
	"fmt"
)

func main() {

	name := getName()
	idade := 34.000
	fmt.Println("Hello", name, idade)

}

func getName() string {
	return "Victor"
}
