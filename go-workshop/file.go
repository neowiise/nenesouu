package main 

import (
    "fmt"
    "io/ioutil"
)

func main() {

    data, err := ioutil.ReadFile("hello.txt")
 
    if err != nil {
        fmt.Println(err)
    }
    fmt.Print(string(data))

}
