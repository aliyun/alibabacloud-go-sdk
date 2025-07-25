// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressVerifyIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressType(v string) *AddressVerifyIntlRequest
	GetAddressType() *string
	SetDefaultCity(v string) *AddressVerifyIntlRequest
	GetDefaultCity() *string
	SetDefaultCountry(v string) *AddressVerifyIntlRequest
	GetDefaultCountry() *string
	SetDefaultDistrict(v string) *AddressVerifyIntlRequest
	GetDefaultDistrict() *string
	SetDefaultProvince(v string) *AddressVerifyIntlRequest
	GetDefaultProvince() *string
	SetLatitude(v string) *AddressVerifyIntlRequest
	GetLatitude() *string
	SetLongitude(v string) *AddressVerifyIntlRequest
	GetLongitude() *string
	SetMobile(v string) *AddressVerifyIntlRequest
	GetMobile() *string
	SetProductCode(v string) *AddressVerifyIntlRequest
	GetProductCode() *string
	SetText(v string) *AddressVerifyIntlRequest
	GetText() *string
	SetVerifyType(v string) *AddressVerifyIntlRequest
	GetVerifyType() *string
}

type AddressVerifyIntlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// “0”
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	DefaultCity *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	// This parameter is required.
	DefaultCountry  *string `json:"DefaultCountry,omitempty" xml:"DefaultCountry,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	// example:
	//
	// “31.2304”
	Latitude *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// example:
	//
	// “121.4737”
	Longitude *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1872334****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ADD_VERIFY_PRO
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOME
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s AddressVerifyIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s AddressVerifyIntlRequest) GoString() string {
	return s.String()
}

func (s *AddressVerifyIntlRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *AddressVerifyIntlRequest) GetDefaultCity() *string {
	return s.DefaultCity
}

func (s *AddressVerifyIntlRequest) GetDefaultCountry() *string {
	return s.DefaultCountry
}

func (s *AddressVerifyIntlRequest) GetDefaultDistrict() *string {
	return s.DefaultDistrict
}

func (s *AddressVerifyIntlRequest) GetDefaultProvince() *string {
	return s.DefaultProvince
}

func (s *AddressVerifyIntlRequest) GetLatitude() *string {
	return s.Latitude
}

func (s *AddressVerifyIntlRequest) GetLongitude() *string {
	return s.Longitude
}

func (s *AddressVerifyIntlRequest) GetMobile() *string {
	return s.Mobile
}

func (s *AddressVerifyIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *AddressVerifyIntlRequest) GetText() *string {
	return s.Text
}

func (s *AddressVerifyIntlRequest) GetVerifyType() *string {
	return s.VerifyType
}

func (s *AddressVerifyIntlRequest) SetAddressType(v string) *AddressVerifyIntlRequest {
	s.AddressType = &v
	return s
}

func (s *AddressVerifyIntlRequest) SetDefaultCity(v string) *AddressVerifyIntlRequest {
	s.DefaultCity = &v
	return s
}

func (s *AddressVerifyIntlRequest) SetDefaultCountry(v string) *AddressVerifyIntlRequest {
	s.DefaultCountry = &v
	return s
}

func (s *AddressVerifyIntlRequest) SetDefaultDistrict(v string) *AddressVerifyIntlRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *AddressVerifyIntlRequest) SetDefaultProvince(v string) *AddressVerifyIntlRequest {
	s.DefaultProvince = &v
	return s
}

func (s *AddressVerifyIntlRequest) SetLatitude(v string) *AddressVerifyIntlRequest {
	s.Latitude = &v
	return s
}

func (s *AddressVerifyIntlRequest) SetLongitude(v string) *AddressVerifyIntlRequest {
	s.Longitude = &v
	return s
}

func (s *AddressVerifyIntlRequest) SetMobile(v string) *AddressVerifyIntlRequest {
	s.Mobile = &v
	return s
}

func (s *AddressVerifyIntlRequest) SetProductCode(v string) *AddressVerifyIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *AddressVerifyIntlRequest) SetText(v string) *AddressVerifyIntlRequest {
	s.Text = &v
	return s
}

func (s *AddressVerifyIntlRequest) SetVerifyType(v string) *AddressVerifyIntlRequest {
	s.VerifyType = &v
	return s
}

func (s *AddressVerifyIntlRequest) Validate() error {
	return dara.Validate(s)
}
