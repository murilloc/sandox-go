package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type Task struct {
	Name string
	Done bool
}

func (task Task) Create() {
	// TODO implementar o método
}

func main() {
	http.HandleFunc("/hello", HelloHandle)
	http.HandleFunc("/tasks", TaskHandler)
	http.ListenAndServe(":8887", nil)

}

func HelloHandle(writer http.ResponseWriter, reader *http.Request) {
	id := uuid.NewV4().String()
	fmt.Fprintf(writer, "Hello Full Cycle, my ID is: %s", id)
}

func TaskHandler(writer http.ResponseWriter, reader *http.Request) {

	tasks := []Task{
		{"Tarefa-1", false},
		{"Tarefa-2", true},
	}

	// como fazer uma chamada no método create do struct task
	x := Task{}
	x.Create()

	tmpl := template.Must(template.ParseFiles("server/tasks.html"))
	tmpl.Execute(writer, tasks)
}
