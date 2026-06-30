// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawAgentToolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DescribePolarClawAgentToolsRequest
	GetAgentId() *string
	SetApplicationId(v string) *DescribePolarClawAgentToolsRequest
	GetApplicationId() *string
	SetIncludePlugins(v bool) *DescribePolarClawAgentToolsRequest
	GetIncludePlugins() *bool
}

type DescribePolarClawAgentToolsRequest struct {
	// Agent ID
	//
	// This parameter is required.
	//
	// example:
	//
	// main
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Specifies whether to include plugin tools.
	//
	// example:
	//
	// true
	IncludePlugins *bool `json:"IncludePlugins,omitempty" xml:"IncludePlugins,omitempty"`
}

func (s DescribePolarClawAgentToolsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentToolsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentToolsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *DescribePolarClawAgentToolsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawAgentToolsRequest) GetIncludePlugins() *bool {
	return s.IncludePlugins
}

func (s *DescribePolarClawAgentToolsRequest) SetAgentId(v string) *DescribePolarClawAgentToolsRequest {
	s.AgentId = &v
	return s
}

func (s *DescribePolarClawAgentToolsRequest) SetApplicationId(v string) *DescribePolarClawAgentToolsRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawAgentToolsRequest) SetIncludePlugins(v bool) *DescribePolarClawAgentToolsRequest {
	s.IncludePlugins = &v
	return s
}

func (s *DescribePolarClawAgentToolsRequest) Validate() error {
	return dara.Validate(s)
}
