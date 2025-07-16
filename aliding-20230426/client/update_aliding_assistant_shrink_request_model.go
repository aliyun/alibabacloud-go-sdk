// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlidingAssistantShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantId(v string) *UpdateAlidingAssistantShrinkRequest
	GetAssistantId() *string
	SetDescription(v string) *UpdateAlidingAssistantShrinkRequest
	GetDescription() *string
	SetExtShrink(v string) *UpdateAlidingAssistantShrinkRequest
	GetExtShrink() *string
	SetFallbackContent(v string) *UpdateAlidingAssistantShrinkRequest
	GetFallbackContent() *string
	SetFeatureShrink(v string) *UpdateAlidingAssistantShrinkRequest
	GetFeatureShrink() *string
	SetIcon(v string) *UpdateAlidingAssistantShrinkRequest
	GetIcon() *string
	SetInstructions(v string) *UpdateAlidingAssistantShrinkRequest
	GetInstructions() *string
	SetName(v string) *UpdateAlidingAssistantShrinkRequest
	GetName() *string
	SetRecommendPromptsShrink(v string) *UpdateAlidingAssistantShrinkRequest
	GetRecommendPromptsShrink() *string
	SetTenantContextShrink(v string) *UpdateAlidingAssistantShrinkRequest
	GetTenantContextShrink() *string
	SetWelcomeContent(v string) *UpdateAlidingAssistantShrinkRequest
	GetWelcomeContent() *string
}

type UpdateAlidingAssistantShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	AssistantId *string `json:"AssistantId,omitempty" xml:"AssistantId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExtShrink   *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// fallbackContent
	FallbackContent *string `json:"FallbackContent,omitempty" xml:"FallbackContent,omitempty"`
	FeatureShrink   *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	// example:
	//
	// @lADPDetfgMsFFUvNAkjNAkg
	Icon                   *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Instructions           *string `json:"Instructions,omitempty" xml:"Instructions,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RecommendPromptsShrink *string `json:"RecommendPrompts,omitempty" xml:"RecommendPrompts,omitempty"`
	TenantContextShrink    *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	WelcomeContent         *string `json:"WelcomeContent,omitempty" xml:"WelcomeContent,omitempty"`
}

func (s UpdateAlidingAssistantShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlidingAssistantShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlidingAssistantShrinkRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *UpdateAlidingAssistantShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAlidingAssistantShrinkRequest) GetExtShrink() *string {
	return s.ExtShrink
}

func (s *UpdateAlidingAssistantShrinkRequest) GetFallbackContent() *string {
	return s.FallbackContent
}

func (s *UpdateAlidingAssistantShrinkRequest) GetFeatureShrink() *string {
	return s.FeatureShrink
}

func (s *UpdateAlidingAssistantShrinkRequest) GetIcon() *string {
	return s.Icon
}

func (s *UpdateAlidingAssistantShrinkRequest) GetInstructions() *string {
	return s.Instructions
}

func (s *UpdateAlidingAssistantShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAlidingAssistantShrinkRequest) GetRecommendPromptsShrink() *string {
	return s.RecommendPromptsShrink
}

func (s *UpdateAlidingAssistantShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateAlidingAssistantShrinkRequest) GetWelcomeContent() *string {
	return s.WelcomeContent
}

func (s *UpdateAlidingAssistantShrinkRequest) SetAssistantId(v string) *UpdateAlidingAssistantShrinkRequest {
	s.AssistantId = &v
	return s
}

func (s *UpdateAlidingAssistantShrinkRequest) SetDescription(v string) *UpdateAlidingAssistantShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateAlidingAssistantShrinkRequest) SetExtShrink(v string) *UpdateAlidingAssistantShrinkRequest {
	s.ExtShrink = &v
	return s
}

func (s *UpdateAlidingAssistantShrinkRequest) SetFallbackContent(v string) *UpdateAlidingAssistantShrinkRequest {
	s.FallbackContent = &v
	return s
}

func (s *UpdateAlidingAssistantShrinkRequest) SetFeatureShrink(v string) *UpdateAlidingAssistantShrinkRequest {
	s.FeatureShrink = &v
	return s
}

func (s *UpdateAlidingAssistantShrinkRequest) SetIcon(v string) *UpdateAlidingAssistantShrinkRequest {
	s.Icon = &v
	return s
}

func (s *UpdateAlidingAssistantShrinkRequest) SetInstructions(v string) *UpdateAlidingAssistantShrinkRequest {
	s.Instructions = &v
	return s
}

func (s *UpdateAlidingAssistantShrinkRequest) SetName(v string) *UpdateAlidingAssistantShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateAlidingAssistantShrinkRequest) SetRecommendPromptsShrink(v string) *UpdateAlidingAssistantShrinkRequest {
	s.RecommendPromptsShrink = &v
	return s
}

func (s *UpdateAlidingAssistantShrinkRequest) SetTenantContextShrink(v string) *UpdateAlidingAssistantShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateAlidingAssistantShrinkRequest) SetWelcomeContent(v string) *UpdateAlidingAssistantShrinkRequest {
	s.WelcomeContent = &v
	return s
}

func (s *UpdateAlidingAssistantShrinkRequest) Validate() error {
	return dara.Validate(s)
}
