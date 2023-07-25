package entity

import "errors"

type Order struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.Validate()

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o Order) Validate() error {
	if o.ID == "" {
		return errors.New("id is required")
	}

	if o.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}

	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax

	err := o.Validate()

	if err != nil {
		return err
	}

	return nil
}
