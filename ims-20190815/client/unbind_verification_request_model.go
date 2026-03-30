// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *UnbindVerificationRequest
	GetEmail() *string
	SetMobilePhone(v string) *UnbindVerificationRequest
	GetMobilePhone() *string
	SetUserPrincipalName(v string) *UnbindVerificationRequest
	GetUserPrincipalName() *string
	SetVerifyType(v string) *UnbindVerificationRequest
	GetVerifyType() *string
}

type UnbindVerificationRequest struct {
	// The email address.
	//
	// >  If you set `VerifyType` to `email`, you must specify this parameter.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The mobile phone number.
	//
	// >  If you set `VerifyType` to `sms`, you must specify this parameter.
	//
	// example:
	//
	// 86-13900001234
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
	// The multi-factor authentication (MFA) method. Valid values:
	//
	// 	- sms: mobile phone.
	//
	// 	- email: email.
	//
	// example:
	//
	// sms
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s UnbindVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindVerificationRequest) GoString() string {
	return s.String()
}

func (s *UnbindVerificationRequest) GetEmail() *string {
	return s.Email
}

func (s *UnbindVerificationRequest) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *UnbindVerificationRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *UnbindVerificationRequest) GetVerifyType() *string {
	return s.VerifyType
}

func (s *UnbindVerificationRequest) SetEmail(v string) *UnbindVerificationRequest {
	s.Email = &v
	return s
}

func (s *UnbindVerificationRequest) SetMobilePhone(v string) *UnbindVerificationRequest {
	s.MobilePhone = &v
	return s
}

func (s *UnbindVerificationRequest) SetUserPrincipalName(v string) *UnbindVerificationRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *UnbindVerificationRequest) SetVerifyType(v string) *UnbindVerificationRequest {
	s.VerifyType = &v
	return s
}

func (s *UnbindVerificationRequest) Validate() error {
	return dara.Validate(s)
}
