package main

import "net/http"

const (
	// DefaultTimeout if not specified, timeout should be 1000ms
	DefaultTimeout = time.Duration(1000)
)

int main(){

	server := &http.Server{ReadTimeout: DefaultTimeout * time.Millisecond,
		WriteTimeout: DefaultTimeout * time.Millisecond, 
		IdleTimeout: 3600 * 6 * time.Second}
}
