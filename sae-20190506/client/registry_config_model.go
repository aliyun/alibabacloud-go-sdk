// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegistryConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v *RegistryAuthConfig) *RegistryConfig
	GetAuthConfig() *RegistryAuthConfig
	SetCertConfig(v *RegistryCertConfig) *RegistryConfig
	GetCertConfig() *RegistryCertConfig
}

type RegistryConfig struct {
	AuthConfig *RegistryAuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	CertConfig *RegistryCertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
}

func (s RegistryConfig) String() string {
	return dara.Prettify(s)
}

func (s RegistryConfig) GoString() string {
	return s.String()
}

func (s *RegistryConfig) GetAuthConfig() *RegistryAuthConfig {
	return s.AuthConfig
}

func (s *RegistryConfig) GetCertConfig() *RegistryCertConfig {
	return s.CertConfig
}

func (s *RegistryConfig) SetAuthConfig(v *RegistryAuthConfig) *RegistryConfig {
	s.AuthConfig = v
	return s
}

func (s *RegistryConfig) SetCertConfig(v *RegistryCertConfig) *RegistryConfig {
	s.CertConfig = v
	return s
}

func (s *RegistryConfig) Validate() error {
	if s.AuthConfig != nil {
		if err := s.AuthConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CertConfig != nil {
		if err := s.CertConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
