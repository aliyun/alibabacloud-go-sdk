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
	SetBackgroundMusicName(v string) *ModifyTTSConfigRequest
	GetBackgroundMusicName() *string
	SetEngine(v string) *ModifyTTSConfigRequest
	GetEngine() *string
	SetEngineXunfei(v string) *ModifyTTSConfigRequest
	GetEngineXunfei() *string
	SetExtParams(v string) *ModifyTTSConfigRequest
	GetExtParams() *string
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
	SetTtsOverridesUuid(v string) *ModifyTTSConfigRequest
	GetTtsOverridesUuid() *string
	SetVoice(v string) *ModifyTTSConfigRequest
	GetVoice() *string
	SetVolume(v string) *ModifyTTSConfigRequest
	GetVolume() *string
}

type ModifyTTSConfigRequest struct {
	// The personalized custom voice ID.
	//
	// example:
	//
	// dc458bba-5a25-4cbc-b5c2
	AliCustomizedVoice *string `json:"AliCustomizedVoice,omitempty" xml:"AliCustomizedVoice,omitempty"`
	// The AppKey of the third-party voice configuration.
	//
	// example:
	//
	// 5b358afc
	AppKey              *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	BackgroundMusicName *string `json:"BackgroundMusicName,omitempty" xml:"BackgroundMusicName,omitempty"`
	// The TTS engine.
	//
	// example:
	//
	// bailian
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The iFLYTEK engine parameters.
	//
	// example:
	//
	// {\\"Voice\\":\\"aisjinger\\"}
	EngineXunfei *string `json:"EngineXunfei,omitempty" xml:"EngineXunfei,omitempty"`
	ExtParams    *string `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	// The scenario ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8a9bdaa895a748528a15b50c281e6474
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The TTS service invoke type.
	//
	// example:
	//
	// Managed
	NlsServiceType *string `json:"NlsServiceType,omitempty" xml:"NlsServiceType,omitempty"`
	// The pitch rate.
	//
	// example:
	//
	// 50
	PitchRate *string `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// The speech rate.
	//
	// example:
	//
	// 50
	SpeechRate *string `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// The TTS error correction dictionary.
	//
	// example:
	//
	// [{\\"pronunciation\\":\\"huanqian\\",\\"word\\":\\"huaiqian\\"}]
	TtsOverrides *string `json:"TtsOverrides,omitempty" xml:"TtsOverrides,omitempty"`
	// The TTS error correction dictionary ID.
	//
	// example:
	//
	// 94001ae8-72fd-4f93-84dc-e58e2b20363b
	TtsOverridesUuid *string `json:"TtsOverridesUuid,omitempty" xml:"TtsOverridesUuid,omitempty"`
	// The voice.
	//
	// example:
	//
	// aixia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// The volume.
	//
	// example:
	//
	// 50
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

func (s *ModifyTTSConfigRequest) GetBackgroundMusicName() *string {
	return s.BackgroundMusicName
}

func (s *ModifyTTSConfigRequest) GetEngine() *string {
	return s.Engine
}

func (s *ModifyTTSConfigRequest) GetEngineXunfei() *string {
	return s.EngineXunfei
}

func (s *ModifyTTSConfigRequest) GetExtParams() *string {
	return s.ExtParams
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

func (s *ModifyTTSConfigRequest) GetTtsOverridesUuid() *string {
	return s.TtsOverridesUuid
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

func (s *ModifyTTSConfigRequest) SetBackgroundMusicName(v string) *ModifyTTSConfigRequest {
	s.BackgroundMusicName = &v
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

func (s *ModifyTTSConfigRequest) SetExtParams(v string) *ModifyTTSConfigRequest {
	s.ExtParams = &v
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

func (s *ModifyTTSConfigRequest) SetTtsOverridesUuid(v string) *ModifyTTSConfigRequest {
	s.TtsOverridesUuid = &v
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
