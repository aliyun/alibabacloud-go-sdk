// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateMarkSuccessWorkflowExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateMarkSuccessWorkflowExecutionResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateMarkSuccessWorkflowExecutionResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateMarkSuccessWorkflowExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateMarkSuccessWorkflowExecutionResponseBody
	GetSuccess() *bool
}

type OperateMarkSuccessWorkflowExecutionResponseBody struct {
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
	// 1DF6732E-15D8-5E1F-95E3-C10077F556B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateMarkSuccessWorkflowExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateMarkSuccessWorkflowExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *OperateMarkSuccessWorkflowExecutionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateMarkSuccessWorkflowExecutionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateMarkSuccessWorkflowExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateMarkSuccessWorkflowExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateMarkSuccessWorkflowExecutionResponseBody) SetCode(v int32) *OperateMarkSuccessWorkflowExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *OperateMarkSuccessWorkflowExecutionResponseBody) SetMessage(v string) *OperateMarkSuccessWorkflowExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *OperateMarkSuccessWorkflowExecutionResponseBody) SetRequestId(v string) *OperateMarkSuccessWorkflowExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateMarkSuccessWorkflowExecutionResponseBody) SetSuccess(v bool) *OperateMarkSuccessWorkflowExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *OperateMarkSuccessWorkflowExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
