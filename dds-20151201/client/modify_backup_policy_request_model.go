// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupInterval(v string) *ModifyBackupPolicyRequest
	GetBackupInterval() *string
	SetBackupRetentionPeriod(v int64) *ModifyBackupPolicyRequest
	GetBackupRetentionPeriod() *int64
	SetBackupRetentionPolicyOnClusterDeletion(v int32) *ModifyBackupPolicyRequest
	GetBackupRetentionPolicyOnClusterDeletion() *int32
	SetCrossBackupPeriod(v string) *ModifyBackupPolicyRequest
	GetCrossBackupPeriod() *string
	SetCrossBackupType(v string) *ModifyBackupPolicyRequest
	GetCrossBackupType() *string
	SetCrossLogRetentionType(v string) *ModifyBackupPolicyRequest
	GetCrossLogRetentionType() *string
	SetCrossLogRetentionValue(v int32) *ModifyBackupPolicyRequest
	GetCrossLogRetentionValue() *int32
	SetCrossRetentionType(v string) *ModifyBackupPolicyRequest
	GetCrossRetentionType() *string
	SetCrossRetentionValue(v int32) *ModifyBackupPolicyRequest
	GetCrossRetentionValue() *int32
	SetDBInstanceId(v string) *ModifyBackupPolicyRequest
	GetDBInstanceId() *string
	SetDestRegion(v string) *ModifyBackupPolicyRequest
	GetDestRegion() *string
	SetEnableBackupLog(v int64) *ModifyBackupPolicyRequest
	GetEnableBackupLog() *int64
	SetEnableCrossLogBackup(v int32) *ModifyBackupPolicyRequest
	GetEnableCrossLogBackup() *int32
	SetHighFrequencyBackupRetention(v int64) *ModifyBackupPolicyRequest
	GetHighFrequencyBackupRetention() *int64
	SetInstanceType(v string) *ModifyBackupPolicyRequest
	GetInstanceType() *string
	SetLogBackupRetentionPeriod(v int64) *ModifyBackupPolicyRequest
	GetLogBackupRetentionPeriod() *int64
	SetOwnerAccount(v string) *ModifyBackupPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyBackupPolicyRequest
	GetOwnerId() *int64
	SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupTime() *string
	SetPreserveOneEachHour(v bool) *ModifyBackupPolicyRequest
	GetPreserveOneEachHour() *bool
	SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest
	GetResourceOwnerId() *int64
	SetSnapshotBackupType(v string) *ModifyBackupPolicyRequest
	GetSnapshotBackupType() *string
	SetSrcRegion(v string) *ModifyBackupPolicyRequest
	GetSrcRegion() *string
}

type ModifyBackupPolicyRequest struct {
	// The high-frequency backup frequency. Valid values:
	//
	// - **-1**: Disables high-frequency backups.
	//
	// - **30**: Every 30 minutes.
	//
	// - **60**: Every 1 hour.
	//
	// - **120**: Every 2 hours.
	//
	// - **180**: Every 3 hours.
	//
	// - **240**: Every 4 hours.
	//
	// - **360**: Every 6 hours.
	//
	// - **480**: Every 8 hours.
	//
	// - **720**: Every 12 hours.
	//
	// > - If **SnapshotBackupType*	- is set to **Standard**, this parameter must be set to **-1**.
	//
	// > - If **SnapshotBackupType*	- is set to **Flash*	- and this parameter is set to a value greater than 0, high-frequency backups take effect.
	//
	// example:
	//
	// -1
	BackupInterval *string `json:"BackupInterval,omitempty" xml:"BackupInterval,omitempty"`
	// The number of days for which full backups are retained.
	//
	// > - For users who created instances before September 10, 2021, the default retention period is 7 days.
	//
	// > - For new users who created instances on or after September 10, 2021, the default retention period is 30 days.
	//
	// example:
	//
	// 30
	BackupRetentionPeriod *int64 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The backup data retention policy. Valid values:
	//
	// - 0: Deletes all backup sets of the instance immediately when the instance is released.
	//
	// - 1: An automatic backup is performed when the instance is released, and the backup set is retained for long-term retention (LTR).
	//
	// - 2: An automatic backup is performed when the instance is released, and all backup sets of the cluster are retained for long-term retention (LTR).
	//
	// For more information, see [Long-term retention backups](https://help.aliyun.com/document_detail/2779111.html).
	//
	// example:
	//
	// 2
	BackupRetentionPolicyOnClusterDeletion *int32 `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// The days of the week when cross-region backup is performed. Valid values:
	//
	// 1. Monday: Monday.
	//
	// 2. Tuesday: Tuesday.
	//
	// 3. Wednesday: Wednesday.
	//
	// 4. Thursday: Thursday.
	//
	// 5. Friday: Friday.
	//
	// 6. Saturday: Saturday.
	//
	// 7. Sunday: Sunday.
	//
	// > This parameter is required for cross-region backup.
	//
	// > - If you specify multiple values, separate them with commas (,).
	//
	// > - When the backup settings are set to standard backup, the value must be a subset of the PreferredBackupPeriod backup epoch.
	//
	// example:
	//
	// Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday
	CrossBackupPeriod *string `json:"CrossBackupPeriod,omitempty" xml:"CrossBackupPeriod,omitempty"`
	// The operation policy for cross-region backup. Valid values:
	//
	// - update: Modifies the cross-region backup policy.
	//
	// - delete: Deletes the cross-region backup policy.
	//
	// > This parameter is required for cross-region backup.
	//
	// example:
	//
	// update
	CrossBackupType *string `json:"CrossBackupType,omitempty" xml:"CrossBackupType,omitempty"`
	// The retention type for cross-region log backup. Valid values:
	//
	// - delay: Retains backups for a specified period.
	//
	// - never: Never expires.
	//
	// > This parameter is required for cross-region backup.
	//
	// example:
	//
	// delay
	CrossLogRetentionType *string `json:"CrossLogRetentionType,omitempty" xml:"CrossLogRetentionType,omitempty"`
	// The retention period for cross-region log backup. Valid values: 3 to 1825 days. The value must be less than or equal to the value of CrossRetentionValue.
	//
	//
	// > This parameter is required for cross-region backup.
	//
	// example:
	//
	// 3
	CrossLogRetentionValue *int32 `json:"CrossLogRetentionValue,omitempty" xml:"CrossLogRetentionValue,omitempty"`
	// The retention type for cross-region backup. Valid values:
	//
	// - delay: Retains backups for a specified period.
	//
	// - never: Never expires.
	//
	// > This parameter is required for cross-region backup.
	//
	// example:
	//
	// delay
	CrossRetentionType *string `json:"CrossRetentionType,omitempty" xml:"CrossRetentionType,omitempty"`
	// The retention period for cross-region backup. Valid values: 3 to 1825 days.
	//
	// >
	//
	// > - This parameter is required for cross-region backup.
	//
	// > - This parameter is required and must be specified when CrossRetentionType is set to delay.
	//
	// example:
	//
	// 7
	CrossRetentionValue *int32 `json:"CrossRetentionValue,omitempty" xml:"CrossRetentionValue,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp16cb162771****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region where backups are stored.
	//
	//
	// > This parameter is required for cross-region backup.
	//
	// example:
	//
	// cn-hangzhou
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// Specifies whether to enable log backup. Valid values:
	//
	// - **0**: Disables log backup. This is the default value.
	//
	// - **1**: Enables log backup.
	//
	// 	Notice: Sharded cluster instances do not support disabling log backup.</notice>
	//
	// example:
	//
	// 0
	EnableBackupLog *int64 `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// Specifies whether to enable cross-region log backup.
	//
	// > This parameter is required for cross-region backup.
	//
	// > - 1: Enables cross-region log backup. This value is required for sharded cluster instances. For replica set instances, this value is required if you want to support cross-region point-in-time recovery.
	//
	// > - 0: Disables cross-region log backup.
	//
	// example:
	//
	// 1
	EnableCrossLogBackup *int32 `json:"EnableCrossLogBackup,omitempty" xml:"EnableCrossLogBackup,omitempty"`
	// The number of days for which high-frequency backups are retained. Before you use this parameter, make sure that the BackupInterval parameter is specified. Default value: 1.
	//
	// example:
	//
	// 1
	HighFrequencyBackupRetention *int64 `json:"HighFrequencyBackupRetention,omitempty" xml:"HighFrequencyBackupRetention,omitempty"`
	// The instance type. Valid values:
	//
	// - replicate
	//
	// - sharding
	//
	// > - This parameter is required when the backup recovery type is set to rebuilding a deleted instance.
	//
	// > - This parameter is required when the backup recovery type is set to cloning an instance from cross-region backup.
	//
	// example:
	//
	// replicate
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of days for which log backups are retained. Default value: 7.
	//
	// Valid values: 7 to 730.
	//
	// example:
	//
	// 7
	LogBackupRetentionPeriod *int64  `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The backup cycle. Valid values:
	//
	// - **Monday**: Monday.
	//
	// - **Tuesday**: Tuesday.
	//
	// - **Wednesday**: Wednesday.
	//
	// - **Thursday**: Thursday.
	//
	// - **Friday**: Friday.
	//
	// - **Saturday**: Saturday.
	//
	// - **Sunday**: Sunday.
	//
	// 	Notice: To ensure data security, back up your MongoDB data at least twice a week.</notice>
	//
	// > If you specify multiple values, separate them with commas (,).
	//
	// example:
	//
	// Monday,Wednesday,Friday,Sunday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The time period during which backups are performed. The format is <i>HH:mm</i>Z-<i>HH:mm</i>Z (UTC).
	//
	// > The time period is limited to 1 hour.
	//
	// example:
	//
	// 03:00Z-04:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// Specifies whether to enable hourly sparse backup. Valid values:
	//
	// - true: If the backup frequency is set to minutes, all snapshots taken within 1 hour from the current time are retained. For snapshots taken more than 1 hour but within 24 hours from the current time, only the first snapshot taken after the hour is retained.
	//
	// - false: All snapshots taken within the high-frequency backup retention period are retained.
	//
	// example:
	//
	// true
	PreserveOneEachHour  *bool   `json:"PreserveOneEachHour,omitempty" xml:"PreserveOneEachHour,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The snapshot backup type. Valid values:
	//
	// - **Flash**: Backup within seconds.
	//
	// - **Standard**: Standard backup. This is the default value.
	//
	// example:
	//
	// Standard
	SnapshotBackupType *string `json:"SnapshotBackupType,omitempty" xml:"SnapshotBackupType,omitempty"`
	// The region where the instance is deployed.
	//
	// >
	//
	// > - This parameter is required to recover deleted instances.
	//
	// > - This parameter is required for cross-region backup.
	//
	// example:
	//
	// cn-beijing
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) GetBackupInterval() *string {
	return s.BackupInterval
}

func (s *ModifyBackupPolicyRequest) GetBackupRetentionPeriod() *int64 {
	return s.BackupRetentionPeriod
}

func (s *ModifyBackupPolicyRequest) GetBackupRetentionPolicyOnClusterDeletion() *int32 {
	return s.BackupRetentionPolicyOnClusterDeletion
}

func (s *ModifyBackupPolicyRequest) GetCrossBackupPeriod() *string {
	return s.CrossBackupPeriod
}

func (s *ModifyBackupPolicyRequest) GetCrossBackupType() *string {
	return s.CrossBackupType
}

func (s *ModifyBackupPolicyRequest) GetCrossLogRetentionType() *string {
	return s.CrossLogRetentionType
}

func (s *ModifyBackupPolicyRequest) GetCrossLogRetentionValue() *int32 {
	return s.CrossLogRetentionValue
}

func (s *ModifyBackupPolicyRequest) GetCrossRetentionType() *string {
	return s.CrossRetentionType
}

func (s *ModifyBackupPolicyRequest) GetCrossRetentionValue() *int32 {
	return s.CrossRetentionValue
}

func (s *ModifyBackupPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyBackupPolicyRequest) GetDestRegion() *string {
	return s.DestRegion
}

func (s *ModifyBackupPolicyRequest) GetEnableBackupLog() *int64 {
	return s.EnableBackupLog
}

func (s *ModifyBackupPolicyRequest) GetEnableCrossLogBackup() *int32 {
	return s.EnableCrossLogBackup
}

func (s *ModifyBackupPolicyRequest) GetHighFrequencyBackupRetention() *int64 {
	return s.HighFrequencyBackupRetention
}

func (s *ModifyBackupPolicyRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyBackupPolicyRequest) GetLogBackupRetentionPeriod() *int64 {
	return s.LogBackupRetentionPeriod
}

func (s *ModifyBackupPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyBackupPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *ModifyBackupPolicyRequest) GetPreserveOneEachHour() *bool {
	return s.PreserveOneEachHour
}

func (s *ModifyBackupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyBackupPolicyRequest) GetSnapshotBackupType() *string {
	return s.SnapshotBackupType
}

func (s *ModifyBackupPolicyRequest) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *ModifyBackupPolicyRequest) SetBackupInterval(v string) *ModifyBackupPolicyRequest {
	s.BackupInterval = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPeriod(v int64) *ModifyBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPolicyOnClusterDeletion(v int32) *ModifyBackupPolicyRequest {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetCrossBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.CrossBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetCrossBackupType(v string) *ModifyBackupPolicyRequest {
	s.CrossBackupType = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetCrossLogRetentionType(v string) *ModifyBackupPolicyRequest {
	s.CrossLogRetentionType = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetCrossLogRetentionValue(v int32) *ModifyBackupPolicyRequest {
	s.CrossLogRetentionValue = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetCrossRetentionType(v string) *ModifyBackupPolicyRequest {
	s.CrossRetentionType = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetCrossRetentionValue(v int32) *ModifyBackupPolicyRequest {
	s.CrossRetentionValue = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDBInstanceId(v string) *ModifyBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDestRegion(v string) *ModifyBackupPolicyRequest {
	s.DestRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableBackupLog(v int64) *ModifyBackupPolicyRequest {
	s.EnableBackupLog = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableCrossLogBackup(v int32) *ModifyBackupPolicyRequest {
	s.EnableCrossLogBackup = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetHighFrequencyBackupRetention(v int64) *ModifyBackupPolicyRequest {
	s.HighFrequencyBackupRetention = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetInstanceType(v string) *ModifyBackupPolicyRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetLogBackupRetentionPeriod(v int64) *ModifyBackupPolicyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreserveOneEachHour(v bool) *ModifyBackupPolicyRequest {
	s.PreserveOneEachHour = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetSnapshotBackupType(v string) *ModifyBackupPolicyRequest {
	s.SnapshotBackupType = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetSrcRegion(v string) *ModifyBackupPolicyRequest {
	s.SrcRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
