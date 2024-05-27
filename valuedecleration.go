/*package main

import (
	"fmt"
)

var name string
var age int

func main() {
	name = "john"
	age = 19
	fmt.Println(name, age)
}
*/

/*package main

import (
	"fmt"
)

func main() {
	var a, b = 6, "Hello"
	c, d := 7, "World!"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
*/

package main

import (
	"fmt"
)

func main() {
	prices := [3]int{10, 20, 30}

	prices[2] = 50
	fmt.Println(prices)
}
