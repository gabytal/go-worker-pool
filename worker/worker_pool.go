import (
        "fmt"
        "sync"
        "time"
)

// Task def
type Task interface {
        Process()
}

type EmailTask struct {
        Email string
        Body  string
}

type DeletionTask struct {
        Name string
}

// Process tasks

func (t *DeletionTask) Process() {
        fmt.Printf("Processing DeletionTask process %s \n", t.Name)
        time.Sleep(2 * time.Second)
}

func (t *EmailTask) Process() {
        fmt.Printf("Processing EmailTask process %s \n", t.Email)
        time.Sleep(2 * time.Second)
}

type WorkerPool struct {
        Tasks       []Task
        Concurrency int
        TasksChan   chan Task
        wg          sync.WaitGroup
}

func (wp *WorkerPool) worker() {
        for task := range wp.TasksChan {
                task.Process()
                wp.wg.Done()
        }
}

func (wp *WorkerPool) Run() {

        for i := 0; i < wp.Concurrency; i++ {
                go wp.worker()
        }

        for _, task := range wp.Tasks {
                wp.AddTask(task)
        }

        wp.wg.Wait()

}

func (wp *WorkerPool) AddTask(task Task) {
        wp.wg.Add(1)
        wp.TasksChan <- task
}
