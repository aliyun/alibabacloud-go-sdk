// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateInput interface {
	dara.Model
	String() string
	GoString() string
	SetContainerConfiguration(v *ContainerConfiguration) *UpdateTemplateInput
	GetContainerConfiguration() *ContainerConfiguration
	SetLogConfiguration(v *LogConfiguration) *UpdateTemplateInput
	GetLogConfiguration() *LogConfiguration
	SetNetworkConfiguration(v *NetworkConfiguration) *UpdateTemplateInput
	GetNetworkConfiguration() *NetworkConfiguration
	SetTeamID(v string) *UpdateTemplateInput
	GetTeamID() *string
}

type UpdateTemplateInput struct {
	ContainerConfiguration *ContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty"`
	LogConfiguration       *LogConfiguration       `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	NetworkConfiguration   *NetworkConfiguration   `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	TeamID                 *string                 `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s UpdateTemplateInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateInput) GoString() string {
	return s.String()
}

func (s *UpdateTemplateInput) GetContainerConfiguration() *ContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *UpdateTemplateInput) GetLogConfiguration() *LogConfiguration {
	return s.LogConfiguration
}

func (s *UpdateTemplateInput) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *UpdateTemplateInput) GetTeamID() *string {
	return s.TeamID
}

func (s *UpdateTemplateInput) SetContainerConfiguration(v *ContainerConfiguration) *UpdateTemplateInput {
	s.ContainerConfiguration = v
	return s
}

func (s *UpdateTemplateInput) SetLogConfiguration(v *LogConfiguration) *UpdateTemplateInput {
	s.LogConfiguration = v
	return s
}

func (s *UpdateTemplateInput) SetNetworkConfiguration(v *NetworkConfiguration) *UpdateTemplateInput {
	s.NetworkConfiguration = v
	return s
}

func (s *UpdateTemplateInput) SetTeamID(v string) *UpdateTemplateInput {
	s.TeamID = &v
	return s
}

func (s *UpdateTemplateInput) Validate() error {
	if s.ContainerConfiguration != nil {
		if err := s.ContainerConfiguration.Validate(); err != nil {
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
	return nil
}
