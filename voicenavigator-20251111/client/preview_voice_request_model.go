// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PreviewVoiceRequest
	GetInstanceId() *string
	SetModel(v string) *PreviewVoiceRequest
	GetModel() *string
	SetNlsAccessType(v string) *PreviewVoiceRequest
	GetNlsAccessType() *string
	SetNlsEngine(v string) *PreviewVoiceRequest
	GetNlsEngine() *string
	SetParams(v *PreviewVoiceRequestParams) *PreviewVoiceRequest
	GetParams() *PreviewVoiceRequestParams
	SetText(v string) *PreviewVoiceRequest
	GetText() *string
	SetVoice(v string) *PreviewVoiceRequest
	GetVoice() *string
}

type PreviewVoiceRequest struct {
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Qwen
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// example:
	//
	// BAILIAN
	NlsEngine *string                    `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	Params    *PreviewVoiceRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
	Text      *string                    `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// Cherry
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
}

func (s PreviewVoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewVoiceRequest) GoString() string {
	return s.String()
}

func (s *PreviewVoiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PreviewVoiceRequest) GetModel() *string {
	return s.Model
}

func (s *PreviewVoiceRequest) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *PreviewVoiceRequest) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *PreviewVoiceRequest) GetParams() *PreviewVoiceRequestParams {
	return s.Params
}

func (s *PreviewVoiceRequest) GetText() *string {
	return s.Text
}

func (s *PreviewVoiceRequest) GetVoice() *string {
	return s.Voice
}

func (s *PreviewVoiceRequest) SetInstanceId(v string) *PreviewVoiceRequest {
	s.InstanceId = &v
	return s
}

func (s *PreviewVoiceRequest) SetModel(v string) *PreviewVoiceRequest {
	s.Model = &v
	return s
}

func (s *PreviewVoiceRequest) SetNlsAccessType(v string) *PreviewVoiceRequest {
	s.NlsAccessType = &v
	return s
}

func (s *PreviewVoiceRequest) SetNlsEngine(v string) *PreviewVoiceRequest {
	s.NlsEngine = &v
	return s
}

func (s *PreviewVoiceRequest) SetParams(v *PreviewVoiceRequestParams) *PreviewVoiceRequest {
	s.Params = v
	return s
}

func (s *PreviewVoiceRequest) SetText(v string) *PreviewVoiceRequest {
	s.Text = &v
	return s
}

func (s *PreviewVoiceRequest) SetVoice(v string) *PreviewVoiceRequest {
	s.Voice = &v
	return s
}

func (s *PreviewVoiceRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewVoiceRequestParams struct {
	// example:
	//
	// 0
	PitchRate *float32 `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// example:
	//
	// 0
	SpeechRate *float32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s PreviewVoiceRequestParams) String() string {
	return dara.Prettify(s)
}

func (s PreviewVoiceRequestParams) GoString() string {
	return s.String()
}

func (s *PreviewVoiceRequestParams) GetPitchRate() *float32 {
	return s.PitchRate
}

func (s *PreviewVoiceRequestParams) GetSpeechRate() *float32 {
	return s.SpeechRate
}

func (s *PreviewVoiceRequestParams) GetVolume() *int32 {
	return s.Volume
}

func (s *PreviewVoiceRequestParams) SetPitchRate(v float32) *PreviewVoiceRequestParams {
	s.PitchRate = &v
	return s
}

func (s *PreviewVoiceRequestParams) SetSpeechRate(v float32) *PreviewVoiceRequestParams {
	s.SpeechRate = &v
	return s
}

func (s *PreviewVoiceRequestParams) SetVolume(v int32) *PreviewVoiceRequestParams {
	s.Volume = &v
	return s
}

func (s *PreviewVoiceRequestParams) Validate() error {
	return dara.Validate(s)
}
