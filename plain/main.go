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

type Pack struct {
	V interface{}
}

func main() {
	buf := &bytes.Buffer{}

	enc := gob.NewEncoder(buf)
	gob.Register(&A{})
	gob.Register(&B{})

	enc.Encode(&Pack{&A{42}})
	enc.Encode(&Pack{&B{"hello"}})

	dec := gob.NewDecoder(buf)
	for {
		p := &Pack{}
		if err := dec.Decode(p); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("V:%+v / Type:%T\n", p.V, p.V)
	}
}
