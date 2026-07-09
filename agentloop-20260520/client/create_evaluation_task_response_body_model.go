// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEvaluationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEvaluationTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateEvaluationTaskResponseBody
	GetStatus() *string
	SetTaskId(v string) *CreateEvaluationTaskResponseBody
	GetTaskId() *string
}

type CreateEvaluationTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The evaluation task status. After creation, the status is typically `Pending`. After asynchronous orchestration, the status may change to `Running` or `Scheduling`.
	//
	// example:
	//
	// Pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The evaluation task ID.
	//
	// example:
	//
	// eval-task-8b36f2e2b1f94f9c91ce7a4b0f6d9c25
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateEvaluationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEvaluationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEvaluationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEvaluationTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateEvaluationTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateEvaluationTaskResponseBody) SetRequestId(v string) *CreateEvaluationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEvaluationTaskResponseBody) SetStatus(v string) *CreateEvaluationTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CreateEvaluationTaskResponseBody) SetTaskId(v string) *CreateEvaluationTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateEvaluationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
