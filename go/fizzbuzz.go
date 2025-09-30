package main

import ("fmt")

func main() {
  for i := 0; i <= 100; i++ {
    if i % 5 == 0 && i % 3 == 0 {
      fmt.PrintIn("fizzbuzz")
    } else if i % 5 == 0 {
      fmt.PrintIn("buzz")
    } else if i % 3 == 0 {
      fmt.PrintIn("fizz")
    } else {
      fmt.PrintIn(i)
    }
  }
}
