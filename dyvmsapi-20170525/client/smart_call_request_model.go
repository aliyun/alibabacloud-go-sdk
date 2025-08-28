// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionCodeBreak(v bool) *SmartCallRequest
	GetActionCodeBreak() *bool
	SetActionCodeTimeBreak(v int32) *SmartCallRequest
	GetActionCodeTimeBreak() *int32
	SetAsrBaseId(v string) *SmartCallRequest
	GetAsrBaseId() *string
	SetAsrModelId(v string) *SmartCallRequest
	GetAsrModelId() *string
	SetBackgroundFileCode(v string) *SmartCallRequest
	GetBackgroundFileCode() *string
	SetBackgroundSpeed(v int32) *SmartCallRequest
	GetBackgroundSpeed() *int32
	SetBackgroundVolume(v int32) *SmartCallRequest
	GetBackgroundVolume() *int32
	SetCalledNumber(v string) *SmartCallRequest
	GetCalledNumber() *string
	SetCalledShowNumber(v string) *SmartCallRequest
	GetCalledShowNumber() *string
	SetDynamicId(v string) *SmartCallRequest
	GetDynamicId() *string
	SetEarlyMediaAsr(v bool) *SmartCallRequest
	GetEarlyMediaAsr() *bool
	SetEnableITN(v bool) *SmartCallRequest
	GetEnableITN() *bool
	SetMuteTime(v int32) *SmartCallRequest
	GetMuteTime() *int32
	SetNoiseThreshold(v float64) *SmartCallRequest
	GetNoiseThreshold() *float64
	SetOutId(v string) *SmartCallRequest
	GetOutId() *string
	SetOwnerId(v int64) *SmartCallRequest
	GetOwnerId() *int64
	SetPauseTime(v int32) *SmartCallRequest
	GetPauseTime() *int32
	SetRecordFlag(v bool) *SmartCallRequest
	GetRecordFlag() *bool
	SetResourceOwnerAccount(v string) *SmartCallRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SmartCallRequest
	GetResourceOwnerId() *int64
	SetSessionTimeout(v int32) *SmartCallRequest
	GetSessionTimeout() *int32
	SetSpeed(v int32) *SmartCallRequest
	GetSpeed() *int32
	SetStreamAsr(v int32) *SmartCallRequest
	GetStreamAsr() *int32
	SetTtsConf(v bool) *SmartCallRequest
	GetTtsConf() *bool
	SetTtsSpeed(v int32) *SmartCallRequest
	GetTtsSpeed() *int32
	SetTtsStyle(v string) *SmartCallRequest
	GetTtsStyle() *string
	SetTtsVolume(v int32) *SmartCallRequest
	GetTtsVolume() *int32
	SetVoiceCode(v string) *SmartCallRequest
	GetVoiceCode() *string
	SetVoiceCodeParam(v string) *SmartCallRequest
	GetVoiceCodeParam() *string
	SetVolume(v int32) *SmartCallRequest
	GetVolume() *int32
}

type SmartCallRequest struct {
	// Specifies whether the playback of the recording file can be interrupted. Default value: **true**. The default value indicates that the playback of the recording file can be interrupted.
	//
	// If you set the value of this parameter to false, the playback of the recording file cannot be interrupted even if the value of action_break is set to true.
	//
	// > The value of action_code_break takes precedence over the value of action_break.
	//
	// example:
	//
	// true
	ActionCodeBreak *bool `json:"ActionCodeBreak,omitempty" xml:"ActionCodeBreak,omitempty"`
	// The duration that the user keeps speaking. The playback of the recording file is interrupted when this duration is reached. Unit: milliseconds.
	//
	// If the value of ActionCodeBreak is set to **true*	- for the recording file and the duration that the user keeps speaking reaches the specified duration, the playback of the recording file is interrupted. If you do not specify ActionCodeTimeBreak or set the value of ActionCodeTimeBreak to 0, the setting of ActionCodeBreak does not take effect.
	//
	// example:
	//
	// 120
	ActionCodeTimeBreak *int32 `json:"ActionCodeTimeBreak,omitempty" xml:"ActionCodeTimeBreak,omitempty"`
	// The ASR base model. Valid values:
	//
	// 	- **customer_service_8k*	- (default): Chinese Mandarin.
	//
	// 	- **dialect_customer_service_8k**: a heavy accent.
	//
	// > You must specify the ASR model when you call the SmartCall operation. We recommend that you specify either of the AsrModelId and AsrBaseId parameters.
	//
	// 	- If you specify only the AsrModelId parameter, the specified ASR model is used.
	//
	// 	- If you specify only the AsrBaseId parameter, the ASR base model is used.
	//
	// 	- If you specify neither of the two parameters, the default ASR base model is used, that is, the default value customer_service_8k is used for the AsrBaseId parameter.
	//
	// 	- If you specify both parameters, make sure that their values do not conflict with each other.
	//
	// example:
	//
	// customer_service_8k
	AsrBaseId *string `json:"AsrBaseId,omitempty" xml:"AsrBaseId,omitempty"`
	// The ID of the Automatic Speech Recognition (ASR) model.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home) and view the ID of the ASR model on the **ASR Model Management*	- page.
	//
	// > You must specify the ASR model when you call the SmartCall operation. We recommend that you specify either of the AsrModelId and AsrBaseId parameters.
	//
	// 	- If you specify only the AsrModelId parameter, the specified ASR model is used.
	//
	// 	- If you specify only the AsrBaseId parameter, the specified ASR base model is used.
	//
	// 	- If you specify neither of the two parameters, the default value customer_service_8k is used for the AsrBaseId parameter. This means that the default Mandarin ASR base model is used.
	//
	// 	- If you specify both parameters, make sure that their values do not conflict with each other.
	//
	// example:
	//
	// 2070aca1eff146f9a7bc826f1c3d****
	AsrModelId *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	// The ID of the background voice file that is played back when the user talks with the robot.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice File Management**, click the **Intelligent Speech Interaction Recording File*	- tab, and then click **Details*	- in the Actions column to view the voice ID.
	//
	// example:
	//
	// 2d4c-4e78-8d2a-afbb06cf****.wav
	BackgroundFileCode *string `json:"BackgroundFileCode,omitempty" xml:"BackgroundFileCode,omitempty"`
	// This parameter is unavailable.
	//
	// example:
	//
	// 1
	BackgroundSpeed *int32 `json:"BackgroundSpeed,omitempty" xml:"BackgroundSpeed,omitempty"`
	// This parameter is unavailable.
	//
	// example:
	//
	// 1
	BackgroundVolume *int32 `json:"BackgroundVolume,omitempty" xml:"BackgroundVolume,omitempty"`
	// The called number. Only phone numbers in the Chinese mainland are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1590****0000
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The number displayed to the called party. The value must be the number you purchased.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home) and choose **Voice Numbers*	- > **Real Number Management*	- to view the number you purchased.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0571****5678
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// The dynamic extension ID that is reserved for the caller of the operation. This ID is returned in the callback URL and is used as the development identifier of the customer.
	//
	// example:
	//
	// abcdefgh
	DynamicId *string `json:"DynamicId,omitempty" xml:"DynamicId,omitempty"`
	// Specifies whether to enable speech recognition of early media. Valid values:
	//
	// 	- **false*	- (default): Speech recognition of early media is disabled.
	//
	// 	- **true**: Speech recognition of early media is enabled.
	//
	// > If you set the value of this parameter to **true**, the reason why the call is not answered is recorded.
	//
	// example:
	//
	// true
	EarlyMediaAsr *bool `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	// Specifies whether to enable Inverse Text Normalization (ITN) during post-processing. Default value: **false**. If you set the value to false, ITN is not enabled during post-processing.
	//
	// If you set the value to **true**, Chinese numerals are converted into Arabic numerals for output.
	//
	// example:
	//
	// false
	EnableITN *bool `json:"EnableITN,omitempty" xml:"EnableITN,omitempty"`
	// The silence duration. The system determines the end of the conversation based on the silence duration of the user. Unit: milliseconds. Valid values: 1000 to 20000.****
	//
	// >
	//
	// 	- If you specify a value out of the valid range, the default value **10000*	- is used.
	//
	// 	- The parameter value can be adjusted during the conversation. The last setting prevails.
	//
	// example:
	//
	// 10000
	MuteTime       *int32   `json:"MuteTime,omitempty" xml:"MuteTime,omitempty"`
	NoiseThreshold *float64 `json:"NoiseThreshold,omitempty" xml:"NoiseThreshold,omitempty"`
	// The ID that is reserved for the caller of the operation. This ID is returned to the caller in a receipt message.
	//
	// The value is of the STRING type and must be 1 to 15 bytes in length.
	//
	// example:
	//
	// 342268*****
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The pause duration. The system determines the end of a sentence based on the pause duration of the user in the conversation. Unit: milliseconds. Valid values: 300 to 1200.****
	//
	// >
	//
	// 	- If you specify a value out of the valid range, the default value **800*	- is used.
	//
	// 	- You cannot change the parameter value after specifying it.
	//
	// example:
	//
	// 800
	PauseTime *int32 `json:"PauseTime,omitempty" xml:"PauseTime,omitempty"`
	// Specifies whether to record the conversation. Valid values:
	//
	// 	- **true**: The conversation is recorded.
	//
	// 	- **false**: The conversation is not recorded.
	//
	// example:
	//
	// true
	RecordFlag           *bool   `json:"RecordFlag,omitempty" xml:"RecordFlag,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The maximum call duration. The call is automatically hung up when the maximum call duration is reached. Unit: seconds.
	//
	// > The maximum call duration is 3,600 seconds.
	//
	// example:
	//
	// 120
	SessionTimeout *int32 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// This parameter is unavailable.
	//
	// example:
	//
	// 1
	Speed *int32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// Specifies whether to enable streaming ASR, which intelligently judges what the user wants to express based on the first few words spoken by the user. Valid values:
	//
	// 	- **0**: Streaming ASR is disabled.
	//
	// 	- **1**: Streaming ASR is enabled.
	//
	// example:
	//
	// 1
	StreamAsr *int32 `json:"StreamAsr,omitempty" xml:"StreamAsr,omitempty"`
	// Specifies whether to set TTS sound parameters. Valid values:
	//
	// 	- **true**: TTS sound parameters must be set. You must set the **TtsStyle**, **TtsColume**, and **TtsSpeed*	- parameters to specify a sound style.
	//
	// 	- **false**: TTS sound parameters do not need to be set. The values of TTS sound parameters do not take effect even if you set them.
	//
	// example:
	//
	// true
	TtsConf *bool `json:"TtsConf,omitempty" xml:"TtsConf,omitempty"`
	// The speed of TTS variable playback. Valid values: -200 to 200. Default value: 0.
	//
	// example:
	//
	// 100
	TtsSpeed *int32 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	// The sound style for TTS variable playback. Default value: **xiaoyun**. For more information about the sound styles, see the **Sound styles*	- table below.
	//
	// example:
	//
	// xiaoyun
	TtsStyle *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	// The volume of TTS variable playback. Valid values: 0 to 100. Default value: 0.
	//
	// example:
	//
	// 10
	TtsVolume *int32 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	// The recording file that is played back in the intelligent outbound call.
	//
	// The file can be an online file, a voice file uploaded from the Voice Messaging Service console, or a text-to-speech (TTS) template that contains variables. You can specify multiple files and a TTS variable together. Separate them with commas (,). The values of the variables in the TTS template are specified by the **VoiceCodeParam*	- parameter.
	//
	// 	- If you use an online file as the recording file, set the value of **VoiceCode*	- to the URL of the file that can be accessed over the Internet.
	//
	// 	- If you use a voice file uploaded from the Voice Messaging Service console as the recording file, set the value of **VoiceCode*	- to the voice ID of the file. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice File Management**, click the **Intelligent Speech Interaction Recording File*	- tab, and then click **Details*	- in the Actions column to view the voice ID.
	//
	// 	- If you use a TTS template that contains variables as the recording file, set the value of **VoiceCode*	- to a variable name such as $name$, and also set a value for the variable in the **VoiceCodeParam*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2d4c-4e78-8d2a-afbb06cf****.wav,$name$
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// The value of the TTS variable, in the JSON format. This value must match the TTS variable specified by the **VoiceCode*	- parameter.
	//
	// example:
	//
	// {"name":"hello"}
	VoiceCodeParam *string `json:"VoiceCodeParam,omitempty" xml:"VoiceCodeParam,omitempty"`
	// The volume at which the recording file is played. Valid values: -4 to 4. We recommend that you set the value of this parameter to **1**.
	//
	// example:
	//
	// 1
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s SmartCallRequest) String() string {
	return dara.Prettify(s)
}

func (s SmartCallRequest) GoString() string {
	return s.String()
}

func (s *SmartCallRequest) GetActionCodeBreak() *bool {
	return s.ActionCodeBreak
}

func (s *SmartCallRequest) GetActionCodeTimeBreak() *int32 {
	return s.ActionCodeTimeBreak
}

func (s *SmartCallRequest) GetAsrBaseId() *string {
	return s.AsrBaseId
}

func (s *SmartCallRequest) GetAsrModelId() *string {
	return s.AsrModelId
}

func (s *SmartCallRequest) GetBackgroundFileCode() *string {
	return s.BackgroundFileCode
}

func (s *SmartCallRequest) GetBackgroundSpeed() *int32 {
	return s.BackgroundSpeed
}

func (s *SmartCallRequest) GetBackgroundVolume() *int32 {
	return s.BackgroundVolume
}

func (s *SmartCallRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *SmartCallRequest) GetCalledShowNumber() *string {
	return s.CalledShowNumber
}

func (s *SmartCallRequest) GetDynamicId() *string {
	return s.DynamicId
}

func (s *SmartCallRequest) GetEarlyMediaAsr() *bool {
	return s.EarlyMediaAsr
}

func (s *SmartCallRequest) GetEnableITN() *bool {
	return s.EnableITN
}

func (s *SmartCallRequest) GetMuteTime() *int32 {
	return s.MuteTime
}

func (s *SmartCallRequest) GetNoiseThreshold() *float64 {
	return s.NoiseThreshold
}

func (s *SmartCallRequest) GetOutId() *string {
	return s.OutId
}

func (s *SmartCallRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SmartCallRequest) GetPauseTime() *int32 {
	return s.PauseTime
}

func (s *SmartCallRequest) GetRecordFlag() *bool {
	return s.RecordFlag
}

func (s *SmartCallRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SmartCallRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SmartCallRequest) GetSessionTimeout() *int32 {
	return s.SessionTimeout
}

func (s *SmartCallRequest) GetSpeed() *int32 {
	return s.Speed
}

func (s *SmartCallRequest) GetStreamAsr() *int32 {
	return s.StreamAsr
}

func (s *SmartCallRequest) GetTtsConf() *bool {
	return s.TtsConf
}

func (s *SmartCallRequest) GetTtsSpeed() *int32 {
	return s.TtsSpeed
}

func (s *SmartCallRequest) GetTtsStyle() *string {
	return s.TtsStyle
}

func (s *SmartCallRequest) GetTtsVolume() *int32 {
	return s.TtsVolume
}

func (s *SmartCallRequest) GetVoiceCode() *string {
	return s.VoiceCode
}

func (s *SmartCallRequest) GetVoiceCodeParam() *string {
	return s.VoiceCodeParam
}

func (s *SmartCallRequest) GetVolume() *int32 {
	return s.Volume
}

func (s *SmartCallRequest) SetActionCodeBreak(v bool) *SmartCallRequest {
	s.ActionCodeBreak = &v
	return s
}

func (s *SmartCallRequest) SetActionCodeTimeBreak(v int32) *SmartCallRequest {
	s.ActionCodeTimeBreak = &v
	return s
}

func (s *SmartCallRequest) SetAsrBaseId(v string) *SmartCallRequest {
	s.AsrBaseId = &v
	return s
}

func (s *SmartCallRequest) SetAsrModelId(v string) *SmartCallRequest {
	s.AsrModelId = &v
	return s
}

func (s *SmartCallRequest) SetBackgroundFileCode(v string) *SmartCallRequest {
	s.BackgroundFileCode = &v
	return s
}

func (s *SmartCallRequest) SetBackgroundSpeed(v int32) *SmartCallRequest {
	s.BackgroundSpeed = &v
	return s
}

func (s *SmartCallRequest) SetBackgroundVolume(v int32) *SmartCallRequest {
	s.BackgroundVolume = &v
	return s
}

func (s *SmartCallRequest) SetCalledNumber(v string) *SmartCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *SmartCallRequest) SetCalledShowNumber(v string) *SmartCallRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *SmartCallRequest) SetDynamicId(v string) *SmartCallRequest {
	s.DynamicId = &v
	return s
}

func (s *SmartCallRequest) SetEarlyMediaAsr(v bool) *SmartCallRequest {
	s.EarlyMediaAsr = &v
	return s
}

func (s *SmartCallRequest) SetEnableITN(v bool) *SmartCallRequest {
	s.EnableITN = &v
	return s
}

func (s *SmartCallRequest) SetMuteTime(v int32) *SmartCallRequest {
	s.MuteTime = &v
	return s
}

func (s *SmartCallRequest) SetNoiseThreshold(v float64) *SmartCallRequest {
	s.NoiseThreshold = &v
	return s
}

func (s *SmartCallRequest) SetOutId(v string) *SmartCallRequest {
	s.OutId = &v
	return s
}

func (s *SmartCallRequest) SetOwnerId(v int64) *SmartCallRequest {
	s.OwnerId = &v
	return s
}

func (s *SmartCallRequest) SetPauseTime(v int32) *SmartCallRequest {
	s.PauseTime = &v
	return s
}

func (s *SmartCallRequest) SetRecordFlag(v bool) *SmartCallRequest {
	s.RecordFlag = &v
	return s
}

func (s *SmartCallRequest) SetResourceOwnerAccount(v string) *SmartCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SmartCallRequest) SetResourceOwnerId(v int64) *SmartCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SmartCallRequest) SetSessionTimeout(v int32) *SmartCallRequest {
	s.SessionTimeout = &v
	return s
}

func (s *SmartCallRequest) SetSpeed(v int32) *SmartCallRequest {
	s.Speed = &v
	return s
}

func (s *SmartCallRequest) SetStreamAsr(v int32) *SmartCallRequest {
	s.StreamAsr = &v
	return s
}

func (s *SmartCallRequest) SetTtsConf(v bool) *SmartCallRequest {
	s.TtsConf = &v
	return s
}

func (s *SmartCallRequest) SetTtsSpeed(v int32) *SmartCallRequest {
	s.TtsSpeed = &v
	return s
}

func (s *SmartCallRequest) SetTtsStyle(v string) *SmartCallRequest {
	s.TtsStyle = &v
	return s
}

func (s *SmartCallRequest) SetTtsVolume(v int32) *SmartCallRequest {
	s.TtsVolume = &v
	return s
}

func (s *SmartCallRequest) SetVoiceCode(v string) *SmartCallRequest {
	s.VoiceCode = &v
	return s
}

func (s *SmartCallRequest) SetVoiceCodeParam(v string) *SmartCallRequest {
	s.VoiceCodeParam = &v
	return s
}

func (s *SmartCallRequest) SetVolume(v int32) *SmartCallRequest {
	s.Volume = &v
	return s
}

func (s *SmartCallRequest) Validate() error {
	return dara.Validate(s)
}
