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
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryAiVoiceAgentDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 742C9243-2870-B8D6-0C68-C60BEB2DF09A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 12311212******
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	AgentName               *string                                                         `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	AiVoiceAgentCallConfig  *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfig  `json:"AiVoiceAgentCallConfig,omitempty" xml:"AiVoiceAgentCallConfig,omitempty" type:"Struct"`
	AiVoiceAgentModelConfig *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfig `json:"AiVoiceAgentModelConfig,omitempty" xml:"AiVoiceAgentModelConfig,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值示例值
	BusinessTypeName *string `json:"BusinessTypeName,omitempty" xml:"BusinessTypeName,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	KnowledgeName *string `json:"KnowledgeName,omitempty" xml:"KnowledgeName,omitempty"`
	// example:
	//
	// 7
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 示例值示例值
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
	EventConfig *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigEventConfig `json:"EventConfig,omitempty" xml:"EventConfig,omitempty" type:"Struct"`
	TtsConfig   *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentCallConfigTtsConfig   `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty" type:"Struct"`
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
	// example:
	//
	// false
	CallAssistantHangup *bool `json:"CallAssistantHangup,omitempty" xml:"CallAssistantHangup,omitempty"`
	// example:
	//
	// true
	CallAssistantRecognize *bool `json:"CallAssistantRecognize,omitempty" xml:"CallAssistantRecognize,omitempty"`
	// example:
	//
	// false
	MuteActive *bool `json:"MuteActive,omitempty" xml:"MuteActive,omitempty"`
	// example:
	//
	// 10
	MuteDuration *int64 `json:"MuteDuration,omitempty" xml:"MuteDuration,omitempty"`
	// example:
	//
	// 1
	MuteHangupNum *int64 `json:"MuteHangupNum,omitempty" xml:"MuteHangupNum,omitempty"`
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
	// example:
	//
	// true
	BackgroundEnabled *bool `json:"BackgroundEnabled,omitempty" xml:"BackgroundEnabled,omitempty"`
	// example:
	//
	// 1
	BackgroundSound *int64 `json:"BackgroundSound,omitempty" xml:"BackgroundSound,omitempty"`
	// example:
	//
	// 1
	BackgroundVolume *int64 `json:"BackgroundVolume,omitempty" xml:"BackgroundVolume,omitempty"`
	// example:
	//
	// true
	MixingEnabled *bool `json:"MixingEnabled,omitempty" xml:"MixingEnabled,omitempty"`
	// example:
	//
	// 1
	MixingTemplate *int64 `json:"MixingTemplate,omitempty" xml:"MixingTemplate,omitempty"`
	// example:
	//
	// 34
	TtsSpeed *int64 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	// example:
	//
	// longxiaoxia_v2p1
	TtsStyle *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	// example:
	//
	// 72
	TtsVolume *int64 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	// example:
	//
	// 示例值
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// example:
	//
	// 示例值
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
	// example:
	//
	// 示例值示例值示例值
	BasicTaskDescription *string `json:"BasicTaskDescription,omitempty" xml:"BasicTaskDescription,omitempty"`
	// example:
	//
	// 1
	BusinessType  *int64                                                                         `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	ChildTaskList []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigChildTaskList `json:"ChildTaskList,omitempty" xml:"ChildTaskList,omitempty" type:"Repeated"`
	// example:
	//
	// false
	CustomExceptionEnable *bool `json:"CustomExceptionEnable,omitempty" xml:"CustomExceptionEnable,omitempty"`
	// example:
	//
	// 示例值示例值
	CustomExceptionFileId *string `json:"CustomExceptionFileId,omitempty" xml:"CustomExceptionFileId,omitempty"`
	// example:
	//
	// 示例值
	CustomExceptionFileName *string                                                                              `json:"CustomExceptionFileName,omitempty" xml:"CustomExceptionFileName,omitempty"`
	CustomExceptionList     []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigCustomExceptionList `json:"CustomExceptionList,omitempty" xml:"CustomExceptionList,omitempty" type:"Repeated"`
	// example:
	//
	// -
	CustomExceptionUrlPath *string `json:"CustomExceptionUrlPath,omitempty" xml:"CustomExceptionUrlPath,omitempty"`
	// example:
	//
	// 2
	CustomExceptionVoiceStyle *int64 `json:"CustomExceptionVoiceStyle,omitempty" xml:"CustomExceptionVoiceStyle,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	FlowDesc                     *string   `json:"FlowDesc,omitempty" xml:"FlowDesc,omitempty"`
	KnowledgeDocIdList           []*string `json:"KnowledgeDocIdList,omitempty" xml:"KnowledgeDocIdList,omitempty" type:"Repeated"`
	KnowledgeDocNameList         []*string `json:"KnowledgeDocNameList,omitempty" xml:"KnowledgeDocNameList,omitempty" type:"Repeated"`
	KnowledgeDocOriginalNameList []*string `json:"KnowledgeDocOriginalNameList,omitempty" xml:"KnowledgeDocOriginalNameList,omitempty" type:"Repeated"`
	// example:
	//
	// false
	KnowledgeEnable *bool `json:"KnowledgeEnable,omitempty" xml:"KnowledgeEnable,omitempty"`
	// example:
	//
	// 1232131*******
	KnowledgeId     *string                                                                          `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	MainPurpose     *QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurpose       `json:"MainPurpose,omitempty" xml:"MainPurpose,omitempty" type:"Struct"`
	OutputTagConfig []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfig `json:"OutputTagConfig,omitempty" xml:"OutputTagConfig,omitempty" type:"Repeated"`
	PhoneTagConfig  []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfig  `json:"PhoneTagConfig,omitempty" xml:"PhoneTagConfig,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值示例值
	Prologue *string `json:"Prologue,omitempty" xml:"Prologue,omitempty"`
	// example:
	//
	// 示例值示例值
	RecordingFile *string `json:"RecordingFile,omitempty" xml:"RecordingFile,omitempty"`
	// example:
	//
	// 1
	StartWordType *int64 `json:"StartWordType,omitempty" xml:"StartWordType,omitempty"`
	// example:
	//
	// 示例值示例值
	SysRole *string `json:"SysRole,omitempty" xml:"SysRole,omitempty"`
	// example:
	//
	// ADVANCED
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 示例值示例值
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
	// example:
	//
	// 示例值
	ChildTaskDescription *string `json:"ChildTaskDescription,omitempty" xml:"ChildTaskDescription,omitempty"`
	// example:
	//
	// 示例值示例值
	ChildTaskName *string `json:"ChildTaskName,omitempty" xml:"ChildTaskName,omitempty"`
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
	// example:
	//
	// true
	ExceptionSign *bool `json:"ExceptionSign,omitempty" xml:"ExceptionSign,omitempty"`
	// example:
	//
	// -
	ExceptionType *string `json:"ExceptionType,omitempty" xml:"ExceptionType,omitempty"`
	// example:
	//
	// -
	Reply *string `json:"Reply,omitempty" xml:"Reply,omitempty"`
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
	// example:
	//
	// 1ee6e994-08e0-5e87-f662-1659cc54d409
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 示例值示例值
	MainPurposeDescription *string                                                                                     `json:"MainPurposeDescription,omitempty" xml:"MainPurposeDescription,omitempty"`
	MainPurposeEnum        []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigMainPurposeMainPurposeEnum `json:"MainPurposeEnum,omitempty" xml:"MainPurposeEnum,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值
	MainPurposeName *string `json:"MainPurposeName,omitempty" xml:"MainPurposeName,omitempty"`
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
	// example:
	//
	// 示例
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// d5606d80-7625-dcea-8610-17f66fbb564a
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// example:
	//
	// 8757************2c499fa
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 示例值示例值
	OutputTagDescription *string                                                                                       `json:"OutputTagDescription,omitempty" xml:"OutputTagDescription,omitempty"`
	OutputTagEnum        []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigOutputTagConfigOutputTagEnum `json:"OutputTagEnum,omitempty" xml:"OutputTagEnum,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值
	OutputTagName *string `json:"OutputTagName,omitempty" xml:"OutputTagName,omitempty"`
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
	// example:
	//
	// 示例值示例值
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 8757************2c499fa
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 示例值示例值示例值
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
	// example:
	//
	// 280cd4bf-*******df472c
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 示例值示例值
	PhoneTagDescription *string                                                                                     `json:"PhoneTagDescription,omitempty" xml:"PhoneTagDescription,omitempty"`
	PhoneTagEnum        []*QueryAiVoiceAgentDetailResponseBodyDataAiVoiceAgentModelConfigPhoneTagConfigPhoneTagEnum `json:"PhoneTagEnum,omitempty" xml:"PhoneTagEnum,omitempty" type:"Repeated"`
	// example:
	//
	// testParam
	PhoneTagKey *string `json:"PhoneTagKey,omitempty" xml:"PhoneTagKey,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	PhoneTagName *string `json:"PhoneTagName,omitempty" xml:"PhoneTagName,omitempty"`
	// example:
	//
	// true
	PhoneTagRequired *bool `json:"PhoneTagRequired,omitempty" xml:"PhoneTagRequired,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	PhoneTagSource *string `json:"PhoneTagSource,omitempty" xml:"PhoneTagSource,omitempty"`
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
	// example:
	//
	// 示例值示例值示例值
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// c3d4ff4e-*********bc26dc044682
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 示例值示例值示例值
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
