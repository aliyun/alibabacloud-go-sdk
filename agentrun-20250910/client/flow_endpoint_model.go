// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowEndpoint interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *FlowEndpoint
	GetCreatedAt() *string
	SetDescription(v string) *FlowEndpoint
	GetDescription() *string
	SetDisablePublicNetworkAccess(v bool) *FlowEndpoint
	GetDisablePublicNetworkAccess() *bool
	SetFlowEndpointArn(v string) *FlowEndpoint
	GetFlowEndpointArn() *string
	SetFlowEndpointId(v string) *FlowEndpoint
	GetFlowEndpointId() *string
	SetFlowEndpointName(v string) *FlowEndpoint
	GetFlowEndpointName() *string
	SetFlowId(v string) *FlowEndpoint
	GetFlowId() *string
	SetLastUpdatedAt(v string) *FlowEndpoint
	GetLastUpdatedAt() *string
	SetRoutingConfiguration(v []*FlowEndpointRoutingConfig) *FlowEndpoint
	GetRoutingConfiguration() []*FlowEndpointRoutingConfig
	SetTargetVersion(v string) *FlowEndpoint
	GetTargetVersion() *string
}

type FlowEndpoint struct {
	// The creation timestamp for the flow endpoint, in ISO 8601 format.
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The description of the flow endpoint.
	//
	// example:
	//
	// Production endpoint for flow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Set to true to disable public network access for this endpoint.
	DisablePublicNetworkAccess *bool `json:"disablePublicNetworkAccess,omitempty" xml:"disablePublicNetworkAccess,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the flow endpoint.
	//
	// example:
	//
	// acs:agentrun:cn-hangzhou:1760720386195983:workspaces/ws-xxx/flows/flow-xxx/endpoints/fe-xxx
	FlowEndpointArn *string `json:"flowEndpointArn,omitempty" xml:"flowEndpointArn,omitempty"`
	// The unique identifier of the flow endpoint.
	//
	// example:
	//
	// fe-1234567890abcdef
	FlowEndpointId *string `json:"flowEndpointId,omitempty" xml:"flowEndpointId,omitempty"`
	// The name of the flow endpoint.
	//
	// example:
	//
	// my-flow-endpoint
	FlowEndpointName *string `json:"flowEndpointName,omitempty" xml:"flowEndpointName,omitempty"`
	// The unique identifier of the flow.
	//
	// example:
	//
	// flow-1234567890abcdef
	FlowId *string `json:"flowId,omitempty" xml:"flowId,omitempty"`
	// The last update timestamp for the flow endpoint, in ISO 8601 format.
	//
	// example:
	//
	// 2025-01-10T11:45:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// The routing configuration that distributes traffic to different versions of the flow.
	//
	// example:
	//
	// []
	RoutingConfiguration []*FlowEndpointRoutingConfig `json:"routingConfiguration" xml:"routingConfiguration" type:"Repeated"`
	// The flow version targeted by the endpoint.
	//
	// example:
	//
	// 1
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
}

func (s FlowEndpoint) String() string {
	return dara.Prettify(s)
}

func (s FlowEndpoint) GoString() string {
	return s.String()
}

func (s *FlowEndpoint) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *FlowEndpoint) GetDescription() *string {
	return s.Description
}

func (s *FlowEndpoint) GetDisablePublicNetworkAccess() *bool {
	return s.DisablePublicNetworkAccess
}

func (s *FlowEndpoint) GetFlowEndpointArn() *string {
	return s.FlowEndpointArn
}

func (s *FlowEndpoint) GetFlowEndpointId() *string {
	return s.FlowEndpointId
}

func (s *FlowEndpoint) GetFlowEndpointName() *string {
	return s.FlowEndpointName
}

func (s *FlowEndpoint) GetFlowId() *string {
	return s.FlowId
}

func (s *FlowEndpoint) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *FlowEndpoint) GetRoutingConfiguration() []*FlowEndpointRoutingConfig {
	return s.RoutingConfiguration
}

func (s *FlowEndpoint) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *FlowEndpoint) SetCreatedAt(v string) *FlowEndpoint {
	s.CreatedAt = &v
	return s
}

func (s *FlowEndpoint) SetDescription(v string) *FlowEndpoint {
	s.Description = &v
	return s
}

func (s *FlowEndpoint) SetDisablePublicNetworkAccess(v bool) *FlowEndpoint {
	s.DisablePublicNetworkAccess = &v
	return s
}

func (s *FlowEndpoint) SetFlowEndpointArn(v string) *FlowEndpoint {
	s.FlowEndpointArn = &v
	return s
}

func (s *FlowEndpoint) SetFlowEndpointId(v string) *FlowEndpoint {
	s.FlowEndpointId = &v
	return s
}

func (s *FlowEndpoint) SetFlowEndpointName(v string) *FlowEndpoint {
	s.FlowEndpointName = &v
	return s
}

func (s *FlowEndpoint) SetFlowId(v string) *FlowEndpoint {
	s.FlowId = &v
	return s
}

func (s *FlowEndpoint) SetLastUpdatedAt(v string) *FlowEndpoint {
	s.LastUpdatedAt = &v
	return s
}

func (s *FlowEndpoint) SetRoutingConfiguration(v []*FlowEndpointRoutingConfig) *FlowEndpoint {
	s.RoutingConfiguration = v
	return s
}

func (s *FlowEndpoint) SetTargetVersion(v string) *FlowEndpoint {
	s.TargetVersion = &v
	return s
}

func (s *FlowEndpoint) Validate() error {
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
