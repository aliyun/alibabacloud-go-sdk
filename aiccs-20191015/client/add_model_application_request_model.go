// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddModelApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCps(v int64) *AddModelApplicationRequest
	GetApplicationCps() *int64
	SetApplicationName(v string) *AddModelApplicationRequest
	GetApplicationName() *string
	SetCallConnectedTriggerModel(v bool) *AddModelApplicationRequest
	GetCallConnectedTriggerModel() *bool
	SetDyvmsSceneName(v string) *AddModelApplicationRequest
	GetDyvmsSceneName() *string
	SetModelCode(v string) *AddModelApplicationRequest
	GetModelCode() *string
	SetModelVersion(v string) *AddModelApplicationRequest
	GetModelVersion() *string
	SetMuteActive(v bool) *AddModelApplicationRequest
	GetMuteActive() *bool
	SetMuteDuration(v int64) *AddModelApplicationRequest
	GetMuteDuration() *int64
	SetMuteHangupNum(v int64) *AddModelApplicationRequest
	GetMuteHangupNum() *int64
	SetOwnerId(v int64) *AddModelApplicationRequest
	GetOwnerId() *int64
	SetPrompt(v string) *AddModelApplicationRequest
	GetPrompt() *string
	SetQualificationId(v int64) *AddModelApplicationRequest
	GetQualificationId() *int64
	SetQualificationName(v string) *AddModelApplicationRequest
	GetQualificationName() *string
	SetRecordingFile(v string) *AddModelApplicationRequest
	GetRecordingFile() *string
	SetResourceOwnerAccount(v string) *AddModelApplicationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddModelApplicationRequest
	GetResourceOwnerId() *int64
	SetSource(v string) *AddModelApplicationRequest
	GetSource() *string
	SetSpeechContent(v string) *AddModelApplicationRequest
	GetSpeechContent() *string
	SetSpeechId(v int64) *AddModelApplicationRequest
	GetSpeechId() *int64
	SetStartWord(v string) *AddModelApplicationRequest
	GetStartWord() *string
	SetStartWordType(v int64) *AddModelApplicationRequest
	GetStartWordType() *int64
	SetTtsConfig(v *AddModelApplicationRequestTtsConfig) *AddModelApplicationRequest
	GetTtsConfig() *AddModelApplicationRequestTtsConfig
	SetUsageDesc(v string) *AddModelApplicationRequest
	GetUsageDesc() *string
}

type AddModelApplicationRequest struct {
	// The number of concurrent requests per second (CPS).
	//
	// This parameter is required.
	//
	// example:
	//
	// 25
	ApplicationCps *int64 `json:"ApplicationCps,omitempty" xml:"ApplicationCps,omitempty"`
	// The name of the model application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试应用
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// Specifies whether to push an event notification when a call is connected. The default value is false.
	//
	// example:
	//
	// false
	CallConnectedTriggerModel *bool `json:"CallConnectedTriggerModel,omitempty" xml:"CallConnectedTriggerModel,omitempty"`
	// The scene name.
	//
	// example:
	//
	// 测试场景
	DyvmsSceneName *string `json:"DyvmsSceneName,omitempty" xml:"DyvmsSceneName,omitempty"`
	// The model code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ModelCode *string `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	// The model version.
	//
	// example:
	//
	// 1
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// Specifies whether the first mute event triggers the model.
	//
	// example:
	//
	// false
	MuteActive *bool `json:"MuteActive,omitempty" xml:"MuteActive,omitempty"`
	// The mute duration.
	//
	// example:
	//
	// 70
	MuteDuration *int64 `json:"MuteDuration,omitempty" xml:"MuteDuration,omitempty"`
	// The number of consecutive mute events that trigger an automatic hang-up.
	//
	// example:
	//
	// 5
	MuteHangupNum *int64 `json:"MuteHangupNum,omitempty" xml:"MuteHangupNum,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The prompt.
	//
	// example:
	//
	// 测试提示词。
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The qualification ID.
	//
	// example:
	//
	// 85
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// The name of the qualification.
	//
	// example:
	//
	// 测试资质
	QualificationName *string `json:"QualificationName,omitempty" xml:"QualificationName,omitempty"`
	// The URL of the audio file for the opening line. This parameter is required if `StartWordType` is set to `1`.
	//
	// example:
	//
	// https://xxxxxxxx.wav
	RecordingFile        *string `json:"RecordingFile,omitempty" xml:"RecordingFile,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source. The value must be `USER`.
	//
	// example:
	//
	// USER
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The speech script content.
	//
	// example:
	//
	// 测试话术。
	SpeechContent *string `json:"SpeechContent,omitempty" xml:"SpeechContent,omitempty"`
	// The speech script ID.
	//
	// example:
	//
	// 88
	SpeechId *int64 `json:"SpeechId,omitempty" xml:"SpeechId,omitempty"`
	// The opening line.
	//
	// This parameter is required.
	//
	// example:
	//
	// 你好，这是一句开场白。
	StartWord *string `json:"StartWord,omitempty" xml:"StartWord,omitempty"`
	// The type of the opening line.
	//
	// example:
	//
	// 0：文本
	//
	// 1：录音
	StartWordType *int64 `json:"StartWordType,omitempty" xml:"StartWordType,omitempty"`
	// The TTS configuration, including voice, volume, speech speed, and more.
	//
	// This parameter is required.
	TtsConfig *AddModelApplicationRequestTtsConfig `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty" type:"Struct"`
	// The purpose of the application.
	//
	// example:
	//
	// 测试用途
	UsageDesc *string `json:"UsageDesc,omitempty" xml:"UsageDesc,omitempty"`
}

func (s AddModelApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s AddModelApplicationRequest) GoString() string {
	return s.String()
}

func (s *AddModelApplicationRequest) GetApplicationCps() *int64 {
	return s.ApplicationCps
}

func (s *AddModelApplicationRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *AddModelApplicationRequest) GetCallConnectedTriggerModel() *bool {
	return s.CallConnectedTriggerModel
}

func (s *AddModelApplicationRequest) GetDyvmsSceneName() *string {
	return s.DyvmsSceneName
}

func (s *AddModelApplicationRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *AddModelApplicationRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *AddModelApplicationRequest) GetMuteActive() *bool {
	return s.MuteActive
}

func (s *AddModelApplicationRequest) GetMuteDuration() *int64 {
	return s.MuteDuration
}

func (s *AddModelApplicationRequest) GetMuteHangupNum() *int64 {
	return s.MuteHangupNum
}

func (s *AddModelApplicationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddModelApplicationRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *AddModelApplicationRequest) GetQualificationId() *int64 {
	return s.QualificationId
}

func (s *AddModelApplicationRequest) GetQualificationName() *string {
	return s.QualificationName
}

func (s *AddModelApplicationRequest) GetRecordingFile() *string {
	return s.RecordingFile
}

func (s *AddModelApplicationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddModelApplicationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddModelApplicationRequest) GetSource() *string {
	return s.Source
}

func (s *AddModelApplicationRequest) GetSpeechContent() *string {
	return s.SpeechContent
}

func (s *AddModelApplicationRequest) GetSpeechId() *int64 {
	return s.SpeechId
}

func (s *AddModelApplicationRequest) GetStartWord() *string {
	return s.StartWord
}

func (s *AddModelApplicationRequest) GetStartWordType() *int64 {
	return s.StartWordType
}

func (s *AddModelApplicationRequest) GetTtsConfig() *AddModelApplicationRequestTtsConfig {
	return s.TtsConfig
}

func (s *AddModelApplicationRequest) GetUsageDesc() *string {
	return s.UsageDesc
}

func (s *AddModelApplicationRequest) SetApplicationCps(v int64) *AddModelApplicationRequest {
	s.ApplicationCps = &v
	return s
}

func (s *AddModelApplicationRequest) SetApplicationName(v string) *AddModelApplicationRequest {
	s.ApplicationName = &v
	return s
}

func (s *AddModelApplicationRequest) SetCallConnectedTriggerModel(v bool) *AddModelApplicationRequest {
	s.CallConnectedTriggerModel = &v
	return s
}

func (s *AddModelApplicationRequest) SetDyvmsSceneName(v string) *AddModelApplicationRequest {
	s.DyvmsSceneName = &v
	return s
}

func (s *AddModelApplicationRequest) SetModelCode(v string) *AddModelApplicationRequest {
	s.ModelCode = &v
	return s
}

func (s *AddModelApplicationRequest) SetModelVersion(v string) *AddModelApplicationRequest {
	s.ModelVersion = &v
	return s
}

func (s *AddModelApplicationRequest) SetMuteActive(v bool) *AddModelApplicationRequest {
	s.MuteActive = &v
	return s
}

func (s *AddModelApplicationRequest) SetMuteDuration(v int64) *AddModelApplicationRequest {
	s.MuteDuration = &v
	return s
}

func (s *AddModelApplicationRequest) SetMuteHangupNum(v int64) *AddModelApplicationRequest {
	s.MuteHangupNum = &v
	return s
}

func (s *AddModelApplicationRequest) SetOwnerId(v int64) *AddModelApplicationRequest {
	s.OwnerId = &v
	return s
}

func (s *AddModelApplicationRequest) SetPrompt(v string) *AddModelApplicationRequest {
	s.Prompt = &v
	return s
}

func (s *AddModelApplicationRequest) SetQualificationId(v int64) *AddModelApplicationRequest {
	s.QualificationId = &v
	return s
}

func (s *AddModelApplicationRequest) SetQualificationName(v string) *AddModelApplicationRequest {
	s.QualificationName = &v
	return s
}

func (s *AddModelApplicationRequest) SetRecordingFile(v string) *AddModelApplicationRequest {
	s.RecordingFile = &v
	return s
}

func (s *AddModelApplicationRequest) SetResourceOwnerAccount(v string) *AddModelApplicationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddModelApplicationRequest) SetResourceOwnerId(v int64) *AddModelApplicationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddModelApplicationRequest) SetSource(v string) *AddModelApplicationRequest {
	s.Source = &v
	return s
}

func (s *AddModelApplicationRequest) SetSpeechContent(v string) *AddModelApplicationRequest {
	s.SpeechContent = &v
	return s
}

func (s *AddModelApplicationRequest) SetSpeechId(v int64) *AddModelApplicationRequest {
	s.SpeechId = &v
	return s
}

func (s *AddModelApplicationRequest) SetStartWord(v string) *AddModelApplicationRequest {
	s.StartWord = &v
	return s
}

func (s *AddModelApplicationRequest) SetStartWordType(v int64) *AddModelApplicationRequest {
	s.StartWordType = &v
	return s
}

func (s *AddModelApplicationRequest) SetTtsConfig(v *AddModelApplicationRequestTtsConfig) *AddModelApplicationRequest {
	s.TtsConfig = v
	return s
}

func (s *AddModelApplicationRequest) SetUsageDesc(v string) *AddModelApplicationRequest {
	s.UsageDesc = &v
	return s
}

func (s *AddModelApplicationRequest) Validate() error {
	if s.TtsConfig != nil {
		if err := s.TtsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddModelApplicationRequestTtsConfig struct {
	// Specifies whether to enable background sound.
	//
	// example:
	//
	// true
	BackgroundEnabled *bool `json:"BackgroundEnabled,omitempty" xml:"BackgroundEnabled,omitempty"`
	// The background sound ID.
	//
	// example:
	//
	// 111
	BackgroundSound *int64 `json:"BackgroundSound,omitempty" xml:"BackgroundSound,omitempty"`
	// The background sound volume.
	//
	// example:
	//
	// 0：low
	//
	// 1：medium
	//
	// 2：high
	BackgroundVolume *int64 `json:"BackgroundVolume,omitempty" xml:"BackgroundVolume,omitempty"`
	// The account ID.
	//
	// example:
	//
	// 45
	CustomerAccountId *int64 `json:"CustomerAccountId,omitempty" xml:"CustomerAccountId,omitempty"`
	// Specifies whether to enable mixing.
	//
	// example:
	//
	// true
	MixingEnabled *bool `json:"MixingEnabled,omitempty" xml:"MixingEnabled,omitempty"`
	// The mixing template ID.
	//
	// example:
	//
	// 111
	MixingTemplate *int64 `json:"MixingTemplate,omitempty" xml:"MixingTemplate,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// example
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The speech speed for TTS playback. Valid values: -200–200. The default value is 0.
	//
	// example:
	//
	// 13
	TtsSpeed *int64 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	// The voice style.
	//
	// example:
	//
	// 龙小夏
	TtsStyle *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	// The volume for TTS playback. Valid values: 0–100. The default value is 0.
	//
	// example:
	//
	// 55
	TtsVolume *int64 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	// The voice code.
	//
	// example:
	//
	// ddddfd
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// The voice type.
	//
	// ```
	//
	// SYSTEM: System voice
	//
	// COSYCLONE: Cloned voice
	//
	// BL-CUSTOM: Premium custom cloned voice
	//
	// ```
	//
	// example:
	//
	// SYSTEM
	VoiceType *string `json:"VoiceType,omitempty" xml:"VoiceType,omitempty"`
}

func (s AddModelApplicationRequestTtsConfig) String() string {
	return dara.Prettify(s)
}

func (s AddModelApplicationRequestTtsConfig) GoString() string {
	return s.String()
}

func (s *AddModelApplicationRequestTtsConfig) GetBackgroundEnabled() *bool {
	return s.BackgroundEnabled
}

func (s *AddModelApplicationRequestTtsConfig) GetBackgroundSound() *int64 {
	return s.BackgroundSound
}

func (s *AddModelApplicationRequestTtsConfig) GetBackgroundVolume() *int64 {
	return s.BackgroundVolume
}

func (s *AddModelApplicationRequestTtsConfig) GetCustomerAccountId() *int64 {
	return s.CustomerAccountId
}

func (s *AddModelApplicationRequestTtsConfig) GetMixingEnabled() *bool {
	return s.MixingEnabled
}

func (s *AddModelApplicationRequestTtsConfig) GetMixingTemplate() *int64 {
	return s.MixingTemplate
}

func (s *AddModelApplicationRequestTtsConfig) GetResourceId() *string {
	return s.ResourceId
}

func (s *AddModelApplicationRequestTtsConfig) GetTtsSpeed() *int64 {
	return s.TtsSpeed
}

func (s *AddModelApplicationRequestTtsConfig) GetTtsStyle() *string {
	return s.TtsStyle
}

func (s *AddModelApplicationRequestTtsConfig) GetTtsVolume() *int64 {
	return s.TtsVolume
}

func (s *AddModelApplicationRequestTtsConfig) GetVoiceCode() *string {
	return s.VoiceCode
}

func (s *AddModelApplicationRequestTtsConfig) GetVoiceType() *string {
	return s.VoiceType
}

func (s *AddModelApplicationRequestTtsConfig) SetBackgroundEnabled(v bool) *AddModelApplicationRequestTtsConfig {
	s.BackgroundEnabled = &v
	return s
}

func (s *AddModelApplicationRequestTtsConfig) SetBackgroundSound(v int64) *AddModelApplicationRequestTtsConfig {
	s.BackgroundSound = &v
	return s
}

func (s *AddModelApplicationRequestTtsConfig) SetBackgroundVolume(v int64) *AddModelApplicationRequestTtsConfig {
	s.BackgroundVolume = &v
	return s
}

func (s *AddModelApplicationRequestTtsConfig) SetCustomerAccountId(v int64) *AddModelApplicationRequestTtsConfig {
	s.CustomerAccountId = &v
	return s
}

func (s *AddModelApplicationRequestTtsConfig) SetMixingEnabled(v bool) *AddModelApplicationRequestTtsConfig {
	s.MixingEnabled = &v
	return s
}

func (s *AddModelApplicationRequestTtsConfig) SetMixingTemplate(v int64) *AddModelApplicationRequestTtsConfig {
	s.MixingTemplate = &v
	return s
}

func (s *AddModelApplicationRequestTtsConfig) SetResourceId(v string) *AddModelApplicationRequestTtsConfig {
	s.ResourceId = &v
	return s
}

func (s *AddModelApplicationRequestTtsConfig) SetTtsSpeed(v int64) *AddModelApplicationRequestTtsConfig {
	s.TtsSpeed = &v
	return s
}

func (s *AddModelApplicationRequestTtsConfig) SetTtsStyle(v string) *AddModelApplicationRequestTtsConfig {
	s.TtsStyle = &v
	return s
}

func (s *AddModelApplicationRequestTtsConfig) SetTtsVolume(v int64) *AddModelApplicationRequestTtsConfig {
	s.TtsVolume = &v
	return s
}

func (s *AddModelApplicationRequestTtsConfig) SetVoiceCode(v string) *AddModelApplicationRequestTtsConfig {
	s.VoiceCode = &v
	return s
}

func (s *AddModelApplicationRequestTtsConfig) SetVoiceType(v string) *AddModelApplicationRequestTtsConfig {
	s.VoiceType = &v
	return s
}

func (s *AddModelApplicationRequestTtsConfig) Validate() error {
	return dara.Validate(s)
}
