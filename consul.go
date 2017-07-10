package goutils

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"errors"
	"fmt"
)

func GetConfigFromConsul(url, key, token string) (configContent []byte, err error) {
	req, err := http.NewRequest("GET", url+"/v1/kv/"+key, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Consul-Token", token)
	cl := http.DefaultClient
	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	} else {
		if resp.StatusCode != 200 {
			return nil, errors.New(fmt.Sprintf("failed get config. server response code:%s", resp.StatusCode))
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		consulKey := make([]ConsulKey, 1)
		err = json.Unmarshal(body, &consulKey)
		if err != nil {
			return nil, err
		}
		var keyValueContent []byte = make([]byte, len(consulKey[0].Value)*3/4)
		_, err = base64.StdEncoding.Decode(keyValueContent, []byte(consulKey[0].Value))
		if err != nil {
			return nil, err
		}
		return keyValueContent, nil
	}
}

func GetConfigMapFromConsul(url, key, token string) (Config, error){
	configInBytes, err := GetConfigFromConsul(url, key, token)
	if err != nil {
		return nil, err
	}
	return ParseProperties(configInBytes)
}

type ConsulKey struct {
	Key         string
	CreateIndex uint64
	ModifyIndex uint64
	LockIndex   uint64
	Flags       uint64
	Value       string
	Session     string
}
