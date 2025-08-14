// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartWorkflowResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StartWorkflowResponseBody
	GetTaskId() *string
}

type StartWorkflowResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******42-E8E1-4FBB-8E52-F4225C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the workflow task.
	//
	// example:
	//
	// ******22dad741d086a50325f9******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *StartWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartWorkflowResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StartWorkflowResponseBody) SetRequestId(v string) *StartWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartWorkflowResponseBody) SetTaskId(v string) *StartWorkflowResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}
