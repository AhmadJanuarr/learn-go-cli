# GoDo CLI 🚀

A simple, interactive Command Line Interface (CLI) Todo List application built with **Golang**.
This project serves as a personal exercise to master Golang fundamental concepts.

## 📚 Learning Objectives

This project implements core Golang features:

- **Structs & Interfaces**: For data modeling and abstraction.
- **Pointers**: Efficient memory management and state updates.
- **Slices**: Dynamic storage for tasks.
- **Packages**: Modular code structure (`cmd` vs `internal`).
- **Input/Output**: Using `bufio` for capturing user input.
- **Error Handling**: Proper error checking logic.

## ✨ Features

- [x] **Add Task**: Create new todo items easily.
- [x] **List Tasks**: View all tasks with their current status.
- [x] **Filter Views**: See only "Completed" or "Incomplete" tasks.
- [x] **Mark as Done**: Update task status by ID.
- [x] **Persistent Loop**: Interactive menu that doesn't exit until you say so.
- [x] **Clear Screen**: Keeps your terminal clean after every action.

## 🛠️ Project Structure

todo-app/
├── cmd/
│ └── app/
│ └── main.go # Entry point & Menu Logic
├── internal/
│ └── task/ # Business Logic & Data Structures
│ ├── entity.go
│ ├── service.go
│ └── storage.go
└── go.mod

## Usage Example

=== MENU ===

1. Tambah tugas
2. Lihat semua tugas
3. Tugas selesai
4. Keluar

Choice: 1
Masukan judul tugas: Belajar Golang Pointer
-> Berhasil memasukan data!
