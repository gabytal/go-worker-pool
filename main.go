package main

import (
	"worker-pool/worker"
)

func main() {
	// Create new Tasks
	deletionTasks := []pool.DeletionTask{
		{Name: "task1"},
		{Name: "task2"},
		{Name: "task3"},
		{Name: "task4"},
	}

	emailTasks := []pool.EmailTask{
		{Email: "email1", Body: "body1"},
		{Email: "email2", Body: "body2"},
		{Email: "email3", Body: "body3"},
		{Email: "email4", Body: "body4"},
	}

	var tasks []pool.Task

	// Append tasks to tasks slice
	for _, task := range deletionTasks {
		del := task
		tasks = append(tasks, &del)
	}

	for _, task := range emailTasks {
		emailTask := task
		tasks = append(tasks, &emailTask)
	}

	wp := pool.WorkerPool{
		Tasks:       tasks,
		Concurrency: 1,
		TasksChan:   make(chan pool.Task),
	}

	wp.Run()
}
