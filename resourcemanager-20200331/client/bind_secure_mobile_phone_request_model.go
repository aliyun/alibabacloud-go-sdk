// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSecureMobilePhoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *BindSecureMobilePhoneRequest
	GetAccountId() *string
	SetSecureMobilePhone(v string) *BindSecureMobilePhoneRequest
	GetSecureMobilePhone() *string
	SetVerificationCode(v string) *BindSecureMobilePhoneRequest
	GetVerificationCode() *string
}

type BindSecureMobilePhoneRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 138660628348****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The mobile phone number that you want to bind to the member for security purposes.
	//
	// The mobile phone number you specify must be the same as the mobile phone number that you specify when you call the [SendVerificationCodeForBindSecureMobilePhone](https://help.aliyun.com/document_detail/372556.html) operation to obtain a verification code.
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
	// The verification code.
	//
	// You can call the [SendVerificationCodeForBindSecureMobilePhone](https://help.aliyun.com/document_detail/372556.html) operation to obtain the verification code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	VerificationCode *string `json:"VerificationCode,omitempty" xml:"VerificationCode,omitempty"`
}

func (s BindSecureMobilePhoneRequest) String() string {
	return dara.Prettify(s)
}

func (s BindSecureMobilePhoneRequest) GoString() string {
	return s.String()
}

func (s *BindSecureMobilePhoneRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *BindSecureMobilePhoneRequest) GetSecureMobilePhone() *string {
	return s.SecureMobilePhone
}

func (s *BindSecureMobilePhoneRequest) GetVerificationCode() *string {
	return s.VerificationCode
}

func (s *BindSecureMobilePhoneRequest) SetAccountId(v string) *BindSecureMobilePhoneRequest {
	s.AccountId = &v
	return s
}

func (s *BindSecureMobilePhoneRequest) SetSecureMobilePhone(v string) *BindSecureMobilePhoneRequest {
	s.SecureMobilePhone = &v
	return s
}

func (s *BindSecureMobilePhoneRequest) SetVerificationCode(v string) *BindSecureMobilePhoneRequest {
	s.VerificationCode = &v
	return s
}

func (s *BindSecureMobilePhoneRequest) Validate() error {
	return dara.Validate(s)
}
