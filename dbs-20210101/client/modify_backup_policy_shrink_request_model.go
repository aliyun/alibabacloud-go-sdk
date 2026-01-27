// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvanceDataPoliciesShrink(v string) *ModifyBackupPolicyShrinkRequest
	GetAdvanceDataPoliciesShrink() *string
	SetAdvanceIncPoliciesShrink(v string) *ModifyBackupPolicyShrinkRequest
	GetAdvanceIncPoliciesShrink() *string
	SetAdvanceLogPoliciesShrink(v string) *ModifyBackupPolicyShrinkRequest
	GetAdvanceLogPoliciesShrink() *string
	SetBackupMethod(v string) *ModifyBackupPolicyShrinkRequest
	GetBackupMethod() *string
	SetBackupPriority(v int32) *ModifyBackupPolicyShrinkRequest
	GetBackupPriority() *int32
	SetBackupRetentionPolicyOnClusterDeletion(v string) *ModifyBackupPolicyShrinkRequest
	GetBackupRetentionPolicyOnClusterDeletion() *string
	SetCategory(v string) *ModifyBackupPolicyShrinkRequest
	GetCategory() *string
	SetEnableIncBackup(v bool) *ModifyBackupPolicyShrinkRequest
	GetEnableIncBackup() *bool
	SetInstanceName(v string) *ModifyBackupPolicyShrinkRequest
	GetInstanceName() *string
	SetPreferredBackupWindowBegin(v string) *ModifyBackupPolicyShrinkRequest
	GetPreferredBackupWindowBegin() *string
	SetRegionCode(v string) *ModifyBackupPolicyShrinkRequest
	GetRegionCode() *string
}

type ModifyBackupPolicyShrinkRequest struct {
	// The details of data backup policies.
	AdvanceDataPoliciesShrink              *string `json:"AdvanceDataPolicies,omitempty" xml:"AdvanceDataPolicies,omitempty"`
	AdvanceIncPoliciesShrink               *string `json:"AdvanceIncPolicies,omitempty" xml:"AdvanceIncPolicies,omitempty"`
	AdvanceLogPoliciesShrink               *string `json:"AdvanceLogPolicies,omitempty" xml:"AdvanceLogPolicies,omitempty"`
	BackupMethod                           *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	BackupPriority                         *int32  `json:"BackupPriority,omitempty" xml:"BackupPriority,omitempty"`
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	Category                               *string `json:"Category,omitempty" xml:"Category,omitempty"`
	EnableIncBackup                        *bool   `json:"EnableIncBackup,omitempty" xml:"EnableIncBackup,omitempty"`
	// The ID of the PolarDB instance.
	//
	// example:
	//
	// pc-2ze3nrr64c5****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The start time of a basic backup.
	//
	// example:
	//
	// 17:00Z
	PreferredBackupWindowBegin *string `json:"PreferredBackupWindowBegin,omitempty" xml:"PreferredBackupWindowBegin,omitempty"`
	// The region in which backup sets reside.
	//
	// example:
	//
	// cn-shanghai
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s ModifyBackupPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyShrinkRequest) GetAdvanceDataPoliciesShrink() *string {
	return s.AdvanceDataPoliciesShrink
}

func (s *ModifyBackupPolicyShrinkRequest) GetAdvanceIncPoliciesShrink() *string {
	return s.AdvanceIncPoliciesShrink
}

func (s *ModifyBackupPolicyShrinkRequest) GetAdvanceLogPoliciesShrink() *string {
	return s.AdvanceLogPoliciesShrink
}

func (s *ModifyBackupPolicyShrinkRequest) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *ModifyBackupPolicyShrinkRequest) GetBackupPriority() *int32 {
	return s.BackupPriority
}

func (s *ModifyBackupPolicyShrinkRequest) GetBackupRetentionPolicyOnClusterDeletion() *string {
	return s.BackupRetentionPolicyOnClusterDeletion
}

func (s *ModifyBackupPolicyShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyBackupPolicyShrinkRequest) GetEnableIncBackup() *bool {
	return s.EnableIncBackup
}

func (s *ModifyBackupPolicyShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyBackupPolicyShrinkRequest) GetPreferredBackupWindowBegin() *string {
	return s.PreferredBackupWindowBegin
}

func (s *ModifyBackupPolicyShrinkRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *ModifyBackupPolicyShrinkRequest) SetAdvanceDataPoliciesShrink(v string) *ModifyBackupPolicyShrinkRequest {
	s.AdvanceDataPoliciesShrink = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetAdvanceIncPoliciesShrink(v string) *ModifyBackupPolicyShrinkRequest {
	s.AdvanceIncPoliciesShrink = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetAdvanceLogPoliciesShrink(v string) *ModifyBackupPolicyShrinkRequest {
	s.AdvanceLogPoliciesShrink = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetBackupMethod(v string) *ModifyBackupPolicyShrinkRequest {
	s.BackupMethod = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetBackupPriority(v int32) *ModifyBackupPolicyShrinkRequest {
	s.BackupPriority = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetBackupRetentionPolicyOnClusterDeletion(v string) *ModifyBackupPolicyShrinkRequest {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetCategory(v string) *ModifyBackupPolicyShrinkRequest {
	s.Category = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetEnableIncBackup(v bool) *ModifyBackupPolicyShrinkRequest {
	s.EnableIncBackup = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetInstanceName(v string) *ModifyBackupPolicyShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetPreferredBackupWindowBegin(v string) *ModifyBackupPolicyShrinkRequest {
	s.PreferredBackupWindowBegin = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetRegionCode(v string) *ModifyBackupPolicyShrinkRequest {
	s.RegionCode = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
