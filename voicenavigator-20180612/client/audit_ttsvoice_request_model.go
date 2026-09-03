// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditTTSVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKey(v string) *AuditTTSVoiceRequest
	GetAccessKey() *string
	SetAppKey(v string) *AuditTTSVoiceRequest
	GetAppKey() *string
	SetEngine(v string) *AuditTTSVoiceRequest
	GetEngine() *string
	SetExtParams(v string) *AuditTTSVoiceRequest
	GetExtParams() *string
	SetInstanceId(v string) *AuditTTSVoiceRequest
	GetInstanceId() *string
	SetNlsServiceType(v string) *AuditTTSVoiceRequest
	GetNlsServiceType() *string
	SetPitchRate(v string) *AuditTTSVoiceRequest
	GetPitchRate() *string
	SetSecretKey(v string) *AuditTTSVoiceRequest
	GetSecretKey() *string
	SetSpeechRate(v string) *AuditTTSVoiceRequest
	GetSpeechRate() *string
	SetText(v string) *AuditTTSVoiceRequest
	GetText() *string
	SetVoice(v string) *AuditTTSVoiceRequest
	GetVoice() *string
	SetVolume(v string) *AuditTTSVoiceRequest
	GetVolume() *string
}

type AuditTTSVoiceRequest struct {
	// The AccessKey ID of the namespace.
	//
	// example:
	//
	// b4331******a4640ce1f88e27ac8df0
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The AppKey of the third-party voice configuration.
	//
	// example:
	//
	// be******
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The TTS engine.
	//
	// example:
	//
	// ali
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The extended parameters.
	//
	// example:
	//
	// {}
	ExtParams *string `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	// The navigation instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NlsServiceType *string `json:"NlsServiceType,omitempty" xml:"NlsServiceType,omitempty"`
	// example:
	//
	// 0
	PitchRate *string `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// The AccessKey secret.
	//
	// example:
	//
	// ZDc3********DAzM2E0YjM5NTFkMDQ1
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The speech rate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	SpeechRate *string `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// The audition text.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hello
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The voice.
	//
	// This parameter is required.
	//
	// example:
	//
	// aixia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// The volume.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Volume *string `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s AuditTTSVoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s AuditTTSVoiceRequest) GoString() string {
	return s.String()
}

func (s *AuditTTSVoiceRequest) GetAccessKey() *string {
	return s.AccessKey
}

func (s *AuditTTSVoiceRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *AuditTTSVoiceRequest) GetEngine() *string {
	return s.Engine
}

func (s *AuditTTSVoiceRequest) GetExtParams() *string {
	return s.ExtParams
}

func (s *AuditTTSVoiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AuditTTSVoiceRequest) GetNlsServiceType() *string {
	return s.NlsServiceType
}

func (s *AuditTTSVoiceRequest) GetPitchRate() *string {
	return s.PitchRate
}

func (s *AuditTTSVoiceRequest) GetSecretKey() *string {
	return s.SecretKey
}

func (s *AuditTTSVoiceRequest) GetSpeechRate() *string {
	return s.SpeechRate
}

func (s *AuditTTSVoiceRequest) GetText() *string {
	return s.Text
}

func (s *AuditTTSVoiceRequest) GetVoice() *string {
	return s.Voice
}

func (s *AuditTTSVoiceRequest) GetVolume() *string {
	return s.Volume
}

func (s *AuditTTSVoiceRequest) SetAccessKey(v string) *AuditTTSVoiceRequest {
	s.AccessKey = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetAppKey(v string) *AuditTTSVoiceRequest {
	s.AppKey = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetEngine(v string) *AuditTTSVoiceRequest {
	s.Engine = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetExtParams(v string) *AuditTTSVoiceRequest {
	s.ExtParams = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetInstanceId(v string) *AuditTTSVoiceRequest {
	s.InstanceId = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetNlsServiceType(v string) *AuditTTSVoiceRequest {
	s.NlsServiceType = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetPitchRate(v string) *AuditTTSVoiceRequest {
	s.PitchRate = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetSecretKey(v string) *AuditTTSVoiceRequest {
	s.SecretKey = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetSpeechRate(v string) *AuditTTSVoiceRequest {
	s.SpeechRate = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetText(v string) *AuditTTSVoiceRequest {
	s.Text = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetVoice(v string) *AuditTTSVoiceRequest {
	s.Voice = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetVolume(v string) *AuditTTSVoiceRequest {
	s.Volume = &v
	return s
}

func (s *AuditTTSVoiceRequest) Validate() error {
	return dara.Validate(s)
}
