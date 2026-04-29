// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolarClawBindingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentListShrink(v string) *ListPolarClawBindingsShrinkRequest
	GetAgentListShrink() *string
	SetApplicationId(v string) *ListPolarClawBindingsShrinkRequest
	GetApplicationId() *string
}

type ListPolarClawBindingsShrinkRequest struct {
	// example:
	//
	// work,research
	AgentListShrink *string `json:"AgentList,omitempty" xml:"AgentList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s ListPolarClawBindingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolarClawBindingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPolarClawBindingsShrinkRequest) GetAgentListShrink() *string {
	return s.AgentListShrink
}

func (s *ListPolarClawBindingsShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListPolarClawBindingsShrinkRequest) SetAgentListShrink(v string) *ListPolarClawBindingsShrinkRequest {
	s.AgentListShrink = &v
	return s
}

func (s *ListPolarClawBindingsShrinkRequest) SetApplicationId(v string) *ListPolarClawBindingsShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListPolarClawBindingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
