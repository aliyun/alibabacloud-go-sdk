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
	// The AccessKey (AK) for this namespace.
	//
	// > Enter the AK when the engine is xunfei.
	//
	// example:
	//
	// 5d0f37**********ef56db601****
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// Alibaba Cloud custom voice ID
	//
	// example:
	//
	// voice-e1be3a6
	AliCustomizedVoice *string `json:"AliCustomizedVoice,omitempty" xml:"AliCustomizedVoice,omitempty"`
	// Speech service type
	//
	// - When using **ali*	- as a custom service, enter the appKey of your Intelligent Speech Interaction project.
	//
	// - When using **xunfei*	- as a custom service, enter its appKey.
	//
	// example:
	//
	// xusi*******RnP7
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// Storage engine. Choose from ali, volc, or xunfei.
	//
	// - Enter **ali*	- when using the default service or Alibaba Cloud as a custom service.
	//
	// - Enter **volc*	- when using the doubao service.
	//
	// - Enter **xunfei*	- when using iFLYTEK as a service provider. This option is only available for small-model scenarios.
	//
	// example:
	//
	// ali
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// a8eccb3c-2b26-4b6d-a54f-696b953e33a6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Service type
	//
	// Managed: The default Intelligent Speech Interaction service for Intelligent Outbound Calling (public service).
	//
	// Authorized: An Intelligent Speech Interaction service you purchased on Alibaba Cloud public cloud (your private service). You can grant authorization by going to Scenario Management > Edit > Call Service > Custom Service.
	//
	// > Set this parameter to Authorized when using Alibaba Cloud\\"s Intelligent Speech Interaction as your custom service provider.
	//
	// example:
	//
	// Authorized
	NlsServiceType *string `json:"NlsServiceType,omitempty" xml:"NlsServiceType,omitempty"`
	// Pitch. An integer between -500 and 500. Default is 0.
	//
	// A value greater than 0 raises pitch.
	//
	// A value less than 0 lowers pitch.
	//
	// example:
	//
	// 0
	PitchRate *int32 `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// Scenario ID
	//
	// example:
	//
	// 5ab2d935-306c-478a-88bf-d08e4e25c1b7
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The AccessKey secret (SK) for this namespace.
	//
	// > Enter the SK when the engine is xunfei.
	//
	// example:
	//
	// OTdhNDE3Z********zQ****
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// Speech rate. An integer between -500 and 500. Default is 0.
	//
	// A value greater than 0 increases speech speed.
	//
	// A value less than 0 decreases speech speed.
	//
	// example:
	//
	// 0
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// Text to convert to speech
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// Voice ID. Examples include aixia, siyue, and xiaoyun. For the full list of available voices, see the voice list below.
	//
	// > Cloned voices use dynamic Voice IDs that are generated during voice cloning. Therefore, specific Voice IDs for cloned voices are not listed here. To get a cloned voice’s Voice ID, call ListVoiceClone from the voice cloning page.
	//
	// example:
	//
	// aixia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// Volume. An integer between 0 and 100. Default is 50.
	//
	// A value greater than 50 increases volume.
	//
	// A value less than 50 decreases volume.
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
