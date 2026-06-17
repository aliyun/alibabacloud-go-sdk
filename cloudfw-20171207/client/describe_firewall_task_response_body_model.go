// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsFound(v bool) *DescribeFirewallTaskResponseBody
	GetIsFound() *bool
	SetRequestId(v string) *DescribeFirewallTaskResponseBody
	GetRequestId() *string
	SetTaskFinishTimestamp(v string) *DescribeFirewallTaskResponseBody
	GetTaskFinishTimestamp() *string
	SetTaskId(v int64) *DescribeFirewallTaskResponseBody
	GetTaskId() *int64
	SetTaskName(v string) *DescribeFirewallTaskResponseBody
	GetTaskName() *string
	SetTaskStartTimestamp(v string) *DescribeFirewallTaskResponseBody
	GetTaskStartTimestamp() *string
	SetTaskStatus(v string) *DescribeFirewallTaskResponseBody
	GetTaskStatus() *string
	SetTaskSteps(v []*DescribeFirewallTaskResponseBodyTaskSteps) *DescribeFirewallTaskResponseBody
	GetTaskSteps() []*DescribeFirewallTaskResponseBodyTaskSteps
	SetTaskWaitingTime(v string) *DescribeFirewallTaskResponseBody
	GetTaskWaitingTime() *string
}

type DescribeFirewallTaskResponseBody struct {
	// Indicates whether the task exists.
	//
	// example:
	//
	// false
	IsFound *bool `json:"IsFound,omitempty" xml:"IsFound,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7D5483BF-2262-586D-8706-BDDB8B42****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timestamp when the task was completed.
	//
	// example:
	//
	// 17151381075
	TaskFinishTimestamp *string `json:"TaskFinishTimestamp,omitempty" xml:"TaskFinishTimestamp,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 189997648
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// egressgw
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The timestamp when the task started.
	//
	// example:
	//
	// 17151361285
	TaskStartTimestamp *string `json:"TaskStartTimestamp,omitempty" xml:"TaskStartTimestamp,omitempty"`
	// The status of the task. Valid values:
	//
	// - **init**
	//
	// - **running**
	//
	// - **finished**
	//
	// - **rollback**
	//
	// - **rollbackDone**
	//
	// example:
	//
	// init
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The steps of the task.
	TaskSteps []*DescribeFirewallTaskResponseBodyTaskSteps `json:"TaskSteps,omitempty" xml:"TaskSteps,omitempty" type:"Repeated"`
	// The waiting time in minutes.
	//
	// example:
	//
	// 30
	TaskWaitingTime *string `json:"TaskWaitingTime,omitempty" xml:"TaskWaitingTime,omitempty"`
}

func (s DescribeFirewallTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTaskResponseBody) GetIsFound() *bool {
	return s.IsFound
}

func (s *DescribeFirewallTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFirewallTaskResponseBody) GetTaskFinishTimestamp() *string {
	return s.TaskFinishTimestamp
}

func (s *DescribeFirewallTaskResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeFirewallTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeFirewallTaskResponseBody) GetTaskStartTimestamp() *string {
	return s.TaskStartTimestamp
}

func (s *DescribeFirewallTaskResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeFirewallTaskResponseBody) GetTaskSteps() []*DescribeFirewallTaskResponseBodyTaskSteps {
	return s.TaskSteps
}

func (s *DescribeFirewallTaskResponseBody) GetTaskWaitingTime() *string {
	return s.TaskWaitingTime
}

func (s *DescribeFirewallTaskResponseBody) SetIsFound(v bool) *DescribeFirewallTaskResponseBody {
	s.IsFound = &v
	return s
}

func (s *DescribeFirewallTaskResponseBody) SetRequestId(v string) *DescribeFirewallTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirewallTaskResponseBody) SetTaskFinishTimestamp(v string) *DescribeFirewallTaskResponseBody {
	s.TaskFinishTimestamp = &v
	return s
}

func (s *DescribeFirewallTaskResponseBody) SetTaskId(v int64) *DescribeFirewallTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeFirewallTaskResponseBody) SetTaskName(v string) *DescribeFirewallTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *DescribeFirewallTaskResponseBody) SetTaskStartTimestamp(v string) *DescribeFirewallTaskResponseBody {
	s.TaskStartTimestamp = &v
	return s
}

func (s *DescribeFirewallTaskResponseBody) SetTaskStatus(v string) *DescribeFirewallTaskResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeFirewallTaskResponseBody) SetTaskSteps(v []*DescribeFirewallTaskResponseBodyTaskSteps) *DescribeFirewallTaskResponseBody {
	s.TaskSteps = v
	return s
}

func (s *DescribeFirewallTaskResponseBody) SetTaskWaitingTime(v string) *DescribeFirewallTaskResponseBody {
	s.TaskWaitingTime = &v
	return s
}

func (s *DescribeFirewallTaskResponseBody) Validate() error {
	if s.TaskSteps != nil {
		for _, item := range s.TaskSteps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFirewallTaskResponseBodyTaskSteps struct {
	// The information about the task step.
	//
	// example:
	//
	// abcd
	StepInfo *string `json:"StepInfo,omitempty" xml:"StepInfo,omitempty"`
	// Creating the Cloud Firewall.
	//
	// example:
	//
	// Create Firewall
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The progress of the task step.
	//
	// example:
	//
	// 80
	StepProgress *string `json:"StepProgress,omitempty" xml:"StepProgress,omitempty"`
	// The status of the task step. Valid values:
	//
	// - **init**
	//
	// - **running**
	//
	// - **finished**
	//
	// - **failed**
	//
	// example:
	//
	// init
	StepStatus *string `json:"StepStatus,omitempty" xml:"StepStatus,omitempty"`
}

func (s DescribeFirewallTaskResponseBodyTaskSteps) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTaskResponseBodyTaskSteps) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTaskResponseBodyTaskSteps) GetStepInfo() *string {
	return s.StepInfo
}

func (s *DescribeFirewallTaskResponseBodyTaskSteps) GetStepName() *string {
	return s.StepName
}

func (s *DescribeFirewallTaskResponseBodyTaskSteps) GetStepProgress() *string {
	return s.StepProgress
}

func (s *DescribeFirewallTaskResponseBodyTaskSteps) GetStepStatus() *string {
	return s.StepStatus
}

func (s *DescribeFirewallTaskResponseBodyTaskSteps) SetStepInfo(v string) *DescribeFirewallTaskResponseBodyTaskSteps {
	s.StepInfo = &v
	return s
}

func (s *DescribeFirewallTaskResponseBodyTaskSteps) SetStepName(v string) *DescribeFirewallTaskResponseBodyTaskSteps {
	s.StepName = &v
	return s
}

func (s *DescribeFirewallTaskResponseBodyTaskSteps) SetStepProgress(v string) *DescribeFirewallTaskResponseBodyTaskSteps {
	s.StepProgress = &v
	return s
}

func (s *DescribeFirewallTaskResponseBodyTaskSteps) SetStepStatus(v string) *DescribeFirewallTaskResponseBodyTaskSteps {
	s.StepStatus = &v
	return s
}

func (s *DescribeFirewallTaskResponseBodyTaskSteps) Validate() error {
	return dara.Validate(s)
}
