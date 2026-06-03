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
}

type SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	City    *string `json:"City,omitempty" xml:"City,omitempty"`
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// This parameter is required.
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	Email      *string   `json:"Email,omitempty" xml:"Email,omitempty"`
	// This parameter is required.
	IdentityCredential *string `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	// This parameter is required.
	IdentityCredentialNo *string `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	// This parameter is required.
	IdentityCredentialType *string `json:"IdentityCredentialType,omitempty" xml:"IdentityCredentialType,omitempty"`
	Lang                   *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PostalCode             *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Province               *string `json:"Province,omitempty" xml:"Province,omitempty"`
	RegistrantName         *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	// This parameter is required.
	RegistrantType *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	// This parameter is required.
	TelArea *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt  *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	// This parameter is required.
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	// This parameter is required.
	TransferOutProhibited *bool   `json:"TransferOutProhibited,omitempty" xml:"TransferOutProhibited,omitempty"`
	UserClientIp          *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
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

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) Validate() error {
	return dara.Validate(s)
}
