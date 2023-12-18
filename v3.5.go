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
			tambahData(&A)
		} else if ans == 2 {
			editData(&A)
		} else if ans == 5 {
			daftarData(A)
		}
		if ans != 7 {
			uimenu()
			fmt.Scan(&ans)
		}

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

func uiEditData() {
	fmt.Println("1. Edit data pabrikan")
	fmt.Println("2. Edit data mobil")
	fmt.Println("3. Keluar")
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func uiEditDataMobil() {
	fmt.Println("1. Edit data nama mobil")
	fmt.Println("2. Edit data jumlah unit")
	fmt.Println("3. Edit data tahun keluar")
	fmt.Println("4. Edit data jumlah penjualan")
	fmt.Println("5. Keluar")
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func pilihan() {
	fmt.Println("1. Setuju")
	fmt.Println("2. Batal")
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func pilihanEditPabrik() {
	fmt.Println("1. Edit nama pabrik")
	fmt.Println("2. Kelaur")
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func tambahData(A *arrPabrikan) {
	var temp pabrikan
	var ans, ans1, idx int
	var tMob mobil
	var t string

	uiTambahData()
	fmt.Scan(&ans)
	for ans != 3 {
		for ans < 1 || ans > 3 {
			uiTambahData()
			fmt.Scan(&ans)
		}
		//opsi pilihan 1
		if ans == 1 {

			if A.n < NMAX {
				fmt.Print("Masukkan nama pabrikan : ")
				fmt.Scan(&temp.nama)

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
			} else {
				fmt.Println("Data sudah penuh")
			}

			//opsi 2
		} else if ans == 2 {

			cetakPabrik(*A)
			fmt.Println()
			fmt.Print("Masukkan nama pabrik yang ingin diisi :")
			fmt.Scan(&t)
			fmt.Println()
			idx = cariPabrik(*A, t)
			if idx == -1 {
				fmt.Println("Data tidak ditemukan")
			} else {
				if A.Pabrik[idx].arrMobil.n < NMAX {
					fmt.Println("Ketikan data secara horizontal dan ditambahi spasi dengan format:")
					fmt.Println("Nama mobil | Jumlah unit | Jumlah penjualan | Tahun keluar ")
					fmt.Scan(&tMob.namaMobil, &tMob.jumlahUnit, &tMob.jumlahPenjualan, &tMob.tahunKeluar)
					tMob.namaPabrikan = A.Pabrik[idx].nama
				} else {
					fmt.Println("Data sudah penuh")
				}

				pilihan()
				fmt.Scan(&ans1)

				//jika pilihan user lebih dari opsi
				for ans1 < 1 && ans1 > 2 {
					pilihan()
					fmt.Scan(&ans1)
				}
				//Jika memilih opsi simpan
				if ans1 == 1 {
					A.Pabrik[idx].arrMobil.M[A.Pabrik[idx].arrMobil.n] = tMob
					A.Pabrik[idx].arrMobil.n++
					fmt.Println("Data berhasil ditambahkan")
				}

			}

		}
		if ans != 3 {
			uiTambahData()
			fmt.Scan(&ans)
		}

	}
}

func editData(A *arrPabrikan) {
	var ans, idx, ans1, ans2, idxM, tempInt int
	var t, temp, tMobil string
	var tPab pabrikan

	uiEditData()
	fmt.Scan(&ans)
	for ans != 3 {
		for ans < 1 || ans > 3 {
			uiEditData()
			fmt.Scan(&ans)
		}

		if ans == 1 {
			cetakPabrik(*A)
			fmt.Print("Masukkan nama pabrik yang ingin diedit : ")
			fmt.Scan(&t)

			idx = cariPabrik(*A, t)

			if idx == -1 {
				fmt.Println("Data tidak ditemukan")
			} else {
				fmt.Println("Data ditemukan")
				pilihanEditPabrik()
				fmt.Scan(&ans1)
				for ans < 1 || ans > 2 {
					pilihanEditPabrik()
					fmt.Scan(&ans1)
				}

				if ans1 == 1 {
					fmt.Print("Masukkan nama pabrik yang baru :")
					fmt.Scan(&temp)

					pilihan()
					fmt.Scan(&ans2)

					//jika pilihan user lebih dari opsi
					for ans2 < 1 && ans2 > 2 {
						pilihan()
						fmt.Scan(&ans2)
					}
					//Jika memilih opsi simpan
					if ans2 == 1 {
						A.Pabrik[idx].nama = temp
						fmt.Println("Data berhasil diubah")
					}
				}
			}
		} else if ans == 2 {
			cetakPabrik(*A)
			fmt.Print("Masukkan nama pabrikan mobil yang ingin diedit : ")
			fmt.Scan(&t)

			idx = cariPabrik(*A, t)

			if idx == -1 {
				fmt.Println("Data tidak ditemukan")
			} else {
				fmt.Println("Data ditemukan")

				tPab = A.Pabrik[idx]

				cetakMobil(tPab)

				fmt.Print("Masukkan nama mobil yang ingin diubah datanya : ")
				fmt.Scan(&tMobil)
				idxM = cariMobil(tPab, tMobil)

				if idxM == -1 {
					fmt.Println("Data tidak ditemukan")
				} else {
					uiEditDataMobil()
					fmt.Scan(&ans1)

					for ans1 < 1 || ans1 > 5 {
						uiEditDataMobil()
						fmt.Scan(&ans1)
					}

					if ans1 == 1 {
						fmt.Print("Masukkan nama mobil yang baru :")
						fmt.Scan(&temp)

						pilihan()
						fmt.Scan(&ans2)

						//jika pilihan user lebih dari opsi
						for ans2 < 1 && ans2 > 2 {
							pilihan()
							fmt.Scan(&ans2)
						}
						//Jika memilih opsi simpan
						if ans2 == 1 {
							A.Pabrik[idx].arrMobil.M[idxM].namaMobil = temp
							fmt.Println("Data berhasil diubah")
						}
					} else if ans1 == 2 {
						fmt.Print("Masukkan data jumlah unit yang baru : ")
						fmt.Scan(&tempInt)

						pilihan()
						fmt.Scan(&ans2)

						//jika pilihan user lebih dari opsi
						for ans2 < 1 && ans2 > 2 {
							pilihan()
							fmt.Scan(&ans2)
						}

						if ans2 == 1 {
							A.Pabrik[idx].arrMobil.M[idxM].jumlahUnit = tempInt
							fmt.Println("Data berhasil diubah")
						}
					} else if ans1 == 3 {
						fmt.Print("Masukkan data tahun keluar yang baru : ")
						fmt.Scan(&tempInt)

						pilihan()
						fmt.Scan(&ans2)

						//jika pilihan user lebih dari opsi
						for ans2 < 1 && ans2 > 2 {
							pilihan()
							fmt.Scan(&ans2)
						}

						if ans2 == 1 {
							A.Pabrik[idx].arrMobil.M[idxM].tahunKeluar = tempInt
							fmt.Println("Data berhasil diubah")
						}
					} else if ans1 == 4 {
						fmt.Print("Masukkan data jumlah penjualan yang baru : ")
						fmt.Scan(&tempInt)

						pilihan()
						fmt.Scan(&ans2)

						//jika pilihan user lebih dari opsi
						for ans2 < 1 && ans2 > 2 {
							pilihan()
							fmt.Scan(&ans2)
						}

						if ans2 == 1 {
							A.Pabrik[idx].arrMobil.M[idxM].jumlahPenjualan = tempInt
							fmt.Println("Data berhasil diubah")
						}
					}
				}

			}

		}
		if ans != 3 {
			uiEditData()
			fmt.Scan(&ans)
		}
	}
}

func daftarData(A arrPabrikan) {
	var ans, ans1, ans2 int

	menghitungPenjualanPabrik(&A)

	uiDaftarData()
	fmt.Scan(&ans)

	for ans != 3 {
		for ans < 1 || ans > 3 {
			uiDaftarData()
			fmt.Scan(&ans)
		}

		if ans == 1 {
			uiDaftarData2()
			fmt.Scan(&ans1)

			for ans1 != 3 {
				for ans1 < 1 || ans1 > 3 {
					uiDaftarData2()
					fmt.Scan(&ans1)
				}

				if A.n == 0 {
					fmt.Println("Tidak ada data")
				} else if ans1 == 1 {
					uiDaftarDataPabrik()
					fmt.Scan(&ans2)

					for ans2 != 3 {
						for ans2 < 1 || ans2 > 3 {
							uiDaftarDataPabrik()
							fmt.Scan(&ans2)
						}
						if ans2 == 1 {
							sortAscendingPabrikan(&A, 1)
							cetakPabrik(A)
						} else if ans2 == 2 {
							sortAscendingPabrikan(&A, 2)
							cetakTotalPenjualanPabrik(A)
						}
						if ans2 != 3 {
							uiDaftarDataPabrik()
							fmt.Scan(&ans2)
						}
					}

				}

				if ans1 != 3 {
					uiDaftarData2()
					fmt.Scan(&ans1)
				}
			}
		} else if ans == 2 {
			uiDaftarData2()
			fmt.Scan(&ans1)

			for ans1 != 3 {
				for ans1 < 1 || ans1 > 3 {
					uiDaftarData2()
					fmt.Scan(&ans1)
				}

				if A.n == 0 {
					fmt.Println("Tidak ada data")
				} else if ans1 == 1 {
					uiDaftarDataPabrik()
					fmt.Scan(&ans2)

					for ans2 != 3 {
						for ans2 < 1 || ans2 > 3 {
							uiDaftarDataPabrik()
							fmt.Scan(&ans2)
						}
						if ans2 == 1 {
							sortDescendingPabrikan(&A, 1)
							cetakPabrik(A)
						} else if ans2 == 2 {
							sortDescendingPabrikan(&A, 2)
							cetakTotalPenjualanPabrik(A)
						}
						if ans2 != 3 {
							uiDaftarDataPabrik()
							fmt.Scan(&ans2)
						}
					}
				}
				if ans != 3 {
					uiDaftarData2()
					fmt.Scan(&ans1)
				}
			}

		}

		if ans != 3 {
			uiDaftarData()
			fmt.Scan(&ans)
		}
	}
}

func uiDaftarData() {
	fmt.Println("1. Daftar terurut secara ascending")
	fmt.Println("2. Daftar terurut secara descending")
	fmt.Println("3. Keluar")
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func uiDaftarData2() {
	fmt.Println("1. Data Pabrikan")
	fmt.Println("2. Data Mobil")
	fmt.Println("3. Keluar")
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func uiDaftarDataMobil() {
	fmt.Println("Diurutkan berdasarkan: ")
	fmt.Println("1. Nama Pabrikan")
	fmt.Println("2. Nama Mobil")
	fmt.Println("3. Tahun Keluar")
	fmt.Println("4. Jumlah Unit")
	fmt.Println("5. Jumlah Penjualan")
	fmt.Println("6. Keluar")
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func uiDaftarDataPabrik() {
	fmt.Println("Diurutkan berdasarkan: ")
	fmt.Println("1. Nama Pabrikan")
	fmt.Println("2. Total Penjualan")
	fmt.Println("3. Keluar")
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

//sequential search
func cariPabrik(A arrPabrikan, target string) int {
	for i := 0; i < A.n; i++ {
		if A.Pabrik[i].nama == target {
			return i
		}
	}
	return -1
}

func cetakPabrik(A arrPabrikan) {
	fmt.Println("Daftar Pabrik")
	for i := 0; i < A.n; i++ {
		fmt.Println(i+1, ". ", A.Pabrik[i].nama)
	}
}

func cetakTotalPenjualanPabrik(A arrPabrikan) {
	fmt.Println("Daftar Pabrik")
	for i := 0; i < A.n; i++ {
		fmt.Println(i+1, ". ", A.Pabrik[i].nama, "\t", A.Pabrik[i].jumlahPenjualanPabrik)
	}
}

func cariMobil(A pabrikan, target string) int {
	for i := 0; i < A.arrMobil.n; i++ {
		if A.arrMobil.M[i].namaMobil == target {
			return i
		}
	}
	return -1
}

func cetakMobil(A pabrikan) {
	fmt.Println("Daftar Mobil Pabrik", A.nama)
	fmt.Println("Nama Mobil|Nama Pabrikan|Jumlah Unit|Jumlah Penjualan|Tahun Keluar")
	for i := 0; i < A.arrMobil.n; i++ {
		fmt.Println(i+1, ".", A.arrMobil.M[i].namaMobil, "\t", A.arrMobil.M[i].namaPabrikan, "\t", A.arrMobil.M[i].jumlahUnit, "\t", A.arrMobil.M[i].jumlahPenjualan, "\t", A.arrMobil.M[i].tahunKeluar)
	}
}

//insertion sort
func sortAscendingPabrikan(A *arrPabrikan, code int) {
	var temp pabrikan
	var j int

	if code == 1 {
		for i := 1; i < A.n; i++ {
			j = i
			temp = A.Pabrik[j]
			for j > 0 && temp.nama < A.Pabrik[j-1].nama {
				A.Pabrik[j] = A.Pabrik[j-1]
				j--
			}
			A.Pabrik[j] = temp
		}
	} else if code == 2 {
		for i := 1; i < A.n; i++ {
			j = i
			temp = A.Pabrik[j]
			for j > 0 && temp.jumlahPenjualanPabrik < A.Pabrik[j-1].jumlahPenjualanPabrik {
				A.Pabrik[j] = A.Pabrik[j-1]
				j--
			}
			A.Pabrik[j] = temp
		}
	}
}

//insertion sort
func sortDescendingPabrikan(A *arrPabrikan, code int) {
	var temp pabrikan
	var j int

	if code == 1 {
		for i := 1; i < A.n; i++ {
			j = i
			temp = A.Pabrik[j]
			for j > 0 && temp.nama > A.Pabrik[j-1].nama {
				A.Pabrik[j] = A.Pabrik[j-1]
				j--
			}
			A.Pabrik[j] = temp
		}
	} else if code == 2 {
		for i := 1; i < A.n; i++ {
			j = i
			temp = A.Pabrik[j]
			for j > 0 && temp.jumlahPenjualanPabrik > A.Pabrik[j-1].jumlahPenjualanPabrik {
				A.Pabrik[j] = A.Pabrik[j-1]
				j--
			}
			A.Pabrik[j] = temp
		}
	}
}

//selection sort
func sortAscendingMobil(B arrMobil, code int) {
	var iMax int
	var temp mobil

	if code == 1 {
		for i := 0; i < B.n; i++ {
			iMax = i
			for j := i + 1; j < B.n; j++ {
				if B.M[j].namaMobil > B.M[iMax].namaMobil {
					iMax = j
				}
			}
			temp = B.M[iMax]
			B.M[iMax] = B.M[i]
			B.M[i] = temp
		}
	}

}

func gabungSemuaMobil(A arrPabrikan, B *arrMobil) {
	var t int
	for i := 0; i < A.n; i++ {
		for j := 0; j < A.Pabrik[i].arrMobil.n; i++ {
			B.M[t] = A.Pabrik[i].arrMobil.M[j]
			t++
		}
	}
	B.n = t
}

func menghitungPenjualanPabrik(A *arrPabrikan) {
	for i := 0; i < A.n; i++ {
		for j := 0; j < A.Pabrik[i].arrMobil.n; j++ {
			A.Pabrik[i].jumlahPenjualanPabrik += A.Pabrik[i].arrMobil.M[j].jumlahPenjualan
		}
	}
}
