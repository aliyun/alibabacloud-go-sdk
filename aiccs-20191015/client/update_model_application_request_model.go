// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCode(v string) *UpdateModelApplicationRequest
	GetApplicationCode() *string
	SetApplicationCps(v int64) *UpdateModelApplicationRequest
	GetApplicationCps() *int64
	SetApplicationName(v string) *UpdateModelApplicationRequest
	GetApplicationName() *string
	SetCallAssistantHangup(v bool) *UpdateModelApplicationRequest
	GetCallAssistantHangup() *bool
	SetCallAssistantRecognize(v bool) *UpdateModelApplicationRequest
	GetCallAssistantRecognize() *bool
	SetCallConnectedTriggerModel(v bool) *UpdateModelApplicationRequest
	GetCallConnectedTriggerModel() *bool
	SetDtmfAllowedDigits(v string) *UpdateModelApplicationRequest
	GetDtmfAllowedDigits() *string
	SetDtmfAutoValidateEnable(v bool) *UpdateModelApplicationRequest
	GetDtmfAutoValidateEnable() *bool
	SetDtmfDigitCount(v int64) *UpdateModelApplicationRequest
	GetDtmfDigitCount() *int64
	SetDtmfInputTimeout(v int64) *UpdateModelApplicationRequest
	GetDtmfInputTimeout() *int64
	SetDtmfOutOfRangeAction(v string) *UpdateModelApplicationRequest
	GetDtmfOutOfRangeAction() *string
	SetDtmfRetryPlayTimes(v int64) *UpdateModelApplicationRequest
	GetDtmfRetryPlayTimes() *int64
	SetDtmfRetryPromptText(v string) *UpdateModelApplicationRequest
	GetDtmfRetryPromptText() *string
	SetDyvmsSceneName(v string) *UpdateModelApplicationRequest
	GetDyvmsSceneName() *string
	SetEnableDtmfReceive(v bool) *UpdateModelApplicationRequest
	GetEnableDtmfReceive() *bool
	SetEnableMorse(v bool) *UpdateModelApplicationRequest
	GetEnableMorse() *bool
	SetInterruptConfig(v *UpdateModelApplicationRequestInterruptConfig) *UpdateModelApplicationRequest
	GetInterruptConfig() *UpdateModelApplicationRequestInterruptConfig
	SetModelCode(v string) *UpdateModelApplicationRequest
	GetModelCode() *string
	SetModelVersion(v string) *UpdateModelApplicationRequest
	GetModelVersion() *string
	SetMuteActive(v bool) *UpdateModelApplicationRequest
	GetMuteActive() *bool
	SetMuteDuration(v int64) *UpdateModelApplicationRequest
	GetMuteDuration() *int64
	SetMuteHangupNum(v int64) *UpdateModelApplicationRequest
	GetMuteHangupNum() *int64
	SetOwnerId(v int64) *UpdateModelApplicationRequest
	GetOwnerId() *int64
	SetPrompt(v string) *UpdateModelApplicationRequest
	GetPrompt() *string
	SetQualificationId(v int64) *UpdateModelApplicationRequest
	GetQualificationId() *int64
	SetQualificationName(v string) *UpdateModelApplicationRequest
	GetQualificationName() *string
	SetRecordingFile(v string) *UpdateModelApplicationRequest
	GetRecordingFile() *string
	SetResourceOwnerAccount(v string) *UpdateModelApplicationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateModelApplicationRequest
	GetResourceOwnerId() *int64
	SetSessionTimeout(v int64) *UpdateModelApplicationRequest
	GetSessionTimeout() *int64
	SetSource(v string) *UpdateModelApplicationRequest
	GetSource() *string
	SetSpeechContent(v string) *UpdateModelApplicationRequest
	GetSpeechContent() *string
	SetSpeechId(v int64) *UpdateModelApplicationRequest
	GetSpeechId() *int64
	SetStartWord(v string) *UpdateModelApplicationRequest
	GetStartWord() *string
	SetStartWordType(v int64) *UpdateModelApplicationRequest
	GetStartWordType() *int64
	SetTtsConfig(v *UpdateModelApplicationRequestTtsConfig) *UpdateModelApplicationRequest
	GetTtsConfig() *UpdateModelApplicationRequestTtsConfig
	SetUsageDesc(v string) *UpdateModelApplicationRequest
	GetUsageDesc() *string
}

type UpdateModelApplicationRequest struct {
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
	InterruptConfig *UpdateModelApplicationRequestInterruptConfig `json:"InterruptConfig,omitempty" xml:"InterruptConfig,omitempty" type:"Struct"`
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
	TtsConfig *UpdateModelApplicationRequestTtsConfig `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty" type:"Struct"`
	// 用途
	//
	// example:
	//
	// 示例值示例值
	UsageDesc *string `json:"UsageDesc,omitempty" xml:"UsageDesc,omitempty"`
}

func (s UpdateModelApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelApplicationRequest) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *UpdateModelApplicationRequest) GetApplicationCps() *int64 {
	return s.ApplicationCps
}

func (s *UpdateModelApplicationRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *UpdateModelApplicationRequest) GetCallAssistantHangup() *bool {
	return s.CallAssistantHangup
}

func (s *UpdateModelApplicationRequest) GetCallAssistantRecognize() *bool {
	return s.CallAssistantRecognize
}

func (s *UpdateModelApplicationRequest) GetCallConnectedTriggerModel() *bool {
	return s.CallConnectedTriggerModel
}

func (s *UpdateModelApplicationRequest) GetDtmfAllowedDigits() *string {
	return s.DtmfAllowedDigits
}

func (s *UpdateModelApplicationRequest) GetDtmfAutoValidateEnable() *bool {
	return s.DtmfAutoValidateEnable
}

func (s *UpdateModelApplicationRequest) GetDtmfDigitCount() *int64 {
	return s.DtmfDigitCount
}

func (s *UpdateModelApplicationRequest) GetDtmfInputTimeout() *int64 {
	return s.DtmfInputTimeout
}

func (s *UpdateModelApplicationRequest) GetDtmfOutOfRangeAction() *string {
	return s.DtmfOutOfRangeAction
}

func (s *UpdateModelApplicationRequest) GetDtmfRetryPlayTimes() *int64 {
	return s.DtmfRetryPlayTimes
}

func (s *UpdateModelApplicationRequest) GetDtmfRetryPromptText() *string {
	return s.DtmfRetryPromptText
}

func (s *UpdateModelApplicationRequest) GetDyvmsSceneName() *string {
	return s.DyvmsSceneName
}

func (s *UpdateModelApplicationRequest) GetEnableDtmfReceive() *bool {
	return s.EnableDtmfReceive
}

func (s *UpdateModelApplicationRequest) GetEnableMorse() *bool {
	return s.EnableMorse
}

func (s *UpdateModelApplicationRequest) GetInterruptConfig() *UpdateModelApplicationRequestInterruptConfig {
	return s.InterruptConfig
}

func (s *UpdateModelApplicationRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *UpdateModelApplicationRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *UpdateModelApplicationRequest) GetMuteActive() *bool {
	return s.MuteActive
}

func (s *UpdateModelApplicationRequest) GetMuteDuration() *int64 {
	return s.MuteDuration
}

func (s *UpdateModelApplicationRequest) GetMuteHangupNum() *int64 {
	return s.MuteHangupNum
}

func (s *UpdateModelApplicationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateModelApplicationRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *UpdateModelApplicationRequest) GetQualificationId() *int64 {
	return s.QualificationId
}

func (s *UpdateModelApplicationRequest) GetQualificationName() *string {
	return s.QualificationName
}

func (s *UpdateModelApplicationRequest) GetRecordingFile() *string {
	return s.RecordingFile
}

func (s *UpdateModelApplicationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateModelApplicationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateModelApplicationRequest) GetSessionTimeout() *int64 {
	return s.SessionTimeout
}

func (s *UpdateModelApplicationRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateModelApplicationRequest) GetSpeechContent() *string {
	return s.SpeechContent
}

func (s *UpdateModelApplicationRequest) GetSpeechId() *int64 {
	return s.SpeechId
}

func (s *UpdateModelApplicationRequest) GetStartWord() *string {
	return s.StartWord
}

func (s *UpdateModelApplicationRequest) GetStartWordType() *int64 {
	return s.StartWordType
}

func (s *UpdateModelApplicationRequest) GetTtsConfig() *UpdateModelApplicationRequestTtsConfig {
	return s.TtsConfig
}

func (s *UpdateModelApplicationRequest) GetUsageDesc() *string {
	return s.UsageDesc
}

func (s *UpdateModelApplicationRequest) SetApplicationCode(v string) *UpdateModelApplicationRequest {
	s.ApplicationCode = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetApplicationCps(v int64) *UpdateModelApplicationRequest {
	s.ApplicationCps = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetApplicationName(v string) *UpdateModelApplicationRequest {
	s.ApplicationName = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetCallAssistantHangup(v bool) *UpdateModelApplicationRequest {
	s.CallAssistantHangup = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetCallAssistantRecognize(v bool) *UpdateModelApplicationRequest {
	s.CallAssistantRecognize = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetCallConnectedTriggerModel(v bool) *UpdateModelApplicationRequest {
	s.CallConnectedTriggerModel = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetDtmfAllowedDigits(v string) *UpdateModelApplicationRequest {
	s.DtmfAllowedDigits = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetDtmfAutoValidateEnable(v bool) *UpdateModelApplicationRequest {
	s.DtmfAutoValidateEnable = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetDtmfDigitCount(v int64) *UpdateModelApplicationRequest {
	s.DtmfDigitCount = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetDtmfInputTimeout(v int64) *UpdateModelApplicationRequest {
	s.DtmfInputTimeout = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetDtmfOutOfRangeAction(v string) *UpdateModelApplicationRequest {
	s.DtmfOutOfRangeAction = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetDtmfRetryPlayTimes(v int64) *UpdateModelApplicationRequest {
	s.DtmfRetryPlayTimes = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetDtmfRetryPromptText(v string) *UpdateModelApplicationRequest {
	s.DtmfRetryPromptText = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetDyvmsSceneName(v string) *UpdateModelApplicationRequest {
	s.DyvmsSceneName = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetEnableDtmfReceive(v bool) *UpdateModelApplicationRequest {
	s.EnableDtmfReceive = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetEnableMorse(v bool) *UpdateModelApplicationRequest {
	s.EnableMorse = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetInterruptConfig(v *UpdateModelApplicationRequestInterruptConfig) *UpdateModelApplicationRequest {
	s.InterruptConfig = v
	return s
}

func (s *UpdateModelApplicationRequest) SetModelCode(v string) *UpdateModelApplicationRequest {
	s.ModelCode = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetModelVersion(v string) *UpdateModelApplicationRequest {
	s.ModelVersion = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetMuteActive(v bool) *UpdateModelApplicationRequest {
	s.MuteActive = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetMuteDuration(v int64) *UpdateModelApplicationRequest {
	s.MuteDuration = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetMuteHangupNum(v int64) *UpdateModelApplicationRequest {
	s.MuteHangupNum = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetOwnerId(v int64) *UpdateModelApplicationRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetPrompt(v string) *UpdateModelApplicationRequest {
	s.Prompt = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetQualificationId(v int64) *UpdateModelApplicationRequest {
	s.QualificationId = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetQualificationName(v string) *UpdateModelApplicationRequest {
	s.QualificationName = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetRecordingFile(v string) *UpdateModelApplicationRequest {
	s.RecordingFile = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetResourceOwnerAccount(v string) *UpdateModelApplicationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetResourceOwnerId(v int64) *UpdateModelApplicationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetSessionTimeout(v int64) *UpdateModelApplicationRequest {
	s.SessionTimeout = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetSource(v string) *UpdateModelApplicationRequest {
	s.Source = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetSpeechContent(v string) *UpdateModelApplicationRequest {
	s.SpeechContent = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetSpeechId(v int64) *UpdateModelApplicationRequest {
	s.SpeechId = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetStartWord(v string) *UpdateModelApplicationRequest {
	s.StartWord = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetStartWordType(v int64) *UpdateModelApplicationRequest {
	s.StartWordType = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetTtsConfig(v *UpdateModelApplicationRequestTtsConfig) *UpdateModelApplicationRequest {
	s.TtsConfig = v
	return s
}

func (s *UpdateModelApplicationRequest) SetUsageDesc(v string) *UpdateModelApplicationRequest {
	s.UsageDesc = &v
	return s
}

func (s *UpdateModelApplicationRequest) Validate() error {
	if s.InterruptConfig != nil {
		if err := s.InterruptConfig.Validate(); err != nil {
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

type UpdateModelApplicationRequestInterruptConfig struct {
	// 防止连续抢话功能配置
	AvoidInterruptDTO *UpdateModelApplicationRequestInterruptConfigAvoidInterruptDTO `json:"AvoidInterruptDTO,omitempty" xml:"AvoidInterruptDTO,omitempty" type:"Struct"`
	// 防止连续抢话功能是否开启
	//
	// example:
	//
	// true
	EnableAvoidInterrupt *bool `json:"EnableAvoidInterrupt,omitempty" xml:"EnableAvoidInterrupt,omitempty"`
	// example:
	//
	// true
	EnableInterruptBackchannel *bool `json:"EnableInterruptBackchannel,omitempty" xml:"EnableInterruptBackchannel,omitempty"`
	// 开场白全程不打断
	//
	// example:
	//
	// true
	EnableStartwordEntireNotInterrupt *bool `json:"EnableStartwordEntireNotInterrupt,omitempty" xml:"EnableStartwordEntireNotInterrupt,omitempty"`
	// 开场白不打断配置是否开启
	EnableStartwordNotInterrupt *bool `json:"EnableStartwordNotInterrupt,omitempty" xml:"EnableStartwordNotInterrupt,omitempty"`
	// 开场白保护时长
	//
	// example:
	//
	// 1.4699
	StartwordProtectDuration *float64 `json:"StartwordProtectDuration,omitempty" xml:"StartwordProtectDuration,omitempty"`
}

func (s UpdateModelApplicationRequestInterruptConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelApplicationRequestInterruptConfig) GoString() string {
	return s.String()
}

func (s *UpdateModelApplicationRequestInterruptConfig) GetAvoidInterruptDTO() *UpdateModelApplicationRequestInterruptConfigAvoidInterruptDTO {
	return s.AvoidInterruptDTO
}

func (s *UpdateModelApplicationRequestInterruptConfig) GetEnableAvoidInterrupt() *bool {
	return s.EnableAvoidInterrupt
}

func (s *UpdateModelApplicationRequestInterruptConfig) GetEnableInterruptBackchannel() *bool {
	return s.EnableInterruptBackchannel
}

func (s *UpdateModelApplicationRequestInterruptConfig) GetEnableStartwordEntireNotInterrupt() *bool {
	return s.EnableStartwordEntireNotInterrupt
}

func (s *UpdateModelApplicationRequestInterruptConfig) GetEnableStartwordNotInterrupt() *bool {
	return s.EnableStartwordNotInterrupt
}

func (s *UpdateModelApplicationRequestInterruptConfig) GetStartwordProtectDuration() *float64 {
	return s.StartwordProtectDuration
}

func (s *UpdateModelApplicationRequestInterruptConfig) SetAvoidInterruptDTO(v *UpdateModelApplicationRequestInterruptConfigAvoidInterruptDTO) *UpdateModelApplicationRequestInterruptConfig {
	s.AvoidInterruptDTO = v
	return s
}

func (s *UpdateModelApplicationRequestInterruptConfig) SetEnableAvoidInterrupt(v bool) *UpdateModelApplicationRequestInterruptConfig {
	s.EnableAvoidInterrupt = &v
	return s
}

func (s *UpdateModelApplicationRequestInterruptConfig) SetEnableInterruptBackchannel(v bool) *UpdateModelApplicationRequestInterruptConfig {
	s.EnableInterruptBackchannel = &v
	return s
}

func (s *UpdateModelApplicationRequestInterruptConfig) SetEnableStartwordEntireNotInterrupt(v bool) *UpdateModelApplicationRequestInterruptConfig {
	s.EnableStartwordEntireNotInterrupt = &v
	return s
}

func (s *UpdateModelApplicationRequestInterruptConfig) SetEnableStartwordNotInterrupt(v bool) *UpdateModelApplicationRequestInterruptConfig {
	s.EnableStartwordNotInterrupt = &v
	return s
}

func (s *UpdateModelApplicationRequestInterruptConfig) SetStartwordProtectDuration(v float64) *UpdateModelApplicationRequestInterruptConfig {
	s.StartwordProtectDuration = &v
	return s
}

func (s *UpdateModelApplicationRequestInterruptConfig) Validate() error {
	if s.AvoidInterruptDTO != nil {
		if err := s.AvoidInterruptDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateModelApplicationRequestInterruptConfigAvoidInterruptDTO struct {
	// example:
	//
	// 3
	InterruptNum *int64 `json:"InterruptNum,omitempty" xml:"InterruptNum,omitempty"`
	// example:
	//
	// 16.417547
	InterruptProtectDuration *float64 `json:"InterruptProtectDuration,omitempty" xml:"InterruptProtectDuration,omitempty"`
}

func (s UpdateModelApplicationRequestInterruptConfigAvoidInterruptDTO) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelApplicationRequestInterruptConfigAvoidInterruptDTO) GoString() string {
	return s.String()
}

func (s *UpdateModelApplicationRequestInterruptConfigAvoidInterruptDTO) GetInterruptNum() *int64 {
	return s.InterruptNum
}

func (s *UpdateModelApplicationRequestInterruptConfigAvoidInterruptDTO) GetInterruptProtectDuration() *float64 {
	return s.InterruptProtectDuration
}

func (s *UpdateModelApplicationRequestInterruptConfigAvoidInterruptDTO) SetInterruptNum(v int64) *UpdateModelApplicationRequestInterruptConfigAvoidInterruptDTO {
	s.InterruptNum = &v
	return s
}

func (s *UpdateModelApplicationRequestInterruptConfigAvoidInterruptDTO) SetInterruptProtectDuration(v float64) *UpdateModelApplicationRequestInterruptConfigAvoidInterruptDTO {
	s.InterruptProtectDuration = &v
	return s
}

func (s *UpdateModelApplicationRequestInterruptConfigAvoidInterruptDTO) Validate() error {
	return dara.Validate(s)
}

type UpdateModelApplicationRequestTtsConfig struct {
	// example:
	//
	// true
	BackgroundEnabled *bool `json:"BackgroundEnabled,omitempty" xml:"BackgroundEnabled,omitempty"`
	// 背景音id
	//
	// example:
	//
	// 111
	BackgroundSound *int64 `json:"BackgroundSound,omitempty" xml:"BackgroundSound,omitempty"`
	// 背景音音量(id)
	//
	// example:
	//
	// 1
	BackgroundVolume *int64 `json:"BackgroundVolume,omitempty" xml:"BackgroundVolume,omitempty"`
	// example:
	//
	// 47
	CustomerAccountId *int64 `json:"CustomerAccountId,omitempty" xml:"CustomerAccountId,omitempty"`
	// example:
	//
	// true
	MixingEnabled *bool `json:"MixingEnabled,omitempty" xml:"MixingEnabled,omitempty"`
	// 混音模版id
	//
	// example:
	//
	// 111
	MixingTemplate *int64 `json:"MixingTemplate,omitempty" xml:"MixingTemplate,omitempty"`
	// example:
	//
	// 示例值示例值
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// TTS 变量播放时的声音速度。取值范围：-200~200，默认值为 0。
	//
	// example:
	//
	// 7
	TtsSpeed *int64 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	// 声音风格
	//
	// example:
	//
	// voice
	TtsStyle *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	// TTS 变量播放的音量。取值范围：0~100，默认值为 0。
	//
	// example:
	//
	// 11
	TtsVolume *int64 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	// 声音编码
	//
	// example:
	//
	// 示例值示例值
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// 声音类型
	//
	// example:
	//
	// 示例值示例值
	VoiceType *string `json:"VoiceType,omitempty" xml:"VoiceType,omitempty"`
}

func (s UpdateModelApplicationRequestTtsConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelApplicationRequestTtsConfig) GoString() string {
	return s.String()
}

func (s *UpdateModelApplicationRequestTtsConfig) GetBackgroundEnabled() *bool {
	return s.BackgroundEnabled
}

func (s *UpdateModelApplicationRequestTtsConfig) GetBackgroundSound() *int64 {
	return s.BackgroundSound
}

func (s *UpdateModelApplicationRequestTtsConfig) GetBackgroundVolume() *int64 {
	return s.BackgroundVolume
}

func (s *UpdateModelApplicationRequestTtsConfig) GetCustomerAccountId() *int64 {
	return s.CustomerAccountId
}

func (s *UpdateModelApplicationRequestTtsConfig) GetMixingEnabled() *bool {
	return s.MixingEnabled
}

func (s *UpdateModelApplicationRequestTtsConfig) GetMixingTemplate() *int64 {
	return s.MixingTemplate
}

func (s *UpdateModelApplicationRequestTtsConfig) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateModelApplicationRequestTtsConfig) GetTtsSpeed() *int64 {
	return s.TtsSpeed
}

func (s *UpdateModelApplicationRequestTtsConfig) GetTtsStyle() *string {
	return s.TtsStyle
}

func (s *UpdateModelApplicationRequestTtsConfig) GetTtsVolume() *int64 {
	return s.TtsVolume
}

func (s *UpdateModelApplicationRequestTtsConfig) GetVoiceCode() *string {
	return s.VoiceCode
}

func (s *UpdateModelApplicationRequestTtsConfig) GetVoiceType() *string {
	return s.VoiceType
}

func (s *UpdateModelApplicationRequestTtsConfig) SetBackgroundEnabled(v bool) *UpdateModelApplicationRequestTtsConfig {
	s.BackgroundEnabled = &v
	return s
}

func (s *UpdateModelApplicationRequestTtsConfig) SetBackgroundSound(v int64) *UpdateModelApplicationRequestTtsConfig {
	s.BackgroundSound = &v
	return s
}

func (s *UpdateModelApplicationRequestTtsConfig) SetBackgroundVolume(v int64) *UpdateModelApplicationRequestTtsConfig {
	s.BackgroundVolume = &v
	return s
}

func (s *UpdateModelApplicationRequestTtsConfig) SetCustomerAccountId(v int64) *UpdateModelApplicationRequestTtsConfig {
	s.CustomerAccountId = &v
	return s
}

func (s *UpdateModelApplicationRequestTtsConfig) SetMixingEnabled(v bool) *UpdateModelApplicationRequestTtsConfig {
	s.MixingEnabled = &v
	return s
}

func (s *UpdateModelApplicationRequestTtsConfig) SetMixingTemplate(v int64) *UpdateModelApplicationRequestTtsConfig {
	s.MixingTemplate = &v
	return s
}

func (s *UpdateModelApplicationRequestTtsConfig) SetResourceId(v string) *UpdateModelApplicationRequestTtsConfig {
	s.ResourceId = &v
	return s
}

func (s *UpdateModelApplicationRequestTtsConfig) SetTtsSpeed(v int64) *UpdateModelApplicationRequestTtsConfig {
	s.TtsSpeed = &v
	return s
}

func (s *UpdateModelApplicationRequestTtsConfig) SetTtsStyle(v string) *UpdateModelApplicationRequestTtsConfig {
	s.TtsStyle = &v
	return s
}

func (s *UpdateModelApplicationRequestTtsConfig) SetTtsVolume(v int64) *UpdateModelApplicationRequestTtsConfig {
	s.TtsVolume = &v
	return s
}

func (s *UpdateModelApplicationRequestTtsConfig) SetVoiceCode(v string) *UpdateModelApplicationRequestTtsConfig {
	s.VoiceCode = &v
	return s
}

func (s *UpdateModelApplicationRequestTtsConfig) SetVoiceType(v string) *UpdateModelApplicationRequestTtsConfig {
	s.VoiceType = &v
	return s
}

func (s *UpdateModelApplicationRequestTtsConfig) Validate() error {
	return dara.Validate(s)
}
