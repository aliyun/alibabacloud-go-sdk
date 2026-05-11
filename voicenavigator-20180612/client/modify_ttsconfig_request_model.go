// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTTSConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliCustomizedVoice(v string) *ModifyTTSConfigRequest
	GetAliCustomizedVoice() *string
	SetAppKey(v string) *ModifyTTSConfigRequest
	GetAppKey() *string
	SetEngine(v string) *ModifyTTSConfigRequest
	GetEngine() *string
	SetEngineXunfei(v string) *ModifyTTSConfigRequest
	GetEngineXunfei() *string
	SetInstanceId(v string) *ModifyTTSConfigRequest
	GetInstanceId() *string
	SetNlsServiceType(v string) *ModifyTTSConfigRequest
	GetNlsServiceType() *string
	SetPitchRate(v string) *ModifyTTSConfigRequest
	GetPitchRate() *string
	SetSpeechRate(v string) *ModifyTTSConfigRequest
	GetSpeechRate() *string
	SetTtsOverrides(v string) *ModifyTTSConfigRequest
	GetTtsOverrides() *string
	SetVoice(v string) *ModifyTTSConfigRequest
	GetVoice() *string
	SetVolume(v string) *ModifyTTSConfigRequest
	GetVolume() *string
}

type ModifyTTSConfigRequest struct {
	AliCustomizedVoice *string `json:"AliCustomizedVoice,omitempty" xml:"AliCustomizedVoice,omitempty"`
	AppKey             *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Engine             *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineXunfei       *string `json:"EngineXunfei,omitempty" xml:"EngineXunfei,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12f407b22cbe4890ac595f09985848d5
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NlsServiceType *string `json:"NlsServiceType,omitempty" xml:"NlsServiceType,omitempty"`
	PitchRate      *string `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// example:
	//
	// 100
	SpeechRate   *string `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	TtsOverrides *string `json:"TtsOverrides,omitempty" xml:"TtsOverrides,omitempty"`
	// example:
	//
	// aixia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 10
	Volume *string `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s ModifyTTSConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTTSConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyTTSConfigRequest) GetAliCustomizedVoice() *string {
	return s.AliCustomizedVoice
}

func (s *ModifyTTSConfigRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *ModifyTTSConfigRequest) GetEngine() *string {
	return s.Engine
}

func (s *ModifyTTSConfigRequest) GetEngineXunfei() *string {
	return s.EngineXunfei
}

func (s *ModifyTTSConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyTTSConfigRequest) GetNlsServiceType() *string {
	return s.NlsServiceType
}

func (s *ModifyTTSConfigRequest) GetPitchRate() *string {
	return s.PitchRate
}

func (s *ModifyTTSConfigRequest) GetSpeechRate() *string {
	return s.SpeechRate
}

func (s *ModifyTTSConfigRequest) GetTtsOverrides() *string {
	return s.TtsOverrides
}

func (s *ModifyTTSConfigRequest) GetVoice() *string {
	return s.Voice
}

func (s *ModifyTTSConfigRequest) GetVolume() *string {
	return s.Volume
}

func (s *ModifyTTSConfigRequest) SetAliCustomizedVoice(v string) *ModifyTTSConfigRequest {
	s.AliCustomizedVoice = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetAppKey(v string) *ModifyTTSConfigRequest {
	s.AppKey = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetEngine(v string) *ModifyTTSConfigRequest {
	s.Engine = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetEngineXunfei(v string) *ModifyTTSConfigRequest {
	s.EngineXunfei = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetInstanceId(v string) *ModifyTTSConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetNlsServiceType(v string) *ModifyTTSConfigRequest {
	s.NlsServiceType = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetPitchRate(v string) *ModifyTTSConfigRequest {
	s.PitchRate = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetSpeechRate(v string) *ModifyTTSConfigRequest {
	s.SpeechRate = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetTtsOverrides(v string) *ModifyTTSConfigRequest {
	s.TtsOverrides = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetVoice(v string) *ModifyTTSConfigRequest {
	s.Voice = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetVolume(v string) *ModifyTTSConfigRequest {
	s.Volume = &v
	return s
}

func (s *ModifyTTSConfigRequest) Validate() error {
	return dara.Validate(s)
}
