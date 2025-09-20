// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressForStory interface {
	dara.Model
	String() string
	GoString() string
	SetCity(v string) *AddressForStory
	GetCity() *string
	SetCountry(v string) *AddressForStory
	GetCountry() *string
	SetDistrict(v string) *AddressForStory
	GetDistrict() *string
	SetProvince(v string) *AddressForStory
	GetProvince() *string
	SetTownship(v string) *AddressForStory
	GetTownship() *string
}

type AddressForStory struct {
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
	Country  *string `json:"Country,omitempty" xml:"Country,omitempty"`
	District *string `json:"District,omitempty" xml:"District,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Township *string `json:"Township,omitempty" xml:"Township,omitempty"`
}

func (s AddressForStory) String() string {
	return dara.Prettify(s)
}

func (s AddressForStory) GoString() string {
	return s.String()
}

func (s *AddressForStory) GetCity() *string {
	return s.City
}

func (s *AddressForStory) GetCountry() *string {
	return s.Country
}

func (s *AddressForStory) GetDistrict() *string {
	return s.District
}

func (s *AddressForStory) GetProvince() *string {
	return s.Province
}

func (s *AddressForStory) GetTownship() *string {
	return s.Township
}

func (s *AddressForStory) SetCity(v string) *AddressForStory {
	s.City = &v
	return s
}

func (s *AddressForStory) SetCountry(v string) *AddressForStory {
	s.Country = &v
	return s
}

func (s *AddressForStory) SetDistrict(v string) *AddressForStory {
	s.District = &v
	return s
}

func (s *AddressForStory) SetProvince(v string) *AddressForStory {
	s.Province = &v
	return s
}

func (s *AddressForStory) SetTownship(v string) *AddressForStory {
	s.Township = &v
	return s
}

func (s *AddressForStory) Validate() error {
	return dara.Validate(s)
}
