// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountAddressInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *UpdateAccountAddressInfoShrinkRequest
	GetAddress() *string
	SetAddress2(v string) *UpdateAccountAddressInfoShrinkRequest
	GetAddress2() *string
	SetCityJsonStringShrink(v string) *UpdateAccountAddressInfoShrinkRequest
	GetCityJsonStringShrink() *string
	SetDistrictJsonStringShrink(v string) *UpdateAccountAddressInfoShrinkRequest
	GetDistrictJsonStringShrink() *string
	SetPK(v string) *UpdateAccountAddressInfoShrinkRequest
	GetPK() *string
	SetPostCode(v string) *UpdateAccountAddressInfoShrinkRequest
	GetPostCode() *string
	SetProvinceJsonStringShrink(v string) *UpdateAccountAddressInfoShrinkRequest
	GetProvinceJsonStringShrink() *string
}

type UpdateAccountAddressInfoShrinkRequest struct {
	Address                  *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Address2                 *string `json:"Address2,omitempty" xml:"Address2,omitempty"`
	CityJsonStringShrink     *string `json:"CityJsonString,omitempty" xml:"CityJsonString,omitempty"`
	DistrictJsonStringShrink *string `json:"DistrictJsonString,omitempty" xml:"DistrictJsonString,omitempty"`
	PK                       *string `json:"PK,omitempty" xml:"PK,omitempty"`
	PostCode                 *string `json:"PostCode,omitempty" xml:"PostCode,omitempty"`
	ProvinceJsonStringShrink *string `json:"ProvinceJsonString,omitempty" xml:"ProvinceJsonString,omitempty"`
}

func (s UpdateAccountAddressInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountAddressInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountAddressInfoShrinkRequest) GetAddress() *string {
	return s.Address
}

func (s *UpdateAccountAddressInfoShrinkRequest) GetAddress2() *string {
	return s.Address2
}

func (s *UpdateAccountAddressInfoShrinkRequest) GetCityJsonStringShrink() *string {
	return s.CityJsonStringShrink
}

func (s *UpdateAccountAddressInfoShrinkRequest) GetDistrictJsonStringShrink() *string {
	return s.DistrictJsonStringShrink
}

func (s *UpdateAccountAddressInfoShrinkRequest) GetPK() *string {
	return s.PK
}

func (s *UpdateAccountAddressInfoShrinkRequest) GetPostCode() *string {
	return s.PostCode
}

func (s *UpdateAccountAddressInfoShrinkRequest) GetProvinceJsonStringShrink() *string {
	return s.ProvinceJsonStringShrink
}

func (s *UpdateAccountAddressInfoShrinkRequest) SetAddress(v string) *UpdateAccountAddressInfoShrinkRequest {
	s.Address = &v
	return s
}

func (s *UpdateAccountAddressInfoShrinkRequest) SetAddress2(v string) *UpdateAccountAddressInfoShrinkRequest {
	s.Address2 = &v
	return s
}

func (s *UpdateAccountAddressInfoShrinkRequest) SetCityJsonStringShrink(v string) *UpdateAccountAddressInfoShrinkRequest {
	s.CityJsonStringShrink = &v
	return s
}

func (s *UpdateAccountAddressInfoShrinkRequest) SetDistrictJsonStringShrink(v string) *UpdateAccountAddressInfoShrinkRequest {
	s.DistrictJsonStringShrink = &v
	return s
}

func (s *UpdateAccountAddressInfoShrinkRequest) SetPK(v string) *UpdateAccountAddressInfoShrinkRequest {
	s.PK = &v
	return s
}

func (s *UpdateAccountAddressInfoShrinkRequest) SetPostCode(v string) *UpdateAccountAddressInfoShrinkRequest {
	s.PostCode = &v
	return s
}

func (s *UpdateAccountAddressInfoShrinkRequest) SetProvinceJsonStringShrink(v string) *UpdateAccountAddressInfoShrinkRequest {
	s.ProvinceJsonStringShrink = &v
	return s
}

func (s *UpdateAccountAddressInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
