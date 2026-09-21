// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProviderEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *ListModelProviderEndpointsRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *ListModelProviderEndpointsRequest
	GetAgentProvider() *string
	SetBizType(v int32) *ListModelProviderEndpointsRequest
	GetBizType() *int32
	SetProviderName(v string) *ListModelProviderEndpointsRequest
	GetProviderName() *string
}

type ListModelProviderEndpointsRequest struct {
	// The agent platform name. This parameter is not processed and is only passed through as a redundant field, such as ENTERPRISE.
	//
	// example:
	//
	// ENTERPRISE_AGENTIC_COMPUTER
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The agent provider name, such as HermesAgent or OpenClaw. If this parameter is specified, the providerName in the returned endpoints is the alias from the perspective of the specified agent provider.
	//
	// example:
	//
	// AgenticComputer
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The business type.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The name of the model provider, such as bailian, moonshot, or siliconflow. If this parameter is not specified, the endpoint configurations of all managed providers are returned.
	//
	// example:
	//
	// bailian
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s ListModelProviderEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListModelProviderEndpointsRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *ListModelProviderEndpointsRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *ListModelProviderEndpointsRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ListModelProviderEndpointsRequest) GetProviderName() *string {
	return s.ProviderName
}

func (s *ListModelProviderEndpointsRequest) SetAgentPlatform(v string) *ListModelProviderEndpointsRequest {
	s.AgentPlatform = &v
	return s
}

func (s *ListModelProviderEndpointsRequest) SetAgentProvider(v string) *ListModelProviderEndpointsRequest {
	s.AgentProvider = &v
	return s
}

func (s *ListModelProviderEndpointsRequest) SetBizType(v int32) *ListModelProviderEndpointsRequest {
	s.BizType = &v
	return s
}

func (s *ListModelProviderEndpointsRequest) SetProviderName(v string) *ListModelProviderEndpointsRequest {
	s.ProviderName = &v
	return s
}

func (s *ListModelProviderEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
