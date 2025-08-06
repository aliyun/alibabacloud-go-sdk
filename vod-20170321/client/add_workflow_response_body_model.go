// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddWorkflowResponseBody
	GetRequestId() *string
	SetWorkflowId(v string) *AddWorkflowResponseBody
	GetWorkflowId() *string
}

type AddWorkflowResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s AddWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *AddWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddWorkflowResponseBody) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *AddWorkflowResponseBody) SetRequestId(v string) *AddWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWorkflowResponseBody) SetWorkflowId(v string) *AddWorkflowResponseBody {
	s.WorkflowId = &v
	return s
}

func (s *AddWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}
