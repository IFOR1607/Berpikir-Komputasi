
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

var (
	tasks  []Task
	nextID = 1
	reader = bufio.NewReader(os.Stdin)
)

func main() {
	fmt.Println("================================")
	fmt.Println("       GO TASK MANAGER")
	fmt.Println("================================")

	for {
		showMenu()

		choice := readInput("Pilih menu: ")

		switch choice {
		case "1":
			addTask()
		case "2":
			showTasks()
		case "3":
			completeTask()
		case "4":
			deleteTask()
		case "5":
			showStats()
		case "0":
			fmt.Println("Terima kasih sudah menggunakan aplikasi!")
			return
		default:
			fmt.Println("Menu tidak valid. Coba lagi.")
		}

		fmt.Println()
	}
}

func showMenu() {
	fmt.Println("\n--- MENU ---")
	fmt.Println("1. Tambah tugas")
	fmt.Println("2. Lihat semua tugas")
	fmt.Println("3. Selesaikan tugas")
	fmt.Println("4. Hapus tugas")
	fmt.Println("5. Statistik tugas")
	fmt.Println("0. Keluar")
}

func readInput(prompt string) string {
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')
	if err != nil && len(input) == 0 {
		return ""
	}

	return strings.TrimSpace(input)
}

func addTask() {
	title := readInput("Nama tugas: ")

	if title == "" {
		fmt.Println("Nama tugas tidak boleh kosong.")
		return
	}

	task := Task{
		ID:        nextID,
		Title:     title,
		Completed: false,
	}

	tasks = append(tasks, task)
	nextID++

	fmt.Println("Tugas berhasil ditambahkan!")
}

func showTasks() {
	if len(tasks) == 0 {
		fmt.Println("Belum ada tugas.")
		return
	}

	fmt.Println("\n--- DAFTAR TUGAS ---")

	for _, task := range tasks {
		status := "Belum selesai"

		if task.Completed {
			status = "Selesai"
		}

		fmt.Printf("[%d] %s - %s\n",
			task.ID, task.Title, status)
	}
}

func completeTask() {
	showTasks()

	if len(tasks) == 0 {
		return
	}

	input := readInput("Masukkan ID tugas: ")
	id, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("ID harus berupa angka.")
		return
	}

	for i := range tasks {
		if tasks[i].ID == id {
			if tasks[i].Completed {
				fmt.Println("Tugas sudah selesai.")
				return
			}

			tasks[i].Completed = true
			fmt.Println("Tugas berhasil diselesaikan!")
			return
		}
	}

	fmt.Println("Tugas tidak ditemukan.")
}

func deleteTask() {
	showTasks()

	if len(tasks) == 0 {
		return
	}

	input := readInput("Masukkan ID tugas yang dihapus: ")
	id, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("ID harus berupa angka.")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Tugas berhasil dihapus!")
			return
		}
	}

	fmt.Println("Tugas tidak ditemukan.")
}

func showStats() {
	total := len(tasks)
	completed := 0

	for _, task := range tasks {
		if task.Completed {
			completed++
		}
	}

	pending := total - completed

	fmt.Println("\n--- STATISTIK ---")
	fmt.Printf("Total tugas     : %d\n", total)
	fmt.Printf("Selesai         : %d\n", completed)
	fmt.Printf("Belum selesai   : %d\n", pending)

	if total > 0 {
		percentage := float64(completed) /
			float64(total) * 100

		fmt.Printf("Progress        : %.1f%%\n",
			percentage)
	}
}