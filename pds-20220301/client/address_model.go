// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddress interface {
	dara.Model
	String() string
	GoString() string
	SetCity(v string) *Address
	GetCity() *string
	SetCountry(v string) *Address
	GetCountry() *string
	SetDistrict(v string) *Address
	GetDistrict() *string
	SetProvince(v string) *Address
	GetProvince() *string
	SetTownship(v string) *Address
	GetTownship() *string
}

type Address struct {
	// The city.
	//
	// example:
	//
	// None
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// The country or region.
	//
	// example:
	//
	// None
	Country *string `json:"country,omitempty" xml:"country,omitempty"`
	// The district.
	//
	// example:
	//
	// None
	District *string `json:"district,omitempty" xml:"district,omitempty"`
	// The province.
	//
	// example:
	//
	// None
	Province *string `json:"province,omitempty" xml:"province,omitempty"`
	// The street.
	//
	// example:
	//
	// None
	Township *string `json:"township,omitempty" xml:"township,omitempty"`
}

func (s Address) String() string {
	return dara.Prettify(s)
}

func (s Address) GoString() string {
	return s.String()
}

func (s *Address) GetCity() *string {
	return s.City
}

func (s *Address) GetCountry() *string {
	return s.Country
}

func (s *Address) GetDistrict() *string {
	return s.District
}

func (s *Address) GetProvince() *string {
	return s.Province
}

func (s *Address) GetTownship() *string {
	return s.Township
}

func (s *Address) SetCity(v string) *Address {
	s.City = &v
	return s
}

func (s *Address) SetCountry(v string) *Address {
	s.Country = &v
	return s
}

func (s *Address) SetDistrict(v string) *Address {
	s.District = &v
	return s
}

func (s *Address) SetProvince(v string) *Address {
	s.Province = &v
	return s
}

func (s *Address) SetTownship(v string) *Address {
	s.Township = &v
	return s
}

func (s *Address) Validate() error {
	return dara.Validate(s)
}
