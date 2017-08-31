package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    person := Person{"rafael", 25}
    s := reflect.ValueOf(&person).Elem()
    typeOfPerson := s.Type()
    for i := 0; i < s.NumField(); i++ {
        field := s.Field(i)
        fmt.Println(typeOfPerson.Field(i).Name)
        fmt.Println(field.Type())
        fmt.Println(field.Interface())
    }
}
