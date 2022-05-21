package rakha_say_hello

import "fmt"

type Orang struct {
	id        int
	name      string
	type_user string
}
type Filter func(string, int) int

type FilterTypeUser func(string) bool

type HasName interface {
	GetName() string
}

func (orang Orang) GetName() string {
	return orang.name
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

		Is_Valid(data.type_user, func(s string) bool {
			return s == "Admin"
		})
		resultFilter := filter(data.name, data.id)
		total = total + resultFilter

	}

	fmt.Println(total)
}
