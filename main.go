package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
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

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}
	fmt.Println("Escriba una opción en consola y presione la tecla ENTER en su teclado.")
	for {
		OptionsMenu()
		reader := bufio.NewReader(os.Stdin)
		entrada, _ := reader.ReadString('\n')
		entrada = strings.TrimSpace(entrada)
		palabras := strings.Fields(entrada)
		if len(palabras) > 0 {
			if palabras[0] == "list" {
				if len(tasks) == 0 {
					fmt.Println("No hay tareas.")
				} else {
					task.ListTask(tasks)
				}
			} else if palabras[0] == "add" {
				fmt.Println("Cuál es tu tarea?")
				entrada, _ = reader.ReadString('\n')
				entrada = strings.TrimSpace(entrada)
				if entrada != "" {
					tasks = task.AddTask(entrada, tasks)
					task.SaveTask(file, tasks)
				} else {
					fmt.Println("No has ingresado ningún dato.")
				}

			} else if palabras[0] == "delete" {
				if len(tasks) == 0 {
					fmt.Println("No hay tareas.")
					continue
				}
				fmt.Println("Que tarea quieres eliminar?")
				task.ListTask(tasks)
				entrada, _ = reader.ReadString('\n')
				entrada = strings.TrimSpace(entrada)
				arr := strings.Fields(entrada)
				if len(arr) > 1 {
					fmt.Println("La entrada no es valida. Debe ingresar un ID único perteneciente a la tarea a eliminar. No son validos multiples números ni letras entre espacios.")
				} else {
					newValue, err := strconv.Atoi(entrada)
					if err != nil {
						fmt.Println("La entrada no es valida. Debes ingresar un ID (valor de tipo Integer. Ej:´459´) de la tarea a eliminar.")
					} else {
						tasks = task.DeleteTask(newValue, tasks)
						task.SaveTask(file, tasks)
					}
				}
			} else if palabras[0] == "complete" {
				if len(tasks) == 0 {
					fmt.Println("No hay tareas.")
					continue
				}
				fmt.Println("Que tarea quieres marcar/desmarcar?")
				task.ListTask(tasks)
				entrada, _ = reader.ReadString('\n')
				entrada = strings.TrimSpace(entrada)
				arr := strings.Fields(entrada)
				if len(arr) > 1 {
					fmt.Println("La entrada no es valida. Debe ingresar un ID único perteneciente a la tarea a marcar/desmarcar. No son validos multiples números ni letras entre espacios.")
				} else {
					newValue, err := strconv.Atoi(entrada)
					if err != nil {
						fmt.Println("La entrada no es valida. Debes ingresar un ID (valor de tipo Integer. Ej:´459´) de la tarea a marcar/desmarcar.")
					} else {
						tasks = task.CompleteTask(newValue, tasks)
						task.SaveTask(file, tasks)
					}
				}
			} else if palabras[0] == "exit" {
				os.Exit(0)
			}
			file.Stat()
		}
	}
}

func OptionsMenu() {
	fmt.Println("Menú de opciones: list | add | delete | complete | exit")
}
