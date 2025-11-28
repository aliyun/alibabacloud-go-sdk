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
	// The backup set ID.
	//
	// >  You can call the [DescribeDataBackups](https://help.aliyun.com/document_detail/210093.html) operation to query the IDs of all backup sets in the instance.
	//
	// example:
	//
	// 1111111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup job ID.
	//
	// example:
	//
	// 123
	BackupJobId *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **Automated**
	//
	// 	- **Manual**
	//
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The backup status. Valid values:
	//
	// 	- **schedule**
	//
	// 	- **check**
	//
	// 	- **backup**
	//
	// 	- **finish**
	//
	// example:
	//
	// backup
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The progress of the backup job.
	//
	// example:
	//
	// 50%
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the backup job started. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
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
