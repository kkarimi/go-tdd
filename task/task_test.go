package work

import (
	"testing"
)

func newTaskOrFatal(t *testing.T, title string) *Task {
	task, err := NewTask(title)
	if err != nil {
		t.Fatalf("Unexpected error, %v", err)
	}
	return task
}

func TestNewTask(t *testing.T) {
	task := newTaskOrFatal(t, "TDD")
	if task.Title != "TDD" {
		t.Errorf("Expected TDD, got %v", task.Title)
	}
}

func TestTaskWithEmptyTitle(t *testing.T) {
	_, err := NewTask("")
	if err == nil {
		t.Errorf("Expected 'empty title' error, got nil")
	}
}

func TestSaveTask(t *testing.T) {
	task := newTaskOrFatal(t, "TDD")

	m := NewTaskManager()
	err := m.Save(task)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestSaveTaskAndRetrieve(t *testing.T) {
	task := newTaskOrFatal(t, "TDD")
	m := NewTaskManager()
	err := m.Save(task)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	all := m.All()
	if len(all) != 1 {
		t.Fatalf("Expected 1 task, got %v", len(all))
	}
	if all[0].Title != task.Title {
		t.Fatalf("Expected task title %q, got %v", task.Title, all[0].Title)
	}
}
