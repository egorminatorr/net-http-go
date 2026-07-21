package middlewares

import (
	"net/http"
)


func LogMiddleware (next http.Handler) http.Handler {
	return http.Handle
}



