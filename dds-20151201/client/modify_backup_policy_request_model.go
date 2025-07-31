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
	// The frequency at which high-frequency backups are generated. Valid values:
	//
	// 	- **-1**: High-frequency backup is disabled.
	//
	// 	- **30**: High-frequency backups are generated every 30 minutes.
	//
	// 	- **60**: High-frequency backups are generated every 1 hour.
	//
	// 	- **120**: High-frequency backups are generated every 2 hours.
	//
	// 	- **180**: High-frequency backups are generated every 3 hours.
	//
	// 	- **240**: High-frequency backups are generated every 4 hours.
	//
	// 	- **360**: High-frequency backups are generated every 6 hours.
	//
	// 	- **480**: High-frequency backups are generated every 8 hours.
	//
	// 	- **720**: High-frequency backups are generated every 12 hours.
	//
	// >
	//
	// 	- If you set the **SnapshotBackupType*	- parameter to **Standard**, you must fix the value of this parameter to -1.
	//
	// 	- High-frequency backup takes effect only when you set the **SnapshotBackupType*	- parameter to **Flash*	- and this parameter to a value greater than 0.
	//
	// example:
	//
	// -1
	BackupInterval *string `json:"BackupInterval,omitempty" xml:"BackupInterval,omitempty"`
	// The retention period of full backups.
	//
	// >
	//
	// 	- If your instance is created before September 10, 2021, backups are retained for seven days by default.
	//
	// 	- If your instance is created after September 10, 2021, backups are retained for 30 days by default.
	//
	// example:
	//
	// 30
	BackupRetentionPeriod *int64 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The backup retention policy configured for the instance. Valid values:
	//
	// 	- 0: All backup sets are immediately deleted when the instance is released.
	//
	// 	- 1: Automatic backup is performed when the instance is released and the backup set is retained for a long period of time.
	//
	// 	- 2: Automatic backup is performed when the instance is released and all backup sets are retained for a long period of time.
	//
	// For more information, see [Retain the backup files of an ApsaraDB for MongoDB instance for a long period of time](https://help.aliyun.com/document_detail/2779111.html).
	//
	// example:
	//
	// 2
	BackupRetentionPolicyOnClusterDeletion *int32 `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// The day of the week on which the cross-region backup files are retained. Valid values:
	//
	// 1.  Monday
	//
	// 2.  Tuesday
	//
	// 3.  Wednesday
	//
	// 4.  Thursday
	//
	// 5.  Friday
	//
	// 6.  Saturday
	//
	// 7.  Sunday
	//
	// >  This parameter is required for a cross-region backup operation.
	//
	// 	- Separate multiple values with commas (,).
	//
	// 	- If you set the SnapshotBackupType parameter to Standard, the parameter value must fall within the value of the PreferredBackupPeriod parameter that specifies the standard backup period.
	//
	// example:
	//
	// Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday
	CrossBackupPeriod *string `json:"CrossBackupPeriod,omitempty" xml:"CrossBackupPeriod,omitempty"`
	// The action performed for the cross-region backup policy. Valid values:
	//
	// 	- update: modifies the cross-region backup policy.
	//
	// 	- delete: deletes the cross-region backup policy.
	//
	// >  This parameter is required for a cross-region backup operation.
	//
	// example:
	//
	// update
	CrossBackupType *string `json:"CrossBackupType,omitempty" xml:"CrossBackupType,omitempty"`
	// The retention type of the cross-region log backup files. Valid values:
	//
	// 	- delay: retains the cross-region backup files for a period of time.
	//
	// 	- never: permanently retains the cross-region backup files.
	//
	// >  This parameter is required for a cross-region backup operation.
	//
	// example:
	//
	// delay
	CrossLogRetentionType *string `json:"CrossLogRetentionType,omitempty" xml:"CrossLogRetentionType,omitempty"`
	// The retention period of the cross-region log backup files. Valid values: 3 to 1825. Unit: day. The parameter value must be less than or equal to the value of the CrossRetentionValue parameter.
	//
	// >  This parameter is required for a cross-region backup operation.
	//
	// example:
	//
	// 3
	CrossLogRetentionValue *int32 `json:"CrossLogRetentionValue,omitempty" xml:"CrossLogRetentionValue,omitempty"`
	// The retention type of the cross-region backup files. Valid values:
	//
	// 	- delay: retains the cross-region backup files for a period of time.
	//
	// 	- never: permanently retains the cross-region backup files.
	//
	// >  This parameter is required for a cross-region backup operation.
	//
	// example:
	//
	// delay
	CrossRetentionType *string `json:"CrossRetentionType,omitempty" xml:"CrossRetentionType,omitempty"`
	// The retention period of the cross-region backup files. Valid values: 3 to 1825. Unit: day.
	//
	// >
	//
	// 	- This parameter is required for a cross-region backup operation.
	//
	// 	- This parameter is required when you set the CrossRetentionType parameter to delay.
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
	// The region in which the backup files are retained.
	//
	// >  This parameter is required for a cross-region backup operation.
	//
	// example:
	//
	// cn-hangzhou
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// Specifies whether to enable the log backup feature. Valid values:
	//
	// 	- **0*	- (default): The log backup feature is disabled.
	//
	// 	- **1**: The log backup feature is enabled.
	//
	// example:
	//
	// 0
	EnableBackupLog *int64 `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// Specifies whether to enable the cross-region log backup feature.
	//
	// >  This parameter is required for a cross-region backup operation.
	//
	// 	- Valid values:1: enables the feature. The parameter value must be 1 for sharded cluster instances.
	//
	// 	- 0: disables the feature. The parameter value must be 0 for replica set instances.
	//
	// example:
	//
	// 1
	EnableCrossLogBackup *int32 `json:"EnableCrossLogBackup,omitempty" xml:"EnableCrossLogBackup,omitempty"`
	// The number of days for which high-frequency backup files are retained. Before you use this parameter, make sure that you specify the BackupInterval parameter. By default, high-frequency backup files are retained for one day.
	//
	// example:
	//
	// 1
	HighFrequencyBackupRetention *int64 `json:"HighFrequencyBackupRetention,omitempty" xml:"HighFrequencyBackupRetention,omitempty"`
	// The instance architecture. Valid values:
	//
	// 	- replicate
	//
	// 	- sharding
	//
	// >
	//
	// 	- This parameter is required when you set the RestoreType parameter to 2.
	//
	// 	- This parameter is required when you set the RestoreType parameter to 3.
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
	// The day of a week when the system regularly backs up data. Valid values:
	//
	// 	- **Monday**
	//
	// 	- **Tuesday**
	//
	// 	- **Wednesday**
	//
	// 	- **Thursday**
	//
	// 	- **Friday**
	//
	// 	- **Saturday**
	//
	// 	- **Sunday**
	//
	// **
	//
	// **Notice**: To ensure data security, make sure that the system backs up data at least twice a week.
	//
	// >  Separate multiple values with commas (,).
	//
	// example:
	//
	// Monday,Wednesday,Friday,Sunday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The start time of the backup. Specify the time in the ISO 8601 standard in the *HH:mm*Z-*HH:mm*Z format. The time must be in UTC.
	//
	// >  The time range is 1 hour.
	//
	// example:
	//
	// 03:00Z-04:00Z
	PreferredBackupTime  *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	PreserveOneEachHour  *bool   `json:"PreserveOneEachHour,omitempty" xml:"PreserveOneEachHour,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The snapshot backup type. Valid values:
	//
	// 	- **Flash**: single-digit second backup
	//
	// 	- **Standard*	- (default): standard backup
	//
	// example:
	//
	// Standard
	SnapshotBackupType *string `json:"SnapshotBackupType,omitempty" xml:"SnapshotBackupType,omitempty"`
	// The region ID of the instance.
	//
	// >
	//
	// 	- This parameter is required for the data restoration of a deleted instance.
	//
	// 	- This parameter is required for a cross-region backup operation.
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
