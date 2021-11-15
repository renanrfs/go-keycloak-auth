package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

// keycloak app variables
var (
	// the Clients section
	clientID     = "app"
	// the Clients -> app -> Credentials -> Secret section
	clientSecret = "431522db-b9e0-4551-bd62-c1981f865c94"
)

func main() {

	log.Printf("GO-KEYCLOACK-AUTH is starting...")
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/demo")
	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8081/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	// CSRF - cross site request forgery
	state := "example"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)
	})

	http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			http.Error(w, "State doesn't match", http.StatusBadRequest)
			return
		}

		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "State doesn't match", http.StatusBadRequest)
			return
		}

		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "Id token doesn't found", http.StatusBadRequest)
			return
		}

		res := struct {
			OAuth2Token *oauth2.Token
			IDToken     string
		}{
			oauth2Token, rawIDToken,
		}

		data, _ := json.MarshalIndent(res, "", " ")
		w.Write(data)
	})

	log.Printf("Server is ready to handle requests at 8081 port!")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
