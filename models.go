package main

import (
    "reflect"
)

var models = make(map[string]reflect.Type)

type Person struct {
    Name string
    Age  int
}

type Animal struct {
    Race string
    Legs int
}

func Init() {
    models["Person"] = reflect.TypeOf(Person{})
    models["Animal"] = reflect.TypeOf(Animal{})
}
