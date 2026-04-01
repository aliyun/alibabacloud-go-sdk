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
	Items *DescribeBackupTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 90496720-2319-42A8-87CD-FCE4DF95EBED
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
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.BackupJob != nil {
		for _, item := range s.BackupJob {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupTasksResponseBodyItemsBackupJob struct {
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupJobId          *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	BackupProgressStatus *string `json:"BackupProgressStatus,omitempty" xml:"BackupProgressStatus,omitempty"`
	BackupStatus         *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	JobMode              *string `json:"JobMode,omitempty" xml:"JobMode,omitempty"`
	Process              *string `json:"Process,omitempty" xml:"Process,omitempty"`
	TaskAction           *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
}

func (s DescribeBackupTasksResponseBodyItemsBackupJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTasksResponseBodyItemsBackupJob) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) GetBackupJobId() *string {
	return s.BackupJobId
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) GetBackupProgressStatus() *string {
	return s.BackupProgressStatus
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) GetJobMode() *string {
	return s.JobMode
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) GetProcess() *string {
	return s.Process
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) GetTaskAction() *string {
	return s.TaskAction
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetBackupId(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetBackupJobId(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.BackupJobId = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetBackupProgressStatus(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.BackupProgressStatus = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetBackupStatus(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.BackupStatus = &v
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

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetTaskAction(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.TaskAction = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) Validate() error {
	return dara.Validate(s)
}
