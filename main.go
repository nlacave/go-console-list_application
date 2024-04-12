package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	task "github.com/nlacave/go-cli-prop/tasks"
)

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	var tasks []task.Task

	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	for {
		OptionsMenu()
		reader := bufio.NewReader(os.Stdin)
		entrada, _ := reader.ReadString('\n')
		entrada = strings.TrimSpace(entrada)

		palabras := strings.Fields(entrada)

		if palabras[0] == "list" {
			if info.Size() != 0 {
				task.ListTask(tasks)
			} else {
				fmt.Println("No hay tareas")
			}
		} else if palabras[0] == "add" {
			fmt.Println("Cuál es tu tarea?")
			entrada, _ := reader.ReadString('\n')
			tasks = task.AddTask(entrada, tasks)
			task.SaveTask(file, tasks)
			file.Stat()

		} else if palabras[0] == "delete" {
			fmt.Println("Que tarea quieres eliminar?")
			task.ListTask(tasks)
			var idTask int
			idTask, err = fmt.Scan(&idTask)
			if err != nil {
				panic(err)
			}
			tasks = task.DeleteTask(idTask, tasks)
			task.SaveTask(file, tasks)
			file.Stat()
		}
	}
}

func OptionsMenu() {
	fmt.Println("Opciones: list | add | delete | update | complete | exit")
}
