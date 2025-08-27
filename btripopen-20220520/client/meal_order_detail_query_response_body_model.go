// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealOrderDetailQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MealOrderDetailQueryResponseBody
	GetCode() *string
	SetMessage(v string) *MealOrderDetailQueryResponseBody
	GetMessage() *string
	SetModule(v *MealOrderDetailQueryResponseBodyModule) *MealOrderDetailQueryResponseBody
	GetModule() *MealOrderDetailQueryResponseBodyModule
	SetRequestId(v string) *MealOrderDetailQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MealOrderDetailQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *MealOrderDetailQueryResponseBody
	GetTraceId() *string
}

type MealOrderDetailQueryResponseBody struct {
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// None
	Message *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	Module  *MealOrderDetailQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 9BCDD5DE-E6CB-5C25-93B9-9BE178A0AA56
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210f079e16603757182131635d866a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s MealOrderDetailQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MealOrderDetailQueryResponseBody) GoString() string {
	return s.String()
}

func (s *MealOrderDetailQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *MealOrderDetailQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MealOrderDetailQueryResponseBody) GetModule() *MealOrderDetailQueryResponseBodyModule {
	return s.Module
}

func (s *MealOrderDetailQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MealOrderDetailQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MealOrderDetailQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *MealOrderDetailQueryResponseBody) SetCode(v string) *MealOrderDetailQueryResponseBody {
	s.Code = &v
	return s
}

func (s *MealOrderDetailQueryResponseBody) SetMessage(v string) *MealOrderDetailQueryResponseBody {
	s.Message = &v
	return s
}

func (s *MealOrderDetailQueryResponseBody) SetModule(v *MealOrderDetailQueryResponseBodyModule) *MealOrderDetailQueryResponseBody {
	s.Module = v
	return s
}

func (s *MealOrderDetailQueryResponseBody) SetRequestId(v string) *MealOrderDetailQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *MealOrderDetailQueryResponseBody) SetSuccess(v bool) *MealOrderDetailQueryResponseBody {
	s.Success = &v
	return s
}

func (s *MealOrderDetailQueryResponseBody) SetTraceId(v string) *MealOrderDetailQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *MealOrderDetailQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type MealOrderDetailQueryResponseBodyModule struct {
	ApplyId *int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// 2024032522001423601409393178
	CorpCodeOrderId *string `json:"corp_code_order_id,omitempty" xml:"corp_code_order_id,omitempty"`
	// example:
	//
	// dingaa15ca45cba9ee744a5
	CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// example:
	//
	// 100
	CorpPayAmount *int64 `json:"corp_pay_amount,omitempty" xml:"corp_pay_amount,omitempty"`
	// example:
	//
	// 100
	CorpRefundAmount *int64  `json:"corp_refund_amount,omitempty" xml:"corp_refund_amount,omitempty"`
	MealReason       *string `json:"meal_reason,omitempty" xml:"meal_reason,omitempty"`
	MerchantName     *string `json:"merchant_name,omitempty" xml:"merchant_name,omitempty"`
	// example:
	//
	// 1034088398073015303
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1
	OrderStatus *int32 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// example:
	//
	// 1
	OrderSubStatus *int32  `json:"order_sub_status,omitempty" xml:"order_sub_status,omitempty"`
	OrderType      *string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// example:
	//
	// 100
	PayAmount *int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	// example:
	//
	// 1
	PayType *int32 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// example:
	//
	// 100
	PersonPayAmount *int64 `json:"person_pay_amount,omitempty" xml:"person_pay_amount,omitempty"`
	// example:
	//
	// 100
	PersonRefundAmount *int64 `json:"person_refund_amount,omitempty" xml:"person_refund_amount,omitempty"`
	// example:
	//
	// 100
	RefundAmount *int64  `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	SceneName    *string `json:"scene_name,omitempty" xml:"scene_name,omitempty"`
	// example:
	//
	// 1711702782
	SettleTime       *string `json:"settle_time,omitempty" xml:"settle_time,omitempty"`
	ThirdPartApplyId *string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	// example:
	//
	// 17509140000
	UserAlipayId *string `json:"user_alipay_id,omitempty" xml:"user_alipay_id,omitempty"`
	// example:
	//
	// 014134681304653773
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s MealOrderDetailQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s MealOrderDetailQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *MealOrderDetailQueryResponseBodyModule) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *MealOrderDetailQueryResponseBodyModule) GetCorpCodeOrderId() *string {
	return s.CorpCodeOrderId
}

func (s *MealOrderDetailQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *MealOrderDetailQueryResponseBodyModule) GetCorpPayAmount() *int64 {
	return s.CorpPayAmount
}

func (s *MealOrderDetailQueryResponseBodyModule) GetCorpRefundAmount() *int64 {
	return s.CorpRefundAmount
}

func (s *MealOrderDetailQueryResponseBodyModule) GetMealReason() *string {
	return s.MealReason
}

func (s *MealOrderDetailQueryResponseBodyModule) GetMerchantName() *string {
	return s.MerchantName
}

func (s *MealOrderDetailQueryResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *MealOrderDetailQueryResponseBodyModule) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *MealOrderDetailQueryResponseBodyModule) GetOrderSubStatus() *int32 {
	return s.OrderSubStatus
}

func (s *MealOrderDetailQueryResponseBodyModule) GetOrderType() *string {
	return s.OrderType
}

func (s *MealOrderDetailQueryResponseBodyModule) GetPayAmount() *int64 {
	return s.PayAmount
}

func (s *MealOrderDetailQueryResponseBodyModule) GetPayType() *int32 {
	return s.PayType
}

func (s *MealOrderDetailQueryResponseBodyModule) GetPersonPayAmount() *int64 {
	return s.PersonPayAmount
}

func (s *MealOrderDetailQueryResponseBodyModule) GetPersonRefundAmount() *int64 {
	return s.PersonRefundAmount
}

func (s *MealOrderDetailQueryResponseBodyModule) GetRefundAmount() *int64 {
	return s.RefundAmount
}

func (s *MealOrderDetailQueryResponseBodyModule) GetSceneName() *string {
	return s.SceneName
}

func (s *MealOrderDetailQueryResponseBodyModule) GetSettleTime() *string {
	return s.SettleTime
}

func (s *MealOrderDetailQueryResponseBodyModule) GetThirdPartApplyId() *string {
	return s.ThirdPartApplyId
}

func (s *MealOrderDetailQueryResponseBodyModule) GetUserAlipayId() *string {
	return s.UserAlipayId
}

func (s *MealOrderDetailQueryResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *MealOrderDetailQueryResponseBodyModule) SetApplyId(v int64) *MealOrderDetailQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetCorpCodeOrderId(v string) *MealOrderDetailQueryResponseBodyModule {
	s.CorpCodeOrderId = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetCorpId(v string) *MealOrderDetailQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetCorpPayAmount(v int64) *MealOrderDetailQueryResponseBodyModule {
	s.CorpPayAmount = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetCorpRefundAmount(v int64) *MealOrderDetailQueryResponseBodyModule {
	s.CorpRefundAmount = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetMealReason(v string) *MealOrderDetailQueryResponseBodyModule {
	s.MealReason = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetMerchantName(v string) *MealOrderDetailQueryResponseBodyModule {
	s.MerchantName = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetOrderId(v string) *MealOrderDetailQueryResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetOrderStatus(v int32) *MealOrderDetailQueryResponseBodyModule {
	s.OrderStatus = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetOrderSubStatus(v int32) *MealOrderDetailQueryResponseBodyModule {
	s.OrderSubStatus = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetOrderType(v string) *MealOrderDetailQueryResponseBodyModule {
	s.OrderType = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetPayAmount(v int64) *MealOrderDetailQueryResponseBodyModule {
	s.PayAmount = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetPayType(v int32) *MealOrderDetailQueryResponseBodyModule {
	s.PayType = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetPersonPayAmount(v int64) *MealOrderDetailQueryResponseBodyModule {
	s.PersonPayAmount = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetPersonRefundAmount(v int64) *MealOrderDetailQueryResponseBodyModule {
	s.PersonRefundAmount = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetRefundAmount(v int64) *MealOrderDetailQueryResponseBodyModule {
	s.RefundAmount = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetSceneName(v string) *MealOrderDetailQueryResponseBodyModule {
	s.SceneName = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetSettleTime(v string) *MealOrderDetailQueryResponseBodyModule {
	s.SettleTime = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetThirdPartApplyId(v string) *MealOrderDetailQueryResponseBodyModule {
	s.ThirdPartApplyId = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetUserAlipayId(v string) *MealOrderDetailQueryResponseBodyModule {
	s.UserAlipayId = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) SetUserId(v string) *MealOrderDetailQueryResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *MealOrderDetailQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
