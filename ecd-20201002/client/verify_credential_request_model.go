// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *VerifyCredentialRequest
	GetClientId() *string
	SetCredential(v string) *VerifyCredentialRequest
	GetCredential() *string
	SetCredentialType(v string) *VerifyCredentialRequest
	GetCredentialType() *string
	SetEncryptedKey(v string) *VerifyCredentialRequest
	GetEncryptedKey() *string
	SetLoginToken(v string) *VerifyCredentialRequest
	GetLoginToken() *string
	SetOfficeSiteId(v string) *VerifyCredentialRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *VerifyCredentialRequest
	GetRegionId() *string
	SetSessionId(v string) *VerifyCredentialRequest
	GetSessionId() *string
}

type VerifyCredentialRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// d0b95762-0541-4b53-a0e4-7ed09f39****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The credential.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789cyG
	Credential *string `json:"Credential,omitempty" xml:"Credential,omitempty"`
	// The type of the logon credential that you want to clear.
	//
	// Valid values:
	//
	// 	- MfaPasscode: the multi-factor verification code.
	//
	// 	- FingerPrint: the fingerprint.
	//
	// 	- Password: the password.
	//
	// example:
	//
	// Password
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// The ID of the key that you want to encrypt.
	//
	// example:
	//
	// drjfs****
	EncryptedKey *string `json:"EncryptedKey,omitempty" xml:"EncryptedKey,omitempty"`
	// The logon token.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1f5772a1c60dbea9fd8e1648567079018086448d234b5bc8e30bec0ba6e80c41c767c4dd0db51e9e5c4e0f111431a****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai+dir-227468****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// eb17af2e-1dd6-4cc4-a3ee-3a14d0d7****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s VerifyCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyCredentialRequest) GoString() string {
	return s.String()
}

func (s *VerifyCredentialRequest) GetClientId() *string {
	return s.ClientId
}

func (s *VerifyCredentialRequest) GetCredential() *string {
	return s.Credential
}

func (s *VerifyCredentialRequest) GetCredentialType() *string {
	return s.CredentialType
}

func (s *VerifyCredentialRequest) GetEncryptedKey() *string {
	return s.EncryptedKey
}

func (s *VerifyCredentialRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *VerifyCredentialRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *VerifyCredentialRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *VerifyCredentialRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *VerifyCredentialRequest) SetClientId(v string) *VerifyCredentialRequest {
	s.ClientId = &v
	return s
}

func (s *VerifyCredentialRequest) SetCredential(v string) *VerifyCredentialRequest {
	s.Credential = &v
	return s
}

func (s *VerifyCredentialRequest) SetCredentialType(v string) *VerifyCredentialRequest {
	s.CredentialType = &v
	return s
}

func (s *VerifyCredentialRequest) SetEncryptedKey(v string) *VerifyCredentialRequest {
	s.EncryptedKey = &v
	return s
}

func (s *VerifyCredentialRequest) SetLoginToken(v string) *VerifyCredentialRequest {
	s.LoginToken = &v
	return s
}

func (s *VerifyCredentialRequest) SetOfficeSiteId(v string) *VerifyCredentialRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *VerifyCredentialRequest) SetRegionId(v string) *VerifyCredentialRequest {
	s.RegionId = &v
	return s
}

func (s *VerifyCredentialRequest) SetSessionId(v string) *VerifyCredentialRequest {
	s.SessionId = &v
	return s
}

func (s *VerifyCredentialRequest) Validate() error {
	return dara.Validate(s)
}
