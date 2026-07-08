// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateInput interface {
	dara.Model
	String() string
	GoString() string
	SetLogConfiguration(v *LogConfiguration) *UpdateTemplateInput
	GetLogConfiguration() *LogConfiguration
	SetNetworkConfiguration(v *NetworkConfiguration) *UpdateTemplateInput
	GetNetworkConfiguration() *NetworkConfiguration
}

type UpdateTemplateInput struct {
	LogConfiguration     *LogConfiguration     `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
}

func (s UpdateTemplateInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateInput) GoString() string {
	return s.String()
}

func (s *UpdateTemplateInput) GetLogConfiguration() *LogConfiguration {
	return s.LogConfiguration
}

func (s *UpdateTemplateInput) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *UpdateTemplateInput) SetLogConfiguration(v *LogConfiguration) *UpdateTemplateInput {
	s.LogConfiguration = v
	return s
}

func (s *UpdateTemplateInput) SetNetworkConfiguration(v *NetworkConfiguration) *UpdateTemplateInput {
	s.NetworkConfiguration = v
	return s
}

func (s *UpdateTemplateInput) Validate() error {
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
	return nil
}
