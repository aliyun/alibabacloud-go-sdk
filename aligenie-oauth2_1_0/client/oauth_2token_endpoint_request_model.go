// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOAuth2TokenEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OAuth2TokenEndpointRequest
	GetCode() *string
	SetGrantType(v string) *OAuth2TokenEndpointRequest
	GetGrantType() *string
	SetRedirectUri(v string) *OAuth2TokenEndpointRequest
	GetRedirectUri() *string
	SetRefreshToken(v string) *OAuth2TokenEndpointRequest
	GetRefreshToken() *string
}

type OAuth2TokenEndpointRequest struct {
	// example:
	//
	// rf3mi4JOU-xRIX2zEuRLHi-U9mPnvISeSphbwiBHJ5mEKZtG-xJsbBWrq8RmhQEPRYh0JOd3DaS_VZ90soD_YrsT4OBtgD06DmdIKL2_5KFfI6p_SjXX2-UMJuGfXDkB
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// authorization_code
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
	// example:
	//
	// https://xxx.xxx.com/xxx
	RedirectUri *string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty"`
	// example:
	//
	// zsEcmaUeb8-NZW4IIUDD7qdgBNflrj6fH8BXJYbW9iXihZTgvbcr1_utC9p5HJLn_lXVwhfivBTgUQZBCGvGl5lxqaxFhmFtt-OrBduFQKL9x8p2lpEMKlxuKHZZZJ3A
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
}

func (s OAuth2TokenEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s OAuth2TokenEndpointRequest) GoString() string {
	return s.String()
}

func (s *OAuth2TokenEndpointRequest) GetCode() *string {
	return s.Code
}

func (s *OAuth2TokenEndpointRequest) GetGrantType() *string {
	return s.GrantType
}

func (s *OAuth2TokenEndpointRequest) GetRedirectUri() *string {
	return s.RedirectUri
}

func (s *OAuth2TokenEndpointRequest) GetRefreshToken() *string {
	return s.RefreshToken
}

func (s *OAuth2TokenEndpointRequest) SetCode(v string) *OAuth2TokenEndpointRequest {
	s.Code = &v
	return s
}

func (s *OAuth2TokenEndpointRequest) SetGrantType(v string) *OAuth2TokenEndpointRequest {
	s.GrantType = &v
	return s
}

func (s *OAuth2TokenEndpointRequest) SetRedirectUri(v string) *OAuth2TokenEndpointRequest {
	s.RedirectUri = &v
	return s
}

func (s *OAuth2TokenEndpointRequest) SetRefreshToken(v string) *OAuth2TokenEndpointRequest {
	s.RefreshToken = &v
	return s
}

func (s *OAuth2TokenEndpointRequest) Validate() error {
	return dara.Validate(s)
}
