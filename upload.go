package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type Payload struct {
	Message string `json:"message"`
	Content string `json:"content"`
}

func Upload(filePath string) (url string, err error) {
	fileName := fmt.Sprintf("%s-%s", strconv.Itoa(int(time.Now().Unix())), filepath.Base(filePath))
	api := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", Owner, Repo, fileName)

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	payload := &Payload{
		Message: fmt.Sprintf("Create file: %s", fileName),
		Content: base64.StdEncoding.EncodeToString(fileData),
	}
	reqData, err := json.Marshal(payload)
	if err != nil {
		return
	}
	req, err := http.NewRequest("PUT", api, bytes.NewReader(reqData))
	if err != nil {
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("token %s", Token))

	var client = &http.Client{
		Timeout: 1 * time.Minute,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	response := &CreateResponseJson{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}
	// fmt.Printf("%v %d", response, resp.StatusCode)
	if resp.StatusCode != 201 {
		err = fmt.Errorf("返回码错误: %d", resp.StatusCode)
		return
	}
	url = fmt.Sprintf("https://cdn.jsdelivr.net/gh/%s/%s/%s", Owner, Repo, fileName)
	return
}
