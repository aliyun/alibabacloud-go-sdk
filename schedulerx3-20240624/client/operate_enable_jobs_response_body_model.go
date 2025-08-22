// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateEnableJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateEnableJobsResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateEnableJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateEnableJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateEnableJobsResponseBody
	GetSuccess() *bool
}

type OperateEnableJobsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4CC4132F-B798-5D6E-9F06-D44B33E417E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateEnableJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateEnableJobsResponseBody) GoString() string {
	return s.String()
}

func (s *OperateEnableJobsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateEnableJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateEnableJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateEnableJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateEnableJobsResponseBody) SetCode(v int32) *OperateEnableJobsResponseBody {
	s.Code = &v
	return s
}

func (s *OperateEnableJobsResponseBody) SetMessage(v string) *OperateEnableJobsResponseBody {
	s.Message = &v
	return s
}

func (s *OperateEnableJobsResponseBody) SetRequestId(v string) *OperateEnableJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateEnableJobsResponseBody) SetSuccess(v bool) *OperateEnableJobsResponseBody {
	s.Success = &v
	return s
}

func (s *OperateEnableJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
