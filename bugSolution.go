```go
package main

import "fmt"

func main() {

    var x int = 10
    var y *int = &x // y is a pointer to x

    fmt.Println(*y) // Output: 10

    *y = 20 // Modifying the value of x through the pointer y

    fmt.Println(x) // Output: 20

    // Demonstrating correct use of pointers:
    z := 30
    y = &z // Now, y points to a new variable z
    *y = 40

    fmt.Println(*y) // Output: 40
    fmt.Println(z) // Output: 40
    fmt.Println(x) // Output: 20 (x remains unchanged)
}
```