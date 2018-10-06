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
	graph *gographviz.Graph
}

func (r *reflectviz) reflectMethod(i interface{}) {
	r.graph = gographviz.NewGraph()
	r.reflectValue(reflect.ValueOf(i))
}

func (r *reflectviz) reflectValue(data reflect.Value) {
	nodeName := fmt.Sprintf("%s", data.Kind())
	_, ok := r.node[nodeName]
	if ok {
		return
	}
	r.node[nodeName] = "a"
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
	r.reflectValue(value.Elem())
}

func (r *reflectviz) showStruct(value reflect.Value) {
	fmt.Println(value.Type())
	for i := 0; i < value.Type().NumField(); i++ {
		inValue := value.Field(i)
		r.reflectValue(inValue)

	}
	fmt.Println(value.Type().NumField())
}

func (r *reflectviz) showString(value reflect.Value) {
	fmt.Println(value.String())
}

// createNode provides creating of teh new node
func (r *reflectviz) createNode(value reflect.Value) error {
	return r.graph.AddNode("G", value.String(), nil)
}

// showGraph returns result graph
func (r *reflectviz) showGraph(value reflect.Value) {
	output := r.graph.String()
	fmt.Println(output)
}

func main() {
	str := "bar"
	t := test{
		foo: &str,
		bar: 10,
	}
	r := &reflectviz{node: map[string]string{}}
	r.reflectMethod(t)
}
