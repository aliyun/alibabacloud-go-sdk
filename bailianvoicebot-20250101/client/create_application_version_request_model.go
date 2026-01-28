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
	SetScriptProfile(v *CreateApplicationVersionRequestScriptProfile) *CreateApplicationVersionRequest
	GetScriptProfile() *CreateApplicationVersionRequestScriptProfile
	SetSourceVersionId(v string) *CreateApplicationVersionRequest
	GetSourceVersionId() *string
	SetSynthesizerConfig(v *CreateApplicationVersionRequestSynthesizerConfig) *CreateApplicationVersionRequest
	GetSynthesizerConfig() *CreateApplicationVersionRequestSynthesizerConfig
	SetTranscriberConfig(v *CreateApplicationVersionRequestTranscriberConfig) *CreateApplicationVersionRequest
	GetTranscriberConfig() *CreateApplicationVersionRequestTranscriberConfig
}

type CreateApplicationVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a395011f-a247-400f-bc69-28796749fd52
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId    *string                                           `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	InteractionConfig *CreateApplicationVersionRequestInteractionConfig `json:"InteractionConfig,omitempty" xml:"InteractionConfig,omitempty" type:"Struct"`
	ScriptProfile     *CreateApplicationVersionRequestScriptProfile     `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty" type:"Struct"`
	// example:
	//
	// 20904943-f711-494f-9f1f-e7f340f37707
	SourceVersionId   *string                                           `json:"SourceVersionId,omitempty" xml:"SourceVersionId,omitempty"`
	SynthesizerConfig *CreateApplicationVersionRequestSynthesizerConfig `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty" type:"Struct"`
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

func (s *CreateApplicationVersionRequest) GetScriptProfile() *CreateApplicationVersionRequestScriptProfile {
	return s.ScriptProfile
}

func (s *CreateApplicationVersionRequest) GetSourceVersionId() *string {
	return s.SourceVersionId
}

func (s *CreateApplicationVersionRequest) GetSynthesizerConfig() *CreateApplicationVersionRequestSynthesizerConfig {
	return s.SynthesizerConfig
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

type CreateApplicationVersionRequestInteractionConfig struct {
	SilenceDetectionConfig *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig `json:"SilenceDetectionConfig,omitempty" xml:"SilenceDetectionConfig,omitempty" type:"Struct"`
}

func (s CreateApplicationVersionRequestInteractionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestInteractionConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestInteractionConfig) GetSilenceDetectionConfig() *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig {
	return s.SilenceDetectionConfig
}

func (s *CreateApplicationVersionRequestInteractionConfig) SetSilenceDetectionConfig(v *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) *CreateApplicationVersionRequestInteractionConfig {
	s.SilenceDetectionConfig = v
	return s
}

func (s *CreateApplicationVersionRequestInteractionConfig) Validate() error {
	if s.SilenceDetectionConfig != nil {
		if err := s.SilenceDetectionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig struct {
	// example:
	//
	// 5
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) SetTimeout(v int32) *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig {
	s.Timeout = &v
	return s
}

func (s *CreateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationVersionRequestScriptProfile struct {
	AgentProfile *CreateApplicationVersionRequestScriptProfileAgentProfile `json:"AgentProfile,omitempty" xml:"AgentProfile,omitempty" type:"Struct"`
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s CreateApplicationVersionRequestScriptProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestScriptProfile) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestScriptProfile) GetAgentProfile() *CreateApplicationVersionRequestScriptProfileAgentProfile {
	return s.AgentProfile
}

func (s *CreateApplicationVersionRequestScriptProfile) GetModel() *string {
	return s.Model
}

func (s *CreateApplicationVersionRequestScriptProfile) SetAgentProfile(v *CreateApplicationVersionRequestScriptProfileAgentProfile) *CreateApplicationVersionRequestScriptProfile {
	s.AgentProfile = v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfile) SetModel(v string) *CreateApplicationVersionRequestScriptProfile {
	s.Model = &v
	return s
}

func (s *CreateApplicationVersionRequestScriptProfile) Validate() error {
	if s.AgentProfile != nil {
		if err := s.AgentProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApplicationVersionRequestScriptProfileAgentProfile struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PromptsJson *string `json:"PromptsJson,omitempty" xml:"PromptsJson,omitempty"`
	// example:
	//
	// SFM_PROMPTS_DEFAULT
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

type CreateApplicationVersionRequestSynthesizerConfig struct {
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// example:
	//
	// ALIYUN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// example:
	//
	// 50
	PitchRate *int32 `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// example:
	//
	// -156
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// example:
	//
	// aiqi
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s CreateApplicationVersionRequestSynthesizerConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestSynthesizerConfig) GoString() string {
	return s.String()
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

func (s *CreateApplicationVersionRequestSynthesizerConfig) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) GetVoice() *string {
	return s.Voice
}

func (s *CreateApplicationVersionRequestSynthesizerConfig) GetVolume() *int32 {
	return s.Volume
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
	return dara.Validate(s)
}

type CreateApplicationVersionRequestTranscriberConfig struct {
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// example:
	//
	// ALIYUN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
}

func (s CreateApplicationVersionRequestTranscriberConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionRequestTranscriberConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionRequestTranscriberConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *CreateApplicationVersionRequestTranscriberConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *CreateApplicationVersionRequestTranscriberConfig) SetNlsAccessType(v string) *CreateApplicationVersionRequestTranscriberConfig {
	s.NlsAccessType = &v
	return s
}

func (s *CreateApplicationVersionRequestTranscriberConfig) SetNlsEngine(v string) *CreateApplicationVersionRequestTranscriberConfig {
	s.NlsEngine = &v
	return s
}

func (s *CreateApplicationVersionRequestTranscriberConfig) Validate() error {
	return dara.Validate(s)
}
