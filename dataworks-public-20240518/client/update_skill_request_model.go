// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBundleUrl(v string) *UpdateSkillRequest
	GetBundleUrl() *string
	SetDescription(v string) *UpdateSkillRequest
	GetDescription() *string
	SetExpectedVersion(v int32) *UpdateSkillRequest
	GetExpectedVersion() *int32
	SetExtra(v map[string]interface{}) *UpdateSkillRequest
	GetExtra() map[string]interface{}
	SetName(v string) *UpdateSkillRequest
	GetName() *string
	SetSkillMdOverride(v string) *UpdateSkillRequest
	GetSkillMdOverride() *string
	SetVersionNote(v string) *UpdateSkillRequest
	GetVersionNote() *string
	SetVisibilityScope(v *UpdateSkillRequestVisibilityScope) *UpdateSkillRequest
	GetVisibilityScope() *UpdateSkillRequestVisibilityScope
}

type UpdateSkillRequest struct {
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
	// 1
	ExpectedVersion *int32 `json:"ExpectedVersion,omitempty" xml:"ExpectedVersion,omitempty"`
	// The extended metadata (key-value pairs).
	//
	// example:
	//
	// {\\"appId\\":\\"APP_Q2SDWKIGFWNZTR68K1GQ\\"}
	Extra map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
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
	VisibilityScope *UpdateSkillRequestVisibilityScope `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty" type:"Struct"`
}

func (s UpdateSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillRequest) GetBundleUrl() *string {
	return s.BundleUrl
}

func (s *UpdateSkillRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSkillRequest) GetExpectedVersion() *int32 {
	return s.ExpectedVersion
}

func (s *UpdateSkillRequest) GetExtra() map[string]interface{} {
	return s.Extra
}

func (s *UpdateSkillRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSkillRequest) GetSkillMdOverride() *string {
	return s.SkillMdOverride
}

func (s *UpdateSkillRequest) GetVersionNote() *string {
	return s.VersionNote
}

func (s *UpdateSkillRequest) GetVisibilityScope() *UpdateSkillRequestVisibilityScope {
	return s.VisibilityScope
}

func (s *UpdateSkillRequest) SetBundleUrl(v string) *UpdateSkillRequest {
	s.BundleUrl = &v
	return s
}

func (s *UpdateSkillRequest) SetDescription(v string) *UpdateSkillRequest {
	s.Description = &v
	return s
}

func (s *UpdateSkillRequest) SetExpectedVersion(v int32) *UpdateSkillRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpdateSkillRequest) SetExtra(v map[string]interface{}) *UpdateSkillRequest {
	s.Extra = v
	return s
}

func (s *UpdateSkillRequest) SetName(v string) *UpdateSkillRequest {
	s.Name = &v
	return s
}

func (s *UpdateSkillRequest) SetSkillMdOverride(v string) *UpdateSkillRequest {
	s.SkillMdOverride = &v
	return s
}

func (s *UpdateSkillRequest) SetVersionNote(v string) *UpdateSkillRequest {
	s.VersionNote = &v
	return s
}

func (s *UpdateSkillRequest) SetVisibilityScope(v *UpdateSkillRequestVisibilityScope) *UpdateSkillRequest {
	s.VisibilityScope = v
	return s
}

func (s *UpdateSkillRequest) Validate() error {
	if s.VisibilityScope != nil {
		if err := s.VisibilityScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSkillRequestVisibilityScope struct {
	// The list of visible project IDs.
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The list of visible user IDs.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s UpdateSkillRequestVisibilityScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillRequestVisibilityScope) GoString() string {
	return s.String()
}

func (s *UpdateSkillRequestVisibilityScope) GetProjectIds() []*string {
	return s.ProjectIds
}

func (s *UpdateSkillRequestVisibilityScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *UpdateSkillRequestVisibilityScope) SetProjectIds(v []*string) *UpdateSkillRequestVisibilityScope {
	s.ProjectIds = v
	return s
}

func (s *UpdateSkillRequestVisibilityScope) SetUserIds(v []*string) *UpdateSkillRequestVisibilityScope {
	s.UserIds = v
	return s
}

func (s *UpdateSkillRequestVisibilityScope) Validate() error {
	return dara.Validate(s)
}
