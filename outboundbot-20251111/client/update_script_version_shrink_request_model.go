// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScriptVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateScriptVersionShrinkRequest
	GetInstanceId() *string
	SetInteractionConfigShrink(v string) *UpdateScriptVersionShrinkRequest
	GetInteractionConfigShrink() *string
	SetLabelConfigsShrink(v string) *UpdateScriptVersionShrinkRequest
	GetLabelConfigsShrink() *string
	SetScriptId(v string) *UpdateScriptVersionShrinkRequest
	GetScriptId() *string
	SetScriptProfileShrink(v string) *UpdateScriptVersionShrinkRequest
	GetScriptProfileShrink() *string
	SetSynthesizerConfigShrink(v string) *UpdateScriptVersionShrinkRequest
	GetSynthesizerConfigShrink() *string
	SetTranscriberConfigShrink(v string) *UpdateScriptVersionShrinkRequest
	GetTranscriberConfigShrink() *string
	SetVersionId(v string) *UpdateScriptVersionShrinkRequest
	GetVersionId() *string
}

type UpdateScriptVersionShrinkRequest struct {
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
	// The TTS configuration.
	SynthesizerConfigShrink *string `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty"`
	// The ASR configuration.
	TranscriberConfigShrink *string `json:"TranscriberConfig,omitempty" xml:"TranscriberConfig,omitempty"`
	// The version ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b26
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s UpdateScriptVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateScriptVersionShrinkRequest) GetInteractionConfigShrink() *string {
	return s.InteractionConfigShrink
}

func (s *UpdateScriptVersionShrinkRequest) GetLabelConfigsShrink() *string {
	return s.LabelConfigsShrink
}

func (s *UpdateScriptVersionShrinkRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *UpdateScriptVersionShrinkRequest) GetScriptProfileShrink() *string {
	return s.ScriptProfileShrink
}

func (s *UpdateScriptVersionShrinkRequest) GetSynthesizerConfigShrink() *string {
	return s.SynthesizerConfigShrink
}

func (s *UpdateScriptVersionShrinkRequest) GetTranscriberConfigShrink() *string {
	return s.TranscriberConfigShrink
}

func (s *UpdateScriptVersionShrinkRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *UpdateScriptVersionShrinkRequest) SetInstanceId(v string) *UpdateScriptVersionShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateScriptVersionShrinkRequest) SetInteractionConfigShrink(v string) *UpdateScriptVersionShrinkRequest {
	s.InteractionConfigShrink = &v
	return s
}

func (s *UpdateScriptVersionShrinkRequest) SetLabelConfigsShrink(v string) *UpdateScriptVersionShrinkRequest {
	s.LabelConfigsShrink = &v
	return s
}

func (s *UpdateScriptVersionShrinkRequest) SetScriptId(v string) *UpdateScriptVersionShrinkRequest {
	s.ScriptId = &v
	return s
}

func (s *UpdateScriptVersionShrinkRequest) SetScriptProfileShrink(v string) *UpdateScriptVersionShrinkRequest {
	s.ScriptProfileShrink = &v
	return s
}

func (s *UpdateScriptVersionShrinkRequest) SetSynthesizerConfigShrink(v string) *UpdateScriptVersionShrinkRequest {
	s.SynthesizerConfigShrink = &v
	return s
}

func (s *UpdateScriptVersionShrinkRequest) SetTranscriberConfigShrink(v string) *UpdateScriptVersionShrinkRequest {
	s.TranscriberConfigShrink = &v
	return s
}

func (s *UpdateScriptVersionShrinkRequest) SetVersionId(v string) *UpdateScriptVersionShrinkRequest {
	s.VersionId = &v
	return s
}

func (s *UpdateScriptVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
