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
	// The frequency of high-frequency backups. Valid values:
	//
	// - **-1**: High-frequency backup is disabled.
	//
	// - **30**: every 30 minutes.
	//
	// - **60**: every 1 hour.
	//
	// - **120**: every 2 hours.
	//
	// - **180**: every 3 hours.
	//
	// - **240**: every 4 hours.
	//
	// - **360**: every 6 hours.
	//
	// - **480**: every 8 hours.
	//
	// - **720**: every 12 hours.
	//
	// > 	- If you set **SnapshotBackupType*	- to **Standard**, the value of this parameter is -1.
	//
	// >
	//
	// > 	- High-frequency backup takes effect only if you set **SnapshotBackupType*	- to **Flash*	- and set this parameter to a value greater than 0.
	//
	// example:
	//
	// -1
	BackupInterval *string `json:"BackupInterval,omitempty" xml:"BackupInterval,omitempty"`
	// The number of days to retain full backups.
	//
	// > - For instances that were created before September 10, 2021, the default retention period is 7 days.
	//
	// >
	//
	// > - For instances that are created after September 10, 2021, the default retention period is 30 days.
	//
	// example:
	//
	// 30
	BackupRetentionPeriod *int64 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The policy to retain backups when you release the instance.
	//
	// - 0: All backup sets of the instance are deleted when the instance is released.
	//
	// - 1: An automatic backup is performed when the instance is released, and this backup is retained for a long time.
	//
	// - 2: An automatic backup is performed when the instance is released, and all backup sets of the instance are retained for a long time.
	//
	// For more information, see [Long-term backup retention](https://help.aliyun.com/document_detail/2779111.html).
	//
	// example:
	//
	// 2
	BackupRetentionPolicyOnClusterDeletion *int32 `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// The days of the week to perform geo-redundant backups. Valid values:
	//
	// 1. Monday
	//
	// 2. Tuesday
	//
	// 3. Wednesday
	//
	// 4. Thursday
	//
	// 5. Friday
	//
	// 6. Saturday
	//
	// 7. Sunday
	//
	// > This parameter is required if you enable geo-redundancy.
	//
	// >
	//
	// > - To specify multiple days, separate them with commas (,).
	//
	// >
	//
	// > - If you set the backup method to conventional backup, the days of the week specified by this parameter must be a subset of the days of the week specified by PreferredBackupPeriod.
	//
	// example:
	//
	// Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday
	CrossBackupPeriod *string `json:"CrossBackupPeriod,omitempty" xml:"CrossBackupPeriod,omitempty"`
	// The policy for geo-redundant backups. Valid values:
	//
	// - update: Modify the geo-redundancy policy.
	//
	// - delete: Delete the geo-redundancy policy.
	//
	// > This parameter is required if you enable geo-redundancy.
	//
	// example:
	//
	// update
	CrossBackupType *string `json:"CrossBackupType,omitempty" xml:"CrossBackupType,omitempty"`
	// The retention policy for cross-region log backups. Valid values:
	//
	// - delay: Retain the backup for a specified period.
	//
	// - never: Retain the backup permanently.
	//
	// > This parameter is required if you enable geo-redundancy.
	//
	// example:
	//
	// delay
	CrossLogRetentionType *string `json:"CrossLogRetentionType,omitempty" xml:"CrossLogRetentionType,omitempty"`
	// The number of days to retain cross-region log backups. Valid values: 3 to 1825. The value must be less than or equal to the value of CrossRetentionValue.
	//
	// > This parameter is required if you enable geo-redundancy.
	//
	// example:
	//
	// 3
	CrossLogRetentionValue *int32 `json:"CrossLogRetentionValue,omitempty" xml:"CrossLogRetentionValue,omitempty"`
	// The retention policy for geo-redundant backups. Valid values:
	//
	// - delay: Retain the backup for a specified period.
	//
	// - never: Retain the backup permanently.
	//
	// > This parameter is required if you enable geo-redundancy.
	//
	// example:
	//
	// delay
	CrossRetentionType *string `json:"CrossRetentionType,omitempty" xml:"CrossRetentionType,omitempty"`
	// The number of days to retain geo-redundant backups. Valid values: 3 to 1825.
	//
	// > - This parameter is required if you enable geo-redundancy.
	//
	// >
	//
	// > - This parameter is required if you set CrossRetentionType to delay.
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
	// The region ID of the geo-redundant backup.
	//
	// > This parameter is required if you enable geo-redundancy.
	//
	// example:
	//
	// cn-hangzhou
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// Specifies whether to enable log backup. Valid values:
	//
	// - **0**: Disable log backup. This is the default value.
	//
	// - **1**: Enable log backup.
	//
	// 	Notice:
	//
	// You cannot disable log backup for sharded cluster instances.
	//
	// example:
	//
	// 0
	EnableBackupLog *int64 `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// Specifies whether to enable cross-region log backup. Valid values:
	//
	// > This parameter is required if you enable geo-redundancy.
	//
	// >
	//
	// > - 1: Enable cross-region log backup. This value is required for sharded cluster instances. This value is also required for replica set instances if you want to enable geo-redundant point-in-time recovery.
	//
	// >
	//
	// > - 0: Disable cross-region log backup.
	//
	// example:
	//
	// 1
	EnableCrossLogBackup *int32 `json:"EnableCrossLogBackup,omitempty" xml:"EnableCrossLogBackup,omitempty"`
	// The number of days to retain high-frequency backups. Before you specify this parameter, you must set the BackupInterval parameter. The default retention period is 1 day.
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
	// > 	- This parameter is required when you restore a deleted instance.
	//
	// >
	//
	// > 	- This parameter is required when you clone an instance from a geo-redundant backup.
	//
	// example:
	//
	// replicate
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of days to retain log backups. Default value: 7.
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
	// - **Monday**
	//
	// - **Tuesday**
	//
	// - **Wednesday**
	//
	// - **Thursday**
	//
	// - **Friday**
	//
	// - **Saturday**
	//
	// - **Sunday**
	//
	// 	Notice:
	//
	// To ensure data security, back up the MongoDB instance at least twice a week.
	//
	//
	//
	// > To specify multiple backup cycles, separate them with commas (,).
	//
	// example:
	//
	// Monday,Wednesday,Friday,Sunday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The time range to perform a backup. Specify the time in the *HH:mm*Z-*HH:mm*Z format. The time is displayed in Coordinated Universal Time (UTC).
	//
	// > The time range must be 1 hour.
	//
	// example:
	//
	// 03:00Z-04:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// Specifies whether to enable hourly sparse backup. Valid values:
	//
	// - true: If the backup frequency is in minutes, all snapshots that are generated within the last hour are retained. For snapshots that were generated more than 1 hour ago but less than 24 hours ago, only the first snapshot that is generated after each full hour is retained.
	//
	// - false: All snapshots are retained within the high-frequency backup retention period.
	//
	// example:
	//
	// true
	PreserveOneEachHour  *bool   `json:"PreserveOneEachHour,omitempty" xml:"PreserveOneEachHour,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The snapshot backup type. Valid values:
	//
	// - **Flash**: second-level backup.
	//
	// - **Standard**: conventional backup. This is the default value.
	//
	// example:
	//
	// Standard
	SnapshotBackupType *string `json:"SnapshotBackupType,omitempty" xml:"SnapshotBackupType,omitempty"`
	// The region ID of the instance.
	//
	// > - This parameter is required if you restore a deleted instance.
	//
	// >
	//
	// > - This parameter is required if you enable geo-redundancy.
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
