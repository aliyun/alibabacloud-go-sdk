// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveRegistrantProfileRealNameVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetAddress() *string
	SetCity(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetCity() *string
	SetCountry(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetCountry() *string
	SetEmail(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetEmail() *string
	SetIdentityCredential(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetIdentityCredential() *string
	SetIdentityCredentialNo(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetIdentityCredentialNo() *string
	SetIdentityCredentialType(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetIdentityCredentialType() *string
	SetLang(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetLang() *string
	SetPostalCode(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetPostalCode() *string
	SetProvince(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetProvince() *string
	SetRegistrantName(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetRegistrantName() *string
	SetRegistrantOrganization(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetRegistrantOrganization() *string
	SetRegistrantProfileId(v int64) *SaveRegistrantProfileRealNameVerificationRequest
	GetRegistrantProfileId() *int64
	SetRegistrantProfileType(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetRegistrantProfileType() *string
	SetRegistrantType(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetRegistrantType() *string
	SetTelArea(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetTelArea() *string
	SetTelExt(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetTelExt() *string
	SetTelephone(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetTelephone() *string
	SetUserClientIp(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetUserClientIp() *string
	SetZhAddress(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetZhAddress() *string
	SetZhCity(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetZhCity() *string
	SetZhProvince(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetZhProvince() *string
	SetZhRegistrantName(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetZhRegistrantName() *string
	SetZhRegistrantOrganization(v string) *SaveRegistrantProfileRealNameVerificationRequest
	GetZhRegistrantOrganization() *string
}

type SaveRegistrantProfileRealNameVerificationRequest struct {
	// example:
	//
	// chao yang qu
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// bei jing shi
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// dGVzdA==
	IdentityCredential *string `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	// example:
	//
	// 4111111111111110**
	IdentityCredentialNo *string `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	// example:
	//
	// SFZ
	IdentityCredentialType *string `json:"IdentityCredentialType,omitempty" xml:"IdentityCredentialType,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1234567
	PostalCode *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	// example:
	//
	// bei jing
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// example:
	//
	// ce shi
	RegistrantName *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	// example:
	//
	// ce shi
	RegistrantOrganization *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	// example:
	//
	// 1234567
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
	// 1234
	TelExt *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	// example:
	//
	// 12345678
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

func (s SaveRegistrantProfileRealNameVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveRegistrantProfileRealNameVerificationRequest) GoString() string {
	return s.String()
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetAddress() *string {
	return s.Address
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetCity() *string {
	return s.City
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetCountry() *string {
	return s.Country
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetEmail() *string {
	return s.Email
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetIdentityCredential() *string {
	return s.IdentityCredential
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetIdentityCredentialNo() *string {
	return s.IdentityCredentialNo
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetIdentityCredentialType() *string {
	return s.IdentityCredentialType
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetPostalCode() *string {
	return s.PostalCode
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetProvince() *string {
	return s.Province
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetRegistrantName() *string {
	return s.RegistrantName
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetRegistrantOrganization() *string {
	return s.RegistrantOrganization
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetRegistrantProfileType() *string {
	return s.RegistrantProfileType
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetRegistrantType() *string {
	return s.RegistrantType
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetTelArea() *string {
	return s.TelArea
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetTelExt() *string {
	return s.TelExt
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetTelephone() *string {
	return s.Telephone
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetZhAddress() *string {
	return s.ZhAddress
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetZhCity() *string {
	return s.ZhCity
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetZhProvince() *string {
	return s.ZhProvince
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetZhRegistrantName() *string {
	return s.ZhRegistrantName
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) GetZhRegistrantOrganization() *string {
	return s.ZhRegistrantOrganization
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetAddress(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.Address = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetCity(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.City = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetCountry(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.Country = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetEmail(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.Email = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetIdentityCredential(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.IdentityCredential = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetIdentityCredentialNo(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.IdentityCredentialNo = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetIdentityCredentialType(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.IdentityCredentialType = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetLang(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.Lang = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetPostalCode(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.PostalCode = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetProvince(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.Province = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetRegistrantName(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.RegistrantName = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetRegistrantOrganization(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.RegistrantOrganization = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetRegistrantProfileId(v int64) *SaveRegistrantProfileRealNameVerificationRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetRegistrantProfileType(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.RegistrantProfileType = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetRegistrantType(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.RegistrantType = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetTelArea(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.TelArea = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetTelExt(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.TelExt = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetTelephone(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.Telephone = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetUserClientIp(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetZhAddress(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.ZhAddress = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetZhCity(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.ZhCity = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetZhProvince(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.ZhProvince = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetZhRegistrantName(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.ZhRegistrantName = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetZhRegistrantOrganization(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.ZhRegistrantOrganization = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) Validate() error {
	return dara.Validate(s)
}
