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
	SetRegistrantType(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetRegistrantType() *string
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
	SetZhAddress(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetZhAddress() *string
	SetZhCity(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetZhCity() *string
	SetZhProvince(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetZhProvince() *string
	SetZhRegistrantName(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetZhRegistrantName() *string
	SetZhRegistrantOrganization(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest
	GetZhRegistrantOrganization() *string
}

type SaveBatchTaskForUpdatingContactInfoByNewContactRequest struct {
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
	// Contact type. Valid values:
	//
	// - **registrant**: Registrant.
	//
	// - **admin**: Administrator.
	//
	// - **billing**: Billing contact.
	//
	// - **tech**: Technical contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// registrant
	ContactType *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	// Country code, such as **CN*	- or **US**.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// Domain name list.
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
	// Language of error messages returned by the API. Valid values:
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
	// - **2**: Enterprise.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RegistrantType *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	// Telephone country code.
	//
	// example:
	//
	// 86
	TelArea *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	// Extension number.
	//
	// example:
	//
	// 1235
	TelExt *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	// Telephone number.
	//
	// example:
	//
	// 1234567890
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	// Whether to add a transfer-out prohibition restriction. This parameter only takes effect when **ContactType*	- is **registrant**, indicating whether the domain name is restricted from transfer-out for 60 days after the registrant is modified. The default value is **false**, which means transfer-out is not restricted.
	//
	// example:
	//
	// false
	TransferOutProhibited *bool `json:"TransferOutProhibited,omitempty" xml:"TransferOutProhibited,omitempty"`
	// User IP.
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
	// Chinese city.
	//
	// example:
	//
	// 北京市
	ZhCity *string `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	// Chinese province.
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

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetRegistrantType() *string {
	return s.RegistrantType
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

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetZhAddress() *string {
	return s.ZhAddress
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetZhCity() *string {
	return s.ZhCity
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetZhProvince() *string {
	return s.ZhProvince
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetZhRegistrantName() *string {
	return s.ZhRegistrantName
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GetZhRegistrantOrganization() *string {
	return s.ZhRegistrantOrganization
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

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetRegistrantType(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.RegistrantType = &v
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

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetZhAddress(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.ZhAddress = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetZhCity(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.ZhCity = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetZhProvince(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.ZhProvince = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetZhRegistrantName(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.ZhRegistrantName = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetZhRegistrantOrganization(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.ZhRegistrantOrganization = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) Validate() error {
	return dara.Validate(s)
}
