// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreateApplicationVersionRequest
	GetApplicationId() *string
	SetBusinessUnitId(v string) *CreateApplicationVersionRequest
	GetBusinessUnitId() *string
	SetInteractionConfig(v *CreateApplicationVersionRequestInteractionConfig) *CreateApplicationVersionRequest
	GetInteractionConfig() *CreateApplicationVersionRequestInteractionConfig
	SetLabelConfig(v []*CreateApplicationVersionRequestLabelConfig) *CreateApplicationVersionRequest
	GetLabelConfig() []*CreateApplicationVersionRequestLabelConfig
	SetRagConfig(v *CreateApplicationVersionRequestRagConfig) *CreateApplicationVersionRequest
	GetRagConfig() *CreateApplicationVersionRequestRagConfig
	SetScriptProfile(v *CreateApplicationVersionRequestScriptProfile) *CreateApplicationVersionRequest
	GetScriptProfile() *CreateApplicationVersionRequestScriptProfile
	SetSourceVersionId(v string) *CreateApplicationVersionRequest
	GetSourceVersionId() *string
	SetSynthesizerConfig(v *CreateApplicationVersionRequestSynthesizerConfig) *CreateApplicationVersionRequest
	GetSynthesizerConfig() *CreateApplicationVersionRequestSynthesizerConfig
	SetToolConfig(v *CreateApplicationVersionRequestToolConfig) *CreateApplicationVersionRequest
	GetToolConfig() *CreateApplicationVersionRequestToolConfig
	SetTranscriberConfig(v *CreateApplicationVersionRequestTranscriberConfig) *CreateApplicationVersionRequest
	GetTranscriberConfig() *CreateApplicationVersionRequestTranscriberConfig
}

type CreateApplicationVersionRequest struct {
	// This parameter is required.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	BusinessUnitId    *string                                           `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	InteractionConfig *CreateApplicationVersionRequestInteractionConfig `json:"InteractionConfig,omitempty" xml:"InteractionConfig,omitempty" type:"Struct"`
	LabelConfig       []*CreateApplicationVersionRequestLabelConfig     `json:"LabelConfig,omitempty" xml:"LabelConfig,omitempty" type:"Repeated"`
	RagConfig         *CreateApplicationVersionRequestRagConfig         `json:"RagConfig,omitempty" xml:"RagConfig,omitempty" type:"Struct"`
	ScriptProfile     *CreateApplicationVersionRequestScriptProfile     `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty" type:"Struct"`
	SourceVersionId   *string                                           `json:"SourceVersionId,omitempty" xml:"SourceVersionId,omitempty"`
	SynthesizerConfig *CreateApplicationVersionRequestSynthesizerConfig `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty" type:"Struct"`
	ToolConfig        *CreateApplicationVersionRequestToolConfig        `json:"ToolConfig,omitempty" xml:"ToolConfig,omitempty" type:"Struct"`
	TranscriberConfig *CreateApplicationVersionRequestTranscriberConfig `json:"TranscriberConfig,omitempty" xml:"TranscriberConfig,omitempty" type:"Struct"`
}

func (s CreateApplicationVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateApplicationVersionRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *CreateApplicationVersionRequest) GetInteractionConfig() *CreateApplicationVersionRequestInteractionConfig {
	return s.InteractionConfig
}

func (s *CreateApplicationVersionRequest) GetLabelConfig() []*CreateApplicationVersionRequestLabelConfig {
	return s.LabelConfig
}

func (s *CreateApplicationVersionRequest) GetRagConfig() *CreateApplicationVersionRequestRagConfig {
	return s.RagConfig
}

func (s *CreateApplicationVersionRequest) GetScriptProfile() *CreateApplicationVersionRequestScriptProfile {
	return s.ScriptProfile
}

func (s *CreateApplicationVersionRequest) GetSourceVersionId() *string {
	return s.SourceVersionId
}

func (s *CreateApplicationVersionRequest) GetSynthesizerConfig() *CreateApplicationVersionRequestSynthesizerConfig {
	return s.SynthesizerConfig
}

func (s *CreateApplicationVersionRequest) GetToolConfig() *CreateApplicationVersionRequestToolConfig {
	return s.ToolConfig
}

func (s *CreateApplicationVersionRequest) GetTranscriberConfig() *CreateApplicationVersionRequestTranscriberConfig {
	return s.TranscriberConfig
}

func (s *CreateApplicationVersionRequest) SetApplicationId(v string) *CreateApplicationVersionRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationVersionRequest) SetBusinessUnitId(v string) *CreateApplicationVersionRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *CreateApplicationVersionRequest) SetInteractionConfig(v *CreateApplicationVersionRequestInteractionConfig) *CreateApplicationVersionRequest {
	s.InteractionConfig = v
	return s
}

func (s *CreateApplicationVersionRequest) SetLabelConfig(v []*CreateApplicationVersionRequestLabelConfig) *CreateApplicationVersionRequest {
	s.LabelConfig = v
	return s
}

func (s *CreateApplicationVersionRequest) SetRagConfig(v *CreateApplicationVersionRequestRagConfig) *CreateApplicationVersionRequest {
	s.RagConfig = v
	return s
}

func (s *CreateApplicationVersionRequest) SetScriptProfile(v *CreateApplicationVersionRequestScriptProfile) *CreateApplicationVersionRequest {
	s.ScriptProfile = v
	return s
}

func (s *CreateApplicationVersionRequest) SetSourceVersionId(v string) *CreateApplicationVersionRequest {
	s.SourceVersionId = &v
	return s
}

func (s *CreateApplicationVersionRequest) SetSynthesizerConfig(v *CreateApplicationVersionRequestSynthesizerConfig) *CreateApplicationVersionRequest {
	s.SynthesizerConfig = v
	return s
}

func (s *CreateApplicationVersionRequest) SetToolConfig(v *CreateApplicationVersionRequestToolConfig) *CreateApplicationVersionRequest {
	s.ToolConfig = v
	return s
}

func (s *CreateApplicationVersionRequest) SetTranscriberConfig(v *CreateApplicationVersionRequestTranscriberConfig) *CreateApplicationVersionRequest {
	s.TranscriberConfig = v
	return s
}

func (s *CreateApplicationVersionRequest) Validate() error {
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

type CreateApplicationVersionRequestInteractionConfig struct {
	BackgroundMusicId                *string                                                                 `json:"BackgroundMusicId,omitempty" xml:"BackgroundMusicId,omitempty"`
	EndConversationConfig            *CreateApplicationVersionRequestInteractionConfigEndConversationConfig  `json:"EndConversationConfig,omitempty" xml:"EndConversationConfig,omitempty" type:"Struct"`
	InitialGreetingDelayMilliseconds *int32                                                                  `json:"InitialGreetingDelayMilliseconds,omitempty" xml:"InitialGreetingDelayMilliseconds,omitempty"`
	SilenceDetectionConfig           *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig `json:"SilenceDetectionConfig,omitempty" xml:"SilenceDetectionConfig,omitempty" type:"Struct"`
}

func (s CreateApplicationVersionRequestInteractionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestInteractionConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestInteractionConfig) GetBackgroundMusicId() *string {
	return s.BackgroundMusicId
}

func (s *CreateApplicationVersionRequestInteractionConfig) GetEndConversationConfig() *CreateApplicationVersionRequestInteractionConfigEndConversationConfig {
	return s.EndConversationConfig
}

func (s *CreateApplicationVersionRequestInteractionConfig) GetInitialGreetingDelayMilliseconds() *int32 {
	return s.InitialGreetingDelayMilliseconds
}

func (s *CreateApplicationVersionRequestInteractionConfig) GetSilenceDetectionConfig() *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig {
	return s.SilenceDetectionConfig
}

func (s *CreateApplicationVersionRequestInteractionConfig) SetBackgroundMusicId(v string) *CreateApplicationVersionRequestInteractionConfig {
	s.BackgroundMusicId = &v
	return s
}

func (s *CreateApplicationVersionRequestInteractionConfig) SetEndConversationConfig(v *CreateApplicationVersionRequestInteractionConfigEndConversationConfig) *CreateApplicationVersionRequestInteractionConfig {
	s.EndConversationConfig = v
	return s
}

func (s *CreateApplicationVersionRequestInteractionConfig) SetInitialGreetingDelayMilliseconds(v int32) *CreateApplicationVersionRequestInteractionConfig {
	s.InitialGreetingDelayMilliseconds = &v
	return s
}

func (s *CreateApplicationVersionRequestInteractionConfig) SetSilenceDetectionConfig(v *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) *CreateApplicationVersionRequestInteractionConfig {
	s.SilenceDetectionConfig = v
	return s
}

func (s *CreateApplicationVersionRequestInteractionConfig) Validate() error {
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

type CreateApplicationVersionRequestInteractionConfigEndConversationConfig struct {
	Delay    *int32                                                                           `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Triggers []*CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s CreateApplicationVersionRequestInteractionConfigEndConversationConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestInteractionConfigEndConversationConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestInteractionConfigEndConversationConfig) GetDelay() *int32 {
	return s.Delay
}

func (s *CreateApplicationVersionRequestInteractionConfigEndConversationConfig) GetTriggers() []*CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers {
	return s.Triggers
}

func (s *CreateApplicationVersionRequestInteractionConfigEndConversationConfig) SetDelay(v int32) *CreateApplicationVersionRequestInteractionConfigEndConversationConfig {
	s.Delay = &v
	return s
}

func (s *CreateApplicationVersionRequestInteractionConfigEndConversationConfig) SetTriggers(v []*CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) *CreateApplicationVersionRequestInteractionConfigEndConversationConfig {
	s.Triggers = v
	return s
}

func (s *CreateApplicationVersionRequestInteractionConfigEndConversationConfig) Validate() error {
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

type CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers struct {
	// example:
	//
	// 感谢您的接听，祝您生活愉快，再见!
	ClosingStatement *string   `json:"ClosingStatement,omitempty" xml:"ClosingStatement,omitempty"`
	KeyWords         []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	TriggerType      *string   `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// example:
	//
	// 20
	TurnLimit *int32 `json:"TurnLimit,omitempty" xml:"TurnLimit,omitempty"`
}

func (s CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) GetClosingStatement() *string {
	return s.ClosingStatement
}

func (s *CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) GetKeyWords() []*string {
	return s.KeyWords
}

func (s *CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) GetTriggerType() *string {
	return s.TriggerType
}

func (s *CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) GetTurnLimit() *int32 {
	return s.TurnLimit
}

func (s *CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) SetClosingStatement(v string) *CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.ClosingStatement = &v
	return s
}

func (s *CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) SetKeyWords(v []*string) *CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.KeyWords = v
	return s
}

func (s *CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) SetTriggerType(v string) *CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.TriggerType = &v
	return s
}

func (s *CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) SetTurnLimit(v int32) *CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.TurnLimit = &v
	return s
}

func (s *CreateApplicationVersionRequestInteractionConfigEndConversationConfigTriggers) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig struct {
	MaxRepeats *int32 `json:"MaxRepeats,omitempty" xml:"MaxRepeats,omitempty"`
	Timeout    *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) GetMaxRepeats() *int32 {
	return s.MaxRepeats
}

func (s *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) SetMaxRepeats(v int32) *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig {
	s.MaxRepeats = &v
	return s
}

func (s *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) SetTimeout(v int32) *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig {
	s.Timeout = &v
	return s
}

func (s *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationVersionRequestLabelConfig struct {
	CandidateValues []*string `json:"CandidateValues,omitempty" xml:"CandidateValues,omitempty" type:"Repeated"`
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateApplicationVersionRequestLabelConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestLabelConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestLabelConfig) GetCandidateValues() []*string {
	return s.CandidateValues
}

func (s *CreateApplicationVersionRequestLabelConfig) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationVersionRequestLabelConfig) GetName() *string {
	return s.Name
}

func (s *CreateApplicationVersionRequestLabelConfig) SetCandidateValues(v []*string) *CreateApplicationVersionRequestLabelConfig {
	s.CandidateValues = v
	return s
}

func (s *CreateApplicationVersionRequestLabelConfig) SetDescription(v string) *CreateApplicationVersionRequestLabelConfig {
	s.Description = &v
	return s
}

func (s *CreateApplicationVersionRequestLabelConfig) SetName(v string) *CreateApplicationVersionRequestLabelConfig {
	s.Name = &v
	return s
}

func (s *CreateApplicationVersionRequestLabelConfig) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationVersionRequestRagConfig struct {
	Enabled          *bool     `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	KnowledgeBaseIds []*string `json:"KnowledgeBaseIds,omitempty" xml:"KnowledgeBaseIds,omitempty" type:"Repeated"`
	MaxContentLength *int32    `json:"MaxContentLength,omitempty" xml:"MaxContentLength,omitempty"`
	RagEngine        *string   `json:"RagEngine,omitempty" xml:"RagEngine,omitempty"`
	TopN             *int32    `json:"TopN,omitempty" xml:"TopN,omitempty"`
}

func (s CreateApplicationVersionRequestRagConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestRagConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestRagConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateApplicationVersionRequestRagConfig) GetKnowledgeBaseIds() []*string {
	return s.KnowledgeBaseIds
}

func (s *CreateApplicationVersionRequestRagConfig) GetMaxContentLength() *int32 {
	return s.MaxContentLength
}

func (s *CreateApplicationVersionRequestRagConfig) GetRagEngine() *string {
	return s.RagEngine
}

func (s *CreateApplicationVersionRequestRagConfig) GetTopN() *int32 {
	return s.TopN
}

func (s *CreateApplicationVersionRequestRagConfig) SetEnabled(v bool) *CreateApplicationVersionRequestRagConfig {
	s.Enabled = &v
	return s
}

func (s *CreateApplicationVersionRequestRagConfig) SetKnowledgeBaseIds(v []*string) *CreateApplicationVersionRequestRagConfig {
	s.KnowledgeBaseIds = v
	return s
}

func (s *CreateApplicationVersionRequestRagConfig) SetMaxContentLength(v int32) *CreateApplicationVersionRequestRagConfig {
	s.MaxContentLength = &v
	return s
}

func (s *CreateApplicationVersionRequestRagConfig) SetRagEngine(v string) *CreateApplicationVersionRequestRagConfig {
	s.RagEngine = &v
	return s
}

func (s *CreateApplicationVersionRequestRagConfig) SetTopN(v int32) *CreateApplicationVersionRequestRagConfig {
	s.TopN = &v
	return s
}

func (s *CreateApplicationVersionRequestRagConfig) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationVersionRequestScriptProfile struct {
	// example:
	//
	// 6f444ecf21d94238b516735916c98666
	AgentKey     *string                                                   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	AgentProfile *CreateApplicationVersionRequestScriptProfileAgentProfile `json:"AgentProfile,omitempty" xml:"AgentProfile,omitempty" type:"Struct"`
	// example:
	//
	// chatbot-cn-MQuyjjb666
	ChatbotId        *string                                                       `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	FunctionMeta     *CreateApplicationVersionRequestScriptProfileFunctionMeta     `json:"FunctionMeta,omitempty" xml:"FunctionMeta,omitempty" type:"Struct"`
	Model            *string                                                       `json:"Model,omitempty" xml:"Model,omitempty"`
	NluAccessProfile *CreateApplicationVersionRequestScriptProfileNluAccessProfile `json:"NluAccessProfile,omitempty" xml:"NluAccessProfile,omitempty" type:"Struct"`
	NluAccessType    *string                                                       `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	OmniModel        *bool                                                         `json:"OmniModel,omitempty" xml:"OmniModel,omitempty"`
}

func (s CreateApplicationVersionRequestScriptProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestScriptProfile) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestScriptProfile) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateApplicationVersionRequestScriptProfile) GetAgentProfile() *CreateApplicationVersionRequestScriptProfileAgentProfile {
	return s.AgentProfile
}

func (s *CreateApplicationVersionRequestScriptProfile) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *CreateApplicationVersionRequestScriptProfile) GetFunctionMeta() *CreateApplicationVersionRequestScriptProfileFunctionMeta {
	return s.FunctionMeta
}

func (s *CreateApplicationVersionRequestScriptProfile) GetModel() *string {
	return s.Model
}

func (s *CreateApplicationVersionRequestScriptProfile) GetNluAccessProfile() *CreateApplicationVersionRequestScriptProfileNluAccessProfile {
	return s.NluAccessProfile
}

func (s *CreateApplicationVersionRequestScriptProfile) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *CreateApplicationVersionRequestScriptProfile) GetOmniModel() *bool {
	return s.OmniModel
}

func (s *CreateApplicationVersionRequestScriptProfile) SetAgentKey(v string) *CreateApplicationVersionRequestScriptProfile {
	s.AgentKey = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfile) SetAgentProfile(v *CreateApplicationVersionRequestScriptProfileAgentProfile) *CreateApplicationVersionRequestScriptProfile {
	s.AgentProfile = v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfile) SetChatbotId(v string) *CreateApplicationVersionRequestScriptProfile {
	s.ChatbotId = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfile) SetFunctionMeta(v *CreateApplicationVersionRequestScriptProfileFunctionMeta) *CreateApplicationVersionRequestScriptProfile {
	s.FunctionMeta = v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfile) SetModel(v string) *CreateApplicationVersionRequestScriptProfile {
	s.Model = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfile) SetNluAccessProfile(v *CreateApplicationVersionRequestScriptProfileNluAccessProfile) *CreateApplicationVersionRequestScriptProfile {
	s.NluAccessProfile = v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfile) SetNluAccessType(v string) *CreateApplicationVersionRequestScriptProfile {
	s.NluAccessType = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfile) SetOmniModel(v bool) *CreateApplicationVersionRequestScriptProfile {
	s.OmniModel = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfile) Validate() error {
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

type CreateApplicationVersionRequestScriptProfileAgentProfile struct {
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PromptsJson             *string `json:"PromptsJson,omitempty" xml:"PromptsJson,omitempty"`
	ScriptProfileTemplateId *string `json:"ScriptProfileTemplateId,omitempty" xml:"ScriptProfileTemplateId,omitempty"`
}

func (s CreateApplicationVersionRequestScriptProfileAgentProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestScriptProfileAgentProfile) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestScriptProfileAgentProfile) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationVersionRequestScriptProfileAgentProfile) GetName() *string {
	return s.Name
}

func (s *CreateApplicationVersionRequestScriptProfileAgentProfile) GetPromptsJson() *string {
	return s.PromptsJson
}

func (s *CreateApplicationVersionRequestScriptProfileAgentProfile) GetScriptProfileTemplateId() *string {
	return s.ScriptProfileTemplateId
}

func (s *CreateApplicationVersionRequestScriptProfileAgentProfile) SetDescription(v string) *CreateApplicationVersionRequestScriptProfileAgentProfile {
	s.Description = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfileAgentProfile) SetName(v string) *CreateApplicationVersionRequestScriptProfileAgentProfile {
	s.Name = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfileAgentProfile) SetPromptsJson(v string) *CreateApplicationVersionRequestScriptProfileAgentProfile {
	s.PromptsJson = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfileAgentProfile) SetScriptProfileTemplateId(v string) *CreateApplicationVersionRequestScriptProfileAgentProfile {
	s.ScriptProfileTemplateId = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfileAgentProfile) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationVersionRequestScriptProfileFunctionMeta struct {
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

func (s CreateApplicationVersionRequestScriptProfileFunctionMeta) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestScriptProfileFunctionMeta) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestScriptProfileFunctionMeta) GetFunctionId() *string {
	return s.FunctionId
}

func (s *CreateApplicationVersionRequestScriptProfileFunctionMeta) GetFunctionName() *string {
	return s.FunctionName
}

func (s *CreateApplicationVersionRequestScriptProfileFunctionMeta) GetHttpTriggerName() *string {
	return s.HttpTriggerName
}

func (s *CreateApplicationVersionRequestScriptProfileFunctionMeta) GetHttpTriggerUrl() *string {
	return s.HttpTriggerUrl
}

func (s *CreateApplicationVersionRequestScriptProfileFunctionMeta) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApplicationVersionRequestScriptProfileFunctionMeta) SetFunctionId(v string) *CreateApplicationVersionRequestScriptProfileFunctionMeta {
	s.FunctionId = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfileFunctionMeta) SetFunctionName(v string) *CreateApplicationVersionRequestScriptProfileFunctionMeta {
	s.FunctionName = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfileFunctionMeta) SetHttpTriggerName(v string) *CreateApplicationVersionRequestScriptProfileFunctionMeta {
	s.HttpTriggerName = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfileFunctionMeta) SetHttpTriggerUrl(v string) *CreateApplicationVersionRequestScriptProfileFunctionMeta {
	s.HttpTriggerUrl = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfileFunctionMeta) SetRegionId(v string) *CreateApplicationVersionRequestScriptProfileFunctionMeta {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfileFunctionMeta) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationVersionRequestScriptProfileNluAccessProfile struct {
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s CreateApplicationVersionRequestScriptProfileNluAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestScriptProfileNluAccessProfile) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestScriptProfileNluAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *CreateApplicationVersionRequestScriptProfileNluAccessProfile) SetAccessProfileId(v string) *CreateApplicationVersionRequestScriptProfileNluAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfileNluAccessProfile) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationVersionRequestSynthesizerConfig struct {
	Model            *string                                                           `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile *CreateApplicationVersionRequestSynthesizerConfigNlsAccessProfile `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	NlsAccessType    *string                                                           `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	NlsEngine        *string                                                           `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	PitchRate        *int32                                                            `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	PronRules        []*CreateApplicationVersionRequestSynthesizerConfigPronRules      `json:"PronRules,omitempty" xml:"PronRules,omitempty" type:"Repeated"`
	SpeechRate       *int32                                                            `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Voice            *string                                                           `json:"Voice,omitempty" xml:"Voice,omitempty"`
	Volume           *int32                                                            `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s CreateApplicationVersionRequestSynthesizerConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestSynthesizerConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) GetModel() *string {
	return s.Model
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) GetNlsAccessProfile() *CreateApplicationVersionRequestSynthesizerConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) GetPitchRate() *int32 {
	return s.PitchRate
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) GetPronRules() []*CreateApplicationVersionRequestSynthesizerConfigPronRules {
	return s.PronRules
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) GetVoice() *string {
	return s.Voice
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) SetModel(v string) *CreateApplicationVersionRequestSynthesizerConfig {
	s.Model = &v
	return s
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) SetNlsAccessProfile(v *CreateApplicationVersionRequestSynthesizerConfigNlsAccessProfile) *CreateApplicationVersionRequestSynthesizerConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) SetNlsAccessType(v string) *CreateApplicationVersionRequestSynthesizerConfig {
	s.NlsAccessType = &v
	return s
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) SetNlsEngine(v string) *CreateApplicationVersionRequestSynthesizerConfig {
	s.NlsEngine = &v
	return s
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) SetPitchRate(v int32) *CreateApplicationVersionRequestSynthesizerConfig {
	s.PitchRate = &v
	return s
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) SetPronRules(v []*CreateApplicationVersionRequestSynthesizerConfigPronRules) *CreateApplicationVersionRequestSynthesizerConfig {
	s.PronRules = v
	return s
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) SetSpeechRate(v int32) *CreateApplicationVersionRequestSynthesizerConfig {
	s.SpeechRate = &v
	return s
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) SetVoice(v string) *CreateApplicationVersionRequestSynthesizerConfig {
	s.Voice = &v
	return s
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) SetVolume(v int32) *CreateApplicationVersionRequestSynthesizerConfig {
	s.Volume = &v
	return s
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) Validate() error {
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

type CreateApplicationVersionRequestSynthesizerConfigNlsAccessProfile struct {
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s CreateApplicationVersionRequestSynthesizerConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestSynthesizerConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestSynthesizerConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *CreateApplicationVersionRequestSynthesizerConfigNlsAccessProfile) SetAccessProfileId(v string) *CreateApplicationVersionRequestSynthesizerConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *CreateApplicationVersionRequestSynthesizerConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationVersionRequestSynthesizerConfigPronRules struct {
	Pattern     *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
}

func (s CreateApplicationVersionRequestSynthesizerConfigPronRules) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestSynthesizerConfigPronRules) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestSynthesizerConfigPronRules) GetPattern() *string {
	return s.Pattern
}

func (s *CreateApplicationVersionRequestSynthesizerConfigPronRules) GetReplacement() *string {
	return s.Replacement
}

func (s *CreateApplicationVersionRequestSynthesizerConfigPronRules) SetPattern(v string) *CreateApplicationVersionRequestSynthesizerConfigPronRules {
	s.Pattern = &v
	return s
}

func (s *CreateApplicationVersionRequestSynthesizerConfigPronRules) SetReplacement(v string) *CreateApplicationVersionRequestSynthesizerConfigPronRules {
	s.Replacement = &v
	return s
}

func (s *CreateApplicationVersionRequestSynthesizerConfigPronRules) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationVersionRequestToolConfig struct {
	McpServers []*CreateApplicationVersionRequestToolConfigMcpServers `json:"McpServers,omitempty" xml:"McpServers,omitempty" type:"Repeated"`
}

func (s CreateApplicationVersionRequestToolConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestToolConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestToolConfig) GetMcpServers() []*CreateApplicationVersionRequestToolConfigMcpServers {
	return s.McpServers
}

func (s *CreateApplicationVersionRequestToolConfig) SetMcpServers(v []*CreateApplicationVersionRequestToolConfigMcpServers) *CreateApplicationVersionRequestToolConfig {
	s.McpServers = v
	return s
}

func (s *CreateApplicationVersionRequestToolConfig) Validate() error {
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

type CreateApplicationVersionRequestToolConfigMcpServers struct {
	BaseUrl     *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SseEndpoint *string `json:"SseEndpoint,omitempty" xml:"SseEndpoint,omitempty"`
}

func (s CreateApplicationVersionRequestToolConfigMcpServers) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestToolConfigMcpServers) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestToolConfigMcpServers) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *CreateApplicationVersionRequestToolConfigMcpServers) GetName() *string {
	return s.Name
}

func (s *CreateApplicationVersionRequestToolConfigMcpServers) GetSseEndpoint() *string {
	return s.SseEndpoint
}

func (s *CreateApplicationVersionRequestToolConfigMcpServers) SetBaseUrl(v string) *CreateApplicationVersionRequestToolConfigMcpServers {
	s.BaseUrl = &v
	return s
}

func (s *CreateApplicationVersionRequestToolConfigMcpServers) SetName(v string) *CreateApplicationVersionRequestToolConfigMcpServers {
	s.Name = &v
	return s
}

func (s *CreateApplicationVersionRequestToolConfigMcpServers) SetSseEndpoint(v string) *CreateApplicationVersionRequestToolConfigMcpServers {
	s.SseEndpoint = &v
	return s
}

func (s *CreateApplicationVersionRequestToolConfigMcpServers) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationVersionRequestTranscriberConfig struct {
	CorrectionRules      []*CreateApplicationVersionRequestTranscriberConfigCorrectionRules `json:"CorrectionRules,omitempty" xml:"CorrectionRules,omitempty" type:"Repeated"`
	CustomizationId      *string                                                            `json:"CustomizationId,omitempty" xml:"CustomizationId,omitempty"`
	EndSilenceTimeout    *int32                                                             `json:"EndSilenceTimeout,omitempty" xml:"EndSilenceTimeout,omitempty"`
	Model                *string                                                            `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile     *CreateApplicationVersionRequestTranscriberConfigNlsAccessProfile  `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	NlsAccessType        *string                                                            `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	NlsEngine            *string                                                            `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	SpeechNoiseThreshold *int32                                                             `json:"SpeechNoiseThreshold,omitempty" xml:"SpeechNoiseThreshold,omitempty"`
	VocabularyId         *string                                                            `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
}

func (s CreateApplicationVersionRequestTranscriberConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestTranscriberConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestTranscriberConfig) GetCorrectionRules() []*CreateApplicationVersionRequestTranscriberConfigCorrectionRules {
	return s.CorrectionRules
}

func (s *CreateApplicationVersionRequestTranscriberConfig) GetCustomizationId() *string {
	return s.CustomizationId
}

func (s *CreateApplicationVersionRequestTranscriberConfig) GetEndSilenceTimeout() *int32 {
	return s.EndSilenceTimeout
}

func (s *CreateApplicationVersionRequestTranscriberConfig) GetModel() *string {
	return s.Model
}

func (s *CreateApplicationVersionRequestTranscriberConfig) GetNlsAccessProfile() *CreateApplicationVersionRequestTranscriberConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *CreateApplicationVersionRequestTranscriberConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *CreateApplicationVersionRequestTranscriberConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *CreateApplicationVersionRequestTranscriberConfig) GetSpeechNoiseThreshold() *int32 {
	return s.SpeechNoiseThreshold
}

func (s *CreateApplicationVersionRequestTranscriberConfig) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *CreateApplicationVersionRequestTranscriberConfig) SetCorrectionRules(v []*CreateApplicationVersionRequestTranscriberConfigCorrectionRules) *CreateApplicationVersionRequestTranscriberConfig {
	s.CorrectionRules = v
	return s
}

func (s *CreateApplicationVersionRequestTranscriberConfig) SetCustomizationId(v string) *CreateApplicationVersionRequestTranscriberConfig {
	s.CustomizationId = &v
	return s
}

func (s *CreateApplicationVersionRequestTranscriberConfig) SetEndSilenceTimeout(v int32) *CreateApplicationVersionRequestTranscriberConfig {
	s.EndSilenceTimeout = &v
	return s
}

func (s *CreateApplicationVersionRequestTranscriberConfig) SetModel(v string) *CreateApplicationVersionRequestTranscriberConfig {
	s.Model = &v
	return s
}

func (s *CreateApplicationVersionRequestTranscriberConfig) SetNlsAccessProfile(v *CreateApplicationVersionRequestTranscriberConfigNlsAccessProfile) *CreateApplicationVersionRequestTranscriberConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *CreateApplicationVersionRequestTranscriberConfig) SetNlsAccessType(v string) *CreateApplicationVersionRequestTranscriberConfig {
	s.NlsAccessType = &v
	return s
}

func (s *CreateApplicationVersionRequestTranscriberConfig) SetNlsEngine(v string) *CreateApplicationVersionRequestTranscriberConfig {
	s.NlsEngine = &v
	return s
}

func (s *CreateApplicationVersionRequestTranscriberConfig) SetSpeechNoiseThreshold(v int32) *CreateApplicationVersionRequestTranscriberConfig {
	s.SpeechNoiseThreshold = &v
	return s
}

func (s *CreateApplicationVersionRequestTranscriberConfig) SetVocabularyId(v string) *CreateApplicationVersionRequestTranscriberConfig {
	s.VocabularyId = &v
	return s
}

func (s *CreateApplicationVersionRequestTranscriberConfig) Validate() error {
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

type CreateApplicationVersionRequestTranscriberConfigCorrectionRules struct {
	Pattern     *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
}

func (s CreateApplicationVersionRequestTranscriberConfigCorrectionRules) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestTranscriberConfigCorrectionRules) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestTranscriberConfigCorrectionRules) GetPattern() *string {
	return s.Pattern
}

func (s *CreateApplicationVersionRequestTranscriberConfigCorrectionRules) GetReplacement() *string {
	return s.Replacement
}

func (s *CreateApplicationVersionRequestTranscriberConfigCorrectionRules) SetPattern(v string) *CreateApplicationVersionRequestTranscriberConfigCorrectionRules {
	s.Pattern = &v
	return s
}

func (s *CreateApplicationVersionRequestTranscriberConfigCorrectionRules) SetReplacement(v string) *CreateApplicationVersionRequestTranscriberConfigCorrectionRules {
	s.Replacement = &v
	return s
}

func (s *CreateApplicationVersionRequestTranscriberConfigCorrectionRules) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationVersionRequestTranscriberConfigNlsAccessProfile struct {
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s CreateApplicationVersionRequestTranscriberConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestTranscriberConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestTranscriberConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *CreateApplicationVersionRequestTranscriberConfigNlsAccessProfile) SetAccessProfileId(v string) *CreateApplicationVersionRequestTranscriberConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *CreateApplicationVersionRequestTranscriberConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}
