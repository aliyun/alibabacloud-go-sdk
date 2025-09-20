// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddress interface {
	dara.Model
	String() string
	GoString() string
	SetAddressLine(v string) *Address
	GetAddressLine() *string
	SetCity(v string) *Address
	GetCity() *string
	SetCountry(v string) *Address
	GetCountry() *string
	SetDistrict(v string) *Address
	GetDistrict() *string
	SetLanguage(v string) *Address
	GetLanguage() *string
	SetProvince(v string) *Address
	GetProvince() *string
	SetTownship(v string) *Address
	GetTownship() *string
}

type Address struct {
	AddressLine *string `json:"AddressLine,omitempty" xml:"AddressLine,omitempty"`
	City        *string `json:"City,omitempty" xml:"City,omitempty"`
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	District    *string `json:"District,omitempty" xml:"District,omitempty"`
	Language    *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Township    *string `json:"Township,omitempty" xml:"Township,omitempty"`
}

func (s Address) String() string {
	return dara.Prettify(s)
}

func (s Address) GoString() string {
	return s.String()
}

func (s *Address) GetAddressLine() *string {
	return s.AddressLine
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

func (s *Address) GetLanguage() *string {
	return s.Language
}

func (s *Address) GetProvince() *string {
	return s.Province
}

func (s *Address) GetTownship() *string {
	return s.Township
}

func (s *Address) SetAddressLine(v string) *Address {
	s.AddressLine = &v
	return s
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

func (s *Address) SetLanguage(v string) *Address {
	s.Language = &v
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
