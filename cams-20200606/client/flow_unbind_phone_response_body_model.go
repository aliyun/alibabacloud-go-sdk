// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowUnbindPhoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *FlowUnbindPhoneResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *FlowUnbindPhoneResponseBody
	GetCode() *string
	SetMessage(v string) *FlowUnbindPhoneResponseBody
	GetMessage() *string
	SetModel(v bool) *FlowUnbindPhoneResponseBody
	GetModel() *bool
	SetRequestId(v string) *FlowUnbindPhoneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlowUnbindPhoneResponseBody
	GetSuccess() *bool
}

type FlowUnbindPhoneResponseBody struct {
	// Details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
	Model     *bool   `json:"Model,omitempty" xml:"Model,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FlowUnbindPhoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlowUnbindPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *FlowUnbindPhoneResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *FlowUnbindPhoneResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlowUnbindPhoneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlowUnbindPhoneResponseBody) GetModel() *bool {
	return s.Model
}

func (s *FlowUnbindPhoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlowUnbindPhoneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlowUnbindPhoneResponseBody) SetAccessDeniedDetail(v string) *FlowUnbindPhoneResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *FlowUnbindPhoneResponseBody) SetCode(v string) *FlowUnbindPhoneResponseBody {
	s.Code = &v
	return s
}

func (s *FlowUnbindPhoneResponseBody) SetMessage(v string) *FlowUnbindPhoneResponseBody {
	s.Message = &v
	return s
}

func (s *FlowUnbindPhoneResponseBody) SetModel(v bool) *FlowUnbindPhoneResponseBody {
	s.Model = &v
	return s
}

func (s *FlowUnbindPhoneResponseBody) SetRequestId(v string) *FlowUnbindPhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlowUnbindPhoneResponseBody) SetSuccess(v bool) *FlowUnbindPhoneResponseBody {
	s.Success = &v
	return s
}

func (s *FlowUnbindPhoneResponseBody) Validate() error {
	return dara.Validate(s)
}
