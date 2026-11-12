// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCmsV1MigrationPreview interface {
	dara.Model
	String() string
	GoString() string
	SetChecks(v []*CmsV1MigrationCheck) *CmsV1MigrationPreview
	GetChecks() []*CmsV1MigrationCheck
	SetGroupName(v string) *CmsV1MigrationPreview
	GetGroupName() *string
	SetMigratable(v bool) *CmsV1MigrationPreview
	GetMigratable() *bool
	SetObserveGroupId(v string) *CmsV1MigrationPreview
	GetObserveGroupId() *string
	SetOriginGroupId(v string) *CmsV1MigrationPreview
	GetOriginGroupId() *string
	SetReason(v string) *CmsV1MigrationPreview
	GetReason() *string
	SetSourceReadOnly(v bool) *CmsV1MigrationPreview
	GetSourceReadOnly() *bool
	SetSourceType(v string) *CmsV1MigrationPreview
	GetSourceType() *string
	SetTargetGroupName(v string) *CmsV1MigrationPreview
	GetTargetGroupName() *string
}

type CmsV1MigrationPreview struct {
	// The itemized check results for name and resource support.
	Checks []*CmsV1MigrationCheck `json:"checks,omitempty" xml:"checks,omitempty" type:"Repeated"`
	// The name of the source application group.
	//
	// example:
	//
	// order-service
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// Indicates whether the dry run is passed.
	//
	// example:
	//
	// false
	Migratable *bool `json:"migratable,omitempty" xml:"migratable,omitempty"`
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
	// The reason for the dry run failure.
	//
	// example:
	//
	// NAME_INVALID
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// Indicates whether the source 1.0 group is persistently read-only.
	//
	// example:
	//
	// false
	SourceReadOnly *bool `json:"sourceReadOnly,omitempty" xml:"sourceReadOnly,omitempty"`
	// The type of the source application group. custom indicates a custom group, tag indicates a tag-based group, and regMgr indicates a resource group-based group.
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

func (s CmsV1MigrationPreview) String() string {
	return dara.Prettify(s)
}

func (s CmsV1MigrationPreview) GoString() string {
	return s.String()
}

func (s *CmsV1MigrationPreview) GetChecks() []*CmsV1MigrationCheck {
	return s.Checks
}

func (s *CmsV1MigrationPreview) GetGroupName() *string {
	return s.GroupName
}

func (s *CmsV1MigrationPreview) GetMigratable() *bool {
	return s.Migratable
}

func (s *CmsV1MigrationPreview) GetObserveGroupId() *string {
	return s.ObserveGroupId
}

func (s *CmsV1MigrationPreview) GetOriginGroupId() *string {
	return s.OriginGroupId
}

func (s *CmsV1MigrationPreview) GetReason() *string {
	return s.Reason
}

func (s *CmsV1MigrationPreview) GetSourceReadOnly() *bool {
	return s.SourceReadOnly
}

func (s *CmsV1MigrationPreview) GetSourceType() *string {
	return s.SourceType
}

func (s *CmsV1MigrationPreview) GetTargetGroupName() *string {
	return s.TargetGroupName
}

func (s *CmsV1MigrationPreview) SetChecks(v []*CmsV1MigrationCheck) *CmsV1MigrationPreview {
	s.Checks = v
	return s
}

func (s *CmsV1MigrationPreview) SetGroupName(v string) *CmsV1MigrationPreview {
	s.GroupName = &v
	return s
}

func (s *CmsV1MigrationPreview) SetMigratable(v bool) *CmsV1MigrationPreview {
	s.Migratable = &v
	return s
}

func (s *CmsV1MigrationPreview) SetObserveGroupId(v string) *CmsV1MigrationPreview {
	s.ObserveGroupId = &v
	return s
}

func (s *CmsV1MigrationPreview) SetOriginGroupId(v string) *CmsV1MigrationPreview {
	s.OriginGroupId = &v
	return s
}

func (s *CmsV1MigrationPreview) SetReason(v string) *CmsV1MigrationPreview {
	s.Reason = &v
	return s
}

func (s *CmsV1MigrationPreview) SetSourceReadOnly(v bool) *CmsV1MigrationPreview {
	s.SourceReadOnly = &v
	return s
}

func (s *CmsV1MigrationPreview) SetSourceType(v string) *CmsV1MigrationPreview {
	s.SourceType = &v
	return s
}

func (s *CmsV1MigrationPreview) SetTargetGroupName(v string) *CmsV1MigrationPreview {
	s.TargetGroupName = &v
	return s
}

func (s *CmsV1MigrationPreview) Validate() error {
	if s.Checks != nil {
		for _, item := range s.Checks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
