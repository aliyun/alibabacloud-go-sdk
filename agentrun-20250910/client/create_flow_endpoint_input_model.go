// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowEndpointInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateFlowEndpointInput
	GetDescription() *string
	SetDisablePublicNetworkAccess(v bool) *CreateFlowEndpointInput
	GetDisablePublicNetworkAccess() *bool
	SetFlowEndpointName(v string) *CreateFlowEndpointInput
	GetFlowEndpointName() *string
	SetRoutingConfiguration(v []*FlowEndpointRoutingConfig) *CreateFlowEndpointInput
	GetRoutingConfiguration() []*FlowEndpointRoutingConfig
	SetTargetVersion(v string) *CreateFlowEndpointInput
	GetTargetVersion() *string
}

type CreateFlowEndpointInput struct {
	// The description of the flow endpoint.
	//
	// example:
	//
	// Production endpoint for flow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Whether to disable public network access for the endpoint. If unspecified, the endpoint inherits this setting from its parent workflow.
	DisablePublicNetworkAccess *bool `json:"disablePublicNetworkAccess,omitempty" xml:"disablePublicNetworkAccess,omitempty"`
	// The unique name of the flow endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-flow-endpoint
	FlowEndpointName *string `json:"flowEndpointName,omitempty" xml:"flowEndpointName,omitempty"`
	// The routing configuration that defines traffic distribution across different flow versions.
	//
	// example:
	//
	// []
	RoutingConfiguration []*FlowEndpointRoutingConfig `json:"routingConfiguration" xml:"routingConfiguration" type:"Repeated"`
	// The target version for the flow endpoint.
	//
	// example:
	//
	// 1
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
}

func (s CreateFlowEndpointInput) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowEndpointInput) GoString() string {
	return s.String()
}

func (s *CreateFlowEndpointInput) GetDescription() *string {
	return s.Description
}

func (s *CreateFlowEndpointInput) GetDisablePublicNetworkAccess() *bool {
	return s.DisablePublicNetworkAccess
}

func (s *CreateFlowEndpointInput) GetFlowEndpointName() *string {
	return s.FlowEndpointName
}

func (s *CreateFlowEndpointInput) GetRoutingConfiguration() []*FlowEndpointRoutingConfig {
	return s.RoutingConfiguration
}

func (s *CreateFlowEndpointInput) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *CreateFlowEndpointInput) SetDescription(v string) *CreateFlowEndpointInput {
	s.Description = &v
	return s
}

func (s *CreateFlowEndpointInput) SetDisablePublicNetworkAccess(v bool) *CreateFlowEndpointInput {
	s.DisablePublicNetworkAccess = &v
	return s
}

func (s *CreateFlowEndpointInput) SetFlowEndpointName(v string) *CreateFlowEndpointInput {
	s.FlowEndpointName = &v
	return s
}

func (s *CreateFlowEndpointInput) SetRoutingConfiguration(v []*FlowEndpointRoutingConfig) *CreateFlowEndpointInput {
	s.RoutingConfiguration = v
	return s
}

func (s *CreateFlowEndpointInput) SetTargetVersion(v string) *CreateFlowEndpointInput {
	s.TargetVersion = &v
	return s
}

func (s *CreateFlowEndpointInput) Validate() error {
	if s.RoutingConfiguration != nil {
		for _, item := range s.RoutingConfiguration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
