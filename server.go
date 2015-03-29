package main

import (
	"database/sql"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
	"log"
	//	"os"
)

func main() {

	log.Println("main start")

	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, content from todo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var content string
		rows.Scan(&id, &content)
		fmt.Println(id, content)
	}
	rows.Close()

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Get("/todo/list.json", func(r render.Render) {

		var todoList []Todo

		todoList = append(todoList, Todo{
			1, "study golang",
		},
		)

		todoList = append(todoList, Todo{
			2, "study Engish",
		},
		)

		r.JSON(200, todoList)
	})

	m.Get("/hello", func(r render.Render) {
		r.HTML(200, "hello", "you")
	})

	m.Run()

}

type Todo struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}
