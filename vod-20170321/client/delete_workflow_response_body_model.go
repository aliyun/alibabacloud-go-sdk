// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWorkflowResponseBody
	GetRequestId() *string
	SetWorkflowId(v string) *DeleteWorkflowResponseBody
	GetWorkflowId() *string
}

type DeleteWorkflowResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s DeleteWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkflowResponseBody) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *DeleteWorkflowResponseBody) SetRequestId(v string) *DeleteWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkflowResponseBody) SetWorkflowId(v string) *DeleteWorkflowResponseBody {
	s.WorkflowId = &v
	return s
}

func (s *DeleteWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}
