package shared

// MethodMessage says hello.
func MethodMessage(m string) string {
	request := "Normal Request"
	if m == "POST" {
		request = "POST Request"
	}
	return "Hello, " + request
}
