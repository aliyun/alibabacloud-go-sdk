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
	SetInstanceId(v string) *AuditTTSVoiceRequest
	GetInstanceId() *string
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
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Engine    *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PitchRate  *string `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	SecretKey  *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	SpeechRate *string `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// This parameter is required.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aixia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
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

func (s *AuditTTSVoiceRequest) GetInstanceId() *string {
	return s.InstanceId
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

func (s *AuditTTSVoiceRequest) SetInstanceId(v string) *AuditTTSVoiceRequest {
	s.InstanceId = &v
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
