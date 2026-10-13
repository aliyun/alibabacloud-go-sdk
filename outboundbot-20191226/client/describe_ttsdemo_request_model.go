// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTTSDemoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKey(v string) *DescribeTTSDemoRequest
	GetAccessKey() *string
	SetAliCustomizedVoice(v string) *DescribeTTSDemoRequest
	GetAliCustomizedVoice() *string
	SetAppKey(v string) *DescribeTTSDemoRequest
	GetAppKey() *string
	SetEngine(v string) *DescribeTTSDemoRequest
	GetEngine() *string
	SetExtParams(v string) *DescribeTTSDemoRequest
	GetExtParams() *string
	SetInstanceId(v string) *DescribeTTSDemoRequest
	GetInstanceId() *string
	SetNlsServiceType(v string) *DescribeTTSDemoRequest
	GetNlsServiceType() *string
	SetPitchRate(v int32) *DescribeTTSDemoRequest
	GetPitchRate() *int32
	SetScriptId(v string) *DescribeTTSDemoRequest
	GetScriptId() *string
	SetSecretKey(v string) *DescribeTTSDemoRequest
	GetSecretKey() *string
	SetSpeechRate(v int32) *DescribeTTSDemoRequest
	GetSpeechRate() *int32
	SetText(v string) *DescribeTTSDemoRequest
	GetText() *string
	SetVoice(v string) *DescribeTTSDemoRequest
	GetVoice() *string
	SetVolume(v int32) *DescribeTTSDemoRequest
	GetVolume() *int32
}

type DescribeTTSDemoRequest struct {
	// The AccessKey (AK) of the namespace.
	//
	// > When the engine is xunfei, enter the AK here.
	//
	// example:
	//
	// 5d0f37**********ef56db601****
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The Alibaba Cloud custom voice, which is the voice invocation ID.
	//
	// example:
	//
	// voice-e1be3a6
	AliCustomizedVoice *string `json:"AliCustomizedVoice,omitempty" xml:"AliCustomizedVoice,omitempty"`
	// The speech service type.
	//
	// - When you use **ali*	- as the custom service, this field stores the appKey of the interactive voice response product project.
	//
	// - When you use **xunfei*	- as the custom service provider, this field stores the corresponding appKey.
	//
	// example:
	//
	// xusi*******RnP7
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The storage engine. Valid values: ali, volc, xunfei.
	//
	// - When using traditional voices, CosyVoice, or selecting Alibaba Cloud as the custom service, set this to **ali**.
	//
	// - When using the doubao service, set this to **volc**.
	//
	// - When using xunfei as the service provider, set this to **xunfei*	- (can only be used in small model scenarios).
	//
	// - When using your own service, set this to **dock**, and set NlsServiceType to Authorized.
	//
	// - When using CV-v3-flash and Qwen voices, set this to **bailian**.
	//
	// example:
	//
	// ali
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The extended parameter. When Engine is set to **dock**, enter the configuration ID.
	//
	// example:
	//
	// {"thirdVoiceConfigUuid":"3cd82813ece049c4aff3ea3d***********"}
	ExtParams *string `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a8eccb3c-2b26-4b6d-a54f-696b953e33a6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The service type.
	//
	// Managed: the default service (public service) of the outbound robot product.
	//
	// Authorized: the interactive voice response service purchased by the public cloud customer (custom service), configured through system administration > Third-party Voice Configuration.
	//
	// > When using Alibaba Cloud interactive voice response as the custom service provider, set this to Authorized.
	//
	// example:
	//
	// Authorized
	NlsServiceType *string `json:"NlsServiceType,omitempty" xml:"NlsServiceType,omitempty"`
	// The pitch.
	//
	// example:
	//
	// 0
	PitchRate *int32 `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// The scenario ID.
	//
	// example:
	//
	// 5ab2d935-306c-478a-88bf-d08e4e25c1b7
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The AccessKey Secret (SK) of the namespace.
	//
	// > When the engine is xunfei, enter the SK here.
	//
	// example:
	//
	// OTdhNDE3Z********zQ****
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The speech rate. The value is an integer in the range of [-500, 500]. Default value: 0.
	//
	// A value greater than 0 increases the pitch.
	//
	// A value less than 0 decreases the pitch.
	//
	// example:
	//
	// 0
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// The text content.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The voice information, such as aixia, siyue, or xiaoyun. For the complete list of available voices, refer to the voice list below.
	//
	// > Because the Voice value for cloned voices is a non-fixed value generated during voice cloning, a specific Voice value cannot be provided at this stage. You can obtain the specific VoiceID from the voice cloning page by calling the ListVoiceClone operation.
	//
	// example:
	//
	// aixia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// The volume. The value is an integer in the range of [0, 100]. Default value: 50.
	//
	// A value greater than 50 increases the volume.
	//
	// A value less than 50 decreases the volume.
	//
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s DescribeTTSDemoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTTSDemoRequest) GoString() string {
	return s.String()
}

func (s *DescribeTTSDemoRequest) GetAccessKey() *string {
	return s.AccessKey
}

func (s *DescribeTTSDemoRequest) GetAliCustomizedVoice() *string {
	return s.AliCustomizedVoice
}

func (s *DescribeTTSDemoRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *DescribeTTSDemoRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeTTSDemoRequest) GetExtParams() *string {
	return s.ExtParams
}

func (s *DescribeTTSDemoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTTSDemoRequest) GetNlsServiceType() *string {
	return s.NlsServiceType
}

func (s *DescribeTTSDemoRequest) GetPitchRate() *int32 {
	return s.PitchRate
}

func (s *DescribeTTSDemoRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DescribeTTSDemoRequest) GetSecretKey() *string {
	return s.SecretKey
}

func (s *DescribeTTSDemoRequest) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *DescribeTTSDemoRequest) GetText() *string {
	return s.Text
}

func (s *DescribeTTSDemoRequest) GetVoice() *string {
	return s.Voice
}

func (s *DescribeTTSDemoRequest) GetVolume() *int32 {
	return s.Volume
}

func (s *DescribeTTSDemoRequest) SetAccessKey(v string) *DescribeTTSDemoRequest {
	s.AccessKey = &v
	return s
}

func (s *DescribeTTSDemoRequest) SetAliCustomizedVoice(v string) *DescribeTTSDemoRequest {
	s.AliCustomizedVoice = &v
	return s
}

func (s *DescribeTTSDemoRequest) SetAppKey(v string) *DescribeTTSDemoRequest {
	s.AppKey = &v
	return s
}

func (s *DescribeTTSDemoRequest) SetEngine(v string) *DescribeTTSDemoRequest {
	s.Engine = &v
	return s
}

func (s *DescribeTTSDemoRequest) SetExtParams(v string) *DescribeTTSDemoRequest {
	s.ExtParams = &v
	return s
}

func (s *DescribeTTSDemoRequest) SetInstanceId(v string) *DescribeTTSDemoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTTSDemoRequest) SetNlsServiceType(v string) *DescribeTTSDemoRequest {
	s.NlsServiceType = &v
	return s
}

func (s *DescribeTTSDemoRequest) SetPitchRate(v int32) *DescribeTTSDemoRequest {
	s.PitchRate = &v
	return s
}

func (s *DescribeTTSDemoRequest) SetScriptId(v string) *DescribeTTSDemoRequest {
	s.ScriptId = &v
	return s
}

func (s *DescribeTTSDemoRequest) SetSecretKey(v string) *DescribeTTSDemoRequest {
	s.SecretKey = &v
	return s
}

func (s *DescribeTTSDemoRequest) SetSpeechRate(v int32) *DescribeTTSDemoRequest {
	s.SpeechRate = &v
	return s
}

func (s *DescribeTTSDemoRequest) SetText(v string) *DescribeTTSDemoRequest {
	s.Text = &v
	return s
}

func (s *DescribeTTSDemoRequest) SetVoice(v string) *DescribeTTSDemoRequest {
	s.Voice = &v
	return s
}

func (s *DescribeTTSDemoRequest) SetVolume(v int32) *DescribeTTSDemoRequest {
	s.Volume = &v
	return s
}

func (s *DescribeTTSDemoRequest) Validate() error {
	return dara.Validate(s)
}
