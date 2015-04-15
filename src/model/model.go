// model project model.go
package model

type User struct {
	Userid   int64  `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Changed  int    `json:"changed"`
	Picture  string `json:"picture"`
	Birthday string `json:"birthday"`
	Hobby    string `json:"hobby"`
	Friends  string `json:"friends"`
}
