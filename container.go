package example

type User struct{
	Name string
}

func (u User) NameString() string{
	return u.Name
}