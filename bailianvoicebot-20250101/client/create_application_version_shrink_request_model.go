// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreateApplicationVersionShrinkRequest
	GetApplicationId() *string
	SetBusinessUnitId(v string) *CreateApplicationVersionShrinkRequest
	GetBusinessUnitId() *string
	SetInteractionConfigShrink(v string) *CreateApplicationVersionShrinkRequest
	GetInteractionConfigShrink() *string
	SetScriptProfileShrink(v string) *CreateApplicationVersionShrinkRequest
	GetScriptProfileShrink() *string
	SetSourceVersionId(v string) *CreateApplicationVersionShrinkRequest
	GetSourceVersionId() *string
	SetSynthesizerConfigShrink(v string) *CreateApplicationVersionShrinkRequest
	GetSynthesizerConfigShrink() *string
	SetTranscriberConfigShrink(v string) *CreateApplicationVersionShrinkRequest
	GetTranscriberConfigShrink() *string
}

type CreateApplicationVersionShrinkRequest struct {
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
	ScriptProfileShrink     *string `json:"ScriptProfile,omitempty" xml:"ScriptProfile,omitempty"`
	// example:
	//
	// 20904943-f711-494f-9f1f-e7f340f37707
	SourceVersionId         *string `json:"SourceVersionId,omitempty" xml:"SourceVersionId,omitempty"`
	SynthesizerConfigShrink *string `json:"SynthesizerConfig,omitempty" xml:"SynthesizerConfig,omitempty"`
	TranscriberConfigShrink *string `json:"TranscriberConfig,omitempty" xml:"TranscriberConfig,omitempty"`
}

func (s CreateApplicationVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateApplicationVersionShrinkRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *CreateApplicationVersionShrinkRequest) GetInteractionConfigShrink() *string {
	return s.InteractionConfigShrink
}

func (s *CreateApplicationVersionShrinkRequest) GetScriptProfileShrink() *string {
	return s.ScriptProfileShrink
}

func (s *CreateApplicationVersionShrinkRequest) GetSourceVersionId() *string {
	return s.SourceVersionId
}

func (s *CreateApplicationVersionShrinkRequest) GetSynthesizerConfigShrink() *string {
	return s.SynthesizerConfigShrink
}

func (s *CreateApplicationVersionShrinkRequest) GetTranscriberConfigShrink() *string {
	return s.TranscriberConfigShrink
}

func (s *CreateApplicationVersionShrinkRequest) SetApplicationId(v string) *CreateApplicationVersionShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationVersionShrinkRequest) SetBusinessUnitId(v string) *CreateApplicationVersionShrinkRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *CreateApplicationVersionShrinkRequest) SetInteractionConfigShrink(v string) *CreateApplicationVersionShrinkRequest {
	s.InteractionConfigShrink = &v
	return s
}

func (s *CreateApplicationVersionShrinkRequest) SetScriptProfileShrink(v string) *CreateApplicationVersionShrinkRequest {
	s.ScriptProfileShrink = &v
	return s
}

func (s *CreateApplicationVersionShrinkRequest) SetSourceVersionId(v string) *CreateApplicationVersionShrinkRequest {
	s.SourceVersionId = &v
	return s
}

func (s *CreateApplicationVersionShrinkRequest) SetSynthesizerConfigShrink(v string) *CreateApplicationVersionShrinkRequest {
	s.SynthesizerConfigShrink = &v
	return s
}

func (s *CreateApplicationVersionShrinkRequest) SetTranscriberConfigShrink(v string) *CreateApplicationVersionShrinkRequest {
	s.TranscriberConfigShrink = &v
	return s
}

func (s *CreateApplicationVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
