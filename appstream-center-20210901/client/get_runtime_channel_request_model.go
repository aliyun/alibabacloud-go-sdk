// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuntimeChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *GetRuntimeChannelRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *GetRuntimeChannelRequest
	GetAgentProvider() *string
	SetIncludeRiskInfo(v bool) *GetRuntimeChannelRequest
	GetIncludeRiskInfo() *bool
	SetRuntimeId(v string) *GetRuntimeChannelRequest
	GetRuntimeId() *string
	SetRuntimeType(v string) *GetRuntimeChannelRequest
	GetRuntimeType() *string
}

type GetRuntimeChannelRequest struct {
	// The Agent platform.
	//
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The Agent provider.
	//
	// - JVS Computer: Set to OpenClaw.
	//
	// - OpenClaw: Set to OpenClaw.
	//
	// - Hermes Agent: Set to HermesAgent.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// Specifies whether to include risk information.
	//
	// example:
	//
	// true
	IncludeRiskInfo *bool `json:"IncludeRiskInfo,omitempty" xml:"IncludeRiskInfo,omitempty"`
	// The Agent runtime ID. The ID mapping is as follows:
	//
	// - JVS Computer: JVS Computer ID, in the format of jvs-xxxx.
	//
	// - OpenClaw: Cloud computer ID, in the format of ecd-xxxx.
	//
	// - Hermes Agent: Hermes Agent ID, in the format of jvs-xxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// jvs-xxxx
	RuntimeId *string `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// The Agent runtime type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudDesktop
	RuntimeType *string `json:"RuntimeType,omitempty" xml:"RuntimeType,omitempty"`
}

func (s GetRuntimeChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRuntimeChannelRequest) GoString() string {
	return s.String()
}

func (s *GetRuntimeChannelRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *GetRuntimeChannelRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *GetRuntimeChannelRequest) GetIncludeRiskInfo() *bool {
	return s.IncludeRiskInfo
}

func (s *GetRuntimeChannelRequest) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *GetRuntimeChannelRequest) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *GetRuntimeChannelRequest) SetAgentPlatform(v string) *GetRuntimeChannelRequest {
	s.AgentPlatform = &v
	return s
}

func (s *GetRuntimeChannelRequest) SetAgentProvider(v string) *GetRuntimeChannelRequest {
	s.AgentProvider = &v
	return s
}

func (s *GetRuntimeChannelRequest) SetIncludeRiskInfo(v bool) *GetRuntimeChannelRequest {
	s.IncludeRiskInfo = &v
	return s
}

func (s *GetRuntimeChannelRequest) SetRuntimeId(v string) *GetRuntimeChannelRequest {
	s.RuntimeId = &v
	return s
}

func (s *GetRuntimeChannelRequest) SetRuntimeType(v string) *GetRuntimeChannelRequest {
	s.RuntimeType = &v
	return s
}

func (s *GetRuntimeChannelRequest) Validate() error {
	return dara.Validate(s)
}
