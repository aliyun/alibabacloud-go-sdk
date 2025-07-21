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
	// Access denied details.
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
	// Error message.
	//
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request result data.
	//
	// example:
	//
	// false
	Model *bool `json:"Model,omitempty" xml:"Model,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the operation was successful. Values: true: success; false: failure.
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
