package main

import (
	"fmt"
	"time"
)

// Definisikan struktur data untuk barang
type Barang struct {
	Kode  string
	Nama  string
	Harga float64
	Stok  int
}

// Definisikan struktur data untuk transaksi
type Transaksi struct {
	Nomor   int
	Tanggal string
	Barang  []Barang
	Total   float64
}

// Modul Barang
func TambahBarang(barang *[]Barang, kode, nama string, harga float64, stok int) {
	newBarang := Barang{Kode: kode, Nama: nama, Harga: harga, Stok: stok}
	*barang = append(*barang, newBarang)
}

func EditBarang(barang *[]Barang, kode, nama string, harga float64, stok int) {
	index := CariBarangSequential(*barang, kode)
	if index != -1 {
		(*barang)[index].Nama = nama
		(*barang)[index].Harga = harga
		(*barang)[index].Stok = stok
		fmt.Println("Barang berhasil diubah.")
	} else {
		fmt.Println("Barang tidak ditemukan.")
	}
}

func HapusBarang(barang *[]Barang, kode string) {
	index := CariBarangSequential(*barang, kode)
	if index != -1 {
		*barang = append((*barang)[:index], (*barang)[index+1:]...)
		fmt.Println("Barang berhasil dihapus.")
	} else {
		fmt.Println("Barang tidak ditemukan.")
	}
}

// Modul Transaksi
func CatatTransaksi(transaksi *[]Transaksi, nomor int, barang []Barang, total float64) {
	newTransaksi := Transaksi{Nomor: nomor, Tanggal: time.Now().Format("2006-01-02"), Barang: barang, Total: total}
	*transaksi = append(*transaksi, newTransaksi)
}

func TampilkanDaftarTransaksi(transaksi []Transaksi) {
	if len(transaksi) == 0 {
		fmt.Println("Belum ada transaksi yang dicatat.")
		return
	}

	fmt.Println("Daftar Transaksi:")
	for _, tr := range transaksi {
		fmt.Printf("Nomor Transaksi: %d\n", tr.Nomor)
		fmt.Printf("Tanggal: %s\n", tr.Tanggal)
		fmt.Println("Barang yang dibeli:")
		for _, brg := range tr.Barang {
			fmt.Printf("- %s (%s) x %d = %.2f\n", brg.Nama, brg.Kode, brg.Stok, brg.Harga*float64(brg.Stok))
		}
		fmt.Printf("Total: %.2f\n", tr.Total)
		fmt.Println("========================================")
	}
}

func HitungOmzetHarian(transaksi []Transaksi, tanggal string) float64 {
	var omzet float64
	for _, tr := range transaksi {
		if tr.Tanggal == tanggal {
			omzet += tr.Total
		}
	}
	return omzet
}

// Modul Pencarian
func CariBarangSequential(data []Barang, kode string) int {
	for i, brg := range data {
		if brg.Kode == kode {
			return i
		}
	}
	return -1
}

func CariBarangBinary(data []Barang, kode string) int {
	low, high := 0, len(data)-1
	for low <= high {
		mid := (low + high) / 2
		if data[mid].Kode == kode {
			return mid
		} else if data[mid].Kode < kode {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Modul Pengurutan
func SelectionSortBarang(data []Barang, kategori string, ascending bool) {
	for i := 0; i < len(data)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(data); j++ {
			if (ascending && data[j].Kode < data[minIndex].Kode) || (!ascending && data[j].Kode > data[minIndex].Kode) {
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
}

func InsertionSortBarang(data []Barang, kategori string, ascending bool) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && (ascending && data[j].Kode > key.Kode || !ascending && data[j].Kode < key.Kode) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func main() {
	// Inisialisasi array barang dan array transaksi
	var daftarBarang []Barang
	var daftarTransaksi []Transaksi

	// Loop utama program untuk interaksi dengan pengguna
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Edit Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Catat Transaksi")
		fmt.Println("5. Tampilkan Daftar Transaksi")
		fmt.Println("6. Hitung Omzet Harian")
		fmt.Println("7. Keluar")

		var pilihan int
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var kode, nama string
			var harga float64
			var stok int
			fmt.Print("Masukkan kode barang: ")
			fmt.Scan(&kode)
			fmt.Print("Masukkan nama barang: ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan harga barang: ")
			fmt.Scan(&harga)
			fmt.Print("Masukkan stok barang: ")
			fmt.Scan(&stok)
			TambahBarang(&daftarBarang, kode, nama, harga, stok)
			fmt.Println("Barang berhasil ditambahkan.")
		case 2:
			var kode, nama string
			var harga float64
			var stok int
			fmt.Print("Masukkan kode barang yang ingin diubah: ")
			fmt.Scan(&kode)
			fmt.Print("Masukkan nama baru barang: ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan harga baru barang: ")
			fmt.Scan(&harga)
			fmt.Print("Masukkan stok baru barang: ")
			fmt.Scan(&stok)
			EditBarang(&daftarBarang, kode, nama, harga, stok)
		case 3:
			var kode string
			fmt.Print("Masukkan kode barang yang ingin dihapus: ")
			fmt.Scan(&kode)
			HapusBarang(&daftarBarang, kode)
		case 4:
			var nomorTransaksi int
			fmt.Print("Masukkan nomor transaksi: ")
			fmt.Scan(&nomorTransaksi)
			var kodeBarang string
			var barangDibeli []Barang
			var totalTransaksi float64
			for {
				fmt.Print("Masukkan kode barang yang dibeli (ketik 'selesai' untuk mengakhiri): ")
				fmt.Scan(&kodeBarang)
				if kodeBarang == "selesai" {
					break
				}
				index := CariBarangSequential(daftarBarang, kodeBarang)
				if index == -1 {
					fmt.Println("Barang tidak ditemukan.")
				} else {
					var jumlah int
					fmt.Printf("Masukkan jumlah %s yang dibeli: ", daftarBarang[index].Nama)
					fmt.Scan(&jumlah)
					barangDibeli = append(barangDibeli, daftarBarang[index])
					totalTransaksi += daftarBarang[index].Harga * float64(jumlah)
				}
			}
			CatatTransaksi(&daftarTransaksi, nomorTransaksi, barangDibeli, totalTransaksi)
			fmt.Println("Transaksi berhasil dicatat.")
		case 5:
			TampilkanDaftarTransaksi(daftarTransaksi)
		case 6:
			var tanggal string
			fmt.Print("Masukkan tanggal untuk menghitung omzet (format: YYYY-MM-DD): ")
			fmt.Scan(&tanggal)
			omzet := HitungOmzetHarian(daftarTransaksi, tanggal)
			fmt.Printf("Omzet pada tanggal %s adalah %.2f\n", tanggal, omzet)
		case 7:
			fmt.Println("Terima kasih, program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
