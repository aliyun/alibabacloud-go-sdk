// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegistryAuthConfig interface {
	dara.Model
	String() string
	GoString() string
	SetPassword(v string) *RegistryAuthConfig
	GetPassword() *string
	SetUserName(v string) *RegistryAuthConfig
	GetUserName() *string
}

type RegistryAuthConfig struct {
	// example:
	//
	// abc***
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// example:
	//
	// admin
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s RegistryAuthConfig) String() string {
	return dara.Prettify(s)
}

func (s RegistryAuthConfig) GoString() string {
	return s.String()
}

func (s *RegistryAuthConfig) GetPassword() *string {
	return s.Password
}

func (s *RegistryAuthConfig) GetUserName() *string {
	return s.UserName
}

func (s *RegistryAuthConfig) SetPassword(v string) *RegistryAuthConfig {
	s.Password = &v
	return s
}

func (s *RegistryAuthConfig) SetUserName(v string) *RegistryAuthConfig {
	s.UserName = &v
	return s
}

func (s *RegistryAuthConfig) Validate() error {
	return dara.Validate(s)
}
