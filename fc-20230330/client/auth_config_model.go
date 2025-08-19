// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAuthInfo(v string) *AuthConfig
	GetAuthInfo() *string
	SetAuthType(v string) *AuthConfig
	GetAuthType() *string
}

type AuthConfig struct {
	// example:
	//
	// {}
	AuthInfo *string `json:"authInfo,omitempty" xml:"authInfo,omitempty"`
	// example:
	//
	// anonymous, function, jwt
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
}

func (s AuthConfig) String() string {
	return dara.Prettify(s)
}

func (s AuthConfig) GoString() string {
	return s.String()
}

func (s *AuthConfig) GetAuthInfo() *string {
	return s.AuthInfo
}

func (s *AuthConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *AuthConfig) SetAuthInfo(v string) *AuthConfig {
	s.AuthInfo = &v
	return s
}

func (s *AuthConfig) SetAuthType(v string) *AuthConfig {
	s.AuthType = &v
	return s
}

func (s *AuthConfig) Validate() error {
	return dara.Validate(s)
}
