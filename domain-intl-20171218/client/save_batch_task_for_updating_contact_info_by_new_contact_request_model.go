// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForUpdatingContactInfoByNewContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetAddress() *string
	SetCity(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetCity() *string
	SetContactType(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetContactType() *string
	SetCountry(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetCountry() *string
	SetDomainName(v []*string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetDomainName() []*string
	SetEmail(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetEmail() *string
	SetLang(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetLang() *string
	SetPostalCode(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetPostalCode() *string
	SetProvince(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetProvince() *string
	SetRegistrantName(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetRegistrantName() *string
	SetRegistrantOrganization(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetRegistrantOrganization() *string
	SetTelArea(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetTelArea() *string
	SetTelExt(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetTelExt() *string
	SetTelephone(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetTelephone() *string
	SetTransferOutProhibited(v bool) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetTransferOutProhibited() *bool
	SetUserClientIp(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetUserClientIp() *string
}

type SaveBatchTaskForUpdatingContactInfoByNewContactRequest struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	City    *string `json:"City,omitempty" xml:"City,omitempty"`
	// This parameter is required.
	ContactType *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// This parameter is required.
	DomainName             []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	Email                  *string   `json:"Email,omitempty" xml:"Email,omitempty"`
	Lang                   *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PostalCode             *string   `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Province               *string   `json:"Province,omitempty" xml:"Province,omitempty"`
	RegistrantName         *string   `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization *string   `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	TelArea                *string   `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt                 *string   `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	Telephone              *string   `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	TransferOutProhibited  *bool     `json:"TransferOutProhibited,omitempty" xml:"TransferOutProhibited,omitempty"`
	UserClientIp           *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForUpdatingContactInfoByNewContactRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetAddress() *string {
	return s.Address
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetCity() *string {
	return s.City
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetContactType() *string {
	return s.ContactType
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetCountry() *string {
	return s.Country
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetDomainName() []*string {
	return s.DomainName
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetEmail() *string {
	return s.Email
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetPostalCode() *string {
	return s.PostalCode
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetProvince() *string {
	return s.Province
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetRegistrantName() *string {
	return s.RegistrantName
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetRegistrantOrganization() *string {
	return s.RegistrantOrganization
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetTelArea() *string {
	return s.TelArea
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetTelExt() *string {
	return s.TelExt
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetTelephone() *string {
	return s.Telephone
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetTransferOutProhibited() *bool {
	return s.TransferOutProhibited
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetAddress(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.Address = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetCity(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.City = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetContactType(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.ContactType = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetCountry(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.Country = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetDomainName(v []*string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.DomainName = v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetEmail(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.Email = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetLang(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetPostalCode(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.PostalCode = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetProvince(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.Province = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetRegistrantName(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.RegistrantName = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetRegistrantOrganization(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.RegistrantOrganization = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetTelArea(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.TelArea = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetTelExt(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.TelExt = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetTelephone(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.Telephone = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetTransferOutProhibited(v bool) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.TransferOutProhibited = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetUserClientIp(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) Validate() error {
	return dara.Validate(s)
}
