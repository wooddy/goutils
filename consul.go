package goutils

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"encoding/base64"
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

type ConsulKey struct {
	Key         string
	CreateIndex uint64
	ModifyIndex uint64
	LockIndex   uint64
	Flags       uint64
	Value       string
	Session     string
}
