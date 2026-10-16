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
	// The secure phone number of the new management account. If you leave this parameter empty, the secure phone number of the current logon account is used.
	//
	// Format: `<Country code>-<Phone number>`.
	//
	// <props="intl">
	//
	// > Phone numbers in the `86-<Phone number>` format from the Chinese mainland are not supported.
	//
	// <props="china">
	//
	// > Only phone numbers in the `86-<Phone number>` format from the Chinese mainland are supported.
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
