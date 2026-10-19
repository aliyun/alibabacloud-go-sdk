// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvancedBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedBackupDataPolicies(v []*DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) *DescribeAdvancedBackupPolicyResponseBody
	GetAdvancedBackupDataPolicies() []*DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies
	SetAdvancedBackupLogPolicies(v []*DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) *DescribeAdvancedBackupPolicyResponseBody
	GetAdvancedBackupLogPolicies() []*DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies
	SetBackupRetentionPolicyOnClusterDeletion(v int32) *DescribeAdvancedBackupPolicyResponseBody
	GetBackupRetentionPolicyOnClusterDeletion() *int32
	SetDBInstanceId(v string) *DescribeAdvancedBackupPolicyResponseBody
	GetDBInstanceId() *string
	SetPreferredBackupTime(v string) *DescribeAdvancedBackupPolicyResponseBody
	GetPreferredBackupTime() *string
	SetPreferredNextBackupTime(v string) *DescribeAdvancedBackupPolicyResponseBody
	GetPreferredNextBackupTime() *string
	SetRequestId(v string) *DescribeAdvancedBackupPolicyResponseBody
	GetRequestId() *string
	SetSnapshotBackupType(v string) *DescribeAdvancedBackupPolicyResponseBody
	GetSnapshotBackupType() *string
}

type DescribeAdvancedBackupPolicyResponseBody struct {
	// The list of backup policies.
	AdvancedBackupDataPolicies []*DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies `json:"AdvancedBackupDataPolicies,omitempty" xml:"AdvancedBackupDataPolicies,omitempty" type:"Repeated"`
	// The list of log backup policies.
	AdvancedBackupLogPolicies []*DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies `json:"AdvancedBackupLogPolicies,omitempty" xml:"AdvancedBackupLogPolicies,omitempty" type:"Repeated"`
	// The data retention policy. Valid values:
	//
	// - 0: All backup sets of the instance are immediately deleted when the instance is released.
	//
	// - 1: The last backup set is retained for long-term retention (LTR) when the instance is released.
	//
	// - 2: All backup sets are retained for long-term retention (LTR) when the instance is released.
	//
	// example:
	//
	// 2
	BackupRetentionPolicyOnClusterDeletion *int32 `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dds-bp14bf67a76d****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The backup time in the <i>HH:mm</i>Z-<i>HH:mm</i>Z format (UTC).
	//
	// example:
	//
	// 09:00Z-10:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The time range for starting a backup. The time is in the `HH:mmZ-HH:mmZ` format (UTC). Example: 01:00Z-02:00Z. The backup starts within this time range.
	//
	// example:
	//
	// 2024-06-19T19:11Z
	PreferredNextBackupTime *string `json:"PreferredNextBackupTime,omitempty" xml:"PreferredNextBackupTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 111E7B16-0A87-4CBA-B271-F34AD61E099F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The snapshot backup type. Valid values:
	//
	// - **Flash**: backup within seconds
	//
	// - **Standard**: standard backup
	//
	// example:
	//
	// Standard
	SnapshotBackupType *string `json:"SnapshotBackupType,omitempty" xml:"SnapshotBackupType,omitempty"`
}

func (s DescribeAdvancedBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedBackupPolicyResponseBody) GetAdvancedBackupDataPolicies() []*DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies {
	return s.AdvancedBackupDataPolicies
}

func (s *DescribeAdvancedBackupPolicyResponseBody) GetAdvancedBackupLogPolicies() []*DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies {
	return s.AdvancedBackupLogPolicies
}

func (s *DescribeAdvancedBackupPolicyResponseBody) GetBackupRetentionPolicyOnClusterDeletion() *int32 {
	return s.BackupRetentionPolicyOnClusterDeletion
}

func (s *DescribeAdvancedBackupPolicyResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAdvancedBackupPolicyResponseBody) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *DescribeAdvancedBackupPolicyResponseBody) GetPreferredNextBackupTime() *string {
	return s.PreferredNextBackupTime
}

func (s *DescribeAdvancedBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdvancedBackupPolicyResponseBody) GetSnapshotBackupType() *string {
	return s.SnapshotBackupType
}

func (s *DescribeAdvancedBackupPolicyResponseBody) SetAdvancedBackupDataPolicies(v []*DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) *DescribeAdvancedBackupPolicyResponseBody {
	s.AdvancedBackupDataPolicies = v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBody) SetAdvancedBackupLogPolicies(v []*DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) *DescribeAdvancedBackupPolicyResponseBody {
	s.AdvancedBackupLogPolicies = v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBody) SetBackupRetentionPolicyOnClusterDeletion(v int32) *DescribeAdvancedBackupPolicyResponseBody {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBody) SetDBInstanceId(v string) *DescribeAdvancedBackupPolicyResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBody) SetPreferredBackupTime(v string) *DescribeAdvancedBackupPolicyResponseBody {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBody) SetPreferredNextBackupTime(v string) *DescribeAdvancedBackupPolicyResponseBody {
	s.PreferredNextBackupTime = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBody) SetRequestId(v string) *DescribeAdvancedBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBody) SetSnapshotBackupType(v string) *DescribeAdvancedBackupPolicyResponseBody {
	s.SnapshotBackupType = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBody) Validate() error {
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

type DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies struct {
	// The action type.
	//
	// example:
	//
	// xx
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The scheduling type of the advanced policy. Valid values:
	//
	// - dayOfWeek: weekly scheduling
	//
	// - dayOfMonth: monthly scheduling
	//
	// - dayOfYear: yearly scheduling
	//
	// - backupInterval: fixed-interval scheduling (high-frequency backup)
	//
	// example:
	//
	// dayOfWeek
	BackupPolicy *string `json:"BackupPolicy,omitempty" xml:"BackupPolicy,omitempty"`
	// The destination region of the backup policy.
	//
	// example:
	//
	// cn-beijing
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// The destination type of the backup policy. Valid values:
	//
	// - level1: same-region backup
	//
	// - level2Cross: cross-region backup
	//
	// example:
	//
	// level1
	DestType *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	// The ID of the backup policy.
	//
	// example:
	//
	// xxxxxx
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The backup cycle.
	//
	// - When **BackupPolicy*	- is set to **dayOfWeek**, the values 1, 2, 3, 4, 5, 6, and 7 represent Monday to Sunday.
	//
	// - When **BackupPolicy*	- is set to **dayOfMonth**, the values 1, 3, and -1 represent the 1st, 3rd, and last day of each month.
	//
	// - When **BackupPolicy*	- is set to **dayOfYear**, the value 01-07 represents January 7.
	//
	// - When **BackupPolicy*	- is set to **backupInterval**, the value 30 represents a backup every 30 minutes.
	//
	// example:
	//
	// 1,2,3
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// Specifies whether to enable sparse backup within an hour. Valid values:
	//
	// - true: When the backup frequency is minute-level, all snapshots taken within 1 hour from the current time are retained. For snapshots taken more than 1 hour but within 24 hours from the current time, only the first snapshot after each hour is retained.
	//
	// - false: Sparse backup within an hour is disabled. All snapshots within the retention period of high-frequency backups are retained.
	//
	// example:
	//
	// False
	PreserveOneEachHour *bool `json:"PreserveOneEachHour,omitempty" xml:"PreserveOneEachHour,omitempty"`
	// The number of days for which backups are retained.
	//
	// example:
	//
	// 7
	RetentionPeriod *int32 `json:"RetentionPeriod,omitempty" xml:"RetentionPeriod,omitempty"`
	// The retention epoch type of the backup set. Valid values:
	//
	// - **never**: long-term retention (LTR)
	//
	// - **delay**: fixed-day retention
	//
	// example:
	//
	// delay
	RetentionType *string `json:"RetentionType,omitempty" xml:"RetentionType,omitempty"`
	// The region where the instance resides.
	//
	// example:
	//
	// cn-beijing
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The source type of the backup policy. Valid values:
	//
	// - db: same-region backup
	//
	// - level1: cross-region backup
	//
	// example:
	//
	// db
	SrcType *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) GetActionType() *string {
	return s.ActionType
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) GetBackupPolicy() *string {
	return s.BackupPolicy
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) GetDestType() *string {
	return s.DestType
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) GetPreserveOneEachHour() *bool {
	return s.PreserveOneEachHour
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) GetRetentionPeriod() *int32 {
	return s.RetentionPeriod
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) GetRetentionType() *string {
	return s.RetentionType
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) SetActionType(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies {
	s.ActionType = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) SetBackupPolicy(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies {
	s.BackupPolicy = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) SetDestRegion(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies {
	s.DestRegion = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) SetDestType(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies {
	s.DestType = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) SetPolicyId(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies {
	s.PolicyId = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) SetPreferredBackupPeriod(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) SetPreserveOneEachHour(v bool) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies {
	s.PreserveOneEachHour = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) SetRetentionPeriod(v int32) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies {
	s.RetentionPeriod = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) SetRetentionType(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies {
	s.RetentionType = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) SetSrcRegion(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies {
	s.SrcRegion = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) SetSrcType(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies {
	s.SrcType = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupDataPolicies) Validate() error {
	return dara.Validate(s)
}

type DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies struct {
	// The destination region of the log backup policy.
	//
	// example:
	//
	// cn-guangzhou
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// The destination type of the log backup policy. Valid values:
	//
	// - level1: same-region backup
	//
	// - level2Cross: cross-region backup
	//
	// example:
	//
	// RDS
	DestType *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	// Specifies whether log backup is enabled. Valid values:
	//
	// - **1**: enabled
	//
	// - **0**: disabled
	//
	// example:
	//
	// 1
	EnableLogBackup *int32 `json:"EnableLogBackup,omitempty" xml:"EnableLogBackup,omitempty"`
	// The ID of the log backup policy. This parameter is returned only for cross-region log backup policies.
	//
	// example:
	//
	// xxxx
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The number of days for which log backups are retained. This parameter takes effect only when RetentionType is set to delay.
	//
	// example:
	//
	// 4
	RetentionPeriod *string `json:"RetentionPeriod,omitempty" xml:"RetentionPeriod,omitempty"`
	// The retention epoch type of log backups. Valid values:
	//
	// - **never**: long-term retention (LTR)
	//
	// - **delay**: fixed-day retention
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
	// - db: same-region backup
	//
	// - level1: cross-region backup
	//
	// example:
	//
	// et
	SrcType *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) GetDestType() *string {
	return s.DestType
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) GetEnableLogBackup() *int32 {
	return s.EnableLogBackup
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) GetRetentionPeriod() *string {
	return s.RetentionPeriod
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) GetRetentionType() *string {
	return s.RetentionType
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) SetDestRegion(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies {
	s.DestRegion = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) SetDestType(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies {
	s.DestType = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) SetEnableLogBackup(v int32) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies {
	s.EnableLogBackup = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) SetPolicyId(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies {
	s.PolicyId = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) SetRetentionPeriod(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies {
	s.RetentionPeriod = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) SetRetentionType(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies {
	s.RetentionType = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) SetSrcRegion(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies {
	s.SrcRegion = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) SetSrcType(v string) *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies {
	s.SrcType = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponseBodyAdvancedBackupLogPolicies) Validate() error {
	return dara.Validate(s)
}
