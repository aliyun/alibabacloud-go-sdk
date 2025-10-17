// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobExecutorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExexutors(v []*ListJobExecutorsResponseBodyExexutors) *ListJobExecutorsResponseBody
	GetExexutors() []*ListJobExecutorsResponseBodyExexutors
	SetMaxResults(v int32) *ListJobExecutorsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListJobExecutorsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListJobExecutorsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListJobExecutorsResponseBody
	GetTotalCount() *int32
}

type ListJobExecutorsResponseBody struct {
	Exexutors []*ListJobExecutorsResponseBodyExexutors `json:"exexutors,omitempty" xml:"exexutors,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 2
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListJobExecutorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobExecutorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsResponseBody) GetExexutors() []*ListJobExecutorsResponseBodyExexutors {
	return s.Exexutors
}

func (s *ListJobExecutorsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListJobExecutorsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListJobExecutorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobExecutorsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListJobExecutorsResponseBody) SetExexutors(v []*ListJobExecutorsResponseBodyExexutors) *ListJobExecutorsResponseBody {
	s.Exexutors = v
	return s
}

func (s *ListJobExecutorsResponseBody) SetMaxResults(v int32) *ListJobExecutorsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListJobExecutorsResponseBody) SetNextToken(v string) *ListJobExecutorsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListJobExecutorsResponseBody) SetRequestId(v string) *ListJobExecutorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobExecutorsResponseBody) SetTotalCount(v int32) *ListJobExecutorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobExecutorsResponseBody) Validate() error {
	if s.Exexutors != nil {
		for _, item := range s.Exexutors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobExecutorsResponseBodyExexutors struct {
	// example:
	//
	// 10
	ActiveTasks *int64 `json:"activeTasks,omitempty" xml:"activeTasks,omitempty"`
	// example:
	//
	// 1760601***
	AddTime *int64 `json:"addTime,omitempty" xml:"addTime,omitempty"`
	// example:
	//
	// 8
	CompletedTasks *int64 `json:"completedTasks,omitempty" xml:"completedTasks,omitempty"`
	// example:
	//
	// 20
	DiskUsed *int64 `json:"diskUsed,omitempty" xml:"diskUsed,omitempty"`
	// example:
	//
	// 1
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	// example:
	//
	// driver
	ExecutorType *string `json:"executorType,omitempty" xml:"executorType,omitempty"`
	// example:
	//
	// 2
	FailedTasks *int64 `json:"failedTasks,omitempty" xml:"failedTasks,omitempty"`
	// example:
	//
	// 21.10.x.x:1201x
	HostPort *string `json:"hostPort,omitempty" xml:"hostPort,omitempty"`
	// example:
	//
	// jr-1fe145df8ade366a
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// example:
	//
	// 4294967296
	MaxMemory *int64 `json:"maxMemory,omitempty" xml:"maxMemory,omitempty"`
	// example:
	//
	// 30
	MemoryUsed *int64 `json:"memoryUsed,omitempty" xml:"memoryUsed,omitempty"`
	// example:
	//
	// 10
	RddBlocks *int64 `json:"rddBlocks,omitempty" xml:"rddBlocks,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 4
	TotalCores *int64 `json:"totalCores,omitempty" xml:"totalCores,omitempty"`
	// example:
	//
	// 123
	TotalDuration *int64 `json:"totalDuration,omitempty" xml:"totalDuration,omitempty"`
	// example:
	//
	// 3
	TotalGCTime *int64 `json:"totalGCTime,omitempty" xml:"totalGCTime,omitempty"`
	// example:
	//
	// 1024
	TotalInputBytes *int64 `json:"totalInputBytes,omitempty" xml:"totalInputBytes,omitempty"`
	// example:
	//
	// 2048
	TotalShuffleRead *int64 `json:"totalShuffleRead,omitempty" xml:"totalShuffleRead,omitempty"`
	// example:
	//
	// 2048
	TotalShuffleWrite *int64 `json:"totalShuffleWrite,omitempty" xml:"totalShuffleWrite,omitempty"`
	// example:
	//
	// 50
	TotalTasks *int64 `json:"totalTasks,omitempty" xml:"totalTasks,omitempty"`
	// example:
	//
	// w-78faee4da118f02e
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListJobExecutorsResponseBodyExexutors) String() string {
	return dara.Prettify(s)
}

func (s ListJobExecutorsResponseBodyExexutors) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsResponseBodyExexutors) GetActiveTasks() *int64 {
	return s.ActiveTasks
}

func (s *ListJobExecutorsResponseBodyExexutors) GetAddTime() *int64 {
	return s.AddTime
}

func (s *ListJobExecutorsResponseBodyExexutors) GetCompletedTasks() *int64 {
	return s.CompletedTasks
}

func (s *ListJobExecutorsResponseBodyExexutors) GetDiskUsed() *int64 {
	return s.DiskUsed
}

func (s *ListJobExecutorsResponseBodyExexutors) GetExecutorId() *string {
	return s.ExecutorId
}

func (s *ListJobExecutorsResponseBodyExexutors) GetExecutorType() *string {
	return s.ExecutorType
}

func (s *ListJobExecutorsResponseBodyExexutors) GetFailedTasks() *int64 {
	return s.FailedTasks
}

func (s *ListJobExecutorsResponseBodyExexutors) GetHostPort() *string {
	return s.HostPort
}

func (s *ListJobExecutorsResponseBodyExexutors) GetJobRunId() *string {
	return s.JobRunId
}

func (s *ListJobExecutorsResponseBodyExexutors) GetMaxMemory() *int64 {
	return s.MaxMemory
}

func (s *ListJobExecutorsResponseBodyExexutors) GetMemoryUsed() *int64 {
	return s.MemoryUsed
}

func (s *ListJobExecutorsResponseBodyExexutors) GetRddBlocks() *int64 {
	return s.RddBlocks
}

func (s *ListJobExecutorsResponseBodyExexutors) GetStatus() *string {
	return s.Status
}

func (s *ListJobExecutorsResponseBodyExexutors) GetTotalCores() *int64 {
	return s.TotalCores
}

func (s *ListJobExecutorsResponseBodyExexutors) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *ListJobExecutorsResponseBodyExexutors) GetTotalGCTime() *int64 {
	return s.TotalGCTime
}

func (s *ListJobExecutorsResponseBodyExexutors) GetTotalInputBytes() *int64 {
	return s.TotalInputBytes
}

func (s *ListJobExecutorsResponseBodyExexutors) GetTotalShuffleRead() *int64 {
	return s.TotalShuffleRead
}

func (s *ListJobExecutorsResponseBodyExexutors) GetTotalShuffleWrite() *int64 {
	return s.TotalShuffleWrite
}

func (s *ListJobExecutorsResponseBodyExexutors) GetTotalTasks() *int64 {
	return s.TotalTasks
}

func (s *ListJobExecutorsResponseBodyExexutors) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListJobExecutorsResponseBodyExexutors) SetActiveTasks(v int64) *ListJobExecutorsResponseBodyExexutors {
	s.ActiveTasks = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetAddTime(v int64) *ListJobExecutorsResponseBodyExexutors {
	s.AddTime = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetCompletedTasks(v int64) *ListJobExecutorsResponseBodyExexutors {
	s.CompletedTasks = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetDiskUsed(v int64) *ListJobExecutorsResponseBodyExexutors {
	s.DiskUsed = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetExecutorId(v string) *ListJobExecutorsResponseBodyExexutors {
	s.ExecutorId = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetExecutorType(v string) *ListJobExecutorsResponseBodyExexutors {
	s.ExecutorType = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetFailedTasks(v int64) *ListJobExecutorsResponseBodyExexutors {
	s.FailedTasks = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetHostPort(v string) *ListJobExecutorsResponseBodyExexutors {
	s.HostPort = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetJobRunId(v string) *ListJobExecutorsResponseBodyExexutors {
	s.JobRunId = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetMaxMemory(v int64) *ListJobExecutorsResponseBodyExexutors {
	s.MaxMemory = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetMemoryUsed(v int64) *ListJobExecutorsResponseBodyExexutors {
	s.MemoryUsed = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetRddBlocks(v int64) *ListJobExecutorsResponseBodyExexutors {
	s.RddBlocks = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetStatus(v string) *ListJobExecutorsResponseBodyExexutors {
	s.Status = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetTotalCores(v int64) *ListJobExecutorsResponseBodyExexutors {
	s.TotalCores = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetTotalDuration(v int64) *ListJobExecutorsResponseBodyExexutors {
	s.TotalDuration = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetTotalGCTime(v int64) *ListJobExecutorsResponseBodyExexutors {
	s.TotalGCTime = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetTotalInputBytes(v int64) *ListJobExecutorsResponseBodyExexutors {
	s.TotalInputBytes = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetTotalShuffleRead(v int64) *ListJobExecutorsResponseBodyExexutors {
	s.TotalShuffleRead = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetTotalShuffleWrite(v int64) *ListJobExecutorsResponseBodyExexutors {
	s.TotalShuffleWrite = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetTotalTasks(v int64) *ListJobExecutorsResponseBodyExexutors {
	s.TotalTasks = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) SetWorkspaceId(v string) *ListJobExecutorsResponseBodyExexutors {
	s.WorkspaceId = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExexutors) Validate() error {
	return dara.Validate(s)
}
