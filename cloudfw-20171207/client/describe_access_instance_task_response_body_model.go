// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessInstanceTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsFound(v bool) *DescribeAccessInstanceTaskResponseBody
	GetIsFound() *bool
	SetRequestId(v string) *DescribeAccessInstanceTaskResponseBody
	GetRequestId() *string
	SetTaskFinishTimestamp(v int64) *DescribeAccessInstanceTaskResponseBody
	GetTaskFinishTimestamp() *int64
	SetTaskId(v string) *DescribeAccessInstanceTaskResponseBody
	GetTaskId() *string
	SetTaskName(v string) *DescribeAccessInstanceTaskResponseBody
	GetTaskName() *string
	SetTaskStartTimestamp(v int64) *DescribeAccessInstanceTaskResponseBody
	GetTaskStartTimestamp() *int64
	SetTaskStatus(v string) *DescribeAccessInstanceTaskResponseBody
	GetTaskStatus() *string
	SetTaskSteps(v []*DescribeAccessInstanceTaskResponseBodyTaskSteps) *DescribeAccessInstanceTaskResponseBody
	GetTaskSteps() []*DescribeAccessInstanceTaskResponseBodyTaskSteps
}

type DescribeAccessInstanceTaskResponseBody struct {
	// example:
	//
	// true
	IsFound *bool `json:"IsFound,omitempty" xml:"IsFound,omitempty"`
	// example:
	//
	// 15FCCC52-1E23-57AE-B5EF-3E00A3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 17151381075
	TaskFinishTimestamp *int64 `json:"TaskFinishTimestamp,omitempty" xml:"TaskFinishTimestamp,omitempty"`
	// example:
	//
	// 3c9d576f-fce0-4caa-9116-15033509bdb6
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 17151361285
	TaskStartTimestamp *int64 `json:"TaskStartTimestamp,omitempty" xml:"TaskStartTimestamp,omitempty"`
	// example:
	//
	// running
	TaskStatus *string                                            `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskSteps  []*DescribeAccessInstanceTaskResponseBodyTaskSteps `json:"TaskSteps,omitempty" xml:"TaskSteps,omitempty" type:"Repeated"`
}

func (s DescribeAccessInstanceTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceTaskResponseBody) GetIsFound() *bool {
	return s.IsFound
}

func (s *DescribeAccessInstanceTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessInstanceTaskResponseBody) GetTaskFinishTimestamp() *int64 {
	return s.TaskFinishTimestamp
}

func (s *DescribeAccessInstanceTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeAccessInstanceTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeAccessInstanceTaskResponseBody) GetTaskStartTimestamp() *int64 {
	return s.TaskStartTimestamp
}

func (s *DescribeAccessInstanceTaskResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeAccessInstanceTaskResponseBody) GetTaskSteps() []*DescribeAccessInstanceTaskResponseBodyTaskSteps {
	return s.TaskSteps
}

func (s *DescribeAccessInstanceTaskResponseBody) SetIsFound(v bool) *DescribeAccessInstanceTaskResponseBody {
	s.IsFound = &v
	return s
}

func (s *DescribeAccessInstanceTaskResponseBody) SetRequestId(v string) *DescribeAccessInstanceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessInstanceTaskResponseBody) SetTaskFinishTimestamp(v int64) *DescribeAccessInstanceTaskResponseBody {
	s.TaskFinishTimestamp = &v
	return s
}

func (s *DescribeAccessInstanceTaskResponseBody) SetTaskId(v string) *DescribeAccessInstanceTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeAccessInstanceTaskResponseBody) SetTaskName(v string) *DescribeAccessInstanceTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *DescribeAccessInstanceTaskResponseBody) SetTaskStartTimestamp(v int64) *DescribeAccessInstanceTaskResponseBody {
	s.TaskStartTimestamp = &v
	return s
}

func (s *DescribeAccessInstanceTaskResponseBody) SetTaskStatus(v string) *DescribeAccessInstanceTaskResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeAccessInstanceTaskResponseBody) SetTaskSteps(v []*DescribeAccessInstanceTaskResponseBodyTaskSteps) *DescribeAccessInstanceTaskResponseBody {
	s.TaskSteps = v
	return s
}

func (s *DescribeAccessInstanceTaskResponseBody) Validate() error {
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

type DescribeAccessInstanceTaskResponseBodyTaskSteps struct {
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// example:
	//
	// 100%
	StepProgress *string `json:"StepProgress,omitempty" xml:"StepProgress,omitempty"`
	// example:
	//
	// finished
	StepStatus *string `json:"StepStatus,omitempty" xml:"StepStatus,omitempty"`
}

func (s DescribeAccessInstanceTaskResponseBodyTaskSteps) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceTaskResponseBodyTaskSteps) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceTaskResponseBodyTaskSteps) GetStepName() *string {
	return s.StepName
}

func (s *DescribeAccessInstanceTaskResponseBodyTaskSteps) GetStepProgress() *string {
	return s.StepProgress
}

func (s *DescribeAccessInstanceTaskResponseBodyTaskSteps) GetStepStatus() *string {
	return s.StepStatus
}

func (s *DescribeAccessInstanceTaskResponseBodyTaskSteps) SetStepName(v string) *DescribeAccessInstanceTaskResponseBodyTaskSteps {
	s.StepName = &v
	return s
}

func (s *DescribeAccessInstanceTaskResponseBodyTaskSteps) SetStepProgress(v string) *DescribeAccessInstanceTaskResponseBodyTaskSteps {
	s.StepProgress = &v
	return s
}

func (s *DescribeAccessInstanceTaskResponseBodyTaskSteps) SetStepStatus(v string) *DescribeAccessInstanceTaskResponseBodyTaskSteps {
	s.StepStatus = &v
	return s
}

func (s *DescribeAccessInstanceTaskResponseBodyTaskSteps) Validate() error {
	return dara.Validate(s)
}
