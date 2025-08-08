// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOAuthCredential interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v int64) *OAuthCredential
	GetCreatedTime() *int64
	SetExpiration(v int64) *OAuthCredential
	GetExpiration() *int64
	SetRefreshToken(v string) *OAuthCredential
	GetRefreshToken() *string
	SetScope(v string) *OAuthCredential
	GetScope() *string
	SetToken(v string) *OAuthCredential
	GetToken() *string
	SetType(v string) *OAuthCredential
	GetType() *string
}

type OAuthCredential struct {
	// This parameter is required.
	//
	// example:
	//
	// 1716176924603
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1716263324603
	Expiration *int64 `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// example:
	//
	// 4d77bfae284770d94ebeed6b0199ebfd65e3943ba4f1e44dc36d792a93ba0d13
	RefreshToken *string `json:"refreshToken,omitempty" xml:"refreshToken,omitempty"`
	// example:
	//
	// user_info projects pull_requests hook gists emails
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4e84246b6b3962cd3d207aad1ea2f911
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// example:
	//
	// bearer
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s OAuthCredential) String() string {
	return dara.Prettify(s)
}

func (s OAuthCredential) GoString() string {
	return s.String()
}

func (s *OAuthCredential) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *OAuthCredential) GetExpiration() *int64 {
	return s.Expiration
}

func (s *OAuthCredential) GetRefreshToken() *string {
	return s.RefreshToken
}

func (s *OAuthCredential) GetScope() *string {
	return s.Scope
}

func (s *OAuthCredential) GetToken() *string {
	return s.Token
}

func (s *OAuthCredential) GetType() *string {
	return s.Type
}

func (s *OAuthCredential) SetCreatedTime(v int64) *OAuthCredential {
	s.CreatedTime = &v
	return s
}

func (s *OAuthCredential) SetExpiration(v int64) *OAuthCredential {
	s.Expiration = &v
	return s
}

func (s *OAuthCredential) SetRefreshToken(v string) *OAuthCredential {
	s.RefreshToken = &v
	return s
}

func (s *OAuthCredential) SetScope(v string) *OAuthCredential {
	s.Scope = &v
	return s
}

func (s *OAuthCredential) SetToken(v string) *OAuthCredential {
	s.Token = &v
	return s
}

func (s *OAuthCredential) SetType(v string) *OAuthCredential {
	s.Type = &v
	return s
}

func (s *OAuthCredential) Validate() error {
	return dara.Validate(s)
}
