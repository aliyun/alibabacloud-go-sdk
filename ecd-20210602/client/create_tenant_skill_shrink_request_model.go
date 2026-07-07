// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTenantSkillShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateTenantSkillShrinkRequest
	GetApiKey() *string
	SetDescription(v string) *CreateTenantSkillShrinkRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateTenantSkillShrinkRequest
	GetDisplayName() *string
	SetEnvVarsShrink(v string) *CreateTenantSkillShrinkRequest
	GetEnvVarsShrink() *string
	SetIconETag(v string) *CreateTenantSkillShrinkRequest
	GetIconETag() *string
	SetSkillChannel(v string) *CreateTenantSkillShrinkRequest
	GetSkillChannel() *string
	SetSkillIcon(v string) *CreateTenantSkillShrinkRequest
	GetSkillIcon() *string
	SetSkillVersion(v string) *CreateTenantSkillShrinkRequest
	GetSkillVersion() *string
	SetSlug(v string) *CreateTenantSkillShrinkRequest
	GetSlug() *string
	SetTaskKey(v string) *CreateTenantSkillShrinkRequest
	GetTaskKey() *string
}

type CreateTenantSkillShrinkRequest struct {
	// The API key of the skill.
	//
	// example:
	//
	// akm-98f66829***
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The description of the skill. Maximum length: 500 characters.
	//
	// example:
	//
	// This skill is used for****
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name.
	//
	// example:
	//
	// name****
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The environment variables.
	EnvVarsShrink *string `json:"EnvVars,omitempty" xml:"EnvVars,omitempty"`
	// The icon parsing tag. This parameter is required when SkillIcon is specified.
	//
	// example:
	//
	// 21E9A5B273CB8EC0675*********
	IconETag *string `json:"IconETag,omitempty" xml:"IconETag,omitempty"`
	// The skill channel. Valid values:
	//
	// - ENTERPRISE: Enterprise Edition.
	//
	// - BUSINESS: Business Edition.
	//
	// example:
	//
	// BUSINESS
	SkillChannel *string `json:"SkillChannel,omitempty" xml:"SkillChannel,omitempty"`
	// The skill icon.
	SkillIcon *string `json:"SkillIcon,omitempty" xml:"SkillIcon,omitempty"`
	// The skill version.
	//
	// example:
	//
	// 0.0.1
	SkillVersion *string `json:"SkillVersion,omitempty" xml:"SkillVersion,omitempty"`
	// The slug identifier of the skill. This parameter is user-defined and must be unique within the tenant.
	//
	// example:
	//
	// find-skills****
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// The file parsing task key.
	//
	// example:
	//
	// E1CF3D69-529D-****
	TaskKey *string `json:"TaskKey,omitempty" xml:"TaskKey,omitempty"`
}

func (s CreateTenantSkillShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTenantSkillShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTenantSkillShrinkRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateTenantSkillShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTenantSkillShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateTenantSkillShrinkRequest) GetEnvVarsShrink() *string {
	return s.EnvVarsShrink
}

func (s *CreateTenantSkillShrinkRequest) GetIconETag() *string {
	return s.IconETag
}

func (s *CreateTenantSkillShrinkRequest) GetSkillChannel() *string {
	return s.SkillChannel
}

func (s *CreateTenantSkillShrinkRequest) GetSkillIcon() *string {
	return s.SkillIcon
}

func (s *CreateTenantSkillShrinkRequest) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *CreateTenantSkillShrinkRequest) GetSlug() *string {
	return s.Slug
}

func (s *CreateTenantSkillShrinkRequest) GetTaskKey() *string {
	return s.TaskKey
}

func (s *CreateTenantSkillShrinkRequest) SetApiKey(v string) *CreateTenantSkillShrinkRequest {
	s.ApiKey = &v
	return s
}

func (s *CreateTenantSkillShrinkRequest) SetDescription(v string) *CreateTenantSkillShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateTenantSkillShrinkRequest) SetDisplayName(v string) *CreateTenantSkillShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateTenantSkillShrinkRequest) SetEnvVarsShrink(v string) *CreateTenantSkillShrinkRequest {
	s.EnvVarsShrink = &v
	return s
}

func (s *CreateTenantSkillShrinkRequest) SetIconETag(v string) *CreateTenantSkillShrinkRequest {
	s.IconETag = &v
	return s
}

func (s *CreateTenantSkillShrinkRequest) SetSkillChannel(v string) *CreateTenantSkillShrinkRequest {
	s.SkillChannel = &v
	return s
}

func (s *CreateTenantSkillShrinkRequest) SetSkillIcon(v string) *CreateTenantSkillShrinkRequest {
	s.SkillIcon = &v
	return s
}

func (s *CreateTenantSkillShrinkRequest) SetSkillVersion(v string) *CreateTenantSkillShrinkRequest {
	s.SkillVersion = &v
	return s
}

func (s *CreateTenantSkillShrinkRequest) SetSlug(v string) *CreateTenantSkillShrinkRequest {
	s.Slug = &v
	return s
}

func (s *CreateTenantSkillShrinkRequest) SetTaskKey(v string) *CreateTenantSkillShrinkRequest {
	s.TaskKey = &v
	return s
}

func (s *CreateTenantSkillShrinkRequest) Validate() error {
	return dara.Validate(s)
}
