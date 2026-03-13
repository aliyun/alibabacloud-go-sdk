// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCompenInfosForOpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCompenInfosForOpResponseBody
	GetCode() *string
	SetMessage(v string) *QueryCompenInfosForOpResponseBody
	GetMessage() *string
	SetModule(v []*QueryCompenInfosForOpResponseBodyModule) *QueryCompenInfosForOpResponseBody
	GetModule() []*QueryCompenInfosForOpResponseBodyModule
	SetRequestId(v string) *QueryCompenInfosForOpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCompenInfosForOpResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *QueryCompenInfosForOpResponseBody
	GetTraceId() *string
}

type QueryCompenInfosForOpResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                    `json:"message,omitempty" xml:"message,omitempty"`
	Module  []*QueryCompenInfosForOpResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
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
	// 210f079e16603757182131635d866a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s QueryCompenInfosForOpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCompenInfosForOpResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCompenInfosForOpResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCompenInfosForOpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCompenInfosForOpResponseBody) GetModule() []*QueryCompenInfosForOpResponseBodyModule {
	return s.Module
}

func (s *QueryCompenInfosForOpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCompenInfosForOpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCompenInfosForOpResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *QueryCompenInfosForOpResponseBody) SetCode(v string) *QueryCompenInfosForOpResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCompenInfosForOpResponseBody) SetMessage(v string) *QueryCompenInfosForOpResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCompenInfosForOpResponseBody) SetModule(v []*QueryCompenInfosForOpResponseBodyModule) *QueryCompenInfosForOpResponseBody {
	s.Module = v
	return s
}

func (s *QueryCompenInfosForOpResponseBody) SetRequestId(v string) *QueryCompenInfosForOpResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCompenInfosForOpResponseBody) SetSuccess(v bool) *QueryCompenInfosForOpResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCompenInfosForOpResponseBody) SetTraceId(v string) *QueryCompenInfosForOpResponseBody {
	s.TraceId = &v
	return s
}

func (s *QueryCompenInfosForOpResponseBody) Validate() error {
	if s.Module != nil {
		for _, item := range s.Module {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryCompenInfosForOpResponseBodyModule struct {
	// example:
	//
	// 500
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// 1
	Category *int32 `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// 82508250249123456
	CompenId *string `json:"compen_id,omitempty" xml:"compen_id,omitempty"`
	// example:
	//
	// 1
	CompensationType *string `json:"compensation_type,omitempty" xml:"compensation_type,omitempty"`
	// example:
	//
	// 1002086203277812345
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 2025-09-22 20:13:58
	SettleTime *string `json:"settle_time,omitempty" xml:"settle_time,omitempty"`
	// example:
	//
	// 8
	SettleType *int32 `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	// example:
	//
	// 999-12345678
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
}

func (s QueryCompenInfosForOpResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryCompenInfosForOpResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryCompenInfosForOpResponseBodyModule) GetAmount() *int64 {
	return s.Amount
}

func (s *QueryCompenInfosForOpResponseBodyModule) GetCategory() *int32 {
	return s.Category
}

func (s *QueryCompenInfosForOpResponseBodyModule) GetCompenId() *string {
	return s.CompenId
}

func (s *QueryCompenInfosForOpResponseBodyModule) GetCompensationType() *string {
	return s.CompensationType
}

func (s *QueryCompenInfosForOpResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryCompenInfosForOpResponseBodyModule) GetSettleTime() *string {
	return s.SettleTime
}

func (s *QueryCompenInfosForOpResponseBodyModule) GetSettleType() *int32 {
	return s.SettleType
}

func (s *QueryCompenInfosForOpResponseBodyModule) GetTicketNo() *string {
	return s.TicketNo
}

func (s *QueryCompenInfosForOpResponseBodyModule) SetAmount(v int64) *QueryCompenInfosForOpResponseBodyModule {
	s.Amount = &v
	return s
}

func (s *QueryCompenInfosForOpResponseBodyModule) SetCategory(v int32) *QueryCompenInfosForOpResponseBodyModule {
	s.Category = &v
	return s
}

func (s *QueryCompenInfosForOpResponseBodyModule) SetCompenId(v string) *QueryCompenInfosForOpResponseBodyModule {
	s.CompenId = &v
	return s
}

func (s *QueryCompenInfosForOpResponseBodyModule) SetCompensationType(v string) *QueryCompenInfosForOpResponseBodyModule {
	s.CompensationType = &v
	return s
}

func (s *QueryCompenInfosForOpResponseBodyModule) SetOrderId(v string) *QueryCompenInfosForOpResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *QueryCompenInfosForOpResponseBodyModule) SetSettleTime(v string) *QueryCompenInfosForOpResponseBodyModule {
	s.SettleTime = &v
	return s
}

func (s *QueryCompenInfosForOpResponseBodyModule) SetSettleType(v int32) *QueryCompenInfosForOpResponseBodyModule {
	s.SettleType = &v
	return s
}

func (s *QueryCompenInfosForOpResponseBodyModule) SetTicketNo(v string) *QueryCompenInfosForOpResponseBodyModule {
	s.TicketNo = &v
	return s
}

func (s *QueryCompenInfosForOpResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
