// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelProxyInput interface {
	dara.Model
	String() string
	GoString() string
	SetArmsConfiguration(v *ArmsConfiguration) *UpdateModelProxyInput
	GetArmsConfiguration() *ArmsConfiguration
	SetCredentialName(v string) *UpdateModelProxyInput
	GetCredentialName() *string
	SetDescription(v string) *UpdateModelProxyInput
	GetDescription() *string
	SetLogConfiguration(v *LogConfiguration) *UpdateModelProxyInput
	GetLogConfiguration() *LogConfiguration
	SetNetworkConfiguration(v *NetworkConfiguration) *UpdateModelProxyInput
	GetNetworkConfiguration() *NetworkConfiguration
	SetProxyConfig(v *ProxyConfig) *UpdateModelProxyInput
	GetProxyConfig() *ProxyConfig
}

type UpdateModelProxyInput struct {
	ArmsConfiguration    *ArmsConfiguration    `json:"armsConfiguration,omitempty" xml:"armsConfiguration,omitempty"`
	CredentialName       *string               `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	Description          *string               `json:"description,omitempty" xml:"description,omitempty"`
	LogConfiguration     *LogConfiguration     `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	ProxyConfig          *ProxyConfig          `json:"proxyConfig,omitempty" xml:"proxyConfig,omitempty"`
}

func (s UpdateModelProxyInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelProxyInput) GoString() string {
	return s.String()
}

func (s *UpdateModelProxyInput) GetArmsConfiguration() *ArmsConfiguration {
	return s.ArmsConfiguration
}

func (s *UpdateModelProxyInput) GetCredentialName() *string {
	return s.CredentialName
}

func (s *UpdateModelProxyInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelProxyInput) GetLogConfiguration() *LogConfiguration {
	return s.LogConfiguration
}

func (s *UpdateModelProxyInput) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *UpdateModelProxyInput) GetProxyConfig() *ProxyConfig {
	return s.ProxyConfig
}

func (s *UpdateModelProxyInput) SetArmsConfiguration(v *ArmsConfiguration) *UpdateModelProxyInput {
	s.ArmsConfiguration = v
	return s
}

func (s *UpdateModelProxyInput) SetCredentialName(v string) *UpdateModelProxyInput {
	s.CredentialName = &v
	return s
}

func (s *UpdateModelProxyInput) SetDescription(v string) *UpdateModelProxyInput {
	s.Description = &v
	return s
}

func (s *UpdateModelProxyInput) SetLogConfiguration(v *LogConfiguration) *UpdateModelProxyInput {
	s.LogConfiguration = v
	return s
}

func (s *UpdateModelProxyInput) SetNetworkConfiguration(v *NetworkConfiguration) *UpdateModelProxyInput {
	s.NetworkConfiguration = v
	return s
}

func (s *UpdateModelProxyInput) SetProxyConfig(v *ProxyConfig) *UpdateModelProxyInput {
	s.ProxyConfig = v
	return s
}

func (s *UpdateModelProxyInput) Validate() error {
	if s.ArmsConfiguration != nil {
		if err := s.ArmsConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.LogConfiguration != nil {
		if err := s.LogConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ProxyConfig != nil {
		if err := s.ProxyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
