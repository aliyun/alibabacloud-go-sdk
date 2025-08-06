// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWorkflowResponseBody
	GetRequestId() *string
	SetWorkflowId(v string) *UpdateWorkflowResponseBody
	GetWorkflowId() *string
}

type UpdateWorkflowResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s UpdateWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkflowResponseBody) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *UpdateWorkflowResponseBody) SetRequestId(v string) *UpdateWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowResponseBody) SetWorkflowId(v string) *UpdateWorkflowResponseBody {
	s.WorkflowId = &v
	return s
}

func (s *UpdateWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}
