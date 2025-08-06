// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitWorkflowJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitWorkflowJobResponseBody
	GetRequestId() *string
	SetTaskId(v string) *SubmitWorkflowJobResponseBody
	GetTaskId() *string
}

type SubmitWorkflowJobResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A01C8FF4-C106-4431-418F973DADB7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitWorkflowJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitWorkflowJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitWorkflowJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitWorkflowJobResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitWorkflowJobResponseBody) SetRequestId(v string) *SubmitWorkflowJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitWorkflowJobResponseBody) SetTaskId(v string) *SubmitWorkflowJobResponseBody {
	s.TaskId = &v
	return s
}

func (s *SubmitWorkflowJobResponseBody) Validate() error {
	return dara.Validate(s)
}
