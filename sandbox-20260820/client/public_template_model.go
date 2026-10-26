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
	SetResourceTags(v []*PublicTemplateResourceTags) *PublicTemplate
	GetResourceTags() []*PublicTemplateResourceTags
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
	// The time when the template was created, in UTC and RFC 3339 format.
	//
	// example:
	//
	// 2026-08-20T08:30:00Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The name of the template, which is unique within a team.
	//
	// example:
	//
	// python-data-analysis
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the Alibaba Cloud resource group to which the team belongs.
	//
	// example:
	//
	// rg-****
	ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	// The resource tags.
	ResourceTags []*PublicTemplateResourceTags `json:"resourceTags" xml:"resourceTags" type:"Repeated"`
	// The runtime configuration of the template, which determines the specifications and environment of the sandbox instances created from the template.
	RuntimeConfig *PublicTemplateRuntimeConfig `json:"runtimeConfig,omitempty" xml:"runtimeConfig,omitempty"`
	// The current status of the template.
	Status *PublicTemplateStatus `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the team to which the resource belongs, in UUID format.
	//
	// example:
	//
	// 5f4a2c18-****
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// The name of the team.
	//
	// example:
	//
	// sandbox-dev
	TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
	// The ID of the sandbox template.
	//
	// example:
	//
	// tm-8f3a2c7b5e14d806
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

func (s *PublicTemplate) GetResourceTags() []*PublicTemplateResourceTags {
	return s.ResourceTags
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

func (s *PublicTemplate) SetResourceTags(v []*PublicTemplateResourceTags) *PublicTemplate {
	s.ResourceTags = v
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
	if s.ResourceTags != nil {
		for _, item := range s.ResourceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type PublicTemplateResourceTags struct {
	// The tag key.
	//
	// example:
	//
	// project
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// demo
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s PublicTemplateResourceTags) String() string {
	return dara.Prettify(s)
}

func (s PublicTemplateResourceTags) GoString() string {
	return s.String()
}

func (s *PublicTemplateResourceTags) GetTagKey() *string {
	return s.TagKey
}

func (s *PublicTemplateResourceTags) GetTagValue() *string {
	return s.TagValue
}

func (s *PublicTemplateResourceTags) SetTagKey(v string) *PublicTemplateResourceTags {
	s.TagKey = &v
	return s
}

func (s *PublicTemplateResourceTags) SetTagValue(v string) *PublicTemplateResourceTags {
	s.TagValue = &v
	return s
}

func (s *PublicTemplateResourceTags) Validate() error {
	return dara.Validate(s)
}
