package route

type Position struct {
	Lat float64
	Lon float64
}

type Route struct {
	ID        string
	ClientID  string
	Positions []Position
}
