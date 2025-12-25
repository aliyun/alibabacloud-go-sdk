// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateMarkSuccessJobExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateMarkSuccessJobExecutionResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateMarkSuccessJobExecutionResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateMarkSuccessJobExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateMarkSuccessJobExecutionResponseBody
	GetSuccess() *bool
}

type OperateMarkSuccessJobExecutionResponseBody struct {
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

func (s OperateMarkSuccessJobExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateMarkSuccessJobExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *OperateMarkSuccessJobExecutionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateMarkSuccessJobExecutionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateMarkSuccessJobExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateMarkSuccessJobExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateMarkSuccessJobExecutionResponseBody) SetCode(v int32) *OperateMarkSuccessJobExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *OperateMarkSuccessJobExecutionResponseBody) SetMessage(v string) *OperateMarkSuccessJobExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *OperateMarkSuccessJobExecutionResponseBody) SetRequestId(v string) *OperateMarkSuccessJobExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateMarkSuccessJobExecutionResponseBody) SetSuccess(v bool) *OperateMarkSuccessJobExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *OperateMarkSuccessJobExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
