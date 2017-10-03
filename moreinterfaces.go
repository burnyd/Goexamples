package main

import "fmt"
import "reflect"


type Test1 struct {
 A string
}

func main() {
var t interface{}
t = Test1{"Hello world!"}
b := []string{"r","u","n"}
c := b[1:2]
data := reflect.ValueOf(t)
fmt.Printf("%s\n", data.FieldByName("A"))
fmt.Println(b)
fmt.Println(c)
}
