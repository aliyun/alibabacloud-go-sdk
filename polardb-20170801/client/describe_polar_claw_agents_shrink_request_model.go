// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawAgentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentListShrink(v string) *DescribePolarClawAgentsShrinkRequest
	GetAgentListShrink() *string
	SetApplicationId(v string) *DescribePolarClawAgentsShrinkRequest
	GetApplicationId() *string
}

type DescribePolarClawAgentsShrinkRequest struct {
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

func (s DescribePolarClawAgentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentsShrinkRequest) GetAgentListShrink() *string {
	return s.AgentListShrink
}

func (s *DescribePolarClawAgentsShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawAgentsShrinkRequest) SetAgentListShrink(v string) *DescribePolarClawAgentsShrinkRequest {
	s.AgentListShrink = &v
	return s
}

func (s *DescribePolarClawAgentsShrinkRequest) SetApplicationId(v string) *DescribePolarClawAgentsShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawAgentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
