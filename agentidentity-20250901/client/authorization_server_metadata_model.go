// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizationServerMetadata interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRequest(v *AuthorizationRequest) *AuthorizationServerMetadata
	GetAuthorizationRequest() *AuthorizationRequest
	SetIssuer(v string) *AuthorizationServerMetadata
	GetIssuer() *string
	SetPKCE(v *PKCE) *AuthorizationServerMetadata
	GetPKCE() *PKCE
	SetTokenRequest(v *TokenReqeust) *AuthorizationServerMetadata
	GetTokenRequest() *TokenReqeust
}

type AuthorizationServerMetadata struct {
	AuthorizationRequest *AuthorizationRequest `json:"AuthorizationRequest,omitempty" xml:"AuthorizationRequest,omitempty"`
	Issuer               *string               `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	PKCE                 *PKCE                 `json:"PKCE,omitempty" xml:"PKCE,omitempty"`
	TokenRequest         *TokenReqeust         `json:"TokenRequest,omitempty" xml:"TokenRequest,omitempty"`
}

func (s AuthorizationServerMetadata) String() string {
	return dara.Prettify(s)
}

func (s AuthorizationServerMetadata) GoString() string {
	return s.String()
}

func (s *AuthorizationServerMetadata) GetAuthorizationRequest() *AuthorizationRequest {
	return s.AuthorizationRequest
}

func (s *AuthorizationServerMetadata) GetIssuer() *string {
	return s.Issuer
}

func (s *AuthorizationServerMetadata) GetPKCE() *PKCE {
	return s.PKCE
}

func (s *AuthorizationServerMetadata) GetTokenRequest() *TokenReqeust {
	return s.TokenRequest
}

func (s *AuthorizationServerMetadata) SetAuthorizationRequest(v *AuthorizationRequest) *AuthorizationServerMetadata {
	s.AuthorizationRequest = v
	return s
}

func (s *AuthorizationServerMetadata) SetIssuer(v string) *AuthorizationServerMetadata {
	s.Issuer = &v
	return s
}

func (s *AuthorizationServerMetadata) SetPKCE(v *PKCE) *AuthorizationServerMetadata {
	s.PKCE = v
	return s
}

func (s *AuthorizationServerMetadata) SetTokenRequest(v *TokenReqeust) *AuthorizationServerMetadata {
	s.TokenRequest = v
	return s
}

func (s *AuthorizationServerMetadata) Validate() error {
	if s.AuthorizationRequest != nil {
		if err := s.AuthorizationRequest.Validate(); err != nil {
			return err
		}
	}
	if s.PKCE != nil {
		if err := s.PKCE.Validate(); err != nil {
			return err
		}
	}
	if s.TokenRequest != nil {
		if err := s.TokenRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}
