package main

import (
	"encoding/json"
	"net/url"
	"strings"
)

func make_filter() string {
	filter_data := map[string]interface{}{
		"filter": map[string]interface{}{
			"field": map[string]interface{}{
				"key":    "type",
				"sign":   "LIKE",
				"values": [1]string{"MATRIX_REQUEST"},
			},
		},
		"sort": map[string]interface{}{
			"fields":    [1]string{"time"},
			"direction": "DESC",
		},
		"limit": 10,
	}
	jsonData, err := json.Marshal(filter_data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func get_user(row map[string]interface{}) string {
	author := row["author"].(map[string]interface{})
	return author["user_name"].(string)
}
func get_time(row map[string]interface{}) string {
	return row["time"].(string)
}

func get_params(row map[string]interface{}) string {
	comment := row["params"].(map[string]interface{})
	jsonData, err := json.Marshal(comment)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func make_form(time string, comment string, user string) string {
	form := url.Values{}
	form.Add("period_start", "2023-09-01")
	form.Add("period_end", "2023-09-30")
	form.Add("period_key", "month")
	form.Add("indicator_to_mo_id", "315914")
	form.Add("indicator_to_mo_fact_id", "0")
	form.Add("value", "1")
	form.Add("fact_time", strings.Split(time, "T")[0])
	form.Add("is_plan", "0")
	form.Add("supertags", "[{\"tag\":{\"id\":2,\"name\":\"Клиент\",\"key\":\"client\",\"values_source\":0},\"value\":\""+user+"\"}]") //:[{"tag":{"id":2,"name":"Клиент","key":"client","values_source":0},"value":"Иванов И. И."}]
	form.Add("auth_user_id", "40")
	form.Add("comment", comment)
	return form.Encode()
}
func get_supertags(user string) string {
	supertags := map[string]interface{}{
		"tag": map[string]interface{}{
			"id":            2,
			"name":          "Клиент",
			"key":           "client",
			"values_source": 0},
		"value": user,
	}
	var arrsupertags []map[string]interface{}
	arrsupertags = append(arrsupertags, supertags)
	jsonData, err := json.Marshal(arrsupertags)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
