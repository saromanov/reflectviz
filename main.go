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

// reflectviz implements main structire for walking
// thought objeject
type reflectviz struct {
	level int
	node  map[string]string
}

func (r *reflectviz) reflectMethod(i interface{}) {
	data := reflect.ValueOf(i)
	if data.CanAddr() {
		data = data.Elem()
	}
	fmt.Println(data.Kind())
	r.level++
	switch data.Kind() {
	case reflect.Struct:
		r.showStruct(data)
	case reflect.String:
		r.showString(data)
	case reflect.Ptr:
		r.showPtr(data)
	}
}

func (r *reflectviz) showPtr(value reflect.Value) {
	fmt.Println("Result: ", value.Elem().Kind())
}

func (r *reflectviz) showStruct(value reflect.Value) {
	fmt.Println(value.Type())
	for i := 0; i < value.Type().NumField(); i++ {
		inValue := value.Field(i)
		r.reflectMethod(inValue)

	}
	fmt.Println(value.Type().NumField())
}

func (r *reflectviz) showString(value reflect.Value) {
	fmt.Println(value.String())
}
func main() {
	str := "bar"
	t := test{
		foo: &str,
		bar: 10,
	}
	r := &reflectviz{node: map[string]string{}}
	r.reflectMethod(t)
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
