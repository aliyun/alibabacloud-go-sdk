// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyContactFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *VerifyContactFieldRequest
	GetAddress() *string
	SetCity(v string) *VerifyContactFieldRequest
	GetCity() *string
	SetCountry(v string) *VerifyContactFieldRequest
	GetCountry() *string
	SetDomainName(v string) *VerifyContactFieldRequest
	GetDomainName() *string
	SetEmail(v string) *VerifyContactFieldRequest
	GetEmail() *string
	SetLang(v string) *VerifyContactFieldRequest
	GetLang() *string
	SetPostalCode(v string) *VerifyContactFieldRequest
	GetPostalCode() *string
	SetProvince(v string) *VerifyContactFieldRequest
	GetProvince() *string
	SetRegistrantName(v string) *VerifyContactFieldRequest
	GetRegistrantName() *string
	SetRegistrantOrganization(v string) *VerifyContactFieldRequest
	GetRegistrantOrganization() *string
	SetRegistrantType(v string) *VerifyContactFieldRequest
	GetRegistrantType() *string
	SetTelArea(v string) *VerifyContactFieldRequest
	GetTelArea() *string
	SetTelExt(v string) *VerifyContactFieldRequest
	GetTelExt() *string
	SetTelephone(v string) *VerifyContactFieldRequest
	GetTelephone() *string
	SetUserClientIp(v string) *VerifyContactFieldRequest
	GetUserClientIp() *string
	SetZhAddress(v string) *VerifyContactFieldRequest
	GetZhAddress() *string
	SetZhCity(v string) *VerifyContactFieldRequest
	GetZhCity() *string
	SetZhProvince(v string) *VerifyContactFieldRequest
	GetZhProvince() *string
	SetZhRegistrantName(v string) *VerifyContactFieldRequest
	GetZhRegistrantName() *string
	SetZhRegistrantOrganization(v string) *VerifyContactFieldRequest
	GetZhRegistrantOrganization() *string
}

type VerifyContactFieldRequest struct {
	// example:
	//
	// Rd. xitucheng
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// Bei jing
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// en
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PostalCode *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	// example:
	//
	// Bei jing
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// example:
	//
	// wang xian sheng
	RegistrantName *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	// example:
	//
	// wang xian sheng
	RegistrantOrganization *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	// example:
	//
	// 1
	RegistrantType *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	// example:
	//
	// 86
	TelArea *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	// example:
	//
	// 01
	TelExt *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	// example:
	//
	// 1390000****
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp             *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ZhAddress                *string `json:"ZhAddress,omitempty" xml:"ZhAddress,omitempty"`
	ZhCity                   *string `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	ZhProvince               *string `json:"ZhProvince,omitempty" xml:"ZhProvince,omitempty"`
	ZhRegistrantName         *string `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s VerifyContactFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyContactFieldRequest) GoString() string {
	return s.String()
}

func (s *VerifyContactFieldRequest) GetAddress() *string {
	return s.Address
}

func (s *VerifyContactFieldRequest) GetCity() *string {
	return s.City
}

func (s *VerifyContactFieldRequest) GetCountry() *string {
	return s.Country
}

func (s *VerifyContactFieldRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *VerifyContactFieldRequest) GetEmail() *string {
	return s.Email
}

func (s *VerifyContactFieldRequest) GetLang() *string {
	return s.Lang
}

func (s *VerifyContactFieldRequest) GetPostalCode() *string {
	return s.PostalCode
}

func (s *VerifyContactFieldRequest) GetProvince() *string {
	return s.Province
}

func (s *VerifyContactFieldRequest) GetRegistrantName() *string {
	return s.RegistrantName
}

func (s *VerifyContactFieldRequest) GetRegistrantOrganization() *string {
	return s.RegistrantOrganization
}

func (s *VerifyContactFieldRequest) GetRegistrantType() *string {
	return s.RegistrantType
}

func (s *VerifyContactFieldRequest) GetTelArea() *string {
	return s.TelArea
}

func (s *VerifyContactFieldRequest) GetTelExt() *string {
	return s.TelExt
}

func (s *VerifyContactFieldRequest) GetTelephone() *string {
	return s.Telephone
}

func (s *VerifyContactFieldRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *VerifyContactFieldRequest) GetZhAddress() *string {
	return s.ZhAddress
}

func (s *VerifyContactFieldRequest) GetZhCity() *string {
	return s.ZhCity
}

func (s *VerifyContactFieldRequest) GetZhProvince() *string {
	return s.ZhProvince
}

func (s *VerifyContactFieldRequest) GetZhRegistrantName() *string {
	return s.ZhRegistrantName
}

func (s *VerifyContactFieldRequest) GetZhRegistrantOrganization() *string {
	return s.ZhRegistrantOrganization
}

func (s *VerifyContactFieldRequest) SetAddress(v string) *VerifyContactFieldRequest {
	s.Address = &v
	return s
}

func (s *VerifyContactFieldRequest) SetCity(v string) *VerifyContactFieldRequest {
	s.City = &v
	return s
}

func (s *VerifyContactFieldRequest) SetCountry(v string) *VerifyContactFieldRequest {
	s.Country = &v
	return s
}

func (s *VerifyContactFieldRequest) SetDomainName(v string) *VerifyContactFieldRequest {
	s.DomainName = &v
	return s
}

func (s *VerifyContactFieldRequest) SetEmail(v string) *VerifyContactFieldRequest {
	s.Email = &v
	return s
}

func (s *VerifyContactFieldRequest) SetLang(v string) *VerifyContactFieldRequest {
	s.Lang = &v
	return s
}

func (s *VerifyContactFieldRequest) SetPostalCode(v string) *VerifyContactFieldRequest {
	s.PostalCode = &v
	return s
}

func (s *VerifyContactFieldRequest) SetProvince(v string) *VerifyContactFieldRequest {
	s.Province = &v
	return s
}

func (s *VerifyContactFieldRequest) SetRegistrantName(v string) *VerifyContactFieldRequest {
	s.RegistrantName = &v
	return s
}

func (s *VerifyContactFieldRequest) SetRegistrantOrganization(v string) *VerifyContactFieldRequest {
	s.RegistrantOrganization = &v
	return s
}

func (s *VerifyContactFieldRequest) SetRegistrantType(v string) *VerifyContactFieldRequest {
	s.RegistrantType = &v
	return s
}

func (s *VerifyContactFieldRequest) SetTelArea(v string) *VerifyContactFieldRequest {
	s.TelArea = &v
	return s
}

func (s *VerifyContactFieldRequest) SetTelExt(v string) *VerifyContactFieldRequest {
	s.TelExt = &v
	return s
}

func (s *VerifyContactFieldRequest) SetTelephone(v string) *VerifyContactFieldRequest {
	s.Telephone = &v
	return s
}

func (s *VerifyContactFieldRequest) SetUserClientIp(v string) *VerifyContactFieldRequest {
	s.UserClientIp = &v
	return s
}

func (s *VerifyContactFieldRequest) SetZhAddress(v string) *VerifyContactFieldRequest {
	s.ZhAddress = &v
	return s
}

func (s *VerifyContactFieldRequest) SetZhCity(v string) *VerifyContactFieldRequest {
	s.ZhCity = &v
	return s
}

func (s *VerifyContactFieldRequest) SetZhProvince(v string) *VerifyContactFieldRequest {
	s.ZhProvince = &v
	return s
}

func (s *VerifyContactFieldRequest) SetZhRegistrantName(v string) *VerifyContactFieldRequest {
	s.ZhRegistrantName = &v
	return s
}

func (s *VerifyContactFieldRequest) SetZhRegistrantOrganization(v string) *VerifyContactFieldRequest {
	s.ZhRegistrantOrganization = &v
	return s
}

func (s *VerifyContactFieldRequest) Validate() error {
	return dara.Validate(s)
}
