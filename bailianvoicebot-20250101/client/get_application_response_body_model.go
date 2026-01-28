// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetApplicationResponseBody
	GetCode() *string
	SetData(v *GetApplicationResponseBodyData) *GetApplicationResponseBody
	GetData() *GetApplicationResponseBodyData
	SetHttpStatusCode(v int32) *GetApplicationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetApplicationResponseBody
	GetRequestId() *string
}

type GetApplicationResponseBody struct {
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D771A1B6-3D5F-174A-BEE1-98CE1000D337
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetApplicationResponseBody) GetData() *GetApplicationResponseBodyData {
	return s.Data
}

func (s *GetApplicationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationResponseBody) SetCode(v string) *GetApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *GetApplicationResponseBody) SetData(v *GetApplicationResponseBodyData) *GetApplicationResponseBody {
	s.Data = v
	return s
}

func (s *GetApplicationResponseBody) SetHttpStatusCode(v int32) *GetApplicationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetApplicationResponseBody) SetMessage(v string) *GetApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationResponseBodyData struct {
	// example:
	//
	// a395011f-a247-400f-bc69-28796749fd52
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 10
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// example:
	//
	// 1730081561000
	CreatedTime  *int64                                      `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	DraftVersion *GetApplicationResponseBodyDataDraftVersion `json:"DraftVersion,omitempty" xml:"DraftVersion,omitempty" type:"Struct"`
	Name         *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MANAGED
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// example:
	//
	// PROMPTS
	NluEngine        *string                                         `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	PublishedVersion *GetApplicationResponseBodyDataPublishedVersion `json:"PublishedVersion,omitempty" xml:"PublishedVersion,omitempty" type:"Struct"`
	// example:
	//
	// 1730081561000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyData) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationResponseBodyData) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *GetApplicationResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetApplicationResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyData) GetDraftVersion() *GetApplicationResponseBodyDataDraftVersion {
	return s.DraftVersion
}

func (s *GetApplicationResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetApplicationResponseBodyData) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *GetApplicationResponseBodyData) GetNluEngine() *string {
	return s.NluEngine
}

func (s *GetApplicationResponseBodyData) GetPublishedVersion() *GetApplicationResponseBodyDataPublishedVersion {
	return s.PublishedVersion
}

func (s *GetApplicationResponseBodyData) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetApplicationResponseBodyData) SetApplicationId(v string) *GetApplicationResponseBodyData {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetConcurrency(v int32) *GetApplicationResponseBodyData {
	s.Concurrency = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetCreatedTime(v int64) *GetApplicationResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetDescription(v string) *GetApplicationResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetDraftVersion(v *GetApplicationResponseBodyDataDraftVersion) *GetApplicationResponseBodyData {
	s.DraftVersion = v
	return s
}

func (s *GetApplicationResponseBodyData) SetName(v string) *GetApplicationResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetNluAccessType(v string) *GetApplicationResponseBodyData {
	s.NluAccessType = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetNluEngine(v string) *GetApplicationResponseBodyData {
	s.NluEngine = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetPublishedVersion(v *GetApplicationResponseBodyDataPublishedVersion) *GetApplicationResponseBodyData {
	s.PublishedVersion = v
	return s
}

func (s *GetApplicationResponseBodyData) SetUpdatedTime(v int64) *GetApplicationResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetApplicationResponseBodyData) Validate() error {
	if s.DraftVersion != nil {
		if err := s.DraftVersion.Validate(); err != nil {
			return err
		}
	}
	if s.PublishedVersion != nil {
		if err := s.PublishedVersion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationResponseBodyDataDraftVersion struct {
	InteractionConfig *GetApplicationResponseBodyDataDraftVersionInteractionConfig `json:"InteractionConfig,omitempty" xml:"InteractionConfig,omitempty" type:"Struct"`
	ScriptProfile     *GetApplicationResponseBodyDataDraftVersionScriptProfile     `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty" type:"Struct"`
	SynthesizerConfig *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty" type:"Struct"`
	TranscriberConfig *GetApplicationResponseBodyDataDraftVersionTranscriberConfig `json:"TranscriberConfig,omitempty" xml:"TranscriberConfig,omitempty" type:"Struct"`
	// example:
	//
	// 743219815472857088
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersion) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersion) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersion) GetInteractionConfig() *GetApplicationResponseBodyDataDraftVersionInteractionConfig {
	return s.InteractionConfig
}

func (s *GetApplicationResponseBodyDataDraftVersion) GetScriptProfile() *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	return s.ScriptProfile
}

func (s *GetApplicationResponseBodyDataDraftVersion) GetSynthesizerConfig() *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	return s.SynthesizerConfig
}

func (s *GetApplicationResponseBodyDataDraftVersion) GetTranscriberConfig() *GetApplicationResponseBodyDataDraftVersionTranscriberConfig {
	return s.TranscriberConfig
}

func (s *GetApplicationResponseBodyDataDraftVersion) GetVersionId() *string {
	return s.VersionId
}

func (s *GetApplicationResponseBodyDataDraftVersion) SetInteractionConfig(v *GetApplicationResponseBodyDataDraftVersionInteractionConfig) *GetApplicationResponseBodyDataDraftVersion {
	s.InteractionConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersion) SetScriptProfile(v *GetApplicationResponseBodyDataDraftVersionScriptProfile) *GetApplicationResponseBodyDataDraftVersion {
	s.ScriptProfile = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersion) SetSynthesizerConfig(v *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) *GetApplicationResponseBodyDataDraftVersion {
	s.SynthesizerConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersion) SetTranscriberConfig(v *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) *GetApplicationResponseBodyDataDraftVersion {
	s.TranscriberConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersion) SetVersionId(v string) *GetApplicationResponseBodyDataDraftVersion {
	s.VersionId = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersion) Validate() error {
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

type GetApplicationResponseBodyDataDraftVersionInteractionConfig struct {
	SilenceDetectionConfig *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig `json:"SilenceDetectionConfig,omitempty" xml:"SilenceDetectionConfig,omitempty" type:"Struct"`
}

func (s GetApplicationResponseBodyDataDraftVersionInteractionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionInteractionConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfig) GetSilenceDetectionConfig() *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig {
	return s.SilenceDetectionConfig
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfig) SetSilenceDetectionConfig(v *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) *GetApplicationResponseBodyDataDraftVersionInteractionConfig {
	s.SilenceDetectionConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfig) Validate() error {
	if s.SilenceDetectionConfig != nil {
		if err := s.SilenceDetectionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig struct {
	// example:
	//
	// 3
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) SetTimeout(v int32) *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig {
	s.Timeout = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionInteractionConfigSilenceDetectionConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataDraftVersionScriptProfile struct {
	AgentProfile *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile `json:"AgentProfile,omitempty" xml:"AgentProfile,omitempty" type:"Struct"`
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 0.8
	Temperature *string `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// example:
	//
	// 0.1
	TopP *string `json:"TopP,omitempty" xml:"TopP,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionScriptProfile) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionScriptProfile) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) GetAgentProfile() *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile {
	return s.AgentProfile
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) GetModel() *string {
	return s.Model
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) GetTemperature() *string {
	return s.Temperature
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) GetTopP() *string {
	return s.TopP
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) SetAgentProfile(v *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	s.AgentProfile = v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) SetModel(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	s.Model = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) SetTemperature(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	s.Temperature = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) SetTopP(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfile {
	s.TopP = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfile) Validate() error {
	if s.AgentProfile != nil {
		if err := s.AgentProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile struct {
	// example:
	//
	// 6a50b67072d44788951de29758432d94
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PromptsJson    *string `json:"PromptsJson,omitempty" xml:"PromptsJson,omitempty"`
	// example:
	//
	// SFM_PROMPTS_DEFAULT
	ScriptProfileTemplateId *string `json:"ScriptProfileTemplateId,omitempty" xml:"ScriptProfileTemplateId,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) GetAgentProfileId() *string {
	return s.AgentProfileId
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) GetPromptsJson() *string {
	return s.PromptsJson
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) GetScriptProfileTemplateId() *string {
	return s.ScriptProfileTemplateId
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) SetAgentProfileId(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile {
	s.AgentProfileId = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) SetDescription(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) SetPromptsJson(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile {
	s.PromptsJson = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) SetScriptProfileTemplateId(v string) *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile {
	s.ScriptProfileTemplateId = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionScriptProfileAgentProfile) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataDraftVersionSynthesizerConfig struct {
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
	// 5
	PitchRate *int32 `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// example:
	//
	// 1
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

func (s GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GetPitchRate() *int32 {
	return s.PitchRate
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GetVoice() *string {
	return s.Voice
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) SetNlsAccessType(v string) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	s.NlsAccessType = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) SetNlsEngine(v string) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	s.NlsEngine = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) SetPitchRate(v int32) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	s.PitchRate = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) SetSpeechRate(v int32) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	s.SpeechRate = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) SetVoice(v string) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	s.Voice = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) SetVolume(v int32) *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig {
	s.Volume = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionSynthesizerConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataDraftVersionTranscriberConfig struct {
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// example:
	//
	// ALIYUN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
}

func (s GetApplicationResponseBodyDataDraftVersionTranscriberConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataDraftVersionTranscriberConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) SetNlsAccessType(v string) *GetApplicationResponseBodyDataDraftVersionTranscriberConfig {
	s.NlsAccessType = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) SetNlsEngine(v string) *GetApplicationResponseBodyDataDraftVersionTranscriberConfig {
	s.NlsEngine = &v
	return s
}

func (s *GetApplicationResponseBodyDataDraftVersionTranscriberConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersion struct {
	InteractionConfig *GetApplicationResponseBodyDataPublishedVersionInteractionConfig `json:"InteractionConfig,omitempty" xml:"InteractionConfig,omitempty" type:"Struct"`
	ScriptProfile     *GetApplicationResponseBodyDataPublishedVersionScriptProfile     `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty" type:"Struct"`
	SynthesizerConfig *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty" type:"Struct"`
	TranscriberConfig *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig `json:"TranscriberConfig,omitempty" xml:"TranscriberConfig,omitempty" type:"Struct"`
	// example:
	//
	// 47889c1f-dd3f-4ace-9587-a13a3563e678
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersion) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersion) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersion) GetInteractionConfig() *GetApplicationResponseBodyDataPublishedVersionInteractionConfig {
	return s.InteractionConfig
}

func (s *GetApplicationResponseBodyDataPublishedVersion) GetScriptProfile() *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	return s.ScriptProfile
}

func (s *GetApplicationResponseBodyDataPublishedVersion) GetSynthesizerConfig() *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	return s.SynthesizerConfig
}

func (s *GetApplicationResponseBodyDataPublishedVersion) GetTranscriberConfig() *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig {
	return s.TranscriberConfig
}

func (s *GetApplicationResponseBodyDataPublishedVersion) GetVersionId() *string {
	return s.VersionId
}

func (s *GetApplicationResponseBodyDataPublishedVersion) SetInteractionConfig(v *GetApplicationResponseBodyDataPublishedVersionInteractionConfig) *GetApplicationResponseBodyDataPublishedVersion {
	s.InteractionConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersion) SetScriptProfile(v *GetApplicationResponseBodyDataPublishedVersionScriptProfile) *GetApplicationResponseBodyDataPublishedVersion {
	s.ScriptProfile = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersion) SetSynthesizerConfig(v *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) *GetApplicationResponseBodyDataPublishedVersion {
	s.SynthesizerConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersion) SetTranscriberConfig(v *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) *GetApplicationResponseBodyDataPublishedVersion {
	s.TranscriberConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersion) SetVersionId(v string) *GetApplicationResponseBodyDataPublishedVersion {
	s.VersionId = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersion) Validate() error {
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

type GetApplicationResponseBodyDataPublishedVersionInteractionConfig struct {
	SilenceDetectionConfig *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig `json:"SilenceDetectionConfig,omitempty" xml:"SilenceDetectionConfig,omitempty" type:"Struct"`
}

func (s GetApplicationResponseBodyDataPublishedVersionInteractionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionInteractionConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfig) GetSilenceDetectionConfig() *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig {
	return s.SilenceDetectionConfig
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfig) SetSilenceDetectionConfig(v *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) *GetApplicationResponseBodyDataPublishedVersionInteractionConfig {
	s.SilenceDetectionConfig = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfig) Validate() error {
	if s.SilenceDetectionConfig != nil {
		if err := s.SilenceDetectionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig struct {
	// example:
	//
	// 30
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) SetTimeout(v int32) *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig {
	s.Timeout = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionInteractionConfigSilenceDetectionConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersionScriptProfile struct {
	AgentProfile *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile `json:"AgentProfile,omitempty" xml:"AgentProfile,omitempty" type:"Struct"`
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 0.8
	Temperature *string `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// example:
	//
	// 0.1
	TopP *string `json:"TopP,omitempty" xml:"TopP,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionScriptProfile) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionScriptProfile) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) GetAgentProfile() *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	return s.AgentProfile
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) GetModel() *string {
	return s.Model
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) GetTemperature() *string {
	return s.Temperature
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) GetTopP() *string {
	return s.TopP
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) SetAgentProfile(v *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	s.AgentProfile = v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) SetModel(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	s.Model = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) SetTemperature(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	s.Temperature = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) SetTopP(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfile {
	s.TopP = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfile) Validate() error {
	if s.AgentProfile != nil {
		if err := s.AgentProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile struct {
	// example:
	//
	// b97b6822dd624c32b6c2a54d717db718
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PromptsJson    *string `json:"PromptsJson,omitempty" xml:"PromptsJson,omitempty"`
	// example:
	//
	// SFM_PROMPTS_DEFAULT
	ScriptProfileTemplateId *string `json:"ScriptProfileTemplateId,omitempty" xml:"ScriptProfileTemplateId,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) GetAgentProfileId() *string {
	return s.AgentProfileId
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) GetPromptsJson() *string {
	return s.PromptsJson
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) GetScriptProfileTemplateId() *string {
	return s.ScriptProfileTemplateId
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) SetAgentProfileId(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	s.AgentProfileId = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) SetDescription(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) SetPromptsJson(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	s.PromptsJson = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) SetScriptProfileTemplateId(v string) *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile {
	s.ScriptProfileTemplateId = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionScriptProfileAgentProfile) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig struct {
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
	// 3
	PitchRate *int32 `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// example:
	//
	// -20
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

func (s GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GetPitchRate() *int32 {
	return s.PitchRate
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GetVoice() *string {
	return s.Voice
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) SetNlsAccessType(v string) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	s.NlsAccessType = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) SetNlsEngine(v string) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	s.NlsEngine = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) SetPitchRate(v int32) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	s.PitchRate = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) SetSpeechRate(v int32) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	s.SpeechRate = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) SetVoice(v string) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	s.Voice = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) SetVolume(v int32) *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig {
	s.Volume = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionSynthesizerConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPublishedVersionTranscriberConfig struct {
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// example:
	//
	// ALIYUN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
}

func (s GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) SetNlsAccessType(v string) *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig {
	s.NlsAccessType = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) SetNlsEngine(v string) *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig {
	s.NlsEngine = &v
	return s
}

func (s *GetApplicationResponseBodyDataPublishedVersionTranscriberConfig) Validate() error {
	return dara.Validate(s)
}
