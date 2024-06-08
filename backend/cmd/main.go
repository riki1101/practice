package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// tests defines the model structure
type Tasks struct {
	gorm.Model
	Title    string
	Priority string
	Status   string
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "P4ssW0rd!"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "todo"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected")
	return db
}

func main() {
	db := gormConnect()
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&Tasks{})
	rows, err := db.Raw("show tables").Rows()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			panic(err.Error())
		}
		fmt.Println(table)
	}

	defer db.Close()
}
