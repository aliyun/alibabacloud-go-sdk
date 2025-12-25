// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateUnskipJobExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateUnskipJobExecutionResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateUnskipJobExecutionResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateUnskipJobExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateUnskipJobExecutionResponseBody
	GetSuccess() *bool
}

type OperateUnskipJobExecutionResponseBody struct {
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
	// AA3538A0-FBE6-5E31-AD88-A02C6FF0DACC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateUnskipJobExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateUnskipJobExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *OperateUnskipJobExecutionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateUnskipJobExecutionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateUnskipJobExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateUnskipJobExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateUnskipJobExecutionResponseBody) SetCode(v int32) *OperateUnskipJobExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *OperateUnskipJobExecutionResponseBody) SetMessage(v string) *OperateUnskipJobExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *OperateUnskipJobExecutionResponseBody) SetRequestId(v string) *OperateUnskipJobExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateUnskipJobExecutionResponseBody) SetSuccess(v bool) *OperateUnskipJobExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *OperateUnskipJobExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
