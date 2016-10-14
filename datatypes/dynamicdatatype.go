package main

import "fmt"

func main() {
   var x float64 = 20.0

   // Dynamic typechecking, here you dont need to write the datatype
   y := 42 

   z := "Hello"

   fmt.Println(x)
   fmt.Println(y)
   fmt.Println(z)

   fmt.Printf("x is of type %T\n", x)
   fmt.Printf("y is of type %T\n", y)
   fmt.Printf("z is of type %T\n", z)
}