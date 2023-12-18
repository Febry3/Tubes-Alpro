package main

import "fmt"

const NMAX int = 10

type mobil struct {
	namaMobil, namaPabrikan                  string
	jumlahUnit, tahunKeluar, jumlahPenjualan int
}

type arrMobil struct {
	M [NMAX]mobil
	n int
}

type pabrikan struct {
	nama                  string
	arrMobil              arrMobil
	jumlahPenjualanPabrik int
}

type arrPabrikan struct {
	Pabrik [NMAX]pabrikan
	n      int
}

func main() {
	var ans int
	var A arrPabrikan

	uimenu()
	fmt.Scan(&ans)
	for ans != 7 {
		for ans < 1 || ans > 7 {
			uimenu()
			fmt.Scan(&ans)
		}

		if ans == 1 {
			tambahDataPabrikan(&A)
		}
		uimenu()
		fmt.Scan(&ans)

	}
	fmt.Print(A)
}

func uimenu() {
	fmt.Println("==============================================")
	fmt.Println("|         Selamat datang di aplikasi         |")
	fmt.Println("|                Dealer Mobil                |")
	fmt.Println("==============================================")
	fmt.Println()
	fmt.Println("=== Main Menu ===")
	fmt.Println()
	fmt.Println("1. Tambah data ")
	fmt.Println("2. Ubah data ")
	fmt.Println("3. Hapus data ")
	fmt.Println("4. Cari data")
	fmt.Println("5. Daftar data")
	fmt.Println("6. Top 3 data tertinggi")
	fmt.Println("7. Keluar")
	fmt.Println()
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func uiTambahData() {
	fmt.Println("1. Tambah data pabrikan")
	fmt.Println("2. Tambah data mobil")
	fmt.Println("3. Keluar")
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func pilihan() {
	fmt.Println("1. Setuju")
	fmt.Println("2. Batal")
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func tambahDataPabrikan(A *arrPabrikan) {
	var temp pabrikan
	var ans, ans1 int
	var tMob arrMobil

	uiTambahData()
	fmt.Scan(&ans)
	for ans != 3 {
		for ans < 1 || ans > 7 {
			uiTambahData()
			fmt.Scan(&ans)
		}
		//opsi pilihan 1
		if ans == 1 {
			var stop bool = false
			for !stop {
				if A.n <= NMAX {
					fmt.Print("Masukkan nama pabrikan : ")
					fmt.Scan(&temp.nama)
				} else {
					fmt.Println("Data sudah penuh")
				}

				pilihan()
				fmt.Scan(&ans1)

				for ans1 < 1 && ans1 > 2 {
					pilihan()
					fmt.Scan(&ans1)
				}
				if ans1 == 1 {
					A.Pabrik[A.n].nama = temp.nama
					A.n++
					fmt.Println("Data berhasil ditambahkan")
				}
				if ans1 == 2 {
					stop = true
				}
			}

			//opsi 2
		} else if ans == 2 {
			var t string
			var idx int
			cetakPabrik(*A)
			fmt.Println()
			fmt.Print("Masukkan nama pabrik yang ingin diisi :")
			fmt.Scan(&t)
			fmt.Println()
			idx = cariPabrik(*A, t)
			if idx == -1 {
				fmt.Println("Data tidak ditemukan")
			} else {
				if A.Pabrik[idx].arrMobil.n <= NMAX {
					fmt.Println("Ketika data secara horizontal dan ditambahi spasi dimulai dari :")
					fmt.Println("Nama mobil | Nama pabrikan | Jumlah unit | Jumlah penjualan | Tahun keluar ")
					fmt.Scan(&tMob.M[tMob.n].namaMobil, &tMob.M[tMob.n].namaPabrikan, &tMob.M[tMob.n].jumlahUnit, &tMob.M[tMob.n].jumlahPenjualan, &tMob.M[tMob.n].tahunKeluar)
					tMob.n++
				} else {
					fmt.Println("Data sudah penuh")
				}

				pilihan()
				fmt.Scan(&ans1)

				for ans1 < 1 && ans1 > 2 {
					pilihan()
					fmt.Scan(&ans1)
				}
				if ans1 == 1 {
					A.Pabrik[idx].arrMobil = tMob
					A.Pabrik[idx].arrMobil.n = tMob.n
					fmt.Println("Data berhasil ditambahkan")
				}

			}

		}
		uiTambahData()
		fmt.Scan(&ans)
	}
}

/*
func tambahDataMobil(A arrPabrikan) {
	var temp pabrikan
	var ans int

	fmt.Println("Ketika data secara horizontal dan ditambahi spasi dimulai dari :")
	fmt.Println("Nama mobil | Nama pabrikan | Jumlah unit | Jumlah penjualan | Tahun keluar ")
	if A.n <= NMAX {
		fmt.Scan(&temp.arrMobil[temp.n].namaMobil, &temp.arrMobil[temp.n].namaPabrikan, &temp.arrMobil[temp.n].jumlahUnit, &temp.arrMobil[temp.n].jumlahPenjualan, &temp.arrMobil[temp.n].tahunKeluar)
		temp.n++
	}
	fmt.Println("1. Simpan")
	fmt.Println("2. Batal")
	fmt.Scan(&ans)
	if ans == 1 {
		A.arrMobil[temp.n] = temp
	}

}
*/

func cariPabrik(A arrPabrikan, target string) int {
	for i := 0; i < A.n; i++ {
		if A.Pabrik[i].nama == target {
			return i
		}
	}
	return -1
}

func cetakPabrik(A arrPabrikan) {
	for i := 0; i < A.n; i++ {
		fmt.Println(i+1, ". ", A.Pabrik[i].nama)
	}
}
