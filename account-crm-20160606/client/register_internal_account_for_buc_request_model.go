// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterInternalAccountForBucRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBid(v string) *RegisterInternalAccountForBucRequest
	GetBid() *string
	SetEmail(v string) *RegisterInternalAccountForBucRequest
	GetEmail() *string
	SetIsEmailConfirmed(v bool) *RegisterInternalAccountForBucRequest
	GetIsEmailConfirmed() *bool
	SetIsMobileConfirmed(v bool) *RegisterInternalAccountForBucRequest
	GetIsMobileConfirmed() *bool
	SetIsMobileLogin(v bool) *RegisterInternalAccountForBucRequest
	GetIsMobileLogin() *bool
	SetMobile(v string) *RegisterInternalAccountForBucRequest
	GetMobile() *string
	SetNationalityCode(v string) *RegisterInternalAccountForBucRequest
	GetNationalityCode() *string
	SetPlainPassword(v string) *RegisterInternalAccountForBucRequest
	GetPlainPassword() *string
	SetPreferredLanguage(v string) *RegisterInternalAccountForBucRequest
	GetPreferredLanguage() *string
	SetAccountTypeCode(v string) *RegisterInternalAccountForBucRequest
	GetAccountTypeCode() *string
}

type RegisterInternalAccountForBucRequest struct {
	// This parameter is required.
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// This parameter is required.
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty"`
	IsEmailConfirmed  *bool   `json:"IsEmailConfirmed,omitempty" xml:"IsEmailConfirmed,omitempty"`
	IsMobileConfirmed *bool   `json:"IsMobileConfirmed,omitempty" xml:"IsMobileConfirmed,omitempty"`
	IsMobileLogin     *bool   `json:"IsMobileLogin,omitempty" xml:"IsMobileLogin,omitempty"`
	Mobile            *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	NationalityCode   *string `json:"NationalityCode,omitempty" xml:"NationalityCode,omitempty"`
	PlainPassword     *string `json:"PlainPassword,omitempty" xml:"PlainPassword,omitempty"`
	PreferredLanguage *string `json:"PreferredLanguage,omitempty" xml:"PreferredLanguage,omitempty"`
	AccountTypeCode   *string `json:"accountTypeCode,omitempty" xml:"accountTypeCode,omitempty"`
}

func (s RegisterInternalAccountForBucRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterInternalAccountForBucRequest) GoString() string {
	return s.String()
}

func (s *RegisterInternalAccountForBucRequest) GetBid() *string {
	return s.Bid
}

func (s *RegisterInternalAccountForBucRequest) GetEmail() *string {
	return s.Email
}

func (s *RegisterInternalAccountForBucRequest) GetIsEmailConfirmed() *bool {
	return s.IsEmailConfirmed
}

func (s *RegisterInternalAccountForBucRequest) GetIsMobileConfirmed() *bool {
	return s.IsMobileConfirmed
}

func (s *RegisterInternalAccountForBucRequest) GetIsMobileLogin() *bool {
	return s.IsMobileLogin
}

func (s *RegisterInternalAccountForBucRequest) GetMobile() *string {
	return s.Mobile
}

func (s *RegisterInternalAccountForBucRequest) GetNationalityCode() *string {
	return s.NationalityCode
}

func (s *RegisterInternalAccountForBucRequest) GetPlainPassword() *string {
	return s.PlainPassword
}

func (s *RegisterInternalAccountForBucRequest) GetPreferredLanguage() *string {
	return s.PreferredLanguage
}

func (s *RegisterInternalAccountForBucRequest) GetAccountTypeCode() *string {
	return s.AccountTypeCode
}

func (s *RegisterInternalAccountForBucRequest) SetBid(v string) *RegisterInternalAccountForBucRequest {
	s.Bid = &v
	return s
}

func (s *RegisterInternalAccountForBucRequest) SetEmail(v string) *RegisterInternalAccountForBucRequest {
	s.Email = &v
	return s
}

func (s *RegisterInternalAccountForBucRequest) SetIsEmailConfirmed(v bool) *RegisterInternalAccountForBucRequest {
	s.IsEmailConfirmed = &v
	return s
}

func (s *RegisterInternalAccountForBucRequest) SetIsMobileConfirmed(v bool) *RegisterInternalAccountForBucRequest {
	s.IsMobileConfirmed = &v
	return s
}

func (s *RegisterInternalAccountForBucRequest) SetIsMobileLogin(v bool) *RegisterInternalAccountForBucRequest {
	s.IsMobileLogin = &v
	return s
}

func (s *RegisterInternalAccountForBucRequest) SetMobile(v string) *RegisterInternalAccountForBucRequest {
	s.Mobile = &v
	return s
}

func (s *RegisterInternalAccountForBucRequest) SetNationalityCode(v string) *RegisterInternalAccountForBucRequest {
	s.NationalityCode = &v
	return s
}

func (s *RegisterInternalAccountForBucRequest) SetPlainPassword(v string) *RegisterInternalAccountForBucRequest {
	s.PlainPassword = &v
	return s
}

func (s *RegisterInternalAccountForBucRequest) SetPreferredLanguage(v string) *RegisterInternalAccountForBucRequest {
	s.PreferredLanguage = &v
	return s
}

func (s *RegisterInternalAccountForBucRequest) SetAccountTypeCode(v string) *RegisterInternalAccountForBucRequest {
	s.AccountTypeCode = &v
	return s
}

func (s *RegisterInternalAccountForBucRequest) Validate() error {
	return dara.Validate(s)
}
