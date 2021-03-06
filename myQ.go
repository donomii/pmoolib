package pmoo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"net/http"
)

func Receiver(url string, callback func([]byte)) {
	for {
		resp, err := http.Get(url + "/subscribe/main")
		if err != nil {
			//log.Fatalln(err)
		} else {
			data, _ := ioutil.ReadAll(resp.Body)
			callback(data)
		}
	}
}

func MyQMessage(url string, mess interface{}) {
	json, _ := json.Marshal(mess)
	Send(url, json)
}

func Send(url string, data []byte) {
	resp, err := http.Post(url+"/publish/main", "who/cares", bytes.NewReader(data))
	if err != nil {
		return
		//log.Fatalln(err)
	}
	defer resp.Body.Close()
}

func StoreObject(url, id string, m *Object) {
	b, _ := json.Marshal(m)
	resp, err := http.Post(url+"/store/"+id, "who/cares", bytes.NewReader(b))
	if err != nil {
		return
		//log.Fatalln(err)
	}
	defer resp.Body.Close()
}

func FetchObject(url, id string) *Object {
	resp, err := http.Get(url + "/fetch/" + id)
	if err != nil {
		return nil
		//log.Fatalln(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return nil
	}
	data, _ := ioutil.ReadAll(resp.Body)

	retrievedVal := new(Object)
	json.Unmarshal(data, retrievedVal)
	return retrievedVal
}
func DatabaseConnection(url string) bool {
	resp, err := http.Get(url + "/operational")
	if err != nil {
		return false
		//log.Fatalln(err)
	}
	defer resp.Body.Close()
	return resp.StatusCode < 299
}

/*
func DeleteObject(id string) {
	err := KVstore.Delete(id)
	if err != nil {
		panic(err)
	}
}
*/
