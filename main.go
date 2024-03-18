package main

import (
    "fmt"
    "geolocation-service/geocoder"
    "geolocation-service/distance"
)

func main() {
    address := "1600 Amphitheatre Parkway, Mountain View, CA"
    lat, lng, err := geocoder.Geocode(address)
    if err != nil {
        fmt.Println("Error geocoding address:", err)
        return
    }
    fmt.Printf("Latitude: %f, Longitude: %f\n", lat, lng)

    distanceInKm := distance.HaversineDistance(lat, lng, 37.7749, -122.4194)
    fmt.Printf("Distance to San Francisco: %.2f km\n", distanceInKm)
}
