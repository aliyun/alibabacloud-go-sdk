// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateJwtAuthenticationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudiences(v []*string) *GenerateJwtAuthenticationTokenRequest
	GetAudiences() []*string
	SetCredentialProviderIdentifier(v string) *GenerateJwtAuthenticationTokenRequest
	GetCredentialProviderIdentifier() *string
	SetCustomClaims(v map[string]interface{}) *GenerateJwtAuthenticationTokenRequest
	GetCustomClaims() map[string]interface{}
	SetExpiration(v int32) *GenerateJwtAuthenticationTokenRequest
	GetExpiration() *int32
	SetIncludeDerivedShortToken(v bool) *GenerateJwtAuthenticationTokenRequest
	GetIncludeDerivedShortToken() *bool
	SetIssuer(v string) *GenerateJwtAuthenticationTokenRequest
	GetIssuer() *string
	SetSubject(v string) *GenerateJwtAuthenticationTokenRequest
	GetSubject() *string
}

type GenerateJwtAuthenticationTokenRequest struct {
	// This parameter is required.
	Audiences []*string `json:"audiences,omitempty" xml:"audiences,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test_example_identifier
	CredentialProviderIdentifier *string                `json:"credentialProviderIdentifier,omitempty" xml:"credentialProviderIdentifier,omitempty"`
	CustomClaims                 map[string]interface{} `json:"customClaims,omitempty" xml:"customClaims,omitempty"`
	// example:
	//
	// 900
	Expiration *int32 `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// example:
	//
	// true
	IncludeDerivedShortToken *bool `json:"includeDerivedShortToken,omitempty" xml:"includeDerivedShortToken,omitempty"`
	// example:
	//
	// https://test.issuer.com
	Issuer *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_jwt_subject
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s GenerateJwtAuthenticationTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateJwtAuthenticationTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateJwtAuthenticationTokenRequest) GetAudiences() []*string {
	return s.Audiences
}

func (s *GenerateJwtAuthenticationTokenRequest) GetCredentialProviderIdentifier() *string {
	return s.CredentialProviderIdentifier
}

func (s *GenerateJwtAuthenticationTokenRequest) GetCustomClaims() map[string]interface{} {
	return s.CustomClaims
}

func (s *GenerateJwtAuthenticationTokenRequest) GetExpiration() *int32 {
	return s.Expiration
}

func (s *GenerateJwtAuthenticationTokenRequest) GetIncludeDerivedShortToken() *bool {
	return s.IncludeDerivedShortToken
}

func (s *GenerateJwtAuthenticationTokenRequest) GetIssuer() *string {
	return s.Issuer
}

func (s *GenerateJwtAuthenticationTokenRequest) GetSubject() *string {
	return s.Subject
}

func (s *GenerateJwtAuthenticationTokenRequest) SetAudiences(v []*string) *GenerateJwtAuthenticationTokenRequest {
	s.Audiences = v
	return s
}

func (s *GenerateJwtAuthenticationTokenRequest) SetCredentialProviderIdentifier(v string) *GenerateJwtAuthenticationTokenRequest {
	s.CredentialProviderIdentifier = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenRequest) SetCustomClaims(v map[string]interface{}) *GenerateJwtAuthenticationTokenRequest {
	s.CustomClaims = v
	return s
}

func (s *GenerateJwtAuthenticationTokenRequest) SetExpiration(v int32) *GenerateJwtAuthenticationTokenRequest {
	s.Expiration = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenRequest) SetIncludeDerivedShortToken(v bool) *GenerateJwtAuthenticationTokenRequest {
	s.IncludeDerivedShortToken = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenRequest) SetIssuer(v string) *GenerateJwtAuthenticationTokenRequest {
	s.Issuer = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenRequest) SetSubject(v string) *GenerateJwtAuthenticationTokenRequest {
	s.Subject = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenRequest) Validate() error {
	return dara.Validate(s)
}
