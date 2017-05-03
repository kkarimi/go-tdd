package work

import (
	"testing"
)

func TestNewTask(t *testing.T) {
	task, err := NewTask("TDD")
	if err != nil {
		t.Fatalf("Unexpected error, %v", err)
	}
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
