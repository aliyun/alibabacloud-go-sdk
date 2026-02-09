// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMode(v string) *AuthConfig
	GetAuthMode() *string
	SetAuthType(v string) *AuthConfig
	GetAuthType() *string
}

type AuthConfig struct {
	// The authentication mode.
	//
	// example:
	//
	// NoAuth
	AuthMode *string `json:"authMode,omitempty" xml:"authMode,omitempty"`
	// The authentication type.
	//
	// example:
	//
	// Jwt
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
}

func (s AuthConfig) String() string {
	return dara.Prettify(s)
}

func (s AuthConfig) GoString() string {
	return s.String()
}

func (s *AuthConfig) GetAuthMode() *string {
	return s.AuthMode
}

func (s *AuthConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *AuthConfig) SetAuthMode(v string) *AuthConfig {
	s.AuthMode = &v
	return s
}

func (s *AuthConfig) SetAuthType(v string) *AuthConfig {
	s.AuthType = &v
	return s
}

func (s *AuthConfig) Validate() error {
	return dara.Validate(s)
}
