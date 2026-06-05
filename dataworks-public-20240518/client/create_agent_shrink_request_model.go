// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallableAgentsShrink(v string) *CreateAgentShrinkRequest
	GetCallableAgentsShrink() *string
	SetDescription(v string) *CreateAgentShrinkRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateAgentShrinkRequest
	GetDisplayName() *string
	SetMetadataShrink(v string) *CreateAgentShrinkRequest
	GetMetadataShrink() *string
	SetModelShrink(v string) *CreateAgentShrinkRequest
	GetModelShrink() *string
	SetName(v string) *CreateAgentShrinkRequest
	GetName() *string
	SetSkillsShrink(v string) *CreateAgentShrinkRequest
	GetSkillsShrink() *string
	SetSystemPrompt(v string) *CreateAgentShrinkRequest
	GetSystemPrompt() *string
	SetToolsShrink(v string) *CreateAgentShrinkRequest
	GetToolsShrink() *string
	SetVisibility(v string) *CreateAgentShrinkRequest
	GetVisibility() *string
	SetVisibilityScopeShrink(v string) *CreateAgentShrinkRequest
	GetVisibilityScopeShrink() *string
}

type CreateAgentShrinkRequest struct {
	CallableAgentsShrink *string `json:"CallableAgents,omitempty" xml:"CallableAgents,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName          *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	MetadataShrink       *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	ModelShrink          *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-agent
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SkillsShrink *string `json:"Skills,omitempty" xml:"Skills,omitempty"`
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	ToolsShrink  *string `json:"Tools,omitempty" xml:"Tools,omitempty"`
	// example:
	//
	// TENANT
	Visibility            *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	VisibilityScopeShrink *string `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty"`
}

func (s CreateAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentShrinkRequest) GetCallableAgentsShrink() *string {
	return s.CallableAgentsShrink
}

func (s *CreateAgentShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateAgentShrinkRequest) GetMetadataShrink() *string {
	return s.MetadataShrink
}

func (s *CreateAgentShrinkRequest) GetModelShrink() *string {
	return s.ModelShrink
}

func (s *CreateAgentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateAgentShrinkRequest) GetSkillsShrink() *string {
	return s.SkillsShrink
}

func (s *CreateAgentShrinkRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *CreateAgentShrinkRequest) GetToolsShrink() *string {
	return s.ToolsShrink
}

func (s *CreateAgentShrinkRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateAgentShrinkRequest) GetVisibilityScopeShrink() *string {
	return s.VisibilityScopeShrink
}

func (s *CreateAgentShrinkRequest) SetCallableAgentsShrink(v string) *CreateAgentShrinkRequest {
	s.CallableAgentsShrink = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetDescription(v string) *CreateAgentShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetDisplayName(v string) *CreateAgentShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetMetadataShrink(v string) *CreateAgentShrinkRequest {
	s.MetadataShrink = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetModelShrink(v string) *CreateAgentShrinkRequest {
	s.ModelShrink = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetName(v string) *CreateAgentShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetSkillsShrink(v string) *CreateAgentShrinkRequest {
	s.SkillsShrink = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetSystemPrompt(v string) *CreateAgentShrinkRequest {
	s.SystemPrompt = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetToolsShrink(v string) *CreateAgentShrinkRequest {
	s.ToolsShrink = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetVisibility(v string) *CreateAgentShrinkRequest {
	s.Visibility = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetVisibilityScopeShrink(v string) *CreateAgentShrinkRequest {
	s.VisibilityScopeShrink = &v
	return s
}

func (s *CreateAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
