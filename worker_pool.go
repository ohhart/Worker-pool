package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker обрабатывает данные из канала
type Worker struct {
	id       int
	jobChan  chan string
	quitChan chan bool
}

// NewWorker создает новый воркер
func NewWorker(id int, jobChan chan string) *Worker {
	return &Worker{
		id:       id,
		jobChan:  jobChan,
		quitChan: make(chan bool),
	}
}

// запуск воркера в отдельной горутине
func (w *Worker) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case job, ok := <-w.jobChan:
				if !ok {
					fmt.Printf("Worker %d: Job channel closed, exiting.\n", w.id)
					return
				}
				fmt.Printf("Worker %d processed: %s\n", w.id, job)
			case <-w.quitChan:
				fmt.Printf("Worker %d stopped\n", w.id)
				return
			}
		}
	}()
}

// посылает сигнал стоп воркеру
func (w *Worker) Stop() {
	w.quitChan <- true
}

func main() {
	var wg sync.WaitGroup
	jobChan := make(chan string, 100)
	workers := make([]*Worker, 0)

	// добавить новый воркер
	addWorker := func() {
		newWorker := NewWorker(len(workers)+1, jobChan)
		newWorker.Start(&wg)
		workers = append(workers, newWorker)
	}

	// добавляем начальный воркер
	for i := 0; i < 3; i++ {
		addWorker()
	}

	// симуляция работы
	go func() {
		for i := 1; i <= 10; i++ {
			jobChan <- fmt.Sprintf("Job %d", i)
			time.Sleep(time.Second)
		}
		close(jobChan)
	}()

	// добавляем и удаляем воркеры динамически
	time.Sleep(2 * time.Second)
	fmt.Println("Adding a new worker")
	addWorker()

	time.Sleep(3 * time.Second)
	fmt.Println("Stopping a worker")
	if len(workers) > 0 {
		workerToStop := workers[0]
		workerToStop.Stop()
		workers = workers[1:]
	}

	wg.Wait() // ждем завершения всех воркеров
	fmt.Println("All workers have finished.")
}
