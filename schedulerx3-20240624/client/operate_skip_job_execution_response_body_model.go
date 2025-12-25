// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateSkipJobExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateSkipJobExecutionResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateSkipJobExecutionResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateSkipJobExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateSkipJobExecutionResponseBody
	GetSuccess() *bool
}

type OperateSkipJobExecutionResponseBody struct {
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
	// C78E2AD2-5985-515B-BAD2-31A248AFC263
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateSkipJobExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateSkipJobExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *OperateSkipJobExecutionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateSkipJobExecutionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateSkipJobExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateSkipJobExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateSkipJobExecutionResponseBody) SetCode(v int32) *OperateSkipJobExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *OperateSkipJobExecutionResponseBody) SetMessage(v string) *OperateSkipJobExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *OperateSkipJobExecutionResponseBody) SetRequestId(v string) *OperateSkipJobExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateSkipJobExecutionResponseBody) SetSuccess(v bool) *OperateSkipJobExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *OperateSkipJobExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
