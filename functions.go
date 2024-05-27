/*package main

import (
	"fmt"
)

func myMessage() {
	fmt.Println("I just got executed!")
}

func main() {
	myMessage()
	myMessage()
	myMessage()
}
*/

package main

import (
	"fmt"
)

func myFunction(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(myFunction(1, 2))
}
