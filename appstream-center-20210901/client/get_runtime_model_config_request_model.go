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
	SetLanguage(v string) *GetRuntimeModelConfigRequest
	GetLanguage() *string
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
	// The model display language. Valid values: zh-CN, en-US, ja-JP, and zh-TW. The value is case-insensitive, and leading and trailing spaces are removed. If this parameter is not specified, left empty, or set to an unsupported value, the default value zh-CN is used. The model pool description mapping is returned based on the normalized interface language, and falls back to Chinese if a translation is missing. Bailian and existing dynamic configurations still use Chinese for zh-CN and zh-TW, and English for en-US and ja-JP. This parameter is used for the names and descriptions of system intelligence tiers, as well as general model descriptions. Template fallback text, custom names, and runtime configurations retain their original values. The values zh_CN, en_US, ja_JP, and zh_TW are also supported, and en and ja are compatible. The value is case-insensitive, and leading and trailing spaces are removed. An invalid language value falls back to Chinese and does not cause a parameter error to be returned for Language. The names and descriptions of the four system intelligence tiers (Auto, Standard, Advanced, and Ultimate) support all four languages, and Chinese is retained for missing Japanese and Traditional Chinese items. User template names retain their original values.
	//
	// example:
	//
	// en_US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The agent runtime ID. The ID mappings are as follows:
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
	// The agent runtime type.
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
	// - Not specified: Returns both smart and non-smart models in the effective configuration.
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

func (s *GetRuntimeModelConfigRequest) GetLanguage() *string {
	return s.Language
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

func (s *GetRuntimeModelConfigRequest) SetLanguage(v string) *GetRuntimeModelConfigRequest {
	s.Language = &v
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
