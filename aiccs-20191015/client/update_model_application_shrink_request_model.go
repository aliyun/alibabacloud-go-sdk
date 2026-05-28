// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCode(v string) *UpdateModelApplicationShrinkRequest
	GetApplicationCode() *string
	SetApplicationCps(v int64) *UpdateModelApplicationShrinkRequest
	GetApplicationCps() *int64
	SetApplicationName(v string) *UpdateModelApplicationShrinkRequest
	GetApplicationName() *string
	SetCallAssistantHangup(v bool) *UpdateModelApplicationShrinkRequest
	GetCallAssistantHangup() *bool
	SetCallAssistantRecognize(v bool) *UpdateModelApplicationShrinkRequest
	GetCallAssistantRecognize() *bool
	SetCallConnectedTriggerModel(v bool) *UpdateModelApplicationShrinkRequest
	GetCallConnectedTriggerModel() *bool
	SetDtmfAllowedDigits(v string) *UpdateModelApplicationShrinkRequest
	GetDtmfAllowedDigits() *string
	SetDtmfAutoValidateEnable(v bool) *UpdateModelApplicationShrinkRequest
	GetDtmfAutoValidateEnable() *bool
	SetDtmfDigitCount(v int64) *UpdateModelApplicationShrinkRequest
	GetDtmfDigitCount() *int64
	SetDtmfInputTimeout(v int64) *UpdateModelApplicationShrinkRequest
	GetDtmfInputTimeout() *int64
	SetDtmfOutOfRangeAction(v string) *UpdateModelApplicationShrinkRequest
	GetDtmfOutOfRangeAction() *string
	SetDtmfRetryPlayTimes(v int64) *UpdateModelApplicationShrinkRequest
	GetDtmfRetryPlayTimes() *int64
	SetDtmfRetryPromptText(v string) *UpdateModelApplicationShrinkRequest
	GetDtmfRetryPromptText() *string
	SetDyvmsSceneName(v string) *UpdateModelApplicationShrinkRequest
	GetDyvmsSceneName() *string
	SetEnableDtmfReceive(v bool) *UpdateModelApplicationShrinkRequest
	GetEnableDtmfReceive() *bool
	SetEnableMorse(v bool) *UpdateModelApplicationShrinkRequest
	GetEnableMorse() *bool
	SetInterruptConfigShrink(v string) *UpdateModelApplicationShrinkRequest
	GetInterruptConfigShrink() *string
	SetModelCode(v string) *UpdateModelApplicationShrinkRequest
	GetModelCode() *string
	SetModelVersion(v string) *UpdateModelApplicationShrinkRequest
	GetModelVersion() *string
	SetMuteActive(v bool) *UpdateModelApplicationShrinkRequest
	GetMuteActive() *bool
	SetMuteDuration(v int64) *UpdateModelApplicationShrinkRequest
	GetMuteDuration() *int64
	SetMuteHangupNum(v int64) *UpdateModelApplicationShrinkRequest
	GetMuteHangupNum() *int64
	SetOwnerId(v int64) *UpdateModelApplicationShrinkRequest
	GetOwnerId() *int64
	SetPrompt(v string) *UpdateModelApplicationShrinkRequest
	GetPrompt() *string
	SetQualificationId(v int64) *UpdateModelApplicationShrinkRequest
	GetQualificationId() *int64
	SetQualificationName(v string) *UpdateModelApplicationShrinkRequest
	GetQualificationName() *string
	SetRecordingFile(v string) *UpdateModelApplicationShrinkRequest
	GetRecordingFile() *string
	SetResourceOwnerAccount(v string) *UpdateModelApplicationShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateModelApplicationShrinkRequest
	GetResourceOwnerId() *int64
	SetSessionTimeout(v int64) *UpdateModelApplicationShrinkRequest
	GetSessionTimeout() *int64
	SetSource(v string) *UpdateModelApplicationShrinkRequest
	GetSource() *string
	SetSpeechContent(v string) *UpdateModelApplicationShrinkRequest
	GetSpeechContent() *string
	SetSpeechId(v int64) *UpdateModelApplicationShrinkRequest
	GetSpeechId() *int64
	SetStartWord(v string) *UpdateModelApplicationShrinkRequest
	GetStartWord() *string
	SetStartWordType(v int64) *UpdateModelApplicationShrinkRequest
	GetStartWordType() *int64
	SetTtsConfigShrink(v string) *UpdateModelApplicationShrinkRequest
	GetTtsConfigShrink() *string
	SetUsageDesc(v string) *UpdateModelApplicationShrinkRequest
	GetUsageDesc() *string
}

type UpdateModelApplicationShrinkRequest struct {
	// 应用编码
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// 应用并发请求数
	//
	// example:
	//
	// 12
	ApplicationCps *int64 `json:"ApplicationCps,omitempty" xml:"ApplicationCps,omitempty"`
	// 模型应用名称
	//
	// example:
	//
	// 示例值示例值
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// example:
	//
	// false
	CallAssistantHangup *bool `json:"CallAssistantHangup,omitempty" xml:"CallAssistantHangup,omitempty"`
	// 通话助手识别
	//
	// example:
	//
	// 示例值
	CallAssistantRecognize *bool `json:"CallAssistantRecognize,omitempty" xml:"CallAssistantRecognize,omitempty"`
	// example:
	//
	// false
	CallConnectedTriggerModel *bool `json:"CallConnectedTriggerModel,omitempty" xml:"CallConnectedTriggerModel,omitempty"`
	// example:
	//
	// 示例值
	DtmfAllowedDigits *string `json:"DtmfAllowedDigits,omitempty" xml:"DtmfAllowedDigits,omitempty"`
	// example:
	//
	// true
	DtmfAutoValidateEnable *bool `json:"DtmfAutoValidateEnable,omitempty" xml:"DtmfAutoValidateEnable,omitempty"`
	// example:
	//
	// 1
	DtmfDigitCount *int64 `json:"DtmfDigitCount,omitempty" xml:"DtmfDigitCount,omitempty"`
	// example:
	//
	// 1
	DtmfInputTimeout *int64 `json:"DtmfInputTimeout,omitempty" xml:"DtmfInputTimeout,omitempty"`
	// example:
	//
	// RETURN_MODEL
	DtmfOutOfRangeAction *string `json:"DtmfOutOfRangeAction,omitempty" xml:"DtmfOutOfRangeAction,omitempty"`
	// example:
	//
	// 1
	DtmfRetryPlayTimes *int64 `json:"DtmfRetryPlayTimes,omitempty" xml:"DtmfRetryPlayTimes,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	DtmfRetryPromptText *string `json:"DtmfRetryPromptText,omitempty" xml:"DtmfRetryPromptText,omitempty"`
	// 场景名称
	//
	// example:
	//
	// 示例值示例值
	DyvmsSceneName *string `json:"DyvmsSceneName,omitempty" xml:"DyvmsSceneName,omitempty"`
	// example:
	//
	// false
	EnableDtmfReceive *bool `json:"EnableDtmfReceive,omitempty" xml:"EnableDtmfReceive,omitempty"`
	// example:
	//
	// false
	EnableMorse *bool `json:"EnableMorse,omitempty" xml:"EnableMorse,omitempty"`
	// 打断配置
	InterruptConfigShrink *string `json:"InterruptConfig,omitempty" xml:"InterruptConfig,omitempty"`
	// 模型编码
	//
	// example:
	//
	// 示例值示例值示例值
	ModelCode *string `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	// 模型版本
	//
	// example:
	//
	// 示例值示例值示例值
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// 第一个静音是否唤起模型
	//
	// example:
	//
	// true
	MuteActive *bool `json:"MuteActive,omitempty" xml:"MuteActive,omitempty"`
	// 静音时长
	//
	// example:
	//
	// 85
	MuteDuration *int64 `json:"MuteDuration,omitempty" xml:"MuteDuration,omitempty"`
	// 连续多少个静音事件主动挂机
	//
	// example:
	//
	// 70
	MuteHangupNum *int64 `json:"MuteHangupNum,omitempty" xml:"MuteHangupNum,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 提示词
	//
	// example:
	//
	// 示例值示例值
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// 资质ID
	//
	// example:
	//
	// 61
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// 资质名称
	//
	// example:
	//
	// 示例值示例值示例值
	QualificationName *string `json:"QualificationName,omitempty" xml:"QualificationName,omitempty"`
	// example:
	//
	// 示例值示例值
	RecordingFile        *string `json:"RecordingFile,omitempty" xml:"RecordingFile,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 最大通话时长
	//
	// example:
	//
	// 49
	SessionTimeout *int64 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// 来源
	//
	// example:
	//
	// 示例值示例值
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 话术内容
	//
	// example:
	//
	// 示例值
	SpeechContent *string `json:"SpeechContent,omitempty" xml:"SpeechContent,omitempty"`
	// 话束id
	//
	// example:
	//
	// 15
	SpeechId *int64 `json:"SpeechId,omitempty" xml:"SpeechId,omitempty"`
	// 开场白
	//
	// example:
	//
	// 示例值示例值
	StartWord *string `json:"StartWord,omitempty" xml:"StartWord,omitempty"`
	// example:
	//
	// 1
	StartWordType *int64 `json:"StartWordType,omitempty" xml:"StartWordType,omitempty"`
	// tts配置，包括音色、音量、音速等。
	TtsConfigShrink *string `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty"`
	// 用途
	//
	// example:
	//
	// 示例值示例值
	UsageDesc *string `json:"UsageDesc,omitempty" xml:"UsageDesc,omitempty"`
}

func (s UpdateModelApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelApplicationShrinkRequest) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *UpdateModelApplicationShrinkRequest) GetApplicationCps() *int64 {
	return s.ApplicationCps
}

func (s *UpdateModelApplicationShrinkRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *UpdateModelApplicationShrinkRequest) GetCallAssistantHangup() *bool {
	return s.CallAssistantHangup
}

func (s *UpdateModelApplicationShrinkRequest) GetCallAssistantRecognize() *bool {
	return s.CallAssistantRecognize
}

func (s *UpdateModelApplicationShrinkRequest) GetCallConnectedTriggerModel() *bool {
	return s.CallConnectedTriggerModel
}

func (s *UpdateModelApplicationShrinkRequest) GetDtmfAllowedDigits() *string {
	return s.DtmfAllowedDigits
}

func (s *UpdateModelApplicationShrinkRequest) GetDtmfAutoValidateEnable() *bool {
	return s.DtmfAutoValidateEnable
}

func (s *UpdateModelApplicationShrinkRequest) GetDtmfDigitCount() *int64 {
	return s.DtmfDigitCount
}

func (s *UpdateModelApplicationShrinkRequest) GetDtmfInputTimeout() *int64 {
	return s.DtmfInputTimeout
}

func (s *UpdateModelApplicationShrinkRequest) GetDtmfOutOfRangeAction() *string {
	return s.DtmfOutOfRangeAction
}

func (s *UpdateModelApplicationShrinkRequest) GetDtmfRetryPlayTimes() *int64 {
	return s.DtmfRetryPlayTimes
}

func (s *UpdateModelApplicationShrinkRequest) GetDtmfRetryPromptText() *string {
	return s.DtmfRetryPromptText
}

func (s *UpdateModelApplicationShrinkRequest) GetDyvmsSceneName() *string {
	return s.DyvmsSceneName
}

func (s *UpdateModelApplicationShrinkRequest) GetEnableDtmfReceive() *bool {
	return s.EnableDtmfReceive
}

func (s *UpdateModelApplicationShrinkRequest) GetEnableMorse() *bool {
	return s.EnableMorse
}

func (s *UpdateModelApplicationShrinkRequest) GetInterruptConfigShrink() *string {
	return s.InterruptConfigShrink
}

func (s *UpdateModelApplicationShrinkRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *UpdateModelApplicationShrinkRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *UpdateModelApplicationShrinkRequest) GetMuteActive() *bool {
	return s.MuteActive
}

func (s *UpdateModelApplicationShrinkRequest) GetMuteDuration() *int64 {
	return s.MuteDuration
}

func (s *UpdateModelApplicationShrinkRequest) GetMuteHangupNum() *int64 {
	return s.MuteHangupNum
}

func (s *UpdateModelApplicationShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateModelApplicationShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *UpdateModelApplicationShrinkRequest) GetQualificationId() *int64 {
	return s.QualificationId
}

func (s *UpdateModelApplicationShrinkRequest) GetQualificationName() *string {
	return s.QualificationName
}

func (s *UpdateModelApplicationShrinkRequest) GetRecordingFile() *string {
	return s.RecordingFile
}

func (s *UpdateModelApplicationShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateModelApplicationShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateModelApplicationShrinkRequest) GetSessionTimeout() *int64 {
	return s.SessionTimeout
}

func (s *UpdateModelApplicationShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateModelApplicationShrinkRequest) GetSpeechContent() *string {
	return s.SpeechContent
}

func (s *UpdateModelApplicationShrinkRequest) GetSpeechId() *int64 {
	return s.SpeechId
}

func (s *UpdateModelApplicationShrinkRequest) GetStartWord() *string {
	return s.StartWord
}

func (s *UpdateModelApplicationShrinkRequest) GetStartWordType() *int64 {
	return s.StartWordType
}

func (s *UpdateModelApplicationShrinkRequest) GetTtsConfigShrink() *string {
	return s.TtsConfigShrink
}

func (s *UpdateModelApplicationShrinkRequest) GetUsageDesc() *string {
	return s.UsageDesc
}

func (s *UpdateModelApplicationShrinkRequest) SetApplicationCode(v string) *UpdateModelApplicationShrinkRequest {
	s.ApplicationCode = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetApplicationCps(v int64) *UpdateModelApplicationShrinkRequest {
	s.ApplicationCps = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetApplicationName(v string) *UpdateModelApplicationShrinkRequest {
	s.ApplicationName = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetCallAssistantHangup(v bool) *UpdateModelApplicationShrinkRequest {
	s.CallAssistantHangup = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetCallAssistantRecognize(v bool) *UpdateModelApplicationShrinkRequest {
	s.CallAssistantRecognize = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetCallConnectedTriggerModel(v bool) *UpdateModelApplicationShrinkRequest {
	s.CallConnectedTriggerModel = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetDtmfAllowedDigits(v string) *UpdateModelApplicationShrinkRequest {
	s.DtmfAllowedDigits = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetDtmfAutoValidateEnable(v bool) *UpdateModelApplicationShrinkRequest {
	s.DtmfAutoValidateEnable = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetDtmfDigitCount(v int64) *UpdateModelApplicationShrinkRequest {
	s.DtmfDigitCount = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetDtmfInputTimeout(v int64) *UpdateModelApplicationShrinkRequest {
	s.DtmfInputTimeout = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetDtmfOutOfRangeAction(v string) *UpdateModelApplicationShrinkRequest {
	s.DtmfOutOfRangeAction = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetDtmfRetryPlayTimes(v int64) *UpdateModelApplicationShrinkRequest {
	s.DtmfRetryPlayTimes = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetDtmfRetryPromptText(v string) *UpdateModelApplicationShrinkRequest {
	s.DtmfRetryPromptText = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetDyvmsSceneName(v string) *UpdateModelApplicationShrinkRequest {
	s.DyvmsSceneName = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetEnableDtmfReceive(v bool) *UpdateModelApplicationShrinkRequest {
	s.EnableDtmfReceive = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetEnableMorse(v bool) *UpdateModelApplicationShrinkRequest {
	s.EnableMorse = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetInterruptConfigShrink(v string) *UpdateModelApplicationShrinkRequest {
	s.InterruptConfigShrink = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetModelCode(v string) *UpdateModelApplicationShrinkRequest {
	s.ModelCode = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetModelVersion(v string) *UpdateModelApplicationShrinkRequest {
	s.ModelVersion = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetMuteActive(v bool) *UpdateModelApplicationShrinkRequest {
	s.MuteActive = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetMuteDuration(v int64) *UpdateModelApplicationShrinkRequest {
	s.MuteDuration = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetMuteHangupNum(v int64) *UpdateModelApplicationShrinkRequest {
	s.MuteHangupNum = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetOwnerId(v int64) *UpdateModelApplicationShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetPrompt(v string) *UpdateModelApplicationShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetQualificationId(v int64) *UpdateModelApplicationShrinkRequest {
	s.QualificationId = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetQualificationName(v string) *UpdateModelApplicationShrinkRequest {
	s.QualificationName = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetRecordingFile(v string) *UpdateModelApplicationShrinkRequest {
	s.RecordingFile = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetResourceOwnerAccount(v string) *UpdateModelApplicationShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetResourceOwnerId(v int64) *UpdateModelApplicationShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetSessionTimeout(v int64) *UpdateModelApplicationShrinkRequest {
	s.SessionTimeout = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetSource(v string) *UpdateModelApplicationShrinkRequest {
	s.Source = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetSpeechContent(v string) *UpdateModelApplicationShrinkRequest {
	s.SpeechContent = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetSpeechId(v int64) *UpdateModelApplicationShrinkRequest {
	s.SpeechId = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetStartWord(v string) *UpdateModelApplicationShrinkRequest {
	s.StartWord = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetStartWordType(v int64) *UpdateModelApplicationShrinkRequest {
	s.StartWordType = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetTtsConfigShrink(v string) *UpdateModelApplicationShrinkRequest {
	s.TtsConfigShrink = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) SetUsageDesc(v string) *UpdateModelApplicationShrinkRequest {
	s.UsageDesc = &v
	return s
}

func (s *UpdateModelApplicationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
