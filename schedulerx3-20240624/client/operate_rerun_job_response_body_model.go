// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateRerunJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateRerunJobResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateRerunJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateRerunJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateRerunJobResponseBody
	GetSuccess() *bool
}

type OperateRerunJobResponseBody struct {
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
	// BCDF4006-C8A1-5F83-9368-588347D3EE84
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateRerunJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateRerunJobResponseBody) GoString() string {
	return s.String()
}

func (s *OperateRerunJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateRerunJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateRerunJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateRerunJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateRerunJobResponseBody) SetCode(v int32) *OperateRerunJobResponseBody {
	s.Code = &v
	return s
}

func (s *OperateRerunJobResponseBody) SetMessage(v string) *OperateRerunJobResponseBody {
	s.Message = &v
	return s
}

func (s *OperateRerunJobResponseBody) SetRequestId(v string) *OperateRerunJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateRerunJobResponseBody) SetSuccess(v bool) *OperateRerunJobResponseBody {
	s.Success = &v
	return s
}

func (s *OperateRerunJobResponseBody) Validate() error {
	return dara.Validate(s)
}
