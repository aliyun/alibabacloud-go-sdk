// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRuntimeStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *GetAgentRuntimeStatisticRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *GetAgentRuntimeStatisticRequest
	GetAgentProvider() *string
	SetBizType(v int32) *GetAgentRuntimeStatisticRequest
	GetBizType() *int32
	SetRuntimeType(v string) *GetAgentRuntimeStatisticRequest
	GetRuntimeType() *string
}

type GetAgentRuntimeStatisticRequest struct {
	// The agent platform used to filter the statistics scope, such as ENTERPRISE_AGENTIC_COMPUTER. If you do not specify this parameter, statistics for all agent platforms are returned. If you specify an unsupported value, a parameter error is returned.
	//
	// example:
	//
	// ENTERPRISE_AGENTIC_COMPUTER
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The agent provider used to filter the statistics scope, such as AgenticComputer. If you do not specify this parameter, statistics for all agent providers are returned. If you specify an unsupported value, a parameter error is returned.
	//
	// example:
	//
	// AgenticComputer
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The business type used to filter runtimes of a specific business type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The runtime environment type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudDesktop
	RuntimeType *string `json:"RuntimeType,omitempty" xml:"RuntimeType,omitempty"`
}

func (s GetAgentRuntimeStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRuntimeStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRuntimeStatisticRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *GetAgentRuntimeStatisticRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *GetAgentRuntimeStatisticRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *GetAgentRuntimeStatisticRequest) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *GetAgentRuntimeStatisticRequest) SetAgentPlatform(v string) *GetAgentRuntimeStatisticRequest {
	s.AgentPlatform = &v
	return s
}

func (s *GetAgentRuntimeStatisticRequest) SetAgentProvider(v string) *GetAgentRuntimeStatisticRequest {
	s.AgentProvider = &v
	return s
}

func (s *GetAgentRuntimeStatisticRequest) SetBizType(v int32) *GetAgentRuntimeStatisticRequest {
	s.BizType = &v
	return s
}

func (s *GetAgentRuntimeStatisticRequest) SetRuntimeType(v string) *GetAgentRuntimeStatisticRequest {
	s.RuntimeType = &v
	return s
}

func (s *GetAgentRuntimeStatisticRequest) Validate() error {
	return dara.Validate(s)
}
