// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEvaluatorSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetEvaluatorSkillResponseBody
	GetRequestId() *string
	SetSkill(v *GetEvaluatorSkillResponseBodySkill) *GetEvaluatorSkillResponseBody
	GetSkill() *GetEvaluatorSkillResponseBodySkill
}

type GetEvaluatorSkillResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The skill details.
	//
	// example:
	//
	// {"skillName":"trace_context_loader","enable":true,"currentVersion":"1782816000000"}
	Skill *GetEvaluatorSkillResponseBodySkill `json:"skill,omitempty" xml:"skill,omitempty" type:"Struct"`
}

func (s GetEvaluatorSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluatorSkillResponseBody) GoString() string {
	return s.String()
}

func (s *GetEvaluatorSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEvaluatorSkillResponseBody) GetSkill() *GetEvaluatorSkillResponseBodySkill {
	return s.Skill
}

func (s *GetEvaluatorSkillResponseBody) SetRequestId(v string) *GetEvaluatorSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEvaluatorSkillResponseBody) SetSkill(v *GetEvaluatorSkillResponseBodySkill) *GetEvaluatorSkillResponseBody {
	s.Skill = v
	return s
}

func (s *GetEvaluatorSkillResponseBody) Validate() error {
	if s.Skill != nil {
		if err := s.Skill.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEvaluatorSkillResponseBodySkill struct {
	// The time when the skill was created. This value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1782816000
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The current version.
	//
	// example:
	//
	// 1782816000000
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The skill description.
	//
	// example:
	//
	// 读取链路上下文辅助评估
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name.
	//
	// example:
	//
	// Trace 上下文读取
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Indicates whether the skill is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The list of skill files.
	//
	// example:
	//
	// [{"name":"SKILL.md","content":"# Trace Context Loader","remark":"主技能说明"}]
	Files []*GetEvaluatorSkillResponseBodySkillFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The latest version.
	//
	// example:
	//
	// 1782816000000
	LatestVersion *string `json:"latestVersion,omitempty" xml:"latestVersion,omitempty"`
	// The skill name.
	//
	// example:
	//
	// trace_context_loader
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// The time when the skill was last updated. This value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1782816600
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The list of skill versions.
	//
	// example:
	//
	// [{"version":"1782816000000","versionDescription":"首次发布版本"}]
	Versions []*GetEvaluatorSkillResponseBodySkillVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s GetEvaluatorSkillResponseBodySkill) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluatorSkillResponseBodySkill) GoString() string {
	return s.String()
}

func (s *GetEvaluatorSkillResponseBodySkill) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *GetEvaluatorSkillResponseBodySkill) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *GetEvaluatorSkillResponseBodySkill) GetDescription() *string {
	return s.Description
}

func (s *GetEvaluatorSkillResponseBodySkill) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetEvaluatorSkillResponseBodySkill) GetEnable() *bool {
	return s.Enable
}

func (s *GetEvaluatorSkillResponseBodySkill) GetFiles() []*GetEvaluatorSkillResponseBodySkillFiles {
	return s.Files
}

func (s *GetEvaluatorSkillResponseBodySkill) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *GetEvaluatorSkillResponseBodySkill) GetSkillName() *string {
	return s.SkillName
}

func (s *GetEvaluatorSkillResponseBodySkill) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *GetEvaluatorSkillResponseBodySkill) GetVersions() []*GetEvaluatorSkillResponseBodySkillVersions {
	return s.Versions
}

func (s *GetEvaluatorSkillResponseBodySkill) SetCreatedAt(v int64) *GetEvaluatorSkillResponseBodySkill {
	s.CreatedAt = &v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkill) SetCurrentVersion(v string) *GetEvaluatorSkillResponseBodySkill {
	s.CurrentVersion = &v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkill) SetDescription(v string) *GetEvaluatorSkillResponseBodySkill {
	s.Description = &v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkill) SetDisplayName(v string) *GetEvaluatorSkillResponseBodySkill {
	s.DisplayName = &v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkill) SetEnable(v bool) *GetEvaluatorSkillResponseBodySkill {
	s.Enable = &v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkill) SetFiles(v []*GetEvaluatorSkillResponseBodySkillFiles) *GetEvaluatorSkillResponseBodySkill {
	s.Files = v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkill) SetLatestVersion(v string) *GetEvaluatorSkillResponseBodySkill {
	s.LatestVersion = &v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkill) SetSkillName(v string) *GetEvaluatorSkillResponseBodySkill {
	s.SkillName = &v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkill) SetUpdatedAt(v int64) *GetEvaluatorSkillResponseBodySkill {
	s.UpdatedAt = &v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkill) SetVersions(v []*GetEvaluatorSkillResponseBodySkillVersions) *GetEvaluatorSkillResponseBodySkill {
	s.Versions = v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkill) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEvaluatorSkillResponseBodySkillFiles struct {
	// The file content.
	//
	// example:
	//
	// # Trace Context Loader
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The file name.
	//
	// example:
	//
	// SKILL.md
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The file remarks.
	//
	// example:
	//
	// 主技能说明
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s GetEvaluatorSkillResponseBodySkillFiles) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluatorSkillResponseBodySkillFiles) GoString() string {
	return s.String()
}

func (s *GetEvaluatorSkillResponseBodySkillFiles) GetContent() *string {
	return s.Content
}

func (s *GetEvaluatorSkillResponseBodySkillFiles) GetName() *string {
	return s.Name
}

func (s *GetEvaluatorSkillResponseBodySkillFiles) GetRemark() *string {
	return s.Remark
}

func (s *GetEvaluatorSkillResponseBodySkillFiles) SetContent(v string) *GetEvaluatorSkillResponseBodySkillFiles {
	s.Content = &v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkillFiles) SetName(v string) *GetEvaluatorSkillResponseBodySkillFiles {
	s.Name = &v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkillFiles) SetRemark(v string) *GetEvaluatorSkillResponseBodySkillFiles {
	s.Remark = &v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkillFiles) Validate() error {
	return dara.Validate(s)
}

type GetEvaluatorSkillResponseBodySkillVersions struct {
	// The time when the version was created. This value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1782816000
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1782816000000
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The version description.
	//
	// example:
	//
	// 首次发布版本
	VersionDescription *string `json:"versionDescription,omitempty" xml:"versionDescription,omitempty"`
}

func (s GetEvaluatorSkillResponseBodySkillVersions) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluatorSkillResponseBodySkillVersions) GoString() string {
	return s.String()
}

func (s *GetEvaluatorSkillResponseBodySkillVersions) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *GetEvaluatorSkillResponseBodySkillVersions) GetVersion() *string {
	return s.Version
}

func (s *GetEvaluatorSkillResponseBodySkillVersions) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *GetEvaluatorSkillResponseBodySkillVersions) SetCreatedAt(v int64) *GetEvaluatorSkillResponseBodySkillVersions {
	s.CreatedAt = &v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkillVersions) SetVersion(v string) *GetEvaluatorSkillResponseBodySkillVersions {
	s.Version = &v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkillVersions) SetVersionDescription(v string) *GetEvaluatorSkillResponseBodySkillVersions {
	s.VersionDescription = &v
	return s
}

func (s *GetEvaluatorSkillResponseBodySkillVersions) Validate() error {
	return dara.Validate(s)
}
