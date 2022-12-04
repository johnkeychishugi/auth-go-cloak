# Auth Keycloak Lib

Auth Keycloak Lib it's Go package to simplify the authentification and authorization to Keycloak using the Go-cloak package

### Installation 

- `go get bitbucket.org/dnda-tech/auth-keycloak-lib`
- import package `import bitbucket.org/dnda-tech/auth-keycloak-lib`

## Features

- LoginClient(hostname string, realm string, clientID string, clientSecret string).
- LoginUser(hostname string, realm string, clientID string, clientSecret string, username string, password string)
- LoginAdmin(hostname string, realm string, username string, password string) (*gocloak.JWT, string)
- Logout(hostname string, realm string, clientID string, clientSecret string, refreshToken string)
- CreateUser(hostname string, realm string, token string, firstname string, lastname string, username string, password string, email string)
- CheckToken(hostname string, realm string, clientID string, clientSecret string, token string)

# License

MIT