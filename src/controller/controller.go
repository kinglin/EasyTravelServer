// controller project controller.go
package controller

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"model"
	"net"
	"strconv"
	"time"
)

func HandleRegister(requestFromClient string, connect net.Conn) {

	type Response struct {
		RegisterResult string `json:"registerResult"`
	}

	var newUser = new(model.User)

	if err := json.Unmarshal([]byte(requestFromClient), &newUser); err == nil {

		response := &Response{
			RegisterResult: "yes"}

		jsonResponse, _ := json.Marshal(response)
		fmt.Println(string(jsonResponse))
		connect.Write([]byte(string(jsonResponse)))

		newUser.Userid = time.Now().UnixNano() / 1000000
		newUser.Changed = 0
		newUser.Picture = ""
		newUser.Hobby = ""
		newUser.Birthday = ""
		newUser.Friends = ""

		session, err := mgo.Dial("127.0.0.1:27017")
		if err != nil {
			panic(err)
		}
		defer session.Close()

		session.SetMode(mgo.Monotonic, true)

		collection := session.DB(strconv.FormatInt(newUser.Userid, 10)).C("user")

		collection.Insert(&model.User{
			newUser.Userid, newUser.Username, newUser.Password,
			newUser.Changed, newUser.Picture, newUser.Birthday,
			newUser.Hobby, newUser.Friends})

	}
}
