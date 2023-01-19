package backend

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello World from port "+os.Getenv("PORT"))
}

func Run(quote string) {
	godotenv.Load(".env")
	http.HandleFunc("/", HelloWorld)
	fmt.Println("Listening on port " + os.Getenv("PORT"))
	http.ListenAndServe(":8080", nil)
}
