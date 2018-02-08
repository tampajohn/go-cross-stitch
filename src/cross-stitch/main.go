package main

import (
  "fmt"
  "os"

  "cross-stitch/convert"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Hello world!")
    os.Exit(0)
  }
  args := os.Args[1:]

  img, err := convert.Open(args[0])
  if err != nil {
    panic(err)
  }
  
  //grey, err := convert.Greyscale(img, "output.png")
  //if err != nil {
  //  panic(err)
  //}
  
  _, _ = convert.DMC(img)
}