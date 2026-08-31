// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateRegistryConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v *CreateTemplateRegistryAuthConfig) *CreateTemplateRegistryConfig
	GetAuthConfig() *CreateTemplateRegistryAuthConfig
	SetCertConfig(v *CreateTemplateRegistryCertConfig) *CreateTemplateRegistryConfig
	GetCertConfig() *CreateTemplateRegistryCertConfig
	SetNetworkConfig(v *CreateTemplateRegistryNetworkConfig) *CreateTemplateRegistryConfig
	GetNetworkConfig() *CreateTemplateRegistryNetworkConfig
}

type CreateTemplateRegistryConfig struct {
	// The image repository authentication configuration.
	AuthConfig *CreateTemplateRegistryAuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// The image repository certificate configuration.
	CertConfig *CreateTemplateRegistryCertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	// The image repository network configuration.
	NetworkConfig *CreateTemplateRegistryNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty"`
}

func (s CreateTemplateRegistryConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRegistryConfig) GoString() string {
	return s.String()
}

func (s *CreateTemplateRegistryConfig) GetAuthConfig() *CreateTemplateRegistryAuthConfig {
	return s.AuthConfig
}

func (s *CreateTemplateRegistryConfig) GetCertConfig() *CreateTemplateRegistryCertConfig {
	return s.CertConfig
}

func (s *CreateTemplateRegistryConfig) GetNetworkConfig() *CreateTemplateRegistryNetworkConfig {
	return s.NetworkConfig
}

func (s *CreateTemplateRegistryConfig) SetAuthConfig(v *CreateTemplateRegistryAuthConfig) *CreateTemplateRegistryConfig {
	s.AuthConfig = v
	return s
}

func (s *CreateTemplateRegistryConfig) SetCertConfig(v *CreateTemplateRegistryCertConfig) *CreateTemplateRegistryConfig {
	s.CertConfig = v
	return s
}

func (s *CreateTemplateRegistryConfig) SetNetworkConfig(v *CreateTemplateRegistryNetworkConfig) *CreateTemplateRegistryConfig {
	s.NetworkConfig = v
	return s
}

func (s *CreateTemplateRegistryConfig) Validate() error {
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
