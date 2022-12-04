package authkeycloaklib

import (
	"fmt"
	"testing"

	"github.com/johnkeychishugi/auth-go-cloak/constants"
	"github.com/rs/xid"
)

var (
	userToken   string
	adminToken  string
	clientToken string
)

func TestLibary(t *testing.T) {
	t.Run("Client Login", func(t *testing.T) {

		_, err := LoginClient(
			constants.Hostname,
			constants.Realm,
			constants.Client_id,
			constants.Client_secret,
		)

		// Assert that you got your expected response
		if err != "" {
			t.Fail()
		}
	})

	t.Run("Client Login Failed", func(t *testing.T) {

		_, err := LoginClient(
			constants.Hostname,
			constants.Realm,
			constants.Client_id,
			"jhfjhfjshsjfshfjs",
		)

		// Assert that you got your expected response
		if err == "" {
			t.Fail()
		}
	})

	t.Run("Login User", func(t *testing.T) {

		result, err := LoginUser(
			constants.Hostname,
			constants.Realm,
			constants.Client_id,
			constants.Client_secret,
			constants.Username,
			constants.Password,
		)

		userToken = result.RefreshToken

		// Assert that you got your expected response
		if err != "" {
			t.Fail()
		}
	})

	t.Run("Login User Failed", func(t *testing.T) {

		_, err := LoginUser(
			constants.Hostname,
			constants.Realm,
			constants.Client_id,
			constants.Client_secret,
			constants.Username,
			"jhfjhfjshsjfshfjs",
		)

		// Assert that you got your expected response
		if err == "" {
			t.Fail()
		}
	})

	t.Run("Login Admin", func(t *testing.T) {

		result, err := LoginAdmin(
			constants.Hostname,
			constants.AdminRealm,
			constants.AdminUsername,
			constants.AdminPassword,
		)

		adminToken = result.AccessToken
		// Assert that you got your expected response
		fmt.Println(err)
		if err != "" {
			t.Fail()
		}
	})

	t.Run("Login Admin Failed", func(t *testing.T) {

		_, err := LoginAdmin(
			constants.Hostname,
			constants.AdminRealm,
			constants.AdminUsername,
			"jhfjhfjshsjfshfjs",
		)

		// Assert that you got your expected response
		if err == "" {
			t.Fail()
		}
	})

	t.Run("Logout", func(t *testing.T) {

		got := Logout(
			constants.Hostname,
			constants.Realm,
			constants.Client_id,
			constants.Client_secret,
			userToken,
		)

		// Assert that you got your expected respons
		if got != "Successully logged out" {
			t.Fail()
		}
	})

	t.Run("Logout failed", func(t *testing.T) {

		got := Logout(
			constants.Hostname,
			constants.Realm,
			constants.Client_id,
			constants.Client_secret,
			"jhrighureitreiute",
		)

		// Assert that you got your expected respons

		if got == "Successully logged out" {
			t.Fail()
		}
	})

	t.Run("Create User", func(t *testing.T) {

		uniqueString := xid.New().String()

		got := CreateUser(
			constants.Hostname,
			constants.Realm,
			adminToken,
			"firstname_"+uniqueString,
			"lastname_"+uniqueString,
			"username_"+uniqueString,
			"password_"+uniqueString,
			"email_"+uniqueString+"@gmail.com",
		)

		// Assert that you got your expected respons
		if got == "" {
			t.Fail()
		}
	})

	t.Run("Create User failed", func(t *testing.T) {

		uniqueString := xid.New().String()

		got := CreateUser(
			constants.Hostname,
			constants.Realm,
			adminToken,
			"firstname_"+uniqueString,
			"lastname_"+uniqueString,
			"username_"+uniqueString,
			"password_"+uniqueString,
			"email_"+uniqueString+"gmailcom",
		)

		// Assert that you got your expected respons
		if got == "" {
			t.Fail()
		}
	})

	t.Run("Check Token", func(t *testing.T) {

		_, err := CheckToken(
			constants.Hostname,
			constants.Realm,
			constants.Client_id,
			constants.Client_secret,
			clientToken,
		)

		// Assert that you got your expected respons
		fmt.Println(err)
		if err != "" {
			t.Fail()
		}
	})

	t.Run("Check Token failed", func(t *testing.T) {

		_, err := CheckToken(
			constants.Hostname,
			constants.Realm,
			constants.Client_id,
			constants.Client_secret,
			"kdfjldfkd",
		)

		// Assert that you got your expected respons
		fmt.Println(err)
		if err != "" {
			t.Fail()
		}
	})
}
