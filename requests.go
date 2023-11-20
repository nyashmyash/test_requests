package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func send_req_filter(cookie *http.Cookie) map[string]interface{} {
	jsonData := make_filter()

	req, err := http.NewRequest("GET", "https://development.kpi-drive.ru/_api/events", bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		fmt.Println("Error create request", err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error create request:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error read request:", err)
		return nil
	}
	//fmt.Println(string(body))

	resp_data := make(map[string]interface{})

	err = json.Unmarshal(body, &resp_data)
	if err != nil {
		fmt.Println("error parse", err)
		return nil
	}
	return resp_data
}
func send_data(data map[string]interface{}) {
	rows := get_rows(data)
	// for _, value := range rows {
	// 	jsonData, err := json.Marshal(value)
	// 	if err != nil {
	// 	}
	// 	var rdata = req_data{time: get_time(value), user: get_user(value), comment: string(jsonData)}
	// 	send_data_row(rdata)
	// }
	jsonData, err := json.Marshal(rows[0])
	if err != nil {
	}
	var rdata = req_data{time: get_time(rows[0]), user: get_user(rows[0]), comment: string(jsonData)}
	send_data_row(rdata)
}
func print_data(data map[string]interface{}) {
	rows := get_rows(data)
	for _, value := range rows {
		jsonData, err := json.Marshal(value)
		if err != nil {
		}
		var rdata = req_data{time: get_time(value), user: get_user(value), comment: string(jsonData)}
		fmt.Println(rdata)
	}
}
func send_data_row(rdata req_data) {
	form := make_form(rdata)
	//fmt.Println(form)
	req, err := http.NewRequest("POST", "https://development.kpi-drive.ru/_api/facts/save_fact", bytes.NewBuffer([]byte(form)))
	if err != nil {
		fmt.Println("Error create request to POST data:", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer 48ab34464a5573519725deb5865cc74c")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error create request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}
	fmt.Println(string(body))
}

func login() *http.Cookie {
	data := url.Values{}
	data.Set("login", "admin")
	data.Set("password", "admin")
	resp, err := http.PostForm("https://development.kpi-drive.ru/_api/auth/login", data)
	if err != nil {
		fmt.Println("Ошибка при выполнении GET-запроса:", err)
		return nil
	}
	defer resp.Body.Close()
	cookie := resp.Cookies()
	return &*cookie[0]
}
