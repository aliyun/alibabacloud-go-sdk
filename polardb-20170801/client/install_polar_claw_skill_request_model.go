// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallPolarClawSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *InstallPolarClawSkillRequest
	GetAgentId() *string
	SetApplicationId(v string) *InstallPolarClawSkillRequest
	GetApplicationId() *string
	SetForce(v bool) *InstallPolarClawSkillRequest
	GetForce() *bool
	SetScope(v string) *InstallPolarClawSkillRequest
	GetScope() *string
	SetSkillVersion(v string) *InstallPolarClawSkillRequest
	GetSkillVersion() *string
	SetSlug(v string) *InstallPolarClawSkillRequest
	GetSlug() *string
	SetSource(v string) *InstallPolarClawSkillRequest
	GetSource() *string
	SetUrl(v string) *InstallPolarClawSkillRequest
	GetUrl() *string
}

type InstallPolarClawSkillRequest struct {
	// The target agent in the workspace. If not specified, the default agent is used.
	//
	// example:
	//
	// test1
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Specifies whether to force reinstallation.
	//
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The scope. Valid values: workspace (default) and global (shared directory). This parameter is mutually exclusive with AgentId.
	//
	// example:
	//
	// global
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1.2.0
	SkillVersion *string `json:"SkillVersion,omitempty" xml:"SkillVersion,omitempty"`
	// The skill identifier.
	//
	// example:
	//
	// alibacloud-rds-copilot
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// The source. Default value: polarhub.
	//
	// example:
	//
	// clawhub
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// URL
	//
	// example:
	//
	// https://example.com/skill.tar.gz
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s InstallPolarClawSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallPolarClawSkillRequest) GoString() string {
	return s.String()
}

func (s *InstallPolarClawSkillRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *InstallPolarClawSkillRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *InstallPolarClawSkillRequest) GetForce() *bool {
	return s.Force
}

func (s *InstallPolarClawSkillRequest) GetScope() *string {
	return s.Scope
}

func (s *InstallPolarClawSkillRequest) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *InstallPolarClawSkillRequest) GetSlug() *string {
	return s.Slug
}

func (s *InstallPolarClawSkillRequest) GetSource() *string {
	return s.Source
}

func (s *InstallPolarClawSkillRequest) GetUrl() *string {
	return s.Url
}

func (s *InstallPolarClawSkillRequest) SetAgentId(v string) *InstallPolarClawSkillRequest {
	s.AgentId = &v
	return s
}

func (s *InstallPolarClawSkillRequest) SetApplicationId(v string) *InstallPolarClawSkillRequest {
	s.ApplicationId = &v
	return s
}

func (s *InstallPolarClawSkillRequest) SetForce(v bool) *InstallPolarClawSkillRequest {
	s.Force = &v
	return s
}

func (s *InstallPolarClawSkillRequest) SetScope(v string) *InstallPolarClawSkillRequest {
	s.Scope = &v
	return s
}

func (s *InstallPolarClawSkillRequest) SetSkillVersion(v string) *InstallPolarClawSkillRequest {
	s.SkillVersion = &v
	return s
}

func (s *InstallPolarClawSkillRequest) SetSlug(v string) *InstallPolarClawSkillRequest {
	s.Slug = &v
	return s
}

func (s *InstallPolarClawSkillRequest) SetSource(v string) *InstallPolarClawSkillRequest {
	s.Source = &v
	return s
}

func (s *InstallPolarClawSkillRequest) SetUrl(v string) *InstallPolarClawSkillRequest {
	s.Url = &v
	return s
}

func (s *InstallPolarClawSkillRequest) Validate() error {
	return dara.Validate(s)
}
