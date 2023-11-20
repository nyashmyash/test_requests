package main

func main() {
	coocies := login()
	data := send_req_filter(coocies)
	print_data(data)
	send_data(data)
}
