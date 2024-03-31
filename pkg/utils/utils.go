package utils

import (
	"encoding/json"
	"io"
	"net/http"
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
