// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderChangeApplyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelOrderChangeApplyResponseBody
	GetCode() *string
	SetMessage(v string) *HotelOrderChangeApplyResponseBody
	GetMessage() *string
	SetModule(v *HotelOrderChangeApplyResponseBodyModule) *HotelOrderChangeApplyResponseBody
	GetModule() *HotelOrderChangeApplyResponseBodyModule
	SetRequestId(v string) *HotelOrderChangeApplyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelOrderChangeApplyResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelOrderChangeApplyResponseBody
	GetTraceId() *string
}

type HotelOrderChangeApplyResponseBody struct {
	// example:
	//
	// 0
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *HotelOrderChangeApplyResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 6E9ABA14-5135-58FB-9DFC-C751B5855605
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 215045ec17018285034106091e8ba9
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s HotelOrderChangeApplyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderChangeApplyResponseBody) GoString() string {
	return s.String()
}

func (s *HotelOrderChangeApplyResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelOrderChangeApplyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelOrderChangeApplyResponseBody) GetModule() *HotelOrderChangeApplyResponseBodyModule {
	return s.Module
}

func (s *HotelOrderChangeApplyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelOrderChangeApplyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelOrderChangeApplyResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelOrderChangeApplyResponseBody) SetCode(v string) *HotelOrderChangeApplyResponseBody {
	s.Code = &v
	return s
}

func (s *HotelOrderChangeApplyResponseBody) SetMessage(v string) *HotelOrderChangeApplyResponseBody {
	s.Message = &v
	return s
}

func (s *HotelOrderChangeApplyResponseBody) SetModule(v *HotelOrderChangeApplyResponseBodyModule) *HotelOrderChangeApplyResponseBody {
	s.Module = v
	return s
}

func (s *HotelOrderChangeApplyResponseBody) SetRequestId(v string) *HotelOrderChangeApplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelOrderChangeApplyResponseBody) SetSuccess(v bool) *HotelOrderChangeApplyResponseBody {
	s.Success = &v
	return s
}

func (s *HotelOrderChangeApplyResponseBody) SetTraceId(v string) *HotelOrderChangeApplyResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelOrderChangeApplyResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelOrderChangeApplyResponseBodyModule struct {
	// example:
	//
	// 123445443444
	ChangeOrderId *string `json:"change_order_id,omitempty" xml:"change_order_id,omitempty"`
}

func (s HotelOrderChangeApplyResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderChangeApplyResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelOrderChangeApplyResponseBodyModule) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *HotelOrderChangeApplyResponseBodyModule) SetChangeOrderId(v string) *HotelOrderChangeApplyResponseBodyModule {
	s.ChangeOrderId = &v
	return s
}

func (s *HotelOrderChangeApplyResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
