// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModulePriceMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetPrice(v int64) *ModulePriceMapValue
	GetPrice() *int64
	SetServiceNo(v string) *ModulePriceMapValue
	GetServiceNo() *string
}

type ModulePriceMapValue struct {
	// The price, in cents.
	//
	// example:
	//
	// 10000
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// The service number, such as a flight number or train number.
	//
	// example:
	//
	// MU1234
	ServiceNo *string `json:"service_no,omitempty" xml:"service_no,omitempty"`
}

func (s ModulePriceMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModulePriceMapValue) GoString() string {
	return s.String()
}

func (s *ModulePriceMapValue) GetPrice() *int64 {
	return s.Price
}

func (s *ModulePriceMapValue) GetServiceNo() *string {
	return s.ServiceNo
}

func (s *ModulePriceMapValue) SetPrice(v int64) *ModulePriceMapValue {
	s.Price = &v
	return s
}

func (s *ModulePriceMapValue) SetServiceNo(v string) *ModulePriceMapValue {
	s.ServiceNo = &v
	return s
}

func (s *ModulePriceMapValue) Validate() error {
	return dara.Validate(s)
}
