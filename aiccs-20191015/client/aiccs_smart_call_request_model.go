// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiccsSmartCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionCodeBreak(v bool) *AiccsSmartCallRequest
	GetActionCodeBreak() *bool
	SetActionCodeTimeBreak(v int32) *AiccsSmartCallRequest
	GetActionCodeTimeBreak() *int32
	SetAsrAlsAmId(v string) *AiccsSmartCallRequest
	GetAsrAlsAmId() *string
	SetAsrBaseId(v string) *AiccsSmartCallRequest
	GetAsrBaseId() *string
	SetAsrModelId(v string) *AiccsSmartCallRequest
	GetAsrModelId() *string
	SetAsrVocabularyId(v string) *AiccsSmartCallRequest
	GetAsrVocabularyId() *string
	SetBackgroundFileCode(v string) *AiccsSmartCallRequest
	GetBackgroundFileCode() *string
	SetBackgroundSpeed(v int32) *AiccsSmartCallRequest
	GetBackgroundSpeed() *int32
	SetBackgroundVolume(v int32) *AiccsSmartCallRequest
	GetBackgroundVolume() *int32
	SetCalledNumber(v string) *AiccsSmartCallRequest
	GetCalledNumber() *string
	SetCalledShowNumber(v string) *AiccsSmartCallRequest
	GetCalledShowNumber() *string
	SetDynamicId(v string) *AiccsSmartCallRequest
	GetDynamicId() *string
	SetEarlyMediaAsr(v bool) *AiccsSmartCallRequest
	GetEarlyMediaAsr() *bool
	SetEnableITN(v bool) *AiccsSmartCallRequest
	GetEnableITN() *bool
	SetMuteTime(v int32) *AiccsSmartCallRequest
	GetMuteTime() *int32
	SetOutId(v string) *AiccsSmartCallRequest
	GetOutId() *string
	SetOwnerId(v int64) *AiccsSmartCallRequest
	GetOwnerId() *int64
	SetPauseTime(v int32) *AiccsSmartCallRequest
	GetPauseTime() *int32
	SetPlayTimes(v int32) *AiccsSmartCallRequest
	GetPlayTimes() *int32
	SetProdCode(v string) *AiccsSmartCallRequest
	GetProdCode() *string
	SetRecordFlag(v bool) *AiccsSmartCallRequest
	GetRecordFlag() *bool
	SetResourceOwnerAccount(v string) *AiccsSmartCallRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AiccsSmartCallRequest
	GetResourceOwnerId() *int64
	SetSessionTimeout(v int32) *AiccsSmartCallRequest
	GetSessionTimeout() *int32
	SetSpeed(v int32) *AiccsSmartCallRequest
	GetSpeed() *int32
	SetTtsConf(v bool) *AiccsSmartCallRequest
	GetTtsConf() *bool
	SetTtsSpeed(v int32) *AiccsSmartCallRequest
	GetTtsSpeed() *int32
	SetTtsStyle(v string) *AiccsSmartCallRequest
	GetTtsStyle() *string
	SetTtsVolume(v int32) *AiccsSmartCallRequest
	GetTtsVolume() *int32
	SetVoiceCode(v string) *AiccsSmartCallRequest
	GetVoiceCode() *string
	SetVoiceCodeParam(v string) *AiccsSmartCallRequest
	GetVoiceCodeParam() *string
	SetVolume(v int32) *AiccsSmartCallRequest
	GetVolume() *int32
}

type AiccsSmartCallRequest struct {
	// Whether the initial audio playback file is interruptible. The default value is **true**, which means the initial audio playback file can be interrupted.
	//
	// example:
	//
	// true
	ActionCodeBreak *bool `json:"ActionCodeBreak,omitempty" xml:"ActionCodeBreak,omitempty"`
	// Interrupts based on the user\\"s continuous speaking duration. Takes effect only when ActionCodeBreak is **true**. Unit: milliseconds.
	//
	// example:
	//
	// 120
	ActionCodeTimeBreak *int32 `json:"ActionCodeTimeBreak,omitempty" xml:"ActionCodeTimeBreak,omitempty"`
	// Acoustic model ID.
	//
	// example:
	//
	// 23387****
	AsrAlsAmId *string `json:"AsrAlsAmId,omitempty" xml:"AsrAlsAmId,omitempty"`
	// ASR foundation model.
	//
	// - **customer_service_8k**: Mandarin.
	//
	// - **dialect_customer_service_8k**: Heavy accent.
	//
	// > - When invoking the **SendCcoSmartCall*	- API, you must specify an ASR model. We recommend that you provide either the **asrModelId*	- or **AsrBaseId*	- parameter.
	//
	// - If only **asrModelId*	- is set, the specified ASR model is used.
	//
	// - If only **AsrBaseId*	- is set, the specified ASR foundation model is used.
	//
	// - If neither parameter is set, the default ASR foundation model is used. By default, **AsrBaseId*	- is **customer_service_8k**, which corresponds to the Mandarin ASR foundation model.
	//
	// - If both parameters are set, confirm that they correctly correspond to each other.
	//
	// example:
	//
	// customer_service_8k
	AsrBaseId *string `json:"AsrBaseId,omitempty" xml:"AsrBaseId,omitempty"`
	// ASR model ID. You can view the ASR model ID on the [ASR Model Management page](https://aiccs.console.aliyun.com/sentence/model/private?spm=a2c4g.11186623.0.0.7f9b2964fYSGv4).
	//
	// example:
	//
	// bf71664d30d2478fb8cb8c39c6b6****
	AsrModelId *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	// Hotword ID. You can view the ASR hotword ID on the [ASR Hotword Management Page](https://aiccs.console.aliyun.com/sentence/vocab?spm=a2c4g.11186623.0.0.7f9bf965IKBpsi).
	//
	// example:
	//
	// 6689****
	AsrVocabularyId *string `json:"AsrVocabularyId,omitempty" xml:"AsrVocabularyId,omitempty"`
	// ID of the background audio file played during the conversation between the user and the robot.
	//
	// You can log on to the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview), choose **Intelligent Interaction > Audio File Management**, and click **View*	- to check the corresponding audio ID.
	//
	// example:
	//
	// 2d4c-4e78-8d2a-afbb06cf****.wav
	BackgroundFileCode *string `json:"BackgroundFileCode,omitempty" xml:"BackgroundFileCode,omitempty"`
	// This parameter is currently not supported.
	//
	// example:
	//
	// 1
	BackgroundSpeed *int32 `json:"BackgroundSpeed,omitempty" xml:"BackgroundSpeed,omitempty"`
	// The parameter is not supported yet.
	//
	// example:
	//
	// 1
	BackgroundVolume *int32 `json:"BackgroundVolume,omitempty" xml:"BackgroundVolume,omitempty"`
	// Called number. Only numbers in the Chinese mainland are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1862222****
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The caller ID displayed to the callee. This must be a number you have purchased.
	//
	// You can log on to the [Contact Center console](https://aiccs.console.aliyun.com/overview?spm=a2c4g.11186623.0.0.7f9bf9658X6jte) to view your purchased numbers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0571000****
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// A dynamic extension ID reserved for the caller, which is returned in the webhook address to serve as the customer\\"s developer identity.
	//
	// example:
	//
	// 2234****
	DynamicId *string `json:"DynamicId,omitempty" xml:"DynamicId,omitempty"`
	// Early media speech recognition identity. When set to **true**, it records the reason why the call was not answered. Default value: **false**, meaning disabled.
	//
	// > To enable early media speech recognition, you must manually set this parameter to **true**.
	//
	// example:
	//
	// fasle
	EarlyMediaAsr *bool `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	// Whether to execute ITN during post-processing.
	//
	// > When set to **true**, Chinese numerals are converted to Arabic numerals in the output. The default value is **false**.
	//
	// example:
	//
	// true
	EnableITN *bool `json:"EnableITN,omitempty" xml:"EnableITN,omitempty"`
	// Silence duration. This parameter defines how long the call waits for user speech before ending the call. The unit is milliseconds, and valid values range from **1000 to 20000**.
	//
	// - If the specified value is outside this range, **MuteTime*	- defaults to **10000**.
	//
	// - This parameter can be dynamically updated during the call. The last set value takes effect.
	//
	// example:
	//
	// 10000
	MuteTime *int32 `json:"MuteTime,omitempty" xml:"MuteTime,omitempty"`
	// An ID reserved for the caller. This ID will be returned to the caller in the receipt message.
	//
	// It is a string with a length of 1 to 15 bytes.
	//
	// example:
	//
	// 222356****
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Pause duration. Specifies how long the user must pause to indicate the end of a sentence. Unit: milliseconds. Valid range: **300–1200**.
	//
	// - If the specified value is outside this range, PauseTime defaults to **800**.
	//
	// - Only the first setting takes effect; subsequent settings are ignored.
	//
	// example:
	//
	// 800
	PauseTime *int32 `json:"PauseTime,omitempty" xml:"PauseTime,omitempty"`
	// The parameter is not supported yet.
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
	// example:
	//
	// true
	RecordFlag           *bool   `json:"RecordFlag,omitempty" xml:"RecordFlag,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Maximum call duration, in seconds. The call is automatically disconnected after timeout.
	//
	// example:
	//
	// 120
	SessionTimeout *int32 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// This parameter is not currently supported.
	//
	// example:
	//
	// 1
	Speed *int32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// Whether to configure TTS voice parameters.
	//
	// - If set to **true**, you must use the TtsStyle, TtsVolume, and TtsSpeed parameters to define the voice style.
	//
	// - If set to **false**, related parameters are not required and will have no effect even if configured.
	//
	// example:
	//
	// true
	TtsConf *bool `json:"TtsConf,omitempty" xml:"TtsConf,omitempty"`
	// Speech speed when playing TTS variables. Valid values range from **-200 to 200**. The default value is **0**.
	//
	// example:
	//
	// 100
	TtsSpeed *int32 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	// Voice style used during TTS variable playback. Default value: **xiaoyun**. For available styles, see the voice style list.
	//
	// example:
	//
	// xiaoyun
	TtsStyle *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	// The volume for TTS variable playback. Valid values range from **0 to 100**. The default value is **0**.
	//
	// example:
	//
	// 10
	TtsVolume *int32 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	// The Intelligent Outbound Call playback audio file supports both network files and TTS. Multiple files and TTS parameters can be mixed for playback, separated by commas (,). The replacement values for TTS parameters are specified in **VoiceCodeParam**.
	//
	// - When the playback file is a network file: Set the VoiceCode parameter to a publicly accessible URL of the audio file. We recommend using a WAV-formatted audio file with a sampling frequency of 8000 Hz or 16000 Hz.
	//
	// - When the playback file uses TTS: Set the VoiceCode parameter to a variable name such as $name$, and define the corresponding content for this variable in VoiceCodeParam.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2d4c-4e78-8d2a-afbb06cf****.wav,$name$
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// TTS parameter string in JSON format. It must correspond to the TTS parameters of VoiceCode.
	//
	// example:
	//
	// “{\\”name\\”:\\”喂，你好\\”}”
	VoiceCodeParam *string `json:"VoiceCodeParam,omitempty" xml:"VoiceCodeParam,omitempty"`
	// The volume for playing user audio. Valid values range from **-4 to 4**. We recommend setting it to **1**.
	//
	// example:
	//
	// 1
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s AiccsSmartCallRequest) String() string {
	return dara.Prettify(s)
}

func (s AiccsSmartCallRequest) GoString() string {
	return s.String()
}

func (s *AiccsSmartCallRequest) GetActionCodeBreak() *bool {
	return s.ActionCodeBreak
}

func (s *AiccsSmartCallRequest) GetActionCodeTimeBreak() *int32 {
	return s.ActionCodeTimeBreak
}

func (s *AiccsSmartCallRequest) GetAsrAlsAmId() *string {
	return s.AsrAlsAmId
}

func (s *AiccsSmartCallRequest) GetAsrBaseId() *string {
	return s.AsrBaseId
}

func (s *AiccsSmartCallRequest) GetAsrModelId() *string {
	return s.AsrModelId
}

func (s *AiccsSmartCallRequest) GetAsrVocabularyId() *string {
	return s.AsrVocabularyId
}

func (s *AiccsSmartCallRequest) GetBackgroundFileCode() *string {
	return s.BackgroundFileCode
}

func (s *AiccsSmartCallRequest) GetBackgroundSpeed() *int32 {
	return s.BackgroundSpeed
}

func (s *AiccsSmartCallRequest) GetBackgroundVolume() *int32 {
	return s.BackgroundVolume
}

func (s *AiccsSmartCallRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *AiccsSmartCallRequest) GetCalledShowNumber() *string {
	return s.CalledShowNumber
}

func (s *AiccsSmartCallRequest) GetDynamicId() *string {
	return s.DynamicId
}

func (s *AiccsSmartCallRequest) GetEarlyMediaAsr() *bool {
	return s.EarlyMediaAsr
}

func (s *AiccsSmartCallRequest) GetEnableITN() *bool {
	return s.EnableITN
}

func (s *AiccsSmartCallRequest) GetMuteTime() *int32 {
	return s.MuteTime
}

func (s *AiccsSmartCallRequest) GetOutId() *string {
	return s.OutId
}

func (s *AiccsSmartCallRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AiccsSmartCallRequest) GetPauseTime() *int32 {
	return s.PauseTime
}

func (s *AiccsSmartCallRequest) GetPlayTimes() *int32 {
	return s.PlayTimes
}

func (s *AiccsSmartCallRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *AiccsSmartCallRequest) GetRecordFlag() *bool {
	return s.RecordFlag
}

func (s *AiccsSmartCallRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AiccsSmartCallRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AiccsSmartCallRequest) GetSessionTimeout() *int32 {
	return s.SessionTimeout
}

func (s *AiccsSmartCallRequest) GetSpeed() *int32 {
	return s.Speed
}

func (s *AiccsSmartCallRequest) GetTtsConf() *bool {
	return s.TtsConf
}

func (s *AiccsSmartCallRequest) GetTtsSpeed() *int32 {
	return s.TtsSpeed
}

func (s *AiccsSmartCallRequest) GetTtsStyle() *string {
	return s.TtsStyle
}

func (s *AiccsSmartCallRequest) GetTtsVolume() *int32 {
	return s.TtsVolume
}

func (s *AiccsSmartCallRequest) GetVoiceCode() *string {
	return s.VoiceCode
}

func (s *AiccsSmartCallRequest) GetVoiceCodeParam() *string {
	return s.VoiceCodeParam
}

func (s *AiccsSmartCallRequest) GetVolume() *int32 {
	return s.Volume
}

func (s *AiccsSmartCallRequest) SetActionCodeBreak(v bool) *AiccsSmartCallRequest {
	s.ActionCodeBreak = &v
	return s
}

func (s *AiccsSmartCallRequest) SetActionCodeTimeBreak(v int32) *AiccsSmartCallRequest {
	s.ActionCodeTimeBreak = &v
	return s
}

func (s *AiccsSmartCallRequest) SetAsrAlsAmId(v string) *AiccsSmartCallRequest {
	s.AsrAlsAmId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetAsrBaseId(v string) *AiccsSmartCallRequest {
	s.AsrBaseId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetAsrModelId(v string) *AiccsSmartCallRequest {
	s.AsrModelId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetAsrVocabularyId(v string) *AiccsSmartCallRequest {
	s.AsrVocabularyId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetBackgroundFileCode(v string) *AiccsSmartCallRequest {
	s.BackgroundFileCode = &v
	return s
}

func (s *AiccsSmartCallRequest) SetBackgroundSpeed(v int32) *AiccsSmartCallRequest {
	s.BackgroundSpeed = &v
	return s
}

func (s *AiccsSmartCallRequest) SetBackgroundVolume(v int32) *AiccsSmartCallRequest {
	s.BackgroundVolume = &v
	return s
}

func (s *AiccsSmartCallRequest) SetCalledNumber(v string) *AiccsSmartCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *AiccsSmartCallRequest) SetCalledShowNumber(v string) *AiccsSmartCallRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *AiccsSmartCallRequest) SetDynamicId(v string) *AiccsSmartCallRequest {
	s.DynamicId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetEarlyMediaAsr(v bool) *AiccsSmartCallRequest {
	s.EarlyMediaAsr = &v
	return s
}

func (s *AiccsSmartCallRequest) SetEnableITN(v bool) *AiccsSmartCallRequest {
	s.EnableITN = &v
	return s
}

func (s *AiccsSmartCallRequest) SetMuteTime(v int32) *AiccsSmartCallRequest {
	s.MuteTime = &v
	return s
}

func (s *AiccsSmartCallRequest) SetOutId(v string) *AiccsSmartCallRequest {
	s.OutId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetOwnerId(v int64) *AiccsSmartCallRequest {
	s.OwnerId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetPauseTime(v int32) *AiccsSmartCallRequest {
	s.PauseTime = &v
	return s
}

func (s *AiccsSmartCallRequest) SetPlayTimes(v int32) *AiccsSmartCallRequest {
	s.PlayTimes = &v
	return s
}

func (s *AiccsSmartCallRequest) SetProdCode(v string) *AiccsSmartCallRequest {
	s.ProdCode = &v
	return s
}

func (s *AiccsSmartCallRequest) SetRecordFlag(v bool) *AiccsSmartCallRequest {
	s.RecordFlag = &v
	return s
}

func (s *AiccsSmartCallRequest) SetResourceOwnerAccount(v string) *AiccsSmartCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AiccsSmartCallRequest) SetResourceOwnerId(v int64) *AiccsSmartCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetSessionTimeout(v int32) *AiccsSmartCallRequest {
	s.SessionTimeout = &v
	return s
}

func (s *AiccsSmartCallRequest) SetSpeed(v int32) *AiccsSmartCallRequest {
	s.Speed = &v
	return s
}

func (s *AiccsSmartCallRequest) SetTtsConf(v bool) *AiccsSmartCallRequest {
	s.TtsConf = &v
	return s
}

func (s *AiccsSmartCallRequest) SetTtsSpeed(v int32) *AiccsSmartCallRequest {
	s.TtsSpeed = &v
	return s
}

func (s *AiccsSmartCallRequest) SetTtsStyle(v string) *AiccsSmartCallRequest {
	s.TtsStyle = &v
	return s
}

func (s *AiccsSmartCallRequest) SetTtsVolume(v int32) *AiccsSmartCallRequest {
	s.TtsVolume = &v
	return s
}

func (s *AiccsSmartCallRequest) SetVoiceCode(v string) *AiccsSmartCallRequest {
	s.VoiceCode = &v
	return s
}

func (s *AiccsSmartCallRequest) SetVoiceCodeParam(v string) *AiccsSmartCallRequest {
	s.VoiceCodeParam = &v
	return s
}

func (s *AiccsSmartCallRequest) SetVolume(v int32) *AiccsSmartCallRequest {
	s.Volume = &v
	return s
}

func (s *AiccsSmartCallRequest) Validate() error {
	return dara.Validate(s)
}
