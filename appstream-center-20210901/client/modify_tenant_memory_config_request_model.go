// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTenantMemoryConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *ModifyTenantMemoryConfigRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *ModifyTenantMemoryConfigRequest
	GetAgentProvider() *string
	SetBizType(v int32) *ModifyTenantMemoryConfigRequest
	GetBizType() *int32
	SetEnabled(v bool) *ModifyTenantMemoryConfigRequest
	GetEnabled() *bool
}

type ModifyTenantMemoryConfigRequest struct {
	// The business platform to which the agent belongs. Specify the platform that corresponds to the tenant memory configuration, such as the enterprise platform `ENTERPRISE`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The agent provider. Used together with AgentPlatform and BizType to locate the memory configuration. A configured business combination must be used.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The business type code. Used together with the provider and business platform. Common values include `1` (Enterprise Edition), `10` (Personal Edition), and `20` (Commercial Edition).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Specifies whether to enable the memory feature for the business combination. A value of `true` enables the feature, and `false` disables it. If this parameter is omitted, the current status is retained.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s ModifyTenantMemoryConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTenantMemoryConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantMemoryConfigRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *ModifyTenantMemoryConfigRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *ModifyTenantMemoryConfigRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ModifyTenantMemoryConfigRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyTenantMemoryConfigRequest) SetAgentPlatform(v string) *ModifyTenantMemoryConfigRequest {
	s.AgentPlatform = &v
	return s
}

func (s *ModifyTenantMemoryConfigRequest) SetAgentProvider(v string) *ModifyTenantMemoryConfigRequest {
	s.AgentProvider = &v
	return s
}

func (s *ModifyTenantMemoryConfigRequest) SetBizType(v int32) *ModifyTenantMemoryConfigRequest {
	s.BizType = &v
	return s
}

func (s *ModifyTenantMemoryConfigRequest) SetEnabled(v bool) *ModifyTenantMemoryConfigRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyTenantMemoryConfigRequest) Validate() error {
	return dara.Validate(s)
}
