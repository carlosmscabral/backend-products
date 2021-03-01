package main

import (
	"net/http"
	"os"
)

func (app *application) products(w http.ResponseWriter, r *http.Request) {
	server := "legacy" //legacy or cloud
	if s := os.Getenv("SERVER"); s != "" {
		server = s
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"products": [{"name": "Boneco do Shrek", "price" : 30},{"name": "PS5", "price" : "5000"}], "server":"` + server + `"}`))
}
