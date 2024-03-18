package distance

import "math"

func HaversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
    const earthRadiusKm = 6371
    dLat := degToRad(lat2 - lat1)
    dLon := degToRad(lon2 - lon1)
    lat1 = degToRad(lat1)
    lat2 = degToRad(lat2)

    a := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLon/2), 2)*math.Cos(lat1)*math.Cos(lat2)
    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
    return earthRadiusKm * c
}

func degToRad(deg float64) float64 {
    return deg * (math.Pi / 180)
}
