//a demonstaration of who to connect to a database in go

package main

import (
	"encoding/json"
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
	var people []*person

	//not actually my password you tricksy hobbitses
	con, err := sql.Open("mysql", "root:thegame@/golang")
	checkErr(err)

	rows, err := con.Query("select name,age from test")
	checkErr(err)

	for rows.Next() {
		var age int
		var name string
		err = rows.Scan(&name,&age)
		checkErr(err)
		me := &person{Name:name,Age:age}
		people = append(people, me)
	}
	peopleJson, err := json.Marshal(people)
	checkErr(err)
	fmt.Println(string(peopleJson))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
