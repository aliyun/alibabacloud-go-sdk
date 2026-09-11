// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateApplicationVersionRequest
	GetApplicationId() *string
	SetBusinessUnitId(v string) *UpdateApplicationVersionRequest
	GetBusinessUnitId() *string
	SetInteractionConfig(v *UpdateApplicationVersionRequestInteractionConfig) *UpdateApplicationVersionRequest
	GetInteractionConfig() *UpdateApplicationVersionRequestInteractionConfig
	SetLabelConfig(v []*UpdateApplicationVersionRequestLabelConfig) *UpdateApplicationVersionRequest
	GetLabelConfig() []*UpdateApplicationVersionRequestLabelConfig
	SetRagConfig(v *UpdateApplicationVersionRequestRagConfig) *UpdateApplicationVersionRequest
	GetRagConfig() *UpdateApplicationVersionRequestRagConfig
	SetScriptProfile(v *UpdateApplicationVersionRequestScriptProfile) *UpdateApplicationVersionRequest
	GetScriptProfile() *UpdateApplicationVersionRequestScriptProfile
	SetSynthesizerConfig(v *UpdateApplicationVersionRequestSynthesizerConfig) *UpdateApplicationVersionRequest
	GetSynthesizerConfig() *UpdateApplicationVersionRequestSynthesizerConfig
	SetToolConfig(v *UpdateApplicationVersionRequestToolConfig) *UpdateApplicationVersionRequest
	GetToolConfig() *UpdateApplicationVersionRequestToolConfig
	SetTranscriberConfig(v *UpdateApplicationVersionRequestTranscriberConfig) *UpdateApplicationVersionRequest
	GetTranscriberConfig() *UpdateApplicationVersionRequestTranscriberConfig
	SetVersionId(v string) *UpdateApplicationVersionRequest
	GetVersionId() *string
}

type UpdateApplicationVersionRequest struct {
	// This parameter is required.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	BusinessUnitId    *string                                           `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	InteractionConfig *UpdateApplicationVersionRequestInteractionConfig `json:"InteractionConfig,omitempty" xml:"InteractionConfig,omitempty" type:"Struct"`
	LabelConfig       []*UpdateApplicationVersionRequestLabelConfig     `json:"LabelConfig,omitempty" xml:"LabelConfig,omitempty" type:"Repeated"`
	RagConfig         *UpdateApplicationVersionRequestRagConfig         `json:"RagConfig,omitempty" xml:"RagConfig,omitempty" type:"Struct"`
	// This parameter is required.
	ScriptProfile *UpdateApplicationVersionRequestScriptProfile `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty" type:"Struct"`
	// if can be null:
	// true
	SynthesizerConfig *UpdateApplicationVersionRequestSynthesizerConfig `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty" type:"Struct"`
	ToolConfig        *UpdateApplicationVersionRequestToolConfig        `json:"ToolConfig,omitempty" xml:"ToolConfig,omitempty" type:"Struct"`
	// if can be null:
	// true
	TranscriberConfig *UpdateApplicationVersionRequestTranscriberConfig `json:"TranscriberConfig,omitempty" xml:"TranscriberConfig,omitempty" type:"Struct"`
	// This parameter is required.
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s UpdateApplicationVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationVersionRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *UpdateApplicationVersionRequest) GetInteractionConfig() *UpdateApplicationVersionRequestInteractionConfig {
	return s.InteractionConfig
}

func (s *UpdateApplicationVersionRequest) GetLabelConfig() []*UpdateApplicationVersionRequestLabelConfig {
	return s.LabelConfig
}

func (s *UpdateApplicationVersionRequest) GetRagConfig() *UpdateApplicationVersionRequestRagConfig {
	return s.RagConfig
}

func (s *UpdateApplicationVersionRequest) GetScriptProfile() *UpdateApplicationVersionRequestScriptProfile {
	return s.ScriptProfile
}

func (s *UpdateApplicationVersionRequest) GetSynthesizerConfig() *UpdateApplicationVersionRequestSynthesizerConfig {
	return s.SynthesizerConfig
}

func (s *UpdateApplicationVersionRequest) GetToolConfig() *UpdateApplicationVersionRequestToolConfig {
	return s.ToolConfig
}

func (s *UpdateApplicationVersionRequest) GetTranscriberConfig() *UpdateApplicationVersionRequestTranscriberConfig {
	return s.TranscriberConfig
}

func (s *UpdateApplicationVersionRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *UpdateApplicationVersionRequest) SetApplicationId(v string) *UpdateApplicationVersionRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationVersionRequest) SetBusinessUnitId(v string) *UpdateApplicationVersionRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *UpdateApplicationVersionRequest) SetInteractionConfig(v *UpdateApplicationVersionRequestInteractionConfig) *UpdateApplicationVersionRequest {
	s.InteractionConfig = v
	return s
}

func (s *UpdateApplicationVersionRequest) SetLabelConfig(v []*UpdateApplicationVersionRequestLabelConfig) *UpdateApplicationVersionRequest {
	s.LabelConfig = v
	return s
}

func (s *UpdateApplicationVersionRequest) SetRagConfig(v *UpdateApplicationVersionRequestRagConfig) *UpdateApplicationVersionRequest {
	s.RagConfig = v
	return s
}

func (s *UpdateApplicationVersionRequest) SetScriptProfile(v *UpdateApplicationVersionRequestScriptProfile) *UpdateApplicationVersionRequest {
	s.ScriptProfile = v
	return s
}

func (s *UpdateApplicationVersionRequest) SetSynthesizerConfig(v *UpdateApplicationVersionRequestSynthesizerConfig) *UpdateApplicationVersionRequest {
	s.SynthesizerConfig = v
	return s
}

func (s *UpdateApplicationVersionRequest) SetToolConfig(v *UpdateApplicationVersionRequestToolConfig) *UpdateApplicationVersionRequest {
	s.ToolConfig = v
	return s
}

func (s *UpdateApplicationVersionRequest) SetTranscriberConfig(v *UpdateApplicationVersionRequestTranscriberConfig) *UpdateApplicationVersionRequest {
	s.TranscriberConfig = v
	return s
}

func (s *UpdateApplicationVersionRequest) SetVersionId(v string) *UpdateApplicationVersionRequest {
	s.VersionId = &v
	return s
}

func (s *UpdateApplicationVersionRequest) Validate() error {
	if s.InteractionConfig != nil {
		if err := s.InteractionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LabelConfig != nil {
		for _, item := range s.LabelConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RagConfig != nil {
		if err := s.RagConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ScriptProfile != nil {
		if err := s.ScriptProfile.Validate(); err != nil {
			return err
		}
	}
	if s.SynthesizerConfig != nil {
		if err := s.SynthesizerConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ToolConfig != nil {
		if err := s.ToolConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TranscriberConfig != nil {
		if err := s.TranscriberConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationVersionRequestInteractionConfig struct {
	BackgroundMusicId                *string                                                                 `json:"BackgroundMusicId,omitempty" xml:"BackgroundMusicId,omitempty"`
	EndConversationConfig            *UpdateApplicationVersionRequestInteractionConfigEndConversationConfig  `json:"EndConversationConfig,omitempty" xml:"EndConversationConfig,omitempty" type:"Struct"`
	InitialGreetingDelayMilliseconds *int32                                                                  `json:"InitialGreetingDelayMilliseconds,omitempty" xml:"InitialGreetingDelayMilliseconds,omitempty"`
	SilenceDetectionConfig           *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig `json:"SilenceDetectionConfig,omitempty" xml:"SilenceDetectionConfig,omitempty" type:"Struct"`
}

func (s UpdateApplicationVersionRequestInteractionConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestInteractionConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestInteractionConfig) GetBackgroundMusicId() *string {
	return s.BackgroundMusicId
}

func (s *UpdateApplicationVersionRequestInteractionConfig) GetEndConversationConfig() *UpdateApplicationVersionRequestInteractionConfigEndConversationConfig {
	return s.EndConversationConfig
}

func (s *UpdateApplicationVersionRequestInteractionConfig) GetInitialGreetingDelayMilliseconds() *int32 {
	return s.InitialGreetingDelayMilliseconds
}

func (s *UpdateApplicationVersionRequestInteractionConfig) GetSilenceDetectionConfig() *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig {
	return s.SilenceDetectionConfig
}

func (s *UpdateApplicationVersionRequestInteractionConfig) SetBackgroundMusicId(v string) *UpdateApplicationVersionRequestInteractionConfig {
	s.BackgroundMusicId = &v
	return s
}

func (s *UpdateApplicationVersionRequestInteractionConfig) SetEndConversationConfig(v *UpdateApplicationVersionRequestInteractionConfigEndConversationConfig) *UpdateApplicationVersionRequestInteractionConfig {
	s.EndConversationConfig = v
	return s
}

func (s *UpdateApplicationVersionRequestInteractionConfig) SetInitialGreetingDelayMilliseconds(v int32) *UpdateApplicationVersionRequestInteractionConfig {
	s.InitialGreetingDelayMilliseconds = &v
	return s
}

func (s *UpdateApplicationVersionRequestInteractionConfig) SetSilenceDetectionConfig(v *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) *UpdateApplicationVersionRequestInteractionConfig {
	s.SilenceDetectionConfig = v
	return s
}

func (s *UpdateApplicationVersionRequestInteractionConfig) Validate() error {
	if s.EndConversationConfig != nil {
		if err := s.EndConversationConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SilenceDetectionConfig != nil {
		if err := s.SilenceDetectionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationVersionRequestInteractionConfigEndConversationConfig struct {
	Delay    *int32                                                                           `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Triggers []*UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s UpdateApplicationVersionRequestInteractionConfigEndConversationConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestInteractionConfigEndConversationConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestInteractionConfigEndConversationConfig) GetDelay() *int32 {
	return s.Delay
}

func (s *UpdateApplicationVersionRequestInteractionConfigEndConversationConfig) GetTriggers() []*UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers {
	return s.Triggers
}

func (s *UpdateApplicationVersionRequestInteractionConfigEndConversationConfig) SetDelay(v int32) *UpdateApplicationVersionRequestInteractionConfigEndConversationConfig {
	s.Delay = &v
	return s
}

func (s *UpdateApplicationVersionRequestInteractionConfigEndConversationConfig) SetTriggers(v []*UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) *UpdateApplicationVersionRequestInteractionConfigEndConversationConfig {
	s.Triggers = v
	return s
}

func (s *UpdateApplicationVersionRequestInteractionConfigEndConversationConfig) Validate() error {
	if s.Triggers != nil {
		for _, item := range s.Triggers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers struct {
	// example:
	//
	// 感谢您的接听，祝您生活愉快，再见!
	ClosingStatement *string   `json:"ClosingStatement,omitempty" xml:"ClosingStatement,omitempty"`
	KeyWords         []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	TriggerType      *string   `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	TurnLimit        *int32    `json:"TurnLimit,omitempty" xml:"TurnLimit,omitempty"`
}

func (s UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) GetClosingStatement() *string {
	return s.ClosingStatement
}

func (s *UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) GetKeyWords() []*string {
	return s.KeyWords
}

func (s *UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) GetTriggerType() *string {
	return s.TriggerType
}

func (s *UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) GetTurnLimit() *int32 {
	return s.TurnLimit
}

func (s *UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) SetClosingStatement(v string) *UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.ClosingStatement = &v
	return s
}

func (s *UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) SetKeyWords(v []*string) *UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.KeyWords = v
	return s
}

func (s *UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) SetTriggerType(v string) *UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.TriggerType = &v
	return s
}

func (s *UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) SetTurnLimit(v int32) *UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.TurnLimit = &v
	return s
}

func (s *UpdateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig struct {
	MaxRepeats *int32 `json:"MaxRepeats,omitempty" xml:"MaxRepeats,omitempty"`
	Timeout    *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) GetMaxRepeats() *int32 {
	return s.MaxRepeats
}

func (s *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) SetMaxRepeats(v int32) *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig {
	s.MaxRepeats = &v
	return s
}

func (s *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) SetTimeout(v int32) *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig {
	s.Timeout = &v
	return s
}

func (s *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationVersionRequestLabelConfig struct {
	CandidateValues []*string `json:"CandidateValues,omitempty" xml:"CandidateValues,omitempty" type:"Repeated"`
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateApplicationVersionRequestLabelConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestLabelConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestLabelConfig) GetCandidateValues() []*string {
	return s.CandidateValues
}

func (s *UpdateApplicationVersionRequestLabelConfig) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationVersionRequestLabelConfig) GetName() *string {
	return s.Name
}

func (s *UpdateApplicationVersionRequestLabelConfig) SetCandidateValues(v []*string) *UpdateApplicationVersionRequestLabelConfig {
	s.CandidateValues = v
	return s
}

func (s *UpdateApplicationVersionRequestLabelConfig) SetDescription(v string) *UpdateApplicationVersionRequestLabelConfig {
	s.Description = &v
	return s
}

func (s *UpdateApplicationVersionRequestLabelConfig) SetName(v string) *UpdateApplicationVersionRequestLabelConfig {
	s.Name = &v
	return s
}

func (s *UpdateApplicationVersionRequestLabelConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationVersionRequestRagConfig struct {
	Enabled          *bool     `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	KnowledgeBaseIds []*string `json:"KnowledgeBaseIds,omitempty" xml:"KnowledgeBaseIds,omitempty" type:"Repeated"`
	MaxContentLength *int32    `json:"MaxContentLength,omitempty" xml:"MaxContentLength,omitempty"`
	RagEngine        *string   `json:"RagEngine,omitempty" xml:"RagEngine,omitempty"`
	TopN             *int32    `json:"TopN,omitempty" xml:"TopN,omitempty"`
}

func (s UpdateApplicationVersionRequestRagConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestRagConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestRagConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateApplicationVersionRequestRagConfig) GetKnowledgeBaseIds() []*string {
	return s.KnowledgeBaseIds
}

func (s *UpdateApplicationVersionRequestRagConfig) GetMaxContentLength() *int32 {
	return s.MaxContentLength
}

func (s *UpdateApplicationVersionRequestRagConfig) GetRagEngine() *string {
	return s.RagEngine
}

func (s *UpdateApplicationVersionRequestRagConfig) GetTopN() *int32 {
	return s.TopN
}

func (s *UpdateApplicationVersionRequestRagConfig) SetEnabled(v bool) *UpdateApplicationVersionRequestRagConfig {
	s.Enabled = &v
	return s
}

func (s *UpdateApplicationVersionRequestRagConfig) SetKnowledgeBaseIds(v []*string) *UpdateApplicationVersionRequestRagConfig {
	s.KnowledgeBaseIds = v
	return s
}

func (s *UpdateApplicationVersionRequestRagConfig) SetMaxContentLength(v int32) *UpdateApplicationVersionRequestRagConfig {
	s.MaxContentLength = &v
	return s
}

func (s *UpdateApplicationVersionRequestRagConfig) SetRagEngine(v string) *UpdateApplicationVersionRequestRagConfig {
	s.RagEngine = &v
	return s
}

func (s *UpdateApplicationVersionRequestRagConfig) SetTopN(v int32) *UpdateApplicationVersionRequestRagConfig {
	s.TopN = &v
	return s
}

func (s *UpdateApplicationVersionRequestRagConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationVersionRequestScriptProfile struct {
	// example:
	//
	// 6f444ecf21d94238b516735916c98666
	AgentKey     *string                                                   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	AgentProfile *UpdateApplicationVersionRequestScriptProfileAgentProfile `json:"AgentProfile,omitempty" xml:"AgentProfile,omitempty" type:"Struct"`
	// example:
	//
	// chatbot-cn-MQuyjjb666
	ChatbotId        *string                                                       `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	FunctionMeta     *UpdateApplicationVersionRequestScriptProfileFunctionMeta     `json:"FunctionMeta,omitempty" xml:"FunctionMeta,omitempty" type:"Struct"`
	Model            *string                                                       `json:"Model,omitempty" xml:"Model,omitempty"`
	NluAccessProfile *UpdateApplicationVersionRequestScriptProfileNluAccessProfile `json:"NluAccessProfile,omitempty" xml:"NluAccessProfile,omitempty" type:"Struct"`
	NluAccessType    *string                                                       `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	OmniModel        *bool                                                         `json:"OmniModel,omitempty" xml:"OmniModel,omitempty"`
}

func (s UpdateApplicationVersionRequestScriptProfile) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestScriptProfile) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestScriptProfile) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateApplicationVersionRequestScriptProfile) GetAgentProfile() *UpdateApplicationVersionRequestScriptProfileAgentProfile {
	return s.AgentProfile
}

func (s *UpdateApplicationVersionRequestScriptProfile) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *UpdateApplicationVersionRequestScriptProfile) GetFunctionMeta() *UpdateApplicationVersionRequestScriptProfileFunctionMeta {
	return s.FunctionMeta
}

func (s *UpdateApplicationVersionRequestScriptProfile) GetModel() *string {
	return s.Model
}

func (s *UpdateApplicationVersionRequestScriptProfile) GetNluAccessProfile() *UpdateApplicationVersionRequestScriptProfileNluAccessProfile {
	return s.NluAccessProfile
}

func (s *UpdateApplicationVersionRequestScriptProfile) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *UpdateApplicationVersionRequestScriptProfile) GetOmniModel() *bool {
	return s.OmniModel
}

func (s *UpdateApplicationVersionRequestScriptProfile) SetAgentKey(v string) *UpdateApplicationVersionRequestScriptProfile {
	s.AgentKey = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfile) SetAgentProfile(v *UpdateApplicationVersionRequestScriptProfileAgentProfile) *UpdateApplicationVersionRequestScriptProfile {
	s.AgentProfile = v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfile) SetChatbotId(v string) *UpdateApplicationVersionRequestScriptProfile {
	s.ChatbotId = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfile) SetFunctionMeta(v *UpdateApplicationVersionRequestScriptProfileFunctionMeta) *UpdateApplicationVersionRequestScriptProfile {
	s.FunctionMeta = v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfile) SetModel(v string) *UpdateApplicationVersionRequestScriptProfile {
	s.Model = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfile) SetNluAccessProfile(v *UpdateApplicationVersionRequestScriptProfileNluAccessProfile) *UpdateApplicationVersionRequestScriptProfile {
	s.NluAccessProfile = v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfile) SetNluAccessType(v string) *UpdateApplicationVersionRequestScriptProfile {
	s.NluAccessType = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfile) SetOmniModel(v bool) *UpdateApplicationVersionRequestScriptProfile {
	s.OmniModel = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfile) Validate() error {
	if s.AgentProfile != nil {
		if err := s.AgentProfile.Validate(); err != nil {
			return err
		}
	}
	if s.FunctionMeta != nil {
		if err := s.FunctionMeta.Validate(); err != nil {
			return err
		}
	}
	if s.NluAccessProfile != nil {
		if err := s.NluAccessProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationVersionRequestScriptProfileAgentProfile struct {
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PromptsJson             *string `json:"PromptsJson,omitempty" xml:"PromptsJson,omitempty"`
	ScriptProfileTemplateId *string `json:"ScriptProfileTemplateId,omitempty" xml:"ScriptProfileTemplateId,omitempty"`
}

func (s UpdateApplicationVersionRequestScriptProfileAgentProfile) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestScriptProfileAgentProfile) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestScriptProfileAgentProfile) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationVersionRequestScriptProfileAgentProfile) GetName() *string {
	return s.Name
}

func (s *UpdateApplicationVersionRequestScriptProfileAgentProfile) GetPromptsJson() *string {
	return s.PromptsJson
}

func (s *UpdateApplicationVersionRequestScriptProfileAgentProfile) GetScriptProfileTemplateId() *string {
	return s.ScriptProfileTemplateId
}

func (s *UpdateApplicationVersionRequestScriptProfileAgentProfile) SetDescription(v string) *UpdateApplicationVersionRequestScriptProfileAgentProfile {
	s.Description = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfileAgentProfile) SetName(v string) *UpdateApplicationVersionRequestScriptProfileAgentProfile {
	s.Name = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfileAgentProfile) SetPromptsJson(v string) *UpdateApplicationVersionRequestScriptProfileAgentProfile {
	s.PromptsJson = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfileAgentProfile) SetScriptProfileTemplateId(v string) *UpdateApplicationVersionRequestScriptProfileAgentProfile {
	s.ScriptProfileTemplateId = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfileAgentProfile) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationVersionRequestScriptProfileFunctionMeta struct {
	// example:
	//
	// 9b752bbb-805a-4d3e-9013-eab5555c3fef
	FunctionId      *string `json:"FunctionId,omitempty" xml:"FunctionId,omitempty"`
	FunctionName    *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	HttpTriggerName *string `json:"HttpTriggerName,omitempty" xml:"HttpTriggerName,omitempty"`
	// example:
	//
	// http://chat-xxxxx-v-yewiundukb.cn-hangzhou-xxx.run
	HttpTriggerUrl *string `json:"HttpTriggerUrl,omitempty" xml:"HttpTriggerUrl,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateApplicationVersionRequestScriptProfileFunctionMeta) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestScriptProfileFunctionMeta) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestScriptProfileFunctionMeta) GetFunctionId() *string {
	return s.FunctionId
}

func (s *UpdateApplicationVersionRequestScriptProfileFunctionMeta) GetFunctionName() *string {
	return s.FunctionName
}

func (s *UpdateApplicationVersionRequestScriptProfileFunctionMeta) GetHttpTriggerName() *string {
	return s.HttpTriggerName
}

func (s *UpdateApplicationVersionRequestScriptProfileFunctionMeta) GetHttpTriggerUrl() *string {
	return s.HttpTriggerUrl
}

func (s *UpdateApplicationVersionRequestScriptProfileFunctionMeta) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateApplicationVersionRequestScriptProfileFunctionMeta) SetFunctionId(v string) *UpdateApplicationVersionRequestScriptProfileFunctionMeta {
	s.FunctionId = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfileFunctionMeta) SetFunctionName(v string) *UpdateApplicationVersionRequestScriptProfileFunctionMeta {
	s.FunctionName = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfileFunctionMeta) SetHttpTriggerName(v string) *UpdateApplicationVersionRequestScriptProfileFunctionMeta {
	s.HttpTriggerName = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfileFunctionMeta) SetHttpTriggerUrl(v string) *UpdateApplicationVersionRequestScriptProfileFunctionMeta {
	s.HttpTriggerUrl = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfileFunctionMeta) SetRegionId(v string) *UpdateApplicationVersionRequestScriptProfileFunctionMeta {
	s.RegionId = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfileFunctionMeta) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationVersionRequestScriptProfileNluAccessProfile struct {
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s UpdateApplicationVersionRequestScriptProfileNluAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestScriptProfileNluAccessProfile) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestScriptProfileNluAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *UpdateApplicationVersionRequestScriptProfileNluAccessProfile) SetAccessProfileId(v string) *UpdateApplicationVersionRequestScriptProfileNluAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfileNluAccessProfile) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationVersionRequestSynthesizerConfig struct {
	Model            *string                                                           `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile *UpdateApplicationVersionRequestSynthesizerConfigNlsAccessProfile `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	NlsAccessType    *string                                                           `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	NlsEngine        *string                                                           `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	PitchRate        *int32                                                            `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	PronRules        []*UpdateApplicationVersionRequestSynthesizerConfigPronRules      `json:"PronRules,omitempty" xml:"PronRules,omitempty" type:"Repeated"`
	SpeechRate       *int32                                                            `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Voice            *string                                                           `json:"Voice,omitempty" xml:"Voice,omitempty"`
	Volume           *int32                                                            `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s UpdateApplicationVersionRequestSynthesizerConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestSynthesizerConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) GetModel() *string {
	return s.Model
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) GetNlsAccessProfile() *UpdateApplicationVersionRequestSynthesizerConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) GetPitchRate() *int32 {
	return s.PitchRate
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) GetPronRules() []*UpdateApplicationVersionRequestSynthesizerConfigPronRules {
	return s.PronRules
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) GetVoice() *string {
	return s.Voice
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) SetModel(v string) *UpdateApplicationVersionRequestSynthesizerConfig {
	s.Model = &v
	return s
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) SetNlsAccessProfile(v *UpdateApplicationVersionRequestSynthesizerConfigNlsAccessProfile) *UpdateApplicationVersionRequestSynthesizerConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) SetNlsAccessType(v string) *UpdateApplicationVersionRequestSynthesizerConfig {
	s.NlsAccessType = &v
	return s
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) SetNlsEngine(v string) *UpdateApplicationVersionRequestSynthesizerConfig {
	s.NlsEngine = &v
	return s
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) SetPitchRate(v int32) *UpdateApplicationVersionRequestSynthesizerConfig {
	s.PitchRate = &v
	return s
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) SetPronRules(v []*UpdateApplicationVersionRequestSynthesizerConfigPronRules) *UpdateApplicationVersionRequestSynthesizerConfig {
	s.PronRules = v
	return s
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) SetSpeechRate(v int32) *UpdateApplicationVersionRequestSynthesizerConfig {
	s.SpeechRate = &v
	return s
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) SetVoice(v string) *UpdateApplicationVersionRequestSynthesizerConfig {
	s.Voice = &v
	return s
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) SetVolume(v int32) *UpdateApplicationVersionRequestSynthesizerConfig {
	s.Volume = &v
	return s
}

func (s *UpdateApplicationVersionRequestSynthesizerConfig) Validate() error {
	if s.NlsAccessProfile != nil {
		if err := s.NlsAccessProfile.Validate(); err != nil {
			return err
		}
	}
	if s.PronRules != nil {
		for _, item := range s.PronRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApplicationVersionRequestSynthesizerConfigNlsAccessProfile struct {
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s UpdateApplicationVersionRequestSynthesizerConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestSynthesizerConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestSynthesizerConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *UpdateApplicationVersionRequestSynthesizerConfigNlsAccessProfile) SetAccessProfileId(v string) *UpdateApplicationVersionRequestSynthesizerConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *UpdateApplicationVersionRequestSynthesizerConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationVersionRequestSynthesizerConfigPronRules struct {
	Pattern     *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
}

func (s UpdateApplicationVersionRequestSynthesizerConfigPronRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestSynthesizerConfigPronRules) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestSynthesizerConfigPronRules) GetPattern() *string {
	return s.Pattern
}

func (s *UpdateApplicationVersionRequestSynthesizerConfigPronRules) GetReplacement() *string {
	return s.Replacement
}

func (s *UpdateApplicationVersionRequestSynthesizerConfigPronRules) SetPattern(v string) *UpdateApplicationVersionRequestSynthesizerConfigPronRules {
	s.Pattern = &v
	return s
}

func (s *UpdateApplicationVersionRequestSynthesizerConfigPronRules) SetReplacement(v string) *UpdateApplicationVersionRequestSynthesizerConfigPronRules {
	s.Replacement = &v
	return s
}

func (s *UpdateApplicationVersionRequestSynthesizerConfigPronRules) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationVersionRequestToolConfig struct {
	McpServers []*UpdateApplicationVersionRequestToolConfigMcpServers `json:"McpServers,omitempty" xml:"McpServers,omitempty" type:"Repeated"`
}

func (s UpdateApplicationVersionRequestToolConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestToolConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestToolConfig) GetMcpServers() []*UpdateApplicationVersionRequestToolConfigMcpServers {
	return s.McpServers
}

func (s *UpdateApplicationVersionRequestToolConfig) SetMcpServers(v []*UpdateApplicationVersionRequestToolConfigMcpServers) *UpdateApplicationVersionRequestToolConfig {
	s.McpServers = v
	return s
}

func (s *UpdateApplicationVersionRequestToolConfig) Validate() error {
	if s.McpServers != nil {
		for _, item := range s.McpServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApplicationVersionRequestToolConfigMcpServers struct {
	BaseUrl     *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SseEndpoint *string `json:"SseEndpoint,omitempty" xml:"SseEndpoint,omitempty"`
}

func (s UpdateApplicationVersionRequestToolConfigMcpServers) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestToolConfigMcpServers) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestToolConfigMcpServers) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *UpdateApplicationVersionRequestToolConfigMcpServers) GetName() *string {
	return s.Name
}

func (s *UpdateApplicationVersionRequestToolConfigMcpServers) GetSseEndpoint() *string {
	return s.SseEndpoint
}

func (s *UpdateApplicationVersionRequestToolConfigMcpServers) SetBaseUrl(v string) *UpdateApplicationVersionRequestToolConfigMcpServers {
	s.BaseUrl = &v
	return s
}

func (s *UpdateApplicationVersionRequestToolConfigMcpServers) SetName(v string) *UpdateApplicationVersionRequestToolConfigMcpServers {
	s.Name = &v
	return s
}

func (s *UpdateApplicationVersionRequestToolConfigMcpServers) SetSseEndpoint(v string) *UpdateApplicationVersionRequestToolConfigMcpServers {
	s.SseEndpoint = &v
	return s
}

func (s *UpdateApplicationVersionRequestToolConfigMcpServers) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationVersionRequestTranscriberConfig struct {
	CorrectionRules      []*UpdateApplicationVersionRequestTranscriberConfigCorrectionRules `json:"CorrectionRules,omitempty" xml:"CorrectionRules,omitempty" type:"Repeated"`
	CustomizationId      *string                                                            `json:"CustomizationId,omitempty" xml:"CustomizationId,omitempty"`
	EndSilenceTimeout    *int32                                                             `json:"EndSilenceTimeout,omitempty" xml:"EndSilenceTimeout,omitempty"`
	Model                *string                                                            `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile     *UpdateApplicationVersionRequestTranscriberConfigNlsAccessProfile  `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	NlsAccessType        *string                                                            `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	NlsEngine            *string                                                            `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	SpeechNoiseThreshold *int32                                                             `json:"SpeechNoiseThreshold,omitempty" xml:"SpeechNoiseThreshold,omitempty"`
	VocabularyId         *string                                                            `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
}

func (s UpdateApplicationVersionRequestTranscriberConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestTranscriberConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) GetCorrectionRules() []*UpdateApplicationVersionRequestTranscriberConfigCorrectionRules {
	return s.CorrectionRules
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) GetCustomizationId() *string {
	return s.CustomizationId
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) GetEndSilenceTimeout() *int32 {
	return s.EndSilenceTimeout
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) GetModel() *string {
	return s.Model
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) GetNlsAccessProfile() *UpdateApplicationVersionRequestTranscriberConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) GetSpeechNoiseThreshold() *int32 {
	return s.SpeechNoiseThreshold
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) SetCorrectionRules(v []*UpdateApplicationVersionRequestTranscriberConfigCorrectionRules) *UpdateApplicationVersionRequestTranscriberConfig {
	s.CorrectionRules = v
	return s
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) SetCustomizationId(v string) *UpdateApplicationVersionRequestTranscriberConfig {
	s.CustomizationId = &v
	return s
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) SetEndSilenceTimeout(v int32) *UpdateApplicationVersionRequestTranscriberConfig {
	s.EndSilenceTimeout = &v
	return s
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) SetModel(v string) *UpdateApplicationVersionRequestTranscriberConfig {
	s.Model = &v
	return s
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) SetNlsAccessProfile(v *UpdateApplicationVersionRequestTranscriberConfigNlsAccessProfile) *UpdateApplicationVersionRequestTranscriberConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) SetNlsAccessType(v string) *UpdateApplicationVersionRequestTranscriberConfig {
	s.NlsAccessType = &v
	return s
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) SetNlsEngine(v string) *UpdateApplicationVersionRequestTranscriberConfig {
	s.NlsEngine = &v
	return s
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) SetSpeechNoiseThreshold(v int32) *UpdateApplicationVersionRequestTranscriberConfig {
	s.SpeechNoiseThreshold = &v
	return s
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) SetVocabularyId(v string) *UpdateApplicationVersionRequestTranscriberConfig {
	s.VocabularyId = &v
	return s
}

func (s *UpdateApplicationVersionRequestTranscriberConfig) Validate() error {
	if s.CorrectionRules != nil {
		for _, item := range s.CorrectionRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NlsAccessProfile != nil {
		if err := s.NlsAccessProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationVersionRequestTranscriberConfigCorrectionRules struct {
	Pattern     *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
}

func (s UpdateApplicationVersionRequestTranscriberConfigCorrectionRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestTranscriberConfigCorrectionRules) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestTranscriberConfigCorrectionRules) GetPattern() *string {
	return s.Pattern
}

func (s *UpdateApplicationVersionRequestTranscriberConfigCorrectionRules) GetReplacement() *string {
	return s.Replacement
}

func (s *UpdateApplicationVersionRequestTranscriberConfigCorrectionRules) SetPattern(v string) *UpdateApplicationVersionRequestTranscriberConfigCorrectionRules {
	s.Pattern = &v
	return s
}

func (s *UpdateApplicationVersionRequestTranscriberConfigCorrectionRules) SetReplacement(v string) *UpdateApplicationVersionRequestTranscriberConfigCorrectionRules {
	s.Replacement = &v
	return s
}

func (s *UpdateApplicationVersionRequestTranscriberConfigCorrectionRules) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationVersionRequestTranscriberConfigNlsAccessProfile struct {
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s UpdateApplicationVersionRequestTranscriberConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestTranscriberConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestTranscriberConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *UpdateApplicationVersionRequestTranscriberConfigNlsAccessProfile) SetAccessProfileId(v string) *UpdateApplicationVersionRequestTranscriberConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *UpdateApplicationVersionRequestTranscriberConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}
