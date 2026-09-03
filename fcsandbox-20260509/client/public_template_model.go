// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *PublicTemplate
	GetCreatedTime() *string
	SetName(v string) *PublicTemplate
	GetName() *string
	SetResourceGroupID(v string) *PublicTemplate
	GetResourceGroupID() *string
	SetRuntimeConfig(v *PublicTemplateRuntimeConfig) *PublicTemplate
	GetRuntimeConfig() *PublicTemplateRuntimeConfig
	SetStatus(v *PublicTemplateStatus) *PublicTemplate
	GetStatus() *PublicTemplateStatus
	SetTeamID(v string) *PublicTemplate
	GetTeamID() *string
	SetTeamName(v string) *PublicTemplate
	GetTeamName() *string
	SetTemplateID(v string) *PublicTemplate
	GetTemplateID() *string
}

type PublicTemplate struct {
	// The creation time.
	//
	// example:
	//
	// 2026-08-28T12:00:00.000Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The template name.
	//
	// example:
	//
	// my-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmz7h4ocksp5y
	ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	// The template runtime configuration.
	RuntimeConfig *PublicTemplateRuntimeConfig `json:"runtimeConfig,omitempty" xml:"runtimeConfig,omitempty"`
	// The template build status.
	Status *PublicTemplateStatus `json:"status,omitempty" xml:"status,omitempty"`
	// The unique identifier of the team.
	//
	// example:
	//
	// 70d1c834-0383-58d8-97ac-5336eb91abcd
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// The team name.
	//
	// example:
	//
	// my-team
	TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
	// The template ID.
	//
	// example:
	//
	// w0aipmi0rvn5xqdnsihg
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
}

func (s PublicTemplate) String() string {
	return dara.Prettify(s)
}

func (s PublicTemplate) GoString() string {
	return s.String()
}

func (s *PublicTemplate) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *PublicTemplate) GetName() *string {
	return s.Name
}

func (s *PublicTemplate) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *PublicTemplate) GetRuntimeConfig() *PublicTemplateRuntimeConfig {
	return s.RuntimeConfig
}

func (s *PublicTemplate) GetStatus() *PublicTemplateStatus {
	return s.Status
}

func (s *PublicTemplate) GetTeamID() *string {
	return s.TeamID
}

func (s *PublicTemplate) GetTeamName() *string {
	return s.TeamName
}

func (s *PublicTemplate) GetTemplateID() *string {
	return s.TemplateID
}

func (s *PublicTemplate) SetCreatedTime(v string) *PublicTemplate {
	s.CreatedTime = &v
	return s
}

func (s *PublicTemplate) SetName(v string) *PublicTemplate {
	s.Name = &v
	return s
}

func (s *PublicTemplate) SetResourceGroupID(v string) *PublicTemplate {
	s.ResourceGroupID = &v
	return s
}

func (s *PublicTemplate) SetRuntimeConfig(v *PublicTemplateRuntimeConfig) *PublicTemplate {
	s.RuntimeConfig = v
	return s
}

func (s *PublicTemplate) SetStatus(v *PublicTemplateStatus) *PublicTemplate {
	s.Status = v
	return s
}

func (s *PublicTemplate) SetTeamID(v string) *PublicTemplate {
	s.TeamID = &v
	return s
}

func (s *PublicTemplate) SetTeamName(v string) *PublicTemplate {
	s.TeamName = &v
	return s
}

func (s *PublicTemplate) SetTemplateID(v string) *PublicTemplate {
	s.TemplateID = &v
	return s
}

func (s *PublicTemplate) Validate() error {
	if s.RuntimeConfig != nil {
		if err := s.RuntimeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}
