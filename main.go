package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

type Doctor struct {
    ID   int
    Name string
}

type DoctorFactory struct {
    doctors map[int]*Doctor
}

func NewDoctorFactory() *DoctorFactory {
    return &DoctorFactory{
        doctors: make(map[int]*Doctor),
    }
}

func (df *DoctorFactory) CreateDoctor(id int, name string) *Doctor {
    doctor := &Doctor{
        ID:   id,
        Name: name,
    }
    df.doctors[id] = doctor
    return doctor
}

func (df *DoctorFactory) GetDoctor(id int) (*Doctor, error) {
    if doctor, ok := df.doctors[id]; ok {
        return doctor, nil
    }
    return nil, fmt.Errorf("Doctor dengan ID %d tidak ditemukan", id)
}

func (df *DoctorFactory) UpdateDoctor(id int, name string) error {
    if _, ok := df.doctors[id]; ok {
        df.doctors[id].Name = name
        return nil
    }
    return fmt.Errorf("Doctor dengan ID %d tidak ditemukan", id)
}

func (df *DoctorFactory) DeleteDoctor(id int) error {
    if _, ok := df.doctors[id]; ok {
        delete(df.doctors, id)
        return nil
    }
    return fmt.Errorf("Doctor dengan ID %d tidak ditemukan", id)
}

func main() {
    factory := NewDoctorFactory()
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("Pilih menu:")
        fmt.Println("1. Buat dokter baru")
        fmt.Println("2. Dapatkan informasi dokter")
        fmt.Println("3. Perbarui nama dokter")
        fmt.Println("4. Hapus dokter")
        fmt.Println("5. Keluar")

        fmt.Print("Masukkan pilihan Anda: ")
        input, _ := reader.ReadString('\n')
        choice, _ := strconv.Atoi(strings.TrimSpace(input))

        switch choice {
        case 1:
            fmt.Print("Masukkan ID dokter: ")
            input, _ = reader.ReadString('\n')
            id, _ := strconv.Atoi(strings.TrimSpace(input))

            fmt.Print("Masukkan nama dokter: ")
            input, _ = reader.ReadString('\n')
            name := strings.TrimSpace(input)

            doctor := factory.CreateDoctor(id, name)
			fmt.Println("Dokter berhasil dibuat dengan ID:", doctor.ID, "dan Nama:", doctor.Name)

        case 2:
            fmt.Print("Masukkan ID dokter: ")
            input, _ = reader.ReadString('\n')
            id, _ := strconv.Atoi(strings.TrimSpace(input))

            doctor, err := factory.GetDoctor(id)
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Println("Informasi dokter:", "dengan ID:", doctor.ID, "dan Nama:", doctor.Name)
            }

        case 3:
            fmt.Print("Masukkan ID dokter: ")
            input, _ = reader.ReadString('\n')
            id, _ := strconv.Atoi(strings.TrimSpace(input))

            fmt.Print("Masukkan nama baru dokter: ")
            input, _ = reader.ReadString('\n')
            name := strings.TrimSpace(input)

            err := factory.UpdateDoctor(id, name)
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Println("Nama dokter berhasil diperbarui.")
            }

        case 4:
            fmt.Print("Masukkan ID dokter: ")
            input, _ = reader.ReadString('\n')
            id, _ := strconv.Atoi(strings.TrimSpace(input))

            err := factory.DeleteDoctor(id)
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Println("Dokter berhasil dihapus.")
            }

        case 5:
            fmt.Println("Keluar dari program.")
            return

        default:
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
        }
    }
}