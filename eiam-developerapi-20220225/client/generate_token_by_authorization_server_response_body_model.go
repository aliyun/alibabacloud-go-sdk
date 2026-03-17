// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTokenByAuthorizationServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GenerateTokenByAuthorizationServerResponseBody
	GetAccessToken() *string
	SetExpiresAt(v int64) *GenerateTokenByAuthorizationServerResponseBody
	GetExpiresAt() *int64
	SetExpiresIn(v int64) *GenerateTokenByAuthorizationServerResponseBody
	GetExpiresIn() *int64
	SetIdToken(v string) *GenerateTokenByAuthorizationServerResponseBody
	GetIdToken() *string
	SetRefreshToken(v string) *GenerateTokenByAuthorizationServerResponseBody
	GetRefreshToken() *string
	SetScope(v string) *GenerateTokenByAuthorizationServerResponseBody
	GetScope() *string
	SetTokenType(v string) *GenerateTokenByAuthorizationServerResponseBody
	GetTokenType() *string
}

type GenerateTokenByAuthorizationServerResponseBody struct {
	// example:
	//
	// eyJraWQiOiJLRVlLZ0Iyxxxxx
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// example:
	//
	// 1653288641
	ExpiresAt *int64 `json:"expires_at,omitempty" xml:"expires_at,omitempty"`
	// example:
	//
	// 1200
	ExpiresIn *int64 `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	// example:
	//
	// eyJraWQiOiJLRVlLZ0Iyxxxxx
	IdToken *string `json:"id_token,omitempty" xml:"id_token,omitempty"`
	// example:
	//
	// ATxxxxx
	RefreshToken *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	// example:
	//
	// openid
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// example:
	//
	// Bearer
	TokenType *string `json:"token_type,omitempty" xml:"token_type,omitempty"`
}

func (s GenerateTokenByAuthorizationServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateTokenByAuthorizationServerResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTokenByAuthorizationServerResponseBody) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GenerateTokenByAuthorizationServerResponseBody) GetExpiresAt() *int64 {
	return s.ExpiresAt
}

func (s *GenerateTokenByAuthorizationServerResponseBody) GetExpiresIn() *int64 {
	return s.ExpiresIn
}

func (s *GenerateTokenByAuthorizationServerResponseBody) GetIdToken() *string {
	return s.IdToken
}

func (s *GenerateTokenByAuthorizationServerResponseBody) GetRefreshToken() *string {
	return s.RefreshToken
}

func (s *GenerateTokenByAuthorizationServerResponseBody) GetScope() *string {
	return s.Scope
}

func (s *GenerateTokenByAuthorizationServerResponseBody) GetTokenType() *string {
	return s.TokenType
}

func (s *GenerateTokenByAuthorizationServerResponseBody) SetAccessToken(v string) *GenerateTokenByAuthorizationServerResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerResponseBody) SetExpiresAt(v int64) *GenerateTokenByAuthorizationServerResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerResponseBody) SetExpiresIn(v int64) *GenerateTokenByAuthorizationServerResponseBody {
	s.ExpiresIn = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerResponseBody) SetIdToken(v string) *GenerateTokenByAuthorizationServerResponseBody {
	s.IdToken = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerResponseBody) SetRefreshToken(v string) *GenerateTokenByAuthorizationServerResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerResponseBody) SetScope(v string) *GenerateTokenByAuthorizationServerResponseBody {
	s.Scope = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerResponseBody) SetTokenType(v string) *GenerateTokenByAuthorizationServerResponseBody {
	s.TokenType = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerResponseBody) Validate() error {
	return dara.Validate(s)
}
