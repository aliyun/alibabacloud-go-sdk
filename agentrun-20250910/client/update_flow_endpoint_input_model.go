// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowEndpointInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateFlowEndpointInput
	GetDescription() *string
	SetFlowEndpointName(v string) *UpdateFlowEndpointInput
	GetFlowEndpointName() *string
	SetRoutingConfiguration(v []*FlowEndpointRoutingConfig) *UpdateFlowEndpointInput
	GetRoutingConfiguration() []*FlowEndpointRoutingConfig
	SetTargetVersion(v string) *UpdateFlowEndpointInput
	GetTargetVersion() *string
}

type UpdateFlowEndpointInput struct {
	// 工作流端点的描述信息
	//
	// example:
	//
	// Production endpoint for flow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 工作流端点的唯一标识名称
	//
	// example:
	//
	// my-flow-endpoint
	FlowEndpointName *string `json:"flowEndpointName,omitempty" xml:"flowEndpointName,omitempty"`
	// 工作流端点的版本路由配置，用于流量分配
	//
	// example:
	//
	// []
	RoutingConfiguration []*FlowEndpointRoutingConfig `json:"routingConfiguration" xml:"routingConfiguration" type:"Repeated"`
	// 工作流端点指向的目标版本号
	//
	// example:
	//
	// 1
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
}

func (s UpdateFlowEndpointInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowEndpointInput) GoString() string {
	return s.String()
}

func (s *UpdateFlowEndpointInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateFlowEndpointInput) GetFlowEndpointName() *string {
	return s.FlowEndpointName
}

func (s *UpdateFlowEndpointInput) GetRoutingConfiguration() []*FlowEndpointRoutingConfig {
	return s.RoutingConfiguration
}

func (s *UpdateFlowEndpointInput) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *UpdateFlowEndpointInput) SetDescription(v string) *UpdateFlowEndpointInput {
	s.Description = &v
	return s
}

func (s *UpdateFlowEndpointInput) SetFlowEndpointName(v string) *UpdateFlowEndpointInput {
	s.FlowEndpointName = &v
	return s
}

func (s *UpdateFlowEndpointInput) SetRoutingConfiguration(v []*FlowEndpointRoutingConfig) *UpdateFlowEndpointInput {
	s.RoutingConfiguration = v
	return s
}

func (s *UpdateFlowEndpointInput) SetTargetVersion(v string) *UpdateFlowEndpointInput {
	s.TargetVersion = &v
	return s
}

func (s *UpdateFlowEndpointInput) Validate() error {
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
