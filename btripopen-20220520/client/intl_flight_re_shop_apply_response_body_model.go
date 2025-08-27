// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopApplyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightReShopApplyResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightReShopApplyResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightReShopApplyResponseBodyModule) *IntlFlightReShopApplyResponseBody
	GetModule() *IntlFlightReShopApplyResponseBodyModule
	SetRequestId(v string) *IntlFlightReShopApplyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightReShopApplyResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightReShopApplyResponseBody
	GetTraceId() *string
}

type IntlFlightReShopApplyResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IntlFlightReShopApplyResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 88BA5020-561C-51F5-8E73-6659729913F0
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

func (s IntlFlightReShopApplyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopApplyResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopApplyResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightReShopApplyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightReShopApplyResponseBody) GetModule() *IntlFlightReShopApplyResponseBodyModule {
	return s.Module
}

func (s *IntlFlightReShopApplyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightReShopApplyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightReShopApplyResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightReShopApplyResponseBody) SetCode(v string) *IntlFlightReShopApplyResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightReShopApplyResponseBody) SetMessage(v string) *IntlFlightReShopApplyResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightReShopApplyResponseBody) SetModule(v *IntlFlightReShopApplyResponseBodyModule) *IntlFlightReShopApplyResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightReShopApplyResponseBody) SetRequestId(v string) *IntlFlightReShopApplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightReShopApplyResponseBody) SetSuccess(v bool) *IntlFlightReShopApplyResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightReShopApplyResponseBody) SetTraceId(v string) *IntlFlightReShopApplyResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightReShopApplyResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopApplyResponseBodyModule struct {
	// example:
	//
	// asyncKey_2390u230slgw023
	AsyncApplyKey *string `json:"async_apply_key,omitempty" xml:"async_apply_key,omitempty"`
	// example:
	//
	// JPM20241024354
	OutReShopApplyId *string `json:"out_re_shop_apply_id,omitempty" xml:"out_re_shop_apply_id,omitempty"`
	// example:
	//
	// 1009035199432
	ReShopApplyId *string `json:"re_shop_apply_id,omitempty" xml:"re_shop_apply_id,omitempty"`
}

func (s IntlFlightReShopApplyResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopApplyResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopApplyResponseBodyModule) GetAsyncApplyKey() *string {
	return s.AsyncApplyKey
}

func (s *IntlFlightReShopApplyResponseBodyModule) GetOutReShopApplyId() *string {
	return s.OutReShopApplyId
}

func (s *IntlFlightReShopApplyResponseBodyModule) GetReShopApplyId() *string {
	return s.ReShopApplyId
}

func (s *IntlFlightReShopApplyResponseBodyModule) SetAsyncApplyKey(v string) *IntlFlightReShopApplyResponseBodyModule {
	s.AsyncApplyKey = &v
	return s
}

func (s *IntlFlightReShopApplyResponseBodyModule) SetOutReShopApplyId(v string) *IntlFlightReShopApplyResponseBodyModule {
	s.OutReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopApplyResponseBodyModule) SetReShopApplyId(v string) *IntlFlightReShopApplyResponseBodyModule {
	s.ReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopApplyResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
