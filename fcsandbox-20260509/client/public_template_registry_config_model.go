// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicTemplateRegistryConfig interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkConfig(v *PublicTemplateRegistryNetworkConfig) *PublicTemplateRegistryConfig
	GetNetworkConfig() *PublicTemplateRegistryNetworkConfig
}

type PublicTemplateRegistryConfig struct {
	NetworkConfig *PublicTemplateRegistryNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty"`
}

func (s PublicTemplateRegistryConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicTemplateRegistryConfig) GoString() string {
	return s.String()
}

func (s *PublicTemplateRegistryConfig) GetNetworkConfig() *PublicTemplateRegistryNetworkConfig {
	return s.NetworkConfig
}

func (s *PublicTemplateRegistryConfig) SetNetworkConfig(v *PublicTemplateRegistryNetworkConfig) *PublicTemplateRegistryConfig {
	s.NetworkConfig = v
	return s
}

func (s *PublicTemplateRegistryConfig) Validate() error {
	if s.NetworkConfig != nil {
		if err := s.NetworkConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
