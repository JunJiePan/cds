package handler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

type response struct {
	Code    int64       `json:"code"`
	Desc    string      `json:"desc,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

func HttpError(w http.ResponseWriter, httpCode int, appCode int64, desc string, message interface{}) {
	httpx.WriteJson(w, httpCode, response{
		Code:    appCode,
		Desc:    desc,
		Message: message,
	})
}

func HttpOk(w http.ResponseWriter, data interface{}) {
	httpx.WriteJson(w, http.StatusOK, response{
		Data: data,
	})
}