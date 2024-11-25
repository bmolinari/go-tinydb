package main

import (
	"fmt"

	"github.com/bmolinari/go-tinydb/internal/engine"
)

func main() {
	db := engine.NewDatabase()

	schema := engine.Schema{
		Columns: []engine.Column{
			{Name: "id", Type: "int"},
			{Name: "name", Type: "string"},
		},
	}
	db.CreateTable("users", schema)

	db.InsertRow("users", []interface{}{1, "Alice"})
	db.InsertRow("users", []interface{}{2, "Bob"})
	fmt.Println(db.InsertRow("users", []interface{}{"invalid", "Alice"}))
	fmt.Println(db.InsertRow("users", []interface{}{3, 12345}))

	rows, err := db.SelectRows("users", []engine.Condition{
		{Column: "id", Operator: "=", Value: 1},
	})
	fmt.Println("Query =", rows, err)

	rows, err = db.SelectRows("users", []engine.Condition{
		{Column: "id", Operator: "!=", Value: 1},
	})
	fmt.Println("Query !=", rows, err)

	rows, err = db.SelectRows("users", []engine.Condition{
		{Column: "id", Operator: ">", Value: 1},
	})
	fmt.Println("Query >", rows, err)

	rows, err = db.SelectRows("users", []engine.Condition{
		{Column: "id", Operator: "<", Value: 2},
	})
	fmt.Println("Query <", rows, err)
}
