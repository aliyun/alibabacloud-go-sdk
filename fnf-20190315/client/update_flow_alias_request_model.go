// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateFlowAliasRequest
	GetDescription() *string
	SetFlowName(v string) *UpdateFlowAliasRequest
	GetFlowName() *string
	SetName(v string) *UpdateFlowAliasRequest
	GetName() *string
	SetRoutingConfigurations(v []*UpdateFlowAliasRequestRoutingConfigurations) *UpdateFlowAliasRequest
	GetRoutingConfigurations() []*UpdateFlowAliasRequestRoutingConfigurations
}

type UpdateFlowAliasRequest struct {
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
	// alias name
	Name                  *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	RoutingConfigurations []*UpdateFlowAliasRequestRoutingConfigurations `json:"RoutingConfigurations,omitempty" xml:"RoutingConfigurations,omitempty" type:"Repeated"`
}

func (s UpdateFlowAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowAliasRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowAliasRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateFlowAliasRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *UpdateFlowAliasRequest) GetName() *string {
	return s.Name
}

func (s *UpdateFlowAliasRequest) GetRoutingConfigurations() []*UpdateFlowAliasRequestRoutingConfigurations {
	return s.RoutingConfigurations
}

func (s *UpdateFlowAliasRequest) SetDescription(v string) *UpdateFlowAliasRequest {
	s.Description = &v
	return s
}

func (s *UpdateFlowAliasRequest) SetFlowName(v string) *UpdateFlowAliasRequest {
	s.FlowName = &v
	return s
}

func (s *UpdateFlowAliasRequest) SetName(v string) *UpdateFlowAliasRequest {
	s.Name = &v
	return s
}

func (s *UpdateFlowAliasRequest) SetRoutingConfigurations(v []*UpdateFlowAliasRequestRoutingConfigurations) *UpdateFlowAliasRequest {
	s.RoutingConfigurations = v
	return s
}

func (s *UpdateFlowAliasRequest) Validate() error {
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

type UpdateFlowAliasRequestRoutingConfigurations struct {
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// 20
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateFlowAliasRequestRoutingConfigurations) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowAliasRequestRoutingConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateFlowAliasRequestRoutingConfigurations) GetVersion() *string {
	return s.Version
}

func (s *UpdateFlowAliasRequestRoutingConfigurations) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateFlowAliasRequestRoutingConfigurations) SetVersion(v string) *UpdateFlowAliasRequestRoutingConfigurations {
	s.Version = &v
	return s
}

func (s *UpdateFlowAliasRequestRoutingConfigurations) SetWeight(v int32) *UpdateFlowAliasRequestRoutingConfigurations {
	s.Weight = &v
	return s
}

func (s *UpdateFlowAliasRequestRoutingConfigurations) Validate() error {
	return dara.Validate(s)
}
