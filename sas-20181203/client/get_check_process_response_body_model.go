// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFinishCount(v int32) *GetCheckProcessResponseBody
	GetFinishCount() *int32
	SetRequestId(v string) *GetCheckProcessResponseBody
	GetRequestId() *string
	SetStatusCode(v string) *GetCheckProcessResponseBody
	GetStatusCode() *string
	SetTaskId(v string) *GetCheckProcessResponseBody
	GetTaskId() *string
	SetTotalCount(v int32) *GetCheckProcessResponseBody
	GetTotalCount() *int32
}

type GetCheckProcessResponseBody struct {
	// The total number of assets on which the task is complete.
	//
	// example:
	//
	// 80
	FinishCount *int32 `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D18B5DAD-BA97-5552-AE48-83F59D5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code of the Cloud Security Posture Management (CSPM) task. Valid values:
	//
	// 	- 0: The task is being initialized. The system is calculating the total number of subtasks.
	//
	// 	- 1: The task is being executed. You can query the total number of tasks and the number of completed tasks.
	//
	// 	- 2: The task is successful.
	//
	// 	- 3: The task times out.
	//
	// 	- 4: The task is invalid. Check whether assets exist.
	//
	// 	- 5: No task record is found. Check whether the TaskId parameter is valid.
	//
	// example:
	//
	// 1
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 5347c7b6-c85c-4070-846a-3029e08e****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The total number of assets on which the task is performed.
	//
	// example:
	//
	// 113
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetCheckProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCheckProcessResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckProcessResponseBody) GetFinishCount() *int32 {
	return s.FinishCount
}

func (s *GetCheckProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCheckProcessResponseBody) GetStatusCode() *string {
	return s.StatusCode
}

func (s *GetCheckProcessResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCheckProcessResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetCheckProcessResponseBody) SetFinishCount(v int32) *GetCheckProcessResponseBody {
	s.FinishCount = &v
	return s
}

func (s *GetCheckProcessResponseBody) SetRequestId(v string) *GetCheckProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCheckProcessResponseBody) SetStatusCode(v string) *GetCheckProcessResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetCheckProcessResponseBody) SetTaskId(v string) *GetCheckProcessResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetCheckProcessResponseBody) SetTotalCount(v int32) *GetCheckProcessResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetCheckProcessResponseBody) Validate() error {
	return dara.Validate(s)
}
