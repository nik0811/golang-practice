package main
import (
	"fmt"
)

func main (){
     //_if()
    is_prime()
}

func _if() {
   x := 1
   if x > 10 {
      fmt.Println("my name is nik")
   }else if x > 2 {
	   fmt.Println("my name is nikhil")
   }else {
	   fmt.Println("my name is nothing")
   }
}

func is_prime(){
     x := 3
     if x == 2 {
        fmt.Println("Number is even")
     }else if x%2 != 0 {
         fmt.Println("Number is prime")
     }else{
        fmt.Println("Number is even")
     }
}
