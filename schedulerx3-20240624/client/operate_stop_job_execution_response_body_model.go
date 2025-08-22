// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateStopJobExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateStopJobExecutionResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateStopJobExecutionResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateStopJobExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateStopJobExecutionResponseBody
	GetSuccess() *bool
}

type OperateStopJobExecutionResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E82D8B33-204D-58E1-8F56-909F6B48F3D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateStopJobExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateStopJobExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *OperateStopJobExecutionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateStopJobExecutionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateStopJobExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateStopJobExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateStopJobExecutionResponseBody) SetCode(v int32) *OperateStopJobExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *OperateStopJobExecutionResponseBody) SetMessage(v string) *OperateStopJobExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *OperateStopJobExecutionResponseBody) SetRequestId(v string) *OperateStopJobExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateStopJobExecutionResponseBody) SetSuccess(v bool) *OperateStopJobExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *OperateStopJobExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
