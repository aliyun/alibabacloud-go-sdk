// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiVoiceAgentDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryAiVoiceAgentDetailResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryAiVoiceAgentDetailResponseBody
	GetCode() *string
	SetData(v *QueryAiVoiceAgentDetailResponseBodyData) *QueryAiVoiceAgentDetailResponseBody
	GetData() *QueryAiVoiceAgentDetailResponseBodyData
	SetMessage(v string) *QueryAiVoiceAgentDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAiVoiceAgentDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAiVoiceAgentDetailResponseBody
	GetSuccess() *bool
}

type QueryAiVoiceAgentDetailResponseBody struct {
	// The detailed reason why access was denied.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The agent details.
	Data *QueryAiVoiceAgentDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The status code description.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 742C9243-2870-B8D6-0C68-C60BEB2DF09A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAiVoiceAgentDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryAiVoiceAgentDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAiVoiceAgentDetailResponseBody) GetData() *QueryAiVoiceAgentDetailResponseBodyData {
	return s.Data
}

func (s *QueryAiVoiceAgentDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAiVoiceAgentDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAiVoiceAgentDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAiVoiceAgentDetailResponseBody) SetAccessDeniedDetail(v string) *QueryAiVoiceAgentDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBody) SetCode(v string) *QueryAiVoiceAgentDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBody) SetData(v *QueryAiVoiceAgentDetailResponseBodyData) *QueryAiVoiceAgentDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBody) SetMessage(v string) *QueryAiVoiceAgentDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBody) SetRequestId(v string) *QueryAiVoiceAgentDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBody) SetSuccess(v bool) *QueryAiVoiceAgentDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailResponseBodyData struct {
	// The agent ID.
	//
	// example:
	//
	// 12311212******
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The agent name.
	//
	// example:
	//
	// Test agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The intelligent outbound call voice configuration.
	AiVoiceAgentCallConfig *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig `json:"AiVoiceAgentCallConfig,omitempty" xml:"AiVoiceAgentCallConfig,omitempty" type:"Struct"`
	// The agent model configuration.
	AiVoiceAgentModelConfig *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig `json:"AiVoiceAgentModelConfig,omitempty" xml:"AiVoiceAgentModelConfig,omitempty" type:"Struct"`
	// The business scenario name.
	//
	// example:
	//
	// Personal lead conversion
	BusinessTypeName *string `json:"BusinessTypeName,omitempty" xml:"BusinessTypeName,omitempty"`
	// The agent description.
	//
	// example:
	//
	// Used for daily testing
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The knowledge base name.
	//
	// example:
	//
	// Test knowledge base
	KnowledgeName *string `json:"KnowledgeName,omitempty" xml:"KnowledgeName,omitempty"`
	// The agent status.
	//
	// example:
	//
	// 7
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The language style name.
	//
	// example:
	//
	// Friendly
	VoiceStyleName *string `json:"VoiceStyleName,omitempty" xml:"VoiceStyleName,omitempty"`
}

func (s QueryAiVoiceAgentDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) GetAgentId() *int64 {
	return s.AgentId
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) GetAgentName() *string {
	return s.AgentName
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) GetAiVoiceAgentCallConfig() *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig {
	return s.AiVoiceAgentCallConfig
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) GetAiVoiceAgentModelConfig() *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	return s.AiVoiceAgentModelConfig
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) GetBusinessTypeName() *string {
	return s.BusinessTypeName
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) GetKnowledgeName() *string {
	return s.KnowledgeName
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) GetStatus() *int64 {
	return s.Status
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) GetVoiceStyleName() *string {
	return s.VoiceStyleName
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) SetAgentId(v int64) *QueryAiVoiceAgentDetailResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) SetAgentName(v string) *QueryAiVoiceAgentDetailResponseBodyData {
	s.AgentName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) SetAiVoiceAgentCallConfig(v *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig) *QueryAiVoiceAgentDetailResponseBodyData {
	s.AiVoiceAgentCallConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) SetAiVoiceAgentModelConfig(v *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) *QueryAiVoiceAgentDetailResponseBodyData {
	s.AiVoiceAgentModelConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) SetBusinessTypeName(v string) *QueryAiVoiceAgentDetailResponseBodyData {
	s.BusinessTypeName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) SetDescription(v string) *QueryAiVoiceAgentDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) SetKnowledgeName(v string) *QueryAiVoiceAgentDetailResponseBodyData {
	s.KnowledgeName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) SetStatus(v int64) *QueryAiVoiceAgentDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) SetVoiceStyleName(v string) *QueryAiVoiceAgentDetailResponseBodyData {
	s.VoiceStyleName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyData) Validate() error {
	if s.AiVoiceAgentCallConfig != nil {
		if err := s.AiVoiceAgentCallConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AiVoiceAgentModelConfig != nil {
		if err := s.AiVoiceAgentModelConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig struct {
	// The call event configuration.
	EventConfig *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig `json:"EventConfig,omitempty" xml:"EventConfig,omitempty" type:"Struct"`
	// The TTS configuration.
	TtsConfig *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty" type:"Struct"`
	// The hot word ID.
	//
	// example:
	//
	// afb2c43**********83e6df30551c11f7
	VocabId *string `json:"VocabId,omitempty" xml:"VocabId,omitempty"`
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig) GetEventConfig() *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig {
	return s.EventConfig
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig) GetTtsConfig() *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig {
	return s.TtsConfig
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig) GetVocabId() *string {
	return s.VocabId
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig) SetEventConfig(v *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig {
	s.EventConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig) SetTtsConfig(v *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig {
	s.TtsConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig) SetVocabId(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig {
	s.VocabId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig) Validate() error {
	if s.EventConfig != nil {
		if err := s.EventConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TtsConfig != nil {
		if err := s.TtsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig struct {
	// Specifies whether to hang up when an intelligent answering service is detected.
	//
	// example:
	//
	// false
	CallAssistantHangup *bool `json:"CallAssistantHangup,omitempty" xml:"CallAssistantHangup,omitempty"`
	// Specifies whether intelligent answering service detection is enabled.
	//
	// example:
	//
	// true
	CallAssistantRecognize *bool `json:"CallAssistantRecognize,omitempty" xml:"CallAssistantRecognize,omitempty"`
	// Specifies whether the first silence triggers the model.
	//
	// example:
	//
	// false
	MuteActive *bool `json:"MuteActive,omitempty" xml:"MuteActive,omitempty"`
	// The silence duration.
	//
	// >
	//
	// >- Maximum value: 15s.
	//
	// >- Minimum value: 3s.
	//
	// example:
	//
	// 10
	MuteDuration *int64 `json:"MuteDuration,omitempty" xml:"MuteDuration,omitempty"`
	// The number of consecutive silence events before the system proactively hangs up.
	//
	// >
	//
	// >- Maximum value: 5.
	//
	// >- Minimum value: 1.
	//
	// example:
	//
	// 1
	MuteHangupNum *int64 `json:"MuteHangupNum,omitempty" xml:"MuteHangupNum,omitempty"`
	// The maximum call duration. The call is automatically hung up after the timeout. Unit: seconds.
	//
	// >
	//
	// >- Maximum value: 3600.
	//
	// >- Minimum value: 600.
	//
	// example:
	//
	// 600
	SessionTimeout *int64 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) GetCallAssistantHangup() *bool {
	return s.CallAssistantHangup
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) GetCallAssistantRecognize() *bool {
	return s.CallAssistantRecognize
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) GetMuteActive() *bool {
	return s.MuteActive
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) GetMuteDuration() *int64 {
	return s.MuteDuration
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) GetMuteHangupNum() *int64 {
	return s.MuteHangupNum
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) GetSessionTimeout() *int64 {
	return s.SessionTimeout
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) SetCallAssistantHangup(v bool) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig {
	s.CallAssistantHangup = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) SetCallAssistantRecognize(v bool) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig {
	s.CallAssistantRecognize = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) SetMuteActive(v bool) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig {
	s.MuteActive = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) SetMuteDuration(v int64) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig {
	s.MuteDuration = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) SetMuteHangupNum(v int64) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig {
	s.MuteHangupNum = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) SetSessionTimeout(v int64) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig {
	s.SessionTimeout = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig) Validate() error {
	return dara.Validate(s)
}

type QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig struct {
	// Indicates whether background sound is enabled.
	//
	// example:
	//
	// true
	BackgroundEnabled *bool `json:"BackgroundEnabled,omitempty" xml:"BackgroundEnabled,omitempty"`
	// The background sound ID.
	//
	// example:
	//
	// 1
	BackgroundSound *int64 `json:"BackgroundSound,omitempty" xml:"BackgroundSound,omitempty"`
	// The background sound volume. Valid values: 0: low. 1: medium. 2: high.
	//
	// example:
	//
	// 1
	BackgroundVolume *int64 `json:"BackgroundVolume,omitempty" xml:"BackgroundVolume,omitempty"`
	// Indicates whether audio mixing is enabled.
	//
	// example:
	//
	// true
	MixingEnabled *bool `json:"MixingEnabled,omitempty" xml:"MixingEnabled,omitempty"`
	// The mixing template ID.
	//
	// example:
	//
	// 1
	MixingTemplate *int64 `json:"MixingTemplate,omitempty" xml:"MixingTemplate,omitempty"`
	// The voice speed during TTS playback.
	//
	// >
	//
	// > - Valid values: -200 to 200. Default value: 0.
	//
	// > - If no value is specified, the voice speed configured in the large model application is used by default.
	//
	// example:
	//
	// 34
	TtsSpeed *int64 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	// The voice style.
	//
	// example:
	//
	// longxiaoxia_v2p1
	TtsStyle *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	// The TTS playback volume.
	//
	// >
	//
	// > - Valid values: 0 to 100. Default value: 0.
	//
	// > - If no value is specified, the volume configured in the large model application is used by default.
	//
	// example:
	//
	// 72
	TtsVolume *int64 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	// The voice code.
	//
	// example:
	//
	// Sample value
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// The voice type. Valid values: SYSTEM: system voice. COSYCLONE: cloned voice. BL-CUSTOM: custom premium cloned voice.
	//
	// example:
	//
	// Sample value
	VoiceType *string `json:"VoiceType,omitempty" xml:"VoiceType,omitempty"`
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) GetBackgroundEnabled() *bool {
	return s.BackgroundEnabled
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) GetBackgroundSound() *int64 {
	return s.BackgroundSound
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) GetBackgroundVolume() *int64 {
	return s.BackgroundVolume
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) GetMixingEnabled() *bool {
	return s.MixingEnabled
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) GetMixingTemplate() *int64 {
	return s.MixingTemplate
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) GetTtsSpeed() *int64 {
	return s.TtsSpeed
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) GetTtsStyle() *string {
	return s.TtsStyle
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) GetTtsVolume() *int64 {
	return s.TtsVolume
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) GetVoiceCode() *string {
	return s.VoiceCode
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) GetVoiceType() *string {
	return s.VoiceType
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) SetBackgroundEnabled(v bool) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig {
	s.BackgroundEnabled = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) SetBackgroundSound(v int64) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig {
	s.BackgroundSound = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) SetBackgroundVolume(v int64) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig {
	s.BackgroundVolume = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) SetMixingEnabled(v bool) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig {
	s.MixingEnabled = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) SetMixingTemplate(v int64) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig {
	s.MixingTemplate = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) SetTtsSpeed(v int64) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig {
	s.TtsSpeed = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) SetTtsStyle(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig {
	s.TtsStyle = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) SetTtsVolume(v int64) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig {
	s.TtsVolume = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) SetVoiceCode(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig {
	s.VoiceCode = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) SetVoiceType(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig {
	s.VoiceType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig) Validate() error {
	return dara.Validate(s)
}

type QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig struct {
	// The basic task configuration.
	//
	// example:
	//
	// Task description
	BasicTaskDescription *string `json:"BasicTaskDescription,omitempty" xml:"BasicTaskDescription,omitempty"`
	// The business scenario.
	//
	// example:
	//
	// 1
	BusinessType *int64 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The list of subtask configurations.
	ChildTaskList []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList `json:"ChildTaskList,omitempty" xml:"ChildTaskList,omitempty" type:"Repeated"`
	// Indicates whether custom exception handling is enabled.
	//
	// example:
	//
	// false
	CustomExceptionEnable *bool `json:"CustomExceptionEnable,omitempty" xml:"CustomExceptionEnable,omitempty"`
	// The custom exception file ID.
	//
	// example:
	//
	// OSS文件ID
	CustomExceptionFileId *string `json:"CustomExceptionFileId,omitempty" xml:"CustomExceptionFileId,omitempty"`
	// The name of the custom exception file.
	//
	// example:
	//
	// 异常测试文件.xlsx
	CustomExceptionFileName *string `json:"CustomExceptionFileName,omitempty" xml:"CustomExceptionFileName,omitempty"`
	// **[Deprecated]*	- This field is deprecated and will be removed in the future.
	CustomExceptionList []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList `json:"CustomExceptionList,omitempty" xml:"CustomExceptionList,omitempty" type:"Repeated"`
	// **[Deprecated]*	- This field is deprecated and will be removed in the future.
	//
	// example:
	//
	// -
	CustomExceptionUrlPath *string `json:"CustomExceptionUrlPath,omitempty" xml:"CustomExceptionUrlPath,omitempty"`
	// The language style.
	//
	// example:
	//
	// 2
	CustomExceptionVoiceStyle *int64 `json:"CustomExceptionVoiceStyle,omitempty" xml:"CustomExceptionVoiceStyle,omitempty"`
	// The description of the advanced task flow.
	//
	// example:
	//
	// This outbound call communicates with parents through three core steps, ........ 3. User needs > Proactive introduction
	FlowDesc *string `json:"FlowDesc,omitempty" xml:"FlowDesc,omitempty"`
	// The list of knowledge document IDs.
	KnowledgeDocIdList []*string `json:"KnowledgeDocIdList,omitempty" xml:"KnowledgeDocIdList,omitempty" type:"Repeated"`
	// The list of knowledge document names.
	KnowledgeDocNameList []*string `json:"KnowledgeDocNameList,omitempty" xml:"KnowledgeDocNameList,omitempty" type:"Repeated"`
	// The list of original file names of knowledge base documents.
	KnowledgeDocOriginalNameList []*string `json:"KnowledgeDocOriginalNameList,omitempty" xml:"KnowledgeDocOriginalNameList,omitempty" type:"Repeated"`
	// Indicates whether a knowledge base is associated.
	//
	// example:
	//
	// false
	KnowledgeEnable *bool `json:"KnowledgeEnable,omitempty" xml:"KnowledgeEnable,omitempty"`
	// The knowledge base ID.
	//
	// example:
	//
	// 1232131*******
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The main intent configuration.
	MainPurpose *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose `json:"MainPurpose,omitempty" xml:"MainPurpose,omitempty" type:"Struct"`
	// The output tag configurations.
	OutputTagConfig []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig `json:"OutputTagConfig,omitempty" xml:"OutputTagConfig,omitempty" type:"Repeated"`
	// The call variable configuration.
	PhoneTagConfig []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig `json:"PhoneTagConfig,omitempty" xml:"PhoneTagConfig,omitempty" type:"Repeated"`
	// The opening statement.
	//
	// example:
	//
	// Hello, this is ******
	Prologue *string `json:"Prologue,omitempty" xml:"Prologue,omitempty"`
	// The URL of the opening greeting recording audio file. This field has a value only when StartWordType is set to 1.
	//
	// example:
	//
	// Sample value
	RecordingFile *string `json:"RecordingFile,omitempty" xml:"RecordingFile,omitempty"`
	// The opening statement type. Valid values: 0: text. 1: recording.
	//
	// example:
	//
	// 1
	StartWordType *int64 `json:"StartWordType,omitempty" xml:"StartWordType,omitempty"`
	// The system role.
	//
	// example:
	//
	// Course sales
	SysRole *string `json:"SysRole,omitempty" xml:"SysRole,omitempty"`
	// The task type.
	//
	// example:
	//
	// ADVANCED
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The user role.
	//
	// example:
	//
	// New user in urgent need of courses
	UserRole *string `json:"UserRole,omitempty" xml:"UserRole,omitempty"`
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetBasicTaskDescription() *string {
	return s.BasicTaskDescription
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetBusinessType() *int64 {
	return s.BusinessType
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetChildTaskList() []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList {
	return s.ChildTaskList
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetCustomExceptionEnable() *bool {
	return s.CustomExceptionEnable
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetCustomExceptionFileId() *string {
	return s.CustomExceptionFileId
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetCustomExceptionFileName() *string {
	return s.CustomExceptionFileName
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetCustomExceptionList() []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList {
	return s.CustomExceptionList
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetCustomExceptionUrlPath() *string {
	return s.CustomExceptionUrlPath
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetCustomExceptionVoiceStyle() *int64 {
	return s.CustomExceptionVoiceStyle
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetFlowDesc() *string {
	return s.FlowDesc
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetKnowledgeDocIdList() []*string {
	return s.KnowledgeDocIdList
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetKnowledgeDocNameList() []*string {
	return s.KnowledgeDocNameList
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetKnowledgeDocOriginalNameList() []*string {
	return s.KnowledgeDocOriginalNameList
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetKnowledgeEnable() *bool {
	return s.KnowledgeEnable
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetKnowledgeId() *string {
	return s.KnowledgeId
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetMainPurpose() *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose {
	return s.MainPurpose
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetOutputTagConfig() []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig {
	return s.OutputTagConfig
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetPhoneTagConfig() []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig {
	return s.PhoneTagConfig
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetPrologue() *string {
	return s.Prologue
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetRecordingFile() *string {
	return s.RecordingFile
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetStartWordType() *int64 {
	return s.StartWordType
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetSysRole() *string {
	return s.SysRole
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) GetUserRole() *string {
	return s.UserRole
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetBasicTaskDescription(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.BasicTaskDescription = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetBusinessType(v int64) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.BusinessType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetChildTaskList(v []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.ChildTaskList = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetCustomExceptionEnable(v bool) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.CustomExceptionEnable = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetCustomExceptionFileId(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.CustomExceptionFileId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetCustomExceptionFileName(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.CustomExceptionFileName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetCustomExceptionList(v []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.CustomExceptionList = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetCustomExceptionUrlPath(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.CustomExceptionUrlPath = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetCustomExceptionVoiceStyle(v int64) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.CustomExceptionVoiceStyle = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetFlowDesc(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.FlowDesc = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetKnowledgeDocIdList(v []*string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.KnowledgeDocIdList = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetKnowledgeDocNameList(v []*string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.KnowledgeDocNameList = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetKnowledgeDocOriginalNameList(v []*string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.KnowledgeDocOriginalNameList = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetKnowledgeEnable(v bool) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.KnowledgeEnable = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetKnowledgeId(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.KnowledgeId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetMainPurpose(v *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.MainPurpose = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetOutputTagConfig(v []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.OutputTagConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetPhoneTagConfig(v []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.PhoneTagConfig = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetPrologue(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.Prologue = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetRecordingFile(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.RecordingFile = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetStartWordType(v int64) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.StartWordType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetSysRole(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.SysRole = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetTaskType(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.TaskType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) SetUserRole(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig {
	s.UserRole = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig) Validate() error {
	if s.ChildTaskList != nil {
		for _, item := range s.ChildTaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CustomExceptionList != nil {
		for _, item := range s.CustomExceptionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MainPurpose != nil {
		if err := s.MainPurpose.Validate(); err != nil {
			return err
		}
	}
	if s.OutputTagConfig != nil {
		for _, item := range s.OutputTagConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PhoneTagConfig != nil {
		for _, item := range s.PhoneTagConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList struct {
	// The subtask description.
	//
	// example:
	//
	// 新用户邀约: \\"喂，家长您好！我是*****的王老师，我们现在有**课程100个试听名额，想帮宝贝安排一下体验，您看什么时候比较方便呢？\\
	ChildTaskDescription *string `json:"ChildTaskDescription,omitempty" xml:"ChildTaskDescription,omitempty"`
	// The subtask name.
	//
	// example:
	//
	// Opening and invitation
	ChildTaskName *string `json:"ChildTaskName,omitempty" xml:"ChildTaskName,omitempty"`
	// The unique ID of the subtask.
	//
	// example:
	//
	// 280cd4bf-*******df472c
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList) GetChildTaskDescription() *string {
	return s.ChildTaskDescription
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList) GetChildTaskName() *string {
	return s.ChildTaskName
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList) GetId() *string {
	return s.Id
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList) SetChildTaskDescription(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList {
	s.ChildTaskDescription = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList) SetChildTaskName(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList {
	s.ChildTaskName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList) SetId(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList {
	s.Id = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList) Validate() error {
	return dara.Validate(s)
}

type QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList struct {
	// Specifies whether to output as an exception tag.
	//
	// example:
	//
	// true
	ExceptionSign *bool `json:"ExceptionSign,omitempty" xml:"ExceptionSign,omitempty"`
	// The exception type.
	//
	// example:
	//
	// -
	ExceptionType *string `json:"ExceptionType,omitempty" xml:"ExceptionType,omitempty"`
	// The reply content.
	//
	// example:
	//
	// -
	Reply *string `json:"Reply,omitempty" xml:"Reply,omitempty"`
	// Specifies whether interruption is supported.
	//
	// example:
	//
	// false
	SupportBreak *bool `json:"SupportBreak,omitempty" xml:"SupportBreak,omitempty"`
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList) GetExceptionSign() *bool {
	return s.ExceptionSign
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList) GetExceptionType() *string {
	return s.ExceptionType
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList) GetReply() *string {
	return s.Reply
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList) GetSupportBreak() *bool {
	return s.SupportBreak
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList) SetExceptionSign(v bool) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList {
	s.ExceptionSign = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList) SetExceptionType(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList {
	s.ExceptionType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList) SetReply(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList {
	s.Reply = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList) SetSupportBreak(v bool) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList {
	s.SupportBreak = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList) Validate() error {
	return dara.Validate(s)
}

type QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose struct {
	// The main intent ID.
	//
	// example:
	//
	// 1ee6e994-08e0-xxxx-f662-1659cc54d409
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the main intent.
	//
	// example:
	//
	// Call effectiveness rating: A, B, C, D (A is the best)
	MainPurposeDescription *string `json:"MainPurposeDescription,omitempty" xml:"MainPurposeDescription,omitempty"`
	// The list of valid values for the main intent.
	MainPurposeEnum []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum `json:"MainPurposeEnum,omitempty" xml:"MainPurposeEnum,omitempty" type:"Repeated"`
	// The main intent name.
	//
	// example:
	//
	// Call effectiveness rating
	MainPurposeName *string `json:"MainPurposeName,omitempty" xml:"MainPurposeName,omitempty"`
	// The value type of the main intent. Currently, only the ENUM type is supported.
	//
	// example:
	//
	// ENUM
	MainPurposeType *string `json:"MainPurposeType,omitempty" xml:"MainPurposeType,omitempty"`
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose) GetId() *string {
	return s.Id
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose) GetMainPurposeDescription() *string {
	return s.MainPurposeDescription
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose) GetMainPurposeEnum() []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum {
	return s.MainPurposeEnum
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose) GetMainPurposeName() *string {
	return s.MainPurposeName
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose) GetMainPurposeType() *string {
	return s.MainPurposeType
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose) SetId(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose {
	s.Id = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose) SetMainPurposeDescription(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose {
	s.MainPurposeDescription = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose) SetMainPurposeEnum(v []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose {
	s.MainPurposeEnum = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose) SetMainPurposeName(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose {
	s.MainPurposeName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose) SetMainPurposeType(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose {
	s.MainPurposeType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose) Validate() error {
	if s.MainPurposeEnum != nil {
		for _, item := range s.MainPurposeEnum {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum struct {
	// The description of the valid value.
	//
	// example:
	//
	// Strong intent
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique ID of the valid value.
	//
	// example:
	//
	// d5606d80-7625-dcea-xxxx-17f66fbb564a
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The valid value.
	//
	// example:
	//
	// A
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum) GetDescription() *string {
	return s.Description
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum) GetId() *string {
	return s.Id
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum) GetValue() *string {
	return s.Value
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum) SetDescription(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum {
	s.Description = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum) SetId(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum {
	s.Id = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum) SetValue(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum {
	s.Value = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum) Validate() error {
	return dara.Validate(s)
}

type QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig struct {
	// The unique ID of the tag.
	//
	// example:
	//
	// 8757************2c499fa
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tag description.
	//
	// example:
	//
	// Records the final status of this call
	OutputTagDescription *string `json:"OutputTagDescription,omitempty" xml:"OutputTagDescription,omitempty"`
	// The tag enum values. This field is available only when the tag value type is ENUM.
	OutputTagEnum []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum `json:"OutputTagEnum,omitempty" xml:"OutputTagEnum,omitempty" type:"Repeated"`
	// The tag name.
	//
	// example:
	//
	// Customer intent level
	OutputTagName *string `json:"OutputTagName,omitempty" xml:"OutputTagName,omitempty"`
	// The tag value type.
	//
	// example:
	//
	// ENUM
	OutputTagType *string `json:"OutputTagType,omitempty" xml:"OutputTagType,omitempty"`
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig) GetId() *string {
	return s.Id
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig) GetOutputTagDescription() *string {
	return s.OutputTagDescription
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig) GetOutputTagEnum() []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum {
	return s.OutputTagEnum
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig) GetOutputTagName() *string {
	return s.OutputTagName
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig) GetOutputTagType() *string {
	return s.OutputTagType
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig) SetId(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig {
	s.Id = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig) SetOutputTagDescription(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig {
	s.OutputTagDescription = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig) SetOutputTagEnum(v []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig {
	s.OutputTagEnum = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig) SetOutputTagName(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig {
	s.OutputTagName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig) SetOutputTagType(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig {
	s.OutputTagType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig) Validate() error {
	if s.OutputTagEnum != nil {
		for _, item := range s.OutputTagEnum {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum struct {
	// The description of the tag enum value.
	//
	// example:
	//
	// High (very positive, high probability of conversion)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique ID of the tag enum value.
	//
	// example:
	//
	// 8757************2c499fa
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tag enum value.
	//
	// example:
	//
	// High
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum) GetDescription() *string {
	return s.Description
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum) GetId() *string {
	return s.Id
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum) GetValue() *string {
	return s.Value
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum) SetDescription(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum {
	s.Description = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum) SetId(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum {
	s.Id = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum) SetValue(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum {
	s.Value = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum) Validate() error {
	return dara.Validate(s)
}

type QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig struct {
	// The unique ID of the variable.
	//
	// example:
	//
	// 280cd4bf-*******df472c
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The variable description.
	//
	// example:
	//
	// The car series the user is interested in
	PhoneTagDescription *string `json:"PhoneTagDescription,omitempty" xml:"PhoneTagDescription,omitempty"`
	// The list of enumeration values for the variable. This field is present only when the variable value type is ENUM.
	PhoneTagEnum []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum `json:"PhoneTagEnum,omitempty" xml:"PhoneTagEnum,omitempty" type:"Repeated"`
	// The variable key name.
	//
	// example:
	//
	// testParam
	PhoneTagKey *string `json:"PhoneTagKey,omitempty" xml:"PhoneTagKey,omitempty"`
	// The Chinese name of the variable.
	//
	// example:
	//
	// 意向车系
	PhoneTagName *string `json:"PhoneTagName,omitempty" xml:"PhoneTagName,omitempty"`
	// Indicates whether the variable is required.
	//
	// example:
	//
	// true
	PhoneTagRequired *bool `json:"PhoneTagRequired,omitempty" xml:"PhoneTagRequired,omitempty"`
	// The source of the call variable.
	//
	// example:
	//
	// Sample value
	PhoneTagSource *string `json:"PhoneTagSource,omitempty" xml:"PhoneTagSource,omitempty"`
	// The variable value type.
	//
	// example:
	//
	// ENUM
	PhoneTagType *string `json:"PhoneTagType,omitempty" xml:"PhoneTagType,omitempty"`
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) GetId() *string {
	return s.Id
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) GetPhoneTagDescription() *string {
	return s.PhoneTagDescription
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) GetPhoneTagEnum() []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum {
	return s.PhoneTagEnum
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) GetPhoneTagKey() *string {
	return s.PhoneTagKey
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) GetPhoneTagName() *string {
	return s.PhoneTagName
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) GetPhoneTagRequired() *bool {
	return s.PhoneTagRequired
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) GetPhoneTagSource() *string {
	return s.PhoneTagSource
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) GetPhoneTagType() *string {
	return s.PhoneTagType
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) SetId(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig {
	s.Id = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) SetPhoneTagDescription(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig {
	s.PhoneTagDescription = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) SetPhoneTagEnum(v []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig {
	s.PhoneTagEnum = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) SetPhoneTagKey(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig {
	s.PhoneTagKey = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) SetPhoneTagName(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig {
	s.PhoneTagName = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) SetPhoneTagRequired(v bool) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig {
	s.PhoneTagRequired = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) SetPhoneTagSource(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig {
	s.PhoneTagSource = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) SetPhoneTagType(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig {
	s.PhoneTagType = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig) Validate() error {
	if s.PhoneTagEnum != nil {
		for _, item := range s.PhoneTagEnum {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum struct {
	// The description of the enumeration value.
	//
	// example:
	//
	// The customer completely rejects retention
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique ID of the enum value.
	//
	// example:
	//
	// c3d4ff4e-*********bc26dc044682
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The enumeration value.
	//
	// example:
	//
	// Completely unacceptable
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum) GetDescription() *string {
	return s.Description
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum) GetId() *string {
	return s.Id
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum) GetValue() *string {
	return s.Value
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum) SetDescription(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum {
	s.Description = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum) SetId(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum {
	s.Id = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum) SetValue(v string) *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum {
	s.Value = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum) Validate() error {
	return dara.Validate(s)
}
