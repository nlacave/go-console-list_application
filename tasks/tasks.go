// Package task proporciona metodos para trabajar con slices de tareas.
package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Task representa una tarea o nota. Posee un Identificador (ID), un contenido y puede estar marcada como completa o no, según queramos
type Task struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	Complete bool   `json:"complete"`
}

// ListTask imprime en consola el listado de tareas guardadas en formato .json.
func ListTask(tasks []Task) {
	fmt.Println("Listado de tareas:")
	for _, v := range tasks {
		if v.Complete {
			fmt.Printf("[✔] %v. %v\n", v.ID, v.Content)
		} else {
			fmt.Printf("[ ] %v. %v\n", v.ID, v.Content)
		}
	}
}

// AddTask recibe un string, genera una nueva tarea a partir del mismo, y retorna un slice de tareas con la tarea agregada.
func AddTask(c string, tasks []Task) []Task {
	newTask := Task{AutoTaskID(tasks), c, false}
	tasks = append(tasks, newTask)
	return tasks
}

// DeleteTask recibe un ID y un slice de tareas y retorna un nuevo slice sin incluir la tarea con el ID proporcionado como argumento.
func DeleteTask(id int, tasks []Task) []Task {
	for i, v := range tasks {
		if v.ID == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

// AutoTaskID recibe un slice de tareas y retorna un ID con el valor siguiente al de la última tarea del slice. Si no hay tareas, retorna el valor 1.
func AutoTaskID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}

// SaveTask recibe un archivo y un arreglo de tareas, y se encarga de escribir en formato .json las tareas en el archivo.
func SaveTask(file *os.File, tasks []Task) {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}

// CommpleteTask recibe un ID de tipo entero y un slice de tareas, y retorna el slice de tareas con el atributo complete modificado, en la tarea que posee el ID proporcionado por parametros.
func CompleteTask(id int, tasks []Task) []Task {
	for i, v := range tasks {
		if tasks[i].ID == id {
			tasks[i].Complete = !v.Complete
		}
	}
	return tasks
}
