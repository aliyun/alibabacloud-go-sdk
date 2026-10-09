// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupabaseBackupJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ListSupabaseBackupJobsResponseBodyItems) *ListSupabaseBackupJobsResponseBody
	GetItems() []*ListSupabaseBackupJobsResponseBodyItems
	SetMaxResults(v int32) *ListSupabaseBackupJobsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSupabaseBackupJobsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSupabaseBackupJobsResponseBody
	GetRequestId() *string
}

type ListSupabaseBackupJobsResponseBody struct {
	// The list of backup tasks.
	Items []*ListSupabaseBackupJobsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum number of entries to return for this request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page, which can be used as the NextToken parameter in the next request.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSupabaseBackupJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseBackupJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSupabaseBackupJobsResponseBody) GetItems() []*ListSupabaseBackupJobsResponseBodyItems {
	return s.Items
}

func (s *ListSupabaseBackupJobsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSupabaseBackupJobsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupabaseBackupJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSupabaseBackupJobsResponseBody) SetItems(v []*ListSupabaseBackupJobsResponseBodyItems) *ListSupabaseBackupJobsResponseBody {
	s.Items = v
	return s
}

func (s *ListSupabaseBackupJobsResponseBody) SetMaxResults(v int32) *ListSupabaseBackupJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSupabaseBackupJobsResponseBody) SetNextToken(v string) *ListSupabaseBackupJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSupabaseBackupJobsResponseBody) SetRequestId(v string) *ListSupabaseBackupJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSupabaseBackupJobsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSupabaseBackupJobsResponseBodyItems struct {
	// The ID of the backup task.
	//
	// example:
	//
	// 123
	BackupJobId *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **Automated**: automatic backup
	//
	// 	- **Manual**: manual backup
	//
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The status of the backup task. Valid statuses include: schedule (waiting to be scheduled) and backup (in progress).
	//
	// example:
	//
	// backup
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The progress percentage of the backup task, such as 0%. This value may be an empty string when the task is in the schedule (waiting to be scheduled) state.
	//
	// example:
	//
	// 0%
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// The start time of the backup task. The time is displayed in UTC in the yyyy-MM-ddTHH:mm:ssZ format. This value may be an empty string when the task is in the schedule (waiting to be scheduled) state.
	//
	// example:
	//
	// 2026-10-09T04:37:01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListSupabaseBackupJobsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseBackupJobsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListSupabaseBackupJobsResponseBodyItems) GetBackupJobId() *string {
	return s.BackupJobId
}

func (s *ListSupabaseBackupJobsResponseBodyItems) GetBackupMode() *string {
	return s.BackupMode
}

func (s *ListSupabaseBackupJobsResponseBodyItems) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *ListSupabaseBackupJobsResponseBodyItems) GetProcess() *string {
	return s.Process
}

func (s *ListSupabaseBackupJobsResponseBodyItems) GetStartTime() *string {
	return s.StartTime
}

func (s *ListSupabaseBackupJobsResponseBodyItems) SetBackupJobId(v string) *ListSupabaseBackupJobsResponseBodyItems {
	s.BackupJobId = &v
	return s
}

func (s *ListSupabaseBackupJobsResponseBodyItems) SetBackupMode(v string) *ListSupabaseBackupJobsResponseBodyItems {
	s.BackupMode = &v
	return s
}

func (s *ListSupabaseBackupJobsResponseBodyItems) SetBackupStatus(v string) *ListSupabaseBackupJobsResponseBodyItems {
	s.BackupStatus = &v
	return s
}

func (s *ListSupabaseBackupJobsResponseBodyItems) SetProcess(v string) *ListSupabaseBackupJobsResponseBodyItems {
	s.Process = &v
	return s
}

func (s *ListSupabaseBackupJobsResponseBodyItems) SetStartTime(v string) *ListSupabaseBackupJobsResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *ListSupabaseBackupJobsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
