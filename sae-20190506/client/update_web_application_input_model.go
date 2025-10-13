// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebApplicationInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateWebApplicationInput
	GetDescription() *string
	SetWebNetworkConfig(v *WebNetworkConfig) *UpdateWebApplicationInput
	GetWebNetworkConfig() *WebNetworkConfig
}

type UpdateWebApplicationInput struct {
	// example:
	//
	// sae-app
	Description      *string           `json:"Description,omitempty" xml:"Description,omitempty"`
	WebNetworkConfig *WebNetworkConfig `json:"WebNetworkConfig,omitempty" xml:"WebNetworkConfig,omitempty"`
}

func (s UpdateWebApplicationInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebApplicationInput) GoString() string {
	return s.String()
}

func (s *UpdateWebApplicationInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateWebApplicationInput) GetWebNetworkConfig() *WebNetworkConfig {
	return s.WebNetworkConfig
}

func (s *UpdateWebApplicationInput) SetDescription(v string) *UpdateWebApplicationInput {
	s.Description = &v
	return s
}

func (s *UpdateWebApplicationInput) SetWebNetworkConfig(v *WebNetworkConfig) *UpdateWebApplicationInput {
	s.WebNetworkConfig = v
	return s
}

func (s *UpdateWebApplicationInput) Validate() error {
	if s.WebNetworkConfig != nil {
		if err := s.WebNetworkConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
