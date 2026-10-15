// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSubscriptionPriceResponseBody
	GetCode() *string
	SetData(v *GetSubscriptionPriceResponseBodyData) *GetSubscriptionPriceResponseBody
	GetData() *GetSubscriptionPriceResponseBodyData
	SetMessage(v string) *GetSubscriptionPriceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSubscriptionPriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSubscriptionPriceResponseBody
	GetSuccess() *bool
}

type GetSubscriptionPriceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// InvalidConfigCod
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The product price information.
	Data *GetSubscriptionPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30E7066E-AE6F-4E59-AFE6-11386CE3AFA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSubscriptionPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionPriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSubscriptionPriceResponseBody) GetData() *GetSubscriptionPriceResponseBodyData {
	return s.Data
}

func (s *GetSubscriptionPriceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSubscriptionPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSubscriptionPriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSubscriptionPriceResponseBody) SetCode(v string) *GetSubscriptionPriceResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubscriptionPriceResponseBody) SetData(v *GetSubscriptionPriceResponseBodyData) *GetSubscriptionPriceResponseBody {
	s.Data = v
	return s
}

func (s *GetSubscriptionPriceResponseBody) SetMessage(v string) *GetSubscriptionPriceResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubscriptionPriceResponseBody) SetRequestId(v string) *GetSubscriptionPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubscriptionPriceResponseBody) SetSuccess(v bool) *GetSubscriptionPriceResponseBody {
	s.Success = &v
	return s
}

func (s *GetSubscriptionPriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSubscriptionPriceResponseBodyData struct {
	// The currency. Valid values:
	//
	// - CNY: Chinese yuan.
	//
	// - USD: US dollar.
	//
	// - JPY: Japanese yen.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discount amount.
	//
	// example:
	//
	// 100
	DiscountPrice *float32                                           `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	ModuleDetails *GetSubscriptionPriceResponseBodyDataModuleDetails `json:"ModuleDetails,omitempty" xml:"ModuleDetails,omitempty" type:"Struct"`
	// The original price.
	//
	// example:
	//
	// 900
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The parameters for fixed-fee installments.
	PaymentScheme    *GetSubscriptionPriceResponseBodyDataPaymentScheme    `json:"PaymentScheme,omitempty" xml:"PaymentScheme,omitempty" type:"Struct"`
	PromotionDetails *GetSubscriptionPriceResponseBodyDataPromotionDetails `json:"PromotionDetails,omitempty" xml:"PromotionDetails,omitempty" type:"Struct"`
	// The quantity.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The discounted price.
	//
	// example:
	//
	// 0
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s GetSubscriptionPriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *GetSubscriptionPriceResponseBodyData) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *GetSubscriptionPriceResponseBodyData) GetModuleDetails() *GetSubscriptionPriceResponseBodyDataModuleDetails {
	return s.ModuleDetails
}

func (s *GetSubscriptionPriceResponseBodyData) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *GetSubscriptionPriceResponseBodyData) GetPaymentScheme() *GetSubscriptionPriceResponseBodyDataPaymentScheme {
	return s.PaymentScheme
}

func (s *GetSubscriptionPriceResponseBodyData) GetPromotionDetails() *GetSubscriptionPriceResponseBodyDataPromotionDetails {
	return s.PromotionDetails
}

func (s *GetSubscriptionPriceResponseBodyData) GetQuantity() *int32 {
	return s.Quantity
}

func (s *GetSubscriptionPriceResponseBodyData) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *GetSubscriptionPriceResponseBodyData) SetCurrency(v string) *GetSubscriptionPriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyData) SetDiscountPrice(v float32) *GetSubscriptionPriceResponseBodyData {
	s.DiscountPrice = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyData) SetModuleDetails(v *GetSubscriptionPriceResponseBodyDataModuleDetails) *GetSubscriptionPriceResponseBodyData {
	s.ModuleDetails = v
	return s
}

func (s *GetSubscriptionPriceResponseBodyData) SetOriginalPrice(v float32) *GetSubscriptionPriceResponseBodyData {
	s.OriginalPrice = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyData) SetPaymentScheme(v *GetSubscriptionPriceResponseBodyDataPaymentScheme) *GetSubscriptionPriceResponseBodyData {
	s.PaymentScheme = v
	return s
}

func (s *GetSubscriptionPriceResponseBodyData) SetPromotionDetails(v *GetSubscriptionPriceResponseBodyDataPromotionDetails) *GetSubscriptionPriceResponseBodyData {
	s.PromotionDetails = v
	return s
}

func (s *GetSubscriptionPriceResponseBodyData) SetQuantity(v int32) *GetSubscriptionPriceResponseBodyData {
	s.Quantity = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyData) SetTradePrice(v float32) *GetSubscriptionPriceResponseBodyData {
	s.TradePrice = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyData) Validate() error {
	if s.ModuleDetails != nil {
		if err := s.ModuleDetails.Validate(); err != nil {
			return err
		}
	}
	if s.PaymentScheme != nil {
		if err := s.PaymentScheme.Validate(); err != nil {
			return err
		}
	}
	if s.PromotionDetails != nil {
		if err := s.PromotionDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSubscriptionPriceResponseBodyDataModuleDetails struct {
	ModuleDetail []*GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail `json:"ModuleDetail,omitempty" xml:"ModuleDetail,omitempty" type:"Repeated"`
}

func (s GetSubscriptionPriceResponseBodyDataModuleDetails) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionPriceResponseBodyDataModuleDetails) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetails) GetModuleDetail() []*GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail {
	return s.ModuleDetail
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetails) SetModuleDetail(v []*GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) *GetSubscriptionPriceResponseBodyDataModuleDetails {
	s.ModuleDetail = v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetails) Validate() error {
	if s.ModuleDetail != nil {
		for _, item := range s.ModuleDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail struct {
	CostAfterDiscount *float32 `json:"CostAfterDiscount,omitempty" xml:"CostAfterDiscount,omitempty"`
	InvoiceDiscount   *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	ModuleCode        *string  `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	OriginalCost      *float32 `json:"OriginalCost,omitempty" xml:"OriginalCost,omitempty"`
	UnitPrice         *float32 `json:"UnitPrice,omitempty" xml:"UnitPrice,omitempty"`
}

func (s GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) GetCostAfterDiscount() *float32 {
	return s.CostAfterDiscount
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) GetInvoiceDiscount() *float32 {
	return s.InvoiceDiscount
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) GetOriginalCost() *float32 {
	return s.OriginalCost
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) GetUnitPrice() *float32 {
	return s.UnitPrice
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) SetCostAfterDiscount(v float32) *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail {
	s.CostAfterDiscount = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) SetInvoiceDiscount(v float32) *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail {
	s.InvoiceDiscount = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) SetModuleCode(v string) *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail {
	s.ModuleCode = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) SetOriginalCost(v float32) *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail {
	s.OriginalCost = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) SetUnitPrice(v float32) *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail {
	s.UnitPrice = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) Validate() error {
	return dara.Validate(s)
}

type GetSubscriptionPriceResponseBodyDataPaymentScheme struct {
	// The payable amount for the order, in cents.
	//
	// example:
	//
	// 35092
	OrderDownPaymentFee *string `json:"OrderDownPaymentFee,omitempty" xml:"OrderDownPaymentFee,omitempty"`
	// The payable amount in the settlement currency for the order, in cents.
	//
	// example:
	//
	// 35092
	OrderDownPaymentSettleFee *string `json:"OrderDownPaymentSettleFee,omitempty" xml:"OrderDownPaymentSettleFee,omitempty"`
	// The total payable amount for the order, including the prepayment and installment parts, in cents.
	//
	// example:
	//
	// 70184
	OrderTotalFee *string `json:"OrderTotalFee,omitempty" xml:"OrderTotalFee,omitempty"`
	// The total payable amount in the settlement currency for the order, including the prepayment and installment parts, in cents.
	//
	// example:
	//
	// 70184
	OrderTotalSettleFee *string `json:"OrderTotalSettleFee,omitempty" xml:"OrderTotalSettleFee,omitempty"`
	// The fixed-fee installment type. Valid values:
	//
	// - PAYMENT_NORMAL: non-fixed-fee installment products.
	//
	// - PAYMENT_PHASED: fixed-fee installment products.
	//
	// example:
	//
	// PAYMENT_NORMAL
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
}

func (s GetSubscriptionPriceResponseBodyDataPaymentScheme) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionPriceResponseBodyDataPaymentScheme) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceResponseBodyDataPaymentScheme) GetOrderDownPaymentFee() *string {
	return s.OrderDownPaymentFee
}

func (s *GetSubscriptionPriceResponseBodyDataPaymentScheme) GetOrderDownPaymentSettleFee() *string {
	return s.OrderDownPaymentSettleFee
}

func (s *GetSubscriptionPriceResponseBodyDataPaymentScheme) GetOrderTotalFee() *string {
	return s.OrderTotalFee
}

func (s *GetSubscriptionPriceResponseBodyDataPaymentScheme) GetOrderTotalSettleFee() *string {
	return s.OrderTotalSettleFee
}

func (s *GetSubscriptionPriceResponseBodyDataPaymentScheme) GetPaymentType() *string {
	return s.PaymentType
}

func (s *GetSubscriptionPriceResponseBodyDataPaymentScheme) SetOrderDownPaymentFee(v string) *GetSubscriptionPriceResponseBodyDataPaymentScheme {
	s.OrderDownPaymentFee = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataPaymentScheme) SetOrderDownPaymentSettleFee(v string) *GetSubscriptionPriceResponseBodyDataPaymentScheme {
	s.OrderDownPaymentSettleFee = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataPaymentScheme) SetOrderTotalFee(v string) *GetSubscriptionPriceResponseBodyDataPaymentScheme {
	s.OrderTotalFee = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataPaymentScheme) SetOrderTotalSettleFee(v string) *GetSubscriptionPriceResponseBodyDataPaymentScheme {
	s.OrderTotalSettleFee = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataPaymentScheme) SetPaymentType(v string) *GetSubscriptionPriceResponseBodyDataPaymentScheme {
	s.PaymentType = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataPaymentScheme) Validate() error {
	return dara.Validate(s)
}

type GetSubscriptionPriceResponseBodyDataPromotionDetails struct {
	PromotionDetail []*GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail `json:"PromotionDetail,omitempty" xml:"PromotionDetail,omitempty" type:"Repeated"`
}

func (s GetSubscriptionPriceResponseBodyDataPromotionDetails) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionPriceResponseBodyDataPromotionDetails) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceResponseBodyDataPromotionDetails) GetPromotionDetail() []*GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail {
	return s.PromotionDetail
}

func (s *GetSubscriptionPriceResponseBodyDataPromotionDetails) SetPromotionDetail(v []*GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) *GetSubscriptionPriceResponseBodyDataPromotionDetails {
	s.PromotionDetail = v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataPromotionDetails) Validate() error {
	if s.PromotionDetail != nil {
		for _, item := range s.PromotionDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail struct {
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionId   *int64  `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
}

func (s GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) GetPromotionId() *int64 {
	return s.PromotionId
}

func (s *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) GetPromotionName() *string {
	return s.PromotionName
}

func (s *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) SetPromotionDesc(v string) *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail {
	s.PromotionDesc = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) SetPromotionId(v int64) *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail {
	s.PromotionId = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) SetPromotionName(v string) *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail {
	s.PromotionName = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) Validate() error {
	return dara.Validate(s)
}
