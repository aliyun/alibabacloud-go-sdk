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
	// example:
	//
	// https://example.com/skill.zip
	BundleUrl   *string `json:"BundleUrl,omitempty" xml:"BundleUrl,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {"appId":"APP_CWJMV36CT9SAFW1QEHX7"}
	Extra map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
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
	Visibility      *string                            `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
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
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	UserIds    []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
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
