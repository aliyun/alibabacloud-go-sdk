// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackupJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *ListBackupJobsResponseBodyItems) *ListBackupJobsResponseBody
	GetItems() *ListBackupJobsResponseBodyItems
	SetRequestId(v string) *ListBackupJobsResponseBody
	GetRequestId() *string
}

type ListBackupJobsResponseBody struct {
	Items *ListBackupJobsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBackupJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBackupJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBackupJobsResponseBody) GetItems() *ListBackupJobsResponseBodyItems {
	return s.Items
}

func (s *ListBackupJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBackupJobsResponseBody) SetItems(v *ListBackupJobsResponseBodyItems) *ListBackupJobsResponseBody {
	s.Items = v
	return s
}

func (s *ListBackupJobsResponseBody) SetRequestId(v string) *ListBackupJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBackupJobsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBackupJobsResponseBodyItems struct {
	BackupJob []*ListBackupJobsResponseBodyItemsBackupJob `json:"BackupJob,omitempty" xml:"BackupJob,omitempty" type:"Repeated"`
}

func (s ListBackupJobsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListBackupJobsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListBackupJobsResponseBodyItems) GetBackupJob() []*ListBackupJobsResponseBodyItemsBackupJob {
	return s.BackupJob
}

func (s *ListBackupJobsResponseBodyItems) SetBackupJob(v []*ListBackupJobsResponseBodyItemsBackupJob) *ListBackupJobsResponseBodyItems {
	s.BackupJob = v
	return s
}

func (s *ListBackupJobsResponseBodyItems) Validate() error {
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

type ListBackupJobsResponseBodyItemsBackupJob struct {
	BackupJobId  *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	BackupMode   *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	Process      *string `json:"Process,omitempty" xml:"Process,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListBackupJobsResponseBodyItemsBackupJob) String() string {
	return dara.Prettify(s)
}

func (s ListBackupJobsResponseBodyItemsBackupJob) GoString() string {
	return s.String()
}

func (s *ListBackupJobsResponseBodyItemsBackupJob) GetBackupJobId() *string {
	return s.BackupJobId
}

func (s *ListBackupJobsResponseBodyItemsBackupJob) GetBackupMode() *string {
	return s.BackupMode
}

func (s *ListBackupJobsResponseBodyItemsBackupJob) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *ListBackupJobsResponseBodyItemsBackupJob) GetProcess() *string {
	return s.Process
}

func (s *ListBackupJobsResponseBodyItemsBackupJob) GetStartTime() *string {
	return s.StartTime
}

func (s *ListBackupJobsResponseBodyItemsBackupJob) SetBackupJobId(v string) *ListBackupJobsResponseBodyItemsBackupJob {
	s.BackupJobId = &v
	return s
}

func (s *ListBackupJobsResponseBodyItemsBackupJob) SetBackupMode(v string) *ListBackupJobsResponseBodyItemsBackupJob {
	s.BackupMode = &v
	return s
}

func (s *ListBackupJobsResponseBodyItemsBackupJob) SetBackupStatus(v string) *ListBackupJobsResponseBodyItemsBackupJob {
	s.BackupStatus = &v
	return s
}

func (s *ListBackupJobsResponseBodyItemsBackupJob) SetProcess(v string) *ListBackupJobsResponseBodyItemsBackupJob {
	s.Process = &v
	return s
}

func (s *ListBackupJobsResponseBodyItemsBackupJob) SetStartTime(v string) *ListBackupJobsResponseBodyItemsBackupJob {
	s.StartTime = &v
	return s
}

func (s *ListBackupJobsResponseBodyItemsBackupJob) Validate() error {
	return dara.Validate(s)
}
