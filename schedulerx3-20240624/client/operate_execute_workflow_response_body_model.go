// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateExecuteWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateExecuteWorkflowResponseBody
	GetCode() *int32
	SetData(v *OperateExecuteWorkflowResponseBodyData) *OperateExecuteWorkflowResponseBody
	GetData() *OperateExecuteWorkflowResponseBodyData
	SetMessage(v string) *OperateExecuteWorkflowResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateExecuteWorkflowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateExecuteWorkflowResponseBody
	GetSuccess() *bool
}

type OperateExecuteWorkflowResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *OperateExecuteWorkflowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3835AA29-2298-5434-BC53-9CC377CDFD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateExecuteWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateExecuteWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *OperateExecuteWorkflowResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateExecuteWorkflowResponseBody) GetData() *OperateExecuteWorkflowResponseBodyData {
	return s.Data
}

func (s *OperateExecuteWorkflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateExecuteWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateExecuteWorkflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateExecuteWorkflowResponseBody) SetCode(v int32) *OperateExecuteWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *OperateExecuteWorkflowResponseBody) SetData(v *OperateExecuteWorkflowResponseBodyData) *OperateExecuteWorkflowResponseBody {
	s.Data = v
	return s
}

func (s *OperateExecuteWorkflowResponseBody) SetMessage(v string) *OperateExecuteWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *OperateExecuteWorkflowResponseBody) SetRequestId(v string) *OperateExecuteWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateExecuteWorkflowResponseBody) SetSuccess(v bool) *OperateExecuteWorkflowResponseBody {
	s.Success = &v
	return s
}

func (s *OperateExecuteWorkflowResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OperateExecuteWorkflowResponseBodyData struct {
	// example:
	//
	// 100
	WorkflowExecutionId *string `json:"WorkflowExecutionId,omitempty" xml:"WorkflowExecutionId,omitempty"`
}

func (s OperateExecuteWorkflowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OperateExecuteWorkflowResponseBodyData) GoString() string {
	return s.String()
}

func (s *OperateExecuteWorkflowResponseBodyData) GetWorkflowExecutionId() *string {
	return s.WorkflowExecutionId
}

func (s *OperateExecuteWorkflowResponseBodyData) SetWorkflowExecutionId(v string) *OperateExecuteWorkflowResponseBodyData {
	s.WorkflowExecutionId = &v
	return s
}

func (s *OperateExecuteWorkflowResponseBodyData) Validate() error {
	return dara.Validate(s)
}
