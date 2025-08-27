// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightCreateOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightCreateOrderResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightCreateOrderResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightCreateOrderResponseBodyModule) *IntlFlightCreateOrderResponseBody
	GetModule() *IntlFlightCreateOrderResponseBodyModule
	SetRequestId(v string) *IntlFlightCreateOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightCreateOrderResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightCreateOrderResponseBody
	GetTraceId() *string
}

type IntlFlightCreateOrderResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IntlFlightCreateOrderResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2103ad1116826479016562032da98c
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IntlFlightCreateOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightCreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightCreateOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightCreateOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightCreateOrderResponseBody) GetModule() *IntlFlightCreateOrderResponseBodyModule {
	return s.Module
}

func (s *IntlFlightCreateOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightCreateOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightCreateOrderResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightCreateOrderResponseBody) SetCode(v string) *IntlFlightCreateOrderResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightCreateOrderResponseBody) SetMessage(v string) *IntlFlightCreateOrderResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightCreateOrderResponseBody) SetModule(v *IntlFlightCreateOrderResponseBodyModule) *IntlFlightCreateOrderResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightCreateOrderResponseBody) SetRequestId(v string) *IntlFlightCreateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightCreateOrderResponseBody) SetSuccess(v bool) *IntlFlightCreateOrderResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightCreateOrderResponseBody) SetTraceId(v string) *IntlFlightCreateOrderResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightCreateOrderResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightCreateOrderResponseBodyModule struct {
	// example:
	//
	// deb6372db8194f1c94c23bc4fadc508d
	AsyncCreateOrderKey *string `json:"async_create_order_key,omitempty" xml:"async_create_order_key,omitempty"`
	// example:
	//
	// 1003038197806523003
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// F11378244642107293696
	OutOrderId    *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	PayLatestTime *string `json:"pay_latest_time,omitempty" xml:"pay_latest_time,omitempty"`
	// example:
	//
	// 0
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// example:
	//
	// -1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 22300
	TotalPrice *int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
}

func (s IntlFlightCreateOrderResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightCreateOrderResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightCreateOrderResponseBodyModule) GetAsyncCreateOrderKey() *string {
	return s.AsyncCreateOrderKey
}

func (s *IntlFlightCreateOrderResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightCreateOrderResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightCreateOrderResponseBodyModule) GetPayLatestTime() *string {
	return s.PayLatestTime
}

func (s *IntlFlightCreateOrderResponseBodyModule) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *IntlFlightCreateOrderResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *IntlFlightCreateOrderResponseBodyModule) GetTotalPrice() *int64 {
	return s.TotalPrice
}

func (s *IntlFlightCreateOrderResponseBodyModule) SetAsyncCreateOrderKey(v string) *IntlFlightCreateOrderResponseBodyModule {
	s.AsyncCreateOrderKey = &v
	return s
}

func (s *IntlFlightCreateOrderResponseBodyModule) SetOrderId(v string) *IntlFlightCreateOrderResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *IntlFlightCreateOrderResponseBodyModule) SetOutOrderId(v string) *IntlFlightCreateOrderResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightCreateOrderResponseBodyModule) SetPayLatestTime(v string) *IntlFlightCreateOrderResponseBodyModule {
	s.PayLatestTime = &v
	return s
}

func (s *IntlFlightCreateOrderResponseBodyModule) SetPayStatus(v int32) *IntlFlightCreateOrderResponseBodyModule {
	s.PayStatus = &v
	return s
}

func (s *IntlFlightCreateOrderResponseBodyModule) SetStatus(v int32) *IntlFlightCreateOrderResponseBodyModule {
	s.Status = &v
	return s
}

func (s *IntlFlightCreateOrderResponseBodyModule) SetTotalPrice(v int64) *IntlFlightCreateOrderResponseBodyModule {
	s.TotalPrice = &v
	return s
}

func (s *IntlFlightCreateOrderResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
