package main
import  (
	"fmt"
	"bufio"
	"os"
)
func main(){
     user_input()
     io()
}

func user_input() {
    fmt.Println("Enter Your Input: ")
    var  input string
    fmt.Scanln(&input)
    fmt.Println("Your input is: ", input)
}

func io() {
	names := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
          fmt.Print("Enter your name: ")
	  scanner.Scan()
	  text := scanner.Text()
	  if len(text) != 0 {
              fmt.Println(text)
	      names = append(names, text)
          }else {
              break
          }

       }
}
