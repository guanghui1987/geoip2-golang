package main

import (
	"fmt"
	"github.com/guanghui1987/geoip2-golang"
	"github.com/guanghui1987/maxminddb-golang"
	"net"
	"time"
)

func main() {
	maxminddb.Open("")
	geoDB, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		return
	}

	ip := net.ParseIP("114.114.114.114")
	var names map[string]string
	var name string
	country_struct := geoip2.Country{}
	country_Country := geoip2.Country{}
	geoDB.GetField(ip, "", []string{}, &country_struct)
	geoDB.GetField(ip, "", []string{"country"}, &country_Country.Country)
	geoDB.GetField(ip, "", []string{"country", "names"}, &names)
	geoDB.GetField(ip, "en", []string{"country", "names"}, &name)

	var country string
	delta1 := time.Now().UnixNano()
	for i:=0; i < 1000000; i++ {
		//if err = geoDB.GetField(ip, "en", []string{"country", "names"}, &country); err == nil {
		//	//country = record.Country.Names["en"]
		//}
		if record, err := geoDB.Country(ip); err != nil {
			country = record.Country.Names["en"]
		}

	}
	delta2 := time.Now().UnixNano()
	delta := (delta2 - delta1) / int64(time.Millisecond)
	fmt.Printf("Country method query country = %s delta2 = %d   delta1 = %d     time = %d ms \n", country, delta2, delta1,  delta)

	delta1 = time.Now().UnixNano()
	for i:=0; i < 1000000; i++ {
		if err = geoDB.GetField(ip, "en", []string{"country", "names"}, &country); err == nil {
			//country = record.Country.Names["en"]
		}
	}
	delta2 = time.Now().UnixNano()
	delta = (delta2 - delta1) / int64(time.Millisecond)
	fmt.Printf("GetField method query country = %s delta2 = %d   delta1 = %d     time = %d ms\n", country, delta2, delta1,  delta)
}
