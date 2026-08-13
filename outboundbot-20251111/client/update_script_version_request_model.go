// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScriptVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateScriptVersionRequest
	GetInstanceId() *string
	SetInteractionConfig(v *UpdateScriptVersionRequestInteractionConfig) *UpdateScriptVersionRequest
	GetInteractionConfig() *UpdateScriptVersionRequestInteractionConfig
	SetLabelConfigs(v []*UpdateScriptVersionRequestLabelConfigs) *UpdateScriptVersionRequest
	GetLabelConfigs() []*UpdateScriptVersionRequestLabelConfigs
	SetScriptId(v string) *UpdateScriptVersionRequest
	GetScriptId() *string
	SetScriptProfile(v *UpdateScriptVersionRequestScriptProfile) *UpdateScriptVersionRequest
	GetScriptProfile() *UpdateScriptVersionRequestScriptProfile
	SetSynthesizerConfig(v *UpdateScriptVersionRequestSynthesizerConfig) *UpdateScriptVersionRequest
	GetSynthesizerConfig() *UpdateScriptVersionRequestSynthesizerConfig
	SetTranscriberConfig(v *UpdateScriptVersionRequestTranscriberConfig) *UpdateScriptVersionRequest
	GetTranscriberConfig() *UpdateScriptVersionRequestTranscriberConfig
	SetVersionId(v string) *UpdateScriptVersionRequest
	GetVersionId() *string
}

type UpdateScriptVersionRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The interaction configuration.
	InteractionConfig *UpdateScriptVersionRequestInteractionConfig `json:"InteractionConfig,omitempty" xml:"InteractionConfig,omitempty" type:"Struct"`
	// The label configurations.
	LabelConfigs []*UpdateScriptVersionRequestLabelConfigs `json:"LabelConfigs,omitempty" xml:"LabelConfigs,omitempty" type:"Repeated"`
	// The scenario ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The dialogue capability configuration.
	ScriptProfile *UpdateScriptVersionRequestScriptProfile `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty" type:"Struct"`
	// The TTS configuration.
	SynthesizerConfig *UpdateScriptVersionRequestSynthesizerConfig `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty" type:"Struct"`
	// The ASR configuration.
	TranscriberConfig *UpdateScriptVersionRequestTranscriberConfig `json:"TranscriberConfig,omitempty" xml:"TranscriberConfig,omitempty" type:"Struct"`
	// The version ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b26
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s UpdateScriptVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateScriptVersionRequest) GetInteractionConfig() *UpdateScriptVersionRequestInteractionConfig {
	return s.InteractionConfig
}

func (s *UpdateScriptVersionRequest) GetLabelConfigs() []*UpdateScriptVersionRequestLabelConfigs {
	return s.LabelConfigs
}

func (s *UpdateScriptVersionRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *UpdateScriptVersionRequest) GetScriptProfile() *UpdateScriptVersionRequestScriptProfile {
	return s.ScriptProfile
}

func (s *UpdateScriptVersionRequest) GetSynthesizerConfig() *UpdateScriptVersionRequestSynthesizerConfig {
	return s.SynthesizerConfig
}

func (s *UpdateScriptVersionRequest) GetTranscriberConfig() *UpdateScriptVersionRequestTranscriberConfig {
	return s.TranscriberConfig
}

func (s *UpdateScriptVersionRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *UpdateScriptVersionRequest) SetInstanceId(v string) *UpdateScriptVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateScriptVersionRequest) SetInteractionConfig(v *UpdateScriptVersionRequestInteractionConfig) *UpdateScriptVersionRequest {
	s.InteractionConfig = v
	return s
}

func (s *UpdateScriptVersionRequest) SetLabelConfigs(v []*UpdateScriptVersionRequestLabelConfigs) *UpdateScriptVersionRequest {
	s.LabelConfigs = v
	return s
}

func (s *UpdateScriptVersionRequest) SetScriptId(v string) *UpdateScriptVersionRequest {
	s.ScriptId = &v
	return s
}

func (s *UpdateScriptVersionRequest) SetScriptProfile(v *UpdateScriptVersionRequestScriptProfile) *UpdateScriptVersionRequest {
	s.ScriptProfile = v
	return s
}

func (s *UpdateScriptVersionRequest) SetSynthesizerConfig(v *UpdateScriptVersionRequestSynthesizerConfig) *UpdateScriptVersionRequest {
	s.SynthesizerConfig = v
	return s
}

func (s *UpdateScriptVersionRequest) SetTranscriberConfig(v *UpdateScriptVersionRequestTranscriberConfig) *UpdateScriptVersionRequest {
	s.TranscriberConfig = v
	return s
}

func (s *UpdateScriptVersionRequest) SetVersionId(v string) *UpdateScriptVersionRequest {
	s.VersionId = &v
	return s
}

func (s *UpdateScriptVersionRequest) Validate() error {
	if s.InteractionConfig != nil {
		if err := s.InteractionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LabelConfigs != nil {
		for _, item := range s.LabelConfigs {
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

type UpdateScriptVersionRequestInteractionConfig struct {
	// The background music ID.
	//
	// example:
	//
	// office-ambience
	BackgroundMusicId *string `json:"BackgroundMusicId,omitempty" xml:"BackgroundMusicId,omitempty"`
	// The barge-in configuration.
	BargeInConfig *UpdateScriptVersionRequestInteractionConfigBargeInConfig `json:"BargeInConfig,omitempty" xml:"BargeInConfig,omitempty" type:"Struct"`
	// The hang-up configuration.
	EndConversationConfig *UpdateScriptVersionRequestInteractionConfigEndConversationConfig `json:"EndConversationConfig,omitempty" xml:"EndConversationConfig,omitempty" type:"Struct"`
	// The delay in milliseconds before playing audio after the call is connected.
	//
	// example:
	//
	// 2000
	InitialGreetingDelayMilliseconds *int32 `json:"InitialGreetingDelayMilliseconds,omitempty" xml:"InitialGreetingDelayMilliseconds,omitempty"`
	// The silence detection configuration.
	SilenceDetectionConfig *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig `json:"SilenceDetectionConfig,omitempty" xml:"SilenceDetectionConfig,omitempty" type:"Struct"`
	// The transition phrase model configuration.
	TransitionConfig *UpdateScriptVersionRequestInteractionConfigTransitionConfig `json:"TransitionConfig,omitempty" xml:"TransitionConfig,omitempty" type:"Struct"`
}

func (s UpdateScriptVersionRequestInteractionConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestInteractionConfig) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestInteractionConfig) GetBackgroundMusicId() *string {
	return s.BackgroundMusicId
}

func (s *UpdateScriptVersionRequestInteractionConfig) GetBargeInConfig() *UpdateScriptVersionRequestInteractionConfigBargeInConfig {
	return s.BargeInConfig
}

func (s *UpdateScriptVersionRequestInteractionConfig) GetEndConversationConfig() *UpdateScriptVersionRequestInteractionConfigEndConversationConfig {
	return s.EndConversationConfig
}

func (s *UpdateScriptVersionRequestInteractionConfig) GetInitialGreetingDelayMilliseconds() *int32 {
	return s.InitialGreetingDelayMilliseconds
}

func (s *UpdateScriptVersionRequestInteractionConfig) GetSilenceDetectionConfig() *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig {
	return s.SilenceDetectionConfig
}

func (s *UpdateScriptVersionRequestInteractionConfig) GetTransitionConfig() *UpdateScriptVersionRequestInteractionConfigTransitionConfig {
	return s.TransitionConfig
}

func (s *UpdateScriptVersionRequestInteractionConfig) SetBackgroundMusicId(v string) *UpdateScriptVersionRequestInteractionConfig {
	s.BackgroundMusicId = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfig) SetBargeInConfig(v *UpdateScriptVersionRequestInteractionConfigBargeInConfig) *UpdateScriptVersionRequestInteractionConfig {
	s.BargeInConfig = v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfig) SetEndConversationConfig(v *UpdateScriptVersionRequestInteractionConfigEndConversationConfig) *UpdateScriptVersionRequestInteractionConfig {
	s.EndConversationConfig = v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfig) SetInitialGreetingDelayMilliseconds(v int32) *UpdateScriptVersionRequestInteractionConfig {
	s.InitialGreetingDelayMilliseconds = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfig) SetSilenceDetectionConfig(v *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig) *UpdateScriptVersionRequestInteractionConfig {
	s.SilenceDetectionConfig = v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfig) SetTransitionConfig(v *UpdateScriptVersionRequestInteractionConfigTransitionConfig) *UpdateScriptVersionRequestInteractionConfig {
	s.TransitionConfig = v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfig) Validate() error {
	if s.BargeInConfig != nil {
		if err := s.BargeInConfig.Validate(); err != nil {
			return err
		}
	}
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
	if s.TransitionConfig != nil {
		if err := s.TransitionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateScriptVersionRequestInteractionConfigBargeInConfig struct {
	// Specifies whether barge-in is supported during the closing statement.
	//
	// example:
	//
	// true
	ClosingBargeInEnabled *bool `json:"ClosingBargeInEnabled,omitempty" xml:"ClosingBargeInEnabled,omitempty"`
	// Specifies whether barge-in is supported during the conversation.
	//
	// example:
	//
	// true
	GlobalBargeInEnabled *bool `json:"GlobalBargeInEnabled,omitempty" xml:"GlobalBargeInEnabled,omitempty"`
	// Specifies whether barge-in is supported during the opening greeting.
	//
	// example:
	//
	// true
	OpeningBargeInEnabled *bool `json:"OpeningBargeInEnabled,omitempty" xml:"OpeningBargeInEnabled,omitempty"`
}

func (s UpdateScriptVersionRequestInteractionConfigBargeInConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestInteractionConfigBargeInConfig) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestInteractionConfigBargeInConfig) GetClosingBargeInEnabled() *bool {
	return s.ClosingBargeInEnabled
}

func (s *UpdateScriptVersionRequestInteractionConfigBargeInConfig) GetGlobalBargeInEnabled() *bool {
	return s.GlobalBargeInEnabled
}

func (s *UpdateScriptVersionRequestInteractionConfigBargeInConfig) GetOpeningBargeInEnabled() *bool {
	return s.OpeningBargeInEnabled
}

func (s *UpdateScriptVersionRequestInteractionConfigBargeInConfig) SetClosingBargeInEnabled(v bool) *UpdateScriptVersionRequestInteractionConfigBargeInConfig {
	s.ClosingBargeInEnabled = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigBargeInConfig) SetGlobalBargeInEnabled(v bool) *UpdateScriptVersionRequestInteractionConfigBargeInConfig {
	s.GlobalBargeInEnabled = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigBargeInConfig) SetOpeningBargeInEnabled(v bool) *UpdateScriptVersionRequestInteractionConfigBargeInConfig {
	s.OpeningBargeInEnabled = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigBargeInConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateScriptVersionRequestInteractionConfigEndConversationConfig struct {
	// Specifies whether barge-in is supported during the delayed hang-up waiting period.
	//
	// example:
	//
	// true
	BargeInEnabled *bool `json:"BargeInEnabled,omitempty" xml:"BargeInEnabled,omitempty"`
	// The number of seconds to wait after the hang-up script finishes playing before executing the hang-up action. Valid values: 0 to 5.
	//
	// example:
	//
	// 1
	Delay *int32 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The special condition interception configuration.
	Triggers []*UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s UpdateScriptVersionRequestInteractionConfigEndConversationConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestInteractionConfigEndConversationConfig) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfig) GetBargeInEnabled() *bool {
	return s.BargeInEnabled
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfig) GetDelay() *int32 {
	return s.Delay
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfig) GetTriggers() []*UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers {
	return s.Triggers
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfig) SetBargeInEnabled(v bool) *UpdateScriptVersionRequestInteractionConfigEndConversationConfig {
	s.BargeInEnabled = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfig) SetDelay(v int32) *UpdateScriptVersionRequestInteractionConfigEndConversationConfig {
	s.Delay = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfig) SetTriggers(v []*UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) *UpdateScriptVersionRequestInteractionConfigEndConversationConfig {
	s.Triggers = v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfig) Validate() error {
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

type UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers struct {
	// The closing script to play when the turn limit is reached and hang-up is executed.
	//
	// example:
	//
	// Thank you for answering the call. Have a nice day. Goodbye!
	ClosingStatement *string `json:"ClosingStatement,omitempty" xml:"ClosingStatement,omitempty"`
	// The list of custom interception keywords.
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	// Valid values:
	//
	// - TurnLimit: maximum interaction turn limit check.
	//
	// example:
	//
	// TurnLimit
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The maximum number of interaction turns before executing hang-up. Valid values: 0 to 100. A value of 0 indicates that the turn-limit hang-up is not enabled.
	//
	// example:
	//
	// 20
	TurnLimit *int32 `json:"TurnLimit,omitempty" xml:"TurnLimit,omitempty"`
}

func (s UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) GetClosingStatement() *string {
	return s.ClosingStatement
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) GetKeywords() []*string {
	return s.Keywords
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) GetTriggerType() *string {
	return s.TriggerType
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) GetTurnLimit() *int32 {
	return s.TurnLimit
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) SetClosingStatement(v string) *UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.ClosingStatement = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) SetKeywords(v []*string) *UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.Keywords = v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) SetTriggerType(v string) *UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.TriggerType = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) SetTurnLimit(v int32) *UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers {
	s.TurnLimit = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigEndConversationConfigTriggers) Validate() error {
	return dara.Validate(s)
}

type UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig struct {
	// The list of actions to perform during consecutive silence.
	FallbackControlParamsList []*UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfigFallbackControlParamsList `json:"FallbackControlParamsList,omitempty" xml:"FallbackControlParamsList,omitempty" type:"Repeated"`
	// The number of consecutive silence rounds before hanging up.
	//
	// example:
	//
	// 3
	MaxRepeats *int32 `json:"MaxRepeats,omitempty" xml:"MaxRepeats,omitempty"`
	// The silence prompt.
	//
	// example:
	//
	// - Repeat the content of the previous conversation round
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The silence timeout period, in milliseconds. When the user remains silent for longer than the specified value, the silence timeout prompt is played. Valid range: 2000 to 10000.
	//
	// example:
	//
	// 5000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig) GetFallbackControlParamsList() []*UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfigFallbackControlParamsList {
	return s.FallbackControlParamsList
}

func (s *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig) GetMaxRepeats() *int32 {
	return s.MaxRepeats
}

func (s *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig) GetPrompt() *string {
	return s.Prompt
}

func (s *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig) SetFallbackControlParamsList(v []*UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfigFallbackControlParamsList) *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig {
	s.FallbackControlParamsList = v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig) SetMaxRepeats(v int32) *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig {
	s.MaxRepeats = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig) SetPrompt(v string) *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig {
	s.Prompt = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig) SetTimeout(v int32) *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig {
	s.Timeout = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfig) Validate() error {
	if s.FallbackControlParamsList != nil {
		for _, item := range s.FallbackControlParamsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfigFallbackControlParamsList struct {
	// The action to perform during consecutive silence.
	//
	// example:
	//
	// HangUp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfigFallbackControlParamsList) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfigFallbackControlParamsList) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfigFallbackControlParamsList) GetType() *string {
	return s.Type
}

func (s *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfigFallbackControlParamsList) SetType(v string) *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfigFallbackControlParamsList {
	s.Type = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigSilenceDetectionConfigFallbackControlParamsList) Validate() error {
	return dara.Validate(s)
}

type UpdateScriptVersionRequestInteractionConfigTransitionConfig struct {
	// The prompt for model-generated transition phrases.
	//
	// example:
	//
	// Based on the user\\"s latest reply in the conversation history below, generate a brief transitional phrase for the customer service agent to naturally and smoothly connect the conversation. Requirements are as follows:
	//
	// 1. Use colloquial expressions common in customer service scenarios, maintaining a natural, polite, and neutral tone......
	AiPhrasePrompt *string `json:"AiPhrasePrompt,omitempty" xml:"AiPhrasePrompt,omitempty"`
	// The list of fixed transition phrases.
	FixedPhraseList []*string `json:"FixedPhraseList,omitempty" xml:"FixedPhraseList,omitempty" type:"Repeated"`
	// The method for generating transition phrases.
	//
	// example:
	//
	// aiGenerated
	PhraseSource *string `json:"PhraseSource,omitempty" xml:"PhraseSource,omitempty"`
	// Specifies whether to enable transition phrases.
	//
	// example:
	//
	// true
	TransitionSwitch *bool `json:"TransitionSwitch,omitempty" xml:"TransitionSwitch,omitempty"`
}

func (s UpdateScriptVersionRequestInteractionConfigTransitionConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestInteractionConfigTransitionConfig) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestInteractionConfigTransitionConfig) GetAiPhrasePrompt() *string {
	return s.AiPhrasePrompt
}

func (s *UpdateScriptVersionRequestInteractionConfigTransitionConfig) GetFixedPhraseList() []*string {
	return s.FixedPhraseList
}

func (s *UpdateScriptVersionRequestInteractionConfigTransitionConfig) GetPhraseSource() *string {
	return s.PhraseSource
}

func (s *UpdateScriptVersionRequestInteractionConfigTransitionConfig) GetTransitionSwitch() *bool {
	return s.TransitionSwitch
}

func (s *UpdateScriptVersionRequestInteractionConfigTransitionConfig) SetAiPhrasePrompt(v string) *UpdateScriptVersionRequestInteractionConfigTransitionConfig {
	s.AiPhrasePrompt = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigTransitionConfig) SetFixedPhraseList(v []*string) *UpdateScriptVersionRequestInteractionConfigTransitionConfig {
	s.FixedPhraseList = v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigTransitionConfig) SetPhraseSource(v string) *UpdateScriptVersionRequestInteractionConfigTransitionConfig {
	s.PhraseSource = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigTransitionConfig) SetTransitionSwitch(v bool) *UpdateScriptVersionRequestInteractionConfigTransitionConfig {
	s.TransitionSwitch = &v
	return s
}

func (s *UpdateScriptVersionRequestInteractionConfigTransitionConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateScriptVersionRequestLabelConfigs struct {
	// The candidate values for the label.
	CandidateValues []*string `json:"CandidateValues,omitempty" xml:"CandidateValues,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// Describes whether the user is satisfied with this service
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label name.
	//
	// example:
	//
	// Satisfaction
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateScriptVersionRequestLabelConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestLabelConfigs) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestLabelConfigs) GetCandidateValues() []*string {
	return s.CandidateValues
}

func (s *UpdateScriptVersionRequestLabelConfigs) GetDescription() *string {
	return s.Description
}

func (s *UpdateScriptVersionRequestLabelConfigs) GetName() *string {
	return s.Name
}

func (s *UpdateScriptVersionRequestLabelConfigs) SetCandidateValues(v []*string) *UpdateScriptVersionRequestLabelConfigs {
	s.CandidateValues = v
	return s
}

func (s *UpdateScriptVersionRequestLabelConfigs) SetDescription(v string) *UpdateScriptVersionRequestLabelConfigs {
	s.Description = &v
	return s
}

func (s *UpdateScriptVersionRequestLabelConfigs) SetName(v string) *UpdateScriptVersionRequestLabelConfigs {
	s.Name = &v
	return s
}

func (s *UpdateScriptVersionRequestLabelConfigs) Validate() error {
	return dara.Validate(s)
}

type UpdateScriptVersionRequestScriptProfile struct {
	// The chatbot AgentKey.
	//
	// example:
	//
	// 1309723684579735_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The dialogue agent configuration.
	AgentProfile *UpdateScriptVersionRequestScriptProfileAgentProfile `json:"AgentProfile,omitempty" xml:"AgentProfile,omitempty" type:"Struct"`
	// The chatbot type.
	//
	// example:
	//
	// LITE
	BuilderType *string `json:"BuilderType,omitempty" xml:"BuilderType,omitempty"`
	// The chatbot ID.
	//
	// example:
	//
	// chatbot-cn-MQuyjjb666
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// The function compute configuration.
	FunctionMeta *UpdateScriptVersionRequestScriptProfileFunctionMeta `json:"FunctionMeta,omitempty" xml:"FunctionMeta,omitempty" type:"Struct"`
	// The dialogue model.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The associated configuration.
	NluAccessProfile *UpdateScriptVersionRequestScriptProfileNluAccessProfile `json:"NluAccessProfile,omitempty" xml:"NluAccessProfile,omitempty" type:"Struct"`
	// The dialogue model invocation method.
	//
	// example:
	//
	// MANAGED
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// Specifies whether the model is an Omni model.
	//
	// example:
	//
	// true
	OmniModel *bool `json:"OmniModel,omitempty" xml:"OmniModel,omitempty"`
}

func (s UpdateScriptVersionRequestScriptProfile) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestScriptProfile) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestScriptProfile) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateScriptVersionRequestScriptProfile) GetAgentProfile() *UpdateScriptVersionRequestScriptProfileAgentProfile {
	return s.AgentProfile
}

func (s *UpdateScriptVersionRequestScriptProfile) GetBuilderType() *string {
	return s.BuilderType
}

func (s *UpdateScriptVersionRequestScriptProfile) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *UpdateScriptVersionRequestScriptProfile) GetFunctionMeta() *UpdateScriptVersionRequestScriptProfileFunctionMeta {
	return s.FunctionMeta
}

func (s *UpdateScriptVersionRequestScriptProfile) GetModel() *string {
	return s.Model
}

func (s *UpdateScriptVersionRequestScriptProfile) GetNluAccessProfile() *UpdateScriptVersionRequestScriptProfileNluAccessProfile {
	return s.NluAccessProfile
}

func (s *UpdateScriptVersionRequestScriptProfile) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *UpdateScriptVersionRequestScriptProfile) GetOmniModel() *bool {
	return s.OmniModel
}

func (s *UpdateScriptVersionRequestScriptProfile) SetAgentKey(v string) *UpdateScriptVersionRequestScriptProfile {
	s.AgentKey = &v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfile) SetAgentProfile(v *UpdateScriptVersionRequestScriptProfileAgentProfile) *UpdateScriptVersionRequestScriptProfile {
	s.AgentProfile = v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfile) SetBuilderType(v string) *UpdateScriptVersionRequestScriptProfile {
	s.BuilderType = &v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfile) SetChatbotId(v string) *UpdateScriptVersionRequestScriptProfile {
	s.ChatbotId = &v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfile) SetFunctionMeta(v *UpdateScriptVersionRequestScriptProfileFunctionMeta) *UpdateScriptVersionRequestScriptProfile {
	s.FunctionMeta = v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfile) SetModel(v string) *UpdateScriptVersionRequestScriptProfile {
	s.Model = &v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfile) SetNluAccessProfile(v *UpdateScriptVersionRequestScriptProfileNluAccessProfile) *UpdateScriptVersionRequestScriptProfile {
	s.NluAccessProfile = v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfile) SetNluAccessType(v string) *UpdateScriptVersionRequestScriptProfile {
	s.NluAccessType = &v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfile) SetOmniModel(v bool) *UpdateScriptVersionRequestScriptProfile {
	s.OmniModel = &v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfile) Validate() error {
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

type UpdateScriptVersionRequestScriptProfileAgentProfile struct {
	// The prompt in JSON format.
	//
	// example:
	//
	// {\\"prompts\\":\\"I am a chatbot.\\"}
	PromptsJson *string `json:"PromptsJson,omitempty" xml:"PromptsJson,omitempty"`
	// The scenario template ID.
	//
	// example:
	//
	// OUTBOUND_BOT_PROMPTS_DEFAULT
	ScriptProfileTemplateId *string `json:"ScriptProfileTemplateId,omitempty" xml:"ScriptProfileTemplateId,omitempty"`
}

func (s UpdateScriptVersionRequestScriptProfileAgentProfile) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestScriptProfileAgentProfile) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestScriptProfileAgentProfile) GetPromptsJson() *string {
	return s.PromptsJson
}

func (s *UpdateScriptVersionRequestScriptProfileAgentProfile) GetScriptProfileTemplateId() *string {
	return s.ScriptProfileTemplateId
}

func (s *UpdateScriptVersionRequestScriptProfileAgentProfile) SetPromptsJson(v string) *UpdateScriptVersionRequestScriptProfileAgentProfile {
	s.PromptsJson = &v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfileAgentProfile) SetScriptProfileTemplateId(v string) *UpdateScriptVersionRequestScriptProfileAgentProfile {
	s.ScriptProfileTemplateId = &v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfileAgentProfile) Validate() error {
	return dara.Validate(s)
}

type UpdateScriptVersionRequestScriptProfileFunctionMeta struct {
	// The function service ID.
	//
	// example:
	//
	// 9b752bbb-805a-4d3e-9013-eab5555c3fef
	FunctionId *string `json:"FunctionId,omitempty" xml:"FunctionId,omitempty"`
	// The function service name.
	//
	// example:
	//
	// my_funciton
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The function trigger name.
	//
	// example:
	//
	// defaultTrigger
	HttpTriggerName *string `json:"HttpTriggerName,omitempty" xml:"HttpTriggerName,omitempty"`
	// The function trigger URL.
	//
	// example:
	//
	// http://chat-xxxxx-v-yewiundukb.cn-hangzhou-xxx.run
	HttpTriggerUrl *string `json:"HttpTriggerUrl,omitempty" xml:"HttpTriggerUrl,omitempty"`
	// The region where the function service resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateScriptVersionRequestScriptProfileFunctionMeta) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestScriptProfileFunctionMeta) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestScriptProfileFunctionMeta) GetFunctionId() *string {
	return s.FunctionId
}

func (s *UpdateScriptVersionRequestScriptProfileFunctionMeta) GetFunctionName() *string {
	return s.FunctionName
}

func (s *UpdateScriptVersionRequestScriptProfileFunctionMeta) GetHttpTriggerName() *string {
	return s.HttpTriggerName
}

func (s *UpdateScriptVersionRequestScriptProfileFunctionMeta) GetHttpTriggerUrl() *string {
	return s.HttpTriggerUrl
}

func (s *UpdateScriptVersionRequestScriptProfileFunctionMeta) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateScriptVersionRequestScriptProfileFunctionMeta) SetFunctionId(v string) *UpdateScriptVersionRequestScriptProfileFunctionMeta {
	s.FunctionId = &v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfileFunctionMeta) SetFunctionName(v string) *UpdateScriptVersionRequestScriptProfileFunctionMeta {
	s.FunctionName = &v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfileFunctionMeta) SetHttpTriggerName(v string) *UpdateScriptVersionRequestScriptProfileFunctionMeta {
	s.HttpTriggerName = &v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfileFunctionMeta) SetHttpTriggerUrl(v string) *UpdateScriptVersionRequestScriptProfileFunctionMeta {
	s.HttpTriggerUrl = &v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfileFunctionMeta) SetRegionId(v string) *UpdateScriptVersionRequestScriptProfileFunctionMeta {
	s.RegionId = &v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfileFunctionMeta) Validate() error {
	return dara.Validate(s)
}

type UpdateScriptVersionRequestScriptProfileNluAccessProfile struct {
	// The third-party dialogue model configuration ID.
	//
	// example:
	//
	// c2c9baae-9351-4c49-a8cb-6f24a83a8718
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s UpdateScriptVersionRequestScriptProfileNluAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestScriptProfileNluAccessProfile) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestScriptProfileNluAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *UpdateScriptVersionRequestScriptProfileNluAccessProfile) SetAccessProfileId(v string) *UpdateScriptVersionRequestScriptProfileNluAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *UpdateScriptVersionRequestScriptProfileNluAccessProfile) Validate() error {
	return dara.Validate(s)
}

type UpdateScriptVersionRequestSynthesizerConfig struct {
	// The TTS model.
	//
	// example:
	//
	// CosyVoice
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The associated configuration.
	NlsAccessProfile *UpdateScriptVersionRequestSynthesizerConfigNlsAccessProfile `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	// The TTS invocation method.
	//
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// The TTS engine.
	//
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// The pitch rate.
	//
	// example:
	//
	// 0
	PitchRate *int32 `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// The TTS correction dictionary.
	PronRules []*UpdateScriptVersionRequestSynthesizerConfigPronRules `json:"PronRules,omitempty" xml:"PronRules,omitempty" type:"Repeated"`
	// The speech rate.
	//
	// example:
	//
	// 0
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// The voice.
	//
	// example:
	//
	// longanyang
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// The volume.
	//
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s UpdateScriptVersionRequestSynthesizerConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestSynthesizerConfig) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) GetModel() *string {
	return s.Model
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) GetNlsAccessProfile() *UpdateScriptVersionRequestSynthesizerConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) GetPitchRate() *int32 {
	return s.PitchRate
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) GetPronRules() []*UpdateScriptVersionRequestSynthesizerConfigPronRules {
	return s.PronRules
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) GetVoice() *string {
	return s.Voice
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) SetModel(v string) *UpdateScriptVersionRequestSynthesizerConfig {
	s.Model = &v
	return s
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) SetNlsAccessProfile(v *UpdateScriptVersionRequestSynthesizerConfigNlsAccessProfile) *UpdateScriptVersionRequestSynthesizerConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) SetNlsAccessType(v string) *UpdateScriptVersionRequestSynthesizerConfig {
	s.NlsAccessType = &v
	return s
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) SetNlsEngine(v string) *UpdateScriptVersionRequestSynthesizerConfig {
	s.NlsEngine = &v
	return s
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) SetPitchRate(v int32) *UpdateScriptVersionRequestSynthesizerConfig {
	s.PitchRate = &v
	return s
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) SetPronRules(v []*UpdateScriptVersionRequestSynthesizerConfigPronRules) *UpdateScriptVersionRequestSynthesizerConfig {
	s.PronRules = v
	return s
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) SetSpeechRate(v int32) *UpdateScriptVersionRequestSynthesizerConfig {
	s.SpeechRate = &v
	return s
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) SetVoice(v string) *UpdateScriptVersionRequestSynthesizerConfig {
	s.Voice = &v
	return s
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) SetVolume(v int32) *UpdateScriptVersionRequestSynthesizerConfig {
	s.Volume = &v
	return s
}

func (s *UpdateScriptVersionRequestSynthesizerConfig) Validate() error {
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

type UpdateScriptVersionRequestSynthesizerConfigNlsAccessProfile struct {
	// The third-party speech configuration ID. This parameter is required when you use a third-party ASR service such as Doubao or iFLYTEK.
	//
	// example:
	//
	// c2c9baae-9351-4c49-a8cb-6f24a83a8718
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s UpdateScriptVersionRequestSynthesizerConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestSynthesizerConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestSynthesizerConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *UpdateScriptVersionRequestSynthesizerConfigNlsAccessProfile) SetAccessProfileId(v string) *UpdateScriptVersionRequestSynthesizerConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *UpdateScriptVersionRequestSynthesizerConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}

type UpdateScriptVersionRequestSynthesizerConfigPronRules struct {
	// The easily mispronounced word or phrase.
	//
	// example:
	//
	// 还钱
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The homophonic word or phrase.
	//
	// example:
	//
	// 环钱
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
}

func (s UpdateScriptVersionRequestSynthesizerConfigPronRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestSynthesizerConfigPronRules) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestSynthesizerConfigPronRules) GetPattern() *string {
	return s.Pattern
}

func (s *UpdateScriptVersionRequestSynthesizerConfigPronRules) GetReplacement() *string {
	return s.Replacement
}

func (s *UpdateScriptVersionRequestSynthesizerConfigPronRules) SetPattern(v string) *UpdateScriptVersionRequestSynthesizerConfigPronRules {
	s.Pattern = &v
	return s
}

func (s *UpdateScriptVersionRequestSynthesizerConfigPronRules) SetReplacement(v string) *UpdateScriptVersionRequestSynthesizerConfigPronRules {
	s.Replacement = &v
	return s
}

func (s *UpdateScriptVersionRequestSynthesizerConfigPronRules) Validate() error {
	return dara.Validate(s)
}

type UpdateScriptVersionRequestTranscriberConfig struct {
	// The ASR correction dictionary.
	CorrectionRules []*UpdateScriptVersionRequestTranscriberConfigCorrectionRules `json:"CorrectionRules,omitempty" xml:"CorrectionRules,omitempty" type:"Repeated"`
	// The custom language model ID for ASR.
	//
	// example:
	//
	// cd97223f-42f2-4cd9-95af-e734e2fe1472
	CustomizationId *string `json:"CustomizationId,omitempty" xml:"CustomizationId,omitempty"`
	// The silence detection threshold. Sentence segmentation is triggered when the speaking interval exceeds x milliseconds, also known as Voice Activity Detection (VAD).
	//
	// example:
	//
	// 700
	EndSilenceTimeout *int32 `json:"EndSilenceTimeout,omitempty" xml:"EndSilenceTimeout,omitempty"`
	// The ASR model.
	//
	// example:
	//
	// Paraformer
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The associated configuration.
	NlsAccessProfile *UpdateScriptVersionRequestTranscriberConfigNlsAccessProfile `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	// The ASR invocation method.
	//
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// The ASR engine.
	//
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// The noise threshold. Valid values: -100 to 100.
	//
	// A value closer to -100 increases the probability that noise is recognized as speech.
	//
	// A value closer to +100 increases the probability that speech is recognized as noise.
	//
	// example:
	//
	// 0
	SpeechNoiseThreshold *int32 `json:"SpeechNoiseThreshold,omitempty" xml:"SpeechNoiseThreshold,omitempty"`
	// The hot word list ID. You can obtain this ID from the hot word management page.
	//
	// example:
	//
	// cd97223f-42f2-4cd9-95af-e734e2fe1fe3
	VocabularyId *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
}

func (s UpdateScriptVersionRequestTranscriberConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestTranscriberConfig) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestTranscriberConfig) GetCorrectionRules() []*UpdateScriptVersionRequestTranscriberConfigCorrectionRules {
	return s.CorrectionRules
}

func (s *UpdateScriptVersionRequestTranscriberConfig) GetCustomizationId() *string {
	return s.CustomizationId
}

func (s *UpdateScriptVersionRequestTranscriberConfig) GetEndSilenceTimeout() *int32 {
	return s.EndSilenceTimeout
}

func (s *UpdateScriptVersionRequestTranscriberConfig) GetModel() *string {
	return s.Model
}

func (s *UpdateScriptVersionRequestTranscriberConfig) GetNlsAccessProfile() *UpdateScriptVersionRequestTranscriberConfigNlsAccessProfile {
	return s.NlsAccessProfile
}

func (s *UpdateScriptVersionRequestTranscriberConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *UpdateScriptVersionRequestTranscriberConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *UpdateScriptVersionRequestTranscriberConfig) GetSpeechNoiseThreshold() *int32 {
	return s.SpeechNoiseThreshold
}

func (s *UpdateScriptVersionRequestTranscriberConfig) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *UpdateScriptVersionRequestTranscriberConfig) SetCorrectionRules(v []*UpdateScriptVersionRequestTranscriberConfigCorrectionRules) *UpdateScriptVersionRequestTranscriberConfig {
	s.CorrectionRules = v
	return s
}

func (s *UpdateScriptVersionRequestTranscriberConfig) SetCustomizationId(v string) *UpdateScriptVersionRequestTranscriberConfig {
	s.CustomizationId = &v
	return s
}

func (s *UpdateScriptVersionRequestTranscriberConfig) SetEndSilenceTimeout(v int32) *UpdateScriptVersionRequestTranscriberConfig {
	s.EndSilenceTimeout = &v
	return s
}

func (s *UpdateScriptVersionRequestTranscriberConfig) SetModel(v string) *UpdateScriptVersionRequestTranscriberConfig {
	s.Model = &v
	return s
}

func (s *UpdateScriptVersionRequestTranscriberConfig) SetNlsAccessProfile(v *UpdateScriptVersionRequestTranscriberConfigNlsAccessProfile) *UpdateScriptVersionRequestTranscriberConfig {
	s.NlsAccessProfile = v
	return s
}

func (s *UpdateScriptVersionRequestTranscriberConfig) SetNlsAccessType(v string) *UpdateScriptVersionRequestTranscriberConfig {
	s.NlsAccessType = &v
	return s
}

func (s *UpdateScriptVersionRequestTranscriberConfig) SetNlsEngine(v string) *UpdateScriptVersionRequestTranscriberConfig {
	s.NlsEngine = &v
	return s
}

func (s *UpdateScriptVersionRequestTranscriberConfig) SetSpeechNoiseThreshold(v int32) *UpdateScriptVersionRequestTranscriberConfig {
	s.SpeechNoiseThreshold = &v
	return s
}

func (s *UpdateScriptVersionRequestTranscriberConfig) SetVocabularyId(v string) *UpdateScriptVersionRequestTranscriberConfig {
	s.VocabularyId = &v
	return s
}

func (s *UpdateScriptVersionRequestTranscriberConfig) Validate() error {
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

type UpdateScriptVersionRequestTranscriberConfigCorrectionRules struct {
	// The incorrectly recognized text.
	//
	// example:
	//
	// Aliabba
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The corrected text.
	//
	// example:
	//
	// Alibaba
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
}

func (s UpdateScriptVersionRequestTranscriberConfigCorrectionRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestTranscriberConfigCorrectionRules) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestTranscriberConfigCorrectionRules) GetPattern() *string {
	return s.Pattern
}

func (s *UpdateScriptVersionRequestTranscriberConfigCorrectionRules) GetReplacement() *string {
	return s.Replacement
}

func (s *UpdateScriptVersionRequestTranscriberConfigCorrectionRules) SetPattern(v string) *UpdateScriptVersionRequestTranscriberConfigCorrectionRules {
	s.Pattern = &v
	return s
}

func (s *UpdateScriptVersionRequestTranscriberConfigCorrectionRules) SetReplacement(v string) *UpdateScriptVersionRequestTranscriberConfigCorrectionRules {
	s.Replacement = &v
	return s
}

func (s *UpdateScriptVersionRequestTranscriberConfigCorrectionRules) Validate() error {
	return dara.Validate(s)
}

type UpdateScriptVersionRequestTranscriberConfigNlsAccessProfile struct {
	// The third-party speech configuration ID. This parameter is required when you use a third-party ASR service such as Doubao or iFLYTEK.
	//
	// example:
	//
	// c2c9baae-9351-4c49-a8cb-6f24a83a8718
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
}

func (s UpdateScriptVersionRequestTranscriberConfigNlsAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionRequestTranscriberConfigNlsAccessProfile) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionRequestTranscriberConfigNlsAccessProfile) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *UpdateScriptVersionRequestTranscriberConfigNlsAccessProfile) SetAccessProfileId(v string) *UpdateScriptVersionRequestTranscriberConfigNlsAccessProfile {
	s.AccessProfileId = &v
	return s
}

func (s *UpdateScriptVersionRequestTranscriberConfigNlsAccessProfile) Validate() error {
	return dara.Validate(s)
}
