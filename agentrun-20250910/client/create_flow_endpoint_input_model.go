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
	SetFlowEndpointName(v string) *CreateFlowEndpointInput
	GetFlowEndpointName() *string
	SetRoutingConfiguration(v []*FlowEndpointRoutingConfig) *CreateFlowEndpointInput
	GetRoutingConfiguration() []*FlowEndpointRoutingConfig
	SetTargetVersion(v string) *CreateFlowEndpointInput
	GetTargetVersion() *string
}

type CreateFlowEndpointInput struct {
	// 工作流端点的描述信息
	//
	// example:
	//
	// Production endpoint for flow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 工作流端点的唯一标识名称
	//
	// This parameter is required.
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

func (s CreateFlowEndpointInput) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowEndpointInput) GoString() string {
	return s.String()
}

func (s *CreateFlowEndpointInput) GetDescription() *string {
	return s.Description
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
