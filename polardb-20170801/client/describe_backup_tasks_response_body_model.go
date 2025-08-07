// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeBackupTasksResponseBodyItems) *DescribeBackupTasksResponseBody
	GetItems() *DescribeBackupTasksResponseBodyItems
	SetRequestId(v string) *DescribeBackupTasksResponseBody
	GetRequestId() *string
}

type DescribeBackupTasksResponseBody struct {
	// The details of the backup task.
	Items *DescribeBackupTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FA8C1EF1-E3D4-44D7-B809-823187******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponseBody) GetItems() *DescribeBackupTasksResponseBodyItems {
	return s.Items
}

func (s *DescribeBackupTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupTasksResponseBody) SetItems(v *DescribeBackupTasksResponseBodyItems) *DescribeBackupTasksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBackupTasksResponseBody) SetRequestId(v string) *DescribeBackupTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupTasksResponseBodyItems struct {
	BackupJob []*DescribeBackupTasksResponseBodyItemsBackupJob `json:"BackupJob,omitempty" xml:"BackupJob,omitempty" type:"Repeated"`
}

func (s DescribeBackupTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponseBodyItems) GetBackupJob() []*DescribeBackupTasksResponseBodyItemsBackupJob {
	return s.BackupJob
}

func (s *DescribeBackupTasksResponseBodyItems) SetBackupJob(v []*DescribeBackupTasksResponseBodyItemsBackupJob) *DescribeBackupTasksResponseBodyItems {
	s.BackupJob = v
	return s
}

func (s *DescribeBackupTasksResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupTasksResponseBodyItemsBackupJob struct {
	// The ID of the backup task.
	//
	// example:
	//
	// 11111111
	BackupJobId *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// The state of the backup task. Valid values:
	//
	// 	- **NoStart**
	//
	// 	- **Preparing**
	//
	// 	- **Waiting**
	//
	// 	- **Uploading**
	//
	// 	- **Checking**
	//
	// 	- **Finished**
	//
	// example:
	//
	// NoStart
	BackupProgressStatus *string `json:"BackupProgressStatus,omitempty" xml:"BackupProgressStatus,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **Automated**
	//
	// 	- **Manual**
	//
	// example:
	//
	// Automated
	JobMode *string `json:"JobMode,omitempty" xml:"JobMode,omitempty"`
	// The progress of the backup task in percentage.
	//
	// example:
	//
	// 0
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// The time when the backup task started. The time follows the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-08-08T07:24:01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the backup task. Valid values:
	//
	// 	- **TempBackupTask**: The backup task is an adhoc backup task.
	//
	// 	- **NormalBackupTask**: The backup task is a common backup task.
	//
	// example:
	//
	// NormalBackupTask
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
}

func (s DescribeBackupTasksResponseBodyItemsBackupJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTasksResponseBodyItemsBackupJob) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) GetBackupJobId() *string {
	return s.BackupJobId
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) GetBackupProgressStatus() *string {
	return s.BackupProgressStatus
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) GetJobMode() *string {
	return s.JobMode
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) GetProcess() *string {
	return s.Process
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) GetTaskAction() *string {
	return s.TaskAction
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetBackupJobId(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.BackupJobId = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetBackupProgressStatus(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.BackupProgressStatus = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetJobMode(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.JobMode = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetProcess(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.Process = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetStartTime(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetTaskAction(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.TaskAction = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) Validate() error {
	return dara.Validate(s)
}
