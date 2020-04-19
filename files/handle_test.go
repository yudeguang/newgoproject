package files

import (
	"io"
	"net/http"
)

func handle_test(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	test := req.FormValue("test")

	io.WriteString(rw, test)
}
