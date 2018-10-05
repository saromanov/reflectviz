package main

import (
	"fmt"
	"reflect"

	"github.com/awalterschulze/gographviz"
)

type test struct {
	foo *string
	bar int
}

func reflectMethod(i interface{}) {
	data := reflect.ValueOf(i)
	if data.CanAddr() {
		data = data.Elem()
	}
	fmt.Println(data.Kind())

	switch data.Kind() {
	case reflect.Struct:
		showStruct(data)
	case reflect.String:
		showString(data)
	}
}

func showStruct(value reflect.Value) {
	fmt.Println(value.Type())

	for i := 0; i < value.Type().NumField(); i++ {
		inValue := value.Field(i)
		reflectMethod(inValue)

	}
	fmt.Println(value.Type().NumField())
}

func showString(value reflect.Value) {
	fmt.Println(value.String())
}
func main() {
	str := "bar"
	t := test{
		foo: &str,
		bar: 10,
	}
	reflectMethod(t)
	graphAst, _ := gographviz.ParseString(`digraph G {}`)
	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		panic(err)
	}
	graph.AddNode("G", "a", nil)
	graph.AddNode("G", "b", nil)
	graph.AddEdge("a", "b", true, nil)
	output := graph.String()
	fmt.Println(output)
}
