// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowAliasShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateFlowAliasShrinkRequest
	GetDescription() *string
	SetFlowName(v string) *UpdateFlowAliasShrinkRequest
	GetFlowName() *string
	SetName(v string) *UpdateFlowAliasShrinkRequest
	GetName() *string
	SetRoutingConfigurationsShrink(v string) *UpdateFlowAliasShrinkRequest
	GetRoutingConfigurationsShrink() *string
}

type UpdateFlowAliasShrinkRequest struct {
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
	Name                        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RoutingConfigurationsShrink *string `json:"RoutingConfigurations,omitempty" xml:"RoutingConfigurations,omitempty"`
}

func (s UpdateFlowAliasShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowAliasShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowAliasShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateFlowAliasShrinkRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *UpdateFlowAliasShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateFlowAliasShrinkRequest) GetRoutingConfigurationsShrink() *string {
	return s.RoutingConfigurationsShrink
}

func (s *UpdateFlowAliasShrinkRequest) SetDescription(v string) *UpdateFlowAliasShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateFlowAliasShrinkRequest) SetFlowName(v string) *UpdateFlowAliasShrinkRequest {
	s.FlowName = &v
	return s
}

func (s *UpdateFlowAliasShrinkRequest) SetName(v string) *UpdateFlowAliasShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateFlowAliasShrinkRequest) SetRoutingConfigurationsShrink(v string) *UpdateFlowAliasShrinkRequest {
	s.RoutingConfigurationsShrink = &v
	return s
}

func (s *UpdateFlowAliasShrinkRequest) Validate() error {
	return dara.Validate(s)
}
