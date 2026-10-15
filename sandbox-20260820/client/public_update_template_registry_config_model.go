// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicUpdateTemplateRegistryConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v *PublicUpdateTemplateRegistryAuthConfig) *PublicUpdateTemplateRegistryConfig
	GetAuthConfig() *PublicUpdateTemplateRegistryAuthConfig
	SetCertConfig(v *PublicUpdateTemplateRegistryCertConfig) *PublicUpdateTemplateRegistryConfig
	GetCertConfig() *PublicUpdateTemplateRegistryCertConfig
	SetNetworkConfig(v *PublicUpdateTemplateRegistryNetworkConfig) *PublicUpdateTemplateRegistryConfig
	GetNetworkConfig() *PublicUpdateTemplateRegistryNetworkConfig
}

type PublicUpdateTemplateRegistryConfig struct {
	// The image repository authentication configuration.
	AuthConfig *PublicUpdateTemplateRegistryAuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// The image repository certificate configuration.
	CertConfig *PublicUpdateTemplateRegistryCertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	// The image repository network configuration.
	NetworkConfig *PublicUpdateTemplateRegistryNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty"`
}

func (s PublicUpdateTemplateRegistryConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicUpdateTemplateRegistryConfig) GoString() string {
	return s.String()
}

func (s *PublicUpdateTemplateRegistryConfig) GetAuthConfig() *PublicUpdateTemplateRegistryAuthConfig {
	return s.AuthConfig
}

func (s *PublicUpdateTemplateRegistryConfig) GetCertConfig() *PublicUpdateTemplateRegistryCertConfig {
	return s.CertConfig
}

func (s *PublicUpdateTemplateRegistryConfig) GetNetworkConfig() *PublicUpdateTemplateRegistryNetworkConfig {
	return s.NetworkConfig
}

func (s *PublicUpdateTemplateRegistryConfig) SetAuthConfig(v *PublicUpdateTemplateRegistryAuthConfig) *PublicUpdateTemplateRegistryConfig {
	s.AuthConfig = v
	return s
}

func (s *PublicUpdateTemplateRegistryConfig) SetCertConfig(v *PublicUpdateTemplateRegistryCertConfig) *PublicUpdateTemplateRegistryConfig {
	s.CertConfig = v
	return s
}

func (s *PublicUpdateTemplateRegistryConfig) SetNetworkConfig(v *PublicUpdateTemplateRegistryNetworkConfig) *PublicUpdateTemplateRegistryConfig {
	s.NetworkConfig = v
	return s
}

func (s *PublicUpdateTemplateRegistryConfig) Validate() error {
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
	if s.NetworkConfig != nil {
		if err := s.NetworkConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
