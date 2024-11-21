package engine

import (
	"errors"
	"fmt"
)

type Column struct {
	Name string
	Type string
}

type Schema struct {
	Columns []Column
}

type Row struct {
	Values []interface{}
}

type Table struct {
	Name   string
	Schema Schema
	Rows   []Row
}

type Database struct {
	Tables map[string]*Table
}

func NewDatabase() *Database {
	return &Database{
		Tables: make(map[string]*Table),
	}
}

func (db *Database) CreateTable(name string, schema Schema) error {
	if _, exists := db.Tables[name]; exists {
		return errors.New("table already exists")
	}
	db.Tables[name] = &Table{
		Name:   name,
		Schema: schema,
		Rows:   []Row{},
	}
	return nil
}

func (db *Database) InsertRow(tableName string, values []interface{}) error {
	table, exists := db.Tables[tableName]
	if !exists {
		return errors.New("table does not exist")
	}
	if len(values) != len(table.Schema.Columns) {
		return errors.New("row does not match table schema")
	}
	table.Rows = append(table.Rows, Row{Values: values})
	return nil
}

func (db *Database) SelectRows(tableName string) ([]Row, error) {
	table, exists := db.Tables[tableName]
	if !exists {
		return nil, errors.New("table does not exist")
	}
	return table.Rows, nil
}

func (db *Database) DebugPrint() {
	for name, table := range db.Tables {
		fmt.Printf("Table: %s\n", name)
		fmt.Printf("Schema: ")
		for _, col := range table.Schema.Columns {
			fmt.Printf("- %s (%s)\n", col.Name, col.Type)
		}
		fmt.Println("Rows: ")
		for _, row := range table.Rows {
			fmt.Println(row.Values)
		}
	}
}
