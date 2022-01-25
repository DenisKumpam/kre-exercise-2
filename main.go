package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Car struct {
	Brand     string
	Model     string
	Year      int
}

func (c *Car) WriteJSON(w io.Writer) error  {
	js, err := json.Marshal(c)
	if err != nil{
		return err
	}
	_, err = w.Write(js)
	return err
}

func main()  {
	c := &Car{Brand: "Lada", Model: "Granta", Year: 2016}

	var buf bytes.Buffer
	err := c.WriteJSON(&buf)
	if err  != nil{
		log.Fatal(err)
	}

	f, err := os.Create("exercise-2")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = c.WriteJSON(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", c)
}