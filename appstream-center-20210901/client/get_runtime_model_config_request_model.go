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
	SetSmartModel(v bool) *GetRuntimeModelConfigRequest
	GetSmartModel() *bool
}

type GetRuntimeModelConfigRequest struct {
	// The agent platform.
	//
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The agent provider.
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
	// The ID of the agent runtime. The ID mappings are as follows:
	//
	// JVS Computer: JVS Computer ID, in the format of jvs-xxxx.
	//
	// OpenClaw: Cloud desktop ID, in the format of ecd-xxxx.
	//
	// Hermes Agent: Hermes Agent ID, in the format of jvs-xxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-xxxx
	RuntimeId *string `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// The type of the agent runtime.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudDesktop
	RuntimeType *string `json:"RuntimeType,omitempty" xml:"RuntimeType,omitempty"`
	// The smart model filter. Valid values:
	//
	// - true: Queries the system wuying-preset provider models based on the business type of the runtime.
	//
	// - false: Excludes wuying-preset from the effective configuration.
	//
	// - Empty: Returns both smart and non-smart models in the effective configuration.
	//
	// The business type is read by the server from runtime.biz_type, and the BizType input parameter is not required. Risk information is still independently controlled by IncludeRiskInfo.
	//
	// example:
	//
	// true
	SmartModel *bool `json:"SmartModel,omitempty" xml:"SmartModel,omitempty"`
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

func (s *GetRuntimeModelConfigRequest) GetSmartModel() *bool {
	return s.SmartModel
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

func (s *GetRuntimeModelConfigRequest) SetSmartModel(v bool) *GetRuntimeModelConfigRequest {
	s.SmartModel = &v
	return s
}

func (s *GetRuntimeModelConfigRequest) Validate() error {
	return dara.Validate(s)
}
