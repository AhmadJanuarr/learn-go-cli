package task

import "fmt"

var LastID int = 0

func AddTodo(title string) {
	todo := Todo{
		Id:          LastID + 1,
		Title:       title,
		IsCompleted: false,
	}
	StorageTask = append(StorageTask, &todo)
	fmt.Println("Berhasil memasukan data", todo.Title)
}
func ShowAllTodo() {
	status := "[ ]"
	for _, todo := range StorageTask {
		if todo.IsCompleted {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s \n", todo.Id, status, todo.Title)
	}
}

func ShowCompletedTodo() {
	for _, todo := range StorageTask {
		if todo.IsCompleted {
			fmt.Println("ID: ", todo.Id)
			fmt.Println("Title: ", todo.Title)
			fmt.Println("IsCompleted: ", todo.IsCompleted)
			fmt.Println("--------------------------------")
		}
		fmt.Println("Tidak ada tugas yang selesai")
	}
}

func ShowInCompletedTodo() {
	for _, todo := range StorageTask {
		if !todo.IsCompleted {
			fmt.Println("ID: ", todo.Id)
			fmt.Println("Title: ", todo.Title)
			fmt.Println("IsCompleted: ", todo.IsCompleted)
			fmt.Println("--------------------------------")
		} else {
			fmt.Println("Tidak ada tugas yang belum selesai")
		}
	}
}

func MarkAsDone(id int) {
	for _, todo := range StorageTask {
		if todo.Id == id {
			todo.IsCompleted = true
			fmt.Println("Berhasil mengubah status tugas menjadi selesai")
			return
		}
	}
}
