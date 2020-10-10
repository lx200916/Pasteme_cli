package api

import (
	"bytes"
	"github.com/atotto/clipboard"
	"log"
	"strconv"
)

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Resp struct {
	Status int
	Key    uint64
}

var client *http.Client

func CreateBase(content string, password string, lang string, key string, once bool, raw bool, copy bool) {
	if client == nil {
		client = &http.Client{}
	}
	data := make(map[string]interface{})
	if content == "" {
		fmt.Println("Content Can NOT Be Empty")
		return
	} else {
		data["content"] = content
	}
	if lang != "" {
		data["lang"] = lang
	} else {
		data["lang"] = "plain"
	}
	if password != "" {
		data["password"] = password
	}
	var url = "https://api.pasteme.cn"
	var method = "POST"
	if once {
		url = "https://api.pasteme.cn/once"
	}
	if key != "" {
		url = "https://api.pasteme.cn/" + key
		method = "PUT"
	}

	bytesData, _ := json.Marshal(data)
	req, _ := http.NewRequest(method, url, bytes.NewReader(bytesData))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		fmt.Println("Something Wrong About Network")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	var response Resp
	err = json.Unmarshal(body, &response)
	if err != nil && response.Status != 201 {
		fmt.Println(err, response.Status)
		fmt.Println("😱Something Wrong About Network")

	} else if response.Status == 500 {
		if key != "" {
			fmt.Println("🤯Key May Exist")
			return
		} else {
			fmt.Println("🤯Something Wrong About Backend Server")
			return

		}
	} else if response.Status == 400 {
		fmt.Println("🥶Bad Paste")
		return

	} else if response.Status != 201 {
		fmt.Println(string(body), response.Status)

		fmt.Println("🤪Something Wrong...")
		return
	}
	var pkey = strconv.FormatUint(response.Key, 10)
	if key != "" {
		pkey = key
	}

	pasteUrl := fmt.Sprintf("https://pasteme.cn/%s", pkey)
	pasteAUrl := fmt.Sprintf("https://api.pasteme.cn/%s", pkey)
	if password != "" {
		pasteAUrl = pasteAUrl + "," + password
	}
	if raw {
		fmt.Print(pasteAUrl)
	} else {
		fmt.Println("🎉A Paste has created successfully!🎉 Visit Pasteme at")
		fmt.Println(fmt.Sprintf("%s or %s  for raw text.", pasteUrl, pasteAUrl))
		if copy {
			return
		}
		err = clipboard.WriteAll(pasteUrl)
		if err != nil {
			fmt.Println("😢Fail to Copy URL.\nOn Linux,this requires 'xclip' or 'xsel' command to be installed.")
		} else {
			fmt.Print("😛Web Address has copied to your keyboard.")
		}
	}
	return

}
