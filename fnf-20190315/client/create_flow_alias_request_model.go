// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateFlowAliasRequest
	GetDescription() *string
	SetFlowName(v string) *CreateFlowAliasRequest
	GetFlowName() *string
	SetName(v string) *CreateFlowAliasRequest
	GetName() *string
	SetRoutingConfigurations(v []*CreateFlowAliasRequestRoutingConfigurations) *CreateFlowAliasRequest
	GetRoutingConfigurations() []*CreateFlowAliasRequestRoutingConfigurations
}

type CreateFlowAliasRequest struct {
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example-flow-name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example-alias-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	RoutingConfigurations []*CreateFlowAliasRequestRoutingConfigurations `json:"RoutingConfigurations,omitempty" xml:"RoutingConfigurations,omitempty" type:"Repeated"`
}

func (s CreateFlowAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowAliasRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowAliasRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFlowAliasRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *CreateFlowAliasRequest) GetName() *string {
	return s.Name
}

func (s *CreateFlowAliasRequest) GetRoutingConfigurations() []*CreateFlowAliasRequestRoutingConfigurations {
	return s.RoutingConfigurations
}

func (s *CreateFlowAliasRequest) SetDescription(v string) *CreateFlowAliasRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowAliasRequest) SetFlowName(v string) *CreateFlowAliasRequest {
	s.FlowName = &v
	return s
}

func (s *CreateFlowAliasRequest) SetName(v string) *CreateFlowAliasRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowAliasRequest) SetRoutingConfigurations(v []*CreateFlowAliasRequestRoutingConfigurations) *CreateFlowAliasRequest {
	s.RoutingConfigurations = v
	return s
}

func (s *CreateFlowAliasRequest) Validate() error {
	if s.RoutingConfigurations != nil {
		for _, item := range s.RoutingConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFlowAliasRequestRoutingConfigurations struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateFlowAliasRequestRoutingConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowAliasRequestRoutingConfigurations) GoString() string {
	return s.String()
}

func (s *CreateFlowAliasRequestRoutingConfigurations) GetVersion() *string {
	return s.Version
}

func (s *CreateFlowAliasRequestRoutingConfigurations) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateFlowAliasRequestRoutingConfigurations) SetVersion(v string) *CreateFlowAliasRequestRoutingConfigurations {
	s.Version = &v
	return s
}

func (s *CreateFlowAliasRequestRoutingConfigurations) SetWeight(v int32) *CreateFlowAliasRequestRoutingConfigurations {
	s.Weight = &v
	return s
}

func (s *CreateFlowAliasRequestRoutingConfigurations) Validate() error {
	return dara.Validate(s)
}
