package main

import "fmt"

func main(){
s:=[6] int {2,3,4,5,6,7}
var y [] int=s[1:6]
fmt.Println(y[0])
fmt.Println(y)
y= y[2:5]
fmt.Println(y)
fmt.Println(y,len(y),cap(y))
}
