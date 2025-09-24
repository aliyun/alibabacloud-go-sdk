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
	// The attribute of the account. To view the attribute of the account, use the account to log on to the Alibaba Cloud Management Console, move the pointer over the profile picture in the upper-right corner, and then click **Security Settings**.
	//
	// example:
	//
	// 1
	AccountAttr *string `json:"AccountAttr,omitempty" xml:"AccountAttr,omitempty"`
	// The name of the city.
	//
	// example:
	//
	// Beijing
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The name of the enterprise.
	//
	// example:
	//
	// Dongguan ChuangNeng Electric Appliance Co., Ltd
	EnterpriseName *string `json:"EnterpriseName,omitempty" xml:"EnterpriseName,omitempty"`
	// The first name.
	//
	// example:
	//
	// Zhicheng
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// The last name.
	//
	// The last name can be up to 64 characters in length.
	//
	// example:
	//
	// Wu
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// The email address used to log on to the account.
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
	// The zip code.
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
