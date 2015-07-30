//a demonstaration of who to connect to a database in go

package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type person struct {
	Name string
	Age int
}

func printPerson(who *person) {
	fmt.Printf(who.Name + " " )
	fmt.Println(who.Age)
}

func main() {
	//not actually my password you tricksy hobbitses
	con, err := sql.Open("mysql", "root:somepassword@/golang")
	checkErr(err)

	rows, err := con.Query("select name,age from test where name=?", "Ryan")
	checkErr(err)

	for rows.Next() {
		var age int
		var name string
		err = rows.Scan(&name,&age)
		checkErr(err)
		me := &person{Name:name,Age:age}
		printPerson(me)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
