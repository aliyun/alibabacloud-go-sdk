// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogRevision(v int64) *CreateSkillResponseBody
	GetCatalogRevision() *int64
	SetContent(v map[string]interface{}) *CreateSkillResponseBody
	GetContent() map[string]interface{}
	SetCreatedAt(v string) *CreateSkillResponseBody
	GetCreatedAt() *string
	SetDbtypes(v []*string) *CreateSkillResponseBody
	GetDbtypes() []*string
	SetDescription(v string) *CreateSkillResponseBody
	GetDescription() *string
	SetId(v string) *CreateSkillResponseBody
	GetId() *string
	SetName(v string) *CreateSkillResponseBody
	GetName() *string
	SetRequestId(v string) *CreateSkillResponseBody
	GetRequestId() *string
	SetSkill(v *CreateSkillResponseBodySkill) *CreateSkillResponseBody
	GetSkill() *CreateSkillResponseBodySkill
	SetSkillType(v string) *CreateSkillResponseBody
	GetSkillType() *string
	SetVersion(v *CreateSkillResponseBodyVersion) *CreateSkillResponseBody
	GetVersion() *CreateSkillResponseBodyVersion
}

type CreateSkillResponseBody struct {
	// The Skill catalog revision number.
	//
	// example:
	//
	// 1
	CatalogRevision *int64 `json:"CatalogRevision,omitempty" xml:"CatalogRevision,omitempty"`
	// The content grouped by database type.
	//
	// example:
	//
	// {"MySQL": "MySQL optimization guide...","PostgreSQL": "PostgreSQL optimization guide..."}
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-02-04T21:14:45Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The list of database types.
	Dbtypes []*string `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty" type:"Repeated"`
	// The Skill description.
	//
	// example:
	//
	// SQL query optimization skill
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique identifier of the Skill.
	//
	// example:
	//
	// 82cf3d62-0add-47bd-869f-877131f7****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The Skill name.
	//
	// example:
	//
	// query-optimization
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unique request identifier.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The created Skill.
	Skill *CreateSkillResponseBodySkill `json:"Skill,omitempty" xml:"Skill,omitempty" type:"Struct"`
	// The Skill type.
	//
	// example:
	//
	// user
	SkillType *string `json:"SkillType,omitempty" xml:"SkillType,omitempty"`
	// The created Skill version.
	Version *CreateSkillResponseBodyVersion `json:"Version,omitempty" xml:"Version,omitempty" type:"Struct"`
}

func (s CreateSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillResponseBody) GetCatalogRevision() *int64 {
	return s.CatalogRevision
}

func (s *CreateSkillResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *CreateSkillResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateSkillResponseBody) GetDbtypes() []*string {
	return s.Dbtypes
}

func (s *CreateSkillResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateSkillResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateSkillResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSkillResponseBody) GetSkill() *CreateSkillResponseBodySkill {
	return s.Skill
}

func (s *CreateSkillResponseBody) GetSkillType() *string {
	return s.SkillType
}

func (s *CreateSkillResponseBody) GetVersion() *CreateSkillResponseBodyVersion {
	return s.Version
}

func (s *CreateSkillResponseBody) SetCatalogRevision(v int64) *CreateSkillResponseBody {
	s.CatalogRevision = &v
	return s
}

func (s *CreateSkillResponseBody) SetContent(v map[string]interface{}) *CreateSkillResponseBody {
	s.Content = v
	return s
}

func (s *CreateSkillResponseBody) SetCreatedAt(v string) *CreateSkillResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateSkillResponseBody) SetDbtypes(v []*string) *CreateSkillResponseBody {
	s.Dbtypes = v
	return s
}

func (s *CreateSkillResponseBody) SetDescription(v string) *CreateSkillResponseBody {
	s.Description = &v
	return s
}

func (s *CreateSkillResponseBody) SetId(v string) *CreateSkillResponseBody {
	s.Id = &v
	return s
}

func (s *CreateSkillResponseBody) SetName(v string) *CreateSkillResponseBody {
	s.Name = &v
	return s
}

func (s *CreateSkillResponseBody) SetRequestId(v string) *CreateSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillResponseBody) SetSkill(v *CreateSkillResponseBodySkill) *CreateSkillResponseBody {
	s.Skill = v
	return s
}

func (s *CreateSkillResponseBody) SetSkillType(v string) *CreateSkillResponseBody {
	s.SkillType = &v
	return s
}

func (s *CreateSkillResponseBody) SetVersion(v *CreateSkillResponseBodyVersion) *CreateSkillResponseBody {
	s.Version = v
	return s
}

func (s *CreateSkillResponseBody) Validate() error {
	if s.Skill != nil {
		if err := s.Skill.Validate(); err != nil {
			return err
		}
	}
	if s.Version != nil {
		if err := s.Version.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSkillResponseBodySkill struct {
	// The ID of the currently active version.
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
	// The Skill creation time.
	//
	// example:
	//
	// 2026-09-15T10:00:00Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The Skill description.
	//
	// example:
	//
	// An example ContextDB Skill
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
	// Skill ID
	//
	// example:
	//
	// skill-example
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the Skill is deleted.
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// The visibility scope of the Skill.
	//
	// example:
	//
	// PRIVATE
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The stable identifier of the Skill.
	//
	// example:
	//
	// example-skill
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// The Skill update time.
	//
	// example:
	//
	// 2026-09-15T10:00:00Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s CreateSkillResponseBodySkill) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillResponseBodySkill) GoString() string {
	return s.String()
}

func (s *CreateSkillResponseBodySkill) GetActiveVersionId() *string {
	return s.ActiveVersionId
}

func (s *CreateSkillResponseBodySkill) GetCategory() *string {
	return s.Category
}

func (s *CreateSkillResponseBodySkill) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateSkillResponseBodySkill) GetDescription() *string {
	return s.Description
}

func (s *CreateSkillResponseBodySkill) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateSkillResponseBodySkill) GetIcon() *string {
	return s.Icon
}

func (s *CreateSkillResponseBodySkill) GetId() *string {
	return s.Id
}

func (s *CreateSkillResponseBodySkill) GetIsDeleted() *bool {
	return s.IsDeleted
}

func (s *CreateSkillResponseBodySkill) GetScope() *string {
	return s.Scope
}

func (s *CreateSkillResponseBodySkill) GetSlug() *string {
	return s.Slug
}

func (s *CreateSkillResponseBodySkill) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateSkillResponseBodySkill) SetActiveVersionId(v string) *CreateSkillResponseBodySkill {
	s.ActiveVersionId = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetCategory(v string) *CreateSkillResponseBodySkill {
	s.Category = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetCreatedAt(v string) *CreateSkillResponseBodySkill {
	s.CreatedAt = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetDescription(v string) *CreateSkillResponseBodySkill {
	s.Description = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetDisplayName(v string) *CreateSkillResponseBodySkill {
	s.DisplayName = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetIcon(v string) *CreateSkillResponseBodySkill {
	s.Icon = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetId(v string) *CreateSkillResponseBodySkill {
	s.Id = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetIsDeleted(v bool) *CreateSkillResponseBodySkill {
	s.IsDeleted = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetScope(v string) *CreateSkillResponseBodySkill {
	s.Scope = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetSlug(v string) *CreateSkillResponseBodySkill {
	s.Slug = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetUpdatedAt(v string) *CreateSkillResponseBodySkill {
	s.UpdatedAt = &v
	return s
}

func (s *CreateSkillResponseBodySkill) Validate() error {
	return dara.Validate(s)
}

type CreateSkillResponseBodyVersion struct {
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
	// The revocation reason of the Skill version.
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
	// The ID of the parent Skill.
	//
	// example:
	//
	// skill-example
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The status of the Skill version.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Skill version number.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateSkillResponseBodyVersion) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillResponseBodyVersion) GoString() string {
	return s.String()
}

func (s *CreateSkillResponseBodyVersion) GetActivatedAt() *string {
	return s.ActivatedAt
}

func (s *CreateSkillResponseBodyVersion) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateSkillResponseBodyVersion) GetCredentialRequired() *bool {
	return s.CredentialRequired
}

func (s *CreateSkillResponseBodyVersion) GetId() *string {
	return s.Id
}

func (s *CreateSkillResponseBodyVersion) GetPackageSize() *int64 {
	return s.PackageSize
}

func (s *CreateSkillResponseBodyVersion) GetRevokeReason() *string {
	return s.RevokeReason
}

func (s *CreateSkillResponseBodyVersion) GetRevokedAt() *string {
	return s.RevokedAt
}

func (s *CreateSkillResponseBodyVersion) GetSha256() *string {
	return s.Sha256
}

func (s *CreateSkillResponseBodyVersion) GetSkillId() *string {
	return s.SkillId
}

func (s *CreateSkillResponseBodyVersion) GetStatus() *string {
	return s.Status
}

func (s *CreateSkillResponseBodyVersion) GetVersion() *string {
	return s.Version
}

func (s *CreateSkillResponseBodyVersion) SetActivatedAt(v string) *CreateSkillResponseBodyVersion {
	s.ActivatedAt = &v
	return s
}

func (s *CreateSkillResponseBodyVersion) SetCreatedAt(v string) *CreateSkillResponseBodyVersion {
	s.CreatedAt = &v
	return s
}

func (s *CreateSkillResponseBodyVersion) SetCredentialRequired(v bool) *CreateSkillResponseBodyVersion {
	s.CredentialRequired = &v
	return s
}

func (s *CreateSkillResponseBodyVersion) SetId(v string) *CreateSkillResponseBodyVersion {
	s.Id = &v
	return s
}

func (s *CreateSkillResponseBodyVersion) SetPackageSize(v int64) *CreateSkillResponseBodyVersion {
	s.PackageSize = &v
	return s
}

func (s *CreateSkillResponseBodyVersion) SetRevokeReason(v string) *CreateSkillResponseBodyVersion {
	s.RevokeReason = &v
	return s
}

func (s *CreateSkillResponseBodyVersion) SetRevokedAt(v string) *CreateSkillResponseBodyVersion {
	s.RevokedAt = &v
	return s
}

func (s *CreateSkillResponseBodyVersion) SetSha256(v string) *CreateSkillResponseBodyVersion {
	s.Sha256 = &v
	return s
}

func (s *CreateSkillResponseBodyVersion) SetSkillId(v string) *CreateSkillResponseBodyVersion {
	s.SkillId = &v
	return s
}

func (s *CreateSkillResponseBodyVersion) SetStatus(v string) *CreateSkillResponseBodyVersion {
	s.Status = &v
	return s
}

func (s *CreateSkillResponseBodyVersion) SetVersion(v string) *CreateSkillResponseBodyVersion {
	s.Version = &v
	return s
}

func (s *CreateSkillResponseBodyVersion) Validate() error {
	return dara.Validate(s)
}
