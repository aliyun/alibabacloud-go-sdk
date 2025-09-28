// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSmsVerifyCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaseAuthPolicy(v int64) *CheckSmsVerifyCodeRequest
	GetCaseAuthPolicy() *int64
	SetCountryCode(v string) *CheckSmsVerifyCodeRequest
	GetCountryCode() *string
	SetOutId(v string) *CheckSmsVerifyCodeRequest
	GetOutId() *string
	SetOwnerId(v int64) *CheckSmsVerifyCodeRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *CheckSmsVerifyCodeRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *CheckSmsVerifyCodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckSmsVerifyCodeRequest
	GetResourceOwnerId() *int64
	SetSchemeName(v string) *CheckSmsVerifyCodeRequest
	GetSchemeName() *string
	SetVerifyCode(v string) *CheckSmsVerifyCodeRequest
	GetVerifyCode() *string
}

type CheckSmsVerifyCodeRequest struct {
	// The verification policy for uppercase and lowercase letters of the verification code. Valid values:
	//
	// 	- 1: The verification policy does not distinguish uppercase and lowercase letters.
	//
	// 	- 2: The verification policy distinguishes uppercase and lowercase letters.
	//
	// example:
	//
	// 1
	CaseAuthPolicy *int64 `json:"CaseAuthPolicy,omitempty" xml:"CaseAuthPolicy,omitempty"`
	// The country code of the phone number. Default value: 86.
	//
	// example:
	//
	// 86
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The external ID.
	//
	// example:
	//
	// 12123231
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 18653529399
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The verification service name. If this parameter is not specified, the default service is used. The name can be up to 20 characters in length.
	//
	// example:
	//
	// Aliyun
	SchemeName *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	// The verification code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1231
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s CheckSmsVerifyCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckSmsVerifyCodeRequest) GoString() string {
	return s.String()
}

func (s *CheckSmsVerifyCodeRequest) GetCaseAuthPolicy() *int64 {
	return s.CaseAuthPolicy
}

func (s *CheckSmsVerifyCodeRequest) GetCountryCode() *string {
	return s.CountryCode
}

func (s *CheckSmsVerifyCodeRequest) GetOutId() *string {
	return s.OutId
}

func (s *CheckSmsVerifyCodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckSmsVerifyCodeRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *CheckSmsVerifyCodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckSmsVerifyCodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckSmsVerifyCodeRequest) GetSchemeName() *string {
	return s.SchemeName
}

func (s *CheckSmsVerifyCodeRequest) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *CheckSmsVerifyCodeRequest) SetCaseAuthPolicy(v int64) *CheckSmsVerifyCodeRequest {
	s.CaseAuthPolicy = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetCountryCode(v string) *CheckSmsVerifyCodeRequest {
	s.CountryCode = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetOutId(v string) *CheckSmsVerifyCodeRequest {
	s.OutId = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetOwnerId(v int64) *CheckSmsVerifyCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetPhoneNumber(v string) *CheckSmsVerifyCodeRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetResourceOwnerAccount(v string) *CheckSmsVerifyCodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetResourceOwnerId(v int64) *CheckSmsVerifyCodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetSchemeName(v string) *CheckSmsVerifyCodeRequest {
	s.SchemeName = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetVerifyCode(v string) *CheckSmsVerifyCodeRequest {
	s.VerifyCode = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) Validate() error {
	return dara.Validate(s)
}
