// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRestoreTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRestoreTasksResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListRestoreTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRestoreTasksResponseBody
	GetRequestId() *string
	SetRestoreTasks(v []*ListRestoreTasksResponseBodyRestoreTasks) *ListRestoreTasksResponseBody
	GetRestoreTasks() []*ListRestoreTasksResponseBodyRestoreTasks
	SetTotalCount(v int32) *ListRestoreTasksResponseBody
	GetTotalCount() *int32
}

type ListRestoreTasksResponseBody struct {
	// The maximum number of records returned in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Indicates the read position returned by the current call. An empty value means all data has been read.
	//
	// This parameter is required.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 464C8CB6-A548-5206-B83C-D32A8E43EC21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of restore tasks.
	RestoreTasks []*ListRestoreTasksResponseBodyRestoreTasks `json:"RestoreTasks,omitempty" xml:"RestoreTasks,omitempty" type:"Repeated"`
	// Total data count under the current request conditions (optional; not returned by default).
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRestoreTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRestoreTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListRestoreTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRestoreTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRestoreTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRestoreTasksResponseBody) GetRestoreTasks() []*ListRestoreTasksResponseBodyRestoreTasks {
	return s.RestoreTasks
}

func (s *ListRestoreTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRestoreTasksResponseBody) SetMaxResults(v int32) *ListRestoreTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRestoreTasksResponseBody) SetNextToken(v string) *ListRestoreTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRestoreTasksResponseBody) SetRequestId(v string) *ListRestoreTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRestoreTasksResponseBody) SetRestoreTasks(v []*ListRestoreTasksResponseBodyRestoreTasks) *ListRestoreTasksResponseBody {
	s.RestoreTasks = v
	return s
}

func (s *ListRestoreTasksResponseBody) SetTotalCount(v int32) *ListRestoreTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRestoreTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRestoreTasksResponseBodyRestoreTasks struct {
	// The backup ID.
	//
	// example:
	//
	// backup-728f128bf92c4e3da970
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2023-09-07T11:37:37Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expiration time of the service instance.
	//
	// example:
	//
	// 2025-01-27T18:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2025-05-07T12:16:16Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the restore task.
	//
	// example:
	//
	// restore-xxxxxx
	RestoreTaskId *string `json:"RestoreTaskId,omitempty" xml:"RestoreTaskId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-8c367c27c84e44a79d36
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The time when the update started.
	//
	// example:
	//
	// 2025-01-27T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the service instance. Valid values:
	//
	// 	- Restoring
	//
	// 	- Restored
	//
	// 	- RestoreFailed
	//
	// example:
	//
	// Restoring
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the service instance deployment information.
	//
	// example:
	//
	// i-xxxx  failed, error message: error
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
}

func (s ListRestoreTasksResponseBodyRestoreTasks) String() string {
	return dara.Prettify(s)
}

func (s ListRestoreTasksResponseBodyRestoreTasks) GoString() string {
	return s.String()
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) GetBackupId() *string {
	return s.BackupId
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) GetEndTime() *string {
	return s.EndTime
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) GetRestoreTaskId() *string {
	return s.RestoreTaskId
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) GetStartTime() *string {
	return s.StartTime
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) GetStatus() *string {
	return s.Status
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) GetStatusDetail() *string {
	return s.StatusDetail
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) SetBackupId(v string) *ListRestoreTasksResponseBodyRestoreTasks {
	s.BackupId = &v
	return s
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) SetCreateTime(v string) *ListRestoreTasksResponseBodyRestoreTasks {
	s.CreateTime = &v
	return s
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) SetEndTime(v string) *ListRestoreTasksResponseBodyRestoreTasks {
	s.EndTime = &v
	return s
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) SetModifiedTime(v string) *ListRestoreTasksResponseBodyRestoreTasks {
	s.ModifiedTime = &v
	return s
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) SetRestoreTaskId(v string) *ListRestoreTasksResponseBodyRestoreTasks {
	s.RestoreTaskId = &v
	return s
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) SetServiceInstanceId(v string) *ListRestoreTasksResponseBodyRestoreTasks {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) SetStartTime(v string) *ListRestoreTasksResponseBodyRestoreTasks {
	s.StartTime = &v
	return s
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) SetStatus(v string) *ListRestoreTasksResponseBodyRestoreTasks {
	s.Status = &v
	return s
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) SetStatusDetail(v string) *ListRestoreTasksResponseBodyRestoreTasks {
	s.StatusDetail = &v
	return s
}

func (s *ListRestoreTasksResponseBodyRestoreTasks) Validate() error {
	return dara.Validate(s)
}
