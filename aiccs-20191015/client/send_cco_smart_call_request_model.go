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
	// Specifies whether the opening audio file can be interrupted. Default value: **true**, which indicates the opening audio file can be interrupted.
	//
	// example:
	//
	// true
	ActionCodeBreak *bool `json:"ActionCodeBreak,omitempty" xml:"ActionCodeBreak,omitempty"`
	// The duration threshold for interrupting based on continuous user speech. This parameter takes effect only when ActionCodeBreak is set to **true**. Unit: milliseconds.
	//
	// example:
	//
	// 120
	ActionCodeTimeBreak *int32 `json:"ActionCodeTimeBreak,omitempty" xml:"ActionCodeTimeBreak,omitempty"`
	// The acoustic model ID.
	//
	// example:
	//
	// 123456
	AsrAlsAmId *string `json:"AsrAlsAmId,omitempty" xml:"AsrAlsAmId,omitempty"`
	// The ASR foundation model.
	//
	// - **customer_service_8k**: Mandarin.
	//
	// - **dialect_customer_service_8k**: Heavy accent.
	//
	// - If only **asrModelId*	- is set, the specified ASR model is used.
	//
	// - If only **AsrBaseId*	- is set, the specified ASR foundation model is used.
	//
	// - If neither is set, the default ASR foundation model is used. The default value of **AsrBaseId*	- is **customer_service_8k**, which indicates the ASR Mandarin foundation model.
	//
	// - If both are set, make sure they correspond correctly.
	//
	// > When you call the **SendCcoSmartCall*	- operation, specify the ASR model to use. Set either **asrModelId*	- or **AsrBaseId**.
	//
	// example:
	//
	// customer_service_8k
	AsrBaseId *string `json:"AsrBaseId,omitempty" xml:"AsrBaseId,omitempty"`
	// The ASR model ID. View the ASR model ID on the [ASR Model Management page](https://aiccs.console.aliyun.com/sentence/model/private?spm=a2c4g.11186623.0.0.7f9b2964fYSGv4).
	//
	// example:
	//
	// bf71664d30d2478fb8cb8c39c6b6****
	AsrModelId *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	// The hot word ID. View the ASR hot word ID on the [ASR Hot Word Management page](https://aiccs.console.aliyun.com/sentence/vocab?spm=a2c4g.11186623.0.0.7f9bf965IKBpsi).
	//
	// example:
	//
	// 123456
	AsrVocabularyId *string `json:"AsrVocabularyId,omitempty" xml:"AsrVocabularyId,omitempty"`
	// The ID of the background audio file played during the conversation between the user and the robot. Log on to the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview), choose **Intelligent Interaction > Voice File Management**, and click **Details*	- to view the corresponding voice ID.
	//
	// example:
	//
	// 2d4c-4e78-8d2a-afbb06cf****.wav
	BackgroundFileCode *string `json:"BackgroundFileCode,omitempty" xml:"BackgroundFileCode,omitempty"`
	// This parameter is not supported.
	//
	// example:
	//
	// 1
	BackgroundSpeed *int32 `json:"BackgroundSpeed,omitempty" xml:"BackgroundSpeed,omitempty"`
	// This parameter is not supported.
	//
	// example:
	//
	// 1
	BackgroundVolume *int32 `json:"BackgroundVolume,omitempty" xml:"BackgroundVolume,omitempty"`
	// The called number. Only numbers in the Chinese mainland are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 137****0000
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The China-only caller ID displayed to the called party. The number must be a purchased number.
	//
	// Log on to the [Contact Center console](https://aiccs.console.aliyun.com/overview?spm=a2c4g.11186623.0.0.7f9bf9658X6jte) to view purchased numbers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0571****0000
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// The dynamic extension ID reserved for the caller. This ID is returned in the callback URL for the caller\\"s development identifier.
	//
	// example:
	//
	// 123456
	DynamicId *string `json:"DynamicId,omitempty" xml:"DynamicId,omitempty"`
	// The early media speech recognition flag. If set to true, the reason for unanswered calls is recorded. Default value: **false**, which indicates the feature is disabled.
	//
	// > To enable early media speech recognition, manually set this parameter to **true**.
	//
	// example:
	//
	// false
	EarlyMediaAsr *bool `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	// Specifies whether to perform Inverse Text Normalization (ITN) in post-processing.
	//
	// > When set to **true**, Chinese numerals are converted to Arabic numerals in the output. Default value: **false**.
	//
	// example:
	//
	// true
	EnableITN *bool `json:"EnableITN,omitempty" xml:"EnableITN,omitempty"`
	// The silence duration. Specifies how long the call ends after the user stops speaking. Unit: milliseconds. Valid values: **1000 to 20000**.
	//
	// - If the specified value is outside this range, the default value of **10000*	- is used.
	//
	// - This parameter can be dynamically set during the call. The last setting takes effect.
	//
	// example:
	//
	// 10000
	MuteTime *int32 `json:"MuteTime,omitempty" xml:"MuteTime,omitempty"`
	// The ID reserved for the caller. This ID is returned to the caller in the receipt message. The value is a string of 1 to 15 bytes.
	//
	// example:
	//
	// 222356****
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The pause duration. Specifies how long a user pause indicates the end of a sentence. Unit: milliseconds. Valid values: **300 to 1200**. If the specified value is outside this range, the default value of **800*	- is used.
	//
	// > Only the first setting takes effect. Subsequent settings are ignored.
	//
	// example:
	//
	// 800
	PauseTime *int32 `json:"PauseTime,omitempty" xml:"PauseTime,omitempty"`
	// This parameter is not supported.
	//
	// example:
	//
	// 1
	PlayTimes *int32 `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	// The product name. Default value: **aiccs**.
	//
	// example:
	//
	// aiccs
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// Specifies whether to record the call. Valid values:
	//
	// - **true**: Record the call.
	//
	// - **false**: Do not record the call.
	//
	// example:
	//
	// true
	RecordFlag           *bool   `json:"RecordFlag,omitempty" xml:"RecordFlag,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The maximum call duration. Unit: seconds. The call is automatically hung up after the timeout.
	//
	// example:
	//
	// 120
	SessionTimeout *int32 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// This parameter is not supported.
	//
	// example:
	//
	// 1
	Speed *int32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// Specifies whether to set TTS voice parameters. Valid values:
	//
	// - **true**: Set the voice style by using the TtsStyle, TtsVolume, and TtsSpeed parameters.
	//
	// - **false**: Do not set the related parameters. Even if they are set, they do not take effect.
	//
	// example:
	//
	// true
	TtsConf *bool `json:"TtsConf,omitempty" xml:"TtsConf,omitempty"`
	// The voice speed for TTS variable playback. Valid values: -200 to 200. Default value: 0.
	//
	// example:
	//
	// 100
	TtsSpeed *int32 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	// The voice style for TTS variable playback. Default value: **xiaoyun**. For specific styles, refer to the voice style list.
	//
	// example:
	//
	// xiaoyun
	TtsStyle *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	// The volume for TTS variable playback. Valid values: **0 to 100**. Default value: **0**.
	//
	// example:
	//
	// 10
	TtsVolume *int32 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	// The intelligent outbound call audio file. Network files and text-to-speech (TTS) are supported. Multiple files and TTS parameters can be mixed and separated by commas (,). The replacement values for TTS parameters are specified in **VoiceCodeParam**.
	//
	// - When the audio file is a network file: Set VoiceCode to a public network access audio file URL. Use a WAV format audio file with a sampling frequency of 8000 Hz or 16000 Hz.
	//
	// - When the audio file is TTS: Set VoiceCode to a variable name such as $name$, and set the corresponding content for the variable in VoiceCodeParam in **Settings**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2d4c-4e78-8d2a-afbb06cf****.wav,$name$
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// The TTS parameter string in JSON format. This parameter must correspond to the TTS parameters in VoiceCode.
	//
	// example:
	//
	// {"name":"Hello, how are you"}
	VoiceCodeParam *string `json:"VoiceCodeParam,omitempty" xml:"VoiceCodeParam,omitempty"`
	// The volume for playing user audio. Valid values: -4 to 4. Set this parameter to 1.
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
