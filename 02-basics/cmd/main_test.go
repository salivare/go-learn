package main

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	manager := NewTaskManager()

	task := manager.Add("Test Task")

	if task.ID != 0 {
		t.Errorf("Expected task ID 0, got %d", task.ID)
	}
	if task.Title != "Test Task" {
		t.Errorf("Expected task title 'Test Task', got '%s'", task.Title)
	}
	if task.Done {
		t.Errorf("Expected task Done to be false, got true")
	}
	if len(manager.tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(manager.tasks))
	}
}

func TestList(t *testing.T) {
	manager := NewTaskManager()
	testStrings := []string{
		"Test task 1",
		"Test task 2",
		"Test task 3",
		"Test task 4",
		"Test task 5",
	}
	for _, str := range testStrings {
		manager.Add(str)
	}
	tasks := manager.List()

	if len(tasks) != len(testStrings) {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

func TestComplete(t *testing.T) {
	manager := NewTaskManager()
	task := manager.Add("Completable Task")

	if !manager.Complete(task.ID) {
		t.Errorf("Expected task to be marked as complete")
	}
	if !task.Done {
		t.Errorf("Expected task Done to be true")
	}

	// Complete again
	if manager.Complete(task.ID) {
		t.Errorf("Expected task to not be marked complete again")
	}
}

func TestStats(t *testing.T) {
	manager := NewTaskManager()
	manager.Add("Task 1")
	manager.Add("Task 2")
	manager.Complete(1)

	total, done := manager.Stats()
	if total != 2 {
		t.Errorf("Expected total tasks to be 2, got %d", total)
	}
	if done != 1 {
		t.Errorf("Expected done tasks to be 1, got %d", done)
	}
}

func TestAddAndList_Table(t *testing.T) {
	m := NewTaskManager()

	inputs := []string{
		"Test task 1",
		"Test task 2",
		"Test task 3",
	}

	// добавляем все задачи
	for _, title := range inputs {
		task := m.Add(title)
		require.NotNil(t, task)
		require.Equal(t, title, task.Title)
	}

	// проверяем список одним ожиданием
	tasks := m.List()
	require.Len(t, tasks, len(inputs))
	for i, exp := range inputs {
		require.Equal(t, exp, tasks[i].Title)
	}
}

func TestComplete_TableDriven(t *testing.T) {
	m := NewTaskManager()
	// создаём несколько задач
	for i := 1; i <= 3; i++ {
		m.Add("task " + strconv.Itoa(i))
	}

	cases := []struct {
		name     string
		id       int
		wantOK   bool
		wantDone bool
	}{
		{"valid first", 1, true, true},
		{"valid second", 2, true, true},
		{"repeat complete", 1, false, true}, // уже выполнена
		{"nonexistent", 999, false, false},
	}

	for _, tc := range cases {
		t.Run(
			tc.name, func(t *testing.T) {
				ok := m.Complete(tc.id)
				require.Equal(t, tc.wantOK, ok)
				// если задача существует, проверяем её флаг Done
				if task, ok := m.index[tc.id]; ok {
					require.Equal(t, tc.wantDone, task.Done)
				}
			},
		)
	}
}

func TestStats_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		add      []string
		complete []int
		wantTot  int
		wantDone int
	}{
		{"none", nil, nil, 0, 0},
		{"some none done", []string{"a", "b"}, nil, 2, 0},
		{"some with done", []string{"a", "b", "c"}, []int{1, 3}, 3, 2},
	}

	for _, tc := range tests {
		t.Run(
			tc.name, func(t *testing.T) {
				m := NewTaskManager()
				for _, title := range tc.add {
					m.Add(title)
				}
				for _, id := range tc.complete {
					m.Complete(id)
				}
				total, done := m.Stats()
				require.Equal(t, tc.wantTot, total)
				require.Equal(t, tc.wantDone, done)
			},
		)
	}
}
