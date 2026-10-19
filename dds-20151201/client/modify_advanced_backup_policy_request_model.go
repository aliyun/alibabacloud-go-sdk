// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAdvancedBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedBackupDataPolicies(v []*ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) *ModifyAdvancedBackupPolicyRequest
	GetAdvancedBackupDataPolicies() []*ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies
	SetAdvancedBackupLogPolicies(v []*ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) *ModifyAdvancedBackupPolicyRequest
	GetAdvancedBackupLogPolicies() []*ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies
	SetBackupRetentionPolicyOnClusterDeletion(v int32) *ModifyAdvancedBackupPolicyRequest
	GetBackupRetentionPolicyOnClusterDeletion() *int32
	SetDBInstanceId(v string) *ModifyAdvancedBackupPolicyRequest
	GetDBInstanceId() *string
	SetPreferredBackupTime(v string) *ModifyAdvancedBackupPolicyRequest
	GetPreferredBackupTime() *string
	SetSnapshotBackupType(v string) *ModifyAdvancedBackupPolicyRequest
	GetSnapshotBackupType() *string
}

type ModifyAdvancedBackupPolicyRequest struct {
	// The list of backup policies.
	AdvancedBackupDataPolicies []*ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies `json:"AdvancedBackupDataPolicies,omitempty" xml:"AdvancedBackupDataPolicies,omitempty" type:"Repeated"`
	// The list of log backup policies.
	AdvancedBackupLogPolicies []*ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies `json:"AdvancedBackupLogPolicies,omitempty" xml:"AdvancedBackupLogPolicies,omitempty" type:"Repeated"`
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

func (s ModifyAdvancedBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvancedBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyAdvancedBackupPolicyRequest) GetAdvancedBackupDataPolicies() []*ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies {
	return s.AdvancedBackupDataPolicies
}

func (s *ModifyAdvancedBackupPolicyRequest) GetAdvancedBackupLogPolicies() []*ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies {
	return s.AdvancedBackupLogPolicies
}

func (s *ModifyAdvancedBackupPolicyRequest) GetBackupRetentionPolicyOnClusterDeletion() *int32 {
	return s.BackupRetentionPolicyOnClusterDeletion
}

func (s *ModifyAdvancedBackupPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyAdvancedBackupPolicyRequest) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *ModifyAdvancedBackupPolicyRequest) GetSnapshotBackupType() *string {
	return s.SnapshotBackupType
}

func (s *ModifyAdvancedBackupPolicyRequest) SetAdvancedBackupDataPolicies(v []*ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) *ModifyAdvancedBackupPolicyRequest {
	s.AdvancedBackupDataPolicies = v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequest) SetAdvancedBackupLogPolicies(v []*ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) *ModifyAdvancedBackupPolicyRequest {
	s.AdvancedBackupLogPolicies = v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequest) SetBackupRetentionPolicyOnClusterDeletion(v int32) *ModifyAdvancedBackupPolicyRequest {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequest) SetDBInstanceId(v string) *ModifyAdvancedBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyAdvancedBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequest) SetSnapshotBackupType(v string) *ModifyAdvancedBackupPolicyRequest {
	s.SnapshotBackupType = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequest) Validate() error {
	if s.AdvancedBackupDataPolicies != nil {
		for _, item := range s.AdvancedBackupDataPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AdvancedBackupLogPolicies != nil {
		for _, item := range s.AdvancedBackupLogPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies struct {
	// The action. Valid values:
	//
	// - **CREATE**: Creates a policy.
	//
	// - **UPDATE**: Updates a policy.
	//
	// - **DELETE**: Deletes a policy. If this value is specified, only the PolicyId parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// CREATE
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The scheduling type of the advanced policy. Valid values:
	//
	// - dayOfWeek: scheduled by week.
	//
	// - dayOfMonth: scheduled by month.
	//
	// - dayOfYear: scheduled by year.
	//
	// - backupInterval: scheduled at a fixed interval (high-frequency backup).
	//
	// >
	//
	// > - Multiple high-frequency backup policies are not supported.
	//
	// > - Multiple geo-redundancy backups are not supported, and geo-redundancy policies only support dayOfWeek.
	//
	// 	Notice: Required when ActionType is not DELETE.
	//
	// example:
	//
	// dayOfWeek
	BackupPolicy *string `json:"BackupPolicy,omitempty" xml:"BackupPolicy,omitempty"`
	// The destination region of the backup policy.
	//
	// 	Notice: Required when ActionType is not DELETE.
	//
	// example:
	//
	// cn-guangzhou
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// The destination type of the backup policy. Valid values:
	//
	//
	//
	// - level1: same-region backup.
	//
	// - level2Cross: cross-region backup.
	//
	// 	Notice: Required when ActionType is not DELETE.
	//
	// example:
	//
	// level1
	DestType *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	// The ID of the backup policy.
	//
	//
	// 	Notice: Required when ActionType is not CREATE.
	//
	// example:
	//
	// xxxxxx
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The backup cycle. Valid values:
	//
	// - When BackupPolicy is set to dayOfWeek: 1, 2, 3, 4, 5, 6, and 7 represent Monday to Sunday. At least two backups must be performed per week.
	//
	// - When BackupPolicy is set to dayOfMonth: 1, 3, and -1 represent the 1st, 3rd, and last day of each month. -1 represents the last day of the month.
	//
	// - When BackupPolicy is set to dayOfYear: 01-07 represents January 7. Only one day can be selected.
	//
	// - When BackupPolicy is set to backupInterval: specify one of the following numbers in minutes. For sharded clusters, the valid values are 30, 60, 120, 180, 240, 360, 480, and 720. For replica sets, the valid values are 15, 30, 60, 120, 180, 240, 360, 480, and 720.
	//
	// Note:
	//
	// 	Notice: Required when ActionType is not DELETE.
	//
	// Geo-redundancy policies only support dayOfWeek.
	//
	// example:
	//
	// 1,4,6
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// Specifies whether to enable hourly sparse backup. Valid values:
	//
	// - true: When the backup frequency is in minutes, all snapshots taken within 1 hour from the current time are retained. For snapshots taken more than 1 hour but within 24 hours from the current time, only the first snapshot after the hour is retained.
	//
	// - false: All snapshots within the retention period of high-frequency backups are retained.
	//
	// example:
	//
	// False
	PreserveOneEachHour *bool `json:"PreserveOneEachHour,omitempty" xml:"PreserveOneEachHour,omitempty"`
	// The retention period of the backup plan in days. Valid values:
	//
	// - Same-region backup: 7 to 7300 days.
	//
	// - Geo-redundancy: 3 to 1825 days.
	//
	// 	Notice: Required when RetentionType is set to delay.
	//
	// example:
	//
	// 7
	RetentionPeriod *int32 `json:"RetentionPeriod,omitempty" xml:"RetentionPeriod,omitempty"`
	// The retention cycle type of the backup set. Valid values:
	//
	// - **never**: long-term retention (LTR).
	//
	// - **delay**: retention for a fixed number of days.
	//
	// 	Notice: Required when ActionType is not DELETE.
	//
	// example:
	//
	// delay
	RetentionType *string `json:"RetentionType,omitempty" xml:"RetentionType,omitempty"`
	// The region where the instance resides.
	//
	// 	Notice: Required when ActionType is not DELETE.
	//
	// example:
	//
	// cn-guangzhou
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The source type of the backup policy. Valid values:
	//
	//
	//
	// - db: same-region backup.
	//
	// - level1: cross-region backup.
	//
	// 	Notice: Required when ActionType is not DELETE.
	//
	// example:
	//
	// db
	SrcType *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) GoString() string {
	return s.String()
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) GetActionType() *string {
	return s.ActionType
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) GetBackupPolicy() *string {
	return s.BackupPolicy
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) GetDestType() *string {
	return s.DestType
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) GetPreserveOneEachHour() *bool {
	return s.PreserveOneEachHour
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) GetRetentionPeriod() *int32 {
	return s.RetentionPeriod
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) GetRetentionType() *string {
	return s.RetentionType
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) SetActionType(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies {
	s.ActionType = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) SetBackupPolicy(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies {
	s.BackupPolicy = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) SetDestRegion(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies {
	s.DestRegion = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) SetDestType(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies {
	s.DestType = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) SetPolicyId(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies {
	s.PolicyId = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) SetPreferredBackupPeriod(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) SetPreserveOneEachHour(v bool) *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies {
	s.PreserveOneEachHour = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) SetRetentionPeriod(v int32) *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies {
	s.RetentionPeriod = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) SetRetentionType(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies {
	s.RetentionType = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) SetSrcRegion(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies {
	s.SrcRegion = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) SetSrcType(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies {
	s.SrcType = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupDataPolicies) Validate() error {
	return dara.Validate(s)
}

type ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies struct {
	// The operation type. Valid values:
	//
	// - CREATE: Creates a policy.
	//
	// - UPDATE: Updates a policy.
	//
	// 	Notice: This parameter takes effect only for geo-redundancy. To delete geo-redundancy, set EnableLogBackup to 0 and set ActionType to UPDATE.
	//
	// This parameter is required.
	//
	// example:
	//
	// CREATE
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The region where the geo-redundancy log backup is stored.
	//
	// > Required for geo-redundancy.
	//
	// example:
	//
	// cn-huhehaote
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// The destination type of the log backup policy. Valid values:
	//
	//
	//
	// - level1: same-region backup.
	//
	// - level2Cross: cross-region backup.
	//
	// example:
	//
	// level1
	DestType *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	// Specifies whether to enable log backup. Valid values:
	//
	// - **1**: enabled.
	//
	// - **0**: disabled.
	//
	// >
	//
	// > - Log backup must be enabled for sharded instances.
	//
	// > - You must enable same-region log backup before you can enable cross-region log backup.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	EnableLogBackup *int32 `json:"EnableLogBackup,omitempty" xml:"EnableLogBackup,omitempty"`
	// The ID of the log backup policy. This parameter is required only for cross-region log backup policies.
	//
	// example:
	//
	// xxxx
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The retention period of the log backup in days. This parameter takes effect when RetentionType is set to delay. Valid values:
	//
	// - Same-region backup: 7 to 7300 days.
	//
	// - Geo-redundancy: 3 to 1825 days. The value must be less than the backup retention period of the geo-redundancy policy.
	//
	// example:
	//
	// 7
	RetentionPeriod *int32 `json:"RetentionPeriod,omitempty" xml:"RetentionPeriod,omitempty"`
	// The retention cycle type of the log backup. Valid values:
	//
	// - **never**: long-term retention (LTR).
	//
	// - **delay**: retention for a fixed number of days.
	//
	// example:
	//
	// delay
	RetentionType *string `json:"RetentionType,omitempty" xml:"RetentionType,omitempty"`
	// The region where the source instance resides.
	//
	// example:
	//
	// cn-hangzhou
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The source type of the log backup policy. Valid values:
	//
	//
	//
	// - db: same-region backup.
	//
	// - level1: cross-region backup.
	//
	// example:
	//
	// db
	SrcType *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) GoString() string {
	return s.String()
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) GetActionType() *string {
	return s.ActionType
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) GetDestType() *string {
	return s.DestType
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) GetEnableLogBackup() *int32 {
	return s.EnableLogBackup
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) GetRetentionPeriod() *int32 {
	return s.RetentionPeriod
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) GetRetentionType() *string {
	return s.RetentionType
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) SetActionType(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies {
	s.ActionType = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) SetDestRegion(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies {
	s.DestRegion = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) SetDestType(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies {
	s.DestType = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) SetEnableLogBackup(v int32) *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies {
	s.EnableLogBackup = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) SetPolicyId(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies {
	s.PolicyId = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) SetRetentionPeriod(v int32) *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies {
	s.RetentionPeriod = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) SetRetentionType(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies {
	s.RetentionType = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) SetSrcRegion(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies {
	s.SrcRegion = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) SetSrcType(v string) *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies {
	s.SrcType = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyRequestAdvancedBackupLogPolicies) Validate() error {
	return dara.Validate(s)
}
