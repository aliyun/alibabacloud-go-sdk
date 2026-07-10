// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealOrderListQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MealOrderListQueryResponseBody
	GetCode() *string
	SetMessage(v string) *MealOrderListQueryResponseBody
	GetMessage() *string
	SetModule(v *MealOrderListQueryResponseBodyModule) *MealOrderListQueryResponseBody
	GetModule() *MealOrderListQueryResponseBodyModule
	SetRequestId(v string) *MealOrderListQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MealOrderListQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *MealOrderListQueryResponseBody
	GetTraceId() *string
}

type MealOrderListQueryResponseBody struct {
	Code      *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                               `json:"message,omitempty" xml:"message,omitempty"`
	Module    *MealOrderListQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                               `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s MealOrderListQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MealOrderListQueryResponseBody) GoString() string {
	return s.String()
}

func (s *MealOrderListQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *MealOrderListQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MealOrderListQueryResponseBody) GetModule() *MealOrderListQueryResponseBodyModule {
	return s.Module
}

func (s *MealOrderListQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MealOrderListQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MealOrderListQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *MealOrderListQueryResponseBody) SetCode(v string) *MealOrderListQueryResponseBody {
	s.Code = &v
	return s
}

func (s *MealOrderListQueryResponseBody) SetMessage(v string) *MealOrderListQueryResponseBody {
	s.Message = &v
	return s
}

func (s *MealOrderListQueryResponseBody) SetModule(v *MealOrderListQueryResponseBodyModule) *MealOrderListQueryResponseBody {
	s.Module = v
	return s
}

func (s *MealOrderListQueryResponseBody) SetRequestId(v string) *MealOrderListQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *MealOrderListQueryResponseBody) SetSuccess(v bool) *MealOrderListQueryResponseBody {
	s.Success = &v
	return s
}

func (s *MealOrderListQueryResponseBody) SetTraceId(v string) *MealOrderListQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *MealOrderListQueryResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MealOrderListQueryResponseBodyModule struct {
	OrderList []*MealOrderListQueryResponseBodyModuleOrderList `json:"order_list,omitempty" xml:"order_list,omitempty" type:"Repeated"`
}

func (s MealOrderListQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s MealOrderListQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *MealOrderListQueryResponseBodyModule) GetOrderList() []*MealOrderListQueryResponseBodyModuleOrderList {
	return s.OrderList
}

func (s *MealOrderListQueryResponseBodyModule) SetOrderList(v []*MealOrderListQueryResponseBodyModuleOrderList) *MealOrderListQueryResponseBodyModule {
	s.OrderList = v
	return s
}

func (s *MealOrderListQueryResponseBodyModule) Validate() error {
	if s.OrderList != nil {
		for _, item := range s.OrderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MealOrderListQueryResponseBodyModuleOrderList struct {
	CorpPayAmount   *int64  `json:"corp_pay_amount,omitempty" xml:"corp_pay_amount,omitempty"`
	MerchantName    *string `json:"merchant_name,omitempty" xml:"merchant_name,omitempty"`
	OrderId         *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderStatus     *int32  `json:"order_status,omitempty" xml:"order_status,omitempty"`
	OrderType       *string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	PayAmount       *int64  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	PersonPayAmount *int64  `json:"person_pay_amount,omitempty" xml:"person_pay_amount,omitempty"`
	SettleTime      *string `json:"settle_time,omitempty" xml:"settle_time,omitempty"`
}

func (s MealOrderListQueryResponseBodyModuleOrderList) String() string {
	return dara.Prettify(s)
}

func (s MealOrderListQueryResponseBodyModuleOrderList) GoString() string {
	return s.String()
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) GetCorpPayAmount() *int64 {
	return s.CorpPayAmount
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) GetMerchantName() *string {
	return s.MerchantName
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) GetOrderId() *string {
	return s.OrderId
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) GetOrderType() *string {
	return s.OrderType
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) GetPayAmount() *int64 {
	return s.PayAmount
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) GetPersonPayAmount() *int64 {
	return s.PersonPayAmount
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) GetSettleTime() *string {
	return s.SettleTime
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) SetCorpPayAmount(v int64) *MealOrderListQueryResponseBodyModuleOrderList {
	s.CorpPayAmount = &v
	return s
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) SetMerchantName(v string) *MealOrderListQueryResponseBodyModuleOrderList {
	s.MerchantName = &v
	return s
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) SetOrderId(v string) *MealOrderListQueryResponseBodyModuleOrderList {
	s.OrderId = &v
	return s
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) SetOrderStatus(v int32) *MealOrderListQueryResponseBodyModuleOrderList {
	s.OrderStatus = &v
	return s
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) SetOrderType(v string) *MealOrderListQueryResponseBodyModuleOrderList {
	s.OrderType = &v
	return s
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) SetPayAmount(v int64) *MealOrderListQueryResponseBodyModuleOrderList {
	s.PayAmount = &v
	return s
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) SetPersonPayAmount(v int64) *MealOrderListQueryResponseBodyModuleOrderList {
	s.PersonPayAmount = &v
	return s
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) SetSettleTime(v string) *MealOrderListQueryResponseBodyModuleOrderList {
	s.SettleTime = &v
	return s
}

func (s *MealOrderListQueryResponseBodyModuleOrderList) Validate() error {
	return dara.Validate(s)
}
