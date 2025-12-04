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
	SetBackupSize(v string) *DescribeBackupPolicyResponseBody
	GetBackupSize() *string
	SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupTime() *string
	SetRequestId(v string) *DescribeBackupPolicyResponseBody
	GetRequestId() *string
	SetSwitch(v string) *DescribeBackupPolicyResponseBody
	GetSwitch() *string
}

type DescribeBackupPolicyResponseBody struct {
	// The retention period for the backup data. By default, the backup data is retained for seven days. Valid values: 7 to 730. Unit: day.
	//
	// example:
	//
	// 7
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The size of the backup data. Unit: MB.
	//
	// example:
	//
	// 123124
	BackupSize *string `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
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
	// example:
	//
	// Monday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The backup window. The time is displayed in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// 15:00Z-16:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the backup feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Switch *string `json:"Switch,omitempty" xml:"Switch,omitempty"`
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

func (s *DescribeBackupPolicyResponseBody) GetBackupSize() *string {
	return s.BackupSize
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *DescribeBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPolicyResponseBody) GetSwitch() *string {
	return s.Switch
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupSize(v string) *DescribeBackupPolicyResponseBody {
	s.BackupSize = &v
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

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetSwitch(v string) *DescribeBackupPolicyResponseBody {
	s.Switch = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
