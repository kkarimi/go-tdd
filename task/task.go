package work

import (
	"errors"
)

type Task struct {
	Title string
}

type TaskManager struct {
	task *Task
}

func NewTask(title string) (*Task, error) {
	if title == "" {
		return nil, errors.New("Empty title")
	}
	return &Task{title}, nil
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
	// return &TaskManager, error
}

func (m *TaskManager) Save(task *Task) error {
	m.task = task
	return nil
}

// nil is empty slice
func (m *TaskManager) All() []*Task {
	return []*Task{m.task}
}
