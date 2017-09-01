package main

import (
    "reflect"
    "fmt"
)

func main() {
    Init()
    for _, model := range models {
        s := reflect.New(model).Elem()
        typeOfPerson := s.Type()
        for i := 0; i < s.NumField(); i++ {
            field := s.Field(i)
            fmt.Println(typeOfPerson.Field(i).Name)
            fmt.Println(field.Type())
            fmt.Println(field.Interface())
        }
    }
}
