package counters

type User struct {
	Name string
	Email string
}

type admin struct {
	User
	Rights int
}