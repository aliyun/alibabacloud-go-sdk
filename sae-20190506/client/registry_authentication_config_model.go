// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegistryAuthenticationConfig interface {
	dara.Model
	String() string
	GoString() string
	SetPassword(v string) *RegistryAuthenticationConfig
	GetPassword() *string
	SetUserName(v string) *RegistryAuthenticationConfig
	GetUserName() *string
}

type RegistryAuthenticationConfig struct {
	// example:
	//
	// abc***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// admin
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s RegistryAuthenticationConfig) String() string {
	return dara.Prettify(s)
}

func (s RegistryAuthenticationConfig) GoString() string {
	return s.String()
}

func (s *RegistryAuthenticationConfig) GetPassword() *string {
	return s.Password
}

func (s *RegistryAuthenticationConfig) GetUserName() *string {
	return s.UserName
}

func (s *RegistryAuthenticationConfig) SetPassword(v string) *RegistryAuthenticationConfig {
	s.Password = &v
	return s
}

func (s *RegistryAuthenticationConfig) SetUserName(v string) *RegistryAuthenticationConfig {
	s.UserName = &v
	return s
}

func (s *RegistryAuthenticationConfig) Validate() error {
	return dara.Validate(s)
}
