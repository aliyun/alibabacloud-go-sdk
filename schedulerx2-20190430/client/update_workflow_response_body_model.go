// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateWorkflowResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateWorkflowResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateWorkflowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWorkflowResponseBody
	GetSuccess() *bool
}

type UpdateWorkflowResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// timetype is invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateWorkflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWorkflowResponseBody) SetCode(v int32) *UpdateWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWorkflowResponseBody) SetMessage(v string) *UpdateWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWorkflowResponseBody) SetRequestId(v string) *UpdateWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowResponseBody) SetSuccess(v bool) *UpdateWorkflowResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}
