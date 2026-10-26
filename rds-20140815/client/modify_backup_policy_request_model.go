// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedDataPolicies(v []*ModifyBackupPolicyRequestAdvancedDataPolicies) *ModifyBackupPolicyRequest
	GetAdvancedDataPolicies() []*ModifyBackupPolicyRequestAdvancedDataPolicies
	SetAdvancedLogPolicies(v []*ModifyBackupPolicyRequestAdvancedLogPolicies) *ModifyBackupPolicyRequest
	GetAdvancedLogPolicies() []*ModifyBackupPolicyRequestAdvancedLogPolicies
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
	SetEnableAdvancedBackupPolicy(v int32) *ModifyBackupPolicyRequest
	GetEnableAdvancedBackupPolicy() *int32
	SetEnableBackupLog(v string) *ModifyBackupPolicyRequest
	GetEnableBackupLog() *string
	SetEnableIncrementDataBackup(v bool) *ModifyBackupPolicyRequest
	GetEnableIncrementDataBackup() *bool
	SetEnablePitrProtection(v bool) *ModifyBackupPolicyRequest
	GetEnablePitrProtection() *bool
	SetHighSpaceUsageProtection(v string) *ModifyBackupPolicyRequest
	GetHighSpaceUsageProtection() *string
	SetIncBackupInterval(v int32) *ModifyBackupPolicyRequest
	GetIncBackupInterval() *int32
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
	// The list of advanced data backup policies.
	AdvancedDataPolicies []*ModifyBackupPolicyRequestAdvancedDataPolicies `json:"AdvancedDataPolicies,omitempty" xml:"AdvancedDataPolicies,omitempty" type:"Repeated"`
	// The list of advanced log backup policies.
	AdvancedLogPolicies []*ModifyBackupPolicyRequestAdvancedLogPolicies `json:"AdvancedLogPolicies,omitempty" xml:"AdvancedLogPolicies,omitempty" type:"Repeated"`
	// The number of archived backups to retain. Default value: **1**. Valid values:
	//
	// 	- When **ArchiveBackupKeepPolicy*	- is set to **ByMonth**, valid values are **1*	- to **31**.
	//
	// 	- When **ArchiveBackupKeepPolicy*	- is set to **ByWeek**, valid values are **1*	- to **7**.
	//
	// > 	- When **ArchiveBackupKeepPolicy*	- is set to **KeepAll**, you do not need to specify this parameter.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// 1
	ArchiveBackupKeepCount *int32 `json:"ArchiveBackupKeepCount,omitempty" xml:"ArchiveBackupKeepCount,omitempty"`
	// The retention period of archived backups. The number of archived backups that can be retained within this period is determined by **ArchiveBackupKeepCount**. Default value: **0**. Valid values:
	//
	// 	- **ByMonth**: month.
	//
	// 	- **ByWeek**: week.
	//
	// 	- **KeepAll**: all archived backups are retained.
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// ByMonth
	ArchiveBackupKeepPolicy *string `json:"ArchiveBackupKeepPolicy,omitempty" xml:"ArchiveBackupKeepPolicy,omitempty"`
	// The number of days for which archived backups are retained. Default value: **0**, which indicates that archive backup is disabled. Valid values: **30*	- to **1095**.
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// 365
	ArchiveBackupRetentionPeriod *string `json:"ArchiveBackupRetentionPeriod,omitempty" xml:"ArchiveBackupRetentionPeriod,omitempty"`
	// The snapshot backup frequency. Valid values:
	//
	// 	- **15**: 15 minutes.
	//
	// 	- **30**: 30 minutes.
	//
	// 	- **60**: 60 minutes.
	//
	// 	- **120**: 120 minutes.
	//
	// 	- **180**: 180 minutes.
	//
	// 	- **240**: 240 minutes.
	//
	// 	- **360**: 360 minutes.
	//
	// 	- **480**: 480 minutes.
	//
	// 	- **720**: 720 minutes.
	//
	// > 	- This parameter and the **PreferredBackupPeriod*	- parameter determine the backup policy together.
	//
	// > 	- The MySQL instance must be a cloud disk instance of MySQL 5.7 or 8.0 **High-availability Edition*	- or **Cluster Edition**.
	//
	// > 	- The PostgreSQL instance must be a cloud disk instance.
	//
	// > 	- [**Snapshot backup**](https://help.aliyun.com/document_detail/211143.html) must be **enabled*	- for the SQL Server instance.
	//
	// > 	- This parameter does not take effect when **Category*	- is set to **Flash**.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// 30
	BackupInterval *string `json:"BackupInterval,omitempty" xml:"BackupInterval,omitempty"`
	// Specifies whether to enable log backup. Valid values:
	//
	// 	- **Enable**: enables log backup.
	//
	// 	- **Disabled**: disables log backup.
	//
	// **For ApsaraDB RDS for SQL Server instances**, log backup is enabled by default and cannot be disabled. However, you can change the log backup frequency as follows:
	//
	// - **Every 5 minutes**: Set BackupLog to Enable and leave LogBackupFrequency empty. For more information, see [5-minute log backup](https://help.aliyun.com/document_detail/2861729.html). **This configuration is not supported when the preferred secondary instance is configured for backup (BackupPriority is set to 1). Otherwise, an error is reported.**
	//
	// - **Every 30 minutes**: Leave BackupLog empty and set LogBackupFrequency to LogInterval.
	//
	// - **Same as data backup frequency**: Leave both BackupLog and LogBackupFrequency empty.
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**. It is used to enable or disable log backup.
	//
	// example:
	//
	// Enable
	BackupLog *string `json:"BackupLog,omitempty" xml:"BackupLog,omitempty"`
	// The backup method for **ApsaraDB RDS for SQL Server instances with cloud disks**. Valid values:
	//
	// 	- **Physical*	- (default): physical backup.
	//
	// 	- **Snapshot**: snapshot backup.
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// Physical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The backup type. Valid values:
	//
	// 	- **DataBackupPolicy**: data backup.
	//
	// 	- **LogBackupPolicy**: log backup.
	//
	// example:
	//
	// DataBackupPolicy
	BackupPolicyMode *string `json:"BackupPolicyMode,omitempty" xml:"BackupPolicyMode,omitempty"`
	// The backup priority setting for [secondary instance backup](https://help.aliyun.com/document_detail/95717.html) of **ApsaraDB RDS for SQL Server Cluster Edition*	- instances. Valid values:
	//
	// - **1**: prefer secondary instance.
	//
	// - **2**: force primary instance.
	//
	//
	//
	// > - This parameter takes effect only when **BackupMethod*	- is set to **Physical**. If **BackupMethod*	- is set to **Snapshot**, backups are forcibly performed on the primary instance for ApsaraDB RDS for SQL Server Cluster Edition instances.
	//
	// > - After you set the backup priority to **prefer secondary instance*	- (BackupPriority is set to 1), the **5-minute log backup*	- policy (BackupLog is set to Enable and LogBackupFrequency is left empty) is **not supported**. Otherwise, an error is reported. Set the log backup frequency to every 30 minutes or the same as the data backup frequency.
	//
	// example:
	//
	// 2
	BackupPriority *int32 `json:"BackupPriority,omitempty" xml:"BackupPriority,omitempty"`
	// The number of days for which data backups are retained. Valid values: **7*	- to **730**.
	//
	// > 	- This parameter is required when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// 7
	BackupRetentionPeriod *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// Specifies whether to enable backup within seconds. Valid values:
	//
	// 	- **Flash**: enables backup within seconds.
	//
	// 	- **Standard**: disables backup within seconds.
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// Standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The backup compression method. Valid values:
	//
	// 	- **0**: no compression.
	//
	// 	- **1**: zlib compression. The format is tar.gz.
	//
	// 	- **2**: parallel zlib compression.
	//
	// 	- **4**: quicklz compression. The format is xb.gz. This method applies only to MySQL 5.6 and 5.7. You can use this compression method to [restore individual databases and tables](https://help.aliyun.com/document_detail/103175.html).
	//
	// 	- **8**: quicklz compression. The format is xb.gz. This method applies only to MySQL 8.0. Restoration of individual databases and tables is not supported.
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// 4
	CompressType *string `json:"CompressType,omitempty" xml:"CompressType,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to enable the advanced backup policy. If you set this parameter to 1, both the advanced data backup policy and the advanced log backup policy are configured.
	//
	// example:
	//
	// 1
	EnableAdvancedBackupPolicy *int32 `json:"EnableAdvancedBackupPolicy,omitempty" xml:"EnableAdvancedBackupPolicy,omitempty"`
	// Specifies whether to enable log backup for **MySQL**, **PostgreSQL**, and **MariaDB*	- instances. Valid values:
	//
	// 	- **True*	- or **1**: enabled.
	//
	// 	- **False*	- or **0**: shutdown.
	//
	// > - Instance log backup is enabled by default and cannot be shut down for **SQL Server*	- instances. Therefore, you do not need to configure this parameter for SQL Server instances.
	//
	// > - This parameter takes effect only when **BackupPolicyMode*	- is set to **LogBackupPolicy**. It is used to enable or shutdown instance log backup.
	//
	// example:
	//
	// 1
	EnableBackupLog *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// Specifies whether to enable incremental backup for **ApsaraDB RDS for SQL Server instances with cloud disks*	- and **ApsaraDB RDS for MySQL instances with local SSDs**. Valid values:
	//
	// 	- **False*	- (default): disables incremental backup.
	//
	// 	- **True**: enables incremental backup.
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// False
	EnableIncrementDataBackup *bool `json:"EnableIncrementDataBackup,omitempty" xml:"EnableIncrementDataBackup,omitempty"`
	// Specifies whether to enable point-in-time recovery (PITR) for **MySQL*	- instances. Valid values:
	//
	// - **True**: enabled.
	//
	// - **False**: shutdown.
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy*	- and **BackupLog*	- is set to **Enable**. For feature details, see [Settings for point-in-time recovery policy](https://help.aliyun.com/document_detail/2666046.html).
	//
	// example:
	//
	// True
	EnablePitrProtection *bool `json:"EnablePitrProtection,omitempty" xml:"EnablePitrProtection,omitempty"`
	// Specifies whether to unconditionally delete binary log files when the storage usage of the **ApsaraDB RDS for MySQL*	- instance exceeds 80% or the remaining storage space is less than 5 GB. Valid values: **Enable*	- and **Disable**. By default, this parameter is not modified.
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **LogBackupPolicy**. In this case, this parameter is required.
	//
	// example:
	//
	// Enable
	HighSpaceUsageProtection *string `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	// The high-frequency incremental backup frequency for **ApsaraDB RDS for MySQL instances with local SSDs**. Valid values:
	//
	// 	- **60**: 60 minutes.
	//
	// 	- **120**: 120 minutes.
	//
	// 	- **240**: 240 minutes.
	//
	// 	- **360**: 360 minutes.
	//
	// 	- **720**: 720 minutes.
	//
	// > This parameter takes effect only when **EnableIncrementDataBackup*	- is set to **True**.
	//
	// example:
	//
	// 120
	IncBackupInterval *int32 `json:"IncBackupInterval,omitempty" xml:"IncBackupInterval,omitempty"`
	// The number of hours for which instance log backups are retained locally on the **MySQL*	- instance. Valid values: **0*	- to **168*	- (7 × 24). A value of 0 indicates that instance log backups are not retained locally.
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **LogBackupPolicy**. In this case, this parameter is required.
	//
	// example:
	//
	// 18
	LocalLogRetentionHours *string `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	// The maximum usage of the local log storage space on the **ApsaraDB RDS for MySQL*	- instance. If the usage exceeds this value, the earliest binary log files are deleted until the usage drops below this value. Valid values: **0*	- to **50**. By default, this parameter is not modified.
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **LogBackupPolicy**. In this case, this parameter is required.
	//
	// example:
	//
	// 30
	LocalLogRetentionSpace *string `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	// The log backup frequency for **ApsaraDB RDS for SQL Server*	- instances. Valid values:
	//
	// 	- **LogInterval**: every 30 minutes.
	//
	// 	- **Empty*	- (no value required): every 5 minutes or same as the data backup frequency.
	//
	// > This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// LogInterval
	LogBackupFrequency *string `json:"LogBackupFrequency,omitempty" xml:"LogBackupFrequency,omitempty"`
	// The number of binary log files retained locally. Default value: **60**. Valid values: **6*	- to **100**.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **LogBackupPolicy**.
	//
	// > 	- If the instance is an ApsaraDB RDS for MySQL instance, you can set this parameter to -1, which indicates that the number of locally retained binary log files is not limited.
	//
	// example:
	//
	// 60
	LogBackupLocalRetentionNumber *int32 `json:"LogBackupLocalRetentionNumber,omitempty" xml:"LogBackupLocalRetentionNumber,omitempty"`
	// The number of days for which log backups are retained. Valid values: **7*	- to **730**. The value cannot be greater than the number of days for which data backups are retained.
	//
	// > 	- When log backup is enabled, you can specify the number of days for which log backup files are retained. This parameter is supported only for ApsaraDB RDS for MySQL and ApsaraDB RDS for PostgreSQL instances.
	//
	// > 	- This parameter takes effect when **BackupPolicyMode*	- is set to **DataBackupPolicy*	- or **LogBackupPolicy**.
	//
	// example:
	//
	// 7
	LogBackupRetentionPeriod *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The backup cycle. You must specify at least two days. Separate multiple values with commas (,). Valid values:
	//
	// 	- **Monday**: Monday.
	//
	// 	- **Tuesday**: Tuesday.
	//
	// 	- **Wednesday**: Wednesday.
	//
	// 	- **Thursday**: Thursday.
	//
	// 	- **Friday**: Friday.
	//
	// 	- **Saturday**: Saturday.
	//
	// 	- **Sunday**: Sunday.
	//
	// > 	- This parameter and the **BackupInterval*	- parameter determine the backup policy together. For example, if you set this parameter to Saturday and Sunday and set **BackupInterval*	- to 30 minutes, a backup is performed every 30 minutes on Saturdays and Sundays.
	//
	// > 	- This parameter is required when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// Monday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The time period during which the backup task is executed. Specify the time in the <i>HH:mm</i>Z-<i>HH:mm</i>Z format. The time must be in UTC.
	//
	// > 	- This parameter is required when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// > 	- This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// example:
	//
	// 00:00Z-01:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The data retention policy for archived backups of deleted **MySQL*	- instances. Valid values:
	//
	// 	- **None**: does not retain archived backups.
	//
	// 	- **Lastest**: retains the last archived backup.
	//
	// 	- **All**: retains all archived backups.
	//
	// > - This parameter takes effect only when **BackupPolicyMode*	- is set to **DataBackupPolicy**.
	//
	// > - For ApsaraDB RDS for MySQL cloud disk instances purchased on or after February 1, 2024, the default value of ReleasedKeepPolicy is **Lastest**. For instances with Premium Local SSDs, the default value is **None**. For feature details, see [Backups of deleted instances](https://help.aliyun.com/document_detail/2836955.html).
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

func (s *ModifyBackupPolicyRequest) GetAdvancedDataPolicies() []*ModifyBackupPolicyRequestAdvancedDataPolicies {
	return s.AdvancedDataPolicies
}

func (s *ModifyBackupPolicyRequest) GetAdvancedLogPolicies() []*ModifyBackupPolicyRequestAdvancedLogPolicies {
	return s.AdvancedLogPolicies
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

func (s *ModifyBackupPolicyRequest) GetEnableAdvancedBackupPolicy() *int32 {
	return s.EnableAdvancedBackupPolicy
}

func (s *ModifyBackupPolicyRequest) GetEnableBackupLog() *string {
	return s.EnableBackupLog
}

func (s *ModifyBackupPolicyRequest) GetEnableIncrementDataBackup() *bool {
	return s.EnableIncrementDataBackup
}

func (s *ModifyBackupPolicyRequest) GetEnablePitrProtection() *bool {
	return s.EnablePitrProtection
}

func (s *ModifyBackupPolicyRequest) GetHighSpaceUsageProtection() *string {
	return s.HighSpaceUsageProtection
}

func (s *ModifyBackupPolicyRequest) GetIncBackupInterval() *int32 {
	return s.IncBackupInterval
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

func (s *ModifyBackupPolicyRequest) SetAdvancedDataPolicies(v []*ModifyBackupPolicyRequestAdvancedDataPolicies) *ModifyBackupPolicyRequest {
	s.AdvancedDataPolicies = v
	return s
}

func (s *ModifyBackupPolicyRequest) SetAdvancedLogPolicies(v []*ModifyBackupPolicyRequestAdvancedLogPolicies) *ModifyBackupPolicyRequest {
	s.AdvancedLogPolicies = v
	return s
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

func (s *ModifyBackupPolicyRequest) SetEnableAdvancedBackupPolicy(v int32) *ModifyBackupPolicyRequest {
	s.EnableAdvancedBackupPolicy = &v
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

func (s *ModifyBackupPolicyRequest) SetEnablePitrProtection(v bool) *ModifyBackupPolicyRequest {
	s.EnablePitrProtection = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetHighSpaceUsageProtection(v string) *ModifyBackupPolicyRequest {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetIncBackupInterval(v int32) *ModifyBackupPolicyRequest {
	s.IncBackupInterval = &v
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
	if s.AdvancedDataPolicies != nil {
		for _, item := range s.AdvancedDataPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AdvancedLogPolicies != nil {
		for _, item := range s.AdvancedLogPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyBackupPolicyRequestAdvancedDataPolicies struct {
	ActionType  *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	BakType     *string `json:"BakType,omitempty" xml:"BakType,omitempty"`
	DestRegion  *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	DestType    *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	FilterKey   *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterType  *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	// Specifies whether to retain only one data backup per day.
	OnlyPreserveOneEachDay *bool `json:"OnlyPreserveOneEachDay,omitempty" xml:"OnlyPreserveOneEachDay,omitempty"`
	// Specifies whether to retain only one data backup per hour.
	OnlyPreserveOneEachHour *bool   `json:"OnlyPreserveOneEachHour,omitempty" xml:"OnlyPreserveOneEachHour,omitempty"`
	RetentionType           *string `json:"RetentionType,omitempty" xml:"RetentionType,omitempty"`
	RetentionValue          *int32  `json:"RetentionValue,omitempty" xml:"RetentionValue,omitempty"`
	SrcRegion               *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	SrcType                 *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	StrategyId              *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s ModifyBackupPolicyRequestAdvancedDataPolicies) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequestAdvancedDataPolicies) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetActionType() *string {
	return s.ActionType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetBakType() *string {
	return s.BakType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetDestType() *string {
	return s.DestType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetFilterKey() *string {
	return s.FilterKey
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetFilterType() *string {
	return s.FilterType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetFilterValue() *string {
	return s.FilterValue
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetOnlyPreserveOneEachDay() *bool {
	return s.OnlyPreserveOneEachDay
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetOnlyPreserveOneEachHour() *bool {
	return s.OnlyPreserveOneEachHour
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetRetentionType() *string {
	return s.RetentionType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetRetentionValue() *int32 {
	return s.RetentionValue
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetStrategyId() *string {
	return s.StrategyId
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetActionType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.ActionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetBakType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.BakType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetDestRegion(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.DestRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetDestType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.DestType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetFilterKey(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.FilterKey = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetFilterType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.FilterType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetFilterValue(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.FilterValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetOnlyPreserveOneEachDay(v bool) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.OnlyPreserveOneEachDay = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetOnlyPreserveOneEachHour(v bool) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.OnlyPreserveOneEachHour = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetRetentionType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.RetentionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetRetentionValue(v int32) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.RetentionValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetSrcRegion(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.SrcRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetSrcType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.SrcType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetStrategyId(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.StrategyId = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) Validate() error {
	return dara.Validate(s)
}

type ModifyBackupPolicyRequestAdvancedLogPolicies struct {
	ActionType        *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	DestRegion        *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	DestType          *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	EnableLogBackup   *int32  `json:"EnableLogBackup,omitempty" xml:"EnableLogBackup,omitempty"`
	FilterKey         *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterValue       *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	LogRetentionType  *string `json:"LogRetentionType,omitempty" xml:"LogRetentionType,omitempty"`
	LogRetentionValue *int32  `json:"LogRetentionValue,omitempty" xml:"LogRetentionValue,omitempty"`
	SrcRegion         *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	SrcType           *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	StrategyId        *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s ModifyBackupPolicyRequestAdvancedLogPolicies) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequestAdvancedLogPolicies) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetActionType() *string {
	return s.ActionType
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetDestType() *string {
	return s.DestType
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetEnableLogBackup() *int32 {
	return s.EnableLogBackup
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetFilterKey() *string {
	return s.FilterKey
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetFilterValue() *string {
	return s.FilterValue
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetLogRetentionType() *string {
	return s.LogRetentionType
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetLogRetentionValue() *int32 {
	return s.LogRetentionValue
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetStrategyId() *string {
	return s.StrategyId
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetActionType(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.ActionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetDestRegion(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.DestRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetDestType(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.DestType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetEnableLogBackup(v int32) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.EnableLogBackup = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetFilterKey(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.FilterKey = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetFilterValue(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.FilterValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetLogRetentionType(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.LogRetentionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetLogRetentionValue(v int32) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.LogRetentionValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetSrcRegion(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.SrcRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetSrcType(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.SrcType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetStrategyId(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.StrategyId = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) Validate() error {
	return dara.Validate(s)
}
