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
	// Details of access denial; this field is returned only when RAM verification fails.
	//
	// example:
	//
	// 无
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Status code.
	//
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error description message.
	//
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Return result.
	//
	// example:
	//
	// false
	Model *bool `json:"Model,omitempty" xml:"Model,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Values: true: success; false: failure.
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
