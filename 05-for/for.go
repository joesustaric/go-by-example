package main

import "fmt"

func main() {

  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  //classic style loop
  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }

  //loop without exit condition
  for {
    fmt.Println("loopy loop")
    break
  }
}
