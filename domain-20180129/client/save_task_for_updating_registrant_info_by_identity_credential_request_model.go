// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetAddress() *string
	SetCity(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetCity() *string
	SetCountry(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetCountry() *string
	SetDomainName(v []*string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetDomainName() []*string
	SetEmail(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetEmail() *string
	SetIdentityCredential(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetIdentityCredential() *string
	SetIdentityCredentialNo(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetIdentityCredentialNo() *string
	SetIdentityCredentialType(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetIdentityCredentialType() *string
	SetLang(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetLang() *string
	SetPostalCode(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetPostalCode() *string
	SetProvince(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetProvince() *string
	SetRegistrantName(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetRegistrantName() *string
	SetRegistrantOrganization(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetRegistrantOrganization() *string
	SetRegistrantType(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetRegistrantType() *string
	SetTelArea(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetTelArea() *string
	SetTelExt(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetTelExt() *string
	SetTelephone(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetTelephone() *string
	SetTransferOutProhibited(v bool) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetTransferOutProhibited() *bool
	SetUserClientIp(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetUserClientIp() *string
	SetZhAddress(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetZhAddress() *string
	SetZhCity(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetZhCity() *string
	SetZhProvince(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetZhProvince() *string
	SetZhRegistrantName(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetZhRegistrantName() *string
	SetZhRegistrantOrganization(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
	GetZhRegistrantOrganization() *string
}

type SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// alibabacloud.com
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	// example:
	//
	// test@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// h6UPhXz/ADP/2Q==
	IdentityCredential *string `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5****************9
	IdentityCredentialNo *string `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	// This parameter is required.
	//
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
	// 123456
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	RegistrantType *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 86
	TelArea *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	// example:
	//
	// 12345
	TelExt *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	TransferOutProhibited *bool `json:"TransferOutProhibited,omitempty" xml:"TransferOutProhibited,omitempty"`
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

func (s SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetAddress() *string {
	return s.Address
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetCity() *string {
	return s.City
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetCountry() *string {
	return s.Country
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetDomainName() []*string {
	return s.DomainName
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetEmail() *string {
	return s.Email
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetIdentityCredential() *string {
	return s.IdentityCredential
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetIdentityCredentialNo() *string {
	return s.IdentityCredentialNo
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetIdentityCredentialType() *string {
	return s.IdentityCredentialType
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetPostalCode() *string {
	return s.PostalCode
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetProvince() *string {
	return s.Province
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetRegistrantName() *string {
	return s.RegistrantName
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetRegistrantOrganization() *string {
	return s.RegistrantOrganization
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetRegistrantType() *string {
	return s.RegistrantType
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetTelArea() *string {
	return s.TelArea
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetTelExt() *string {
	return s.TelExt
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetTelephone() *string {
	return s.Telephone
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetTransferOutProhibited() *bool {
	return s.TransferOutProhibited
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetZhAddress() *string {
	return s.ZhAddress
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetZhCity() *string {
	return s.ZhCity
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetZhProvince() *string {
	return s.ZhProvince
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetZhRegistrantName() *string {
	return s.ZhRegistrantName
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GetZhRegistrantOrganization() *string {
	return s.ZhRegistrantOrganization
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetAddress(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.Address = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetCity(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.City = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetCountry(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.Country = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetDomainName(v []*string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.DomainName = v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetEmail(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.Email = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetIdentityCredential(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.IdentityCredential = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetIdentityCredentialNo(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.IdentityCredentialNo = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetIdentityCredentialType(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.IdentityCredentialType = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetLang(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.Lang = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetPostalCode(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.PostalCode = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetProvince(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.Province = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetRegistrantName(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.RegistrantName = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetRegistrantOrganization(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.RegistrantOrganization = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetRegistrantType(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.RegistrantType = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetTelArea(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.TelArea = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetTelExt(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.TelExt = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetTelephone(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.Telephone = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetTransferOutProhibited(v bool) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.TransferOutProhibited = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetUserClientIp(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetZhAddress(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.ZhAddress = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetZhCity(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.ZhCity = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetZhProvince(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.ZhProvince = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetZhRegistrantName(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.ZhRegistrantName = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetZhRegistrantOrganization(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.ZhRegistrantOrganization = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) Validate() error {
	return dara.Validate(s)
}
