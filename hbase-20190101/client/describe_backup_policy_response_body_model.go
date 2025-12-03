// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody
	GetBackupRetentionPeriod() *string
	SetPreferredBackupEndTimeUTC(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupEndTimeUTC() *string
	SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupPeriod() *string
	SetPreferredBackupStartTimeUTC(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupStartTimeUTC() *string
	SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupTime() *string
	SetRequestId(v string) *DescribeBackupPolicyResponseBody
	GetRequestId() *string
}

type DescribeBackupPolicyResponseBody struct {
	// example:
	//
	// 10
	BackupRetentionPeriod *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// example:
	//
	// 18:00Z
	PreferredBackupEndTimeUTC *string `json:"PreferredBackupEndTimeUTC,omitempty" xml:"PreferredBackupEndTimeUTC,omitempty"`
	// example:
	//
	// Friday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// example:
	//
	// 17:00Z
	PreferredBackupStartTimeUTC *string `json:"PreferredBackupStartTimeUTC,omitempty" xml:"PreferredBackupStartTimeUTC,omitempty"`
	// example:
	//
	// 01:00-02:00
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// example:
	//
	// 94AC38B6-7C6D-45B2-BC03-B8750071A482
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) GetBackupRetentionPeriod() *string {
	return s.BackupRetentionPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredBackupEndTimeUTC() *string {
	return s.PreferredBackupEndTimeUTC
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredBackupStartTimeUTC() *string {
	return s.PreferredBackupStartTimeUTC
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *DescribeBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupEndTimeUTC(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupEndTimeUTC = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupStartTimeUTC(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupStartTimeUTC = &v
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

func (s *DescribeBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
