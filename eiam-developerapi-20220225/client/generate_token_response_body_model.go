// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GenerateTokenResponseBody
	GetAccessToken() *string
	SetExpiresAt(v int64) *GenerateTokenResponseBody
	GetExpiresAt() *int64
	SetExpiresIn(v int64) *GenerateTokenResponseBody
	GetExpiresIn() *int64
	SetIdToken(v string) *GenerateTokenResponseBody
	GetIdToken() *string
	SetRefreshToken(v string) *GenerateTokenResponseBody
	GetRefreshToken() *string
	SetTokenType(v string) *GenerateTokenResponseBody
	GetTokenType() *string
}

type GenerateTokenResponseBody struct {
	// The access token.
	//
	// example:
	//
	// ATxxx
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// The time when the token expires. This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1653288641
	ExpiresAt *int64 `json:"expires_at,omitempty" xml:"expires_at,omitempty"`
	// The remaining validity period of the token. Unit: seconds.
	//
	// example:
	//
	// 1200
	ExpiresIn *int64 `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	// The ID token.
	//
	// example:
	//
	// xxxxx
	IdToken *string `json:"id_token,omitempty" xml:"id_token,omitempty"`
	// The refresh token.
	//
	// example:
	//
	// RTxxx
	RefreshToken *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	// The type of the token. Valid values: Basic Bearer
	//
	// example:
	//
	// Bearer
	TokenType *string `json:"token_type,omitempty" xml:"token_type,omitempty"`
}

func (s GenerateTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTokenResponseBody) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GenerateTokenResponseBody) GetExpiresAt() *int64 {
	return s.ExpiresAt
}

func (s *GenerateTokenResponseBody) GetExpiresIn() *int64 {
	return s.ExpiresIn
}

func (s *GenerateTokenResponseBody) GetIdToken() *string {
	return s.IdToken
}

func (s *GenerateTokenResponseBody) GetRefreshToken() *string {
	return s.RefreshToken
}

func (s *GenerateTokenResponseBody) GetTokenType() *string {
	return s.TokenType
}

func (s *GenerateTokenResponseBody) SetAccessToken(v string) *GenerateTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GenerateTokenResponseBody) SetExpiresAt(v int64) *GenerateTokenResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *GenerateTokenResponseBody) SetExpiresIn(v int64) *GenerateTokenResponseBody {
	s.ExpiresIn = &v
	return s
}

func (s *GenerateTokenResponseBody) SetIdToken(v string) *GenerateTokenResponseBody {
	s.IdToken = &v
	return s
}

func (s *GenerateTokenResponseBody) SetRefreshToken(v string) *GenerateTokenResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *GenerateTokenResponseBody) SetTokenType(v string) *GenerateTokenResponseBody {
	s.TokenType = &v
	return s
}

func (s *GenerateTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
