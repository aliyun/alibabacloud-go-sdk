// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBundleUrl(v string) *CreateSkillShrinkRequest
	GetBundleUrl() *string
	SetDescription(v string) *CreateSkillShrinkRequest
	GetDescription() *string
	SetExtraShrink(v string) *CreateSkillShrinkRequest
	GetExtraShrink() *string
	SetName(v string) *CreateSkillShrinkRequest
	GetName() *string
	SetSkillMdOverride(v string) *CreateSkillShrinkRequest
	GetSkillMdOverride() *string
	SetVersionNote(v string) *CreateSkillShrinkRequest
	GetVersionNote() *string
	SetVisibility(v string) *CreateSkillShrinkRequest
	GetVisibility() *string
	SetVisibilityScopeShrink(v string) *CreateSkillShrinkRequest
	GetVisibilityScopeShrink() *string
}

type CreateSkillShrinkRequest struct {
	// example:
	//
	// https://example.com/skill.zip
	BundleUrl   *string `json:"BundleUrl,omitempty" xml:"BundleUrl,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {"appId":"APP_CWJMV36CT9SAFW1QEHX7"}
	ExtraShrink *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-skill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// -
	SkillMdOverride *string `json:"SkillMdOverride,omitempty" xml:"SkillMdOverride,omitempty"`
	VersionNote     *string `json:"VersionNote,omitempty" xml:"VersionNote,omitempty"`
	// example:
	//
	// TENANT
	Visibility            *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	VisibilityScopeShrink *string `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty"`
}

func (s CreateSkillShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillShrinkRequest) GetBundleUrl() *string {
	return s.BundleUrl
}

func (s *CreateSkillShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSkillShrinkRequest) GetExtraShrink() *string {
	return s.ExtraShrink
}

func (s *CreateSkillShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateSkillShrinkRequest) GetSkillMdOverride() *string {
	return s.SkillMdOverride
}

func (s *CreateSkillShrinkRequest) GetVersionNote() *string {
	return s.VersionNote
}

func (s *CreateSkillShrinkRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateSkillShrinkRequest) GetVisibilityScopeShrink() *string {
	return s.VisibilityScopeShrink
}

func (s *CreateSkillShrinkRequest) SetBundleUrl(v string) *CreateSkillShrinkRequest {
	s.BundleUrl = &v
	return s
}

func (s *CreateSkillShrinkRequest) SetDescription(v string) *CreateSkillShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateSkillShrinkRequest) SetExtraShrink(v string) *CreateSkillShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *CreateSkillShrinkRequest) SetName(v string) *CreateSkillShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateSkillShrinkRequest) SetSkillMdOverride(v string) *CreateSkillShrinkRequest {
	s.SkillMdOverride = &v
	return s
}

func (s *CreateSkillShrinkRequest) SetVersionNote(v string) *CreateSkillShrinkRequest {
	s.VersionNote = &v
	return s
}

func (s *CreateSkillShrinkRequest) SetVisibility(v string) *CreateSkillShrinkRequest {
	s.Visibility = &v
	return s
}

func (s *CreateSkillShrinkRequest) SetVisibilityScopeShrink(v string) *CreateSkillShrinkRequest {
	s.VisibilityScopeShrink = &v
	return s
}

func (s *CreateSkillShrinkRequest) Validate() error {
	return dara.Validate(s)
}
