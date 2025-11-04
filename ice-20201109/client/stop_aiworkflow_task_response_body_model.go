// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAIWorkflowTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopAIWorkflowTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StopAIWorkflowTaskResponseBody
	GetTaskId() *string
}

type StopAIWorkflowTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the workflow task.
	//
	// example:
	//
	// ********-266c-4bb8-b20c-6faa********
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopAIWorkflowTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopAIWorkflowTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopAIWorkflowTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopAIWorkflowTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StopAIWorkflowTaskResponseBody) SetRequestId(v string) *StopAIWorkflowTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopAIWorkflowTaskResponseBody) SetTaskId(v string) *StopAIWorkflowTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *StopAIWorkflowTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
