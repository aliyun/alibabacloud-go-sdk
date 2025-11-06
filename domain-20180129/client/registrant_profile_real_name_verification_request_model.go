// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegistrantProfileRealNameVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityCredential(v string) *RegistrantProfileRealNameVerificationRequest
	GetIdentityCredential() *string
	SetIdentityCredentialNo(v string) *RegistrantProfileRealNameVerificationRequest
	GetIdentityCredentialNo() *string
	SetIdentityCredentialType(v string) *RegistrantProfileRealNameVerificationRequest
	GetIdentityCredentialType() *string
	SetLang(v string) *RegistrantProfileRealNameVerificationRequest
	GetLang() *string
	SetRegistrantProfileID(v int64) *RegistrantProfileRealNameVerificationRequest
	GetRegistrantProfileID() *int64
	SetUserClientIp(v string) *RegistrantProfileRealNameVerificationRequest
	GetUserClientIp() *string
}

type RegistrantProfileRealNameVerificationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dGVzdA==
	IdentityCredential *string `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 43012512345678****
	IdentityCredentialNo *string `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SFZ
	IdentityCredentialType *string `json:"IdentityCredentialType,omitempty" xml:"IdentityCredentialType,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	RegistrantProfileID *int64 `json:"RegistrantProfileID,omitempty" xml:"RegistrantProfileID,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s RegistrantProfileRealNameVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s RegistrantProfileRealNameVerificationRequest) GoString() string {
	return s.String()
}

func (s *RegistrantProfileRealNameVerificationRequest) GetIdentityCredential() *string {
	return s.IdentityCredential
}

func (s *RegistrantProfileRealNameVerificationRequest) GetIdentityCredentialNo() *string {
	return s.IdentityCredentialNo
}

func (s *RegistrantProfileRealNameVerificationRequest) GetIdentityCredentialType() *string {
	return s.IdentityCredentialType
}

func (s *RegistrantProfileRealNameVerificationRequest) GetLang() *string {
	return s.Lang
}

func (s *RegistrantProfileRealNameVerificationRequest) GetRegistrantProfileID() *int64 {
	return s.RegistrantProfileID
}

func (s *RegistrantProfileRealNameVerificationRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *RegistrantProfileRealNameVerificationRequest) SetIdentityCredential(v string) *RegistrantProfileRealNameVerificationRequest {
	s.IdentityCredential = &v
	return s
}

func (s *RegistrantProfileRealNameVerificationRequest) SetIdentityCredentialNo(v string) *RegistrantProfileRealNameVerificationRequest {
	s.IdentityCredentialNo = &v
	return s
}

func (s *RegistrantProfileRealNameVerificationRequest) SetIdentityCredentialType(v string) *RegistrantProfileRealNameVerificationRequest {
	s.IdentityCredentialType = &v
	return s
}

func (s *RegistrantProfileRealNameVerificationRequest) SetLang(v string) *RegistrantProfileRealNameVerificationRequest {
	s.Lang = &v
	return s
}

func (s *RegistrantProfileRealNameVerificationRequest) SetRegistrantProfileID(v int64) *RegistrantProfileRealNameVerificationRequest {
	s.RegistrantProfileID = &v
	return s
}

func (s *RegistrantProfileRealNameVerificationRequest) SetUserClientIp(v string) *RegistrantProfileRealNameVerificationRequest {
	s.UserClientIp = &v
	return s
}

func (s *RegistrantProfileRealNameVerificationRequest) Validate() error {
	return dara.Validate(s)
}
