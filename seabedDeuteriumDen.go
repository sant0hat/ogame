package ogame

// SeabedDeuteriumDen ...
type seabedDeuteriumDen struct {
	BaseBuilding
}

func newSeabedDeuteriumDen() *seabedDeuteriumDen {
	b := new(seabedDeuteriumDen)
	b.ID = SeabedDeuteriumDenID
	b.IncreaseFactor = 2.3
	b.BaseCost = Resources{Metal: 2645, Crystal: 2645}
	return b
}