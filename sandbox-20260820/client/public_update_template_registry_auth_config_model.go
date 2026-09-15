// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicUpdateTemplateRegistryAuthConfig interface {
	dara.Model
	String() string
	GoString() string
	SetPassword(v string) *PublicUpdateTemplateRegistryAuthConfig
	GetPassword() *string
	SetUserName(v string) *PublicUpdateTemplateRegistryAuthConfig
	GetUserName() *string
}

type PublicUpdateTemplateRegistryAuthConfig struct {
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s PublicUpdateTemplateRegistryAuthConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicUpdateTemplateRegistryAuthConfig) GoString() string {
	return s.String()
}

func (s *PublicUpdateTemplateRegistryAuthConfig) GetPassword() *string {
	return s.Password
}

func (s *PublicUpdateTemplateRegistryAuthConfig) GetUserName() *string {
	return s.UserName
}

func (s *PublicUpdateTemplateRegistryAuthConfig) SetPassword(v string) *PublicUpdateTemplateRegistryAuthConfig {
	s.Password = &v
	return s
}

func (s *PublicUpdateTemplateRegistryAuthConfig) SetUserName(v string) *PublicUpdateTemplateRegistryAuthConfig {
	s.UserName = &v
	return s
}

func (s *PublicUpdateTemplateRegistryAuthConfig) Validate() error {
	return dara.Validate(s)
}
