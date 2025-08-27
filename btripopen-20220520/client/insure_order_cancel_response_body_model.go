// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderCancelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InsureOrderCancelResponseBody
	GetCode() *string
	SetMessage(v string) *InsureOrderCancelResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsureOrderCancelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InsureOrderCancelResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InsureOrderCancelResponseBody
	GetTraceId() *string
}

type InsureOrderCancelResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 210bc41416861901778051918d1942
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InsureOrderCancelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderCancelResponseBody) GoString() string {
	return s.String()
}

func (s *InsureOrderCancelResponseBody) GetCode() *string {
	return s.Code
}

func (s *InsureOrderCancelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsureOrderCancelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsureOrderCancelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InsureOrderCancelResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InsureOrderCancelResponseBody) SetCode(v string) *InsureOrderCancelResponseBody {
	s.Code = &v
	return s
}

func (s *InsureOrderCancelResponseBody) SetMessage(v string) *InsureOrderCancelResponseBody {
	s.Message = &v
	return s
}

func (s *InsureOrderCancelResponseBody) SetRequestId(v string) *InsureOrderCancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsureOrderCancelResponseBody) SetSuccess(v bool) *InsureOrderCancelResponseBody {
	s.Success = &v
	return s
}

func (s *InsureOrderCancelResponseBody) SetTraceId(v string) *InsureOrderCancelResponseBody {
	s.TraceId = &v
	return s
}

func (s *InsureOrderCancelResponseBody) Validate() error {
	return dara.Validate(s)
}
