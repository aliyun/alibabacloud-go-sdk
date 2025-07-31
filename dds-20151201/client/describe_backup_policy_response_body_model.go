// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupInterval(v int32) *DescribeBackupPolicyResponseBody
	GetBackupInterval() *int32
	SetBackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody
	GetBackupRetentionPeriod() *string
	SetBackupRetentionPolicyOnClusterDeletion(v int32) *DescribeBackupPolicyResponseBody
	GetBackupRetentionPolicyOnClusterDeletion() *int32
	SetCrossBackupPeriod(v string) *DescribeBackupPolicyResponseBody
	GetCrossBackupPeriod() *string
	SetCrossLogRetentionType(v string) *DescribeBackupPolicyResponseBody
	GetCrossLogRetentionType() *string
	SetCrossLogRetentionValue(v int32) *DescribeBackupPolicyResponseBody
	GetCrossLogRetentionValue() *int32
	SetCrossRetentionType(v string) *DescribeBackupPolicyResponseBody
	GetCrossRetentionType() *string
	SetCrossRetentionValue(v int32) *DescribeBackupPolicyResponseBody
	GetCrossRetentionValue() *int32
	SetDestRegion(v string) *DescribeBackupPolicyResponseBody
	GetDestRegion() *string
	SetEnableBackupLog(v int32) *DescribeBackupPolicyResponseBody
	GetEnableBackupLog() *int32
	SetEnableCrossLogBackup(v int32) *DescribeBackupPolicyResponseBody
	GetEnableCrossLogBackup() *int32
	SetHighFrequencyBackupRetention(v string) *DescribeBackupPolicyResponseBody
	GetHighFrequencyBackupRetention() *string
	SetLogBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody
	GetLogBackupRetentionPeriod() *int32
	SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupTime() *string
	SetPreferredNextBackupTime(v string) *DescribeBackupPolicyResponseBody
	GetPreferredNextBackupTime() *string
	SetPreserveOneEachHour(v bool) *DescribeBackupPolicyResponseBody
	GetPreserveOneEachHour() *bool
	SetRequestId(v string) *DescribeBackupPolicyResponseBody
	GetRequestId() *string
	SetSnapshotBackupType(v string) *DescribeBackupPolicyResponseBody
	GetSnapshotBackupType() *string
	SetSrcRegion(v string) *DescribeBackupPolicyResponseBody
	GetSrcRegion() *string
}

type DescribeBackupPolicyResponseBody struct {
	// The frequency at which high-frequency backup is created. Valid values:
	//
	// 	- **-1**: High-frequency backup is disabled.
	//
	// 	- **15**: every 15 minutes.
	//
	// 	- **30**: every 30 minutes.
	//
	// 	- **60**: every hour.
	//
	// 	- **120**: every 2 hours.
	//
	// 	- **180**: every 3 hours.
	//
	// 	- **240**: every 4 hours.
	//
	// 	- **360**: every 6 hours.
	//
	// 	- **480**: every 8 hours.
	//
	// 	- **720**: every 12 hours.
	//
	// example:
	//
	// -1
	BackupInterval *int32 `json:"BackupInterval,omitempty" xml:"BackupInterval,omitempty"`
	// The retention period of the backup data. Unit: day.
	//
	// example:
	//
	// 30
	BackupRetentionPeriod *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The backup retention policy configured for the instance. Valid values:
	//
	// 1.  0: All backup sets are immediately deleted when the instance is released.
	//
	// 2.  1: Automatic backup is performed and the backup set is retained for a long period of time when the instance is released.
	//
	// 3.  2: Automatic backup is performed and all backup sets are retained for a long period of time when the instance is released.
	//
	// For more information, see [Retain the backup files of an ApsaraDB for MongoDB instance for a long period of time](https://help.aliyun.com/document_detail/2779111.html).
	//
	// example:
	//
	// 0
	BackupRetentionPolicyOnClusterDeletion *int32 `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// The retention period of Cross-regional backup.
	//
	// Valid values:
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
	// Monday
	CrossBackupPeriod *string `json:"CrossBackupPeriod,omitempty" xml:"CrossBackupPeriod,omitempty"`
	// The retention type of Cross-regional  log backup.
	//
	// - delay : retain the backup for a period of time.
	//
	// - never : retain the backup permanently.
	//
	// example:
	//
	// delay
	CrossLogRetentionType *string `json:"CrossLogRetentionType,omitempty" xml:"CrossLogRetentionType,omitempty"`
	// The retention time of Cross-regional log backup.
	//
	// example:
	//
	// 7
	CrossLogRetentionValue *int32 `json:"CrossLogRetentionValue,omitempty" xml:"CrossLogRetentionValue,omitempty"`
	// The retention type of Cross-regional backup.
	//
	// - delay : retain the backup for a period of time.
	//
	// - never : retain the backup permanently.
	//
	// example:
	//
	// delay
	CrossRetentionType *string `json:"CrossRetentionType,omitempty" xml:"CrossRetentionType,omitempty"`
	// The retention time of Cross-regional backup.
	//
	// example:
	//
	// 7
	CrossRetentionValue *int32 `json:"CrossRetentionValue,omitempty" xml:"CrossRetentionValue,omitempty"`
	// The region ID of the cross-regional backup..
	//
	// example:
	//
	// cn-shenzhen
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// Indicates whether the log backup feature is enabled. Valid values:
	//
	// 	- **0*	- (default): The log backup feature is disabled.
	//
	// 	- **1**: The log backup feature is enabled.
	//
	// example:
	//
	// 1
	EnableBackupLog *int32 `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// Whether to turn on cross-regional log backup.
	//
	// - 1: turn on . Used for sharded cluster.
	//
	// - 0: turn off. Used for replicate set.
	//
	// example:
	//
	// 1
	EnableCrossLogBackup *int32 `json:"EnableCrossLogBackup,omitempty" xml:"EnableCrossLogBackup,omitempty"`
	// The retention period of high-frequency backups. Unit: day.
	//
	// example:
	//
	// 1
	HighFrequencyBackupRetention *string `json:"HighFrequencyBackupRetention,omitempty" xml:"HighFrequencyBackupRetention,omitempty"`
	// The number of days for which log backups are retained. Valid values: 7 to 730.
	//
	// example:
	//
	// 7
	LogBackupRetentionPeriod *int32 `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	// The day of a week on which to back up data. Valid values:
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
	// Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The time range during which the backup was created. The time follows the ISO 8601 standard in the *HH:mm*Z-*HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 09:00Z-10:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The time of next standard backup.
	//
	// example:
	//
	// 2024-06-19T19:11Z
	PreferredNextBackupTime *string `json:"PreferredNextBackupTime,omitempty" xml:"PreferredNextBackupTime,omitempty"`
	PreserveOneEachHour     *bool   `json:"PreserveOneEachHour,omitempty" xml:"PreserveOneEachHour,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5A9464CA-F7DC-5434-90B1-DF7F197C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// cn-hangzhou
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) GetBackupInterval() *int32 {
	return s.BackupInterval
}

func (s *DescribeBackupPolicyResponseBody) GetBackupRetentionPeriod() *string {
	return s.BackupRetentionPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetBackupRetentionPolicyOnClusterDeletion() *int32 {
	return s.BackupRetentionPolicyOnClusterDeletion
}

func (s *DescribeBackupPolicyResponseBody) GetCrossBackupPeriod() *string {
	return s.CrossBackupPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetCrossLogRetentionType() *string {
	return s.CrossLogRetentionType
}

func (s *DescribeBackupPolicyResponseBody) GetCrossLogRetentionValue() *int32 {
	return s.CrossLogRetentionValue
}

func (s *DescribeBackupPolicyResponseBody) GetCrossRetentionType() *string {
	return s.CrossRetentionType
}

func (s *DescribeBackupPolicyResponseBody) GetCrossRetentionValue() *int32 {
	return s.CrossRetentionValue
}

func (s *DescribeBackupPolicyResponseBody) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribeBackupPolicyResponseBody) GetEnableBackupLog() *int32 {
	return s.EnableBackupLog
}

func (s *DescribeBackupPolicyResponseBody) GetEnableCrossLogBackup() *int32 {
	return s.EnableCrossLogBackup
}

func (s *DescribeBackupPolicyResponseBody) GetHighFrequencyBackupRetention() *string {
	return s.HighFrequencyBackupRetention
}

func (s *DescribeBackupPolicyResponseBody) GetLogBackupRetentionPeriod() *int32 {
	return s.LogBackupRetentionPeriod
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

func (s *DescribeBackupPolicyResponseBody) GetPreserveOneEachHour() *bool {
	return s.PreserveOneEachHour
}

func (s *DescribeBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPolicyResponseBody) GetSnapshotBackupType() *string {
	return s.SnapshotBackupType
}

func (s *DescribeBackupPolicyResponseBody) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeBackupPolicyResponseBody) SetBackupInterval(v int32) *DescribeBackupPolicyResponseBody {
	s.BackupInterval = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPolicyOnClusterDeletion(v int32) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetCrossBackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.CrossBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetCrossLogRetentionType(v string) *DescribeBackupPolicyResponseBody {
	s.CrossLogRetentionType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetCrossLogRetentionValue(v int32) *DescribeBackupPolicyResponseBody {
	s.CrossLogRetentionValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetCrossRetentionType(v string) *DescribeBackupPolicyResponseBody {
	s.CrossRetentionType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetCrossRetentionValue(v int32) *DescribeBackupPolicyResponseBody {
	s.CrossRetentionValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetDestRegion(v string) *DescribeBackupPolicyResponseBody {
	s.DestRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetEnableBackupLog(v int32) *DescribeBackupPolicyResponseBody {
	s.EnableBackupLog = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetEnableCrossLogBackup(v int32) *DescribeBackupPolicyResponseBody {
	s.EnableCrossLogBackup = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetHighFrequencyBackupRetention(v string) *DescribeBackupPolicyResponseBody {
	s.HighFrequencyBackupRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetLogBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.LogBackupRetentionPeriod = &v
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

func (s *DescribeBackupPolicyResponseBody) SetPreserveOneEachHour(v bool) *DescribeBackupPolicyResponseBody {
	s.PreserveOneEachHour = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetSnapshotBackupType(v string) *DescribeBackupPolicyResponseBody {
	s.SnapshotBackupType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetSrcRegion(v string) *DescribeBackupPolicyResponseBody {
	s.SrcRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
