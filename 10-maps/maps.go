package main

import "fmt"

func main() {

  m := make(map[string]int)

  m["k1"] = 7
  m["k2"] = 13

  fmt.Println("map:",m)

  v1 := m["k1"]
  fmt.Println("v1: ", v1)

  fmt.Println("len:", len(m))

  delete(m, "k2")

  fmt.Println("map:",m)

  _, k2_present := m["k2"]
  k2, k1_present := m["k1"]
  fmt.Println("prs:", k2_present)
  fmt.Println("k1prs:",k1_present)
  fmt.Println("k2:",k2)

  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", n)

}
