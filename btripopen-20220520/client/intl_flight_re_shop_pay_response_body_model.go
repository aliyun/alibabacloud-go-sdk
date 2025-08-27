// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopPayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightReShopPayResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightReShopPayResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightReShopPayResponseBodyModule) *IntlFlightReShopPayResponseBody
	GetModule() *IntlFlightReShopPayResponseBodyModule
	SetRequestId(v string) *IntlFlightReShopPayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightReShopPayResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightReShopPayResponseBody
	GetTraceId() *string
}

type IntlFlightReShopPayResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IntlFlightReShopPayResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// 210bc4b116835992457938931db4de
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IntlFlightReShopPayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopPayResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopPayResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightReShopPayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightReShopPayResponseBody) GetModule() *IntlFlightReShopPayResponseBodyModule {
	return s.Module
}

func (s *IntlFlightReShopPayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightReShopPayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightReShopPayResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightReShopPayResponseBody) SetCode(v string) *IntlFlightReShopPayResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightReShopPayResponseBody) SetMessage(v string) *IntlFlightReShopPayResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightReShopPayResponseBody) SetModule(v *IntlFlightReShopPayResponseBodyModule) *IntlFlightReShopPayResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightReShopPayResponseBody) SetRequestId(v string) *IntlFlightReShopPayResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightReShopPayResponseBody) SetSuccess(v bool) *IntlFlightReShopPayResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightReShopPayResponseBody) SetTraceId(v string) *IntlFlightReShopPayResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightReShopPayResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopPayResponseBodyModule struct {
	// example:
	//
	// 100000
	ActualPayPrice *int64 `json:"actual_pay_price,omitempty" xml:"actual_pay_price,omitempty"`
	// example:
	//
	// 0
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
}

func (s IntlFlightReShopPayResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopPayResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopPayResponseBodyModule) GetActualPayPrice() *int64 {
	return s.ActualPayPrice
}

func (s *IntlFlightReShopPayResponseBodyModule) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *IntlFlightReShopPayResponseBodyModule) SetActualPayPrice(v int64) *IntlFlightReShopPayResponseBodyModule {
	s.ActualPayPrice = &v
	return s
}

func (s *IntlFlightReShopPayResponseBodyModule) SetPayStatus(v int32) *IntlFlightReShopPayResponseBodyModule {
	s.PayStatus = &v
	return s
}

func (s *IntlFlightReShopPayResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
