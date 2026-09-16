// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActiveVersionId(v string) *GetSkillResponseBody
	GetActiveVersionId() *string
	SetCategory(v string) *GetSkillResponseBody
	GetCategory() *string
	SetContent(v map[string]interface{}) *GetSkillResponseBody
	GetContent() map[string]interface{}
	SetCreatedAt(v string) *GetSkillResponseBody
	GetCreatedAt() *string
	SetDbtypes(v []*string) *GetSkillResponseBody
	GetDbtypes() []*string
	SetDescription(v string) *GetSkillResponseBody
	GetDescription() *string
	SetDisplayName(v string) *GetSkillResponseBody
	GetDisplayName() *string
	SetIcon(v string) *GetSkillResponseBody
	GetIcon() *string
	SetId(v string) *GetSkillResponseBody
	GetId() *string
	SetIsDeleted(v bool) *GetSkillResponseBody
	GetIsDeleted() *bool
	SetName(v string) *GetSkillResponseBody
	GetName() *string
	SetRequestId(v string) *GetSkillResponseBody
	GetRequestId() *string
	SetScope(v string) *GetSkillResponseBody
	GetScope() *string
	SetSkillType(v string) *GetSkillResponseBody
	GetSkillType() *string
	SetSlug(v string) *GetSkillResponseBody
	GetSlug() *string
	SetUpdatedAt(v string) *GetSkillResponseBody
	GetUpdatedAt() *string
	SetVersions(v []*GetSkillResponseBodyVersions) *GetSkillResponseBody
	GetVersions() []*GetSkillResponseBodyVersions
}

type GetSkillResponseBody struct {
	// The currently active version ID.
	//
	// example:
	//
	// version-example
	ActiveVersionId *string `json:"ActiveVersionId,omitempty" xml:"ActiveVersionId,omitempty"`
	// The Skill category.
	//
	// example:
	//
	// productivity
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The content.
	//
	// example:
	//
	// {"MySQL": "MySQL optimization guide...","PostgreSQL": "PostgreSQL optimization guide..."}
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-06-04T02:25:43Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The list of database types.
	Dbtypes []*string `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty" type:"Repeated"`
	// The Skill description, up to 1000 characters.
	//
	// example:
	//
	// SQL Review Expert: Comprehensively reviews SQL for security, performance, and compliance, identifies risks, and provides optimization suggestions. Activated immediately when a user submits SQL or asks about "SQL review", "SQL Review", "any risks", or "how to optimize"
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Skill display name.
	//
	// example:
	//
	// Example Skill
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The public HTTPS URL of the current icon. This value is empty if no icon is configured.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// https://example.com/skill-icon.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The unique identifier of the Skill.
	//
	// example:
	//
	// d1b7d639-f34e-44c7-8231-987da14d****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the Skill is deleted.
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// The Skill name. The name can contain only lowercase letters, digits, and hyphens.
	//
	// example:
	//
	// sql-optimization
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unique identifier of the request.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The visibility scope of the Skill.
	//
	// example:
	//
	// PRIVATE
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The Skill type.
	//
	// example:
	//
	// user
	SkillType *string `json:"SkillType,omitempty" xml:"SkillType,omitempty"`
	// The stable identifier of a private Skill.
	//
	// example:
	//
	// example-skill
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-02-04T21:14:45Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// The list of versions visible to the current principal.
	Versions []*GetSkillResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s GetSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillResponseBody) GetActiveVersionId() *string {
	return s.ActiveVersionId
}

func (s *GetSkillResponseBody) GetCategory() *string {
	return s.Category
}

func (s *GetSkillResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *GetSkillResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetSkillResponseBody) GetDbtypes() []*string {
	return s.Dbtypes
}

func (s *GetSkillResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetSkillResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetSkillResponseBody) GetIcon() *string {
	return s.Icon
}

func (s *GetSkillResponseBody) GetId() *string {
	return s.Id
}

func (s *GetSkillResponseBody) GetIsDeleted() *bool {
	return s.IsDeleted
}

func (s *GetSkillResponseBody) GetName() *string {
	return s.Name
}

func (s *GetSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillResponseBody) GetScope() *string {
	return s.Scope
}

func (s *GetSkillResponseBody) GetSkillType() *string {
	return s.SkillType
}

func (s *GetSkillResponseBody) GetSlug() *string {
	return s.Slug
}

func (s *GetSkillResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetSkillResponseBody) GetVersions() []*GetSkillResponseBodyVersions {
	return s.Versions
}

func (s *GetSkillResponseBody) SetActiveVersionId(v string) *GetSkillResponseBody {
	s.ActiveVersionId = &v
	return s
}

func (s *GetSkillResponseBody) SetCategory(v string) *GetSkillResponseBody {
	s.Category = &v
	return s
}

func (s *GetSkillResponseBody) SetContent(v map[string]interface{}) *GetSkillResponseBody {
	s.Content = v
	return s
}

func (s *GetSkillResponseBody) SetCreatedAt(v string) *GetSkillResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetSkillResponseBody) SetDbtypes(v []*string) *GetSkillResponseBody {
	s.Dbtypes = v
	return s
}

func (s *GetSkillResponseBody) SetDescription(v string) *GetSkillResponseBody {
	s.Description = &v
	return s
}

func (s *GetSkillResponseBody) SetDisplayName(v string) *GetSkillResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetSkillResponseBody) SetIcon(v string) *GetSkillResponseBody {
	s.Icon = &v
	return s
}

func (s *GetSkillResponseBody) SetId(v string) *GetSkillResponseBody {
	s.Id = &v
	return s
}

func (s *GetSkillResponseBody) SetIsDeleted(v bool) *GetSkillResponseBody {
	s.IsDeleted = &v
	return s
}

func (s *GetSkillResponseBody) SetName(v string) *GetSkillResponseBody {
	s.Name = &v
	return s
}

func (s *GetSkillResponseBody) SetRequestId(v string) *GetSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillResponseBody) SetScope(v string) *GetSkillResponseBody {
	s.Scope = &v
	return s
}

func (s *GetSkillResponseBody) SetSkillType(v string) *GetSkillResponseBody {
	s.SkillType = &v
	return s
}

func (s *GetSkillResponseBody) SetSlug(v string) *GetSkillResponseBody {
	s.Slug = &v
	return s
}

func (s *GetSkillResponseBody) SetUpdatedAt(v string) *GetSkillResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *GetSkillResponseBody) SetVersions(v []*GetSkillResponseBodyVersions) *GetSkillResponseBody {
	s.Versions = v
	return s
}

func (s *GetSkillResponseBody) Validate() error {
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

type GetSkillResponseBodyVersions struct {
	// The activation time of the Skill version.
	//
	// example:
	//
	// 2026-09-15T10:00:00Z
	ActivatedAt *string `json:"ActivatedAt,omitempty" xml:"ActivatedAt,omitempty"`
	// The creation time of the Skill version.
	//
	// example:
	//
	// 2026-09-15T10:00:00Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// Indicates whether the Skill requires a credential.
	CredentialRequired *bool `json:"CredentialRequired,omitempty" xml:"CredentialRequired,omitempty"`
	// The Skill version ID.
	//
	// example:
	//
	// version-example
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The Skill package size, in bytes.
	//
	// example:
	//
	// 1024
	PackageSize *int64 `json:"PackageSize,omitempty" xml:"PackageSize,omitempty"`
	// The reason for revoking the Skill version.
	//
	// example:
	//
	// Replaced by a newer version
	RevokeReason *string `json:"RevokeReason,omitempty" xml:"RevokeReason,omitempty"`
	// The revocation time of the Skill version.
	//
	// example:
	//
	// 2026-09-15T11:00:00Z
	RevokedAt *string `json:"RevokedAt,omitempty" xml:"RevokedAt,omitempty"`
	// The SHA-256 digest of the Skill package.
	//
	// example:
	//
	// 0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef
	Sha256 *string `json:"Sha256,omitempty" xml:"Sha256,omitempty"`
	// The ID of the Skill to which this version belongs.
	//
	// example:
	//
	// skill-example
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The Markdown content of the Skill.
	//
	// example:
	//
	// # Example Skill
	SkillMarkdown *string `json:"SkillMarkdown,omitempty" xml:"SkillMarkdown,omitempty"`
	// The status of the Skill version.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version number of the Skill.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetSkillResponseBodyVersions) String() string {
	return dara.Prettify(s)
}

func (s GetSkillResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *GetSkillResponseBodyVersions) GetActivatedAt() *string {
	return s.ActivatedAt
}

func (s *GetSkillResponseBodyVersions) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetSkillResponseBodyVersions) GetCredentialRequired() *bool {
	return s.CredentialRequired
}

func (s *GetSkillResponseBodyVersions) GetId() *string {
	return s.Id
}

func (s *GetSkillResponseBodyVersions) GetPackageSize() *int64 {
	return s.PackageSize
}

func (s *GetSkillResponseBodyVersions) GetRevokeReason() *string {
	return s.RevokeReason
}

func (s *GetSkillResponseBodyVersions) GetRevokedAt() *string {
	return s.RevokedAt
}

func (s *GetSkillResponseBodyVersions) GetSha256() *string {
	return s.Sha256
}

func (s *GetSkillResponseBodyVersions) GetSkillId() *string {
	return s.SkillId
}

func (s *GetSkillResponseBodyVersions) GetSkillMarkdown() *string {
	return s.SkillMarkdown
}

func (s *GetSkillResponseBodyVersions) GetStatus() *string {
	return s.Status
}

func (s *GetSkillResponseBodyVersions) GetVersion() *string {
	return s.Version
}

func (s *GetSkillResponseBodyVersions) SetActivatedAt(v string) *GetSkillResponseBodyVersions {
	s.ActivatedAt = &v
	return s
}

func (s *GetSkillResponseBodyVersions) SetCreatedAt(v string) *GetSkillResponseBodyVersions {
	s.CreatedAt = &v
	return s
}

func (s *GetSkillResponseBodyVersions) SetCredentialRequired(v bool) *GetSkillResponseBodyVersions {
	s.CredentialRequired = &v
	return s
}

func (s *GetSkillResponseBodyVersions) SetId(v string) *GetSkillResponseBodyVersions {
	s.Id = &v
	return s
}

func (s *GetSkillResponseBodyVersions) SetPackageSize(v int64) *GetSkillResponseBodyVersions {
	s.PackageSize = &v
	return s
}

func (s *GetSkillResponseBodyVersions) SetRevokeReason(v string) *GetSkillResponseBodyVersions {
	s.RevokeReason = &v
	return s
}

func (s *GetSkillResponseBodyVersions) SetRevokedAt(v string) *GetSkillResponseBodyVersions {
	s.RevokedAt = &v
	return s
}

func (s *GetSkillResponseBodyVersions) SetSha256(v string) *GetSkillResponseBodyVersions {
	s.Sha256 = &v
	return s
}

func (s *GetSkillResponseBodyVersions) SetSkillId(v string) *GetSkillResponseBodyVersions {
	s.SkillId = &v
	return s
}

func (s *GetSkillResponseBodyVersions) SetSkillMarkdown(v string) *GetSkillResponseBodyVersions {
	s.SkillMarkdown = &v
	return s
}

func (s *GetSkillResponseBodyVersions) SetStatus(v string) *GetSkillResponseBodyVersions {
	s.Status = &v
	return s
}

func (s *GetSkillResponseBodyVersions) SetVersion(v string) *GetSkillResponseBodyVersions {
	s.Version = &v
	return s
}

func (s *GetSkillResponseBodyVersions) Validate() error {
	return dara.Validate(s)
}
