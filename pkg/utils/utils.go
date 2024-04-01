package utils

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func ParseBody(r *http.Request, x interface{}) {
	var body, err = io.ReadAll(r.Body)

	if err == nil {
		err = json.Unmarshal(body, x)

		if err != nil {
			return
		}
	}
}

func SendJson(w http.ResponseWriter, statusCode int, b []byte) {
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(b)

	if err != nil {
		return
	}
}

func ExtractParamId(r *http.Request, name string) (int64, error) {
	vars := mux.Vars(r)
	id := vars[name]

	// Caso ocorrer erro retorna 0, err
	return strconv.ParseInt(id, 0, 0)
}
