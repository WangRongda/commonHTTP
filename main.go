package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type ResKitData struct {
	Code   int    `json:"code"`
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type myHTTP interface {
	post()
	get()
	put()
	delete()
	patch()

	process()
	send()

	// postD() interface
}

type commonHTTP struct {
	instant myHTTP
	r       *http.Request
	w       http.ResponseWriter

	_reqData interface{}
	_resData interface{}
}

func (c *commonHTTP) process() {
	// fmt.Println(reflect.TypeOf(c).Elem().Field(0))
	foo := reflect.ValueOf(c.instant).Elem().Field(1).Field(0).Field(0)
	fmt.Println(foo)
	// zz := reflect.New(foo)
	// fmt.Println(zz)
	// zz := reflect.TypeOf(c.instant).Elem()
	// for i := 0; i < zz.NumField(); i++ {
	// fmt.Println(zz.Field(i))
	// }

	// fmt.Println(foo)
	// fmt.Println(reflect.TypeOf(foo.Interface()).FieldByName("post"))
	switch c.r.Method {
	case "POST":
		reqBody, err := ioutil.ReadAll(c.r.Body)
		if nil != err {
			fmt.Println(err)
		}
		if err := json.Unmarshal(reqBody, c._reqData); nil != err {
			fmt.Println(err)
		}
		c.instant.post()
	case "GET":
		c.instant.get()
	case "PUT":
		c.instant.put()
	case "DELETE":
		c.instant.delete()
	case "PATCH":
		c.instant.patch()
	case "OPTIONS":
	default:
		fmt.Println("request method invalid")
	}
}

func (c commonHTTP) send() {

}

func (c commonHTTP) get() {
	fmt.Println("get")
}
func (c commonHTTP) post() {
	fmt.Println("post")
	fmt.Println(c._reqData)
}
func (c commonHTTP) put() {
	fmt.Println("put")
}
func (c commonHTTP) delete() {
	fmt.Println("delete")
}
func (c commonHTTP) patch() {
	fmt.Println("patch")
}

// func newGroup() grouper {
// 	return grouper{}
// }

type grouper struct {
	commonHTTP
	reqData struct {
		post struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
	}
	resData struct {
		post struct {
			ResKitData
			Result struct {
				ID uint `json:"id"`
			} `json:"result"`
		}
	}
	test string
}

func (g grouper) post() {
	fmt.Println("post")
	fmt.Println(g._reqData)
	fmt.Println(g.reqData.post)
}

func group(w http.ResponseWriter, r *http.Request) *grouper {
	foo := &grouper{}
	foo.r, foo.w, foo.instant = r, w, foo
	foo._reqData = &foo.reqData.post
	foo._resData = foo.resData
	// foo.postD = &foo.postData.reqData
	foo.instant = foo
	foo.test = "heelo"
	return foo
}

func groupFunc(w http.ResponseWriter, r *http.Request) {
	// foo := newGroup()
	foo := group(w, r)
	foo.process()
}

func main() {
	http.HandleFunc("/test", groupFunc)
	http.ListenAndServe(":8888", nil)
}
