// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedBackupPolicyEnabled(v bool) *DescribeBackupPolicyResponseBody
	GetAdvancedBackupPolicyEnabled() *bool
	SetAdvancedDataPolicies(v *DescribeBackupPolicyResponseBodyAdvancedDataPolicies) *DescribeBackupPolicyResponseBody
	GetAdvancedDataPolicies() *DescribeBackupPolicyResponseBodyAdvancedDataPolicies
	SetAdvancedLogPolicies(v *DescribeBackupPolicyResponseBodyAdvancedLogPolicies) *DescribeBackupPolicyResponseBody
	GetAdvancedLogPolicies() *DescribeBackupPolicyResponseBodyAdvancedLogPolicies
	SetArchiveBackupKeepCount(v string) *DescribeBackupPolicyResponseBody
	GetArchiveBackupKeepCount() *string
	SetArchiveBackupKeepPolicy(v string) *DescribeBackupPolicyResponseBody
	GetArchiveBackupKeepPolicy() *string
	SetArchiveBackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody
	GetArchiveBackupRetentionPeriod() *string
	SetBackupInterval(v string) *DescribeBackupPolicyResponseBody
	GetBackupInterval() *string
	SetBackupLog(v string) *DescribeBackupPolicyResponseBody
	GetBackupLog() *string
	SetBackupMethod(v string) *DescribeBackupPolicyResponseBody
	GetBackupMethod() *string
	SetBackupPriority(v int32) *DescribeBackupPolicyResponseBody
	GetBackupPriority() *int32
	SetBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody
	GetBackupRetentionPeriod() *int32
	SetCategory(v string) *DescribeBackupPolicyResponseBody
	GetCategory() *string
	SetCompressType(v string) *DescribeBackupPolicyResponseBody
	GetCompressType() *string
	SetEnableBackupLog(v string) *DescribeBackupPolicyResponseBody
	GetEnableBackupLog() *string
	SetEnableIncrementDataBackup(v bool) *DescribeBackupPolicyResponseBody
	GetEnableIncrementDataBackup() *bool
	SetEnablePitrProtection(v bool) *DescribeBackupPolicyResponseBody
	GetEnablePitrProtection() *bool
	SetHighSpaceUsageProtection(v string) *DescribeBackupPolicyResponseBody
	GetHighSpaceUsageProtection() *string
	SetLocalLogRetentionHours(v int32) *DescribeBackupPolicyResponseBody
	GetLocalLogRetentionHours() *int32
	SetLocalLogRetentionSpace(v string) *DescribeBackupPolicyResponseBody
	GetLocalLogRetentionSpace() *string
	SetLogBackupFrequency(v string) *DescribeBackupPolicyResponseBody
	GetLogBackupFrequency() *string
	SetLogBackupLocalRetentionNumber(v int32) *DescribeBackupPolicyResponseBody
	GetLogBackupLocalRetentionNumber() *int32
	SetLogBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody
	GetLogBackupRetentionPeriod() *int32
	SetPitrRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody
	GetPitrRetentionPeriod() *int32
	SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupTime() *string
	SetPreferredNextBackupTime(v string) *DescribeBackupPolicyResponseBody
	GetPreferredNextBackupTime() *string
	SetReleasedKeepPolicy(v string) *DescribeBackupPolicyResponseBody
	GetReleasedKeepPolicy() *string
	SetRequestId(v string) *DescribeBackupPolicyResponseBody
	GetRequestId() *string
	SetSupportModifyBackupPriority(v bool) *DescribeBackupPolicyResponseBody
	GetSupportModifyBackupPriority() *bool
	SetSupportReleasedKeep(v int32) *DescribeBackupPolicyResponseBody
	GetSupportReleasedKeep() *int32
	SetSupportVolumeShadowCopy(v int32) *DescribeBackupPolicyResponseBody
	GetSupportVolumeShadowCopy() *int32
	SetSupportsHighFrequencyBackup(v int64) *DescribeBackupPolicyResponseBody
	GetSupportsHighFrequencyBackup() *int64
}

type DescribeBackupPolicyResponseBody struct {
	AdvancedBackupPolicyEnabled *bool                                                 `json:"AdvancedBackupPolicyEnabled,omitempty" xml:"AdvancedBackupPolicyEnabled,omitempty"`
	AdvancedDataPolicies        *DescribeBackupPolicyResponseBodyAdvancedDataPolicies `json:"AdvancedDataPolicies,omitempty" xml:"AdvancedDataPolicies,omitempty" type:"Struct"`
	AdvancedLogPolicies         *DescribeBackupPolicyResponseBodyAdvancedLogPolicies  `json:"AdvancedLogPolicies,omitempty" xml:"AdvancedLogPolicies,omitempty" type:"Struct"`
	// The number of archived backup files that are retained.
	//
	// example:
	//
	// 1
	ArchiveBackupKeepCount *string `json:"ArchiveBackupKeepCount,omitempty" xml:"ArchiveBackupKeepCount,omitempty"`
	// The cycle based on which archived backup files are retained.
	//
	// example:
	//
	// ByMonth
	ArchiveBackupKeepPolicy *string `json:"ArchiveBackupKeepPolicy,omitempty" xml:"ArchiveBackupKeepPolicy,omitempty"`
	// The number of days for which archived backup files are retained.
	//
	// example:
	//
	// 365
	ArchiveBackupRetentionPeriod *string `json:"ArchiveBackupRetentionPeriod,omitempty" xml:"ArchiveBackupRetentionPeriod,omitempty"`
	// The backup interval. Unit: minutes.
	//
	// 	- If the instance runs MySQL, the interval is the same as the value of the Snapshot Backup Start Time parameter rather than the Snapshot Backup Period parameter in the ApsaraDB RDS console. For more information, see [Back up an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/98818.html).
	//
	// 	- If the instance runs SQL Server, the interval is the same as the log backup frequency.
	//
	// example:
	//
	// 30
	BackupInterval *string `json:"BackupInterval,omitempty" xml:"BackupInterval,omitempty"`
	// Indicates whether the log backup feature is enabled. Valid values:
	//
	// 	- **Enable**
	//
	// 	- **Disabled**
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
	// > This parameter is returned only when the instance runs SQL Server and uses cloud disks.
	//
	// example:
	//
	// Physical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The backup settings of the secondary instance. Valid values:
	//
	// 	- **1**: Secondary instance preferred
	//
	// 	- **2**: Primary instance preferred
	//
	// >  This parameter is available only for instances that run SQL Server on RDS Cluster Edition. This parameter is returned only when SupportModifyBackupPriority is set to True.
	//
	// example:
	//
	// 2
	BackupPriority *int32 `json:"BackupPriority,omitempty" xml:"BackupPriority,omitempty"`
	// The number of days for which data backup files are retained.
	//
	// example:
	//
	// 7
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// Indicates whether to enable the single-digit second backup feature. This feature allows ApsaraDB RDS to complete a backup within single-digit seconds. Valid values:
	//
	// 	- **Flash**: The single-digit second backup feature is enabled.
	//
	// 	- **Standard**: The single-digit second backup feature is disabled.
	//
	// > This parameter takes effect only when you set the **BackupPolicyMode*	- parameter to **DataBackupPolicy**.
	//
	// example:
	//
	// Standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The method that is used to compress backup data. Valid values:
	//
	// 	- **0**: Backup data is not compressed.
	//
	// 	- **1**: Backup data is compressed by using zlib.
	//
	// 	- **2**: Backup data is compressed by using zlib that invokes more than one thread in parallel for each backup.
	//
	// 	- **4**: Backup data is compressed by using QuickLZ and can be used to restore individual databases or tables.
	//
	// 	- **8**: Backup data is compressed by using QuickLZ but cannot be used to restore individual databases or tables.
	//
	// example:
	//
	// 1
	CompressType *string `json:"CompressType,omitempty" xml:"CompressType,omitempty"`
	// Indicates whether the log backup feature is enabled. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	EnableBackupLog *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// Indicates whether incremental backup is enabled. Valid values:
	//
	// 	- **True**: Incremental backup is enabled.
	//
	// 	- **False**: Incremental backup is disabled.
	//
	// example:
	//
	// True
	EnableIncrementDataBackup *bool `json:"EnableIncrementDataBackup,omitempty" xml:"EnableIncrementDataBackup,omitempty"`
	// Indicates whether the point-in-time restoration (PITR) feature is enabled. The PITR feature is an enhancement of the log backup feature. Valid values:
	//
	// 	- **True**
	//
	// 	- **False**
	//
	// >  This parameter is returned only when the instance runs MySQL. For more information, see [Configure the PITR feature](https://help.aliyun.com/document_detail/2666046.html).
	//
	// example:
	//
	// True
	EnablePitrProtection *bool `json:"EnablePitrProtection,omitempty" xml:"EnablePitrProtection,omitempty"`
	// Indicates whether the log backup deletion feature is enabled. If the disk usage exceeds 80% or the remaining disk space is less than 5 GB on the instance, this feature deletes binary log files. Valid values:
	//
	// 	- **Disable**
	//
	// 	- **Enable**
	//
	// example:
	//
	// Enable
	HighSpaceUsageProtection *string `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	// The number of hours for which log backup files are retained on the instance.
	//
	// example:
	//
	// 0
	LocalLogRetentionHours *int32 `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	// The maximum storage usage that is allowed for log files on the instance.
	//
	// example:
	//
	// 30
	LocalLogRetentionSpace *string `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	// The backup frequency of logs. Valid values:
	//
	// 	- **LogInterval**: Log backups are performed every 30 minutes.
	//
	// 	- Default value: same as the value of the **PreferredBackupPeriod*	- parameter.
	//
	// >  This parameter is returned only when the instance runs SQL Server.
	//
	// example:
	//
	// LogInterval
	LogBackupFrequency *string `json:"LogBackupFrequency,omitempty" xml:"LogBackupFrequency,omitempty"`
	// The number of binary log files that you want to retain on the instance.
	//
	// example:
	//
	// 60
	LogBackupLocalRetentionNumber *int32 `json:"LogBackupLocalRetentionNumber,omitempty" xml:"LogBackupLocalRetentionNumber,omitempty"`
	// The number of days for which log backup files are retained.
	//
	// example:
	//
	// 7
	LogBackupRetentionPeriod *int32 `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	// The number of days during which you can restore data of the instance to any point in time.
	//
	// example:
	//
	// 7
	PitrRetentionPeriod *int32 `json:"PitrRetentionPeriod,omitempty" xml:"PitrRetentionPeriod,omitempty"`
	// The cycle based on which you want to perform a backup. Separate multiple values with commas (,). Valid values:
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
	// example:
	//
	// Monday,Wednesday,Friday,Sunday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The time when a data backup is performed. The time follows the ISO 8601 standard in the *HH:mm*Z-*HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 15:00Z-16:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The time when the next backup is performed. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-19T15:15Z
	PreferredNextBackupTime *string `json:"PreferredNextBackupTime,omitempty" xml:"PreferredNextBackupTime,omitempty"`
	// The policy that is used to retain archived backup files if the instance is released. Valid values:
	//
	// 	- **None**: No archived backup files are retained.
	//
	// 	- **Lastest**: Only the last archived backup file is retained.
	//
	// 	- **All**: All archived backup files are retained.
	//
	// example:
	//
	// None
	ReleasedKeepPolicy *string `json:"ReleasedKeepPolicy,omitempty" xml:"ReleasedKeepPolicy,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B87E2AB3-B7C9-4394-9160-7F639F732031
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the backup settings of a secondary instance can be modified. Valid values:
	//
	// 	- **True**
	//
	// 	- **False**
	//
	// example:
	//
	// False
	SupportModifyBackupPriority *bool `json:"SupportModifyBackupPriority,omitempty" xml:"SupportModifyBackupPriority,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// 0
	SupportReleasedKeep *int32 `json:"SupportReleasedKeep,omitempty" xml:"SupportReleasedKeep,omitempty"`
	// Indicates whether the instance supports snapshot backups. Valid values:
	//
	// 	- **1**: The instance supports snapshot backups.
	//
	// 	- **0**: The instance does not support snapshot backups.
	//
	// >  This parameter is returned only when the instance runs SQL Server.
	//
	// example:
	//
	// 1
	SupportVolumeShadowCopy *int32 `json:"SupportVolumeShadowCopy,omitempty" xml:"SupportVolumeShadowCopy,omitempty"`
	// Indicates whether log backups for SQL Server are performed verery five minutes.
	//
	// 	- 0: No
	//
	// 	- 1: Yes
	//
	// example:
	//
	// 0
	SupportsHighFrequencyBackup *int64 `json:"SupportsHighFrequencyBackup,omitempty" xml:"SupportsHighFrequencyBackup,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) GetAdvancedBackupPolicyEnabled() *bool {
	return s.AdvancedBackupPolicyEnabled
}

func (s *DescribeBackupPolicyResponseBody) GetAdvancedDataPolicies() *DescribeBackupPolicyResponseBodyAdvancedDataPolicies {
	return s.AdvancedDataPolicies
}

func (s *DescribeBackupPolicyResponseBody) GetAdvancedLogPolicies() *DescribeBackupPolicyResponseBodyAdvancedLogPolicies {
	return s.AdvancedLogPolicies
}

func (s *DescribeBackupPolicyResponseBody) GetArchiveBackupKeepCount() *string {
	return s.ArchiveBackupKeepCount
}

func (s *DescribeBackupPolicyResponseBody) GetArchiveBackupKeepPolicy() *string {
	return s.ArchiveBackupKeepPolicy
}

func (s *DescribeBackupPolicyResponseBody) GetArchiveBackupRetentionPeriod() *string {
	return s.ArchiveBackupRetentionPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetBackupInterval() *string {
	return s.BackupInterval
}

func (s *DescribeBackupPolicyResponseBody) GetBackupLog() *string {
	return s.BackupLog
}

func (s *DescribeBackupPolicyResponseBody) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeBackupPolicyResponseBody) GetBackupPriority() *int32 {
	return s.BackupPriority
}

func (s *DescribeBackupPolicyResponseBody) GetBackupRetentionPeriod() *int32 {
	return s.BackupRetentionPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetCategory() *string {
	return s.Category
}

func (s *DescribeBackupPolicyResponseBody) GetCompressType() *string {
	return s.CompressType
}

func (s *DescribeBackupPolicyResponseBody) GetEnableBackupLog() *string {
	return s.EnableBackupLog
}

func (s *DescribeBackupPolicyResponseBody) GetEnableIncrementDataBackup() *bool {
	return s.EnableIncrementDataBackup
}

func (s *DescribeBackupPolicyResponseBody) GetEnablePitrProtection() *bool {
	return s.EnablePitrProtection
}

func (s *DescribeBackupPolicyResponseBody) GetHighSpaceUsageProtection() *string {
	return s.HighSpaceUsageProtection
}

func (s *DescribeBackupPolicyResponseBody) GetLocalLogRetentionHours() *int32 {
	return s.LocalLogRetentionHours
}

func (s *DescribeBackupPolicyResponseBody) GetLocalLogRetentionSpace() *string {
	return s.LocalLogRetentionSpace
}

func (s *DescribeBackupPolicyResponseBody) GetLogBackupFrequency() *string {
	return s.LogBackupFrequency
}

func (s *DescribeBackupPolicyResponseBody) GetLogBackupLocalRetentionNumber() *int32 {
	return s.LogBackupLocalRetentionNumber
}

func (s *DescribeBackupPolicyResponseBody) GetLogBackupRetentionPeriod() *int32 {
	return s.LogBackupRetentionPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetPitrRetentionPeriod() *int32 {
	return s.PitrRetentionPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredNextBackupTime() *string {
	return s.PreferredNextBackupTime
}

func (s *DescribeBackupPolicyResponseBody) GetReleasedKeepPolicy() *string {
	return s.ReleasedKeepPolicy
}

func (s *DescribeBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPolicyResponseBody) GetSupportModifyBackupPriority() *bool {
	return s.SupportModifyBackupPriority
}

func (s *DescribeBackupPolicyResponseBody) GetSupportReleasedKeep() *int32 {
	return s.SupportReleasedKeep
}

func (s *DescribeBackupPolicyResponseBody) GetSupportVolumeShadowCopy() *int32 {
	return s.SupportVolumeShadowCopy
}

func (s *DescribeBackupPolicyResponseBody) GetSupportsHighFrequencyBackup() *int64 {
	return s.SupportsHighFrequencyBackup
}

func (s *DescribeBackupPolicyResponseBody) SetAdvancedBackupPolicyEnabled(v bool) *DescribeBackupPolicyResponseBody {
	s.AdvancedBackupPolicyEnabled = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetAdvancedDataPolicies(v *DescribeBackupPolicyResponseBodyAdvancedDataPolicies) *DescribeBackupPolicyResponseBody {
	s.AdvancedDataPolicies = v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetAdvancedLogPolicies(v *DescribeBackupPolicyResponseBodyAdvancedLogPolicies) *DescribeBackupPolicyResponseBody {
	s.AdvancedLogPolicies = v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetArchiveBackupKeepCount(v string) *DescribeBackupPolicyResponseBody {
	s.ArchiveBackupKeepCount = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetArchiveBackupKeepPolicy(v string) *DescribeBackupPolicyResponseBody {
	s.ArchiveBackupKeepPolicy = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetArchiveBackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.ArchiveBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupInterval(v string) *DescribeBackupPolicyResponseBody {
	s.BackupInterval = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupLog(v string) *DescribeBackupPolicyResponseBody {
	s.BackupLog = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupMethod(v string) *DescribeBackupPolicyResponseBody {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupPriority(v int32) *DescribeBackupPolicyResponseBody {
	s.BackupPriority = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetCategory(v string) *DescribeBackupPolicyResponseBody {
	s.Category = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetCompressType(v string) *DescribeBackupPolicyResponseBody {
	s.CompressType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetEnableBackupLog(v string) *DescribeBackupPolicyResponseBody {
	s.EnableBackupLog = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetEnableIncrementDataBackup(v bool) *DescribeBackupPolicyResponseBody {
	s.EnableIncrementDataBackup = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetEnablePitrProtection(v bool) *DescribeBackupPolicyResponseBody {
	s.EnablePitrProtection = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetHighSpaceUsageProtection(v string) *DescribeBackupPolicyResponseBody {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetLocalLogRetentionHours(v int32) *DescribeBackupPolicyResponseBody {
	s.LocalLogRetentionHours = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetLocalLogRetentionSpace(v string) *DescribeBackupPolicyResponseBody {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetLogBackupFrequency(v string) *DescribeBackupPolicyResponseBody {
	s.LogBackupFrequency = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetLogBackupLocalRetentionNumber(v int32) *DescribeBackupPolicyResponseBody {
	s.LogBackupLocalRetentionNumber = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetLogBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPitrRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.PitrRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredNextBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredNextBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetReleasedKeepPolicy(v string) *DescribeBackupPolicyResponseBody {
	s.ReleasedKeepPolicy = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetSupportModifyBackupPriority(v bool) *DescribeBackupPolicyResponseBody {
	s.SupportModifyBackupPriority = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetSupportReleasedKeep(v int32) *DescribeBackupPolicyResponseBody {
	s.SupportReleasedKeep = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetSupportVolumeShadowCopy(v int32) *DescribeBackupPolicyResponseBody {
	s.SupportVolumeShadowCopy = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetSupportsHighFrequencyBackup(v int64) *DescribeBackupPolicyResponseBody {
	s.SupportsHighFrequencyBackup = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) Validate() error {
	if s.AdvancedDataPolicies != nil {
		if err := s.AdvancedDataPolicies.Validate(); err != nil {
			return err
		}
	}
	if s.AdvancedLogPolicies != nil {
		if err := s.AdvancedLogPolicies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupPolicyResponseBodyAdvancedDataPolicies struct {
	AdvancedDataPolicy []*DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy `json:"AdvancedDataPolicy,omitempty" xml:"AdvancedDataPolicy,omitempty" type:"Repeated"`
}

func (s DescribeBackupPolicyResponseBodyAdvancedDataPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyAdvancedDataPolicies) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPolicies) GetAdvancedDataPolicy() []*DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	return s.AdvancedDataPolicy
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPolicies) SetAdvancedDataPolicy(v []*DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) *DescribeBackupPolicyResponseBodyAdvancedDataPolicies {
	s.AdvancedDataPolicy = v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPolicies) Validate() error {
	if s.AdvancedDataPolicy != nil {
		for _, item := range s.AdvancedDataPolicy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy struct {
	ActionType              *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	BakType                 *string `json:"BakType,omitempty" xml:"BakType,omitempty"`
	DestRegion              *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	DestType                *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	FilterKey               *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterType              *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	FilterValue             *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	OnlyPreserveOneEachDay  *bool   `json:"OnlyPreserveOneEachDay,omitempty" xml:"OnlyPreserveOneEachDay,omitempty"`
	OnlyPreserveOneEachHour *bool   `json:"OnlyPreserveOneEachHour,omitempty" xml:"OnlyPreserveOneEachHour,omitempty"`
	RetentionType           *string `json:"RetentionType,omitempty" xml:"RetentionType,omitempty"`
	RetentionValue          *int32  `json:"RetentionValue,omitempty" xml:"RetentionValue,omitempty"`
	SrcRegion               *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	SrcType                 *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	StrategyId              *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetActionType() *string {
	return s.ActionType
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetBakType() *string {
	return s.BakType
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetDestType() *string {
	return s.DestType
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetFilterKey() *string {
	return s.FilterKey
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetFilterType() *string {
	return s.FilterType
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetFilterValue() *string {
	return s.FilterValue
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetOnlyPreserveOneEachDay() *bool {
	return s.OnlyPreserveOneEachDay
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetOnlyPreserveOneEachHour() *bool {
	return s.OnlyPreserveOneEachHour
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetRetentionType() *string {
	return s.RetentionType
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetRetentionValue() *int32 {
	return s.RetentionValue
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetSrcType() *string {
	return s.SrcType
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) GetStrategyId() *string {
	return s.StrategyId
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetActionType(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.ActionType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetBakType(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.BakType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetDestRegion(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.DestRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetDestType(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.DestType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetFilterKey(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.FilterKey = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetFilterType(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.FilterType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetFilterValue(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.FilterValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetOnlyPreserveOneEachDay(v bool) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.OnlyPreserveOneEachDay = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetOnlyPreserveOneEachHour(v bool) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.OnlyPreserveOneEachHour = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetRetentionType(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.RetentionType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetRetentionValue(v int32) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.RetentionValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetSrcRegion(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.SrcRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetSrcType(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.SrcType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) SetStrategyId(v string) *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy {
	s.StrategyId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedDataPoliciesAdvancedDataPolicy) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupPolicyResponseBodyAdvancedLogPolicies struct {
	AdvancedLogPolicy []*DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy `json:"AdvancedLogPolicy,omitempty" xml:"AdvancedLogPolicy,omitempty" type:"Repeated"`
}

func (s DescribeBackupPolicyResponseBodyAdvancedLogPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyAdvancedLogPolicies) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPolicies) GetAdvancedLogPolicy() []*DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	return s.AdvancedLogPolicy
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPolicies) SetAdvancedLogPolicy(v []*DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) *DescribeBackupPolicyResponseBodyAdvancedLogPolicies {
	s.AdvancedLogPolicy = v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPolicies) Validate() error {
	if s.AdvancedLogPolicy != nil {
		for _, item := range s.AdvancedLogPolicy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy struct {
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

func (s DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetActionType() *string {
	return s.ActionType
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetDestType() *string {
	return s.DestType
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetEnableLogBackup() *int32 {
	return s.EnableLogBackup
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetFilterKey() *string {
	return s.FilterKey
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetFilterValue() *string {
	return s.FilterValue
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetLogRetentionType() *string {
	return s.LogRetentionType
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetLogRetentionValue() *int32 {
	return s.LogRetentionValue
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetSrcType() *string {
	return s.SrcType
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetStrategyId() *string {
	return s.StrategyId
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetActionType(v string) *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.ActionType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetDestRegion(v string) *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.DestRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetDestType(v string) *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.DestType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetEnableLogBackup(v int32) *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.EnableLogBackup = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetFilterKey(v string) *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.FilterKey = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetFilterValue(v string) *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.FilterValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetLogRetentionType(v string) *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.LogRetentionType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetLogRetentionValue(v int32) *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.LogRetentionValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetSrcRegion(v string) *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.SrcRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetSrcType(v string) *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.SrcType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetStrategyId(v string) *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.StrategyId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) Validate() error {
	return dara.Validate(s)
}
