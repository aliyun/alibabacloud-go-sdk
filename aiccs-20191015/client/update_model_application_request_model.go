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
	SetDtmfSendMaxCount(v int64) *UpdateModelApplicationRequest
	GetDtmfSendMaxCount() *int64
	SetDtmfSendWaitTimeout(v int64) *UpdateModelApplicationRequest
	GetDtmfSendWaitTimeout() *int64
	SetDyvmsSceneName(v string) *UpdateModelApplicationRequest
	GetDyvmsSceneName() *string
	SetEnableDtmfReceive(v bool) *UpdateModelApplicationRequest
	GetEnableDtmfReceive() *bool
	SetEnableDtmfSend(v bool) *UpdateModelApplicationRequest
	GetEnableDtmfSend() *bool
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
	SetMutePushMode(v string) *UpdateModelApplicationRequest
	GetMutePushMode() *string
	SetOwnerId(v int64) *UpdateModelApplicationRequest
	GetOwnerId() *int64
	SetPauseTime(v int64) *UpdateModelApplicationRequest
	GetPauseTime() *int64
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
	// The code of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// DKSDLSA
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// The number of concurrent requests for the application.
	//
	// example:
	//
	// 12
	ApplicationCps *int64 `json:"ApplicationCps,omitempty" xml:"ApplicationCps,omitempty"`
	// The name of the model application.
	//
	// example:
	//
	// TestApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// Specifies whether to hang up when an intelligent answering service is detected.
	//
	// example:
	//
	// false
	CallAssistantHangup *bool `json:"CallAssistantHangup,omitempty" xml:"CallAssistantHangup,omitempty"`
	// Specifies whether to enable call assistant recognition.
	//
	// example:
	//
	// true
	CallAssistantRecognize *bool `json:"CallAssistantRecognize,omitempty" xml:"CallAssistantRecognize,omitempty"`
	// Specifies whether to enable call-connected event push. This feature is disabled by default.
	//
	// example:
	//
	// false
	CallConnectedTriggerModel *bool `json:"CallConnectedTriggerModel,omitempty" xml:"CallConnectedTriggerModel,omitempty"`
	// The enumeration of allowed key digits, specified as comma-separated text such as 1,2,3. A maximum of 20 values are supported.
	//
	// example:
	//
	// 1
	DtmfAllowedDigits *string `json:"DtmfAllowedDigits,omitempty" xml:"DtmfAllowedDigits,omitempty"`
	// Specifies whether to enable automatic key value validation.
	//
	// example:
	//
	// true
	DtmfAutoValidateEnable *bool `json:"DtmfAutoValidateEnable,omitempty" xml:"DtmfAutoValidateEnable,omitempty"`
	// The number of DTMF key digits. Valid values: 1 to 12.
	//
	// example:
	//
	// 1
	DtmfDigitCount *int64 `json:"DtmfDigitCount,omitempty" xml:"DtmfDigitCount,omitempty"`
	// The DTMF input timeout period in seconds. Valid values: 1 to 10.
	//
	// example:
	//
	// 1
	DtmfInputTimeout *int64 `json:"DtmfInputTimeout,omitempty" xml:"DtmfInputTimeout,omitempty"`
	// The action to take when the input is out of range. Valid values:
	//
	// - RETURN_MODEL
	//
	// - AUTO_RETRY
	//
	// example:
	//
	// RETURN_MODEL
	DtmfOutOfRangeAction *string `json:"DtmfOutOfRangeAction,omitempty" xml:"DtmfOutOfRangeAction,omitempty"`
	// The number of retry attempts (PlayTimes). Valid values: 1 to 3. This parameter takes effect only when DtmfOutOfRangeAction is set to AUTO_RETRY.
	//
	// example:
	//
	// 1
	DtmfRetryPlayTimes *int64 `json:"DtmfRetryPlayTimes,omitempty" xml:"DtmfRetryPlayTimes,omitempty"`
	// The custom retry prompt text. The maximum length is 50 characters. If this parameter is left empty, the default prompt "Invalid input. Enter again." is used.
	//
	// example:
	//
	// Sample text
	DtmfRetryPromptText *string `json:"DtmfRetryPromptText,omitempty" xml:"DtmfRetryPromptText,omitempty"`
	// The maximum number of DTMF key sending attempts.
	//
	// example:
	//
	// 90
	DtmfSendMaxCount *int64 `json:"DtmfSendMaxCount,omitempty" xml:"DtmfSendMaxCount,omitempty"`
	// The timeout period for waiting for DTMF key sending.
	//
	// example:
	//
	// 58
	DtmfSendWaitTimeout *int64 `json:"DtmfSendWaitTimeout,omitempty" xml:"DtmfSendWaitTimeout,omitempty"`
	// The name of the scenario.
	//
	// example:
	//
	// TestScenario
	DyvmsSceneName *string `json:"DyvmsSceneName,omitempty" xml:"DyvmsSceneName,omitempty"`
	// Specifies whether to enable DTMF key collection. Default value: false.
	//
	// example:
	//
	// false
	EnableDtmfReceive *bool `json:"EnableDtmfReceive,omitempty" xml:"EnableDtmfReceive,omitempty"`
	// Specifies whether to enable DTMF key sending.
	//
	// example:
	//
	// true
	EnableDtmfSend *bool `json:"EnableDtmfSend,omitempty" xml:"EnableDtmfSend,omitempty"`
	// Specifies whether to enable Morse code configuration. This feature is disabled by default.
	//
	// example:
	//
	// false
	EnableMorse *bool `json:"EnableMorse,omitempty" xml:"EnableMorse,omitempty"`
	// The interruption configuration.
	InterruptConfig *UpdateModelApplicationRequestInterruptConfig `json:"InterruptConfig,omitempty" xml:"InterruptConfig,omitempty" type:"Struct"`
	// The code of the model.
	//
	// example:
	//
	// 1231
	ModelCode *string `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	// The version of the model.
	//
	// example:
	//
	// 1
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// Specifies whether the first mute event triggers the model.
	//
	// example:
	//
	// true
	MuteActive *bool `json:"MuteActive,omitempty" xml:"MuteActive,omitempty"`
	// The mute duration.
	//
	// example:
	//
	// 85
	MuteDuration *int64 `json:"MuteDuration,omitempty" xml:"MuteDuration,omitempty"`
	// The number of consecutive mute events before the system proactively hangs up.
	//
	// example:
	//
	// 70
	MuteHangupNum *int64 `json:"MuteHangupNum,omitempty" xml:"MuteHangupNum,omitempty"`
	// The push mode for mute events.
	//
	// example:
	//
	// FIRST_ONLY
	MutePushMode *string `json:"MutePushMode,omitempty" xml:"MutePushMode,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ASR recognition pause duration.
	//
	// example:
	//
	// 350
	PauseTime *int64 `json:"PauseTime,omitempty" xml:"PauseTime,omitempty"`
	// The prompt.
	//
	// example:
	//
	// TestPrompt
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The ID of the qualification.
	//
	// example:
	//
	// 61
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// The name of the qualification.
	//
	// example:
	//
	// TestQualification
	QualificationName *string `json:"QualificationName,omitempty" xml:"QualificationName,omitempty"`
	// The URL of the recording audio file.
	//
	// example:
	//
	// https://xxxxxxxxxxxxxxx.wav
	RecordingFile        *string `json:"RecordingFile,omitempty" xml:"RecordingFile,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The maximum call duration.
	//
	// example:
	//
	// 49
	SessionTimeout *int64 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// The source. Fixed value: USER.
	//
	// example:
	//
	// USER
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The content of the speech script.
	//
	// example:
	//
	// TestSpeechScript
	SpeechContent *string `json:"SpeechContent,omitempty" xml:"SpeechContent,omitempty"`
	// The ID of the speech script.
	//
	// example:
	//
	// 15
	SpeechId *int64 `json:"SpeechId,omitempty" xml:"SpeechId,omitempty"`
	// The opening statement.
	//
	// example:
	//
	// Hello, this is a test opening statement
	StartWord *string `json:"StartWord,omitempty" xml:"StartWord,omitempty"`
	// The type of the opening statement.
	//
	// example:
	//
	// 0
	StartWordType *int64 `json:"StartWordType,omitempty" xml:"StartWordType,omitempty"`
	// The text-to-speech (TTS) configuration, including voice, volume, and speed.
	TtsConfig *UpdateModelApplicationRequestTtsConfig `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty" type:"Struct"`
	// The usage description.
	//
	// example:
	//
	// TestUsage
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

func (s *UpdateModelApplicationRequest) GetDtmfSendMaxCount() *int64 {
	return s.DtmfSendMaxCount
}

func (s *UpdateModelApplicationRequest) GetDtmfSendWaitTimeout() *int64 {
	return s.DtmfSendWaitTimeout
}

func (s *UpdateModelApplicationRequest) GetDyvmsSceneName() *string {
	return s.DyvmsSceneName
}

func (s *UpdateModelApplicationRequest) GetEnableDtmfReceive() *bool {
	return s.EnableDtmfReceive
}

func (s *UpdateModelApplicationRequest) GetEnableDtmfSend() *bool {
	return s.EnableDtmfSend
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

func (s *UpdateModelApplicationRequest) GetMutePushMode() *string {
	return s.MutePushMode
}

func (s *UpdateModelApplicationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateModelApplicationRequest) GetPauseTime() *int64 {
	return s.PauseTime
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

func (s *UpdateModelApplicationRequest) SetDtmfSendMaxCount(v int64) *UpdateModelApplicationRequest {
	s.DtmfSendMaxCount = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetDtmfSendWaitTimeout(v int64) *UpdateModelApplicationRequest {
	s.DtmfSendWaitTimeout = &v
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

func (s *UpdateModelApplicationRequest) SetEnableDtmfSend(v bool) *UpdateModelApplicationRequest {
	s.EnableDtmfSend = &v
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

func (s *UpdateModelApplicationRequest) SetMutePushMode(v string) *UpdateModelApplicationRequest {
	s.MutePushMode = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetOwnerId(v int64) *UpdateModelApplicationRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateModelApplicationRequest) SetPauseTime(v int64) *UpdateModelApplicationRequest {
	s.PauseTime = &v
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
	// The continuous interruption prevention configuration.
	AvoidInterruptDTO *UpdateModelApplicationRequestInterruptConfigAvoidInterruptDTO `json:"AvoidInterruptDTO,omitempty" xml:"AvoidInterruptDTO,omitempty" type:"Struct"`
	// Specifies whether to enable the continuous interruption prevention feature.
	//
	// example:
	//
	// true
	EnableAvoidInterrupt *bool `json:"EnableAvoidInterrupt,omitempty" xml:"EnableAvoidInterrupt,omitempty"`
	// Specifies whether to enable the interrupt filler word configuration.
	//
	// example:
	//
	// true
	EnableInterruptBackchannel *bool `json:"EnableInterruptBackchannel,omitempty" xml:"EnableInterruptBackchannel,omitempty"`
	// Specifies whether to disable interruption during the entire opening statement.
	//
	// example:
	//
	// true
	EnableStartwordEntireNotInterrupt *bool `json:"EnableStartwordEntireNotInterrupt,omitempty" xml:"EnableStartwordEntireNotInterrupt,omitempty"`
	// Specifies whether to enable the opening line no-interrupt configuration.
	//
	// example:
	//
	// true
	EnableStartwordNotInterrupt *bool `json:"EnableStartwordNotInterrupt,omitempty" xml:"EnableStartwordNotInterrupt,omitempty"`
	// The opening line protection duration.
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
	// The number of consecutive interruptions.
	//
	// example:
	//
	// 3
	InterruptNum *int64 `json:"InterruptNum,omitempty" xml:"InterruptNum,omitempty"`
	// The interrupt protection duration.
	//
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
	// Specifies whether to enable background sound.
	//
	// example:
	//
	// true
	BackgroundEnabled *bool `json:"BackgroundEnabled,omitempty" xml:"BackgroundEnabled,omitempty"`
	// The ID of the background sound.
	//
	// example:
	//
	// 111
	BackgroundSound *int64 `json:"BackgroundSound,omitempty" xml:"BackgroundSound,omitempty"`
	// The background sound volume (ID). Valid values:
	//
	// - 0: low.
	//
	// - 1: medium.
	//
	// - 2: high.
	//
	// example:
	//
	// 1
	BackgroundVolume *int64 `json:"BackgroundVolume,omitempty" xml:"BackgroundVolume,omitempty"`
	// The account ID.
	//
	// example:
	//
	// 47
	CustomerAccountId *int64 `json:"CustomerAccountId,omitempty" xml:"CustomerAccountId,omitempty"`
	// Specifies whether to enable audio mixing.
	//
	// example:
	//
	// true
	MixingEnabled *bool `json:"MixingEnabled,omitempty" xml:"MixingEnabled,omitempty"`
	// The ID of the mixing template.
	//
	// example:
	//
	// 111
	MixingTemplate *int64 `json:"MixingTemplate,omitempty" xml:"MixingTemplate,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// 122
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The speed for TTS variable playback. Valid values: -200 to 200. Default value: 0.
	//
	// example:
	//
	// 7
	TtsSpeed *int64 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	// The voice style.
	//
	// example:
	//
	// voice
	TtsStyle *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	// The volume for TTS variable playback. Valid values: 0 to 100. Default value: 0.
	//
	// example:
	//
	// 11
	TtsVolume *int64 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	// The code of the voice.
	//
	// example:
	//
	// 12123213123
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// The voice type. Valid values:
	//
	// - SYSTEM: system voice.
	//
	// - COSYCLONE: cloned voice.
	//
	// - BL-CUSTOM: custom premium cloned voice.
	//
	// example:
	//
	// SYSTEM
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
