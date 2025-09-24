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
	AllowedMountAccessLevels []*string `json:"AllowedMountAccessLevels,omitempty" xml:"AllowedMountAccessLevels,omitempty" type:"Repeated"`
	ExpiresAt                *string   `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	IsSecureMode             *bool     `json:"IsSecureMode,omitempty" xml:"IsSecureMode,omitempty"`
	SharedAt                 *string   `json:"SharedAt,omitempty" xml:"SharedAt,omitempty"`
	SourceTenantId           *string   `json:"SourceTenantId,omitempty" xml:"SourceTenantId,omitempty"`
	SourceWorkspaceId        *string   `json:"SourceWorkspaceId,omitempty" xml:"SourceWorkspaceId,omitempty"`
	Status                   *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId                 *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId              *string   `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
