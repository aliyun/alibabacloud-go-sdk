// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentList(v []*string) *DescribePolarClawAgentsRequest
	GetAgentList() []*string
	SetApplicationId(v string) *DescribePolarClawAgentsRequest
	GetApplicationId() *string
}

type DescribePolarClawAgentsRequest struct {
	// example:
	//
	// work,research
	AgentList []*string `json:"AgentList,omitempty" xml:"AgentList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s DescribePolarClawAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentsRequest) GetAgentList() []*string {
	return s.AgentList
}

func (s *DescribePolarClawAgentsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawAgentsRequest) SetAgentList(v []*string) *DescribePolarClawAgentsRequest {
	s.AgentList = v
	return s
}

func (s *DescribePolarClawAgentsRequest) SetApplicationId(v string) *DescribePolarClawAgentsRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawAgentsRequest) Validate() error {
	return dara.Validate(s)
}
