// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPreValidatePhoneIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPhoneNumber(v string) *GetPreValidatePhoneIdRequest
	GetPhoneNumber() *string
	SetVerifyCode(v string) *GetPreValidatePhoneIdRequest
	GetVerifyCode() *string
}

type GetPreValidatePhoneIdRequest struct {
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 861380000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The verification code provided when you purchased the pre-registered phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 208393
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s GetPreValidatePhoneIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPreValidatePhoneIdRequest) GoString() string {
	return s.String()
}

func (s *GetPreValidatePhoneIdRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetPreValidatePhoneIdRequest) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *GetPreValidatePhoneIdRequest) SetPhoneNumber(v string) *GetPreValidatePhoneIdRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetPreValidatePhoneIdRequest) SetVerifyCode(v string) *GetPreValidatePhoneIdRequest {
	s.VerifyCode = &v
	return s
}

func (s *GetPreValidatePhoneIdRequest) Validate() error {
	return dara.Validate(s)
}
