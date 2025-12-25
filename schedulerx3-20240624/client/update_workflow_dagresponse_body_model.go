// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowDAGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateWorkflowDAGResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateWorkflowDAGResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateWorkflowDAGResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWorkflowDAGResponseBody
	GetSuccess() *bool
}

type UpdateWorkflowDAGResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AA3538A0-FBE6-5E31-AD88-A02C6FF0DACC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkflowDAGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDAGResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDAGResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateWorkflowDAGResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWorkflowDAGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkflowDAGResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWorkflowDAGResponseBody) SetCode(v int32) *UpdateWorkflowDAGResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWorkflowDAGResponseBody) SetMessage(v string) *UpdateWorkflowDAGResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWorkflowDAGResponseBody) SetRequestId(v string) *UpdateWorkflowDAGResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowDAGResponseBody) SetSuccess(v bool) *UpdateWorkflowDAGResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkflowDAGResponseBody) Validate() error {
	return dara.Validate(s)
}
