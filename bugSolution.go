```go
package main

import "fmt"

func main() {
    var m map[string]int
    if m == nil {
        m = make(map[string]int)
    }
    m["key"] = 10
    fmt.Println(m["key"])

    //Alternative approach using comma ok idiom
    value, ok := m["key"]
    if ok{
        fmt.Println("value found", value)
    } else {
        fmt.Println("key not found")
    }
}
```