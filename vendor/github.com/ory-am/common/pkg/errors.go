package pkg

import (
	"github.com/go-errors/errors"
	"net/http"
)

var (
	ErrNotFound       = errors.New("Not found")
	ErrInvalidPayload = errors.New("Invalid payload")
)

func WriteError(w http.ResponseWriter, err error) {
	if err == ErrNotFound {
		LogError(err, http.StatusNotFound)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err == ErrInvalidPayload {
		LogError(err, http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	LogError(err, http.StatusInternalServerError)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func HttpError(rw http.ResponseWriter, err error, code int) {
	LogError(err, code)
	http.Error(rw, err.Error(), code)
}

func GetErrorStack(err interface{}) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(*errors.Error); ok {
		return e.ErrorStack()
	}
	return ""
}

func GetErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}