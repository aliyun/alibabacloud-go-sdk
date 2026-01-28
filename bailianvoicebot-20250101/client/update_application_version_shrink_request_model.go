// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateApplicationVersionShrinkRequest
	GetApplicationId() *string
	SetBusinessUnitId(v string) *UpdateApplicationVersionShrinkRequest
	GetBusinessUnitId() *string
	SetInteractionConfigShrink(v string) *UpdateApplicationVersionShrinkRequest
	GetInteractionConfigShrink() *string
	SetScriptProfileShrink(v string) *UpdateApplicationVersionShrinkRequest
	GetScriptProfileShrink() *string
	SetSynthesizerConfigShrink(v string) *UpdateApplicationVersionShrinkRequest
	GetSynthesizerConfigShrink() *string
	SetTranscriberConfigShrink(v string) *UpdateApplicationVersionShrinkRequest
	GetTranscriberConfigShrink() *string
	SetVersionId(v string) *UpdateApplicationVersionShrinkRequest
	GetVersionId() *string
}

type UpdateApplicationVersionShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a395011f-a247-400f-bc69-28796749fd52
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId          *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	InteractionConfigShrink *string `json:"InteractionConfig,omitempty" xml:"InteractionConfig,omitempty"`
	// This parameter is required.
	ScriptProfileShrink *string `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty"`
	// if can be null:
	// true
	SynthesizerConfigShrink *string `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty"`
	// if can be null:
	// true
	TranscriberConfigShrink *string `json:"TranscriberConfig,omitempty" xml:"TranscriberConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20904943-f711-494f-9f1f-e7f340f37707
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s UpdateApplicationVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationVersionShrinkRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *UpdateApplicationVersionShrinkRequest) GetInteractionConfigShrink() *string {
	return s.InteractionConfigShrink
}

func (s *UpdateApplicationVersionShrinkRequest) GetScriptProfileShrink() *string {
	return s.ScriptProfileShrink
}

func (s *UpdateApplicationVersionShrinkRequest) GetSynthesizerConfigShrink() *string {
	return s.SynthesizerConfigShrink
}

func (s *UpdateApplicationVersionShrinkRequest) GetTranscriberConfigShrink() *string {
	return s.TranscriberConfigShrink
}

func (s *UpdateApplicationVersionShrinkRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *UpdateApplicationVersionShrinkRequest) SetApplicationId(v string) *UpdateApplicationVersionShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationVersionShrinkRequest) SetBusinessUnitId(v string) *UpdateApplicationVersionShrinkRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *UpdateApplicationVersionShrinkRequest) SetInteractionConfigShrink(v string) *UpdateApplicationVersionShrinkRequest {
	s.InteractionConfigShrink = &v
	return s
}

func (s *UpdateApplicationVersionShrinkRequest) SetScriptProfileShrink(v string) *UpdateApplicationVersionShrinkRequest {
	s.ScriptProfileShrink = &v
	return s
}

func (s *UpdateApplicationVersionShrinkRequest) SetSynthesizerConfigShrink(v string) *UpdateApplicationVersionShrinkRequest {
	s.SynthesizerConfigShrink = &v
	return s
}

func (s *UpdateApplicationVersionShrinkRequest) SetTranscriberConfigShrink(v string) *UpdateApplicationVersionShrinkRequest {
	s.TranscriberConfigShrink = &v
	return s
}

func (s *UpdateApplicationVersionShrinkRequest) SetVersionId(v string) *UpdateApplicationVersionShrinkRequest {
	s.VersionId = &v
	return s
}

func (s *UpdateApplicationVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
