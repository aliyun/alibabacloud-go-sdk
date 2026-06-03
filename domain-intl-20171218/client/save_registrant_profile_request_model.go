// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveRegistrantProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *SaveRegistrantProfileRequest
	GetAddress() *string
	SetCity(v string) *SaveRegistrantProfileRequest
	GetCity() *string
	SetCountry(v string) *SaveRegistrantProfileRequest
	GetCountry() *string
	SetDefaultRegistrantProfile(v bool) *SaveRegistrantProfileRequest
	GetDefaultRegistrantProfile() *bool
	SetEmail(v string) *SaveRegistrantProfileRequest
	GetEmail() *string
	SetLang(v string) *SaveRegistrantProfileRequest
	GetLang() *string
	SetPostalCode(v string) *SaveRegistrantProfileRequest
	GetPostalCode() *string
	SetProvince(v string) *SaveRegistrantProfileRequest
	GetProvince() *string
	SetRegistrantName(v string) *SaveRegistrantProfileRequest
	GetRegistrantName() *string
	SetRegistrantOrganization(v string) *SaveRegistrantProfileRequest
	GetRegistrantOrganization() *string
	SetRegistrantProfileId(v int64) *SaveRegistrantProfileRequest
	GetRegistrantProfileId() *int64
	SetRegistrantProfileType(v string) *SaveRegistrantProfileRequest
	GetRegistrantProfileType() *string
	SetRegistrantType(v string) *SaveRegistrantProfileRequest
	GetRegistrantType() *string
	SetTelArea(v string) *SaveRegistrantProfileRequest
	GetTelArea() *string
	SetTelExt(v string) *SaveRegistrantProfileRequest
	GetTelExt() *string
	SetTelephone(v string) *SaveRegistrantProfileRequest
	GetTelephone() *string
	SetUserClientIp(v string) *SaveRegistrantProfileRequest
	GetUserClientIp() *string
}

type SaveRegistrantProfileRequest struct {
	// example:
	//
	// *****************************************************
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// long yan shi
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// false
	DefaultRegistrantProfile *bool `json:"DefaultRegistrantProfile,omitempty" xml:"DefaultRegistrantProfile,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 236300
	PostalCode *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	// example:
	//
	// fu jian
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// example:
	//
	// chen zi chen
	RegistrantName *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	// example:
	//
	// liu yang
	RegistrantOrganization *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	// example:
	//
	// 12659727
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// example:
	//
	// common
	RegistrantProfileType *string `json:"RegistrantProfileType,omitempty" xml:"RegistrantProfileType,omitempty"`
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
	// 7381
	TelExt *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	// example:
	//
	// 1829756****
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveRegistrantProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveRegistrantProfileRequest) GoString() string {
	return s.String()
}

func (s *SaveRegistrantProfileRequest) GetAddress() *string {
	return s.Address
}

func (s *SaveRegistrantProfileRequest) GetCity() *string {
	return s.City
}

func (s *SaveRegistrantProfileRequest) GetCountry() *string {
	return s.Country
}

func (s *SaveRegistrantProfileRequest) GetDefaultRegistrantProfile() *bool {
	return s.DefaultRegistrantProfile
}

func (s *SaveRegistrantProfileRequest) GetEmail() *string {
	return s.Email
}

func (s *SaveRegistrantProfileRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveRegistrantProfileRequest) GetPostalCode() *string {
	return s.PostalCode
}

func (s *SaveRegistrantProfileRequest) GetProvince() *string {
	return s.Province
}

func (s *SaveRegistrantProfileRequest) GetRegistrantName() *string {
	return s.RegistrantName
}

func (s *SaveRegistrantProfileRequest) GetRegistrantOrganization() *string {
	return s.RegistrantOrganization
}

func (s *SaveRegistrantProfileRequest) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *SaveRegistrantProfileRequest) GetRegistrantProfileType() *string {
	return s.RegistrantProfileType
}

func (s *SaveRegistrantProfileRequest) GetRegistrantType() *string {
	return s.RegistrantType
}

func (s *SaveRegistrantProfileRequest) GetTelArea() *string {
	return s.TelArea
}

func (s *SaveRegistrantProfileRequest) GetTelExt() *string {
	return s.TelExt
}

func (s *SaveRegistrantProfileRequest) GetTelephone() *string {
	return s.Telephone
}

func (s *SaveRegistrantProfileRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveRegistrantProfileRequest) SetAddress(v string) *SaveRegistrantProfileRequest {
	s.Address = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetCity(v string) *SaveRegistrantProfileRequest {
	s.City = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetCountry(v string) *SaveRegistrantProfileRequest {
	s.Country = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetDefaultRegistrantProfile(v bool) *SaveRegistrantProfileRequest {
	s.DefaultRegistrantProfile = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetEmail(v string) *SaveRegistrantProfileRequest {
	s.Email = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetLang(v string) *SaveRegistrantProfileRequest {
	s.Lang = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetPostalCode(v string) *SaveRegistrantProfileRequest {
	s.PostalCode = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetProvince(v string) *SaveRegistrantProfileRequest {
	s.Province = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetRegistrantName(v string) *SaveRegistrantProfileRequest {
	s.RegistrantName = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetRegistrantOrganization(v string) *SaveRegistrantProfileRequest {
	s.RegistrantOrganization = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetRegistrantProfileId(v int64) *SaveRegistrantProfileRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetRegistrantProfileType(v string) *SaveRegistrantProfileRequest {
	s.RegistrantProfileType = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetRegistrantType(v string) *SaveRegistrantProfileRequest {
	s.RegistrantType = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetTelArea(v string) *SaveRegistrantProfileRequest {
	s.TelArea = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetTelExt(v string) *SaveRegistrantProfileRequest {
	s.TelExt = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetTelephone(v string) *SaveRegistrantProfileRequest {
	s.Telephone = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetUserClientIp(v string) *SaveRegistrantProfileRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveRegistrantProfileRequest) Validate() error {
	return dara.Validate(s)
}
