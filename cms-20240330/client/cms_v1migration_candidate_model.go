// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCmsV1MigrationCandidate interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *CmsV1MigrationCandidate
	GetGroupName() *string
	SetMigrationStatus(v string) *CmsV1MigrationCandidate
	GetMigrationStatus() *string
	SetObserveGroupId(v string) *CmsV1MigrationCandidate
	GetObserveGroupId() *string
	SetOriginGroupId(v string) *CmsV1MigrationCandidate
	GetOriginGroupId() *string
	SetRejectionReason(v string) *CmsV1MigrationCandidate
	GetRejectionReason() *string
	SetSelectable(v bool) *CmsV1MigrationCandidate
	GetSelectable() *bool
	SetSourceReadOnly(v bool) *CmsV1MigrationCandidate
	GetSourceReadOnly() *bool
	SetSourceType(v string) *CmsV1MigrationCandidate
	GetSourceType() *string
	SetTargetGroupName(v string) *CmsV1MigrationCandidate
	GetTargetGroupName() *string
}

type CmsV1MigrationCandidate struct {
	// The name of the source application group.
	//
	// example:
	//
	// order-service
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The current migration or takeover status of the source group. Valid values: NOT_MIGRATED: Not migrated. PROCESSING: In progress. MIGRATED: Migrated. If a historical task lacks proof of source read-only status, NEEDS_REVIEW is returned through rejectionReason.
	//
	// example:
	//
	// PROCESSING
	MigrationStatus *string `json:"migrationStatus,omitempty" xml:"migrationStatus,omitempty"`
	// The ID of the target observation group.
	//
	// example:
	//
	// og-0123456789abcdef
	ObserveGroupId *string `json:"observeGroupId,omitempty" xml:"observeGroupId,omitempty"`
	// The ID of the source CloudMonitor 1.0 application group.
	//
	// example:
	//
	// 247675311
	OriginGroupId *string `json:"originGroupId,omitempty" xml:"originGroupId,omitempty"`
	// The reason why the group cannot be selected.
	//
	// example:
	//
	// NEEDS_REVIEW
	RejectionReason *string `json:"rejectionReason,omitempty" xml:"rejectionReason,omitempty"`
	// Specifies whether the group can be selected for migration.
	//
	// example:
	//
	// false
	Selectable *bool `json:"selectable,omitempty" xml:"selectable,omitempty"`
	// Specifies whether the source 1.0 group is persistently read-only.
	//
	// example:
	//
	// false
	SourceReadOnly *bool `json:"sourceReadOnly,omitempty" xml:"sourceReadOnly,omitempty"`
	// The type of the source application group. Valid values: custom: Custom group. tag: Tag-based group. regMgr: Resource group-based group.
	//
	// example:
	//
	// custom
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The name of the target observation group.
	//
	// example:
	//
	// order-service
	TargetGroupName *string `json:"targetGroupName,omitempty" xml:"targetGroupName,omitempty"`
}

func (s CmsV1MigrationCandidate) String() string {
	return dara.Prettify(s)
}

func (s CmsV1MigrationCandidate) GoString() string {
	return s.String()
}

func (s *CmsV1MigrationCandidate) GetGroupName() *string {
	return s.GroupName
}

func (s *CmsV1MigrationCandidate) GetMigrationStatus() *string {
	return s.MigrationStatus
}

func (s *CmsV1MigrationCandidate) GetObserveGroupId() *string {
	return s.ObserveGroupId
}

func (s *CmsV1MigrationCandidate) GetOriginGroupId() *string {
	return s.OriginGroupId
}

func (s *CmsV1MigrationCandidate) GetRejectionReason() *string {
	return s.RejectionReason
}

func (s *CmsV1MigrationCandidate) GetSelectable() *bool {
	return s.Selectable
}

func (s *CmsV1MigrationCandidate) GetSourceReadOnly() *bool {
	return s.SourceReadOnly
}

func (s *CmsV1MigrationCandidate) GetSourceType() *string {
	return s.SourceType
}

func (s *CmsV1MigrationCandidate) GetTargetGroupName() *string {
	return s.TargetGroupName
}

func (s *CmsV1MigrationCandidate) SetGroupName(v string) *CmsV1MigrationCandidate {
	s.GroupName = &v
	return s
}

func (s *CmsV1MigrationCandidate) SetMigrationStatus(v string) *CmsV1MigrationCandidate {
	s.MigrationStatus = &v
	return s
}

func (s *CmsV1MigrationCandidate) SetObserveGroupId(v string) *CmsV1MigrationCandidate {
	s.ObserveGroupId = &v
	return s
}

func (s *CmsV1MigrationCandidate) SetOriginGroupId(v string) *CmsV1MigrationCandidate {
	s.OriginGroupId = &v
	return s
}

func (s *CmsV1MigrationCandidate) SetRejectionReason(v string) *CmsV1MigrationCandidate {
	s.RejectionReason = &v
	return s
}

func (s *CmsV1MigrationCandidate) SetSelectable(v bool) *CmsV1MigrationCandidate {
	s.Selectable = &v
	return s
}

func (s *CmsV1MigrationCandidate) SetSourceReadOnly(v bool) *CmsV1MigrationCandidate {
	s.SourceReadOnly = &v
	return s
}

func (s *CmsV1MigrationCandidate) SetSourceType(v string) *CmsV1MigrationCandidate {
	s.SourceType = &v
	return s
}

func (s *CmsV1MigrationCandidate) SetTargetGroupName(v string) *CmsV1MigrationCandidate {
	s.TargetGroupName = &v
	return s
}

func (s *CmsV1MigrationCandidate) Validate() error {
	return dara.Validate(s)
}
