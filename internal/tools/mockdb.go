package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"indy": {
		AuthToken: "123ABC",
		Username:  "indy",
	},
	"megan": {
		AuthToken: "123",
		Username:  "megan",
	},
	"riley": {
		AuthToken: "ABC",
		Username:  "riley",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"indy": {
		Coins:    100,
		Username: "indy",
	},
	"megan": {
		Coins:    200,
		Username: "megan",
	},
	"riley": {
		Coins:    150,
		Username: "riley",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
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
