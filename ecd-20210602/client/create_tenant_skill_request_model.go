// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTenantSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateTenantSkillRequest
	GetApiKey() *string
	SetDescription(v string) *CreateTenantSkillRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateTenantSkillRequest
	GetDisplayName() *string
	SetEnvVars(v map[string]*string) *CreateTenantSkillRequest
	GetEnvVars() map[string]*string
	SetIconETag(v string) *CreateTenantSkillRequest
	GetIconETag() *string
	SetSkillChannel(v string) *CreateTenantSkillRequest
	GetSkillChannel() *string
	SetSkillIcon(v string) *CreateTenantSkillRequest
	GetSkillIcon() *string
	SetSkillVersion(v string) *CreateTenantSkillRequest
	GetSkillVersion() *string
	SetSlug(v string) *CreateTenantSkillRequest
	GetSlug() *string
	SetTaskKey(v string) *CreateTenantSkillRequest
	GetTaskKey() *string
}

type CreateTenantSkillRequest struct {
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
	EnvVars map[string]*string `json:"EnvVars,omitempty" xml:"EnvVars,omitempty"`
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

func (s CreateTenantSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTenantSkillRequest) GoString() string {
	return s.String()
}

func (s *CreateTenantSkillRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateTenantSkillRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTenantSkillRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateTenantSkillRequest) GetEnvVars() map[string]*string {
	return s.EnvVars
}

func (s *CreateTenantSkillRequest) GetIconETag() *string {
	return s.IconETag
}

func (s *CreateTenantSkillRequest) GetSkillChannel() *string {
	return s.SkillChannel
}

func (s *CreateTenantSkillRequest) GetSkillIcon() *string {
	return s.SkillIcon
}

func (s *CreateTenantSkillRequest) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *CreateTenantSkillRequest) GetSlug() *string {
	return s.Slug
}

func (s *CreateTenantSkillRequest) GetTaskKey() *string {
	return s.TaskKey
}

func (s *CreateTenantSkillRequest) SetApiKey(v string) *CreateTenantSkillRequest {
	s.ApiKey = &v
	return s
}

func (s *CreateTenantSkillRequest) SetDescription(v string) *CreateTenantSkillRequest {
	s.Description = &v
	return s
}

func (s *CreateTenantSkillRequest) SetDisplayName(v string) *CreateTenantSkillRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateTenantSkillRequest) SetEnvVars(v map[string]*string) *CreateTenantSkillRequest {
	s.EnvVars = v
	return s
}

func (s *CreateTenantSkillRequest) SetIconETag(v string) *CreateTenantSkillRequest {
	s.IconETag = &v
	return s
}

func (s *CreateTenantSkillRequest) SetSkillChannel(v string) *CreateTenantSkillRequest {
	s.SkillChannel = &v
	return s
}

func (s *CreateTenantSkillRequest) SetSkillIcon(v string) *CreateTenantSkillRequest {
	s.SkillIcon = &v
	return s
}

func (s *CreateTenantSkillRequest) SetSkillVersion(v string) *CreateTenantSkillRequest {
	s.SkillVersion = &v
	return s
}

func (s *CreateTenantSkillRequest) SetSlug(v string) *CreateTenantSkillRequest {
	s.Slug = &v
	return s
}

func (s *CreateTenantSkillRequest) SetTaskKey(v string) *CreateTenantSkillRequest {
	s.TaskKey = &v
	return s
}

func (s *CreateTenantSkillRequest) Validate() error {
	return dara.Validate(s)
}
