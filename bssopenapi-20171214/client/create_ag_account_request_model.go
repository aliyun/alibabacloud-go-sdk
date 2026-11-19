// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountAttr(v string) *CreateAgAccountRequest
	GetAccountAttr() *string
	SetCityName(v string) *CreateAgAccountRequest
	GetCityName() *string
	SetEnterpriseName(v string) *CreateAgAccountRequest
	GetEnterpriseName() *string
	SetFirstName(v string) *CreateAgAccountRequest
	GetFirstName() *string
	SetLastName(v string) *CreateAgAccountRequest
	GetLastName() *string
	SetLoginEmail(v string) *CreateAgAccountRequest
	GetLoginEmail() *string
	SetNationCode(v string) *CreateAgAccountRequest
	GetNationCode() *string
	SetPostcode(v string) *CreateAgAccountRequest
	GetPostcode() *string
	SetProvinceName(v string) *CreateAgAccountRequest
	GetProvinceName() *string
}

type CreateAgAccountRequest struct {
	// The attributes of the new account. To view the attributes, click Account Attributes in the top menu bar of the console to go to the **Security Settings*	- page.
	//
	// example:
	//
	// 1
	AccountAttr *string `json:"AccountAttr,omitempty" xml:"AccountAttr,omitempty"`
	// The name of the city in English.
	//
	// example:
	//
	// Beijing
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The name of the enterprise. The name must be the same as the enterprise name associated with the unified social credit code.
	//
	// example:
	//
	// Alibaba Cloud Test Co., Ltd
	EnterpriseName *string `json:"EnterpriseName,omitempty" xml:"EnterpriseName,omitempty"`
	// The name of the first-level category.
	//
	// example:
	//
	// Zhicheng
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// The last name of the user.
	//
	// Maximum length: 64 characters.
	//
	// example:
	//
	// Wu
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// The logon email address.
	//
	// This parameter is required.
	//
	// example:
	//
	// email
	LoginEmail *string `json:"LoginEmail,omitempty" xml:"LoginEmail,omitempty"`
	// The country code.
	//
	// example:
	//
	// CN
	NationCode *string `json:"NationCode,omitempty" xml:"NationCode,omitempty"`
	// The postal code.
	//
	// example:
	//
	// 350000
	Postcode *string `json:"Postcode,omitempty" xml:"Postcode,omitempty"`
	// The name of the province. This parameter is optional.
	//
	// example:
	//
	// Beijing
	ProvinceName *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
}

func (s CreateAgAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAgAccountRequest) GetAccountAttr() *string {
	return s.AccountAttr
}

func (s *CreateAgAccountRequest) GetCityName() *string {
	return s.CityName
}

func (s *CreateAgAccountRequest) GetEnterpriseName() *string {
	return s.EnterpriseName
}

func (s *CreateAgAccountRequest) GetFirstName() *string {
	return s.FirstName
}

func (s *CreateAgAccountRequest) GetLastName() *string {
	return s.LastName
}

func (s *CreateAgAccountRequest) GetLoginEmail() *string {
	return s.LoginEmail
}

func (s *CreateAgAccountRequest) GetNationCode() *string {
	return s.NationCode
}

func (s *CreateAgAccountRequest) GetPostcode() *string {
	return s.Postcode
}

func (s *CreateAgAccountRequest) GetProvinceName() *string {
	return s.ProvinceName
}

func (s *CreateAgAccountRequest) SetAccountAttr(v string) *CreateAgAccountRequest {
	s.AccountAttr = &v
	return s
}

func (s *CreateAgAccountRequest) SetCityName(v string) *CreateAgAccountRequest {
	s.CityName = &v
	return s
}

func (s *CreateAgAccountRequest) SetEnterpriseName(v string) *CreateAgAccountRequest {
	s.EnterpriseName = &v
	return s
}

func (s *CreateAgAccountRequest) SetFirstName(v string) *CreateAgAccountRequest {
	s.FirstName = &v
	return s
}

func (s *CreateAgAccountRequest) SetLastName(v string) *CreateAgAccountRequest {
	s.LastName = &v
	return s
}

func (s *CreateAgAccountRequest) SetLoginEmail(v string) *CreateAgAccountRequest {
	s.LoginEmail = &v
	return s
}

func (s *CreateAgAccountRequest) SetNationCode(v string) *CreateAgAccountRequest {
	s.NationCode = &v
	return s
}

func (s *CreateAgAccountRequest) SetPostcode(v string) *CreateAgAccountRequest {
	s.Postcode = &v
	return s
}

func (s *CreateAgAccountRequest) SetProvinceName(v string) *CreateAgAccountRequest {
	s.ProvinceName = &v
	return s
}

func (s *CreateAgAccountRequest) Validate() error {
	return dara.Validate(s)
}
