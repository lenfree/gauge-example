package stepimpl

import (
	"fmt"
	"strconv"

	"github.com/getgauge-contrib/gauge-go/gauge"
	. "github.com/getgauge-contrib/gauge-go/testsuit"
	"github.com/hashicorp/go-retryablehttp"
)

type me struct {
	StatusCode int `json:"statusCode"`
}

var _ = gauge.Step("If I make a <action> query to <url>.", func(action, URL string) {
	res, _ := retryablehttp.Get("http://google.com")

	storeStringToSpecDataStore("statusCode", strconv.Itoa(res.StatusCode))
	//gauge.GetSpecStore()["statusCode"] = "200" //res.StatusCode
})

var _ = gauge.Step("The status code should return <expectedCount>.", func(expected string) {
	//value := gauge.GetSuiteStore()["statusCode"].(string)
	value := fetchStringFromSpecDataStore("statusCode")
	if value != expected {
		T.Fail(fmt.Errorf("got: %v, want: %v", value, expected))
	}
})

func storeStringToSpecDataStore(key, value string) {
	gauge.GetSpecStore()[key] = value
}

func fetchStringFromSpecDataStore(key string) string {
	return gauge.GetSpecStore()[key].(string)
}
