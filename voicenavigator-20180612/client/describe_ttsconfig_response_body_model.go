// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTTSConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliCustomizedVoice(v string) *DescribeTTSConfigResponseBody
	GetAliCustomizedVoice() *string
	SetAppKey(v string) *DescribeTTSConfigResponseBody
	GetAppKey() *string
	SetEngine(v string) *DescribeTTSConfigResponseBody
	GetEngine() *string
	SetEngineXunfei(v string) *DescribeTTSConfigResponseBody
	GetEngineXunfei() *string
	SetNlsServiceType(v string) *DescribeTTSConfigResponseBody
	GetNlsServiceType() *string
	SetPitchRate(v int32) *DescribeTTSConfigResponseBody
	GetPitchRate() *int32
	SetRequestId(v string) *DescribeTTSConfigResponseBody
	GetRequestId() *string
	SetSpeechRate(v int32) *DescribeTTSConfigResponseBody
	GetSpeechRate() *int32
	SetTtsOverrides(v string) *DescribeTTSConfigResponseBody
	GetTtsOverrides() *string
	SetVoice(v string) *DescribeTTSConfigResponseBody
	GetVoice() *string
	SetVolume(v int32) *DescribeTTSConfigResponseBody
	GetVolume() *int32
}

type DescribeTTSConfigResponseBody struct {
	AliCustomizedVoice *string `json:"AliCustomizedVoice,omitempty" xml:"AliCustomizedVoice,omitempty"`
	AppKey             *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Engine             *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineXunfei       *string `json:"EngineXunfei,omitempty" xml:"EngineXunfei,omitempty"`
	NlsServiceType     *string `json:"NlsServiceType,omitempty" xml:"NlsServiceType,omitempty"`
	PitchRate          *int32  `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// example:
	//
	// F132DDBA-66C4-5BD3-BF7E-9642FE877158
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// -150
	SpeechRate   *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	TtsOverrides *string `json:"TtsOverrides,omitempty" xml:"TtsOverrides,omitempty"`
	// example:
	//
	// aixia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s DescribeTTSConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTTSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTTSConfigResponseBody) GetAliCustomizedVoice() *string {
	return s.AliCustomizedVoice
}

func (s *DescribeTTSConfigResponseBody) GetAppKey() *string {
	return s.AppKey
}

func (s *DescribeTTSConfigResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeTTSConfigResponseBody) GetEngineXunfei() *string {
	return s.EngineXunfei
}

func (s *DescribeTTSConfigResponseBody) GetNlsServiceType() *string {
	return s.NlsServiceType
}

func (s *DescribeTTSConfigResponseBody) GetPitchRate() *int32 {
	return s.PitchRate
}

func (s *DescribeTTSConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTTSConfigResponseBody) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *DescribeTTSConfigResponseBody) GetTtsOverrides() *string {
	return s.TtsOverrides
}

func (s *DescribeTTSConfigResponseBody) GetVoice() *string {
	return s.Voice
}

func (s *DescribeTTSConfigResponseBody) GetVolume() *int32 {
	return s.Volume
}

func (s *DescribeTTSConfigResponseBody) SetAliCustomizedVoice(v string) *DescribeTTSConfigResponseBody {
	s.AliCustomizedVoice = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetAppKey(v string) *DescribeTTSConfigResponseBody {
	s.AppKey = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetEngine(v string) *DescribeTTSConfigResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetEngineXunfei(v string) *DescribeTTSConfigResponseBody {
	s.EngineXunfei = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetNlsServiceType(v string) *DescribeTTSConfigResponseBody {
	s.NlsServiceType = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetPitchRate(v int32) *DescribeTTSConfigResponseBody {
	s.PitchRate = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetRequestId(v string) *DescribeTTSConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetSpeechRate(v int32) *DescribeTTSConfigResponseBody {
	s.SpeechRate = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetTtsOverrides(v string) *DescribeTTSConfigResponseBody {
	s.TtsOverrides = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetVoice(v string) *DescribeTTSConfigResponseBody {
	s.Voice = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetVolume(v int32) *DescribeTTSConfigResponseBody {
	s.Volume = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
