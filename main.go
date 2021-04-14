package main
 
import (
    "os"
    "fmt"
)
 
func main() {

	argLength := len(os.Args[1:])
	if argLength != 2  {
		fmt.Println("Incorrect number of args. Found %d, but expected 2", argLength);
		return
	}

	inputString := ""
	charToUse := ""
	returnString := ""
	
	// the first argument is always program name
	charToUse = string(os.Args[1])
	inputString = string(os.Args[2])
	returnString = ""

	for i := 0; i < len(inputString); i++  {
		returnString += string(inputString[i]) + charToUse
	}
	//fmt.Println();
	fmt.Println(returnString);

}