package rakha_say_hello

type Person struct {
	id   int
	name string
}

func (p Person) SayHello(name string) string {
	return p.name
}
func (p Person) ShowId(id int) int {
	return p.id
}
