package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"tony": {
		AuthToken: "123ABC",
		Username:  "tony",
	},
	"coach": {
		AuthToken: "456DEF",
		Username:  "coach",
	},
	"mark": {
		AuthToken: "789GHI",
		Username:  "mark",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"tony": {
		Coins:    1,
		Username: "tony",
	},
	"coach": {
		Coins:    2,
		Username: "coach",
	},
	"mark": {
		Coins:    3,
		Username: "mark",
	},
}

func (d *mockDB) GetUserLogInDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}

	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
