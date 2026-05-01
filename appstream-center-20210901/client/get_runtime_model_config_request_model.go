// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuntimeModelConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *GetRuntimeModelConfigRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *GetRuntimeModelConfigRequest
	GetAgentProvider() *string
	SetIncludeRiskInfo(v bool) *GetRuntimeModelConfigRequest
	GetIncludeRiskInfo() *bool
	SetRuntimeId(v string) *GetRuntimeModelConfigRequest
	GetRuntimeId() *string
	SetRuntimeType(v string) *GetRuntimeModelConfigRequest
	GetRuntimeType() *string
}

type GetRuntimeModelConfigRequest struct {
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// example:
	//
	// true
	IncludeRiskInfo *bool `json:"IncludeRiskInfo,omitempty" xml:"IncludeRiskInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecd-xxxx
	RuntimeId *string `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CloudDesktop
	RuntimeType *string `json:"RuntimeType,omitempty" xml:"RuntimeType,omitempty"`
}

func (s GetRuntimeModelConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRuntimeModelConfigRequest) GoString() string {
	return s.String()
}

func (s *GetRuntimeModelConfigRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *GetRuntimeModelConfigRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *GetRuntimeModelConfigRequest) GetIncludeRiskInfo() *bool {
	return s.IncludeRiskInfo
}

func (s *GetRuntimeModelConfigRequest) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *GetRuntimeModelConfigRequest) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *GetRuntimeModelConfigRequest) SetAgentPlatform(v string) *GetRuntimeModelConfigRequest {
	s.AgentPlatform = &v
	return s
}

func (s *GetRuntimeModelConfigRequest) SetAgentProvider(v string) *GetRuntimeModelConfigRequest {
	s.AgentProvider = &v
	return s
}

func (s *GetRuntimeModelConfigRequest) SetIncludeRiskInfo(v bool) *GetRuntimeModelConfigRequest {
	s.IncludeRiskInfo = &v
	return s
}

func (s *GetRuntimeModelConfigRequest) SetRuntimeId(v string) *GetRuntimeModelConfigRequest {
	s.RuntimeId = &v
	return s
}

func (s *GetRuntimeModelConfigRequest) SetRuntimeType(v string) *GetRuntimeModelConfigRequest {
	s.RuntimeType = &v
	return s
}

func (s *GetRuntimeModelConfigRequest) Validate() error {
	return dara.Validate(s)
}
