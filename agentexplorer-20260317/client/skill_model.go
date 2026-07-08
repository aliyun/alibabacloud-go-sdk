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
	SetCategoryNameEn(v string) *Skill
	GetCategoryNameEn() *string
	SetCreatedAt(v string) *Skill
	GetCreatedAt() *string
	SetDescription(v string) *Skill
	GetDescription() *string
	SetDescriptionEn(v string) *Skill
	GetDescriptionEn() *string
	SetDisplayName(v string) *Skill
	GetDisplayName() *string
	SetGithubPath(v string) *Skill
	GetGithubPath() *string
	SetInstallCount(v int32) *Skill
	GetInstallCount() *int32
	SetLikeCount(v int32) *Skill
	GetLikeCount() *int32
	SetNameEn(v string) *Skill
	GetNameEn() *string
	SetSkillName(v string) *Skill
	GetSkillName() *string
	SetSubCategoryCode(v string) *Skill
	GetSubCategoryCode() *string
	SetSubCategoryName(v string) *Skill
	GetSubCategoryName() *string
	SetSubCategoryNameEn(v string) *Skill
	GetSubCategoryNameEn() *string
	SetUpdatedAt(v string) *Skill
	GetUpdatedAt() *string
}

type Skill struct {
	// The primary category code.
	//
	// example:
	//
	// compute
	CategoryCode *string `json:"categoryCode,omitempty" xml:"categoryCode,omitempty"`
	// The primary category name.
	//
	// example:
	//
	// 计算
	CategoryName   *string `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	CategoryNameEn *string `json:"categoryNameEn,omitempty" xml:"categoryNameEn,omitempty"`
	// The time when the Agent Skill was created.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The description of the Agent Skill.
	//
	// example:
	//
	// ECS 实例管理
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	DescriptionEn *string `json:"descriptionEn,omitempty" xml:"descriptionEn,omitempty"`
	// The display name of the Agent Skill.
	//
	// example:
	//
	// ECS 实例管理
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	GithubPath  *string `json:"githubPath,omitempty" xml:"githubPath,omitempty"`
	// The number of installations.
	//
	// example:
	//
	// 1024
	InstallCount *int32 `json:"installCount,omitempty" xml:"installCount,omitempty"`
	// The number of likes.
	//
	// example:
	//
	// 128
	LikeCount *int32  `json:"likeCount,omitempty" xml:"likeCount,omitempty"`
	NameEn    *string `json:"nameEn,omitempty" xml:"nameEn,omitempty"`
	// The English name of the Agent Skill, which serves as a unique identifier.
	//
	// example:
	//
	// deploy-to-vercel
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// The secondary category code.
	//
	// example:
	//
	// ecs
	SubCategoryCode *string `json:"subCategoryCode,omitempty" xml:"subCategoryCode,omitempty"`
	// The secondary category name.
	//
	// example:
	//
	// 弹性计算
	SubCategoryName   *string `json:"subCategoryName,omitempty" xml:"subCategoryName,omitempty"`
	SubCategoryNameEn *string `json:"subCategoryNameEn,omitempty" xml:"subCategoryNameEn,omitempty"`
	// The time when the Agent Skill was last updated.
	//
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

func (s *Skill) GetCategoryNameEn() *string {
	return s.CategoryNameEn
}

func (s *Skill) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Skill) GetDescription() *string {
	return s.Description
}

func (s *Skill) GetDescriptionEn() *string {
	return s.DescriptionEn
}

func (s *Skill) GetDisplayName() *string {
	return s.DisplayName
}

func (s *Skill) GetGithubPath() *string {
	return s.GithubPath
}

func (s *Skill) GetInstallCount() *int32 {
	return s.InstallCount
}

func (s *Skill) GetLikeCount() *int32 {
	return s.LikeCount
}

func (s *Skill) GetNameEn() *string {
	return s.NameEn
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

func (s *Skill) GetSubCategoryNameEn() *string {
	return s.SubCategoryNameEn
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

func (s *Skill) SetCategoryNameEn(v string) *Skill {
	s.CategoryNameEn = &v
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

func (s *Skill) SetDescriptionEn(v string) *Skill {
	s.DescriptionEn = &v
	return s
}

func (s *Skill) SetDisplayName(v string) *Skill {
	s.DisplayName = &v
	return s
}

func (s *Skill) SetGithubPath(v string) *Skill {
	s.GithubPath = &v
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

func (s *Skill) SetNameEn(v string) *Skill {
	s.NameEn = &v
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

func (s *Skill) SetSubCategoryNameEn(v string) *Skill {
	s.SubCategoryNameEn = &v
	return s
}

func (s *Skill) SetUpdatedAt(v string) *Skill {
	s.UpdatedAt = &v
	return s
}

func (s *Skill) Validate() error {
	return dara.Validate(s)
}
