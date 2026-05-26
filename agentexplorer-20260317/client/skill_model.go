// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkill interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryCode(v string) *Skill
	GetCategoryCode() *string
	SetCategoryName(v string) *Skill
	GetCategoryName() *string
	SetCreatedAt(v string) *Skill
	GetCreatedAt() *string
	SetDescription(v string) *Skill
	GetDescription() *string
	SetDisplayName(v string) *Skill
	GetDisplayName() *string
	SetInstallCount(v int32) *Skill
	GetInstallCount() *int32
	SetLikeCount(v int32) *Skill
	GetLikeCount() *int32
	SetSkillName(v string) *Skill
	GetSkillName() *string
	SetSubCategoryCode(v string) *Skill
	GetSubCategoryCode() *string
	SetSubCategoryName(v string) *Skill
	GetSubCategoryName() *string
	SetUpdatedAt(v string) *Skill
	GetUpdatedAt() *string
}

type Skill struct {
	// example:
	//
	// compute
	CategoryCode *string `json:"categoryCode,omitempty" xml:"categoryCode,omitempty"`
	CategoryName *string `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreatedAt   *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 1024
	InstallCount *int32 `json:"installCount,omitempty" xml:"installCount,omitempty"`
	// example:
	//
	// 128
	LikeCount *int32 `json:"likeCount,omitempty" xml:"likeCount,omitempty"`
	// example:
	//
	// deploy-to-vercel
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// example:
	//
	// ecs
	SubCategoryCode *string `json:"subCategoryCode,omitempty" xml:"subCategoryCode,omitempty"`
	SubCategoryName *string `json:"subCategoryName,omitempty" xml:"subCategoryName,omitempty"`
	// example:
	//
	// 2026-03-17T00:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s Skill) String() string {
	return dara.Prettify(s)
}

func (s Skill) GoString() string {
	return s.String()
}

func (s *Skill) GetCategoryCode() *string {
	return s.CategoryCode
}

func (s *Skill) GetCategoryName() *string {
	return s.CategoryName
}

func (s *Skill) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Skill) GetDescription() *string {
	return s.Description
}

func (s *Skill) GetDisplayName() *string {
	return s.DisplayName
}

func (s *Skill) GetInstallCount() *int32 {
	return s.InstallCount
}

func (s *Skill) GetLikeCount() *int32 {
	return s.LikeCount
}

func (s *Skill) GetSkillName() *string {
	return s.SkillName
}

func (s *Skill) GetSubCategoryCode() *string {
	return s.SubCategoryCode
}

func (s *Skill) GetSubCategoryName() *string {
	return s.SubCategoryName
}

func (s *Skill) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *Skill) SetCategoryCode(v string) *Skill {
	s.CategoryCode = &v
	return s
}

func (s *Skill) SetCategoryName(v string) *Skill {
	s.CategoryName = &v
	return s
}

func (s *Skill) SetCreatedAt(v string) *Skill {
	s.CreatedAt = &v
	return s
}

func (s *Skill) SetDescription(v string) *Skill {
	s.Description = &v
	return s
}

func (s *Skill) SetDisplayName(v string) *Skill {
	s.DisplayName = &v
	return s
}

func (s *Skill) SetInstallCount(v int32) *Skill {
	s.InstallCount = &v
	return s
}

func (s *Skill) SetLikeCount(v int32) *Skill {
	s.LikeCount = &v
	return s
}

func (s *Skill) SetSkillName(v string) *Skill {
	s.SkillName = &v
	return s
}

func (s *Skill) SetSubCategoryCode(v string) *Skill {
	s.SubCategoryCode = &v
	return s
}

func (s *Skill) SetSubCategoryName(v string) *Skill {
	s.SubCategoryName = &v
	return s
}

func (s *Skill) SetUpdatedAt(v string) *Skill {
	s.UpdatedAt = &v
	return s
}

func (s *Skill) Validate() error {
	return dara.Validate(s)
}
