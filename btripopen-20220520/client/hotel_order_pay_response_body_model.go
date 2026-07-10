// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderPayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelOrderPayResponseBody
	GetCode() *string
	SetMessage(v string) *HotelOrderPayResponseBody
	GetMessage() *string
	SetModule(v bool) *HotelOrderPayResponseBody
	GetModule() *bool
	SetRequestId(v string) *HotelOrderPayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelOrderPayResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelOrderPayResponseBody
	GetTraceId() *string
}

type HotelOrderPayResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Module    *bool   `json:"module,omitempty" xml:"module,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s HotelOrderPayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPayResponseBody) GoString() string {
	return s.String()
}

func (s *HotelOrderPayResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelOrderPayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelOrderPayResponseBody) GetModule() *bool {
	return s.Module
}

func (s *HotelOrderPayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelOrderPayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelOrderPayResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelOrderPayResponseBody) SetCode(v string) *HotelOrderPayResponseBody {
	s.Code = &v
	return s
}

func (s *HotelOrderPayResponseBody) SetMessage(v string) *HotelOrderPayResponseBody {
	s.Message = &v
	return s
}

func (s *HotelOrderPayResponseBody) SetModule(v bool) *HotelOrderPayResponseBody {
	s.Module = &v
	return s
}

func (s *HotelOrderPayResponseBody) SetRequestId(v string) *HotelOrderPayResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelOrderPayResponseBody) SetSuccess(v bool) *HotelOrderPayResponseBody {
	s.Success = &v
	return s
}

func (s *HotelOrderPayResponseBody) SetTraceId(v string) *HotelOrderPayResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelOrderPayResponseBody) Validate() error {
	return dara.Validate(s)
}
