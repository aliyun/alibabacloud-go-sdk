// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTargetConfigurationInput interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *UpdateTargetConfigurationInput
	GetDomainId() *string
	SetTargetConfiguration(v *TargetConfiguration) *UpdateTargetConfigurationInput
	GetTargetConfiguration() *TargetConfiguration
}

type UpdateTargetConfigurationInput struct {
	DomainId            *string              `json:"domainId,omitempty" xml:"domainId,omitempty"`
	TargetConfiguration *TargetConfiguration `json:"targetConfiguration,omitempty" xml:"targetConfiguration,omitempty"`
}

func (s UpdateTargetConfigurationInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateTargetConfigurationInput) GoString() string {
	return s.String()
}

func (s *UpdateTargetConfigurationInput) GetDomainId() *string {
	return s.DomainId
}

func (s *UpdateTargetConfigurationInput) GetTargetConfiguration() *TargetConfiguration {
	return s.TargetConfiguration
}

func (s *UpdateTargetConfigurationInput) SetDomainId(v string) *UpdateTargetConfigurationInput {
	s.DomainId = &v
	return s
}

func (s *UpdateTargetConfigurationInput) SetTargetConfiguration(v *TargetConfiguration) *UpdateTargetConfigurationInput {
	s.TargetConfiguration = v
	return s
}

func (s *UpdateTargetConfigurationInput) Validate() error {
	if s.TargetConfiguration != nil {
		if err := s.TargetConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
