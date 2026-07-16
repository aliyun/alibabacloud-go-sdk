// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightReShopCreateResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightReShopCreateResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightReShopCreateResponseBodyModule) *IntlFlightReShopCreateResponseBody
	GetModule() *IntlFlightReShopCreateResponseBodyModule
	SetRequestId(v string) *IntlFlightReShopCreateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightReShopCreateResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightReShopCreateResponseBody
	GetTraceId() *string
}

type IntlFlightReShopCreateResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message *string                                   `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IntlFlightReShopCreateResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s IntlFlightReShopCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopCreateResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopCreateResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightReShopCreateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightReShopCreateResponseBody) GetModule() *IntlFlightReShopCreateResponseBodyModule {
	return s.Module
}

func (s *IntlFlightReShopCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightReShopCreateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightReShopCreateResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightReShopCreateResponseBody) SetCode(v string) *IntlFlightReShopCreateResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightReShopCreateResponseBody) SetMessage(v string) *IntlFlightReShopCreateResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightReShopCreateResponseBody) SetModule(v *IntlFlightReShopCreateResponseBodyModule) *IntlFlightReShopCreateResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightReShopCreateResponseBody) SetRequestId(v string) *IntlFlightReShopCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightReShopCreateResponseBody) SetSuccess(v bool) *IntlFlightReShopCreateResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightReShopCreateResponseBody) SetTraceId(v string) *IntlFlightReShopCreateResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightReShopCreateResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IntlFlightReShopCreateResponseBodyModule struct {
	AsyncApplyKey     *string `json:"async_apply_key,omitempty" xml:"async_apply_key,omitempty"`
	NeedRetry         *bool   `json:"need_retry,omitempty" xml:"need_retry,omitempty"`
	NextRetryInterval *int64  `json:"next_retry_interval,omitempty" xml:"next_retry_interval,omitempty"`
	OutReShopApplyId  *string `json:"out_re_shop_apply_id,omitempty" xml:"out_re_shop_apply_id,omitempty"`
	ReShopApplyId     *string `json:"re_shop_apply_id,omitempty" xml:"re_shop_apply_id,omitempty"`
}

func (s IntlFlightReShopCreateResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopCreateResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopCreateResponseBodyModule) GetAsyncApplyKey() *string {
	return s.AsyncApplyKey
}

func (s *IntlFlightReShopCreateResponseBodyModule) GetNeedRetry() *bool {
	return s.NeedRetry
}

func (s *IntlFlightReShopCreateResponseBodyModule) GetNextRetryInterval() *int64 {
	return s.NextRetryInterval
}

func (s *IntlFlightReShopCreateResponseBodyModule) GetOutReShopApplyId() *string {
	return s.OutReShopApplyId
}

func (s *IntlFlightReShopCreateResponseBodyModule) GetReShopApplyId() *string {
	return s.ReShopApplyId
}

func (s *IntlFlightReShopCreateResponseBodyModule) SetAsyncApplyKey(v string) *IntlFlightReShopCreateResponseBodyModule {
	s.AsyncApplyKey = &v
	return s
}

func (s *IntlFlightReShopCreateResponseBodyModule) SetNeedRetry(v bool) *IntlFlightReShopCreateResponseBodyModule {
	s.NeedRetry = &v
	return s
}

func (s *IntlFlightReShopCreateResponseBodyModule) SetNextRetryInterval(v int64) *IntlFlightReShopCreateResponseBodyModule {
	s.NextRetryInterval = &v
	return s
}

func (s *IntlFlightReShopCreateResponseBodyModule) SetOutReShopApplyId(v string) *IntlFlightReShopCreateResponseBodyModule {
	s.OutReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopCreateResponseBodyModule) SetReShopApplyId(v string) *IntlFlightReShopCreateResponseBodyModule {
	s.ReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopCreateResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
