package examplerepo

const FirstName string = "FirstName"

func GiveName(m map[string]string) (string, bool) {
	val, exist := m[FirstName]
	return val, exist
}


