// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchiveBackupKeepCount(v int32) *ModifyBackupPolicyRequest
	GetArchiveBackupKeepCount() *int32
	SetArchiveBackupKeepPolicy(v string) *ModifyBackupPolicyRequest
	GetArchiveBackupKeepPolicy() *string
	SetArchiveBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest
	GetArchiveBackupRetentionPeriod() *string
	SetBackupInterval(v string) *ModifyBackupPolicyRequest
	GetBackupInterval() *string
	SetBackupLog(v string) *ModifyBackupPolicyRequest
	GetBackupLog() *string
	SetBackupMethod(v string) *ModifyBackupPolicyRequest
	GetBackupMethod() *string
	SetBackupPolicyMode(v string) *ModifyBackupPolicyRequest
	GetBackupPolicyMode() *string
	SetBackupPriority(v int32) *ModifyBackupPolicyRequest
	GetBackupPriority() *int32
	SetBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest
	GetBackupRetentionPeriod() *string
	SetCategory(v string) *ModifyBackupPolicyRequest
	GetCategory() *string
	SetCompressType(v string) *ModifyBackupPolicyRequest
	GetCompressType() *string
	SetDBInstanceId(v string) *ModifyBackupPolicyRequest
	GetDBInstanceId() *string
	SetEnableBackupLog(v string) *ModifyBackupPolicyRequest
	GetEnableBackupLog() *string
	SetEnableIncrementDataBackup(v bool) *ModifyBackupPolicyRequest
	GetEnableIncrementDataBackup() *bool
	SetHighSpaceUsageProtection(v string) *ModifyBackupPolicyRequest
	GetHighSpaceUsageProtection() *string
	SetLocalLogRetentionHours(v string) *ModifyBackupPolicyRequest
	GetLocalLogRetentionHours() *string
	SetLocalLogRetentionSpace(v string) *ModifyBackupPolicyRequest
	GetLocalLogRetentionSpace() *string
	SetLogBackupFrequency(v string) *ModifyBackupPolicyRequest
	GetLogBackupFrequency() *string
	SetLogBackupLocalRetentionNumber(v int32) *ModifyBackupPolicyRequest
	GetLogBackupLocalRetentionNumber() *int32
	SetLogBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest
	GetLogBackupRetentionPeriod() *string
	SetOwnerAccount(v string) *ModifyBackupPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyBackupPolicyRequest
	GetOwnerId() *int64
	SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupTime() *string
	SetReleasedKeepPolicy(v string) *ModifyBackupPolicyRequest
	GetReleasedKeepPolicy() *string
	SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest
	GetResourceOwnerId() *int64
}

type ModifyBackupPolicyRequest struct {
	// The number of archived backup files that are retained. Default value: **1**. Valid values:
	//
	// 	- Valid values when **ArchiveBackupKeepPolicy*	- is set to **ByMonth**: **1*	- to **31**.
	//
	// 	- Valid values when **ArchiveBackupKeepPolicy*	- is set to **ByWeek**: **1*	- to **7**.
	//
	// > 	- You do not need to specify this parameter when **ArchiveBackupKeepPolicy*	- is set to **KeepAll**.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// 1
	ArchiveBackupKeepCount *int32 `json:"ArchiveBackupKeepCount,omitempty" xml:"ArchiveBackupKeepCount,omitempty"`
	// The retention period of archived backup files. The number of archived backup files that can be retained within the specified retention period is specified by **ArchiveBackupKeepCount**. Default value: **0**. Valid values:
	//
	// 	- **ByMonth**
	//
	// 	- **ByWeek**
	//
	// 	- **KeepAll**
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// ByMonth
	ArchiveBackupKeepPolicy *string `json:"ArchiveBackupKeepPolicy,omitempty" xml:"ArchiveBackupKeepPolicy,omitempty"`
	// The number of days for which the archived backup is retained. The default value **0*	- specifies that the backup archiving feature is disabled. Valid values: **30*	- to **1095**.
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// 365
	ArchiveBackupRetentionPeriod *string `json:"ArchiveBackupRetentionPeriod,omitempty" xml:"ArchiveBackupRetentionPeriod,omitempty"`
	// The frequency at which you want to perform a snapshot backup on the instance. Valid values:
	//
	// 	- **-1**: No backup frequencies are specified.
	//
	// 	- **30**: A snapshot backup is performed every 30 minutes.
	//
	// 	- **60**: A snapshot backup is performed every 60 minutes.
	//
	// 	- **120**: A snapshot backup is performed every 120 minutes.
	//
	// 	- **240**: A snapshot backup is performed every 240 minutes.
	//
	// 	- **480**: A snapshot backup is performed every 480 minutes.
	//
	// > 	- You can configure a backup policy by using this parameter and the **PreferredBackupPeriod*	- parameter. For example, if you set **PreferredBackupPeriod*	- to Saturday,Sunday and BackupInterval to \\*\\*-1\\*\\*, a snapshot backup is performed on every Saturday and Sunday.
	//
	// > 	- If the instance runs PostgreSQL, BackupInterval is supported only when the instance is equipped with cloud disks.
	//
	// > 	- If the instance runs SQL Server, BackupInterval is supported only when the snapshot backup feature is enabled for the instance. For more information, see [Enable snapshot backups for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/211143.html).
	//
	// > 	- If **Category*	- is set to **Flash**, BackupInterval is invalid.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// 30
	BackupInterval *string `json:"BackupInterval,omitempty" xml:"BackupInterval,omitempty"`
	// Specifies whether to enable the log backup feature. Valid values:
	//
	// 	- **Enable**: enables the feature.
	//
	// 	- **Disabled**: disables the feature.
	//
	// > 	- This parameter must be specified when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// Enable
	BackupLog *string `json:"BackupLog,omitempty" xml:"BackupLog,omitempty"`
	// The backup method of the instance. Valid values:
	//
	// 	- **Physical**: physical backup
	//
	// 	- **Snapshot**: snapshot backup
	//
	// Default value: **Physical**.
	//
	// > 	- This parameter takes effect only on instances that run SQL Server with cloud disks.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// Physical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The type of the backup. Valid values:
	//
	// 	- **DataBackupPolicy**: data backup
	//
	// 	- **LogBackupPolicy**: log backup
	//
	// example:
	//
	// DataBackupPolicy
	BackupPolicyMode *string `json:"BackupPolicyMode,omitempty" xml:"BackupPolicyMode,omitempty"`
	// Specifies whether the backup settings of a secondary instance are configured. Valid values:
	//
	// 	- **1**: secondary instance preferred
	//
	// 	- **2**: primary instance preferred
	//
	// > 	- This parameter is suitable only for instances that run SQL Server on RDS Cluster Edition.
	//
	// > 	- This parameter takes effect only when **BackupMethod*	- is set to **Physical**. If **BackupMethod*	- is set to **Snapshot**, backups are forcefully performed on the primary instance that runs SQL Server on RDS Cluster Edition.
	//
	// example:
	//
	// 2
	BackupPriority *int32 `json:"BackupPriority,omitempty" xml:"BackupPriority,omitempty"`
	// The number of days for which you want to retain data backup files. Valid values: **7 to 730**.
	//
	// > 	- This parameter must be specified when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// 7
	BackupRetentionPeriod *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// Specifies whether to enable the single-digit second backup feature. Valid values:
	//
	// 	- **Flash**: enables the feature.
	//
	// 	- **Standard**: disables the feature.
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// Standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The format that is used to compress backup data. Valid values:
	//
	// 	- **0**: Backups are not compressed.
	//
	// 	- **1**: The zlib tool is used to compress backups into .tar.gz files.
	//
	// 	- **2**: The zlib tool is used to compress backups in parallel.
	//
	// 	- **4**: The QuickLZ tool is used to compress backups into .xb.gz files. This compression format is supported for instances that run MySQL 5.6 or MySQL 5.7. Backups in this compression format can be used to restore individual databases and tables. For more information, see [Restore individual databases and tables of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/103175.html).
	//
	// 	- **8**: The QuickLZ tool is used to compress backups into .xb.gz files. This compression format is supported only for instances that run MySQL 8.0. Backups in this compression format cannot be used to restore individual databases and tables.
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// 4
	CompressType *string `json:"CompressType,omitempty" xml:"CompressType,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to enable the log backup feature. Valid values:
	//
	// 	- **True*	- or **1**: enables the log backup feature.
	//
	// 	- **False*	- or **0**: disables the log backup feature.
	//
	// >
	//
	// 	- You must specify this parameter when you set the **BackupPolicyMode*	- parameter to **LogBackupPolicy**.
	//
	// 	- This parameter takes effect only when you set the **BackupPolicyMode*	- parameter to **LogBackupPolicy**.
	//
	// example:
	//
	// 1
	EnableBackupLog *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// Specifies whether to enable incremental backup. Valid values:
	//
	// 	- **false*	- (default): disables the feature.
	//
	// 	- **true**: enables the feature.
	//
	// > 	- This parameter takes effect only on instances that run SQL Server with cloud disks.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// false
	EnableIncrementDataBackup *bool `json:"EnableIncrementDataBackup,omitempty" xml:"EnableIncrementDataBackup,omitempty"`
	// Specifies whether to forcefully delete log backup files from the instance when the storage usage of the instance exceeds 80% or the amount of remaining storage on the instance is less than 5 GB. Valid values: **Enable and Disable**. You can retain the default value.
	//
	// > 	- You must specify this parameter when you set the **BackupPolicyMode*	- parameter to **LogBackupPolicy**.
	//
	// > 	- This parameter takes effect only when you set the **BackupPolicyMode*	- parameter to **LogBackupPolicy**.
	//
	// example:
	//
	// Enable
	HighSpaceUsageProtection *string `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	// The number of hours for which you want to retain log backup files on the instance. Valid values: **0 to 168**. The value 0 specifies that log backup files are not retained on the instance. The value 168 is calculated based on the following formula: 7 × 24.
	//
	// > 	- This parameter must be specified when **BackupPolicyMode*	- is set to **LogBackupPolicy**.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **LogBackupPolicy**.
	//
	// example:
	//
	// 18
	LocalLogRetentionHours *string `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	// The maximum storage usage that is allowed for log backup files on the instance. If the storage usage for log backup files on the instance exceeds the value of this parameter, the system deletes earlier log backup files until the storage usage falls below the value of this parameter. Valid values:**0 to 50**. You can retain the default value.
	//
	// > 	- This parameter must be specified when **BackupPolicyMode*	- is set to **LogBackupPolicy**.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **LogBackupPolicy**.
	//
	// example:
	//
	// 30
	LocalLogRetentionSpace *string `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	// The frequency at which you want to back up the logs of the instance. Valid values:
	//
	// 	- **LogInterval**: A log backup is performed every 30 minutes.
	//
	// 	- The default value is the same as the data backup frequency.
	//
	// > 	- The value **LogInterval*	- is supported only for instances that run SQL Server.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// LogInterval
	LogBackupFrequency *string `json:"LogBackupFrequency,omitempty" xml:"LogBackupFrequency,omitempty"`
	// The number of binary log files that you want to retain on the instance. Default value: **60**. Valid values: **6*	- to **100**.
	//
	// >
	//
	// 	- This parameter takes effect only when you set the **BackupPolicyMode*	- parameter to **LogBackupPolicy**.
	//
	// 	- If the instance runs MySQL, you can set this parameter to \\*\\*-1\\*\\*. The value \\*\\*-1\\*\\	- specifies that an unlimited number of binary log files can be retained on the instance.
	//
	// example:
	//
	// 60
	LogBackupLocalRetentionNumber *int32 `json:"LogBackupLocalRetentionNumber,omitempty" xml:"LogBackupLocalRetentionNumber,omitempty"`
	// The number of days for which the log backup is retained. Valid values: **7 to 730**. The log backup retention period cannot be longer than the data backup retention period.
	//
	// > 	- If you enable the log backup feature, you can specify the log backup retention period. This parameter is supported for instances that run MySQL and PostgreSQL.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy*	- or **LogBackupPolicy**.
	//
	// example:
	//
	// 7
	LogBackupRetentionPeriod *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The backup cycle. Specify at least two days of the week and separate the days with commas (,). Valid values:
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
	// > 	- You can configure a backup policy by using this parameter and the **BackupInterval*	- parameter. For example, if you set this parameter to Saturday,Sunday and the **BackupInterval*	- parameter to 30, a backup is performed every 30 minutes on every Saturday and Sunday.
	//
	// > 	- This parameter must be specified when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// Monday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The time at which you want to perform a backup. Specify the time in the ISO 8601 standard in the *HH:mm*Z-*HH:mm*Z format. The time must be in UTC.
	//
	// > 	- This parameter must be specified when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// 00:00Z-01:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The policy that is used to retain archived backup files if the instance is released. Valid values:
	//
	// 	- **None**: No archived backup files are retained.
	//
	// 	- **Lastest**: Only the last archived backup file is retained.
	//
	// 	- **All**: All archived backup files are retained.
	//
	// > 	- This parameter takes effect only when you set the **BackupPolicyMode*	- parameter to **DataBackupPolicy**.
	//
	// > 	- If the instance uses cloud disks and was created on or after February 1, 2024, this parameter is automatically set to **Lastest**. If the instance uses local disks in the same scenario, this parameter is automatically set to **None**. For more information, see [Backup for deleted instances](https://help.aliyun.com/document_detail/2836955.html).
	//
	// example:
	//
	// None
	ReleasedKeepPolicy   *string `json:"ReleasedKeepPolicy,omitempty" xml:"ReleasedKeepPolicy,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) GetArchiveBackupKeepCount() *int32 {
	return s.ArchiveBackupKeepCount
}

func (s *ModifyBackupPolicyRequest) GetArchiveBackupKeepPolicy() *string {
	return s.ArchiveBackupKeepPolicy
}

func (s *ModifyBackupPolicyRequest) GetArchiveBackupRetentionPeriod() *string {
	return s.ArchiveBackupRetentionPeriod
}

func (s *ModifyBackupPolicyRequest) GetBackupInterval() *string {
	return s.BackupInterval
}

func (s *ModifyBackupPolicyRequest) GetBackupLog() *string {
	return s.BackupLog
}

func (s *ModifyBackupPolicyRequest) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *ModifyBackupPolicyRequest) GetBackupPolicyMode() *string {
	return s.BackupPolicyMode
}

func (s *ModifyBackupPolicyRequest) GetBackupPriority() *int32 {
	return s.BackupPriority
}

func (s *ModifyBackupPolicyRequest) GetBackupRetentionPeriod() *string {
	return s.BackupRetentionPeriod
}

func (s *ModifyBackupPolicyRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyBackupPolicyRequest) GetCompressType() *string {
	return s.CompressType
}

func (s *ModifyBackupPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyBackupPolicyRequest) GetEnableBackupLog() *string {
	return s.EnableBackupLog
}

func (s *ModifyBackupPolicyRequest) GetEnableIncrementDataBackup() *bool {
	return s.EnableIncrementDataBackup
}

func (s *ModifyBackupPolicyRequest) GetHighSpaceUsageProtection() *string {
	return s.HighSpaceUsageProtection
}

func (s *ModifyBackupPolicyRequest) GetLocalLogRetentionHours() *string {
	return s.LocalLogRetentionHours
}

func (s *ModifyBackupPolicyRequest) GetLocalLogRetentionSpace() *string {
	return s.LocalLogRetentionSpace
}

func (s *ModifyBackupPolicyRequest) GetLogBackupFrequency() *string {
	return s.LogBackupFrequency
}

func (s *ModifyBackupPolicyRequest) GetLogBackupLocalRetentionNumber() *int32 {
	return s.LogBackupLocalRetentionNumber
}

func (s *ModifyBackupPolicyRequest) GetLogBackupRetentionPeriod() *string {
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

func (s *ModifyBackupPolicyRequest) GetReleasedKeepPolicy() *string {
	return s.ReleasedKeepPolicy
}

func (s *ModifyBackupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyBackupPolicyRequest) SetArchiveBackupKeepCount(v int32) *ModifyBackupPolicyRequest {
	s.ArchiveBackupKeepCount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetArchiveBackupKeepPolicy(v string) *ModifyBackupPolicyRequest {
	s.ArchiveBackupKeepPolicy = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetArchiveBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.ArchiveBackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupInterval(v string) *ModifyBackupPolicyRequest {
	s.BackupInterval = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupLog(v string) *ModifyBackupPolicyRequest {
	s.BackupLog = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupMethod(v string) *ModifyBackupPolicyRequest {
	s.BackupMethod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupPolicyMode(v string) *ModifyBackupPolicyRequest {
	s.BackupPolicyMode = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupPriority(v int32) *ModifyBackupPolicyRequest {
	s.BackupPriority = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetCategory(v string) *ModifyBackupPolicyRequest {
	s.Category = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetCompressType(v string) *ModifyBackupPolicyRequest {
	s.CompressType = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDBInstanceId(v string) *ModifyBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableBackupLog(v string) *ModifyBackupPolicyRequest {
	s.EnableBackupLog = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableIncrementDataBackup(v bool) *ModifyBackupPolicyRequest {
	s.EnableIncrementDataBackup = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetHighSpaceUsageProtection(v string) *ModifyBackupPolicyRequest {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetLocalLogRetentionHours(v string) *ModifyBackupPolicyRequest {
	s.LocalLogRetentionHours = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetLocalLogRetentionSpace(v string) *ModifyBackupPolicyRequest {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetLogBackupFrequency(v string) *ModifyBackupPolicyRequest {
	s.LogBackupFrequency = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetLogBackupLocalRetentionNumber(v int32) *ModifyBackupPolicyRequest {
	s.LogBackupLocalRetentionNumber = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetLogBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
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

func (s *ModifyBackupPolicyRequest) SetReleasedKeepPolicy(v string) *ModifyBackupPolicyRequest {
	s.ReleasedKeepPolicy = &v
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

func (s *ModifyBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
