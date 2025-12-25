// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateUnholdWorkflowExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateUnholdWorkflowExecutionResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateUnholdWorkflowExecutionResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateUnholdWorkflowExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateUnholdWorkflowExecutionResponseBody
	GetSuccess() *bool
}

type OperateUnholdWorkflowExecutionResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3835AA29-2298-5434-BC53-9CC377CDFD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateUnholdWorkflowExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateUnholdWorkflowExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *OperateUnholdWorkflowExecutionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateUnholdWorkflowExecutionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateUnholdWorkflowExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateUnholdWorkflowExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateUnholdWorkflowExecutionResponseBody) SetCode(v int32) *OperateUnholdWorkflowExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *OperateUnholdWorkflowExecutionResponseBody) SetMessage(v string) *OperateUnholdWorkflowExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *OperateUnholdWorkflowExecutionResponseBody) SetRequestId(v string) *OperateUnholdWorkflowExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateUnholdWorkflowExecutionResponseBody) SetSuccess(v bool) *OperateUnholdWorkflowExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *OperateUnholdWorkflowExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
