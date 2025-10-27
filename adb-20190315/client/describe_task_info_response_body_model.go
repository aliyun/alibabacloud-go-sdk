// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTaskInfoResponseBody
	GetRequestId() *string
	SetTaskInfo(v *DescribeTaskInfoResponseBodyTaskInfo) *DescribeTaskInfoResponseBody
	GetTaskInfo() *DescribeTaskInfoResponseBodyTaskInfo
}

type DescribeTaskInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried task.
	TaskInfo *DescribeTaskInfoResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s DescribeTaskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTaskInfoResponseBody) GetTaskInfo() *DescribeTaskInfoResponseBodyTaskInfo {
	return s.TaskInfo
}

func (s *DescribeTaskInfoResponseBody) SetRequestId(v string) *DescribeTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskInfoResponseBody) SetTaskInfo(v *DescribeTaskInfoResponseBodyTaskInfo) *DescribeTaskInfoResponseBody {
	s.TaskInfo = v
	return s
}

func (s *DescribeTaskInfoResponseBody) Validate() error {
	if s.TaskInfo != nil {
		if err := s.TaskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTaskInfoResponseBodyTaskInfo struct {
	// The start time of the task. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2020-01-07T07:39:56Z
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end time of the task. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2020-01-07T08:08:50Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The progress of the task. Unit: %.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status. Valid values:
	//
	// 	- Waiting
	//
	// 	- Running
	//
	// 	- Finished
	//
	// 	- Failed
	//
	// 	- Closed
	//
	// 	- Cancel
	//
	// 	- Retry
	//
	// 	- Pause
	//
	// 	- Stop
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 225685759
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTaskInfoResponseBodyTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskInfoResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) GetBeginTime() *string {
	return s.BeginTime
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) GetProgress() *string {
	return s.Progress
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetBeginTime(v string) *DescribeTaskInfoResponseBodyTaskInfo {
	s.BeginTime = &v
	return s
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetFinishTime(v string) *DescribeTaskInfoResponseBodyTaskInfo {
	s.FinishTime = &v
	return s
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetProgress(v string) *DescribeTaskInfoResponseBodyTaskInfo {
	s.Progress = &v
	return s
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetStatus(v string) *DescribeTaskInfoResponseBodyTaskInfo {
	s.Status = &v
	return s
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetTaskId(v int32) *DescribeTaskInfoResponseBodyTaskInfo {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) Validate() error {
	return dara.Validate(s)
}
