// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DescribeBackupJobResponseBody
	GetBackupId() *string
	SetBackupJobId(v string) *DescribeBackupJobResponseBody
	GetBackupJobId() *string
	SetBackupMode(v string) *DescribeBackupJobResponseBody
	GetBackupMode() *string
	SetBackupStatus(v string) *DescribeBackupJobResponseBody
	GetBackupStatus() *string
	SetProcess(v string) *DescribeBackupJobResponseBody
	GetProcess() *string
	SetRequestId(v string) *DescribeBackupJobResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeBackupJobResponseBody
	GetStartTime() *string
}

type DescribeBackupJobResponseBody struct {
	// example:
	//
	// 1111111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// example:
	//
	// 123
	BackupJobId *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// example:
	//
	// backup
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// example:
	//
	// 50%
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2023-01-03T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBackupJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobResponseBody) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeBackupJobResponseBody) GetBackupJobId() *string {
	return s.BackupJobId
}

func (s *DescribeBackupJobResponseBody) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeBackupJobResponseBody) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeBackupJobResponseBody) GetProcess() *string {
	return s.Process
}

func (s *DescribeBackupJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupJobResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBackupJobResponseBody) SetBackupId(v string) *DescribeBackupJobResponseBody {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupJobResponseBody) SetBackupJobId(v string) *DescribeBackupJobResponseBody {
	s.BackupJobId = &v
	return s
}

func (s *DescribeBackupJobResponseBody) SetBackupMode(v string) *DescribeBackupJobResponseBody {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupJobResponseBody) SetBackupStatus(v string) *DescribeBackupJobResponseBody {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupJobResponseBody) SetProcess(v string) *DescribeBackupJobResponseBody {
	s.Process = &v
	return s
}

func (s *DescribeBackupJobResponseBody) SetRequestId(v string) *DescribeBackupJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupJobResponseBody) SetStartTime(v string) *DescribeBackupJobResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupJobResponseBody) Validate() error {
	return dara.Validate(s)
}
