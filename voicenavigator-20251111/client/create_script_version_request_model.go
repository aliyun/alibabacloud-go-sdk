// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateScriptVersionRequest
	GetInstanceId() *string
	SetInteractionConfig(v *CreateScriptVersionRequestInteractionConfig) *CreateScriptVersionRequest
	GetInteractionConfig() *CreateScriptVersionRequestInteractionConfig
	SetLabelConfig(v []*CreateScriptVersionRequestLabelConfig) *CreateScriptVersionRequest
	GetLabelConfig() []*CreateScriptVersionRequestLabelConfig
	SetScriptId(v string) *CreateScriptVersionRequest
	GetScriptId() *string
	SetScriptProfile(v *CreateScriptVersionRequestScriptProfile) *CreateScriptVersionRequest
	GetScriptProfile() *CreateScriptVersionRequestScriptProfile
	SetSourceVersionId(v string) *CreateScriptVersionRequest
	GetSourceVersionId() *string
	SetSynthesizerConfig(v *CreateScriptVersionRequestSynthesizerConfig) *CreateScriptVersionRequest
	GetSynthesizerConfig() *CreateScriptVersionRequestSynthesizerConfig
	SetTranscriberConfig(v *CreateScriptVersionRequestTranscriberConfig) *CreateScriptVersionRequest
	GetTranscriberConfig() *CreateScriptVersionRequestTranscriberConfig
}

type CreateScriptVersionRequest struct {
	// example:
	//
	// abb4aa26-3a8e-43dd-82f8-0c3898c9c67f
	InstanceId        *string                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InteractionConfig *CreateScriptVersionRequestInteractionConfig `json:"InteractionConfig,omitempty" xml:"InteractionConfig,omitempty" type:"Struct"`
	LabelConfig       []*CreateScriptVersionRequestLabelConfig     `json:"LabelConfig,omitempty" xml:"LabelConfig,omitempty" type:"Repeated"`
	// example:
	//
	// 64241e64-190c-45d1-af66-06f51c07b090
	ScriptId          *string                                      `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	ScriptProfile     *CreateScriptVersionRequestScriptProfile     `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty" type:"Struct"`
	SourceVersionId   *string                                      `json:"SourceVersionId,omitempty" xml:"SourceVersionId,omitempty"`
	SynthesizerConfig *CreateScriptVersionRequestSynthesizerConfig `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty" type:"Struct"`
	TranscriberConfig *CreateScriptVersionRequestTranscriberConfig `json:"TranscriberConfig,omitempty" xml:"TranscriberConfig,omitempty" type:"Struct"`
}

func (s CreateScriptVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateScriptVersionRequest) GetInteractionConfig() *CreateScriptVersionRequestInteractionConfig {
	return s.InteractionConfig
}

func (s *CreateScriptVersionRequest) GetLabelConfig() []*CreateScriptVersionRequestLabelConfig {
	return s.LabelConfig
}

func (s *CreateScriptVersionRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateScriptVersionRequest) GetScriptProfile() *CreateScriptVersionRequestScriptProfile {
	return s.ScriptProfile
}

func (s *CreateScriptVersionRequest) GetSourceVersionId() *string {
	return s.SourceVersionId
}

func (s *CreateScriptVersionRequest) GetSynthesizerConfig() *CreateScriptVersionRequestSynthesizerConfig {
	return s.SynthesizerConfig
}

func (s *CreateScriptVersionRequest) GetTranscriberConfig() *CreateScriptVersionRequestTranscriberConfig {
	return s.TranscriberConfig
}

func (s *CreateScriptVersionRequest) SetInstanceId(v string) *CreateScriptVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScriptVersionRequest) SetInteractionConfig(v *CreateScriptVersionRequestInteractionConfig) *CreateScriptVersionRequest {
	s.InteractionConfig = v
	return s
}

func (s *CreateScriptVersionRequest) SetLabelConfig(v []*CreateScriptVersionRequestLabelConfig) *CreateScriptVersionRequest {
	s.LabelConfig = v
	return s
}

func (s *CreateScriptVersionRequest) SetScriptId(v string) *CreateScriptVersionRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateScriptVersionRequest) SetScriptProfile(v *CreateScriptVersionRequestScriptProfile) *CreateScriptVersionRequest {
	s.ScriptProfile = v
	return s
}

func (s *CreateScriptVersionRequest) SetSourceVersionId(v string) *CreateScriptVersionRequest {
	s.SourceVersionId = &v
	return s
}

func (s *CreateScriptVersionRequest) SetSynthesizerConfig(v *CreateScriptVersionRequestSynthesizerConfig) *CreateScriptVersionRequest {
	s.SynthesizerConfig = v
	return s
}

func (s *CreateScriptVersionRequest) SetTranscriberConfig(v *CreateScriptVersionRequestTranscriberConfig) *CreateScriptVersionRequest {
	s.TranscriberConfig = v
	return s
}

func (s *CreateScriptVersionRequest) Validate() error {
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
	if s.TranscriberConfig != nil {
		if err := s.TranscriberConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateScriptVersionRequestInteractionConfig struct {
	// example:
	//
	// office-ambience
	BackgroundMusicId     *string                                                           `json:"BackgroundMusicId,omitempty" xml:"BackgroundMusicId,omitempty"`
	EndConversationConfig *CreateScriptVersionRequestInteractionConfigEndConversationConfig `json:"EndConversationConfig,omitempty" xml:"EndConversationConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2000
	InitialGreetingDelayMilliseconds *int32                                                             `json:"InitialGreetingDelayMilliseconds,omitempty" xml:"InitialGreetingDelayMilliseconds,omitempty"`
	SilenceDetectionConfig           *CreateScriptVersionRequestInteractionConfigSilenceDetectionConfig `json:"SilenceDetectionConfig,omitempty" xml:"SilenceDetectionConfig,omitempty" type:"Struct"`
}

func (s CreateScriptVersionRequestInteractionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionRequestInteractionConfig) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionRequestInteractionConfig) GetBackgroundMusicId() *string {
	return s.BackgroundMusicId
}

func (s *CreateScriptVersionRequestInteractionConfig) GetEndConversationConfig() *CreateScriptVersionRequestInteractionConfigEndConversationConfig {
	return s.EndConversationConfig
}

func (s *CreateScriptVersionRequestInteractionConfig) GetInitialGreetingDelayMilliseconds() *int32 {
	return s.InitialGreetingDelayMilliseconds
}

func (s *CreateScriptVersionRequestInteractionConfig) GetSilenceDetectionConfig() *CreateScriptVersionRequestInteractionConfigSilenceDetectionConfig {
	return s.SilenceDetectionConfig
}

func (s *CreateScriptVersionRequestInteractionConfig) SetBackgroundMusicId(v string) *CreateScriptVersionRequestInteractionConfig {
	s.BackgroundMusicId = &v
	return s
}

func (s *CreateScriptVersionRequestInteractionConfig) SetEndConversationConfig(v *CreateScriptVersionRequestInteractionConfigEndConversationConfig) *CreateScriptVersionRequestInteractionConfig {
	s.EndConversationConfig = v
	return s
}

func (s *CreateScriptVersionRequestInteractionConfig) SetInitialGreetingDelayMilliseconds(v int32) *CreateScriptVersionRequestInteractionConfig {
	s.InitialGreetingDelayMilliseconds = &v
	return s
}

func (s *CreateScriptVersionRequestInteractionConfig) SetSilenceDetectionConfig(v *CreateScriptVersionRequestInteractionConfigSilenceDetectionConfig) *CreateScriptVersionRequestInteractionConfig {
	s.SilenceDetectionConfig = v
	return s
}

func (s *CreateScriptVersionRequestInteractionConfig) Validate() error {
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

type CreateScriptVersionRequestInteractionConfigEndConversationConfig struct {
	// example:
	//
	// 1
	Delay    *int32                                                                      `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Triggers []*CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s CreateScriptVersionRequestInteractionConfigEndConversationConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionRequestInteractionConfigEndConversationConfig) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionRequestInteractionConfigEndConversationConfig) GetDelay() *int32 {
	return s.Delay
}

func (s *CreateScriptVersionRequestInteractionConfigEndConversationConfig) GetTriggers() []*CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers {
	return s.Triggers
}

func (s *CreateScriptVersionRequestInteractionConfigEndConversationConfig) SetDelay(v int32) *CreateScriptVersionRequestInteractionConfigEndConversationConfig {
	s.Delay = &v
	return s
}

func (s *CreateScriptVersionRequestInteractionConfigEndConversationConfig) SetTriggers(v []*CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) *CreateScriptVersionRequestInteractionConfigEndConversationConfig {
	s.Triggers = v
	return s
}

func (s *CreateScriptVersionRequestInteractionConfigEndConversationConfig) Validate() error {
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

type CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers struct {
	// example:
	//
	// 感谢您的接听，祝您生活愉快，再见!
	ClosingStatement *string   `json:"ClosingStatement,omitempty" xml:"ClosingStatement,omitempty"`
	KeyWords         []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	// example:
	//
	// TurnLimit
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// example:
	//
	// 20
	TurnLimit *int32 `json:"TurnLimit,omitempty" xml:"TurnLimit,omitempty"`
}

func (s CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) GetClosingStatement() *string {
	return s.ClosingStatement
}

func (s *CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) GetKeyWords() []*string {
	return s.KeyWords
}

func (s *CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) GetTriggerType() *string {
	return s.TriggerType
}

func (s *CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) GetTurnLimit() *int32 {
	return s.TurnLimit
}

func (s *CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) SetClosingStatement(v string) *CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.ClosingStatement = &v
	return s
}

func (s *CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) SetKeyWords(v []*string) *CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.KeyWords = v
	return s
}

func (s *CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) SetTriggerType(v string) *CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.TriggerType = &v
	return s
}

func (s *CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) SetTurnLimit(v int32) *CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.TurnLimit = &v
	return s
}

func (s *CreateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) Validate() error {
	return dara.Validate(s)
}

type CreateScriptVersionRequestInteractionConfigSilenceDetectionConfig struct {
	// example:
	//
	// 3
	MaxRepeats *int32 `json:"MaxRepeats,omitempty" xml:"MaxRepeats,omitempty"`
	// example:
	//
	// 5000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateScriptVersionRequestInteractionConfigSilenceDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionRequestInteractionConfigSilenceDetectionConfig) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionRequestInteractionConfigSilenceDetectionConfig) GetMaxRepeats() *int32 {
	return s.MaxRepeats
}

func (s *CreateScriptVersionRequestInteractionConfigSilenceDetectionConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateScriptVersionRequestInteractionConfigSilenceDetectionConfig) SetMaxRepeats(v int32) *CreateScriptVersionRequestInteractionConfigSilenceDetectionConfig {
	s.MaxRepeats = &v
	return s
}

func (s *CreateScriptVersionRequestInteractionConfigSilenceDetectionConfig) SetTimeout(v int32) *CreateScriptVersionRequestInteractionConfigSilenceDetectionConfig {
	s.Timeout = &v
	return s
}

func (s *CreateScriptVersionRequestInteractionConfigSilenceDetectionConfig) Validate() error {
	return dara.Validate(s)
}

type CreateScriptVersionRequestLabelConfig struct {
	CandidateValues []*string `json:"CandidateValues,omitempty" xml:"CandidateValues,omitempty" type:"Repeated"`
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateScriptVersionRequestLabelConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionRequestLabelConfig) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionRequestLabelConfig) GetCandidateValues() []*string {
	return s.CandidateValues
}

func (s *CreateScriptVersionRequestLabelConfig) GetDescription() *string {
	return s.Description
}

func (s *CreateScriptVersionRequestLabelConfig) GetName() *string {
	return s.Name
}

func (s *CreateScriptVersionRequestLabelConfig) SetCandidateValues(v []*string) *CreateScriptVersionRequestLabelConfig {
	s.CandidateValues = v
	return s
}

func (s *CreateScriptVersionRequestLabelConfig) SetDescription(v string) *CreateScriptVersionRequestLabelConfig {
	s.Description = &v
	return s
}

func (s *CreateScriptVersionRequestLabelConfig) SetName(v string) *CreateScriptVersionRequestLabelConfig {
	s.Name = &v
	return s
}

func (s *CreateScriptVersionRequestLabelConfig) Validate() error {
	return dara.Validate(s)
}

type CreateScriptVersionRequestScriptProfile struct {
	// example:
	//
	// 1309723684579735_p_beebot_public
	AgentKey     *string                                              `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	AgentProfile *CreateScriptVersionRequestScriptProfileAgentProfile `json:"AgentProfile,omitempty" xml:"AgentProfile,omitempty" type:"Struct"`
	// example:
	//
	// chatbot-cn-MQuyjjb666
	ChatbotId    *string                                              `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	FunctionMeta *CreateScriptVersionRequestScriptProfileFunctionMeta `json:"FunctionMeta,omitempty" xml:"FunctionMeta,omitempty" type:"Struct"`
	// example:
	//
	// qwen-plus
	Model            *string                                                  `json:"Model,omitempty" xml:"Model,omitempty"`
	NluAccessProfile *CreateScriptVersionRequestScriptProfileNluAccessProfile `json:"NluAccessProfile,omitempty" xml:"NluAccessProfile,omitempty" type:"Struct"`
	// example:
	//
	// MANAGED
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// example:
	//
	// true
	OmniModel *bool `json:"OmniModel,omitempty" xml:"OmniModel,omitempty"`
}

func (s CreateScriptVersionRequestScriptProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionRequestScriptProfile) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionRequestScriptProfile) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateScriptVersionRequestScriptProfile) GetAgentProfile() *CreateScriptVersionRequestScriptProfileAgentProfile {
	return s.AgentProfile
}

func (s *CreateScriptVersionRequestScriptProfile) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *CreateScriptVersionRequestScriptProfile) GetFunctionMeta() *CreateScriptVersionRequestScriptProfileFunctionMeta {
	return s.FunctionMeta
}

func (s *CreateScriptVersionRequestScriptProfile) GetModel() *string {
	return s.Model
}

func (s *CreateScriptVersionRequestScriptProfile) GetNluAccessProfile() *CreateScriptVersionRequestScriptProfileNluAccessProfile {
	return s.NluAccessProfile
}

func (s *CreateScriptVersionRequestScriptProfile) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *CreateScriptVersionRequestScriptProfile) GetOmniModel() *bool {
	return s.OmniModel
}

func (s *CreateScriptVersionRequestScriptProfile) SetAgentKey(v string) *CreateScriptVersionRequestScriptProfile {
	s.AgentKey = &v
	return s
}

func (s *CreateScriptVersionRequestScriptProfile) SetAgentProfile(v *CreateScriptVersionRequestScriptProfileAgentProfile) *CreateScriptVersionRequestScriptProfile {
	s.AgentProfile = v
	return s
}

func (s *CreateScriptVersionRequestScriptProfile) SetChatbotId(v string) *CreateScriptVersionRequestScriptProfile {
	s.ChatbotId = &v
	return s
}

func (s *CreateScriptVersionRequestScriptProfile) SetFunctionMeta(v *CreateScriptVersionRequestScriptProfileFunctionMeta) *CreateScriptVersionRequestScriptProfile {
	s.FunctionMeta = v
	return s
}

func (s *CreateScriptVersionRequestScriptProfile) SetModel(v string) *CreateScriptVersionRequestScriptProfile {
	s.Model = &v
	return s
}

func (s *CreateScriptVersionRequestScriptProfile) SetNluAccessProfile(v *CreateScriptVersionRequestScriptProfileNluAccessProfile) *CreateScriptVersionRequestScriptProfile {
	s.NluAccessProfile = v
	return s
}

func (s *CreateScriptVersionRequestScriptProfile) SetNluAccessType(v string) *CreateScriptVersionRequestScriptProfile {
	s.NluAccessType = &v
	return s
}

func (s *CreateScriptVersionRequestScriptProfile) SetOmniModel(v bool) *CreateScriptVersionRequestScriptProfile {
	s.OmniModel = &v
	return s
}

func (s *CreateScriptVersionRequestScriptProfile) Validate() error {
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

type CreateScriptVersionRequestScriptProfileAgentProfile struct {
	PromptsJson *string `json:"PromptsJson,omitempty" xml:"PromptsJson,omitempty"`
	// example:
	//
	// CCC_PROMPTS_DEFAULT
	ScriptProfileTemplateId *string `json:"ScriptProfileTemplateId,omitempty" xml:"ScriptProfileTemplateId,omitempty"`
}

func (s CreateScriptVersionRequestScriptProfileAgentProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionRequestScriptProfileAgentProfile) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionRequestScriptProfileAgentProfile) GetPromptsJson() *string {
	return s.PromptsJson
}

func (s *CreateScriptVersionRequestScriptProfileAgentProfile) GetScriptProfileTemplateId() *string {
	return s.ScriptProfileTemplateId
}

func (s *CreateScriptVersionRequestScriptProfileAgentProfile) SetPromptsJson(v string) *CreateScriptVersionRequestScriptProfileAgentProfile {
	s.PromptsJson = &v
	return s
}

func (s *CreateScriptVersionRequestScriptProfileAgentProfile) SetScriptProfileTemplateId(v string) *CreateScriptVersionRequestScriptProfileAgentProfile {
	s.ScriptProfileTemplateId = &v
	return s
}

func (s *CreateScriptVersionRequestScriptProfileAgentProfile) Validate() error {
	return dara.Validate(s)
}

type CreateScriptVersionRequestScriptProfileFunctionMeta struct {
	// example:
	//
	// 9b752bbb-805a-4d3e-9013-eab5555c3fef
	FunctionId *string `json:"FunctionId,omitempty" xml:"FunctionId,omitempty"`
	// example:
	//
	// my_funciton
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// example:
	//
	// defaultTrigger
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

func (s CreateScriptVersionRequestScriptProfileFunctionMeta) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionRequestScriptProfileFunctionMeta) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionRequestScriptProfileFunctionMeta) GetFunctionId() *string {
	return s.FunctionId
}

func (s *CreateScriptVersionRequestScriptProfileFunctionMeta) GetFunctionName() *string {
	return s.FunctionName
}

func (s *CreateScriptVersionRequestScriptProfileFunctionMeta) GetHttpTriggerName() *string {
	return s.HttpTriggerName
}

func (s *CreateScriptVersionRequestScriptProfileFunctionMeta) GetHttpTriggerUrl() *string {
	return s.HttpTriggerUrl
}

func (s *CreateScriptVersionRequestScriptProfileFunctionMeta) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateScriptVersionRequestScriptProfileFunctionMeta) SetFunctionId(v string) *CreateScriptVersionRequestScriptProfileFunctionMeta {
	s.FunctionId = &v
	return s
}

func (s *CreateScriptVersionRequestScriptProfileFunctionMeta) SetFunctionName(v string) *CreateScriptVersionRequestScriptProfileFunctionMeta {
	s.FunctionName = &v
	return s
}

func (s *CreateScriptVersionRequestScriptProfileFunctionMeta) SetHttpTriggerName(v string) *CreateScriptVersionRequestScriptProfileFunctionMeta {
	s.HttpTriggerName = &v
	return s
}

func (s *CreateScriptVersionRequestScriptProfileFunctionMeta) SetHttpTriggerUrl(v string) *CreateScriptVersionRequestScriptProfileFunctionMeta {
	s.HttpTriggerUrl = &v
	return s
}

func (s *CreateScriptVersionRequestScriptProfileFunctionMeta) SetRegionId(v string) *CreateScriptVersionRequestScriptProfileFunctionMeta {
	s.RegionId = &v
	return s
}

func (s *CreateScriptVersionRequestScriptProfileFunctionMeta) Validate() error {
	return dara.Validate(s)
}

type CreateScriptVersionRequestScriptProfileNluAccessProfile struct {
	// example:
	//
	// c2c9baae-9351-4c49-a8cb-6f24a83a8718
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s CreateScriptVersionRequestScriptProfileNluAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionRequestScriptProfileNluAccessProfile) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionRequestScriptProfileNluAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *CreateScriptVersionRequestScriptProfileNluAccessProfile) SetAccessProfileId(v string) *CreateScriptVersionRequestScriptProfileNluAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *CreateScriptVersionRequestScriptProfileNluAccessProfile) Validate() error {
	return dara.Validate(s)
}

type CreateScriptVersionRequestSynthesizerConfig struct {
	// example:
	//
	// CosyVoice
	Model            *string                                                      `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile *CreateScriptVersionRequestSynthesizerConfigNlsAccessProfile `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// example:
	//
	// 0
	PitchRate *int32                                                  `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	PronRules []*CreateScriptVersionRequestSynthesizerConfigPronRules `json:"PronRules,omitempty" xml:"PronRules,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// example:
	//
	// longanyang
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s CreateScriptVersionRequestSynthesizerConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionRequestSynthesizerConfig) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionRequestSynthesizerConfig) GetModel() *string {
	return s.Model
}

func (s *CreateScriptVersionRequestSynthesizerConfig) GetNlsAccessProfile() *CreateScriptVersionRequestSynthesizerConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *CreateScriptVersionRequestSynthesizerConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *CreateScriptVersionRequestSynthesizerConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *CreateScriptVersionRequestSynthesizerConfig) GetPitchRate() *int32 {
	return s.PitchRate
}

func (s *CreateScriptVersionRequestSynthesizerConfig) GetPronRules() []*CreateScriptVersionRequestSynthesizerConfigPronRules {
	return s.PronRules
}

func (s *CreateScriptVersionRequestSynthesizerConfig) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *CreateScriptVersionRequestSynthesizerConfig) GetVoice() *string {
	return s.Voice
}

func (s *CreateScriptVersionRequestSynthesizerConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *CreateScriptVersionRequestSynthesizerConfig) SetModel(v string) *CreateScriptVersionRequestSynthesizerConfig {
	s.Model = &v
	return s
}

func (s *CreateScriptVersionRequestSynthesizerConfig) SetNlsAccessProfile(v *CreateScriptVersionRequestSynthesizerConfigNlsAccessProfile) *CreateScriptVersionRequestSynthesizerConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *CreateScriptVersionRequestSynthesizerConfig) SetNlsAccessType(v string) *CreateScriptVersionRequestSynthesizerConfig {
	s.NlsAccessType = &v
	return s
}

func (s *CreateScriptVersionRequestSynthesizerConfig) SetNlsEngine(v string) *CreateScriptVersionRequestSynthesizerConfig {
	s.NlsEngine = &v
	return s
}

func (s *CreateScriptVersionRequestSynthesizerConfig) SetPitchRate(v int32) *CreateScriptVersionRequestSynthesizerConfig {
	s.PitchRate = &v
	return s
}

func (s *CreateScriptVersionRequestSynthesizerConfig) SetPronRules(v []*CreateScriptVersionRequestSynthesizerConfigPronRules) *CreateScriptVersionRequestSynthesizerConfig {
	s.PronRules = v
	return s
}

func (s *CreateScriptVersionRequestSynthesizerConfig) SetSpeechRate(v int32) *CreateScriptVersionRequestSynthesizerConfig {
	s.SpeechRate = &v
	return s
}

func (s *CreateScriptVersionRequestSynthesizerConfig) SetVoice(v string) *CreateScriptVersionRequestSynthesizerConfig {
	s.Voice = &v
	return s
}

func (s *CreateScriptVersionRequestSynthesizerConfig) SetVolume(v int32) *CreateScriptVersionRequestSynthesizerConfig {
	s.Volume = &v
	return s
}

func (s *CreateScriptVersionRequestSynthesizerConfig) Validate() error {
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

type CreateScriptVersionRequestSynthesizerConfigNlsAccessProfile struct {
	// example:
	//
	// c2c9baae-9351-4c49-a8cb-6f24a83a8718
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s CreateScriptVersionRequestSynthesizerConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionRequestSynthesizerConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionRequestSynthesizerConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *CreateScriptVersionRequestSynthesizerConfigNlsAccessProfile) SetAccessProfileId(v string) *CreateScriptVersionRequestSynthesizerConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *CreateScriptVersionRequestSynthesizerConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}

type CreateScriptVersionRequestSynthesizerConfigPronRules struct {
	Pattern     *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
}

func (s CreateScriptVersionRequestSynthesizerConfigPronRules) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionRequestSynthesizerConfigPronRules) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionRequestSynthesizerConfigPronRules) GetPattern() *string {
	return s.Pattern
}

func (s *CreateScriptVersionRequestSynthesizerConfigPronRules) GetReplacement() *string {
	return s.Replacement
}

func (s *CreateScriptVersionRequestSynthesizerConfigPronRules) SetPattern(v string) *CreateScriptVersionRequestSynthesizerConfigPronRules {
	s.Pattern = &v
	return s
}

func (s *CreateScriptVersionRequestSynthesizerConfigPronRules) SetReplacement(v string) *CreateScriptVersionRequestSynthesizerConfigPronRules {
	s.Replacement = &v
	return s
}

func (s *CreateScriptVersionRequestSynthesizerConfigPronRules) Validate() error {
	return dara.Validate(s)
}

type CreateScriptVersionRequestTranscriberConfig struct {
	// example:
	//
	// cd97223f-42f2-4cd9-95af-e734e2fe1fe4
	CustomizationId *string `json:"CustomizationId,omitempty" xml:"CustomizationId,omitempty"`
	// example:
	//
	// 700
	EndSilenceTimeout *int32 `json:"EndSilenceTimeout,omitempty" xml:"EndSilenceTimeout,omitempty"`
	// example:
	//
	// Paraformer
	Model            *string                                                      `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile *CreateScriptVersionRequestTranscriberConfigNlsAccessProfile `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// example:
	//
	// 0
	SpeechNoiseThreshold *int32 `json:"SpeechNoiseThreshold,omitempty" xml:"SpeechNoiseThreshold,omitempty"`
	// example:
	//
	// cd97223f-42f2-4cd9-95af-e734e2fe1fe3
	VocabularyId *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
}

func (s CreateScriptVersionRequestTranscriberConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionRequestTranscriberConfig) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionRequestTranscriberConfig) GetCustomizationId() *string {
	return s.CustomizationId
}

func (s *CreateScriptVersionRequestTranscriberConfig) GetEndSilenceTimeout() *int32 {
	return s.EndSilenceTimeout
}

func (s *CreateScriptVersionRequestTranscriberConfig) GetModel() *string {
	return s.Model
}

func (s *CreateScriptVersionRequestTranscriberConfig) GetNlsAccessProfile() *CreateScriptVersionRequestTranscriberConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *CreateScriptVersionRequestTranscriberConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *CreateScriptVersionRequestTranscriberConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *CreateScriptVersionRequestTranscriberConfig) GetSpeechNoiseThreshold() *int32 {
	return s.SpeechNoiseThreshold
}

func (s *CreateScriptVersionRequestTranscriberConfig) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *CreateScriptVersionRequestTranscriberConfig) SetCustomizationId(v string) *CreateScriptVersionRequestTranscriberConfig {
	s.CustomizationId = &v
	return s
}

func (s *CreateScriptVersionRequestTranscriberConfig) SetEndSilenceTimeout(v int32) *CreateScriptVersionRequestTranscriberConfig {
	s.EndSilenceTimeout = &v
	return s
}

func (s *CreateScriptVersionRequestTranscriberConfig) SetModel(v string) *CreateScriptVersionRequestTranscriberConfig {
	s.Model = &v
	return s
}

func (s *CreateScriptVersionRequestTranscriberConfig) SetNlsAccessProfile(v *CreateScriptVersionRequestTranscriberConfigNlsAccessProfile) *CreateScriptVersionRequestTranscriberConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *CreateScriptVersionRequestTranscriberConfig) SetNlsAccessType(v string) *CreateScriptVersionRequestTranscriberConfig {
	s.NlsAccessType = &v
	return s
}

func (s *CreateScriptVersionRequestTranscriberConfig) SetNlsEngine(v string) *CreateScriptVersionRequestTranscriberConfig {
	s.NlsEngine = &v
	return s
}

func (s *CreateScriptVersionRequestTranscriberConfig) SetSpeechNoiseThreshold(v int32) *CreateScriptVersionRequestTranscriberConfig {
	s.SpeechNoiseThreshold = &v
	return s
}

func (s *CreateScriptVersionRequestTranscriberConfig) SetVocabularyId(v string) *CreateScriptVersionRequestTranscriberConfig {
	s.VocabularyId = &v
	return s
}

func (s *CreateScriptVersionRequestTranscriberConfig) Validate() error {
	if s.NlsAccessProfile != nil {
		if err := s.NlsAccessProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateScriptVersionRequestTranscriberConfigNlsAccessProfile struct {
	// example:
	//
	// c2c9baae-9351-4c49-a8cb-6f24a83a8718
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s CreateScriptVersionRequestTranscriberConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionRequestTranscriberConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionRequestTranscriberConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *CreateScriptVersionRequestTranscriberConfigNlsAccessProfile) SetAccessProfileId(v string) *CreateScriptVersionRequestTranscriberConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *CreateScriptVersionRequestTranscriberConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}
