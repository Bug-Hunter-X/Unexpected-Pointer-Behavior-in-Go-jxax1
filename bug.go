```go
package main

import "fmt"

func main() {

    var x int = 10
    var y *int = &x // y is a pointer to x

    fmt.Println(*y) // Output: 10

    *y = 20 // Modifying the value of x through the pointer y

    fmt.Println(x) // Output: 20

    fmt.Println(y) // Output: 0xc0000140a8
}
```