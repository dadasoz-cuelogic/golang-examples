package main

import "fmt"

func main() {

    //Constants can be defined using 'const' keywords.
    const LENGTH int = 10
    const WIDTH int = 5   

    var area int

    area = LENGTH * WIDTH
    fmt.Printf("value of area : %d", area)   
}