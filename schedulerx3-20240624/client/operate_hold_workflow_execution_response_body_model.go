// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateHoldWorkflowExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateHoldWorkflowExecutionResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateHoldWorkflowExecutionResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateHoldWorkflowExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateHoldWorkflowExecutionResponseBody
	GetSuccess() *bool
}

type OperateHoldWorkflowExecutionResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// not support query script history, please upgrade engine version to 2.2.2+
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

func (s OperateHoldWorkflowExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateHoldWorkflowExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *OperateHoldWorkflowExecutionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateHoldWorkflowExecutionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateHoldWorkflowExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateHoldWorkflowExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateHoldWorkflowExecutionResponseBody) SetCode(v int32) *OperateHoldWorkflowExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *OperateHoldWorkflowExecutionResponseBody) SetMessage(v string) *OperateHoldWorkflowExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *OperateHoldWorkflowExecutionResponseBody) SetRequestId(v string) *OperateHoldWorkflowExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateHoldWorkflowExecutionResponseBody) SetSuccess(v bool) *OperateHoldWorkflowExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *OperateHoldWorkflowExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
