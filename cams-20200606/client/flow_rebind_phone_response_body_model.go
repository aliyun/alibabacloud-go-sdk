// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowRebindPhoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *FlowRebindPhoneResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *FlowRebindPhoneResponseBody
	GetCode() *string
	SetMessage(v string) *FlowRebindPhoneResponseBody
	GetMessage() *string
	SetModel(v bool) *FlowRebindPhoneResponseBody
	GetModel() *bool
	SetRequestId(v string) *FlowRebindPhoneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlowRebindPhoneResponseBody
	GetSuccess() *bool
}

type FlowRebindPhoneResponseBody struct {
	// The details about the access denial.
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
	// Indicates whether the operation is successful. Valid values:
	//
	// - true: The operation is successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
	Model *bool `json:"Model,omitempty" xml:"Model,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// - true: The operation is successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FlowRebindPhoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlowRebindPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *FlowRebindPhoneResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *FlowRebindPhoneResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlowRebindPhoneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlowRebindPhoneResponseBody) GetModel() *bool {
	return s.Model
}

func (s *FlowRebindPhoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlowRebindPhoneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlowRebindPhoneResponseBody) SetAccessDeniedDetail(v string) *FlowRebindPhoneResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *FlowRebindPhoneResponseBody) SetCode(v string) *FlowRebindPhoneResponseBody {
	s.Code = &v
	return s
}

func (s *FlowRebindPhoneResponseBody) SetMessage(v string) *FlowRebindPhoneResponseBody {
	s.Message = &v
	return s
}

func (s *FlowRebindPhoneResponseBody) SetModel(v bool) *FlowRebindPhoneResponseBody {
	s.Model = &v
	return s
}

func (s *FlowRebindPhoneResponseBody) SetRequestId(v string) *FlowRebindPhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlowRebindPhoneResponseBody) SetSuccess(v bool) *FlowRebindPhoneResponseBody {
	s.Success = &v
	return s
}

func (s *FlowRebindPhoneResponseBody) Validate() error {
	return dara.Validate(s)
}
