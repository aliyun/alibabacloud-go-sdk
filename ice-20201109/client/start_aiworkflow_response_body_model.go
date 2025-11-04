// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartAIWorkflowResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StartAIWorkflowResponseBody
	GetTaskId() *string
}

type StartAIWorkflowResponseBody struct {
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

func (s StartAIWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAIWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *StartAIWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAIWorkflowResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StartAIWorkflowResponseBody) SetRequestId(v string) *StartAIWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAIWorkflowResponseBody) SetTaskId(v string) *StartAIWorkflowResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartAIWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}
