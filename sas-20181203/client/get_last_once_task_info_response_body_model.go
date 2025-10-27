// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLastOnceTaskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCollectTime(v int64) *GetLastOnceTaskInfoResponseBody
	GetCollectTime() *int64
	SetFinishCount(v int32) *GetLastOnceTaskInfoResponseBody
	GetFinishCount() *int32
	SetRequestId(v string) *GetLastOnceTaskInfoResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *GetLastOnceTaskInfoResponseBody
	GetTaskId() *int64
	SetTaskInfo(v *GetLastOnceTaskInfoResponseBodyTaskInfo) *GetLastOnceTaskInfoResponseBody
	GetTaskInfo() *GetLastOnceTaskInfoResponseBodyTaskInfo
	SetTotalCount(v int32) *GetLastOnceTaskInfoResponseBody
	GetTotalCount() *int32
}

type GetLastOnceTaskInfoResponseBody struct {
	// The time at which the task was run.
	//
	// example:
	//
	// 1671184531000
	CollectTime *int64 `json:"CollectTime,omitempty" xml:"CollectTime,omitempty"`
	// The number of tasks that have been completed.
	//
	// example:
	//
	// 67
	FinishCount *int32 `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578AB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the latest scan task.
	//
	// example:
	//
	// 3f65e1f1bb13118891a889d569a3****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The information about the latest task.
	TaskInfo *GetLastOnceTaskInfoResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 44
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetLastOnceTaskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLastOnceTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetLastOnceTaskInfoResponseBody) GetCollectTime() *int64 {
	return s.CollectTime
}

func (s *GetLastOnceTaskInfoResponseBody) GetFinishCount() *int32 {
	return s.FinishCount
}

func (s *GetLastOnceTaskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLastOnceTaskInfoResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetLastOnceTaskInfoResponseBody) GetTaskInfo() *GetLastOnceTaskInfoResponseBodyTaskInfo {
	return s.TaskInfo
}

func (s *GetLastOnceTaskInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetLastOnceTaskInfoResponseBody) SetCollectTime(v int64) *GetLastOnceTaskInfoResponseBody {
	s.CollectTime = &v
	return s
}

func (s *GetLastOnceTaskInfoResponseBody) SetFinishCount(v int32) *GetLastOnceTaskInfoResponseBody {
	s.FinishCount = &v
	return s
}

func (s *GetLastOnceTaskInfoResponseBody) SetRequestId(v string) *GetLastOnceTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLastOnceTaskInfoResponseBody) SetTaskId(v int64) *GetLastOnceTaskInfoResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetLastOnceTaskInfoResponseBody) SetTaskInfo(v *GetLastOnceTaskInfoResponseBodyTaskInfo) *GetLastOnceTaskInfoResponseBody {
	s.TaskInfo = v
	return s
}

func (s *GetLastOnceTaskInfoResponseBody) SetTotalCount(v int32) *GetLastOnceTaskInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetLastOnceTaskInfoResponseBody) Validate() error {
	if s.TaskInfo != nil {
		if err := s.TaskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLastOnceTaskInfoResponseBodyTaskInfo struct {
	// The progress of the task in percentage.
	//
	// example:
	//
	// 69
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The result of the scan task. Valid values:
	//
	// 	- **SUCCESS**: The task is successful.
	//
	// 	- **TASK_NOT_SUPPORT_REGION**: The images are deployed in a region that is not supported by container image scan.
	//
	// 	- **TASK_NOT_EXISTS**: The task does not exist.
	//
	// example:
	//
	// SUCCESS
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **INIT**: The task is not started.
	//
	// 	- **START**: The task is started.
	//
	// 	- **SUCCESS**: The task is complete.
	//
	// 	- **TIMEOUT**: The task timed out.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLastOnceTaskInfoResponseBodyTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s GetLastOnceTaskInfoResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *GetLastOnceTaskInfoResponseBodyTaskInfo) GetProgress() *int32 {
	return s.Progress
}

func (s *GetLastOnceTaskInfoResponseBodyTaskInfo) GetResult() *string {
	return s.Result
}

func (s *GetLastOnceTaskInfoResponseBodyTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *GetLastOnceTaskInfoResponseBodyTaskInfo) SetProgress(v int32) *GetLastOnceTaskInfoResponseBodyTaskInfo {
	s.Progress = &v
	return s
}

func (s *GetLastOnceTaskInfoResponseBodyTaskInfo) SetResult(v string) *GetLastOnceTaskInfoResponseBodyTaskInfo {
	s.Result = &v
	return s
}

func (s *GetLastOnceTaskInfoResponseBodyTaskInfo) SetStatus(v string) *GetLastOnceTaskInfoResponseBodyTaskInfo {
	s.Status = &v
	return s
}

func (s *GetLastOnceTaskInfoResponseBodyTaskInfo) Validate() error {
	return dara.Validate(s)
}
