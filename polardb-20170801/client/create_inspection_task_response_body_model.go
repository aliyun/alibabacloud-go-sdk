// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInspectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateInspectionTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateInspectionTaskResponseBody
	GetTaskId() *string
}

type CreateInspectionTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the created inspection task.
	//
	// example:
	//
	// task-123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateInspectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInspectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInspectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInspectionTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateInspectionTaskResponseBody) SetRequestId(v string) *CreateInspectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInspectionTaskResponseBody) SetTaskId(v string) *CreateInspectionTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateInspectionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
