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
	// The AccessKey (AK) of this namespace.
	//
	// > When the engine is xunfei, you must provide its AK here.
	//
	// example:
	//
	// 5d0f37**********ef56db601****
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// Alibaba customized voice, which is the voice invocation ID
	//
	// example:
	//
	// voice-e1be3a6
	AliCustomizedVoice *string `json:"AliCustomizedVoice,omitempty" xml:"AliCustomizedVoice,omitempty"`
	// Voice service type
	//
	// - When **ali*	- is used as the custom service, this field stores the appKey of the Intelligent Speech Interaction product project.
	//
	// - When **xunfei*	- is used as the custom service provider, this field stores the corresponding appKey.
	//
	// example:
	//
	// xusi*******RnP7
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// Storage engine. Valid values are ali, volc, and xunfei.
	//
	// - Use **ali*	- when using the default service or a custom service provided by Alibaba.
	//
	// - Use **volc*	- when using the doubao service.
	//
	// - Use **xunfei*	- when using xunfei as the service provider (only supported in small model scenarios).
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
	// Service Type
	//
	// Managed: The default Intelligent Speech Interaction service for Intelligent Outbound Calling products (public service).
	//
	// Authorized: The Intelligent Speech Interaction service purchased by the public cloud customer (customer private service), authorized via Scenario Management > Edit > Invoke Service > Custom Service.
	//
	// > When using Alibaba\\"s Intelligent Speech Interaction service—i.e., when Alibaba is specified as the custom service provider—you must set this parameter to Authorized.
	//
	// example:
	//
	// Authorized
	NlsServiceType *string `json:"NlsServiceType,omitempty" xml:"NlsServiceType,omitempty"`
	// Pitch
	//
	// An integer between [-500, 500]. Default value is 0.
	//
	// A value greater than 0 raises the pitch.
	//
	// A value less than 0 lowers the pitch.
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
	// AccessKey secret (SK) of the namespace.
	//
	// > When engine is set to xunfei, you must provide its SK here.
	//
	// example:
	//
	// OTdhNDE3Z********zQ****
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// Speech rate
	//
	// An integer between -500 and 500. Default value is 0.
	//
	// A value greater than 0 increases the speech rate.
	//
	// A value less than 0 decreases the speech rate.
	//
	// example:
	//
	// 0
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// Text content
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// Voice speaker information, such as aixia, siyue, or xiaoyun. For a complete list of available voices, see https://ai.aliyun.com/nls/tts.
	//
	// example:
	//
	// aixia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// Volume
	//
	// An integer between [0, 100]. Default value is 50.
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
