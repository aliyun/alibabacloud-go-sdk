// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlidingAssistantShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *CreateAlidingAssistantShrinkRequest
	GetAppCode() *string
	SetDescription(v string) *CreateAlidingAssistantShrinkRequest
	GetDescription() *string
	SetExtShrink(v string) *CreateAlidingAssistantShrinkRequest
	GetExtShrink() *string
	SetIcon(v string) *CreateAlidingAssistantShrinkRequest
	GetIcon() *string
	SetInstructions(v string) *CreateAlidingAssistantShrinkRequest
	GetInstructions() *string
	SetName(v string) *CreateAlidingAssistantShrinkRequest
	GetName() *string
	SetRecommendPromptsShrink(v string) *CreateAlidingAssistantShrinkRequest
	GetRecommendPromptsShrink() *string
	SetSource(v int32) *CreateAlidingAssistantShrinkRequest
	GetSource() *int32
	SetSourceIdentityId(v string) *CreateAlidingAssistantShrinkRequest
	GetSourceIdentityId() *string
	SetTenantContextShrink(v string) *CreateAlidingAssistantShrinkRequest
	GetTenantContextShrink() *string
	SetWelcomeContent(v string) *CreateAlidingAssistantShrinkRequest
	GetWelcomeContent() *string
}

type CreateAlidingAssistantShrinkRequest struct {
	// example:
	//
	// f5cb37a0fb44441ab7b74c6f4a679dd3
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExtShrink   *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// @lADPDetfgMsFFUvNAkjNAkg
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// This parameter is required.
	Instructions *string `json:"Instructions,omitempty" xml:"Instructions,omitempty"`
	// This parameter is required.
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RecommendPromptsShrink *string `json:"RecommendPrompts,omitempty" xml:"RecommendPrompts,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// chatBot-123
	SourceIdentityId    *string `json:"SourceIdentityId,omitempty" xml:"SourceIdentityId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	WelcomeContent *string `json:"WelcomeContent,omitempty" xml:"WelcomeContent,omitempty"`
}

func (s CreateAlidingAssistantShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlidingAssistantShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAlidingAssistantShrinkRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *CreateAlidingAssistantShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAlidingAssistantShrinkRequest) GetExtShrink() *string {
	return s.ExtShrink
}

func (s *CreateAlidingAssistantShrinkRequest) GetIcon() *string {
	return s.Icon
}

func (s *CreateAlidingAssistantShrinkRequest) GetInstructions() *string {
	return s.Instructions
}

func (s *CreateAlidingAssistantShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateAlidingAssistantShrinkRequest) GetRecommendPromptsShrink() *string {
	return s.RecommendPromptsShrink
}

func (s *CreateAlidingAssistantShrinkRequest) GetSource() *int32 {
	return s.Source
}

func (s *CreateAlidingAssistantShrinkRequest) GetSourceIdentityId() *string {
	return s.SourceIdentityId
}

func (s *CreateAlidingAssistantShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateAlidingAssistantShrinkRequest) GetWelcomeContent() *string {
	return s.WelcomeContent
}

func (s *CreateAlidingAssistantShrinkRequest) SetAppCode(v string) *CreateAlidingAssistantShrinkRequest {
	s.AppCode = &v
	return s
}

func (s *CreateAlidingAssistantShrinkRequest) SetDescription(v string) *CreateAlidingAssistantShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAlidingAssistantShrinkRequest) SetExtShrink(v string) *CreateAlidingAssistantShrinkRequest {
	s.ExtShrink = &v
	return s
}

func (s *CreateAlidingAssistantShrinkRequest) SetIcon(v string) *CreateAlidingAssistantShrinkRequest {
	s.Icon = &v
	return s
}

func (s *CreateAlidingAssistantShrinkRequest) SetInstructions(v string) *CreateAlidingAssistantShrinkRequest {
	s.Instructions = &v
	return s
}

func (s *CreateAlidingAssistantShrinkRequest) SetName(v string) *CreateAlidingAssistantShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateAlidingAssistantShrinkRequest) SetRecommendPromptsShrink(v string) *CreateAlidingAssistantShrinkRequest {
	s.RecommendPromptsShrink = &v
	return s
}

func (s *CreateAlidingAssistantShrinkRequest) SetSource(v int32) *CreateAlidingAssistantShrinkRequest {
	s.Source = &v
	return s
}

func (s *CreateAlidingAssistantShrinkRequest) SetSourceIdentityId(v string) *CreateAlidingAssistantShrinkRequest {
	s.SourceIdentityId = &v
	return s
}

func (s *CreateAlidingAssistantShrinkRequest) SetTenantContextShrink(v string) *CreateAlidingAssistantShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateAlidingAssistantShrinkRequest) SetWelcomeContent(v string) *CreateAlidingAssistantShrinkRequest {
	s.WelcomeContent = &v
	return s
}

func (s *CreateAlidingAssistantShrinkRequest) Validate() error {
	return dara.Validate(s)
}
