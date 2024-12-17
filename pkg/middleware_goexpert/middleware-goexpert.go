package middleware_goexpert

import (
	"fmt"
	"net/http"
)


func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add rate limiter logic here
		
		fmt.Println("Teste rate limiter alterado")
		
		next.ServeHTTP(w, r)
	})
}