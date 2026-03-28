package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const cmdAdd = "  add <текст>      — добавить задачу"
const cmdList = "list             — показать все задачи"

// Task — структура, описывающая задачу
type Task struct {
	ID        int       // уникальный идентификатор
	Title     string    // заголовок задачи
	Done      bool      // флаг выполнения
	CreatedAt time.Time // время создания
}

// TaskManager — хранит слайс задач и мапу для быстрого поиска по id
type TaskManager struct {
	tasks []*Task       // слайс указателей на задачи
	index map[int]*Task // мапа id -> *Task
	next  int           // следующий id для новой задачи
}

// NewTaskManager — конструктор
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make([]*Task, 0),
		index: make(map[int]*Task),
		next:  1,
	}
}

// Add — добавляет новую задачу
func (m *TaskManager) Add(title string) *Task {
	t := &Task{
		ID:        m.next,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	m.tasks = append(m.tasks, t) // демонстрация слайса и append
	m.index[t.ID] = t            // демонстрация мапы
	m.next++
	return t
}

// List — возвращает все задачи (показываем порядок в слайсе)
func (m *TaskManager) List() []*Task {
	return m.tasks // возвращаем слайс (копия шоблоков, не элементов)
}

// Complete — помечает задачу выполненной по id
func (m *TaskManager) Complete(id int) bool {
	// проверка наличия в мапе — демонстрация if
	if t, ok := m.index[id]; ok {
		if !t.Done { // if для логики
			t.Done = true
			return true
		}
		return false // уже была выполнена
	}
	return false // не найден id
}

// Stats — считает выполненные и всего
func (m *TaskManager) Stats() (total, done int) {
	for _, t := range m.tasks { // цикл по слайсу
		total++
		if t.Done { // логический оператор
			done++
		}
	}
	return
}

// formatTask — форматирует вывод задачи (пример небольшого хелпера)
func formatTask(t *Task) string {
	status := " "
	if t.Done {
		status = "x"
	}
	// демонстрация форматирования строки и обращения к полям структуры
	return fmt.Sprintf("[%s] %d: %s (создано %s)", status, t.ID, t.Title, t.CreatedAt.Format("2006-01-02 15:04"))
}

func printHelp() {
	fmt.Println("Команды:")
	fmt.Println(cmdAdd)
	fmt.Println(cmdList)
	fmt.Println("  done <id>        — отметить задачу выполненной")
	fmt.Println("  stats            — показать статистику")
	fmt.Println("  help             — показать справку")
	fmt.Println("  exit             — выйти")
}

func main() {
	m := NewTaskManager()

	// Несколько стартовых задач — демонстрация добавления и инициализации
	//m.Add("Купить молоко")
	//m.Add("Прочитать главу про слайсы")

	fmt.Println("Простой менеджер задач — введите help для списка команд")

	// Читаем команды из stdin построчно
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			// EOF или ошибка: выходим
			fmt.Println("\nВыход")
			return
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line) // разбиваем строку на слова
		cmd := strings.ToLower(parts[0])

		switch cmd {

		case "help":
			printHelp()
		case "add":
			// всё, что после "add" объединяем в заголовок
			if len(parts) < 2 {
				fmt.Println("Использование: add <текст задачи>")
				continue
			}
			title := strings.Join(parts[1:], " ")
			t := m.Add(title)
			fmt.Printf("Добавлена задача: %s\n", formatTask(t))
		case "list":
			all := m.List()
			if len(all) == 0 {
				fmt.Println("Список задач пуст")
				continue
			}
			for _, t := range all { // цикл по слайсу
				fmt.Println(formatTask(t))
			}

		case "done":
			if len(parts) != 2 {
				fmt.Println("Использование: done <id>")
				continue
			}
			id, err := strconv.Atoi(parts[1]) // демонстрация strconv
			if err != nil {
				fmt.Println("Некорректный id:", parts[1])
				continue
			}
			if m.Complete(id) {
				fmt.Printf("Задача %d помечена выполненной\n", id)
			} else {
				fmt.Printf("Задача %d не найдена или уже выполнена\n", id)
			}
		case "stats":
			total, done := m.Stats()
			fmt.Printf("Всего: %d, выполнено: %d\n", total, done)
		case "exit", "quit":
			fmt.Println("Bye")
			return
		default:
			// демонстрация ветвления: непонятная команда
			fmt.Println("Неизвестная команда. Введите help.")
		}
	}
}
