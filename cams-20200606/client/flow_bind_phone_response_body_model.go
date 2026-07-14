// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowBindPhoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *FlowBindPhoneResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *FlowBindPhoneResponseBody
	GetCode() *string
	SetMessage(v string) *FlowBindPhoneResponseBody
	GetMessage() *string
	SetModel(v bool) *FlowBindPhoneResponseBody
	GetModel() *bool
	SetRequestId(v string) *FlowBindPhoneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlowBindPhoneResponseBody
	GetSuccess() *bool
}

type FlowBindPhoneResponseBody struct {
	// Details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code.
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
	// Indicates whether the result was successful. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// false
	Model *bool `json:"Model,omitempty" xml:"Model,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
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

func (s FlowBindPhoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlowBindPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *FlowBindPhoneResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *FlowBindPhoneResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlowBindPhoneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlowBindPhoneResponseBody) GetModel() *bool {
	return s.Model
}

func (s *FlowBindPhoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlowBindPhoneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlowBindPhoneResponseBody) SetAccessDeniedDetail(v string) *FlowBindPhoneResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *FlowBindPhoneResponseBody) SetCode(v string) *FlowBindPhoneResponseBody {
	s.Code = &v
	return s
}

func (s *FlowBindPhoneResponseBody) SetMessage(v string) *FlowBindPhoneResponseBody {
	s.Message = &v
	return s
}

func (s *FlowBindPhoneResponseBody) SetModel(v bool) *FlowBindPhoneResponseBody {
	s.Model = &v
	return s
}

func (s *FlowBindPhoneResponseBody) SetRequestId(v string) *FlowBindPhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlowBindPhoneResponseBody) SetSuccess(v bool) *FlowBindPhoneResponseBody {
	s.Success = &v
	return s
}

func (s *FlowBindPhoneResponseBody) Validate() error {
	return dara.Validate(s)
}
