package main

import (
	"github.com/k0kubun/pp"
	"github.com/ken-aio/gomws/mws"
	"github.com/ken-aio/gomws/mws/products"
	"github.com/pkg/errors"
)

func main() {
	config := mws.Config{
		AccessKey: "access key",
		SecretKey: "secret key",
		Region:    "JP",
		SellerId:  "seller id",
	}
	client, err := products.NewClient(config)
	if err != nil {
		panic(errors.Wrap(err, "init amazon mws products api error"))
	}
	//resp, err := client.GetMyPriceForASIN([]string{"test1", "test2"})
	resp, err := client.GetMyFeesEstimate([]*products.PriceToEstimateFees{
		&products.PriceToEstimateFees{
			IDType:                   "ASIN",
			IDValue:                  "B07HCH85V6",
			IsAmazonFulfilled:        true,
			Identifier:               "test",
			ListingPriceCurrencyCode: "JPY",
			ListingPriceAmount:       22400,
			ShippingCurrencyCode:     "JPY",
			ShippingAmount:           0,
			PointsPointsNumber:       0,
		},
		&products.PriceToEstimateFees{
			IDType:                   "ASIN",
			IDValue:                  "B00E0GMMHO",
			IsAmazonFulfilled:        true,
			Identifier:               "test2",
			ListingPriceCurrencyCode: "JPY",
			ListingPriceAmount:       7980,
			ShippingCurrencyCode:     "JPY",
			ShippingAmount:           0,
			PointsPointsNumber:       0,
		},
	})
	if err != nil {
		panic(errors.Wrap(err, "init amazon mws products api error"))
	}
	defer resp.Close()
	if resp.Error != nil {
		panic(errors.Wrap(resp.Error, "mws response error"))
	}
	parser, err := resp.ResultParser()
	if err != nil {
		panic(errors.Wrap(err, "mws GetMatchingProductForID parser error"))
	}
	pp.Println("resp: ", parser)
}
