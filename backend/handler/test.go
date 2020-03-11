package handler

import (
	"encoding/binary"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// HandleTest to return current timestamp
func HandleTest(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	body = fmt.Sprint("now is ", now.String())
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(binary.Size(body)))
	w.Write(body)
}
