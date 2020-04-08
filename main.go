//Name: Vincent G.
//Date: 4/1/2020
//Description: Grade report



package main

import "fmt"

func Grade(grade int) {

  //print out the approprate grade depending on the users input
  if grade >=90{
    fmt.Println("Great Job you got an A!")
  }else if grade >=80 && grade <90{
    fmt.Println("Good job you got a B!")
  }else if grade >=70 && grade <80{
    fmt.Println("You got a C!")
  }else if grade >=60 && grade <70{
    fmt.Println("You got a D.")
  }else{
    fmt.Println("Sorry but you Failed.")
  } 
}

func main() {
  //declare variable type integer for "i"
  var i int
  //ask for grade percentage, scan for i, call Grade(i)
  fmt.Println("Please Enter your grade percentage: ")
  fmt.Scanln(&i)
  fmt.Println()
  Grade(i)
}