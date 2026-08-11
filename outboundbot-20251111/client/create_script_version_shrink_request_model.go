// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateScriptVersionShrinkRequest
	GetInstanceId() *string
	SetInteractionConfigShrink(v string) *CreateScriptVersionShrinkRequest
	GetInteractionConfigShrink() *string
	SetLabelConfigsShrink(v string) *CreateScriptVersionShrinkRequest
	GetLabelConfigsShrink() *string
	SetScriptId(v string) *CreateScriptVersionShrinkRequest
	GetScriptId() *string
	SetScriptProfileShrink(v string) *CreateScriptVersionShrinkRequest
	GetScriptProfileShrink() *string
	SetSourceVersionId(v string) *CreateScriptVersionShrinkRequest
	GetSourceVersionId() *string
	SetSynthesizerConfigShrink(v string) *CreateScriptVersionShrinkRequest
	GetSynthesizerConfigShrink() *string
	SetTranscriberConfigShrink(v string) *CreateScriptVersionShrinkRequest
	GetTranscriberConfigShrink() *string
}

type CreateScriptVersionShrinkRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The interaction configuration.
	InteractionConfigShrink *string `json:"InteractionConfig,omitempty" xml:"InteractionConfig,omitempty"`
	// The label configurations.
	LabelConfigsShrink *string `json:"LabelConfigs,omitempty" xml:"LabelConfigs,omitempty"`
	// The scenario ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The dialogue capability configuration.
	ScriptProfileShrink *string `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty"`
	// The source version ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b26
	SourceVersionId *string `json:"SourceVersionId,omitempty" xml:"SourceVersionId,omitempty"`
	// The TTS configuration.
	SynthesizerConfigShrink *string `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty"`
	// The ASR configuration.
	TranscriberConfigShrink *string `json:"TranscriberConfig,omitempty" xml:"TranscriberConfig,omitempty"`
}

func (s CreateScriptVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateScriptVersionShrinkRequest) GetInteractionConfigShrink() *string {
	return s.InteractionConfigShrink
}

func (s *CreateScriptVersionShrinkRequest) GetLabelConfigsShrink() *string {
	return s.LabelConfigsShrink
}

func (s *CreateScriptVersionShrinkRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateScriptVersionShrinkRequest) GetScriptProfileShrink() *string {
	return s.ScriptProfileShrink
}

func (s *CreateScriptVersionShrinkRequest) GetSourceVersionId() *string {
	return s.SourceVersionId
}

func (s *CreateScriptVersionShrinkRequest) GetSynthesizerConfigShrink() *string {
	return s.SynthesizerConfigShrink
}

func (s *CreateScriptVersionShrinkRequest) GetTranscriberConfigShrink() *string {
	return s.TranscriberConfigShrink
}

func (s *CreateScriptVersionShrinkRequest) SetInstanceId(v string) *CreateScriptVersionShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScriptVersionShrinkRequest) SetInteractionConfigShrink(v string) *CreateScriptVersionShrinkRequest {
	s.InteractionConfigShrink = &v
	return s
}

func (s *CreateScriptVersionShrinkRequest) SetLabelConfigsShrink(v string) *CreateScriptVersionShrinkRequest {
	s.LabelConfigsShrink = &v
	return s
}

func (s *CreateScriptVersionShrinkRequest) SetScriptId(v string) *CreateScriptVersionShrinkRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateScriptVersionShrinkRequest) SetScriptProfileShrink(v string) *CreateScriptVersionShrinkRequest {
	s.ScriptProfileShrink = &v
	return s
}

func (s *CreateScriptVersionShrinkRequest) SetSourceVersionId(v string) *CreateScriptVersionShrinkRequest {
	s.SourceVersionId = &v
	return s
}

func (s *CreateScriptVersionShrinkRequest) SetSynthesizerConfigShrink(v string) *CreateScriptVersionShrinkRequest {
	s.SynthesizerConfigShrink = &v
	return s
}

func (s *CreateScriptVersionShrinkRequest) SetTranscriberConfigShrink(v string) *CreateScriptVersionShrinkRequest {
	s.TranscriberConfigShrink = &v
	return s
}

func (s *CreateScriptVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
