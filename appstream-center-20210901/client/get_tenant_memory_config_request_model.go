// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTenantMemoryConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentInstanceId(v string) *GetTenantMemoryConfigRequest
	GetAgentInstanceId() *string
	SetAgentPlatform(v string) *GetTenantMemoryConfigRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *GetTenantMemoryConfigRequest
	GetAgentProvider() *string
	SetBizType(v int32) *GetTenantMemoryConfigRequest
	GetBizType() *int32
}

type GetTenantMemoryConfigRequest struct {
	// This parameter is not enabled. Do not specify this parameter. The current API operation queries tenant configurations by provider, business platform, and business type. The example is used to illustrate the field status.
	//
	// example:
	//
	// N/A (not enabled)
	AgentInstanceId *string `json:"AgentInstanceId,omitempty" xml:"AgentInstanceId,omitempty"`
	// The business platform to which the agent belongs. Specify the platform corresponding to the tenant memory configuration, such as `ENTERPRISE` for the enterprise platform.
	//
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The agent provider. This parameter is used together with AgentPlatform and BizType to locate the memory configuration. You must use a configured business combination.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The business type code. This parameter is used together with the provider and business platform. Common values include `1` for Enterprise Edition, `10` for Personal Edition, and `20` for Commercial Edition.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s GetTenantMemoryConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTenantMemoryConfigRequest) GoString() string {
	return s.String()
}

func (s *GetTenantMemoryConfigRequest) GetAgentInstanceId() *string {
	return s.AgentInstanceId
}

func (s *GetTenantMemoryConfigRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *GetTenantMemoryConfigRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *GetTenantMemoryConfigRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *GetTenantMemoryConfigRequest) SetAgentInstanceId(v string) *GetTenantMemoryConfigRequest {
	s.AgentInstanceId = &v
	return s
}

func (s *GetTenantMemoryConfigRequest) SetAgentPlatform(v string) *GetTenantMemoryConfigRequest {
	s.AgentPlatform = &v
	return s
}

func (s *GetTenantMemoryConfigRequest) SetAgentProvider(v string) *GetTenantMemoryConfigRequest {
	s.AgentProvider = &v
	return s
}

func (s *GetTenantMemoryConfigRequest) SetBizType(v int32) *GetTenantMemoryConfigRequest {
	s.BizType = &v
	return s
}

func (s *GetTenantMemoryConfigRequest) Validate() error {
	return dara.Validate(s)
}
