// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallPolarClawSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *InstallPolarClawSkillRequest
	GetApplicationId() *string
	SetForce(v bool) *InstallPolarClawSkillRequest
	GetForce() *bool
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
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Specifies whether to force a reinstallation.
	//
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1.2.0
	SkillVersion *string `json:"SkillVersion,omitempty" xml:"SkillVersion,omitempty"`
	// The Skill identifier.
	//
	// example:
	//
	// alibacloud-rds-copilot
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// The source.
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

func (s *InstallPolarClawSkillRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *InstallPolarClawSkillRequest) GetForce() *bool {
	return s.Force
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

func (s *InstallPolarClawSkillRequest) SetApplicationId(v string) *InstallPolarClawSkillRequest {
	s.ApplicationId = &v
	return s
}

func (s *InstallPolarClawSkillRequest) SetForce(v bool) *InstallPolarClawSkillRequest {
	s.Force = &v
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
