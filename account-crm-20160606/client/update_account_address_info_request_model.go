// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountAddressInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *UpdateAccountAddressInfoRequest
	GetAddress() *string
	SetAddress2(v string) *UpdateAccountAddressInfoRequest
	GetAddress2() *string
	SetCityJsonString(v map[string]interface{}) *UpdateAccountAddressInfoRequest
	GetCityJsonString() map[string]interface{}
	SetDistrictJsonString(v map[string]interface{}) *UpdateAccountAddressInfoRequest
	GetDistrictJsonString() map[string]interface{}
	SetPK(v string) *UpdateAccountAddressInfoRequest
	GetPK() *string
	SetPostCode(v string) *UpdateAccountAddressInfoRequest
	GetPostCode() *string
	SetProvinceJsonString(v map[string]interface{}) *UpdateAccountAddressInfoRequest
	GetProvinceJsonString() map[string]interface{}
}

type UpdateAccountAddressInfoRequest struct {
	Address            *string                `json:"Address,omitempty" xml:"Address,omitempty"`
	Address2           *string                `json:"Address2,omitempty" xml:"Address2,omitempty"`
	CityJsonString     map[string]interface{} `json:"CityJsonString,omitempty" xml:"CityJsonString,omitempty"`
	DistrictJsonString map[string]interface{} `json:"DistrictJsonString,omitempty" xml:"DistrictJsonString,omitempty"`
	PK                 *string                `json:"PK,omitempty" xml:"PK,omitempty"`
	PostCode           *string                `json:"PostCode,omitempty" xml:"PostCode,omitempty"`
	ProvinceJsonString map[string]interface{} `json:"ProvinceJsonString,omitempty" xml:"ProvinceJsonString,omitempty"`
}

func (s UpdateAccountAddressInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountAddressInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountAddressInfoRequest) GetAddress() *string {
	return s.Address
}

func (s *UpdateAccountAddressInfoRequest) GetAddress2() *string {
	return s.Address2
}

func (s *UpdateAccountAddressInfoRequest) GetCityJsonString() map[string]interface{} {
	return s.CityJsonString
}

func (s *UpdateAccountAddressInfoRequest) GetDistrictJsonString() map[string]interface{} {
	return s.DistrictJsonString
}

func (s *UpdateAccountAddressInfoRequest) GetPK() *string {
	return s.PK
}

func (s *UpdateAccountAddressInfoRequest) GetPostCode() *string {
	return s.PostCode
}

func (s *UpdateAccountAddressInfoRequest) GetProvinceJsonString() map[string]interface{} {
	return s.ProvinceJsonString
}

func (s *UpdateAccountAddressInfoRequest) SetAddress(v string) *UpdateAccountAddressInfoRequest {
	s.Address = &v
	return s
}

func (s *UpdateAccountAddressInfoRequest) SetAddress2(v string) *UpdateAccountAddressInfoRequest {
	s.Address2 = &v
	return s
}

func (s *UpdateAccountAddressInfoRequest) SetCityJsonString(v map[string]interface{}) *UpdateAccountAddressInfoRequest {
	s.CityJsonString = v
	return s
}

func (s *UpdateAccountAddressInfoRequest) SetDistrictJsonString(v map[string]interface{}) *UpdateAccountAddressInfoRequest {
	s.DistrictJsonString = v
	return s
}

func (s *UpdateAccountAddressInfoRequest) SetPK(v string) *UpdateAccountAddressInfoRequest {
	s.PK = &v
	return s
}

func (s *UpdateAccountAddressInfoRequest) SetPostCode(v string) *UpdateAccountAddressInfoRequest {
	s.PostCode = &v
	return s
}

func (s *UpdateAccountAddressInfoRequest) SetProvinceJsonString(v map[string]interface{}) *UpdateAccountAddressInfoRequest {
	s.ProvinceJsonString = v
	return s
}

func (s *UpdateAccountAddressInfoRequest) Validate() error {
	return dara.Validate(s)
}
