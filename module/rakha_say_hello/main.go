package rakha_say_hello

import "fmt"

type Orang struct {
	ID        int
	Name      string
	Type_User string
}
type Filter func(string, int) int

type FilterTypeUser func(string) bool

type HasName interface {
	GetName() string
}

func (orang Orang) GetName() string {
	return orang.Name
}

func Say_Helllo(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}
func Is_Valid(name string, block FilterTypeUser) {
	defer EndApp()
	if block(name) {
		panic("Admin Masuk")
	} else {
		fmt.Println("No Blocked")
	}

}
func EndApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error : ", message)

	}
	fmt.Println("Aplikasi Selesai")
}
func Add_Baru() {
	fmt.Println("Rakha")
}

// func main() {
// 	a := []Orang{
// 		{
// 			id:        123,
// 			name:      "Rakha",
// 			type_user: "Admin",
// 		},
// 		{
// 			id:        23,
// 			name:      "Elang",
// 			type_user: "NonAdmin",
// 		},
// 		{
// 			id:        30,
// 			name:      "Elang",
// 			type_user: "NonAdmin",
// 		},
// 	}
// 	do_sum(do_filter, a[0], a[1])
// 	fmt.Println()
// 	say_hello(a[0])
// }
func Do_Filter(name string, id int) int {
	if name == "Rakha" {
		return 0
	} else {
		return id
	}
}
func Do_sum(filter Filter, value ...Orang) {

	total := 0
	for _, data := range value {

		Is_Valid(data.Type_User, func(s string) bool {
			return s == "Admin"
		})
		resultFilter := filter(data.Name, data.ID)
		total = total + resultFilter

	}
	fmt.Println(total)
}
