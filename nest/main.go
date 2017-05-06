package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type A struct {
	Number int
}

type B struct {
	Character string
}

type Root struct {
	A        *A
	Children []*Child
}

type Child struct {
	B     *B
	Float float32
}

type Pack struct {
	V interface{}
}

func main() {
	buf := &bytes.Buffer{}

	enc := gob.NewEncoder(buf)
	gob.Register(&Root{})

	/* Not need */
	// gob.Register(&A{})
	// gob.Register(&B{})
	// gob.Register(&Child{})

	a1 := &A{10}
	b1 := &B{"foo"}
	b2 := &B{"bar"}
	b3 := &B{"baz"}

	children := []*Child{
		&Child{b1, 1.2},
		&Child{b2, 3.4},
		&Child{b3, 5.6},
	}

	root := &Root{
		A:        a1,
		Children: children,
	}

	enc.Encode(&Pack{root})

	dec := gob.NewDecoder(buf)
	for {
		p := &Pack{}
		if err := dec.Decode(p); err != nil {
			fmt.Println(err)
			break
		}
		dump(p.V)
	}
}

func dump(value interface{}) {
	switch v := value.(type) {
	case *A, *B, float32:
		fmt.Printf("V:%+v / Type:%T\n", v, v)
	case *Root:
		dump(v.A)
		for _, c := range v.Children {
			dump(c)
		}
	case *Child:
		dump(v.B)
		dump(v.Float)
	}
}
