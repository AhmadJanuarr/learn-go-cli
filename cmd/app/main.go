package main

import (
	"bufio"
	"fmt"
	"golang-here/internal/task"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func ShowMenu() {
	fmt.Println("Selamat datang di aplikasi todo list")
	fmt.Println("1. Tambah tugas")
	fmt.Println("2. Lihat semua tugas")
	fmt.Println("3. Tugas selesai")
	fmt.Println("4. Tugas belum selesai")
	fmt.Println("5. Keluar")
	fmt.Println("--------------------------------")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		ClearScreen()
		ShowMenu()
		input, _ := reader.ReadString('\n')
		selectOption := strings.TrimSpace(input)

		switch selectOption {
		case "1":
			fmt.Println("Memilih menu Tambah tugas")
			var inputTitle string
			fmt.Print("Masukan judul tugas: ")
			inputTitle, _ = reader.ReadString('\n')
			task.AddTodo(strings.TrimSpace(inputTitle))
			fmt.Println("--------------------------------")
		case "2":
			task.ShowAllTodo()
			fmt.Println("\nTekan Enter untuk kembali ke menu...")
			reader.ReadString('\n')
		case "3":
			var menuItems []string
			var activeTodos []*task.Todo

			for _, todo := range task.StorageTask {
				if !todo.IsCompleted {
					label := fmt.Sprintf("[%d] %s", todo.Id, todo.Title)
					menuItems = append(menuItems, label)
					activeTodos = append(activeTodos, todo)
				}
			}

			if len(menuItems) == 0 {
				fmt.Println("wah, semua tugas sudah selesai! lanjutkan kerjaanmu")
				break
			}

			prompt := promptui.Select{
				Label: "Pilih tugas yang mau diselesaikan (Gunakan Panah Atas dan Panah Bawah untuk memilih)",
				Items: menuItems,
				Size:  5,
			}

			index, result, err := prompt.Run()

			if err != nil {
				fmt.Println("Batal memilih tugas")
				break
			}

			selectedTodo := activeTodos[index]
			selectedTodo.IsCompleted = true
			fmt.Printf("Berhasil mengubah status tugas %s menjadi selesai\n", result)

			reader.ReadString('\n')

		case "4":
			task.ShowInCompletedTodo()
			fmt.Println("\nTekan Enter untuk kembali ke menu...")
			reader.ReadString('\n')
		case "5":
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}

}
