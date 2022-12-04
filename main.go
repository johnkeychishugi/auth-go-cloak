package authkeycloaklib

import (
	"context"

	"github.com/Nerzal/gocloak/v12"
)

var ctx = context.Background()

func LoginClient(hostname string, realm string, clientID string, clientSecret string) (*gocloak.JWT, string) {
	client := gocloak.NewClient(hostname)

	result, err := client.LoginClient(ctx, clientID, clientSecret, realm)
	if err != nil {
		return nil, err.Error()
	} else {
		return result, ""
	}
}

func LoginUser(hostname string, realm string, clientID string, clientSecret string, username string, password string) (*gocloak.JWT, string) {
	client := gocloak.NewClient(hostname)

	result, err := client.Login(ctx, clientID, clientSecret, realm, username, password)
	if err != nil {
		return nil, err.Error()
	} else {
		return result, ""
	}
}

func LoginAdmin(hostname string, realm string, username string, password string) (*gocloak.JWT, string) {
	client := gocloak.NewClient(hostname)

	result, err := client.LoginAdmin(ctx, username, password, realm)
	if err != nil {
		return nil, err.Error()
	} else {
		return result, ""
	}
}

func Logout(hostname string, realm string, clientID string, clientSecret string, refreshToken string) string {
	client := gocloak.NewClient(hostname)

	err := client.Logout(ctx, clientID, clientSecret, realm, refreshToken)
	if err != nil {
		return err.Error()
	} else {
		return "Successully logged out"
	}
}

func CreateUser(hostname string, realm string, token string, firstname string, lastname string, username string, password string, email string) string {
	client := gocloak.NewClient(hostname)

	credentials := []gocloak.CredentialRepresentation{
		{
			Type:      gocloak.StringP("password"),
			Value:     gocloak.StringP(password),
			Temporary: gocloak.BoolP(false),
		},
	}

	user := gocloak.User{
		FirstName:   gocloak.StringP(firstname),
		LastName:    gocloak.StringP(lastname),
		Email:       gocloak.StringP(email),
		Enabled:     gocloak.BoolP(true),
		Username:    gocloak.StringP(username),
		Credentials: &credentials,
	}

	result, err := client.CreateUser(ctx, token, realm, user)
	if err != nil {
		return err.Error()
	} else {
		return result
	}
}

func CheckToken(hostname string, realm string, clientID string, clientSecret string, token string) (*gocloak.IntroSpectTokenResult, string) {
	client := gocloak.NewClient(hostname)

	result, err := client.RetrospectToken(ctx, token, clientID, clientSecret, realm)
	if err != nil {
		return nil, err.Error()
	} else {
		return result, ""
	}
}
