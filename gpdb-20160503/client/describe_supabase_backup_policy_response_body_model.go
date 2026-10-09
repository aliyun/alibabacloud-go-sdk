// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupabaseBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupInterval(v int32) *DescribeSupabaseBackupPolicyResponseBody
	GetBackupInterval() *int32
	SetBackupRetentionPeriod(v int32) *DescribeSupabaseBackupPolicyResponseBody
	GetBackupRetentionPeriod() *int32
	SetEnableRecoveryPoint(v bool) *DescribeSupabaseBackupPolicyResponseBody
	GetEnableRecoveryPoint() *bool
	SetPreferredBackupPeriod(v string) *DescribeSupabaseBackupPolicyResponseBody
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *DescribeSupabaseBackupPolicyResponseBody
	GetPreferredBackupTime() *string
	SetRecoveryPointPeriod(v string) *DescribeSupabaseBackupPolicyResponseBody
	GetRecoveryPointPeriod() *string
	SetRequestId(v string) *DescribeSupabaseBackupPolicyResponseBody
	GetRequestId() *string
}

type DescribeSupabaseBackupPolicyResponseBody struct {
	// The interval between automatic recovery points, in minutes. A value greater than 0 indicates that automatic recovery points are enabled. If the feature is disabled, -1 is returned.
	//
	// example:
	//
	// -1
	BackupInterval *int32 `json:"BackupInterval,omitempty" xml:"BackupInterval,omitempty"`
	// The data backup retention period, in days.
	//
	// example:
	//
	// 7
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// Indicates whether automatic recovery points are enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	EnableRecoveryPoint *bool `json:"EnableRecoveryPoint,omitempty" xml:"EnableRecoveryPoint,omitempty"`
	// The data backup cycle. Separate multiple values with commas (,). Valid values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, and Sunday.
	//
	// example:
	//
	// Wednesday,Friday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The data backup time window in UTC. The format is HH:mmZ-HH:mmZ.
	//
	// example:
	//
	// 01:00Z-02:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The interval for the automatic creation of recovery points, in hours. Valid values: 1/6 (10 minutes), 1/2 (30 minutes), 1, 2, 4, and 8. This value is valid only when EnableRecoveryPoint is set to true. If automatic recovery points are shutdown, 0 is returned.
	//
	// example:
	//
	// 0
	RecoveryPointPeriod *string `json:"RecoveryPointPeriod,omitempty" xml:"RecoveryPointPeriod,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSupabaseBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseBackupPolicyResponseBody) GetBackupInterval() *int32 {
	return s.BackupInterval
}

func (s *DescribeSupabaseBackupPolicyResponseBody) GetBackupRetentionPeriod() *int32 {
	return s.BackupRetentionPeriod
}

func (s *DescribeSupabaseBackupPolicyResponseBody) GetEnableRecoveryPoint() *bool {
	return s.EnableRecoveryPoint
}

func (s *DescribeSupabaseBackupPolicyResponseBody) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *DescribeSupabaseBackupPolicyResponseBody) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *DescribeSupabaseBackupPolicyResponseBody) GetRecoveryPointPeriod() *string {
	return s.RecoveryPointPeriod
}

func (s *DescribeSupabaseBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSupabaseBackupPolicyResponseBody) SetBackupInterval(v int32) *DescribeSupabaseBackupPolicyResponseBody {
	s.BackupInterval = &v
	return s
}

func (s *DescribeSupabaseBackupPolicyResponseBody) SetBackupRetentionPeriod(v int32) *DescribeSupabaseBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeSupabaseBackupPolicyResponseBody) SetEnableRecoveryPoint(v bool) *DescribeSupabaseBackupPolicyResponseBody {
	s.EnableRecoveryPoint = &v
	return s
}

func (s *DescribeSupabaseBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeSupabaseBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeSupabaseBackupPolicyResponseBody) SetPreferredBackupTime(v string) *DescribeSupabaseBackupPolicyResponseBody {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeSupabaseBackupPolicyResponseBody) SetRecoveryPointPeriod(v string) *DescribeSupabaseBackupPolicyResponseBody {
	s.RecoveryPointPeriod = &v
	return s
}

func (s *DescribeSupabaseBackupPolicyResponseBody) SetRequestId(v string) *DescribeSupabaseBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupabaseBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
