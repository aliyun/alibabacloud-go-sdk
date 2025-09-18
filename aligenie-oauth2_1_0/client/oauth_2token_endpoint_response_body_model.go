// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOAuth2TokenEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *OAuth2TokenEndpointResponseBody
	GetAccessToken() *string
	SetExpiresIn(v int64) *OAuth2TokenEndpointResponseBody
	GetExpiresIn() *int64
	SetRefreshToken(v string) *OAuth2TokenEndpointResponseBody
	GetRefreshToken() *string
	SetRequestId(v string) *OAuth2TokenEndpointResponseBody
	GetRequestId() *string
	SetScope(v string) *OAuth2TokenEndpointResponseBody
	GetScope() *string
	SetTokenType(v string) *OAuth2TokenEndpointResponseBody
	GetTokenType() *string
}

type OAuth2TokenEndpointResponseBody struct {
	// example:
	//
	// UJMiksSwuMJvwXrJLULMykSw6qZ6VqaxOkN4qd5cW1Q4HhsLxvUR5xVOIv1WB3br5LoP20lPa8xiYLSMbt8JqHACXdSdw7fNkhRTIHnadxWW5jfDg7BELUB0FcFfPiv0
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// example:
	//
	// 604799
	ExpiresIn *int64 `json:"ExpiresIn,omitempty" xml:"ExpiresIn,omitempty"`
	// example:
	//
	// zsEcmaUeb8-NZW4IIUDD7qdgBNflrj6fH8BXJYbW9iXihZTgvbcr1_utC9p5HJLn_lXVwhfivBTgUQZBCGvGl5lxqaxFhmFtt-OrBduFQKL9x8p2lpEMKlxuKHZZZJ3A
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	// example:
	//
	// 4070039E-5822-1F32-9295-1D2883E48BA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// aligenie:user:basic:read aligenie:iot:scene:read
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// Bearer
	TokenType *string `json:"TokenType,omitempty" xml:"TokenType,omitempty"`
}

func (s OAuth2TokenEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OAuth2TokenEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *OAuth2TokenEndpointResponseBody) GetAccessToken() *string {
	return s.AccessToken
}

func (s *OAuth2TokenEndpointResponseBody) GetExpiresIn() *int64 {
	return s.ExpiresIn
}

func (s *OAuth2TokenEndpointResponseBody) GetRefreshToken() *string {
	return s.RefreshToken
}

func (s *OAuth2TokenEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OAuth2TokenEndpointResponseBody) GetScope() *string {
	return s.Scope
}

func (s *OAuth2TokenEndpointResponseBody) GetTokenType() *string {
	return s.TokenType
}

func (s *OAuth2TokenEndpointResponseBody) SetAccessToken(v string) *OAuth2TokenEndpointResponseBody {
	s.AccessToken = &v
	return s
}

func (s *OAuth2TokenEndpointResponseBody) SetExpiresIn(v int64) *OAuth2TokenEndpointResponseBody {
	s.ExpiresIn = &v
	return s
}

func (s *OAuth2TokenEndpointResponseBody) SetRefreshToken(v string) *OAuth2TokenEndpointResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *OAuth2TokenEndpointResponseBody) SetRequestId(v string) *OAuth2TokenEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *OAuth2TokenEndpointResponseBody) SetScope(v string) *OAuth2TokenEndpointResponseBody {
	s.Scope = &v
	return s
}

func (s *OAuth2TokenEndpointResponseBody) SetTokenType(v string) *OAuth2TokenEndpointResponseBody {
	s.TokenType = &v
	return s
}

func (s *OAuth2TokenEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
