// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAdvancedBackupPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedBackupDataPoliciesShrink(v string) *ModifyAdvancedBackupPolicyShrinkRequest
	GetAdvancedBackupDataPoliciesShrink() *string
	SetAdvancedBackupLogPoliciesShrink(v string) *ModifyAdvancedBackupPolicyShrinkRequest
	GetAdvancedBackupLogPoliciesShrink() *string
	SetBackupRetentionPolicyOnClusterDeletion(v int32) *ModifyAdvancedBackupPolicyShrinkRequest
	GetBackupRetentionPolicyOnClusterDeletion() *int32
	SetDBInstanceId(v string) *ModifyAdvancedBackupPolicyShrinkRequest
	GetDBInstanceId() *string
	SetPreferredBackupTime(v string) *ModifyAdvancedBackupPolicyShrinkRequest
	GetPreferredBackupTime() *string
	SetSnapshotBackupType(v string) *ModifyAdvancedBackupPolicyShrinkRequest
	GetSnapshotBackupType() *string
}

type ModifyAdvancedBackupPolicyShrinkRequest struct {
	// The list of backup policies.
	AdvancedBackupDataPoliciesShrink *string `json:"AdvancedBackupDataPolicies,omitempty" xml:"AdvancedBackupDataPolicies,omitempty"`
	// The list of log backup policies.
	AdvancedBackupLogPoliciesShrink *string `json:"AdvancedBackupLogPolicies,omitempty" xml:"AdvancedBackupLogPolicies,omitempty"`
	// The backup retention policy. Valid values:
	//
	// - 0: All backup sets of the instance are deleted immediately when the instance is released.
	//
	// - 1: The last backup set is retained for long-term retention (LTR) when the instance is released.
	//
	// - 2: All backup sets are retained for long-term retention (LTR) when the instance is released.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	BackupRetentionPolicyOnClusterDeletion *int32 `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-2zeb2d64cb46xxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The time period during which the backup is performed. The format is <i>HH:mm</i>Z-<i>HH:mm</i>Z (UTC).
	//
	// > The time range is limited to 1 hour.
	//
	// example:
	//
	// 09:00Z-10:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The snapshot backup type. Valid values:
	//
	// - **Flash**: backup within seconds.
	//
	// - **Standard**: standard backup. This is the default value.
	//
	// example:
	//
	// Standard
	SnapshotBackupType *string `json:"SnapshotBackupType,omitempty" xml:"SnapshotBackupType,omitempty"`
}

func (s ModifyAdvancedBackupPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvancedBackupPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAdvancedBackupPolicyShrinkRequest) GetAdvancedBackupDataPoliciesShrink() *string {
	return s.AdvancedBackupDataPoliciesShrink
}

func (s *ModifyAdvancedBackupPolicyShrinkRequest) GetAdvancedBackupLogPoliciesShrink() *string {
	return s.AdvancedBackupLogPoliciesShrink
}

func (s *ModifyAdvancedBackupPolicyShrinkRequest) GetBackupRetentionPolicyOnClusterDeletion() *int32 {
	return s.BackupRetentionPolicyOnClusterDeletion
}

func (s *ModifyAdvancedBackupPolicyShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyAdvancedBackupPolicyShrinkRequest) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *ModifyAdvancedBackupPolicyShrinkRequest) GetSnapshotBackupType() *string {
	return s.SnapshotBackupType
}

func (s *ModifyAdvancedBackupPolicyShrinkRequest) SetAdvancedBackupDataPoliciesShrink(v string) *ModifyAdvancedBackupPolicyShrinkRequest {
	s.AdvancedBackupDataPoliciesShrink = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyShrinkRequest) SetAdvancedBackupLogPoliciesShrink(v string) *ModifyAdvancedBackupPolicyShrinkRequest {
	s.AdvancedBackupLogPoliciesShrink = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyShrinkRequest) SetBackupRetentionPolicyOnClusterDeletion(v int32) *ModifyAdvancedBackupPolicyShrinkRequest {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyShrinkRequest) SetDBInstanceId(v string) *ModifyAdvancedBackupPolicyShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyShrinkRequest) SetPreferredBackupTime(v string) *ModifyAdvancedBackupPolicyShrinkRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyShrinkRequest) SetSnapshotBackupType(v string) *ModifyAdvancedBackupPolicyShrinkRequest {
	s.SnapshotBackupType = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
