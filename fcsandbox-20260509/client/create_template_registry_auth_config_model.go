// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateRegistryAuthConfig interface {
	dara.Model
	String() string
	GoString() string
	SetPassword(v string) *CreateTemplateRegistryAuthConfig
	GetPassword() *string
	SetUserName(v string) *CreateTemplateRegistryAuthConfig
	GetUserName() *string
}

type CreateTemplateRegistryAuthConfig struct {
	// The password of the image repository.
	//
	// example:
	//
	// ******
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The username of the image repository.
	//
	// example:
	//
	// my-user
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s CreateTemplateRegistryAuthConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRegistryAuthConfig) GoString() string {
	return s.String()
}

func (s *CreateTemplateRegistryAuthConfig) GetPassword() *string {
	return s.Password
}

func (s *CreateTemplateRegistryAuthConfig) GetUserName() *string {
	return s.UserName
}

func (s *CreateTemplateRegistryAuthConfig) SetPassword(v string) *CreateTemplateRegistryAuthConfig {
	s.Password = &v
	return s
}

func (s *CreateTemplateRegistryAuthConfig) SetUserName(v string) *CreateTemplateRegistryAuthConfig {
	s.UserName = &v
	return s
}

func (s *CreateTemplateRegistryAuthConfig) Validate() error {
	return dara.Validate(s)
}
