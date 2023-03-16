package main

import (
	"fmt"
	"simpleSQL/config"
	"simpleSQL/produk"
)

// ini remark ugi
func main() {
	var koneksi = config.InitSQL()
	var mdl = produk.ProdukModel{}
	mdl.SetSQLConnection(koneksi)

	var menu int
	for menu != 99 {
		fmt.Println("1. Input Barang")
		fmt.Println("2. Tampilkan Daftar Barang")
<<<<<<< HEAD
		fmt.Println("3. Edit Barang")
=======
		fmt.Println("4. Delete Barang")
>>>>>>> feature: delete barang
		fmt.Println("99. Exit")
		fmt.Print("Masukkan pilihan ")
		fmt.Scanln(&menu)
		switch menu {
		case 1:

		case 2:
			res, err := mdl.GetAllProduct()
			if err != nil {
				fmt.Println("Terjadi sebuah kesalahan")
				break
			}

			for i := 0; i < len(res); i++ {
				fmt.Println(res[i])
			}
		case 3:
			var barcode int
			var inputData produk.Produk
			fmt.Print("Inputkan Barcode ")
			fmt.Scanln(&barcode)
			fmt.Print("Masukkan nama produk")
			fmt.Scanln(&inputData.Nama)

			err := mdl.EditBarang(barcode, inputData)
			if err != nil {
				fmt.Println("terjadi kesalahan saat update data")
			}
		case 4:
			var barcode int
			fmt.Print("Masukkan barcode produk:")
			fmt.Scanln(&barcode)
			err := mdl.DeleteProduk(barcode)
			if err != nil {
				fmt.Println("Terjadi sebuah kesalahan")
				break
			}

			fmt.Println("sukses menghapus data")
		}
	}
}
