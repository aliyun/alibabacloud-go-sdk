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
	SetRagConfig(v *UpdateApplicationVersionRequestRagConfig) *UpdateApplicationVersionRequest
	GetRagConfig() *UpdateApplicationVersionRequestRagConfig
	SetScriptProfile(v *UpdateApplicationVersionRequestScriptProfile) *UpdateApplicationVersionRequest
	GetScriptProfile() *UpdateApplicationVersionRequestScriptProfile
	SetSynthesizerConfig(v *UpdateApplicationVersionRequestSynthesizerConfig) *UpdateApplicationVersionRequest
	GetSynthesizerConfig() *UpdateApplicationVersionRequestSynthesizerConfig
	SetTranscriberConfig(v *UpdateApplicationVersionRequestTranscriberConfig) *UpdateApplicationVersionRequest
	GetTranscriberConfig() *UpdateApplicationVersionRequestTranscriberConfig
	SetVersionId(v string) *UpdateApplicationVersionRequest
	GetVersionId() *string
}

type UpdateApplicationVersionRequest struct {
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
	InteractionConfig *UpdateApplicationVersionRequestInteractionConfig `json:"InteractionConfig,omitempty" xml:"InteractionConfig,omitempty" type:"Struct"`
	RagConfig         *UpdateApplicationVersionRequestRagConfig         `json:"RagConfig,omitempty" xml:"RagConfig,omitempty" type:"Struct"`
	// This parameter is required.
	ScriptProfile *UpdateApplicationVersionRequestScriptProfile `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty" type:"Struct"`
	// if can be null:
	// true
	SynthesizerConfig *UpdateApplicationVersionRequestSynthesizerConfig `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty" type:"Struct"`
	// if can be null:
	// true
	TranscriberConfig *UpdateApplicationVersionRequestTranscriberConfig `json:"TranscriberConfig,omitempty" xml:"TranscriberConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 20904943-f711-494f-9f1f-e7f340f37707
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

func (s *UpdateApplicationVersionRequest) GetRagConfig() *UpdateApplicationVersionRequestRagConfig {
	return s.RagConfig
}

func (s *UpdateApplicationVersionRequest) GetScriptProfile() *UpdateApplicationVersionRequestScriptProfile {
	return s.ScriptProfile
}

func (s *UpdateApplicationVersionRequest) GetSynthesizerConfig() *UpdateApplicationVersionRequestSynthesizerConfig {
	return s.SynthesizerConfig
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
	if s.TranscriberConfig != nil {
		if err := s.TranscriberConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationVersionRequestInteractionConfig struct {
	SilenceDetectionConfig *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig `json:"SilenceDetectionConfig,omitempty" xml:"SilenceDetectionConfig,omitempty" type:"Struct"`
}

func (s UpdateApplicationVersionRequestInteractionConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestInteractionConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestInteractionConfig) GetSilenceDetectionConfig() *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig {
	return s.SilenceDetectionConfig
}

func (s *UpdateApplicationVersionRequestInteractionConfig) SetSilenceDetectionConfig(v *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) *UpdateApplicationVersionRequestInteractionConfig {
	s.SilenceDetectionConfig = v
	return s
}

func (s *UpdateApplicationVersionRequestInteractionConfig) Validate() error {
	if s.SilenceDetectionConfig != nil {
		if err := s.SilenceDetectionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig struct {
	// example:
	//
	// 3
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) SetTimeout(v int32) *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig {
	s.Timeout = &v
	return s
}

func (s *UpdateApplicationVersionRequestInteractionConfigSilenceDetectionConfig) Validate() error {
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
	AgentProfile *UpdateApplicationVersionRequestScriptProfileAgentProfile `json:"AgentProfile,omitempty" xml:"AgentProfile,omitempty" type:"Struct"`
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s UpdateApplicationVersionRequestScriptProfile) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionRequestScriptProfile) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionRequestScriptProfile) GetAgentProfile() *UpdateApplicationVersionRequestScriptProfileAgentProfile {
	return s.AgentProfile
}

func (s *UpdateApplicationVersionRequestScriptProfile) GetModel() *string {
	return s.Model
}

func (s *UpdateApplicationVersionRequestScriptProfile) SetAgentProfile(v *UpdateApplicationVersionRequestScriptProfileAgentProfile) *UpdateApplicationVersionRequestScriptProfile {
	s.AgentProfile = v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfile) SetModel(v string) *UpdateApplicationVersionRequestScriptProfile {
	s.Model = &v
	return s
}

func (s *UpdateApplicationVersionRequestScriptProfile) Validate() error {
	if s.AgentProfile != nil {
		if err := s.AgentProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationVersionRequestScriptProfileAgentProfile struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PromptsJson *string `json:"PromptsJson,omitempty" xml:"PromptsJson,omitempty"`
	// example:
	//
	// SFM_PROMPTS_DEFAULT
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

type UpdateApplicationVersionRequestSynthesizerConfig struct {
	Model            *string                                                           `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile *UpdateApplicationVersionRequestSynthesizerConfigNlsAccessProfile `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
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
	// 1
	PitchRate *int32                                                       `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	PronRules []*UpdateApplicationVersionRequestSynthesizerConfigPronRules `json:"PronRules,omitempty" xml:"PronRules,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// example:
	//
	// aixia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
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

type UpdateApplicationVersionRequestTranscriberConfig struct {
	CorrectionRules   []*UpdateApplicationVersionRequestTranscriberConfigCorrectionRules `json:"CorrectionRules,omitempty" xml:"CorrectionRules,omitempty" type:"Repeated"`
	CustomizationId   *string                                                            `json:"CustomizationId,omitempty" xml:"CustomizationId,omitempty"`
	EndSilenceTimeout *int32                                                             `json:"EndSilenceTimeout,omitempty" xml:"EndSilenceTimeout,omitempty"`
	Model             *string                                                            `json:"Model,omitempty" xml:"Model,omitempty"`
	NlsAccessProfile  *UpdateApplicationVersionRequestTranscriberConfigNlsAccessProfile  `json:"NlsAccessProfile,omitempty" xml:"NlsAccessProfile,omitempty" type:"Struct"`
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// example:
	//
	// ALIYUN
	NlsEngine            *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	SpeechNoiseThreshold *int32  `json:"SpeechNoiseThreshold,omitempty" xml:"SpeechNoiseThreshold,omitempty"`
	VocabularyId         *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
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
