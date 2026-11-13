// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBundleUrl(v string) *CreateSkillRequest
	GetBundleUrl() *string
	SetDescription(v string) *CreateSkillRequest
	GetDescription() *string
	SetExtra(v map[string]interface{}) *CreateSkillRequest
	GetExtra() map[string]interface{}
	SetName(v string) *CreateSkillRequest
	GetName() *string
	SetSkillMdOverride(v string) *CreateSkillRequest
	GetSkillMdOverride() *string
	SetVersionNote(v string) *CreateSkillRequest
	GetVersionNote() *string
	SetVisibility(v string) *CreateSkillRequest
	GetVisibility() *string
	SetVisibilityScope(v *CreateSkillRequestVisibilityScope) *CreateSkillRequest
	GetVisibilityScope() *CreateSkillRequestVisibilityScope
}

type CreateSkillRequest struct {
	// The downloadable URL (HTTP or HTTPS) of the bundle.zip file. This parameter and SkillMdOverride are mutually exclusive.
	//
	// example:
	//
	// https://example.com/skill.zip
	BundleUrl *string `json:"BundleUrl,omitempty" xml:"BundleUrl,omitempty"`
	// The description of the skill.
	//
	// example:
	//
	// Data analytics skill
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extended metadata in key-value pairs.
	//
	// example:
	//
	// {"appId":"APP_CWJMV36CT9SAFW1QEHX7"}
	Extra map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The name of the skill. The name must be unique within the current account.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-skill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The body of the SKILL.md file. This parameter and BundleUrl are mutually exclusive. If you do not provide a bundle, use this field to create a lightweight skill that contains only the SKILL.md file.
	//
	// example:
	//
	// ---
	//
	// name: example-skill
	//
	// description: Sample skill.
	//
	// ---
	//
	// # Sample skill
	//
	// Summarize the text provided by the user
	SkillMdOverride *string `json:"SkillMdOverride,omitempty" xml:"SkillMdOverride,omitempty"`
	// The release notes of the version.
	//
	// example:
	//
	// Initial version
	VersionNote *string `json:"VersionNote,omitempty" xml:"VersionNote,omitempty"`
	// The visibility level. Valid values: TENANT: visible within the account. PROJECT: visible to specified projects. USER: visible to specified users.
	//
	// example:
	//
	// TENANT
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	// The visibility scope. Specify the corresponding field based on the value of Visibility.
	VisibilityScope *CreateSkillRequestVisibilityScope `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty" type:"Struct"`
}

func (s CreateSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillRequest) GetBundleUrl() *string {
	return s.BundleUrl
}

func (s *CreateSkillRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSkillRequest) GetExtra() map[string]interface{} {
	return s.Extra
}

func (s *CreateSkillRequest) GetName() *string {
	return s.Name
}

func (s *CreateSkillRequest) GetSkillMdOverride() *string {
	return s.SkillMdOverride
}

func (s *CreateSkillRequest) GetVersionNote() *string {
	return s.VersionNote
}

func (s *CreateSkillRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateSkillRequest) GetVisibilityScope() *CreateSkillRequestVisibilityScope {
	return s.VisibilityScope
}

func (s *CreateSkillRequest) SetBundleUrl(v string) *CreateSkillRequest {
	s.BundleUrl = &v
	return s
}

func (s *CreateSkillRequest) SetDescription(v string) *CreateSkillRequest {
	s.Description = &v
	return s
}

func (s *CreateSkillRequest) SetExtra(v map[string]interface{}) *CreateSkillRequest {
	s.Extra = v
	return s
}

func (s *CreateSkillRequest) SetName(v string) *CreateSkillRequest {
	s.Name = &v
	return s
}

func (s *CreateSkillRequest) SetSkillMdOverride(v string) *CreateSkillRequest {
	s.SkillMdOverride = &v
	return s
}

func (s *CreateSkillRequest) SetVersionNote(v string) *CreateSkillRequest {
	s.VersionNote = &v
	return s
}

func (s *CreateSkillRequest) SetVisibility(v string) *CreateSkillRequest {
	s.Visibility = &v
	return s
}

func (s *CreateSkillRequest) SetVisibilityScope(v *CreateSkillRequestVisibilityScope) *CreateSkillRequest {
	s.VisibilityScope = v
	return s
}

func (s *CreateSkillRequest) Validate() error {
	if s.VisibilityScope != nil {
		if err := s.VisibilityScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSkillRequestVisibilityScope struct {
	// The IDs of the projects in which the skill is visible. This field takes effect when Visibility is set to PROJECT.
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The IDs of the users who can see the skill. This field takes effect when Visibility is set to USER.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s CreateSkillRequestVisibilityScope) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillRequestVisibilityScope) GoString() string {
	return s.String()
}

func (s *CreateSkillRequestVisibilityScope) GetProjectIds() []*string {
	return s.ProjectIds
}

func (s *CreateSkillRequestVisibilityScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *CreateSkillRequestVisibilityScope) SetProjectIds(v []*string) *CreateSkillRequestVisibilityScope {
	s.ProjectIds = v
	return s
}

func (s *CreateSkillRequestVisibilityScope) SetUserIds(v []*string) *CreateSkillRequestVisibilityScope {
	s.UserIds = v
	return s
}

func (s *CreateSkillRequestVisibilityScope) Validate() error {
	return dara.Validate(s)
}
