// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewVoiceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *PreviewVoiceShrinkRequest
	GetBusinessUnitId() *string
	SetModel(v string) *PreviewVoiceShrinkRequest
	GetModel() *string
	SetNlsAccessType(v string) *PreviewVoiceShrinkRequest
	GetNlsAccessType() *string
	SetNlsEngine(v string) *PreviewVoiceShrinkRequest
	GetNlsEngine() *string
	SetParamsShrink(v string) *PreviewVoiceShrinkRequest
	GetParamsShrink() *string
	SetText(v string) *PreviewVoiceShrinkRequest
	GetText() *string
	SetVoice(v string) *PreviewVoiceShrinkRequest
	GetVoice() *string
}

type PreviewVoiceShrinkRequest struct {
	// The ID of the Model Studio business unit.
	//
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	// The TTS model.
	//
	// example:
	//
	// Qwen
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The TTS access type.
	//
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// The TTS engine.
	//
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// The synthesis parameters.
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The text to synthesize for the preview.
	//
	// example:
	//
	// 你好，很高兴认识你
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The voice for synthesis.
	//
	// example:
	//
	// Cherry
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
}

func (s PreviewVoiceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewVoiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *PreviewVoiceShrinkRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *PreviewVoiceShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *PreviewVoiceShrinkRequest) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *PreviewVoiceShrinkRequest) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *PreviewVoiceShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *PreviewVoiceShrinkRequest) GetText() *string {
	return s.Text
}

func (s *PreviewVoiceShrinkRequest) GetVoice() *string {
	return s.Voice
}

func (s *PreviewVoiceShrinkRequest) SetBusinessUnitId(v string) *PreviewVoiceShrinkRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *PreviewVoiceShrinkRequest) SetModel(v string) *PreviewVoiceShrinkRequest {
	s.Model = &v
	return s
}

func (s *PreviewVoiceShrinkRequest) SetNlsAccessType(v string) *PreviewVoiceShrinkRequest {
	s.NlsAccessType = &v
	return s
}

func (s *PreviewVoiceShrinkRequest) SetNlsEngine(v string) *PreviewVoiceShrinkRequest {
	s.NlsEngine = &v
	return s
}

func (s *PreviewVoiceShrinkRequest) SetParamsShrink(v string) *PreviewVoiceShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *PreviewVoiceShrinkRequest) SetText(v string) *PreviewVoiceShrinkRequest {
	s.Text = &v
	return s
}

func (s *PreviewVoiceShrinkRequest) SetVoice(v string) *PreviewVoiceShrinkRequest {
	s.Voice = &v
	return s
}

func (s *PreviewVoiceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
