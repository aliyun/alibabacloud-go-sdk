// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowDagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateWorkflowDagResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateWorkflowDagResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateWorkflowDagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWorkflowDagResponseBody
	GetSuccess() *bool
}

type UpdateWorkflowDagResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
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

func (s UpdateWorkflowDagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDagResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDagResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateWorkflowDagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWorkflowDagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkflowDagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWorkflowDagResponseBody) SetCode(v int32) *UpdateWorkflowDagResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWorkflowDagResponseBody) SetMessage(v string) *UpdateWorkflowDagResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWorkflowDagResponseBody) SetRequestId(v string) *UpdateWorkflowDagResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowDagResponseBody) SetSuccess(v bool) *UpdateWorkflowDagResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkflowDagResponseBody) Validate() error {
	return dara.Validate(s)
}
