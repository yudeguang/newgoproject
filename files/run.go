package files

import (
	"github.com/yudeguang/config"
	"github.com/yudeguang/haserr"
	"github.com/yudeguang/mysql"
	"log"
)

var sql *mysql.MySqlStruct
var httpport string

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	var err error
	conf, err := config.NewConfig("")
	haserr.Fatal(err)
	sql, err = mysql.Open(conf.Get("dbconn"))
	haserr.Fatal(err)
	//
	httpport = conf.Get("httpport")

}

func Run() {
	log.Println("开始运行")
}
