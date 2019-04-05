package entities

import (
	"errors"
)

type Catalog struct {
	ID		int32
	Name	string
	product	[]*Product
}

var (
	InvalidNameErr = errors.New("Name is not valid")
)

func NewCatalog() *Catalog {
	return &Catalog {
		product : make([]*Product,0),
	}
}

func(c *Catalog) Addproduct(name,description string, price float32) error {

	if name == ""{
		return InvalidNameErr
	}
	if len(name)>10 {
		return InvalidNameErr
	}


	return nil
}