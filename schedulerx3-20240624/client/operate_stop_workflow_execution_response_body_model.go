// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateStopWorkflowExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateStopWorkflowExecutionResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateStopWorkflowExecutionResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateStopWorkflowExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateStopWorkflowExecutionResponseBody
	GetSuccess() *bool
}

type OperateStopWorkflowExecutionResponseBody struct {
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
	// D0DE9C33-992A-580B-89C4-B609A292748D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateStopWorkflowExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateStopWorkflowExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *OperateStopWorkflowExecutionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateStopWorkflowExecutionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateStopWorkflowExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateStopWorkflowExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateStopWorkflowExecutionResponseBody) SetCode(v int32) *OperateStopWorkflowExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *OperateStopWorkflowExecutionResponseBody) SetMessage(v string) *OperateStopWorkflowExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *OperateStopWorkflowExecutionResponseBody) SetRequestId(v string) *OperateStopWorkflowExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateStopWorkflowExecutionResponseBody) SetSuccess(v bool) *OperateStopWorkflowExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *OperateStopWorkflowExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
