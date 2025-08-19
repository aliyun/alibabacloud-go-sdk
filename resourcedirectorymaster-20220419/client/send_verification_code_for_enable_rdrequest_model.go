// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendVerificationCodeForEnableRDRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecureMobilePhone(v string) *SendVerificationCodeForEnableRDRequest
	GetSecureMobilePhone() *string
}

type SendVerificationCodeForEnableRDRequest struct {
	// The mobile phone number that is bound to the newly created account. If you leave this parameter empty, the mobile phone number that is bound to the current account is used.
	//
	// Specify the mobile phone number in the `<Country code>-<Mobile phone number>` format.
	//
	// > Mobile phone numbers in the `86-<Mobile phone number>` format in the Chinese mainland are not supported.
	//
	// example:
	//
	// xx-13900001234
	SecureMobilePhone *string `json:"SecureMobilePhone,omitempty" xml:"SecureMobilePhone,omitempty"`
}

func (s SendVerificationCodeForEnableRDRequest) String() string {
	return dara.Prettify(s)
}

func (s SendVerificationCodeForEnableRDRequest) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForEnableRDRequest) GetSecureMobilePhone() *string {
	return s.SecureMobilePhone
}

func (s *SendVerificationCodeForEnableRDRequest) SetSecureMobilePhone(v string) *SendVerificationCodeForEnableRDRequest {
	s.SecureMobilePhone = &v
	return s
}

func (s *SendVerificationCodeForEnableRDRequest) Validate() error {
	return dara.Validate(s)
}
