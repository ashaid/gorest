package middleware

import (
	"errors"
	"net/http"

	"github.com/ashaid/gorest/api"
	"github.com/ashaid/gorest/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthoriziedError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthoriziedError)
			api.RequestErrorHandler(w, UnAuthoriziedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()

		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLogInDetails(username)

		if loginDetails == nil || token != (*loginDetails).AuthToken {
			log.Error(UnAuthoriziedError)
			api.RequestErrorHandler(w, UnAuthoriziedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
