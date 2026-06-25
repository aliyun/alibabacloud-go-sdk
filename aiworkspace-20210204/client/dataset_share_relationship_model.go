// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetShareRelationship interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedMountAccessLevels(v []*string) *DatasetShareRelationship
	GetAllowedMountAccessLevels() []*string
	SetExpiresAt(v string) *DatasetShareRelationship
	GetExpiresAt() *string
	SetExtra(v string) *DatasetShareRelationship
	GetExtra() *string
	SetIsSecureMode(v bool) *DatasetShareRelationship
	GetIsSecureMode() *bool
	SetSharedAt(v string) *DatasetShareRelationship
	GetSharedAt() *string
	SetSourceTenantId(v string) *DatasetShareRelationship
	GetSourceTenantId() *string
	SetSourceWorkspaceId(v string) *DatasetShareRelationship
	GetSourceWorkspaceId() *string
	SetStatus(v string) *DatasetShareRelationship
	GetStatus() *string
	SetTenantId(v string) *DatasetShareRelationship
	GetTenantId() *string
	SetWorkspaceId(v string) *DatasetShareRelationship
	GetWorkspaceId() *string
}

type DatasetShareRelationship struct {
	// The allowed permissions for the shared dataset. When a user accesses the shared dataset, their permissions are limited to this list. The default value is \\`["RO"]\\`.
	//
	// - RO: Read-only permission. The recipient can only read the dataset.
	//
	// - RW: Read and write permission. The recipient can read and modify the dataset.
	AllowedMountAccessLevels []*string `json:"AllowedMountAccessLevels,omitempty" xml:"AllowedMountAccessLevels,omitempty" type:"Repeated"`
	// The expiration time. The time is in ISO 8601 format.
	//
	// > If you do not specify this parameter, the sharing relationship never expires.
	//
	// example:
	//
	// 2026-08-27T12:23:58Z
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// Additional configurations for the sharing relationship. This parameter is a JSON string.
	//
	// - AllowExportModel: Specifies whether to allow the export of trained models.
	//
	// - AllowAccessDLCWebTerminal: Specifies whether to allow users to log on to the container in a DLC task.
	//
	// - AllowAccessDLCFullLog: Specifies whether to allow access to the full task logs.
	//
	// example:
	//
	// {"AllowExportModel":false,"AllowAccessDLCWebTerminal":false,"AllowAccessDLCFullLog":false}
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// Specifies whether to enable security protection for the shared dataset.
	//
	// example:
	//
	// true
	IsSecureMode *bool `json:"IsSecureMode,omitempty" xml:"IsSecureMode,omitempty"`
	// The time when the dataset was shared. The time is in ISO 8601 format.
	//
	// example:
	//
	// 2025-08-27T12:23:58Z
	SharedAt *string `json:"SharedAt,omitempty" xml:"SharedAt,omitempty"`
	// The ID of the tenant that owns the source dataset. The user who shares the dataset must be a workspace administrator or the root account.
	//
	// example:
	//
	// 148***************115
	SourceTenantId *string `json:"SourceTenantId,omitempty" xml:"SourceTenantId,omitempty"`
	// The ID of the workspace that contains the source dataset.
	//
	// example:
	//
	// 33**19
	SourceWorkspaceId *string `json:"SourceWorkspaceId,omitempty" xml:"SourceWorkspaceId,omitempty"`
	// The status of the sharing relationship.
	//
	// - ACTIVE: The sharing relationship is active. Complete dataset information is displayed only in this state.
	//
	// - EXPIRED: The sharing relationship has expired.
	//
	// - REVOKED: The sharing relationship was revoked by the sharer.
	//
	// - INVALID: The sharing relationship is invalid. This can happen if the source dataset is deleted.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the target tenant. This must be a root account ID.
	//
	// > This parameter is required when you set a sharing relationship.
	//
	// example:
	//
	// 153***************249
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the target workspace. This ID must be different from the source workspace ID.
	//
	// > This parameter is required when you set a sharing relationship.
	//
	// example:
	//
	// 42**2
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DatasetShareRelationship) String() string {
	return dara.Prettify(s)
}

func (s DatasetShareRelationship) GoString() string {
	return s.String()
}

func (s *DatasetShareRelationship) GetAllowedMountAccessLevels() []*string {
	return s.AllowedMountAccessLevels
}

func (s *DatasetShareRelationship) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *DatasetShareRelationship) GetExtra() *string {
	return s.Extra
}

func (s *DatasetShareRelationship) GetIsSecureMode() *bool {
	return s.IsSecureMode
}

func (s *DatasetShareRelationship) GetSharedAt() *string {
	return s.SharedAt
}

func (s *DatasetShareRelationship) GetSourceTenantId() *string {
	return s.SourceTenantId
}

func (s *DatasetShareRelationship) GetSourceWorkspaceId() *string {
	return s.SourceWorkspaceId
}

func (s *DatasetShareRelationship) GetStatus() *string {
	return s.Status
}

func (s *DatasetShareRelationship) GetTenantId() *string {
	return s.TenantId
}

func (s *DatasetShareRelationship) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DatasetShareRelationship) SetAllowedMountAccessLevels(v []*string) *DatasetShareRelationship {
	s.AllowedMountAccessLevels = v
	return s
}

func (s *DatasetShareRelationship) SetExpiresAt(v string) *DatasetShareRelationship {
	s.ExpiresAt = &v
	return s
}

func (s *DatasetShareRelationship) SetExtra(v string) *DatasetShareRelationship {
	s.Extra = &v
	return s
}

func (s *DatasetShareRelationship) SetIsSecureMode(v bool) *DatasetShareRelationship {
	s.IsSecureMode = &v
	return s
}

func (s *DatasetShareRelationship) SetSharedAt(v string) *DatasetShareRelationship {
	s.SharedAt = &v
	return s
}

func (s *DatasetShareRelationship) SetSourceTenantId(v string) *DatasetShareRelationship {
	s.SourceTenantId = &v
	return s
}

func (s *DatasetShareRelationship) SetSourceWorkspaceId(v string) *DatasetShareRelationship {
	s.SourceWorkspaceId = &v
	return s
}

func (s *DatasetShareRelationship) SetStatus(v string) *DatasetShareRelationship {
	s.Status = &v
	return s
}

func (s *DatasetShareRelationship) SetTenantId(v string) *DatasetShareRelationship {
	s.TenantId = &v
	return s
}

func (s *DatasetShareRelationship) SetWorkspaceId(v string) *DatasetShareRelationship {
	s.WorkspaceId = &v
	return s
}

func (s *DatasetShareRelationship) Validate() error {
	return dara.Validate(s)
}
