// interfaces?  Structs?

package main

import (
  "fmt"
  "strings"
)
func main() {
  // func Count(s, sep string) int
  fmt.Println(strings.Count("test", "t"))
  fmt.Println(strings.HasPrefix("test", "te"))
  // => 2
}
