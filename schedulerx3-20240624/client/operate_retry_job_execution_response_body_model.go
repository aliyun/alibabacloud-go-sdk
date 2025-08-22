// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateRetryJobExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateRetryJobExecutionResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateRetryJobExecutionResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateRetryJobExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateRetryJobExecutionResponseBody
	GetSuccess() *bool
}

type OperateRetryJobExecutionResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 438737AC-760A-57D9-B646-B7EF79426243
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateRetryJobExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateRetryJobExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *OperateRetryJobExecutionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateRetryJobExecutionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateRetryJobExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateRetryJobExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateRetryJobExecutionResponseBody) SetCode(v int32) *OperateRetryJobExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *OperateRetryJobExecutionResponseBody) SetMessage(v string) *OperateRetryJobExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *OperateRetryJobExecutionResponseBody) SetRequestId(v string) *OperateRetryJobExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateRetryJobExecutionResponseBody) SetSuccess(v bool) *OperateRetryJobExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *OperateRetryJobExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
