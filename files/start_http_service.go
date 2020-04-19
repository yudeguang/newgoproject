package files

import (
	"log"
	"net/http"
)

func startHttpServer() {
	log.Println("启动HTTP服务:", httpport)
	//获得替换号码
	//http://localhost:8080/test?user=xxxx
	http.HandleFunc("/test", handle_test)

	//其它未实现搜索
	http.HandleFunc("/", http.NotFound)
	err := http.ListenAndServe(httpport, nil)
	if err != nil {
		log.Println("启动HTTP服务失败:", err)
	}
}
