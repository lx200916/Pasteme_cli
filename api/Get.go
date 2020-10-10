package api

import (
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Cresp struct {
	Status  int
	Content string
	lang    string
}

func Get(password string, key string, copy bool) {
	if key == "" {
		fmt.Println("👋Please at least Input the Key of Paste")
		return
	}
	flag, _ := regexp.MatchString("[a-z]", key)

	if (len(key) < 2 && len(key) > 9) && flag {
		fmt.Println("👋Key should Be 3-8 letters long.")
		return
	}
	url := "https://api.pasteme.cn/" + key
	if password != "" {
		url = url + "," + password
	}
	resp, err := http.Get(url + "?json=true")
	if err != nil {
		log.Println(err)
		fmt.Println("Something Wrong About Network")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var response Cresp
	err = json.Unmarshal(body, &response)
	if err != nil && response.Status != 200 {
		fmt.Println(err, response.Status)
		fmt.Println("😱Something Wrong About Network")
		return
	}
	if response.Status == 404 {
		fmt.Println("🥶Key Not Exists")
		return
	} else if response.Status == 500 {
		fmt.Println("🤯Something Wrong About Backend Server")
		return
	} else if response.Status == 401 {
		fmt.Println("🤯Password Protected")
		return
	} else if response.Status != 200 {
		fmt.Println(string(body), response.Status)

		fmt.Println("🤪Something Wrong...")
		return
	}
	fmt.Print(response.Content)
	if !copy {
		err = clipboard.WriteAll(response.Content)
	}

}
