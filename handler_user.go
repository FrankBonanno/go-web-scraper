package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/FrankBonanno/go-web-scraper/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	
	decoder := json.NewDecoder(r.Body)
	
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		resondWithError(w, 400, fmt.Sprintf("Error parsing json: %s", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		resondWithError(w, 400, fmt.Sprintf("Couldn't create user: %s", err))
		return
	}

	respondWithJson(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil {
		resondWithError(w, 400, fmt.Sprintf("Couldn't find posts for user: %s", err))
		return
	}

	respondWithJson(w, 200, databasePostsToPosts(posts))
}