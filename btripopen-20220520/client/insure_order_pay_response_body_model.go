// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderPayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InsureOrderPayResponseBody
	GetCode() *string
	SetMessage(v string) *InsureOrderPayResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsureOrderPayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InsureOrderPayResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InsureOrderPayResponseBody
	GetTraceId() *string
}

type InsureOrderPayResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InsureOrderPayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderPayResponseBody) GoString() string {
	return s.String()
}

func (s *InsureOrderPayResponseBody) GetCode() *string {
	return s.Code
}

func (s *InsureOrderPayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsureOrderPayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsureOrderPayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InsureOrderPayResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InsureOrderPayResponseBody) SetCode(v string) *InsureOrderPayResponseBody {
	s.Code = &v
	return s
}

func (s *InsureOrderPayResponseBody) SetMessage(v string) *InsureOrderPayResponseBody {
	s.Message = &v
	return s
}

func (s *InsureOrderPayResponseBody) SetRequestId(v string) *InsureOrderPayResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsureOrderPayResponseBody) SetSuccess(v bool) *InsureOrderPayResponseBody {
	s.Success = &v
	return s
}

func (s *InsureOrderPayResponseBody) SetTraceId(v string) *InsureOrderPayResponseBody {
	s.TraceId = &v
	return s
}

func (s *InsureOrderPayResponseBody) Validate() error {
	return dara.Validate(s)
}
