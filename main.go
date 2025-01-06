package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type URL struct {
	ID           int    `json:"id"`
	OriginalURL  string `json:"originalUrl"`
	ShortenedURL string `json:"shortenedUrl"`
}

type URLInput struct {
	OriginalURL string `json:"originalUrl"`
	Alias       string `json:"alias"`
}

func main() {
	// connect to databases
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS urls(id SERIAL PRIMARY KEY, original_url TEXT NOT NULL, shortened_url TEXT UNIQUE NOT NULL)")

	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		user := User{ID: 1, Name: "hellllo", Email: "abc@gmail.com"}
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("POST /shorten", func(w http.ResponseWriter, r *http.Request) {
		var input URLInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if input.OriginalURL == "" {
			http.Error(w, "Original URL is required", http.StatusBadRequest)
			return
		}

		if _, err := url.ParseRequestURI(input.OriginalURL); err != nil {
			http.Error(w, "Invalid URL format", http.StatusBadRequest)
			return
		}

		alias := input.Alias
		if alias == "" {
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(input.OriginalURL)))[:8]
			alias = hash
		}

		shortenedURL := fmt.Sprintf("http://chibi.url/%s", alias)

		query := `
			INSERT INTO urls (original_url, shortened_url) 
			VALUES ($1, $2) 
			RETURNING id`

		var urlID int
		err := db.QueryRow(query, input.OriginalURL, shortenedURL).Scan(&urlID)
		if err != nil {
			log.Printf("Database error: %v", err)
			http.Error(w, "Error saving URL", http.StatusInternalServerError)
			return
		}

		response := URL{
			ID:           urlID,
			OriginalURL:  input.OriginalURL,
			ShortenedURL: shortenedURL,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	serv := &http.Server{Handler: jsonContentTypeMiddleware(mux), Addr: ":8080"}
	serv.ListenAndServe()
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("listening to server: 8080")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
