// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowAliasShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateFlowAliasShrinkRequest
	GetDescription() *string
	SetFlowName(v string) *CreateFlowAliasShrinkRequest
	GetFlowName() *string
	SetName(v string) *CreateFlowAliasShrinkRequest
	GetName() *string
	SetRoutingConfigurationsShrink(v string) *CreateFlowAliasShrinkRequest
	GetRoutingConfigurationsShrink() *string
}

type CreateFlowAliasShrinkRequest struct {
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
	RoutingConfigurationsShrink *string `json:"RoutingConfigurations,omitempty" xml:"RoutingConfigurations,omitempty"`
}

func (s CreateFlowAliasShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowAliasShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowAliasShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFlowAliasShrinkRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *CreateFlowAliasShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateFlowAliasShrinkRequest) GetRoutingConfigurationsShrink() *string {
	return s.RoutingConfigurationsShrink
}

func (s *CreateFlowAliasShrinkRequest) SetDescription(v string) *CreateFlowAliasShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowAliasShrinkRequest) SetFlowName(v string) *CreateFlowAliasShrinkRequest {
	s.FlowName = &v
	return s
}

func (s *CreateFlowAliasShrinkRequest) SetName(v string) *CreateFlowAliasShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowAliasShrinkRequest) SetRoutingConfigurationsShrink(v string) *CreateFlowAliasShrinkRequest {
	s.RoutingConfigurationsShrink = &v
	return s
}

func (s *CreateFlowAliasShrinkRequest) Validate() error {
	return dara.Validate(s)
}
