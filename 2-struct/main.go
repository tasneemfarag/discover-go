package main

import (
    "fmt"
)

type user struct {
    Name string
    DOB string
    City  string
}
func (u user) PrintName() string {
  return u.Name
  fmt.Println("Hello", u.Name)
}

func (u user) PrintAll(string, string, string)  {
year := u.DBO[9:len(u.DOB)]
  return u.Name
  return u.City
  return u.DOB
  fmt.Println(u.Name,"who was born in"u.City,"would be",u.DOB, "years old today.")
}

func main(){
u := user{"Betty Holberton", "March 7, 1917", "Philadelphia"}
u.PrintName()
//fmt.Println("Hello", u.Name)
u.PrintAll()
//fmt.Println(u.Name,"who was born in"u.City,"would be"u)
}
