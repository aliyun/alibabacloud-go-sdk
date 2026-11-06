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
	// The agent platform, which is used to filter the statistical scope, such as ENTERPRISE_AGENTIC_COMPUTER. If this parameter is not specified, all agent platforms are included in the statistics. If an unsupported value is specified, a parameter error is returned.
	//
	// example:
	//
	// ENTERPRISE_AGENTIC_COMPUTER
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The agent provider, which is used to filter the statistical scope, such as AgenticComputer. If this parameter is not specified, all agent providers are included in the statistics. If an unsupported value is specified, a parameter error is returned.
	//
	// example:
	//
	// AgenticComputer
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The business type, which is used to filter runtimes of a specified business type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The runtime environment type of the runtime resource.
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
