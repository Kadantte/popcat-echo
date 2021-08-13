package pop

import "encoding/json"

type Pop struct {
	Count   int
	Address string
	Region  string
}

func NewPop(count int, address string, region string) *Pop {
	p := new(Pop)
	p.Count = count
	p.Address = address
	p.Region = region
	return p
}

func (p *Pop) JSON() []byte {
	val, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return val
}