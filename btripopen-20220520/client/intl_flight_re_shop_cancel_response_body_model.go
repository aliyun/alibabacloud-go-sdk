// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopCancelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightReShopCancelResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightReShopCancelResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightReShopCancelResponseBodyModule) *IntlFlightReShopCancelResponseBody
	GetModule() *IntlFlightReShopCancelResponseBodyModule
	SetRequestId(v string) *IntlFlightReShopCancelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightReShopCancelResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightReShopCancelResponseBody
	GetTraceId() *string
}

type IntlFlightReShopCancelResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message *string                                   `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IntlFlightReShopCancelResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s IntlFlightReShopCancelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopCancelResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopCancelResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightReShopCancelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightReShopCancelResponseBody) GetModule() *IntlFlightReShopCancelResponseBodyModule {
	return s.Module
}

func (s *IntlFlightReShopCancelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightReShopCancelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightReShopCancelResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightReShopCancelResponseBody) SetCode(v string) *IntlFlightReShopCancelResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightReShopCancelResponseBody) SetMessage(v string) *IntlFlightReShopCancelResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightReShopCancelResponseBody) SetModule(v *IntlFlightReShopCancelResponseBodyModule) *IntlFlightReShopCancelResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightReShopCancelResponseBody) SetRequestId(v string) *IntlFlightReShopCancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightReShopCancelResponseBody) SetSuccess(v bool) *IntlFlightReShopCancelResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightReShopCancelResponseBody) SetTraceId(v string) *IntlFlightReShopCancelResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightReShopCancelResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopCancelResponseBodyModule struct {
	// example:
	//
	// 2023-08-10 17:45:32
	CancelTime *string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
}

func (s IntlFlightReShopCancelResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopCancelResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopCancelResponseBodyModule) GetCancelTime() *string {
	return s.CancelTime
}

func (s *IntlFlightReShopCancelResponseBodyModule) SetCancelTime(v string) *IntlFlightReShopCancelResponseBodyModule {
	s.CancelTime = &v
	return s
}

func (s *IntlFlightReShopCancelResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
