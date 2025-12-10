// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendVerificationCodeForBindSecureMobilePhoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *SendVerificationCodeForBindSecureMobilePhoneRequest
	GetAccountId() *string
	SetSecureMobilePhone(v string) *SendVerificationCodeForBindSecureMobilePhoneRequest
	GetSecureMobilePhone() *string
}

type SendVerificationCodeForBindSecureMobilePhoneRequest struct {
	// The ID of the resource account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 138660628348****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The mobile phone number that you want to bind to the resource account.
	//
	// Specify the mobile phone number in the \\<Country code>-\\<Mobile phone number> format.
	//
	// >  Mobile phone numbers in the `86-<Mobile phone number>` format in the Chinese mainland are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// xx-13900001234
	SecureMobilePhone *string `json:"SecureMobilePhone,omitempty" xml:"SecureMobilePhone,omitempty"`
}

func (s SendVerificationCodeForBindSecureMobilePhoneRequest) String() string {
	return dara.Prettify(s)
}

func (s SendVerificationCodeForBindSecureMobilePhoneRequest) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForBindSecureMobilePhoneRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *SendVerificationCodeForBindSecureMobilePhoneRequest) GetSecureMobilePhone() *string {
	return s.SecureMobilePhone
}

func (s *SendVerificationCodeForBindSecureMobilePhoneRequest) SetAccountId(v string) *SendVerificationCodeForBindSecureMobilePhoneRequest {
	s.AccountId = &v
	return s
}

func (s *SendVerificationCodeForBindSecureMobilePhoneRequest) SetSecureMobilePhone(v string) *SendVerificationCodeForBindSecureMobilePhoneRequest {
	s.SecureMobilePhone = &v
	return s
}

func (s *SendVerificationCodeForBindSecureMobilePhoneRequest) Validate() error {
	return dara.Validate(s)
}
