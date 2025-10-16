// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayTargetInput interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *CreateGatewayTargetInput
	GetDomainId() *string
	SetName(v string) *CreateGatewayTargetInput
	GetName() *string
	SetTargetConfiguration(v *TargetConfiguration) *CreateGatewayTargetInput
	GetTargetConfiguration() *TargetConfiguration
}

type CreateGatewayTargetInput struct {
	DomainId            *string              `json:"domainId,omitempty" xml:"domainId,omitempty"`
	Name                *string              `json:"name,omitempty" xml:"name,omitempty"`
	TargetConfiguration *TargetConfiguration `json:"targetConfiguration,omitempty" xml:"targetConfiguration,omitempty"`
}

func (s CreateGatewayTargetInput) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayTargetInput) GoString() string {
	return s.String()
}

func (s *CreateGatewayTargetInput) GetDomainId() *string {
	return s.DomainId
}

func (s *CreateGatewayTargetInput) GetName() *string {
	return s.Name
}

func (s *CreateGatewayTargetInput) GetTargetConfiguration() *TargetConfiguration {
	return s.TargetConfiguration
}

func (s *CreateGatewayTargetInput) SetDomainId(v string) *CreateGatewayTargetInput {
	s.DomainId = &v
	return s
}

func (s *CreateGatewayTargetInput) SetName(v string) *CreateGatewayTargetInput {
	s.Name = &v
	return s
}

func (s *CreateGatewayTargetInput) SetTargetConfiguration(v *TargetConfiguration) *CreateGatewayTargetInput {
	s.TargetConfiguration = v
	return s
}

func (s *CreateGatewayTargetInput) Validate() error {
	if s.TargetConfiguration != nil {
		if err := s.TargetConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
