// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebApplicationInput interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *CreateWebApplicationInput
	GetApplicationName() *string
	SetDescription(v string) *CreateWebApplicationInput
	GetDescription() *string
	SetRevisionConfig(v *RevisionConfig) *CreateWebApplicationInput
	GetRevisionConfig() *RevisionConfig
	SetWebNetworkConfig(v *WebNetworkConfig) *CreateWebApplicationInput
	GetWebNetworkConfig() *WebNetworkConfig
	SetWebScalingConfig(v *WebScalingConfig) *CreateWebApplicationInput
	GetWebScalingConfig() *WebScalingConfig
	SetWebTrafficConfig(v *WebTrafficConfig) *CreateWebApplicationInput
	GetWebTrafficConfig() *WebTrafficConfig
}

type CreateWebApplicationInput struct {
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// sae-app
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// my sae app
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The configuration on the application revision.
	//
	// This parameter is required.
	RevisionConfig *RevisionConfig `json:"RevisionConfig,omitempty" xml:"RevisionConfig,omitempty"`
	// The network configurations.
	//
	// >  This field is used to configure network for the application.
	WebNetworkConfig *WebNetworkConfig `json:"WebNetworkConfig,omitempty" xml:"WebNetworkConfig,omitempty"`
	// The scaling configurations of the applications.
	WebScalingConfig *WebScalingConfig `json:"WebScalingConfig,omitempty" xml:"WebScalingConfig,omitempty"`
	// The traffic configurations of the applications.
	WebTrafficConfig *WebTrafficConfig `json:"WebTrafficConfig,omitempty" xml:"WebTrafficConfig,omitempty"`
}

func (s CreateWebApplicationInput) String() string {
	return dara.Prettify(s)
}

func (s CreateWebApplicationInput) GoString() string {
	return s.String()
}

func (s *CreateWebApplicationInput) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *CreateWebApplicationInput) GetDescription() *string {
	return s.Description
}

func (s *CreateWebApplicationInput) GetRevisionConfig() *RevisionConfig {
	return s.RevisionConfig
}

func (s *CreateWebApplicationInput) GetWebNetworkConfig() *WebNetworkConfig {
	return s.WebNetworkConfig
}

func (s *CreateWebApplicationInput) GetWebScalingConfig() *WebScalingConfig {
	return s.WebScalingConfig
}

func (s *CreateWebApplicationInput) GetWebTrafficConfig() *WebTrafficConfig {
	return s.WebTrafficConfig
}

func (s *CreateWebApplicationInput) SetApplicationName(v string) *CreateWebApplicationInput {
	s.ApplicationName = &v
	return s
}

func (s *CreateWebApplicationInput) SetDescription(v string) *CreateWebApplicationInput {
	s.Description = &v
	return s
}

func (s *CreateWebApplicationInput) SetRevisionConfig(v *RevisionConfig) *CreateWebApplicationInput {
	s.RevisionConfig = v
	return s
}

func (s *CreateWebApplicationInput) SetWebNetworkConfig(v *WebNetworkConfig) *CreateWebApplicationInput {
	s.WebNetworkConfig = v
	return s
}

func (s *CreateWebApplicationInput) SetWebScalingConfig(v *WebScalingConfig) *CreateWebApplicationInput {
	s.WebScalingConfig = v
	return s
}

func (s *CreateWebApplicationInput) SetWebTrafficConfig(v *WebTrafficConfig) *CreateWebApplicationInput {
	s.WebTrafficConfig = v
	return s
}

func (s *CreateWebApplicationInput) Validate() error {
	if s.RevisionConfig != nil {
		if err := s.RevisionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.WebNetworkConfig != nil {
		if err := s.WebNetworkConfig.Validate(); err != nil {
			return err
		}
	}
	if s.WebScalingConfig != nil {
		if err := s.WebScalingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.WebTrafficConfig != nil {
		if err := s.WebTrafficConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
