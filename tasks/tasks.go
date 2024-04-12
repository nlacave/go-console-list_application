package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	Complete bool   `json:"complete"`
}

func ListTask(tasks []Task) {
	for _, v := range tasks {
		if v.Complete {
			fmt.Println("[✔]", v.ID, v.Content)
		} else {
			fmt.Println("[ ]", v.ID, v.Content)
		}
	}
}

func AddTask(c string, tasks []Task) []Task {
	newTask := Task{AutoTaskID(tasks), c, false}
	tasks = append(tasks, newTask)
	return tasks
}

func DeleteTask(id int, tasks []Task) []Task {
	for _, v := range tasks {
		if v.ID == id {
			return append(tasks[:id], tasks[id+1:]...)
		}
	}
	return tasks
}

func AutoTaskID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}

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
