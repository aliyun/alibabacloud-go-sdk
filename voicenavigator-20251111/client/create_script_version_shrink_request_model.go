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
	SetLabelConfigShrink(v string) *CreateScriptVersionShrinkRequest
	GetLabelConfigShrink() *string
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
	// example:
	//
	// abb4aa26-3a8e-43dd-82f8-0c3898c9c67f
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InteractionConfigShrink *string `json:"InteractionConfig,omitempty" xml:"InteractionConfig,omitempty"`
	LabelConfigShrink       *string `json:"LabelConfig,omitempty" xml:"LabelConfig,omitempty"`
	// example:
	//
	// 64241e64-190c-45d1-af66-06f51c07b090
	ScriptId                *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	ScriptProfileShrink     *string `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty"`
	SourceVersionId         *string `json:"SourceVersionId,omitempty" xml:"SourceVersionId,omitempty"`
	SynthesizerConfigShrink *string `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty"`
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

func (s *CreateScriptVersionShrinkRequest) GetLabelConfigShrink() *string {
	return s.LabelConfigShrink
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

func (s *CreateScriptVersionShrinkRequest) SetLabelConfigShrink(v string) *CreateScriptVersionShrinkRequest {
	s.LabelConfigShrink = &v
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
