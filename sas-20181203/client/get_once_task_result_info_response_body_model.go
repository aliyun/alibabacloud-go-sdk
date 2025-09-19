// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnceTaskResultInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCollectTime(v int64) *GetOnceTaskResultInfoResponseBody
	GetCollectTime() *int64
	SetFinishCount(v int32) *GetOnceTaskResultInfoResponseBody
	GetFinishCount() *int32
	SetRequestId(v string) *GetOnceTaskResultInfoResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *GetOnceTaskResultInfoResponseBody
	GetTaskId() *int64
	SetTaskInfo(v *GetOnceTaskResultInfoResponseBodyTaskInfo) *GetOnceTaskResultInfoResponseBody
	GetTaskInfo() *GetOnceTaskResultInfoResponseBodyTaskInfo
	SetTotalCount(v int32) *GetOnceTaskResultInfoResponseBody
	GetTotalCount() *int32
}

type GetOnceTaskResultInfoResponseBody struct {
	// The execution time of the task.
	//
	// example:
	//
	// 1671184531000
	CollectTime *int64 `json:"CollectTime,omitempty" xml:"CollectTime,omitempty"`
	// The number of tasks that were completed.
	//
	// example:
	//
	// 47
	FinishCount *int32 `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CE500770-42D3-442E-9DDD-156E0F9F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scan task.
	//
	// example:
	//
	// e7b70a4b030db086db52231f1b58****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The information about the task.
	TaskInfo *GetOnceTaskResultInfoResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 44
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetOnceTaskResultInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOnceTaskResultInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetOnceTaskResultInfoResponseBody) GetCollectTime() *int64 {
	return s.CollectTime
}

func (s *GetOnceTaskResultInfoResponseBody) GetFinishCount() *int32 {
	return s.FinishCount
}

func (s *GetOnceTaskResultInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOnceTaskResultInfoResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetOnceTaskResultInfoResponseBody) GetTaskInfo() *GetOnceTaskResultInfoResponseBodyTaskInfo {
	return s.TaskInfo
}

func (s *GetOnceTaskResultInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetOnceTaskResultInfoResponseBody) SetCollectTime(v int64) *GetOnceTaskResultInfoResponseBody {
	s.CollectTime = &v
	return s
}

func (s *GetOnceTaskResultInfoResponseBody) SetFinishCount(v int32) *GetOnceTaskResultInfoResponseBody {
	s.FinishCount = &v
	return s
}

func (s *GetOnceTaskResultInfoResponseBody) SetRequestId(v string) *GetOnceTaskResultInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOnceTaskResultInfoResponseBody) SetTaskId(v int64) *GetOnceTaskResultInfoResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetOnceTaskResultInfoResponseBody) SetTaskInfo(v *GetOnceTaskResultInfoResponseBodyTaskInfo) *GetOnceTaskResultInfoResponseBody {
	s.TaskInfo = v
	return s
}

func (s *GetOnceTaskResultInfoResponseBody) SetTotalCount(v int32) *GetOnceTaskResultInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetOnceTaskResultInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOnceTaskResultInfoResponseBodyTaskInfo struct {
	// The status of the task. Valid values:
	//
	// 	- **INIT**: The task is not started.
	//
	// 	- **START**: The task is started.
	//
	// 	- **SUCCESS**: The task is complete.
	//
	// 	- **TIMEOUT**: The task times out.
	//
	// example:
	//
	// START
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetOnceTaskResultInfoResponseBodyTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s GetOnceTaskResultInfoResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *GetOnceTaskResultInfoResponseBodyTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *GetOnceTaskResultInfoResponseBodyTaskInfo) SetStatus(v string) *GetOnceTaskResultInfoResponseBodyTaskInfo {
	s.Status = &v
	return s
}

func (s *GetOnceTaskResultInfoResponseBodyTaskInfo) Validate() error {
	return dara.Validate(s)
}
