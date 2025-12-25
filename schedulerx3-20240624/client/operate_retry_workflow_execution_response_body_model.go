// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateRetryWorkflowExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateRetryWorkflowExecutionResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateRetryWorkflowExecutionResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateRetryWorkflowExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateRetryWorkflowExecutionResponseBody
	GetSuccess() *bool
}

type OperateRetryWorkflowExecutionResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter check error
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

func (s OperateRetryWorkflowExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateRetryWorkflowExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *OperateRetryWorkflowExecutionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateRetryWorkflowExecutionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateRetryWorkflowExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateRetryWorkflowExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateRetryWorkflowExecutionResponseBody) SetCode(v int32) *OperateRetryWorkflowExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *OperateRetryWorkflowExecutionResponseBody) SetMessage(v string) *OperateRetryWorkflowExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *OperateRetryWorkflowExecutionResponseBody) SetRequestId(v string) *OperateRetryWorkflowExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateRetryWorkflowExecutionResponseBody) SetSuccess(v bool) *OperateRetryWorkflowExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *OperateRetryWorkflowExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
