// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBundleUrl(v string) *UpdateSkillShrinkRequest
	GetBundleUrl() *string
	SetDescription(v string) *UpdateSkillShrinkRequest
	GetDescription() *string
	SetExpectedVersion(v int32) *UpdateSkillShrinkRequest
	GetExpectedVersion() *int32
	SetExtraShrink(v string) *UpdateSkillShrinkRequest
	GetExtraShrink() *string
	SetName(v string) *UpdateSkillShrinkRequest
	GetName() *string
	SetSkillMdOverride(v string) *UpdateSkillShrinkRequest
	GetSkillMdOverride() *string
	SetVersionNote(v string) *UpdateSkillShrinkRequest
	GetVersionNote() *string
	SetVisibilityScopeShrink(v string) *UpdateSkillShrinkRequest
	GetVisibilityScopeShrink() *string
}

type UpdateSkillShrinkRequest struct {
	// The downloadable URL (HTTP/HTTPS) of the bundle.zip file. Mutually exclusive with SkillMdOverride. If specified, the bundle is replaced.
	//
	// example:
	//
	// https://example.com/skill.zip
	BundleUrl *string `json:"BundleUrl,omitempty" xml:"BundleUrl,omitempty"`
	// The Skill description.
	//
	// example:
	//
	// 数据分析技能
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expected version number for optimistic locking. If not specified, the update is based on the current highest version.
	//
	// example:
	//
	// -
	ExpectedVersion *int32 `json:"ExpectedVersion,omitempty" xml:"ExpectedVersion,omitempty"`
	// The extended metadata (key-value pairs).
	//
	// example:
	//
	// {\\"appId\\":\\"APP_Q2SDWKIGFWNZTR68K1GQ\\"}
	ExtraShrink *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The name of the Skill to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-skill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The SKILL.md body content. Mutually exclusive with BundleUrl.
	//
	// example:
	//
	// 把大象放冰箱分为三步，把冰箱门打开，把大象放进去，把冰箱门关上。
	SkillMdOverride *string `json:"SkillMdOverride,omitempty" xml:"SkillMdOverride,omitempty"`
	// The version note.
	//
	// example:
	//
	// 修订说明
	VersionNote *string `json:"VersionNote,omitempty" xml:"VersionNote,omitempty"`
	// The visibility scope. The corresponding field is used based on the visibility level.
	VisibilityScopeShrink *string `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty"`
}

func (s UpdateSkillShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillShrinkRequest) GetBundleUrl() *string {
	return s.BundleUrl
}

func (s *UpdateSkillShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSkillShrinkRequest) GetExpectedVersion() *int32 {
	return s.ExpectedVersion
}

func (s *UpdateSkillShrinkRequest) GetExtraShrink() *string {
	return s.ExtraShrink
}

func (s *UpdateSkillShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSkillShrinkRequest) GetSkillMdOverride() *string {
	return s.SkillMdOverride
}

func (s *UpdateSkillShrinkRequest) GetVersionNote() *string {
	return s.VersionNote
}

func (s *UpdateSkillShrinkRequest) GetVisibilityScopeShrink() *string {
	return s.VisibilityScopeShrink
}

func (s *UpdateSkillShrinkRequest) SetBundleUrl(v string) *UpdateSkillShrinkRequest {
	s.BundleUrl = &v
	return s
}

func (s *UpdateSkillShrinkRequest) SetDescription(v string) *UpdateSkillShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateSkillShrinkRequest) SetExpectedVersion(v int32) *UpdateSkillShrinkRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpdateSkillShrinkRequest) SetExtraShrink(v string) *UpdateSkillShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *UpdateSkillShrinkRequest) SetName(v string) *UpdateSkillShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateSkillShrinkRequest) SetSkillMdOverride(v string) *UpdateSkillShrinkRequest {
	s.SkillMdOverride = &v
	return s
}

func (s *UpdateSkillShrinkRequest) SetVersionNote(v string) *UpdateSkillShrinkRequest {
	s.VersionNote = &v
	return s
}

func (s *UpdateSkillShrinkRequest) SetVisibilityScopeShrink(v string) *UpdateSkillShrinkRequest {
	s.VisibilityScopeShrink = &v
	return s
}

func (s *UpdateSkillShrinkRequest) Validate() error {
	return dara.Validate(s)
}
