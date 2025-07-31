// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupJobs(v []*DescribeBackupTasksResponseBodyBackupJobs) *DescribeBackupTasksResponseBody
	GetBackupJobs() []*DescribeBackupTasksResponseBodyBackupJobs
	SetRequestId(v string) *DescribeBackupTasksResponseBody
	GetRequestId() *string
}

type DescribeBackupTasksResponseBody struct {
	// The details of the backup task.
	BackupJobs []*DescribeBackupTasksResponseBodyBackupJobs `json:"BackupJobs,omitempty" xml:"BackupJobs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D648B367-15B6-1B42-BD4B-47507C9CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponseBody) GetBackupJobs() []*DescribeBackupTasksResponseBodyBackupJobs {
	return s.BackupJobs
}

func (s *DescribeBackupTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupTasksResponseBody) SetBackupJobs(v []*DescribeBackupTasksResponseBodyBackupJobs) *DescribeBackupTasksResponseBody {
	s.BackupJobs = v
	return s
}

func (s *DescribeBackupTasksResponseBody) SetRequestId(v string) *DescribeBackupTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupTasksResponseBodyBackupJobs struct {
	// The backup task status. Valid values:
	//
	// 	- **Scheduled**: The backup task is in planning. Regular backup tasks that have not started are also in this state.
	//
	// 	- **Checking**: The instance is being checked before the backup.
	//
	// 	- **Backuping**: The backup task is in progress.
	//
	// 	- **Finished**: The backup task is completed.
	//
	// example:
	//
	// Scheduled
	BackupSetStatus *string `json:"BackupSetStatus,omitempty" xml:"BackupSetStatus,omitempty"`
	// The start time of the backup task.
	//
	// example:
	//
	// 2024-01-16T11:04:56Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The ID of the backup task.
	//
	// example:
	//
	// 170034
	BackupjobId *string `json:"BackupjobId,omitempty" xml:"BackupjobId,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **Automated**: automatic backup
	//
	// 	- **Manual**: manual backup
	//
	// example:
	//
	// Manual
	JobMode *string `json:"JobMode,omitempty" xml:"JobMode,omitempty"`
	// The progress of the backup task. Unit: %. The progress is returned only for running backup tasks.
	//
	// example:
	//
	// 18%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeBackupTasksResponseBodyBackupJobs) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTasksResponseBodyBackupJobs) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) GetBackupSetStatus() *string {
	return s.BackupSetStatus
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) GetBackupjobId() *string {
	return s.BackupjobId
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) GetJobMode() *string {
	return s.JobMode
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) GetProgress() *string {
	return s.Progress
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetBackupSetStatus(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.BackupSetStatus = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetBackupStartTime(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetBackupjobId(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.BackupjobId = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetJobMode(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.JobMode = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetProgress(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.Progress = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) Validate() error {
	return dara.Validate(s)
}
