// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkillInfoDO interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *SkillInfoDO
	GetCategory() *string
	SetCompatibility(v string) *SkillInfoDO
	GetCompatibility() *string
	SetDescription(v string) *SkillInfoDO
	GetDescription() *string
	SetDisplayName(v string) *SkillInfoDO
	GetDisplayName() *string
	SetInstallCount(v string) *SkillInfoDO
	GetInstallCount() *string
	SetName(v string) *SkillInfoDO
	GetName() *string
	SetSource(v string) *SkillInfoDO
	GetSource() *string
	SetSourceType(v string) *SkillInfoDO
	GetSourceType() *string
	SetTags(v string) *SkillInfoDO
	GetTags() *string
	SetUpdatedAt(v string) *SkillInfoDO
	GetUpdatedAt() *string
	SetVersion(v string) *SkillInfoDO
	GetVersion() *string
}

type SkillInfoDO struct {
	Category      *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Compatibility *string `json:"Compatibility,omitempty" xml:"Compatibility,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	InstallCount  *string `json:"InstallCount,omitempty" xml:"InstallCount,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceType    *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Tags          *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UpdatedAt     *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s SkillInfoDO) String() string {
	return dara.Prettify(s)
}

func (s SkillInfoDO) GoString() string {
	return s.String()
}

func (s *SkillInfoDO) GetCategory() *string {
	return s.Category
}

func (s *SkillInfoDO) GetCompatibility() *string {
	return s.Compatibility
}

func (s *SkillInfoDO) GetDescription() *string {
	return s.Description
}

func (s *SkillInfoDO) GetDisplayName() *string {
	return s.DisplayName
}

func (s *SkillInfoDO) GetInstallCount() *string {
	return s.InstallCount
}

func (s *SkillInfoDO) GetName() *string {
	return s.Name
}

func (s *SkillInfoDO) GetSource() *string {
	return s.Source
}

func (s *SkillInfoDO) GetSourceType() *string {
	return s.SourceType
}

func (s *SkillInfoDO) GetTags() *string {
	return s.Tags
}

func (s *SkillInfoDO) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *SkillInfoDO) GetVersion() *string {
	return s.Version
}

func (s *SkillInfoDO) SetCategory(v string) *SkillInfoDO {
	s.Category = &v
	return s
}

func (s *SkillInfoDO) SetCompatibility(v string) *SkillInfoDO {
	s.Compatibility = &v
	return s
}

func (s *SkillInfoDO) SetDescription(v string) *SkillInfoDO {
	s.Description = &v
	return s
}

func (s *SkillInfoDO) SetDisplayName(v string) *SkillInfoDO {
	s.DisplayName = &v
	return s
}

func (s *SkillInfoDO) SetInstallCount(v string) *SkillInfoDO {
	s.InstallCount = &v
	return s
}

func (s *SkillInfoDO) SetName(v string) *SkillInfoDO {
	s.Name = &v
	return s
}

func (s *SkillInfoDO) SetSource(v string) *SkillInfoDO {
	s.Source = &v
	return s
}

func (s *SkillInfoDO) SetSourceType(v string) *SkillInfoDO {
	s.SourceType = &v
	return s
}

func (s *SkillInfoDO) SetTags(v string) *SkillInfoDO {
	s.Tags = &v
	return s
}

func (s *SkillInfoDO) SetUpdatedAt(v string) *SkillInfoDO {
	s.UpdatedAt = &v
	return s
}

func (s *SkillInfoDO) SetVersion(v string) *SkillInfoDO {
	s.Version = &v
	return s
}

func (s *SkillInfoDO) Validate() error {
	return dara.Validate(s)
}
