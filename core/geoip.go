package core

import (
	"log"
	"time"

	geoip2 "github.com/oschwald/geoip2-golang"
)

var geoDB *geoip2.Reader
var cache *Cache
var expiration time.Duration

func init() {
	defaultExpiration, _ := time.ParseDuration("1d")
	gcInterval, _ := time.ParseDuration("1h")
	expiration, _ = time.ParseDuration("10h")
	cache = NewCache(defaultExpiration, gcInterval)
}

func loadGeoIP(geoFile string) {
	db, err := geoip2.Open(geoFile)
	// defer db.Close()
	if err != nil {
		log.Printf("Could not open GeoIP database\n")
	}
	// log.Println("GeoIP inited.")
	geoDB = db
}
