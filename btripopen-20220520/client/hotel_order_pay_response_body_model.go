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
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Module *bool `json:"module,omitempty" xml:"module,omitempty"`
	// requestId
	//
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210e842b16611337974412836dae27
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
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
