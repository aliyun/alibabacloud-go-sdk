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
	// The high-frequency backup frequency. Valid values:
	//
	// - **-1**: high-frequency backup is disabled
	//
	// - **15**: every 15 minutes
	//
	// - **30**: every 30 minutes
	//
	// - **60**: every 1 hour
	//
	// - **120**: every 2 hours
	//
	// - **180**: every 3 hours
	//
	// - **240**: every 4 hours
	//
	// - **360**: every 6 hours
	//
	// - **480**: every 8 hours
	//
	// - **720**: every 12 hours
	//
	// example:
	//
	// -1
	BackupInterval *int32 `json:"BackupInterval,omitempty" xml:"BackupInterval,omitempty"`
	// The backup retention period in days.
	//
	// example:
	//
	// 30
	BackupRetentionPeriod *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The backup data retention policy. Valid values:
	//
	// - 0: All backup sets of the instance are immediately deleted when the instance is released.
	//
	// - 1: An automatic backup is created when the instance is released, and the backup set is retained for long-term retention (LTR).
	//
	// - 2: An automatic backup is created when the instance is released, and all backup sets of the cluster are retained for long-term retention (LTR).
	//
	// For more information, see [Long-term retention backup](https://help.aliyun.com/document_detail/2779111.html).
	//
	// example:
	//
	// 0
	BackupRetentionPolicyOnClusterDeletion *int32 `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// The days of the week when geo-redundancy backups are performed. Valid values:
	//
	// - Monday: Monday
	//
	// - Tuesday: Tuesday
	//
	// - Wednesday: Wednesday
	//
	// - Thursday: Thursday
	//
	// - Friday: Friday
	//
	// - Saturday: Saturday
	//
	// - Sunday: Sunday
	//
	// > Required for geo-redundancy backups.
	//
	// example:
	//
	// Monday
	CrossBackupPeriod *string `json:"CrossBackupPeriod,omitempty" xml:"CrossBackupPeriod,omitempty"`
	// The retention type for cross-region log backups. Valid values:
	//
	// - delay: retained for a specified period
	//
	// - never: never expires
	//
	// example:
	//
	// delay
	CrossLogRetentionType *string `json:"CrossLogRetentionType,omitempty" xml:"CrossLogRetentionType,omitempty"`
	// The retention period for cross-region log backups.
	//
	// example:
	//
	// 7
	CrossLogRetentionValue *int32 `json:"CrossLogRetentionValue,omitempty" xml:"CrossLogRetentionValue,omitempty"`
	// The retention type for geo-redundancy backups. Valid values:
	//
	// - delay: retained for a specified period
	//
	// - never: never expires
	//
	// example:
	//
	// delay
	CrossRetentionType *string `json:"CrossRetentionType,omitempty" xml:"CrossRetentionType,omitempty"`
	// The retention period for geo-redundancy backups.
	//
	// example:
	//
	// 7
	CrossRetentionValue *int32 `json:"CrossRetentionValue,omitempty" xml:"CrossRetentionValue,omitempty"`
	// The region where geo-redundancy backups are stored.
	//
	// example:
	//
	// cn-shenzhen
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// Specifies whether to enable log backup. Valid values:
	//
	// - **0**: disabled (default)
	//
	// - **1**: enabled
	//
	// example:
	//
	// 1
	EnableBackupLog *int32 `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// Specifies whether cross-region log backup is enabled.
	//
	// > Required for geo-redundancy backups.
	//
	// > - 1: enabled (must be set to 1 for sharded cluster instances)
	//
	// > - 0: disabled (must be set to 0 for replica set instances)
	//
	// example:
	//
	// 1
	EnableCrossLogBackup *int32 `json:"EnableCrossLogBackup,omitempty" xml:"EnableCrossLogBackup,omitempty"`
	// The retention period for high-frequency backups, in days.
	//
	// example:
	//
	// 1
	HighFrequencyBackupRetention *string `json:"HighFrequencyBackupRetention,omitempty" xml:"HighFrequencyBackupRetention,omitempty"`
	// The retention period for log backups, in days. Valid values: 7 to 730.
	//
	// example:
	//
	// 7
	LogBackupRetentionPeriod *int32 `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	// The backup cycle. Valid values:
	//
	// - **Monday**: Monday
	//
	// - **Tuesday**: Tuesday
	//
	// - **Wednesday**: Wednesday
	//
	// - **Thursday**: Thursday
	//
	// - **Friday**: Friday
	//
	// - **Saturday**: Saturday
	//
	// - **Sunday**: Sunday
	//
	// example:
	//
	// Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The backup time. The format is <i>HH:mm</i>Z-<i>HH:mm</i>Z (UTC time).
	//
	// example:
	//
	// 09:00Z-10:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The time of the next regular backup. The format is <i>yyyy-mm-dd</i>t<i>hh:mm</i>z (UTC time).
	//
	// example:
	//
	// 2024-06-19T19:11Z
	PreferredNextBackupTime *string `json:"PreferredNextBackupTime,omitempty" xml:"PreferredNextBackupTime,omitempty"`
	// Specifies whether to enable sparse backup within an hour. Valid values:
	//
	// - true: When the backup frequency is set to minutes, all snapshots taken within 1 hour from the current time are retained. For snapshots taken more than 1 hour but within 24 hours from the current time, only the first snapshot after the hour is retained.
	//
	// - false: Sparse backup within an hour is disabled. All snapshots within the high-frequency backup retention period are retained.
	//
	// example:
	//
	// false
	PreserveOneEachHour *bool `json:"PreserveOneEachHour,omitempty" xml:"PreserveOneEachHour,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5A9464CA-F7DC-5434-90B1-DF7F197C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The snapshot backup type. Valid values:
	//
	// - **Flash**: backup within seconds
	//
	// - **Standard**: standard backup (default)
	//
	// example:
	//
	// Standard
	SnapshotBackupType *string `json:"SnapshotBackupType,omitempty" xml:"SnapshotBackupType,omitempty"`
	// The region where the source instance is deployed.
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
