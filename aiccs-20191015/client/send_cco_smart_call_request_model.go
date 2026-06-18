// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCcoSmartCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionCodeBreak(v bool) *SendCcoSmartCallRequest
	GetActionCodeBreak() *bool
	SetActionCodeTimeBreak(v int32) *SendCcoSmartCallRequest
	GetActionCodeTimeBreak() *int32
	SetAsrAlsAmId(v string) *SendCcoSmartCallRequest
	GetAsrAlsAmId() *string
	SetAsrBaseId(v string) *SendCcoSmartCallRequest
	GetAsrBaseId() *string
	SetAsrModelId(v string) *SendCcoSmartCallRequest
	GetAsrModelId() *string
	SetAsrVocabularyId(v string) *SendCcoSmartCallRequest
	GetAsrVocabularyId() *string
	SetBackgroundFileCode(v string) *SendCcoSmartCallRequest
	GetBackgroundFileCode() *string
	SetBackgroundSpeed(v int32) *SendCcoSmartCallRequest
	GetBackgroundSpeed() *int32
	SetBackgroundVolume(v int32) *SendCcoSmartCallRequest
	GetBackgroundVolume() *int32
	SetCalledNumber(v string) *SendCcoSmartCallRequest
	GetCalledNumber() *string
	SetCalledShowNumber(v string) *SendCcoSmartCallRequest
	GetCalledShowNumber() *string
	SetDynamicId(v string) *SendCcoSmartCallRequest
	GetDynamicId() *string
	SetEarlyMediaAsr(v bool) *SendCcoSmartCallRequest
	GetEarlyMediaAsr() *bool
	SetEnableITN(v bool) *SendCcoSmartCallRequest
	GetEnableITN() *bool
	SetMuteTime(v int32) *SendCcoSmartCallRequest
	GetMuteTime() *int32
	SetOutId(v string) *SendCcoSmartCallRequest
	GetOutId() *string
	SetOwnerId(v int64) *SendCcoSmartCallRequest
	GetOwnerId() *int64
	SetPauseTime(v int32) *SendCcoSmartCallRequest
	GetPauseTime() *int32
	SetPlayTimes(v int32) *SendCcoSmartCallRequest
	GetPlayTimes() *int32
	SetProdCode(v string) *SendCcoSmartCallRequest
	GetProdCode() *string
	SetRecordFlag(v bool) *SendCcoSmartCallRequest
	GetRecordFlag() *bool
	SetResourceOwnerAccount(v string) *SendCcoSmartCallRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SendCcoSmartCallRequest
	GetResourceOwnerId() *int64
	SetSessionTimeout(v int32) *SendCcoSmartCallRequest
	GetSessionTimeout() *int32
	SetSpeed(v int32) *SendCcoSmartCallRequest
	GetSpeed() *int32
	SetTtsConf(v bool) *SendCcoSmartCallRequest
	GetTtsConf() *bool
	SetTtsSpeed(v int32) *SendCcoSmartCallRequest
	GetTtsSpeed() *int32
	SetTtsStyle(v string) *SendCcoSmartCallRequest
	GetTtsStyle() *string
	SetTtsVolume(v int32) *SendCcoSmartCallRequest
	GetTtsVolume() *int32
	SetVoiceCode(v string) *SendCcoSmartCallRequest
	GetVoiceCode() *string
	SetVoiceCodeParam(v string) *SendCcoSmartCallRequest
	GetVoiceCodeParam() *string
	SetVolume(v int32) *SendCcoSmartCallRequest
	GetVolume() *int32
}

type SendCcoSmartCallRequest struct {
	// Indicates whether the initial playback file can be interrupted. The default value is **true**, meaning the initial playback file can be interrupted.
	//
	// example:
	//
	// true
	ActionCodeBreak *bool `json:"ActionCodeBreak,omitempty" xml:"ActionCodeBreak,omitempty"`
	// Interrupts based on the user\\"s continuous speaking duration. Takes effect only when ActionCodeBreak is set to **true**. Unit: milliseconds.
	//
	// example:
	//
	// 120
	ActionCodeTimeBreak *int32 `json:"ActionCodeTimeBreak,omitempty" xml:"ActionCodeTimeBreak,omitempty"`
	// Acoustic model ID.
	//
	// example:
	//
	// 123456
	AsrAlsAmId *string `json:"AsrAlsAmId,omitempty" xml:"AsrAlsAmId,omitempty"`
	// ASR foundation model.
	//
	// - **customer_service_8k**: Mandarin.
	//
	// - **dialect_customer_service_8k**: Strong accent.
	//
	// - If only **asrModelId*	- is set, the specified ASR model is used.
	//
	// - If only **AsrBaseId*	- is set, the specified ASR foundation model is used.
	//
	// - If neither is set, the default ASR foundation model is used. By default, **AsrBaseId*	- is **customer_service_8k**, which corresponds to the Mandarin ASR foundation model.
	//
	// - If both are set, ensure they correctly correspond to each other.
	//
	// > When invoking the **SendCcoSmartCall*	- API, you must specify the ASR model to use. We recommend that you specify either **asrModelId*	- or **AsrBaseId**, but not both.
	//
	// example:
	//
	// customer_service_8k
	AsrBaseId *string `json:"AsrBaseId,omitempty" xml:"AsrBaseId,omitempty"`
	// The ASR model ID. You can view the ASR model ID on the [ASR Model Management Page](https://aiccs.console.aliyun.com/sentence/model/private?spm=a2c4g.11186623.0.0.7f9b2964fYSGv4).
	//
	// example:
	//
	// bf71664d30d2478fb8cb8c39c6b6****
	AsrModelId *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	// Hotword ID. You can view the ASR hotword ID on the [ASR Hotword Management Page](https://aiccs.console.aliyun.com/sentence/vocab?spm=a2c4g.11186623.0.0.7f9bf965IKBpsi).
	//
	// example:
	//
	// 123456
	AsrVocabularyId *string `json:"AsrVocabularyId,omitempty" xml:"AsrVocabularyId,omitempty"`
	// ID of the background audio file played during the conversation between the user and the robot. You can log on to the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview), choose **Intelligent Interaction > Audio File Management**, and click **View Details*	- to obtain the corresponding audio ID.
	//
	// example:
	//
	// 2d4c-4e78-8d2a-afbb06cf****.wav
	BackgroundFileCode *string `json:"BackgroundFileCode,omitempty" xml:"BackgroundFileCode,omitempty"`
	// This parameter is not currently supported.
	//
	// example:
	//
	// 1
	BackgroundSpeed *int32 `json:"BackgroundSpeed,omitempty" xml:"BackgroundSpeed,omitempty"`
	// This parameter is not currently supported.
	//
	// example:
	//
	// 1
	BackgroundVolume *int32 `json:"BackgroundVolume,omitempty" xml:"BackgroundVolume,omitempty"`
	// Callee number. Only numbers from the Chinese mainland are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 137****0000
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// Called party display number. This must be a purchased number.
	//
	// You can log on to the [Contact Center console](https://aiccs.console.aliyun.com/overview?spm=a2c4g.11186623.0.0.7f9bf9658X6jte) to view your purchased numbers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0571****0000
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// A dynamic extension ID reserved for the caller. This ID is returned in the webhook address and used as a developer identity for the customer.
	//
	// example:
	//
	// 123456
	DynamicId *string `json:"DynamicId,omitempty" xml:"DynamicId,omitempty"`
	// The early media speech recognition identity. When set to true, it records the reason why the call was not answered. The default value is **false**, meaning this feature is disabled.
	//
	// > To enable early media speech recognition, you must manually set this parameter to **true**.
	//
	// example:
	//
	// false
	EarlyMediaAsr *bool `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	// Whether to execute ITN in post-processing.
	//
	// > When set to **true**, Chinese numerals are converted to Arabic numerals in the output. The default value is **false**.
	//
	// example:
	//
	// true
	EnableITN *bool `json:"EnableITN,omitempty" xml:"EnableITN,omitempty"`
	// Mute duration, used to set how long the user remains silent before the call ends. The unit is milliseconds, and the valid range is **1000–20000**.
	//
	// - If the specified value is outside this range, the default MuteTime is **10000**.
	//
	// - This parameter can be dynamically set during the call, and the last setting takes effect.
	//
	// example:
	//
	// 10000
	MuteTime *int32 `json:"MuteTime,omitempty" xml:"MuteTime,omitempty"`
	// An ID reserved for the caller, which will ultimately be returned to the caller in the receipt message.
	//
	// It is of string type and must be 1 to 15 bytes in length.
	//
	// example:
	//
	// 222356****
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The pause duration. This parameter specifies how long a user must pause to indicate the end of a sentence. Unit: milliseconds. Valid values are **300–1200**. If the specified value is outside this range, PauseTime defaults to **800**.
	//
	// > The first setting takes effect; subsequent settings are ignored.
	//
	// example:
	//
	// 800
	PauseTime *int32 `json:"PauseTime,omitempty" xml:"PauseTime,omitempty"`
	// This parameter is not currently supported.
	//
	// example:
	//
	// 1
	PlayTimes *int32 `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	// Product name. Default value: **aiccs**.
	//
	// example:
	//
	// aiccs
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// Whether to record during the call.
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	RecordFlag           *bool   `json:"RecordFlag,omitempty" xml:"RecordFlag,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The maximum call duration, in seconds. The call is automatically disconnected after timeout.
	//
	// example:
	//
	// 120
	SessionTimeout *int32 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// This parameter is currently not supported.
	//
	// example:
	//
	// 1
	Speed *int32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// Whether to configure TTS voice parameters.
	//
	// - If set to **true**, you must configure the voice style using the TtsStyle, TtsVolume, and TtsSpeed parameters.
	//
	// - If set to **false**, related parameters do not take effect, even if configured.
	//
	// example:
	//
	// true
	TtsConf *bool `json:"TtsConf,omitempty" xml:"TtsConf,omitempty"`
	// The playback speed of the TTS variable. Valid values range from -200 to 200. The default value is 0.
	//
	// example:
	//
	// 100
	TtsSpeed *int32 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	// The voice style for TTS variable playback. The default value is **xiaoyun**. For available styles, see the voice style list.
	//
	// example:
	//
	// xiaoyun
	TtsStyle *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	// Playback volume for TTS variables. Valid values: **0–100**. Default value: **0**.
	//
	// example:
	//
	// 10
	TtsVolume *int32 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	// The audio playback file for Intelligent outbound calls, which supports both network files and TTS. Multiple files and TTS parameters can be mixed and separated by commas (,). The replacement values for TTS parameters are specified in **VoiceCodeParam**.
	//
	// - When using a network file for playback: Set the VoiceCode parameter to a publicly accessible URL of the audio file. We recommend using a .wav audio file with a sampling frequency of 8000 Hz or 16000 Hz.
	//
	// - When using TTS for playback: Set the VoiceCode parameter to a variable name such as $name$, and define the corresponding content for this variable in VoiceCodeParam.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2d4c-4e78-8d2a-afbb06cf****.wav,$name$
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// A JSON-formatted string for passing TTS parameters. It must correspond to the TTS parameters of VoiceCode.
	//
	// example:
	//
	// {"name":"喂，你好"}
	VoiceCodeParam *string `json:"VoiceCodeParam,omitempty" xml:"VoiceCodeParam,omitempty"`
	// Playback volume for user audio. Valid values: –4 to 4. We recommend setting it to 1.
	//
	// example:
	//
	// 1
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s SendCcoSmartCallRequest) String() string {
	return dara.Prettify(s)
}

func (s SendCcoSmartCallRequest) GoString() string {
	return s.String()
}

func (s *SendCcoSmartCallRequest) GetActionCodeBreak() *bool {
	return s.ActionCodeBreak
}

func (s *SendCcoSmartCallRequest) GetActionCodeTimeBreak() *int32 {
	return s.ActionCodeTimeBreak
}

func (s *SendCcoSmartCallRequest) GetAsrAlsAmId() *string {
	return s.AsrAlsAmId
}

func (s *SendCcoSmartCallRequest) GetAsrBaseId() *string {
	return s.AsrBaseId
}

func (s *SendCcoSmartCallRequest) GetAsrModelId() *string {
	return s.AsrModelId
}

func (s *SendCcoSmartCallRequest) GetAsrVocabularyId() *string {
	return s.AsrVocabularyId
}

func (s *SendCcoSmartCallRequest) GetBackgroundFileCode() *string {
	return s.BackgroundFileCode
}

func (s *SendCcoSmartCallRequest) GetBackgroundSpeed() *int32 {
	return s.BackgroundSpeed
}

func (s *SendCcoSmartCallRequest) GetBackgroundVolume() *int32 {
	return s.BackgroundVolume
}

func (s *SendCcoSmartCallRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *SendCcoSmartCallRequest) GetCalledShowNumber() *string {
	return s.CalledShowNumber
}

func (s *SendCcoSmartCallRequest) GetDynamicId() *string {
	return s.DynamicId
}

func (s *SendCcoSmartCallRequest) GetEarlyMediaAsr() *bool {
	return s.EarlyMediaAsr
}

func (s *SendCcoSmartCallRequest) GetEnableITN() *bool {
	return s.EnableITN
}

func (s *SendCcoSmartCallRequest) GetMuteTime() *int32 {
	return s.MuteTime
}

func (s *SendCcoSmartCallRequest) GetOutId() *string {
	return s.OutId
}

func (s *SendCcoSmartCallRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SendCcoSmartCallRequest) GetPauseTime() *int32 {
	return s.PauseTime
}

func (s *SendCcoSmartCallRequest) GetPlayTimes() *int32 {
	return s.PlayTimes
}

func (s *SendCcoSmartCallRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *SendCcoSmartCallRequest) GetRecordFlag() *bool {
	return s.RecordFlag
}

func (s *SendCcoSmartCallRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SendCcoSmartCallRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SendCcoSmartCallRequest) GetSessionTimeout() *int32 {
	return s.SessionTimeout
}

func (s *SendCcoSmartCallRequest) GetSpeed() *int32 {
	return s.Speed
}

func (s *SendCcoSmartCallRequest) GetTtsConf() *bool {
	return s.TtsConf
}

func (s *SendCcoSmartCallRequest) GetTtsSpeed() *int32 {
	return s.TtsSpeed
}

func (s *SendCcoSmartCallRequest) GetTtsStyle() *string {
	return s.TtsStyle
}

func (s *SendCcoSmartCallRequest) GetTtsVolume() *int32 {
	return s.TtsVolume
}

func (s *SendCcoSmartCallRequest) GetVoiceCode() *string {
	return s.VoiceCode
}

func (s *SendCcoSmartCallRequest) GetVoiceCodeParam() *string {
	return s.VoiceCodeParam
}

func (s *SendCcoSmartCallRequest) GetVolume() *int32 {
	return s.Volume
}

func (s *SendCcoSmartCallRequest) SetActionCodeBreak(v bool) *SendCcoSmartCallRequest {
	s.ActionCodeBreak = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetActionCodeTimeBreak(v int32) *SendCcoSmartCallRequest {
	s.ActionCodeTimeBreak = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetAsrAlsAmId(v string) *SendCcoSmartCallRequest {
	s.AsrAlsAmId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetAsrBaseId(v string) *SendCcoSmartCallRequest {
	s.AsrBaseId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetAsrModelId(v string) *SendCcoSmartCallRequest {
	s.AsrModelId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetAsrVocabularyId(v string) *SendCcoSmartCallRequest {
	s.AsrVocabularyId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetBackgroundFileCode(v string) *SendCcoSmartCallRequest {
	s.BackgroundFileCode = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetBackgroundSpeed(v int32) *SendCcoSmartCallRequest {
	s.BackgroundSpeed = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetBackgroundVolume(v int32) *SendCcoSmartCallRequest {
	s.BackgroundVolume = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetCalledNumber(v string) *SendCcoSmartCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetCalledShowNumber(v string) *SendCcoSmartCallRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetDynamicId(v string) *SendCcoSmartCallRequest {
	s.DynamicId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetEarlyMediaAsr(v bool) *SendCcoSmartCallRequest {
	s.EarlyMediaAsr = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetEnableITN(v bool) *SendCcoSmartCallRequest {
	s.EnableITN = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetMuteTime(v int32) *SendCcoSmartCallRequest {
	s.MuteTime = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetOutId(v string) *SendCcoSmartCallRequest {
	s.OutId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetOwnerId(v int64) *SendCcoSmartCallRequest {
	s.OwnerId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetPauseTime(v int32) *SendCcoSmartCallRequest {
	s.PauseTime = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetPlayTimes(v int32) *SendCcoSmartCallRequest {
	s.PlayTimes = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetProdCode(v string) *SendCcoSmartCallRequest {
	s.ProdCode = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetRecordFlag(v bool) *SendCcoSmartCallRequest {
	s.RecordFlag = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetResourceOwnerAccount(v string) *SendCcoSmartCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetResourceOwnerId(v int64) *SendCcoSmartCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetSessionTimeout(v int32) *SendCcoSmartCallRequest {
	s.SessionTimeout = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetSpeed(v int32) *SendCcoSmartCallRequest {
	s.Speed = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetTtsConf(v bool) *SendCcoSmartCallRequest {
	s.TtsConf = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetTtsSpeed(v int32) *SendCcoSmartCallRequest {
	s.TtsSpeed = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetTtsStyle(v string) *SendCcoSmartCallRequest {
	s.TtsStyle = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetTtsVolume(v int32) *SendCcoSmartCallRequest {
	s.TtsVolume = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetVoiceCode(v string) *SendCcoSmartCallRequest {
	s.VoiceCode = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetVoiceCodeParam(v string) *SendCcoSmartCallRequest {
	s.VoiceCodeParam = &v
	return s
}

func (s *SendCcoSmartCallRequest) SetVolume(v int32) *SendCcoSmartCallRequest {
	s.Volume = &v
	return s
}

func (s *SendCcoSmartCallRequest) Validate() error {
	return dara.Validate(s)
}
