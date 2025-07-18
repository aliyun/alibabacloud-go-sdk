// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody
	GetBackupRetentionPeriod() *int32
	SetEnableRecoveryPoint(v bool) *DescribeBackupPolicyResponseBody
	GetEnableRecoveryPoint() *bool
	SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupTime() *string
	SetRecoveryPointPeriod(v string) *DescribeBackupPolicyResponseBody
	GetRecoveryPointPeriod() *string
	SetRequestId(v string) *DescribeBackupPolicyResponseBody
	GetRequestId() *string
}

type DescribeBackupPolicyResponseBody struct {
	// The number of days for which data backup files are retained.
	//
	// example:
	//
	// 7
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// Indicates whether automatic point-in-time backup is enabled. Valid values:
	//
	// 	- **true**: Automatic point-in-time backup is enabled.
	//
	// 	- **false**: Automatic point-in-time backup is disabled.
	//
	// example:
	//
	// true
	EnableRecoveryPoint *bool `json:"EnableRecoveryPoint,omitempty" xml:"EnableRecoveryPoint,omitempty"`
	// The cycle based on which backups are performed. If more than one day of the week is specified, the days of the week are separated by commas (,). Valid values:
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
	// Wednesday,Friday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The backup time. The time is in the HH:mmZ-HH:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 12:00Z-13:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The frequency of the point-in-time backup. Valid values:
	//
	// 	- **1**: per hour
	//
	// 	- **2**: per 2 hours
	//
	// 	- **4**: per 4 hours
	//
	// 	- **8**: per 8 hours
	//
	// example:
	//
	// 1
	RecoveryPointPeriod *string `json:"RecoveryPointPeriod,omitempty" xml:"RecoveryPointPeriod,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9E6B3287-A3E2-5A87-B8D8-E9**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) GetBackupRetentionPeriod() *int32 {
	return s.BackupRetentionPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetEnableRecoveryPoint() *bool {
	return s.EnableRecoveryPoint
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *DescribeBackupPolicyResponseBody) GetRecoveryPointPeriod() *string {
	return s.RecoveryPointPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetEnableRecoveryPoint(v bool) *DescribeBackupPolicyResponseBody {
	s.EnableRecoveryPoint = &v
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

func (s *DescribeBackupPolicyResponseBody) SetRecoveryPointPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.RecoveryPointPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
