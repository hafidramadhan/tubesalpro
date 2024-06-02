package main

import (
	"fmt"
	"time"
)

const maxBarang = 100    // Ukuran maksimum array barang
const maxTransaksi = 100 // Ukuran maksimum array transaksi

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
	Barang  [maxBarang]Barang
	Total   float64
}

// Inisialisasi array barang dan array transaksi
var daftarBarang [maxBarang]Barang
var daftarTransaksi [maxTransaksi]Transaksi

// Fungsi untuk menambahkan barang baru ke dalam daftar barang
func TambahBarang(kode, nama string, harga float64, stok int) {
	for i := 0; i < maxBarang; i++ {
		if daftarBarang[i].Kode == "" {
			daftarBarang[i] = Barang{Kode: kode, Nama: nama, Harga: harga, Stok: stok}
			fmt.Println("Barang berhasil ditambahkan.")
			return
		}
	}
	fmt.Println("Daftar barang sudah penuh.")
}

// Fungsi untuk mengubah data barang berdasarkan kode barang
func UbahBarang(kode, nama string, harga float64, stok int) {
	for i := 0; i < maxBarang; i++ {
		if daftarBarang[i].Kode == kode {
			daftarBarang[i].Nama = nama
			daftarBarang[i].Harga = harga
			daftarBarang[i].Stok = stok
			fmt.Println("Barang berhasil diubah.")
			return
		}
	}
	fmt.Println("Barang tidak ditemukan.")
}

// Fungsi untuk menghapus data barang berdasarkan kode barang
func HapusBarang(kode string) {
	for i := 0; i < maxBarang; i++ {
		if daftarBarang[i].Kode == kode {
			for j := i; j < maxBarang-1; j++ {
				daftarBarang[j] = daftarBarang[j+1]
			}
			daftarBarang[maxBarang-1] = Barang{} // Menghapus data terakhir
			fmt.Println("Barang berhasil dihapus.")
			return
		}
	}
	fmt.Println("Barang tidak ditemukan.")
}

// Fungsi untuk mencatat transaksi
func CatatTransaksi(nomor int, barang [maxBarang]Barang, total float64) {
	for i := 0; i < maxTransaksi; i++ {
		if daftarTransaksi[i].Nomor == 0 {
			daftarTransaksi[i] = Transaksi{Nomor: nomor, Tanggal: time.Now().Format("2006-01-02"), Barang: barang, Total: total}
			fmt.Println("Transaksi berhasil dicatat.")
			return
		}
	}
	fmt.Println("Daftar transaksi sudah penuh.")
}

// Fungsi untuk menampilkan daftar transaksi
func TampilkanDaftarTransaksi() {
	fmt.Println("\nDaftar Transaksi:")
	for _, transaksi := range daftarTransaksi {
		if transaksi.Nomor != 0 {
			fmt.Printf("Nomor Transaksi: %d\n", transaksi.Nomor)
			fmt.Printf("Tanggal: %s\n", transaksi.Tanggal)
			fmt.Println("Barang yang dibeli:")
			for _, barang := range transaksi.Barang {
				if barang.Kode != "" {
					fmt.Printf("- %s (%s) x %d = %.2f\n", barang.Nama, barang.Kode, barang.Stok, barang.Harga*float64(barang.Stok))
				}
			}
			fmt.Printf("Total: %.2f\n", transaksi.Total)
			fmt.Println("========================================")
		}
	}
}

// Fungsi untuk menghitung omzet harian
func HitungOmzetHarian(tanggal string) float64 {
	var omzet float64
	for _, transaksi := range daftarTransaksi {
		if transaksi.Nomor != 0 && transaksi.Tanggal == tanggal {
			omzet += transaksi.Total
		}
	}
	return omzet
}

// Fungsi untuk mencari barang berdasarkan kode dengan algoritma sequential search
func CariBarangSequential(kode string) int {
	for i, barang := range daftarBarang {
		if barang.Kode == kode {
			return i
		}
	}
	return -1
}

// Fungsi untuk mencari barang berdasarkan kode dengan algoritma binary search
func CariBarangBinary(kode string) int {
	low, high := 0, maxBarang-1
	for low <= high {
		mid := (low + high) / 2
		if daftarBarang[mid].Kode == kode {
			return mid
		} else if daftarBarang[mid].Kode < kode {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Fungsi untuk mengurutkan daftar barang berdasarkan kategori dengan algoritma selection sort
func SelectionSortBarang(kategori string, ascending bool) {
	for i := 0; i < maxBarang-1; i++ {
		minIndex := i
		for j := i + 1; j < maxBarang; j++ {
			switch kategori {
			case "Kode":
				if (ascending && daftarBarang[j].Kode < daftarBarang[minIndex].Kode) || (!ascending && daftarBarang[j].Kode > daftarBarang[minIndex].Kode) {
					minIndex = j
				}
			case "Nama":
				if (ascending && daftarBarang[j].Nama < daftarBarang[minIndex].Nama) || (!ascending && daftarBarang[j].Nama > daftarBarang[minIndex].Nama) {
					minIndex = j
				}
			case "Harga":
				if (ascending && daftarBarang[j].Harga < daftarBarang[minIndex].Harga) || (!ascending && daftarBarang[j].Harga > daftarBarang[minIndex].Harga) {
					minIndex = j
				}
			case "Stok":
				if (ascending && daftarBarang[j].Stok < daftarBarang[minIndex].Stok) || (!ascending && daftarBarang[j].Stok > daftarBarang[minIndex].Stok) {
					minIndex = j
				}
			}
		}
		daftarBarang[i], daftarBarang[minIndex] = daftarBarang[minIndex], daftarBarang[i]
	}
}

// Fungsi untuk mengurutkan daftar barang berdasarkan kategori dengan algoritma insertion sort
func InsertionSortBarang(kategori string, ascending bool) {
	for i := 1; i < maxBarang; i++ {
		key := daftarBarang[i]
		j := i - 1
		for j >= 0 {
			switch kategori {
			case "Kode":
				if (ascending && daftarBarang[j].Kode > key.Kode) || (!ascending && daftarBarang[j].Kode < key.Kode) {
					daftarBarang[j+1] = daftarBarang[j]
					j--
				} else {
					break
				}
			case "Nama":
				if (ascending && daftarBarang[j].Nama > key.Nama) || (!ascending && daftarBarang[j].Nama < key.Nama) {
					daftarBarang[j+1] = daftarBarang[j]
					j--
				} else {
					break
				}
			case "Harga":
				if (ascending && daftarBarang[j].Harga > key.Harga) || (!ascending && daftarBarang[j].Harga < key.Harga) {
					daftarBarang[j+1] = daftarBarang[j]
					j--
				} else {
					break
				}
			case "Stok":
				if (ascending && daftarBarang[j].Stok > key.Stok) || (!ascending && daftarBarang[j].Stok < key.Stok) {
					daftarBarang[j+1] = daftarBarang[j]
					j--
				} else {
					break
				}
			}
		}
		daftarBarang[j+1] = key
	}
}

// Fungsi main sebagai titik masuk utama program
func main() {
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
			TambahBarang(kode, nama, harga, stok)
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
			UbahBarang(kode, nama, harga, stok)
		case 3:
			var kode string
			fmt.Print("Masukkan kode barang yang ingin dihapus: ")
			fmt.Scan(&kode)
			HapusBarang(kode)
		case 4:
			var nomorTransaksi int
			fmt.Print("Masukkan nomor transaksi: ")
			fmt.Scan(&nomorTransaksi)
			var kodeBarang string
			var barangDibeli [maxBarang]Barang
			var totalTransaksi float64
			for {
				fmt.Print("Masukkan kode barang yang dibeli (ketik 'selesai' untuk mengakhiri): ")
				fmt.Scan(&kodeBarang)
				if kodeBarang == "selesai" {
					break
				}
				index := CariBarangSequential(kodeBarang)
				if index == -1 {
					fmt.Println("Barang tidak ditemukan.")
				} else {
					var jumlah int
					fmt.Printf("Masukkan jumlah %s yang dibeli: ", daftarBarang[index].Nama)
					fmt.Scan(&jumlah)
					barangDibeli[index] = daftarBarang[index]
					barangDibeli[index].Stok = jumlah
					totalTransaksi += daftarBarang[index].Harga * float64(jumlah)
				}
			}
			CatatTransaksi(nomorTransaksi, barangDibeli, totalTransaksi)
		case 5:
			TampilkanDaftarTransaksi()
		case 6:
			var tanggal string
			fmt.Print("Masukkan tanggal untuk menghitung omzet (format: YYYY-MM-DD): ")
			fmt.Scan(&tanggal)
			omzet := HitungOmzetHarian(tanggal)
			fmt.Printf("Omzet pada tanggal %s adalah %.2f\n", tanggal, omzet)
		case 7:
			fmt.Println("Terima kasih, program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
