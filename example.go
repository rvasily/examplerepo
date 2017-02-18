package examplerepo

const FirstName string = "FirstName"

func GiveName(m map[string][string]) (string, error) {
	val, err := m[FirstName]
	return val, err
}


