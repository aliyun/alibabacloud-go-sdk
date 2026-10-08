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
	// Specific address.
	//
	// example:
	//
	// chao yang qu
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// City.
	//
	// example:
	//
	// bei jing shi
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// Country code, such as **CN*	- or **US**.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// List of domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// alibabacloud.com
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	// Mailbox.
	//
	// example:
	//
	// test@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Base64-encoded image of the identity verification document. Image requirements:
	//
	// - Format must be **jpg*	- or **bmp**.
	//
	// - Original image size must be between **55 KB and 1 MB**.
	//
	// This parameter is required.
	//
	// example:
	//
	// h6UPhXz/ADP/2Q==
	IdentityCredential *string `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	// Certificate number used for identity verification, such as an ID card number or Unified Social Credit Code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5****************9
	IdentityCredentialNo *string `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	// Identity verification certificate type. Valid values:
	//
	// - **SFZ**: Identity card.
	//
	// - **HZ**: Passport.
	//
	// - **YYZZ**: Business license.
	//
	// - **ORG**: Organization code certificate.
	//
	// - **XYDM**: Unified Social Credit Code certificate.
	//
	// - **TXZ**: Mainland Travel Permits for Hong Kong and Macao Residents.
	//
	// If your certificate type is not listed above, see [Supported identity verification certificate types](https://help.aliyun.com/document_detail/72209.html) for valid values of other certificate types.
	//
	// > You must select the certificate type that matches the document you are submitting.
	//
	// This parameter is required.
	//
	// example:
	//
	// SFZ
	IdentityCredentialType *string `json:"IdentityCredentialType,omitempty" xml:"IdentityCredentialType,omitempty"`
	// Language of the error message returned by the API. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Postal code.
	//
	// example:
	//
	// 123456
	PostalCode *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	// Province.
	//
	// example:
	//
	// bei jing
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// Contact name.
	//
	// example:
	//
	// ce shi
	RegistrantName *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	// Registrant organization name.
	//
	// example:
	//
	// ce shi
	RegistrantOrganization *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	// Domain registrant type. Valid values:
	//
	// - **1**: Individual.
	//
	// - **2**: Organization.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RegistrantType *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	// Telephone country code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86
	TelArea *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	// Telephone extension number.
	//
	// example:
	//
	// 12345
	TelExt *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	// Telephone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	// Whether to add a transfer-out prohibition restriction. This indicates whether modifying the registrant imposes a 60-day restriction on domain name transfer-out. Default value: **false**, which means transfer-out is not restricted.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	TransferOutProhibited *bool `json:"TransferOutProhibited,omitempty" xml:"TransferOutProhibited,omitempty"`
	// User IP address.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// Chinese address.
	//
	// example:
	//
	// 朝阳区
	ZhAddress *string `json:"ZhAddress,omitempty" xml:"ZhAddress,omitempty"`
	// Chinese city name.
	//
	// example:
	//
	// 北京市
	ZhCity *string `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	// Chinese province name.
	//
	// example:
	//
	// 北京
	ZhProvince *string `json:"ZhProvince,omitempty" xml:"ZhProvince,omitempty"`
	// Chinese contact name.
	//
	// example:
	//
	// 测试
	ZhRegistrantName *string `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	// Chinese registrant organization name.
	//
	// example:
	//
	// 测试
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
