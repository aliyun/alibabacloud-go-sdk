// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketPageQueryProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TicketPageQueryProductResponseBodyData) *TicketPageQueryProductResponseBody
	GetData() *TicketPageQueryProductResponseBodyData
	SetErrorCode(v string) *TicketPageQueryProductResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *TicketPageQueryProductResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *TicketPageQueryProductResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketPageQueryProductResponseBody
	GetSuccess() *bool
}

type TicketPageQueryProductResponseBody struct {
	Data *TicketPageQueryProductResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ScenicIdInvalid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// ScenicId不合法
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TicketPageQueryProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBody) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBody) GetData() *TicketPageQueryProductResponseBodyData {
	return s.Data
}

func (s *TicketPageQueryProductResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TicketPageQueryProductResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TicketPageQueryProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketPageQueryProductResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketPageQueryProductResponseBody) SetData(v *TicketPageQueryProductResponseBodyData) *TicketPageQueryProductResponseBody {
	s.Data = v
	return s
}

func (s *TicketPageQueryProductResponseBody) SetErrorCode(v string) *TicketPageQueryProductResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketPageQueryProductResponseBody) SetErrorMsg(v string) *TicketPageQueryProductResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketPageQueryProductResponseBody) SetRequestId(v string) *TicketPageQueryProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketPageQueryProductResponseBody) SetSuccess(v bool) *TicketPageQueryProductResponseBody {
	s.Success = &v
	return s
}

func (s *TicketPageQueryProductResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketPageQueryProductResponseBodyData struct {
	Products []*TicketPageQueryProductResponseBodyDataProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s TicketPageQueryProductResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyData) GetProducts() []*TicketPageQueryProductResponseBodyDataProducts {
	return s.Products
}

func (s *TicketPageQueryProductResponseBodyData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *TicketPageQueryProductResponseBodyData) SetProducts(v []*TicketPageQueryProductResponseBodyDataProducts) *TicketPageQueryProductResponseBodyData {
	s.Products = v
	return s
}

func (s *TicketPageQueryProductResponseBodyData) SetTotalSize(v int64) *TicketPageQueryProductResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyData) Validate() error {
	if s.Products != nil {
		for _, item := range s.Products {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TicketPageQueryProductResponseBodyDataProducts struct {
	// example:
	//
	// 2
	BookingType *int32                                                 `json:"BookingType,omitempty" xml:"BookingType,omitempty"`
	BuyRule     *TicketPageQueryProductResponseBodyDataProductsBuyRule `json:"BuyRule,omitempty" xml:"BuyRule,omitempty" type:"Struct"`
	// example:
	//
	// 含景区大门票一张
	CostIncludeRemark *string `json:"CostIncludeRemark,omitempty" xml:"CostIncludeRemark,omitempty"`
	// example:
	//
	// 120
	DeliverGuaranteeMinutes *int32 `json:"DeliverGuaranteeMinutes,omitempty" xml:"DeliverGuaranteeMinutes,omitempty"`
	// example:
	//
	// ["https://example.com/detail1.jpg"]
	Images []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	InvoiceIssuerType *int32 `json:"InvoiceIssuerType,omitempty" xml:"InvoiceIssuerType,omitempty"`
	// example:
	//
	// 20
	PaymentLimitMinutes *int32 `json:"PaymentLimitMinutes,omitempty" xml:"PaymentLimitMinutes,omitempty"`
	// example:
	//
	// 123456
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 西湖游船成人票
	ProductName *string                                                   `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	RefundRule  *TicketPageQueryProductResponseBodyDataProductsRefundRule `json:"RefundRule,omitempty" xml:"RefundRule,omitempty" type:"Struct"`
	Region      *TicketPageQueryProductResponseBodyDataProductsRegion     `json:"Region,omitempty" xml:"Region,omitempty" type:"Struct"`
	// example:
	//
	// 123456
	ScenicId *int64                                                 `json:"ScenicId,omitempty" xml:"ScenicId,omitempty"`
	Session  *TicketPageQueryProductResponseBodyDataProductsSession `json:"Session,omitempty" xml:"Session,omitempty" type:"Struct"`
	// example:
	//
	// 1
	SettlePriceCalculateType *int32                                             `json:"SettlePriceCalculateType,omitempty" xml:"SettlePriceCalculateType,omitempty"`
	Spu                      *TicketPageQueryProductResponseBodyDataProductsSpu `json:"Spu,omitempty" xml:"Spu,omitempty" type:"Struct"`
	// example:
	//
	// 飞猪景区乐园旗舰店
	SupplierName *string                                                   `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	TicketKind   *TicketPageQueryProductResponseBodyDataProductsTicketKind `json:"TicketKind,omitempty" xml:"TicketKind,omitempty" type:"Struct"`
	UseRule      *TicketPageQueryProductResponseBodyDataProductsUseRule    `json:"UseRule,omitempty" xml:"UseRule,omitempty" type:"Struct"`
}

func (s TicketPageQueryProductResponseBodyDataProducts) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProducts) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetBookingType() *int32 {
	return s.BookingType
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetBuyRule() *TicketPageQueryProductResponseBodyDataProductsBuyRule {
	return s.BuyRule
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetCostIncludeRemark() *string {
	return s.CostIncludeRemark
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetDeliverGuaranteeMinutes() *int32 {
	return s.DeliverGuaranteeMinutes
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetImages() []*string {
	return s.Images
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetInvoiceIssuerType() *int32 {
	return s.InvoiceIssuerType
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetPaymentLimitMinutes() *int32 {
	return s.PaymentLimitMinutes
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetProductId() *string {
	return s.ProductId
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetProductName() *string {
	return s.ProductName
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetRefundRule() *TicketPageQueryProductResponseBodyDataProductsRefundRule {
	return s.RefundRule
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetRegion() *TicketPageQueryProductResponseBodyDataProductsRegion {
	return s.Region
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetScenicId() *int64 {
	return s.ScenicId
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetSession() *TicketPageQueryProductResponseBodyDataProductsSession {
	return s.Session
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetSettlePriceCalculateType() *int32 {
	return s.SettlePriceCalculateType
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetSpu() *TicketPageQueryProductResponseBodyDataProductsSpu {
	return s.Spu
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetSupplierName() *string {
	return s.SupplierName
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetTicketKind() *TicketPageQueryProductResponseBodyDataProductsTicketKind {
	return s.TicketKind
}

func (s *TicketPageQueryProductResponseBodyDataProducts) GetUseRule() *TicketPageQueryProductResponseBodyDataProductsUseRule {
	return s.UseRule
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetBookingType(v int32) *TicketPageQueryProductResponseBodyDataProducts {
	s.BookingType = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetBuyRule(v *TicketPageQueryProductResponseBodyDataProductsBuyRule) *TicketPageQueryProductResponseBodyDataProducts {
	s.BuyRule = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetCostIncludeRemark(v string) *TicketPageQueryProductResponseBodyDataProducts {
	s.CostIncludeRemark = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetDeliverGuaranteeMinutes(v int32) *TicketPageQueryProductResponseBodyDataProducts {
	s.DeliverGuaranteeMinutes = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetImages(v []*string) *TicketPageQueryProductResponseBodyDataProducts {
	s.Images = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetInvoiceIssuerType(v int32) *TicketPageQueryProductResponseBodyDataProducts {
	s.InvoiceIssuerType = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetPaymentLimitMinutes(v int32) *TicketPageQueryProductResponseBodyDataProducts {
	s.PaymentLimitMinutes = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetProductId(v string) *TicketPageQueryProductResponseBodyDataProducts {
	s.ProductId = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetProductName(v string) *TicketPageQueryProductResponseBodyDataProducts {
	s.ProductName = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetRefundRule(v *TicketPageQueryProductResponseBodyDataProductsRefundRule) *TicketPageQueryProductResponseBodyDataProducts {
	s.RefundRule = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetRegion(v *TicketPageQueryProductResponseBodyDataProductsRegion) *TicketPageQueryProductResponseBodyDataProducts {
	s.Region = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetScenicId(v int64) *TicketPageQueryProductResponseBodyDataProducts {
	s.ScenicId = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetSession(v *TicketPageQueryProductResponseBodyDataProductsSession) *TicketPageQueryProductResponseBodyDataProducts {
	s.Session = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetSettlePriceCalculateType(v int32) *TicketPageQueryProductResponseBodyDataProducts {
	s.SettlePriceCalculateType = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetSpu(v *TicketPageQueryProductResponseBodyDataProductsSpu) *TicketPageQueryProductResponseBodyDataProducts {
	s.Spu = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetSupplierName(v string) *TicketPageQueryProductResponseBodyDataProducts {
	s.SupplierName = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetTicketKind(v *TicketPageQueryProductResponseBodyDataProductsTicketKind) *TicketPageQueryProductResponseBodyDataProducts {
	s.TicketKind = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) SetUseRule(v *TicketPageQueryProductResponseBodyDataProductsUseRule) *TicketPageQueryProductResponseBodyDataProducts {
	s.UseRule = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProducts) Validate() error {
	if s.BuyRule != nil {
		if err := s.BuyRule.Validate(); err != nil {
			return err
		}
	}
	if s.RefundRule != nil {
		if err := s.RefundRule.Validate(); err != nil {
			return err
		}
	}
	if s.Region != nil {
		if err := s.Region.Validate(); err != nil {
			return err
		}
	}
	if s.Session != nil {
		if err := s.Session.Validate(); err != nil {
			return err
		}
	}
	if s.Spu != nil {
		if err := s.Spu.Validate(); err != nil {
			return err
		}
	}
	if s.TicketKind != nil {
		if err := s.TicketKind.Validate(); err != nil {
			return err
		}
	}
	if s.UseRule != nil {
		if err := s.UseRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketPageQueryProductResponseBodyDataProductsBuyRule struct {
	AheadBuyTimePointRule           *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule             `json:"AheadBuyTimePointRule,omitempty" xml:"AheadBuyTimePointRule,omitempty" type:"Struct"`
	ContactRule                     *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRule                       `json:"ContactRule,omitempty" xml:"ContactRule,omitempty" type:"Struct"`
	CrossOrderBuyQuantityLimitRules []*TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules `json:"CrossOrderBuyQuantityLimitRules,omitempty" xml:"CrossOrderBuyQuantityLimitRules,omitempty" type:"Repeated"`
	PerOrderBuyQuantityLimitRule    *TicketPageQueryProductResponseBodyDataProductsBuyRulePerOrderBuyQuantityLimitRule      `json:"PerOrderBuyQuantityLimitRule,omitempty" xml:"PerOrderBuyQuantityLimitRule,omitempty" type:"Struct"`
	TravelerRule                    *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule                      `json:"TravelerRule,omitempty" xml:"TravelerRule,omitempty" type:"Struct"`
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRule) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRule) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRule) GetAheadBuyTimePointRule() *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule {
	return s.AheadBuyTimePointRule
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRule) GetContactRule() *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRule {
	return s.ContactRule
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRule) GetCrossOrderBuyQuantityLimitRules() []*TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules {
	return s.CrossOrderBuyQuantityLimitRules
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRule) GetPerOrderBuyQuantityLimitRule() *TicketPageQueryProductResponseBodyDataProductsBuyRulePerOrderBuyQuantityLimitRule {
	return s.PerOrderBuyQuantityLimitRule
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRule) GetTravelerRule() *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule {
	return s.TravelerRule
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRule) SetAheadBuyTimePointRule(v *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule) *TicketPageQueryProductResponseBodyDataProductsBuyRule {
	s.AheadBuyTimePointRule = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRule) SetContactRule(v *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRule) *TicketPageQueryProductResponseBodyDataProductsBuyRule {
	s.ContactRule = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRule) SetCrossOrderBuyQuantityLimitRules(v []*TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) *TicketPageQueryProductResponseBodyDataProductsBuyRule {
	s.CrossOrderBuyQuantityLimitRules = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRule) SetPerOrderBuyQuantityLimitRule(v *TicketPageQueryProductResponseBodyDataProductsBuyRulePerOrderBuyQuantityLimitRule) *TicketPageQueryProductResponseBodyDataProductsBuyRule {
	s.PerOrderBuyQuantityLimitRule = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRule) SetTravelerRule(v *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) *TicketPageQueryProductResponseBodyDataProductsBuyRule {
	s.TravelerRule = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRule) Validate() error {
	if s.AheadBuyTimePointRule != nil {
		if err := s.AheadBuyTimePointRule.Validate(); err != nil {
			return err
		}
	}
	if s.ContactRule != nil {
		if err := s.ContactRule.Validate(); err != nil {
			return err
		}
	}
	if s.CrossOrderBuyQuantityLimitRules != nil {
		for _, item := range s.CrossOrderBuyQuantityLimitRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PerOrderBuyQuantityLimitRule != nil {
		if err := s.PerOrderBuyQuantityLimitRule.Validate(); err != nil {
			return err
		}
	}
	if s.TravelerRule != nil {
		if err := s.TravelerRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule struct {
	// example:
	//
	// 1
	Anchor *int32 `json:"Anchor,omitempty" xml:"Anchor,omitempty"`
	// example:
	//
	// 2026-01-01
	FixedTime *string `json:"FixedTime,omitempty" xml:"FixedTime,omitempty"`
	// example:
	//
	// 18:00
	OffsetDayOfTime *string `json:"OffsetDayOfTime,omitempty" xml:"OffsetDayOfTime,omitempty"`
	// example:
	//
	// 1
	OffsetUnit *int32 `json:"OffsetUnit,omitempty" xml:"OffsetUnit,omitempty"`
	// example:
	//
	// 1
	OffsetValue *int32 `json:"OffsetValue,omitempty" xml:"OffsetValue,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule) GetAnchor() *int32 {
	return s.Anchor
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule) GetFixedTime() *string {
	return s.FixedTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule) GetOffsetDayOfTime() *string {
	return s.OffsetDayOfTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule) GetOffsetUnit() *int32 {
	return s.OffsetUnit
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule) GetOffsetValue() *int32 {
	return s.OffsetValue
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule) SetAnchor(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule {
	s.Anchor = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule) SetFixedTime(v string) *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule {
	s.FixedTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule) SetOffsetDayOfTime(v string) *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule {
	s.OffsetDayOfTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule) SetOffsetUnit(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule {
	s.OffsetUnit = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule) SetOffsetValue(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule {
	s.OffsetValue = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleAheadBuyTimePointRule) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRule struct {
	ContactFieldRule *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule `json:"ContactFieldRule,omitempty" xml:"ContactFieldRule,omitempty" type:"Struct"`
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRule) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRule) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRule) GetContactFieldRule() *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule {
	return s.ContactFieldRule
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRule) SetContactFieldRule(v *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRule {
	s.ContactFieldRule = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRule) Validate() error {
	if s.ContactFieldRule != nil {
		if err := s.ContactFieldRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule struct {
	// example:
	//
	// true
	Certificate *bool `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// example:
	//
	// [1, 2]
	CertificateTypes []*int32 `json:"CertificateTypes,omitempty" xml:"CertificateTypes,omitempty" type:"Repeated"`
	// example:
	//
	// true
	DialingCode *bool `json:"DialingCode,omitempty" xml:"DialingCode,omitempty"`
	// example:
	//
	// false
	Email *bool `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// false
	FirstName *bool `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// example:
	//
	// false
	LastName *bool `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// example:
	//
	// true
	Mobile *bool `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// true
	Name *bool `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) GetCertificate() *bool {
	return s.Certificate
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) GetCertificateTypes() []*int32 {
	return s.CertificateTypes
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) GetDialingCode() *bool {
	return s.DialingCode
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) GetEmail() *bool {
	return s.Email
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) GetFirstName() *bool {
	return s.FirstName
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) GetLastName() *bool {
	return s.LastName
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) GetMobile() *bool {
	return s.Mobile
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) GetName() *bool {
	return s.Name
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) SetCertificate(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule {
	s.Certificate = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) SetCertificateTypes(v []*int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule {
	s.CertificateTypes = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) SetDialingCode(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule {
	s.DialingCode = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) SetEmail(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule {
	s.Email = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) SetFirstName(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule {
	s.FirstName = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) SetLastName(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule {
	s.LastName = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) SetMobile(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule {
	s.Mobile = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) SetName(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule {
	s.Name = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleContactRuleContactFieldRule) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules struct {
	// example:
	//
	// 1
	LimitDayType *int32 `json:"LimitDayType,omitempty" xml:"LimitDayType,omitempty"`
	// example:
	//
	// 7
	LimitDays *int32 `json:"LimitDays,omitempty" xml:"LimitDays,omitempty"`
	// example:
	//
	// 1
	LimitPeriod *int32 `json:"LimitPeriod,omitempty" xml:"LimitPeriod,omitempty"`
	// example:
	//
	// 1
	LimitQuantityType *int32 `json:"LimitQuantityType,omitempty" xml:"LimitQuantityType,omitempty"`
	// example:
	//
	// 1
	LimitType *int32 `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	// example:
	//
	// 5
	MaxBuyQuantity *int32 `json:"MaxBuyQuantity,omitempty" xml:"MaxBuyQuantity,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) GetLimitDayType() *int32 {
	return s.LimitDayType
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) GetLimitDays() *int32 {
	return s.LimitDays
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) GetLimitPeriod() *int32 {
	return s.LimitPeriod
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) GetLimitQuantityType() *int32 {
	return s.LimitQuantityType
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) GetLimitType() *int32 {
	return s.LimitType
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) GetMaxBuyQuantity() *int32 {
	return s.MaxBuyQuantity
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) SetLimitDayType(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules {
	s.LimitDayType = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) SetLimitDays(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules {
	s.LimitDays = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) SetLimitPeriod(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules {
	s.LimitPeriod = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) SetLimitQuantityType(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules {
	s.LimitQuantityType = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) SetLimitType(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules {
	s.LimitType = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) SetMaxBuyQuantity(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules {
	s.MaxBuyQuantity = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleCrossOrderBuyQuantityLimitRules) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsBuyRulePerOrderBuyQuantityLimitRule struct {
	// example:
	//
	// 10
	MaxBuyQuantity *int32 `json:"MaxBuyQuantity,omitempty" xml:"MaxBuyQuantity,omitempty"`
	// example:
	//
	// 1
	MinBuyQuantity *int32 `json:"MinBuyQuantity,omitempty" xml:"MinBuyQuantity,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRulePerOrderBuyQuantityLimitRule) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRulePerOrderBuyQuantityLimitRule) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRulePerOrderBuyQuantityLimitRule) GetMaxBuyQuantity() *int32 {
	return s.MaxBuyQuantity
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRulePerOrderBuyQuantityLimitRule) GetMinBuyQuantity() *int32 {
	return s.MinBuyQuantity
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRulePerOrderBuyQuantityLimitRule) SetMaxBuyQuantity(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRulePerOrderBuyQuantityLimitRule {
	s.MaxBuyQuantity = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRulePerOrderBuyQuantityLimitRule) SetMinBuyQuantity(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRulePerOrderBuyQuantityLimitRule {
	s.MinBuyQuantity = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRulePerOrderBuyQuantityLimitRule) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule struct {
	CrowdLimitRules     []*TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules     `json:"CrowdLimitRules,omitempty" xml:"CrowdLimitRules,omitempty" type:"Repeated"`
	CrowdQuantityLimits []*TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdQuantityLimits `json:"CrowdQuantityLimits,omitempty" xml:"CrowdQuantityLimits,omitempty" type:"Repeated"`
	// example:
	//
	// true
	NeedFillTraveler  *bool                                                                               `json:"NeedFillTraveler,omitempty" xml:"NeedFillTraveler,omitempty"`
	TravelerFieldRule *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule `json:"TravelerFieldRule,omitempty" xml:"TravelerFieldRule,omitempty" type:"Struct"`
	// example:
	//
	// 2
	TravelerFillDimension *int32 `json:"TravelerFillDimension,omitempty" xml:"TravelerFillDimension,omitempty"`
	// example:
	//
	// 3
	TravelerQuantity *int32 `json:"TravelerQuantity,omitempty" xml:"TravelerQuantity,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) GetCrowdLimitRules() []*TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules {
	return s.CrowdLimitRules
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) GetCrowdQuantityLimits() []*TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdQuantityLimits {
	return s.CrowdQuantityLimits
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) GetNeedFillTraveler() *bool {
	return s.NeedFillTraveler
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) GetTravelerFieldRule() *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule {
	return s.TravelerFieldRule
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) GetTravelerFillDimension() *int32 {
	return s.TravelerFillDimension
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) GetTravelerQuantity() *int32 {
	return s.TravelerQuantity
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) SetCrowdLimitRules(v []*TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule {
	s.CrowdLimitRules = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) SetCrowdQuantityLimits(v []*TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdQuantityLimits) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule {
	s.CrowdQuantityLimits = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) SetNeedFillTraveler(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule {
	s.NeedFillTraveler = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) SetTravelerFieldRule(v *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule {
	s.TravelerFieldRule = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) SetTravelerFillDimension(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule {
	s.TravelerFillDimension = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) SetTravelerQuantity(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule {
	s.TravelerQuantity = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRule) Validate() error {
	if s.CrowdLimitRules != nil {
		for _, item := range s.CrowdLimitRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CrowdQuantityLimits != nil {
		for _, item := range s.CrowdQuantityLimits {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TravelerFieldRule != nil {
		if err := s.TravelerFieldRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules struct {
	// example:
	//
	// 2
	AgeBaseTimeType *int32 `json:"AgeBaseTimeType,omitempty" xml:"AgeBaseTimeType,omitempty"`
	// example:
	//
	// 1
	AgeCalculateType *int32 `json:"AgeCalculateType,omitempty" xml:"AgeCalculateType,omitempty"`
	// example:
	//
	// 65
	AgeMax *int32 `json:"AgeMax,omitempty" xml:"AgeMax,omitempty"`
	// example:
	//
	// 18
	AgeMin *int32 `json:"AgeMin,omitempty" xml:"AgeMin,omitempty"`
	// example:
	//
	// 成人票
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules) GetAgeBaseTimeType() *int32 {
	return s.AgeBaseTimeType
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules) GetAgeCalculateType() *int32 {
	return s.AgeCalculateType
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules) GetAgeMax() *int32 {
	return s.AgeMax
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules) GetAgeMin() *int32 {
	return s.AgeMin
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules) GetName() *string {
	return s.Name
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules) SetAgeBaseTimeType(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules {
	s.AgeBaseTimeType = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules) SetAgeCalculateType(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules {
	s.AgeCalculateType = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules) SetAgeMax(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules {
	s.AgeMax = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules) SetAgeMin(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules {
	s.AgeMin = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules) SetName(v string) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules {
	s.Name = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdLimitRules) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdQuantityLimits struct {
	// example:
	//
	// 成人票
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdQuantityLimits) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdQuantityLimits) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdQuantityLimits) GetName() *string {
	return s.Name
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdQuantityLimits) GetQuantity() *int32 {
	return s.Quantity
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdQuantityLimits) SetName(v string) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdQuantityLimits {
	s.Name = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdQuantityLimits) SetQuantity(v int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdQuantityLimits {
	s.Quantity = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleCrowdQuantityLimits) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule struct {
	// example:
	//
	// false
	Birthday *bool `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	// example:
	//
	// true
	Certificate *bool `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// example:
	//
	// [1, 2]
	CertificateTypes []*int32 `json:"CertificateTypes,omitempty" xml:"CertificateTypes,omitempty" type:"Repeated"`
	// example:
	//
	// true
	DialingCode *bool `json:"DialingCode,omitempty" xml:"DialingCode,omitempty"`
	// example:
	//
	// false
	Email *bool `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// false
	FirstName *bool `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// example:
	//
	// false
	Gender *bool `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// example:
	//
	// false
	LastName *bool `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// example:
	//
	// true
	Mobile *bool `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// true
	Name *bool `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	Nationality *bool `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) GetBirthday() *bool {
	return s.Birthday
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) GetCertificate() *bool {
	return s.Certificate
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) GetCertificateTypes() []*int32 {
	return s.CertificateTypes
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) GetDialingCode() *bool {
	return s.DialingCode
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) GetEmail() *bool {
	return s.Email
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) GetFirstName() *bool {
	return s.FirstName
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) GetGender() *bool {
	return s.Gender
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) GetLastName() *bool {
	return s.LastName
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) GetMobile() *bool {
	return s.Mobile
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) GetName() *bool {
	return s.Name
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) GetNationality() *bool {
	return s.Nationality
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) SetBirthday(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule {
	s.Birthday = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) SetCertificate(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule {
	s.Certificate = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) SetCertificateTypes(v []*int32) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule {
	s.CertificateTypes = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) SetDialingCode(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule {
	s.DialingCode = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) SetEmail(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule {
	s.Email = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) SetFirstName(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule {
	s.FirstName = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) SetGender(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule {
	s.Gender = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) SetLastName(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule {
	s.LastName = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) SetMobile(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule {
	s.Mobile = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) SetName(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule {
	s.Name = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) SetNationality(v bool) *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule {
	s.Nationality = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsBuyRuleTravelerRuleTravelerFieldRule) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsRefundRule struct {
	RefundStageRules []*TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules `json:"RefundStageRules,omitempty" xml:"RefundStageRules,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	RefundType *int32 `json:"RefundType,omitempty" xml:"RefundType,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsRefundRule) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsRefundRule) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRule) GetRefundStageRules() []*TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules {
	return s.RefundStageRules
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRule) GetRefundType() *int32 {
	return s.RefundType
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRule) SetRefundStageRules(v []*TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules) *TicketPageQueryProductResponseBodyDataProductsRefundRule {
	s.RefundStageRules = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRule) SetRefundType(v int32) *TicketPageQueryProductResponseBodyDataProductsRefundRule {
	s.RefundType = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRule) Validate() error {
	if s.RefundStageRules != nil {
		for _, item := range s.RefundStageRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules struct {
	// example:
	//
	// 0.2056
	Fee *float64 `json:"Fee,omitempty" xml:"Fee,omitempty"`
	// example:
	//
	// 1
	FeeBase *int32 `json:"FeeBase,omitempty" xml:"FeeBase,omitempty"`
	// example:
	//
	// 1
	FeeType *int32                                                                        `json:"FeeType,omitempty" xml:"FeeType,omitempty"`
	From    *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom `json:"From,omitempty" xml:"From,omitempty" type:"Struct"`
	To      *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo   `json:"To,omitempty" xml:"To,omitempty" type:"Struct"`
}

func (s TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules) GetFee() *float64 {
	return s.Fee
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules) GetFeeBase() *int32 {
	return s.FeeBase
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules) GetFeeType() *int32 {
	return s.FeeType
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules) GetFrom() *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom {
	return s.From
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules) GetTo() *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo {
	return s.To
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules) SetFee(v float64) *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules {
	s.Fee = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules) SetFeeBase(v int32) *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules {
	s.FeeBase = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules) SetFeeType(v int32) *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules {
	s.FeeType = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules) SetFrom(v *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom) *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules {
	s.From = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules) SetTo(v *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo) *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules {
	s.To = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRules) Validate() error {
	if s.From != nil {
		if err := s.From.Validate(); err != nil {
			return err
		}
	}
	if s.To != nil {
		if err := s.To.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom struct {
	// example:
	//
	// 1
	Anchor *int32 `json:"Anchor,omitempty" xml:"Anchor,omitempty"`
	// example:
	//
	// 2026-01-01
	FixedTime *string `json:"FixedTime,omitempty" xml:"FixedTime,omitempty"`
	// example:
	//
	// 18:00
	OffsetDayOfTime *string `json:"OffsetDayOfTime,omitempty" xml:"OffsetDayOfTime,omitempty"`
	// example:
	//
	// 1
	OffsetUnit *int32 `json:"OffsetUnit,omitempty" xml:"OffsetUnit,omitempty"`
	// example:
	//
	// 1
	OffsetValue *int32 `json:"OffsetValue,omitempty" xml:"OffsetValue,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom) GetAnchor() *int32 {
	return s.Anchor
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom) GetFixedTime() *string {
	return s.FixedTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom) GetOffsetDayOfTime() *string {
	return s.OffsetDayOfTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom) GetOffsetUnit() *int32 {
	return s.OffsetUnit
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom) GetOffsetValue() *int32 {
	return s.OffsetValue
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom) SetAnchor(v int32) *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom {
	s.Anchor = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom) SetFixedTime(v string) *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom {
	s.FixedTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom) SetOffsetDayOfTime(v string) *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom {
	s.OffsetDayOfTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom) SetOffsetUnit(v int32) *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom {
	s.OffsetUnit = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom) SetOffsetValue(v int32) *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom {
	s.OffsetValue = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesFrom) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo struct {
	// example:
	//
	// 1
	Anchor *int32 `json:"Anchor,omitempty" xml:"Anchor,omitempty"`
	// example:
	//
	// 2026-01-01
	FixedTime *string `json:"FixedTime,omitempty" xml:"FixedTime,omitempty"`
	// example:
	//
	// 18:00
	OffsetDayOfTime *string `json:"OffsetDayOfTime,omitempty" xml:"OffsetDayOfTime,omitempty"`
	// example:
	//
	// 1
	OffsetUnit *int32 `json:"OffsetUnit,omitempty" xml:"OffsetUnit,omitempty"`
	// example:
	//
	// 1
	OffsetValue *int32 `json:"OffsetValue,omitempty" xml:"OffsetValue,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo) GetAnchor() *int32 {
	return s.Anchor
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo) GetFixedTime() *string {
	return s.FixedTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo) GetOffsetDayOfTime() *string {
	return s.OffsetDayOfTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo) GetOffsetUnit() *int32 {
	return s.OffsetUnit
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo) GetOffsetValue() *int32 {
	return s.OffsetValue
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo) SetAnchor(v int32) *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo {
	s.Anchor = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo) SetFixedTime(v string) *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo {
	s.FixedTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo) SetOffsetDayOfTime(v string) *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo {
	s.OffsetDayOfTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo) SetOffsetUnit(v int32) *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo {
	s.OffsetUnit = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo) SetOffsetValue(v int32) *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo {
	s.OffsetValue = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRefundRuleRefundStageRulesTo) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsRegion struct {
	// example:
	//
	// 普通区
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsRegion) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsRegion) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsRegion) GetName() *string {
	return s.Name
}

func (s *TicketPageQueryProductResponseBodyDataProductsRegion) SetName(v string) *TicketPageQueryProductResponseBodyDataProductsRegion {
	s.Name = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsRegion) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsSession struct {
	// example:
	//
	// 12:00
	SessionEndTime *string `json:"SessionEndTime,omitempty" xml:"SessionEndTime,omitempty"`
	// example:
	//
	// 08:00-12:00
	SessionName *string `json:"SessionName,omitempty" xml:"SessionName,omitempty"`
	// example:
	//
	// 08:00
	SessionStartTime *string `json:"SessionStartTime,omitempty" xml:"SessionStartTime,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsSession) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsSession) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsSession) GetSessionEndTime() *string {
	return s.SessionEndTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsSession) GetSessionName() *string {
	return s.SessionName
}

func (s *TicketPageQueryProductResponseBodyDataProductsSession) GetSessionStartTime() *string {
	return s.SessionStartTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsSession) SetSessionEndTime(v string) *TicketPageQueryProductResponseBodyDataProductsSession {
	s.SessionEndTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsSession) SetSessionName(v string) *TicketPageQueryProductResponseBodyDataProductsSession {
	s.SessionName = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsSession) SetSessionStartTime(v string) *TicketPageQueryProductResponseBodyDataProductsSession {
	s.SessionStartTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsSession) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsSpu struct {
	// example:
	//
	// 门票
	PrimaryTypeName *string `json:"PrimaryTypeName,omitempty" xml:"PrimaryTypeName,omitempty"`
	// example:
	//
	// 门票级别日历详情
	ReserveDetail *string `json:"ReserveDetail,omitempty" xml:"ReserveDetail,omitempty"`
	// example:
	//
	// 门票级别日历说明
	ReserveTitle *string `json:"ReserveTitle,omitempty" xml:"ReserveTitle,omitempty"`
	// example:
	//
	// 首道门票
	SecondaryTypeName *string `json:"SecondaryTypeName,omitempty" xml:"SecondaryTypeName,omitempty"`
	// example:
	//
	// 3507
	SpuId *int64 `json:"SpuId,omitempty" xml:"SpuId,omitempty"`
	// example:
	//
	// 1日票
	SpuName *string `json:"SpuName,omitempty" xml:"SpuName,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsSpu) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsSpu) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsSpu) GetPrimaryTypeName() *string {
	return s.PrimaryTypeName
}

func (s *TicketPageQueryProductResponseBodyDataProductsSpu) GetReserveDetail() *string {
	return s.ReserveDetail
}

func (s *TicketPageQueryProductResponseBodyDataProductsSpu) GetReserveTitle() *string {
	return s.ReserveTitle
}

func (s *TicketPageQueryProductResponseBodyDataProductsSpu) GetSecondaryTypeName() *string {
	return s.SecondaryTypeName
}

func (s *TicketPageQueryProductResponseBodyDataProductsSpu) GetSpuId() *int64 {
	return s.SpuId
}

func (s *TicketPageQueryProductResponseBodyDataProductsSpu) GetSpuName() *string {
	return s.SpuName
}

func (s *TicketPageQueryProductResponseBodyDataProductsSpu) SetPrimaryTypeName(v string) *TicketPageQueryProductResponseBodyDataProductsSpu {
	s.PrimaryTypeName = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsSpu) SetReserveDetail(v string) *TicketPageQueryProductResponseBodyDataProductsSpu {
	s.ReserveDetail = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsSpu) SetReserveTitle(v string) *TicketPageQueryProductResponseBodyDataProductsSpu {
	s.ReserveTitle = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsSpu) SetSecondaryTypeName(v string) *TicketPageQueryProductResponseBodyDataProductsSpu {
	s.SecondaryTypeName = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsSpu) SetSpuId(v int64) *TicketPageQueryProductResponseBodyDataProductsSpu {
	s.SpuId = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsSpu) SetSpuName(v string) *TicketPageQueryProductResponseBodyDataProductsSpu {
	s.SpuName = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsSpu) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsTicketKind struct {
	// example:
	//
	// 适用于18周岁(含)至59周岁(含)的游客
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 60484007
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 成人票
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsTicketKind) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsTicketKind) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsTicketKind) GetDescription() *string {
	return s.Description
}

func (s *TicketPageQueryProductResponseBodyDataProductsTicketKind) GetId() *int64 {
	return s.Id
}

func (s *TicketPageQueryProductResponseBodyDataProductsTicketKind) GetName() *string {
	return s.Name
}

func (s *TicketPageQueryProductResponseBodyDataProductsTicketKind) SetDescription(v string) *TicketPageQueryProductResponseBodyDataProductsTicketKind {
	s.Description = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsTicketKind) SetId(v int64) *TicketPageQueryProductResponseBodyDataProductsTicketKind {
	s.Id = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsTicketKind) SetName(v string) *TicketPageQueryProductResponseBodyDataProductsTicketKind {
	s.Name = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsTicketKind) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsUseRule struct {
	EffectTimePointRule *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule `json:"EffectTimePointRule,omitempty" xml:"EffectTimePointRule,omitempty" type:"Struct"`
	// example:
	//
	// 景区正门入口
	EntryAddress *string `json:"EntryAddress,omitempty" xml:"EntryAddress,omitempty"`
	// example:
	//
	// 请于入园前2小时至香港迪士尼度假区售票处旁人工服务站进行取票预约，服务时间上午 10:00 至 晚上 8:30
	EntryRemark      *string                                                                  `json:"EntryRemark,omitempty" xml:"EntryRemark,omitempty"`
	EntryTimePeriods []*TicketPageQueryProductResponseBodyDataProductsUseRuleEntryTimePeriods `json:"EntryTimePeriods,omitempty" xml:"EntryTimePeriods,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	EntryType            *int32                                                                     `json:"EntryType,omitempty" xml:"EntryType,omitempty"`
	EntryWithVoucherRule *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryWithVoucherRule `json:"EntryWithVoucherRule,omitempty" xml:"EntryWithVoucherRule,omitempty" type:"Struct"`
	// example:
	//
	// true
	NeedAssemble *bool `json:"NeedAssemble,omitempty" xml:"NeedAssemble,omitempty"`
	// example:
	//
	// true
	NeedPrebook *bool `json:"NeedPrebook,omitempty" xml:"NeedPrebook,omitempty"`
	// example:
	//
	// 请携带有效身份证件
	OtherNote          *string                                                                  `json:"OtherNote,omitempty" xml:"OtherNote,omitempty"`
	PickupsRule        *TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule        `json:"PickupsRule,omitempty" xml:"PickupsRule,omitempty" type:"Struct"`
	ValidityPeriodRule *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule `json:"ValidityPeriodRule,omitempty" xml:"ValidityPeriodRule,omitempty" type:"Struct"`
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRule) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRule) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) GetEffectTimePointRule() *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule {
	return s.EffectTimePointRule
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) GetEntryAddress() *string {
	return s.EntryAddress
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) GetEntryRemark() *string {
	return s.EntryRemark
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) GetEntryTimePeriods() []*TicketPageQueryProductResponseBodyDataProductsUseRuleEntryTimePeriods {
	return s.EntryTimePeriods
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) GetEntryType() *int32 {
	return s.EntryType
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) GetEntryWithVoucherRule() *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryWithVoucherRule {
	return s.EntryWithVoucherRule
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) GetNeedAssemble() *bool {
	return s.NeedAssemble
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) GetNeedPrebook() *bool {
	return s.NeedPrebook
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) GetOtherNote() *string {
	return s.OtherNote
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) GetPickupsRule() *TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule {
	return s.PickupsRule
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) GetValidityPeriodRule() *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule {
	return s.ValidityPeriodRule
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) SetEffectTimePointRule(v *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule) *TicketPageQueryProductResponseBodyDataProductsUseRule {
	s.EffectTimePointRule = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) SetEntryAddress(v string) *TicketPageQueryProductResponseBodyDataProductsUseRule {
	s.EntryAddress = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) SetEntryRemark(v string) *TicketPageQueryProductResponseBodyDataProductsUseRule {
	s.EntryRemark = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) SetEntryTimePeriods(v []*TicketPageQueryProductResponseBodyDataProductsUseRuleEntryTimePeriods) *TicketPageQueryProductResponseBodyDataProductsUseRule {
	s.EntryTimePeriods = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) SetEntryType(v int32) *TicketPageQueryProductResponseBodyDataProductsUseRule {
	s.EntryType = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) SetEntryWithVoucherRule(v *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryWithVoucherRule) *TicketPageQueryProductResponseBodyDataProductsUseRule {
	s.EntryWithVoucherRule = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) SetNeedAssemble(v bool) *TicketPageQueryProductResponseBodyDataProductsUseRule {
	s.NeedAssemble = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) SetNeedPrebook(v bool) *TicketPageQueryProductResponseBodyDataProductsUseRule {
	s.NeedPrebook = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) SetOtherNote(v string) *TicketPageQueryProductResponseBodyDataProductsUseRule {
	s.OtherNote = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) SetPickupsRule(v *TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule) *TicketPageQueryProductResponseBodyDataProductsUseRule {
	s.PickupsRule = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) SetValidityPeriodRule(v *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule) *TicketPageQueryProductResponseBodyDataProductsUseRule {
	s.ValidityPeriodRule = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRule) Validate() error {
	if s.EffectTimePointRule != nil {
		if err := s.EffectTimePointRule.Validate(); err != nil {
			return err
		}
	}
	if s.EntryTimePeriods != nil {
		for _, item := range s.EntryTimePeriods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EntryWithVoucherRule != nil {
		if err := s.EntryWithVoucherRule.Validate(); err != nil {
			return err
		}
	}
	if s.PickupsRule != nil {
		if err := s.PickupsRule.Validate(); err != nil {
			return err
		}
	}
	if s.ValidityPeriodRule != nil {
		if err := s.ValidityPeriodRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule struct {
	// example:
	//
	// 1
	Anchor *int32 `json:"Anchor,omitempty" xml:"Anchor,omitempty"`
	// example:
	//
	// 2026-01-01
	FixedTime *string `json:"FixedTime,omitempty" xml:"FixedTime,omitempty"`
	// example:
	//
	// 18:00
	OffsetDayOfTime *string `json:"OffsetDayOfTime,omitempty" xml:"OffsetDayOfTime,omitempty"`
	// example:
	//
	// 1
	OffsetUnit *int32 `json:"OffsetUnit,omitempty" xml:"OffsetUnit,omitempty"`
	// example:
	//
	// 1
	OffsetValue *int32 `json:"OffsetValue,omitempty" xml:"OffsetValue,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule) GetAnchor() *int32 {
	return s.Anchor
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule) GetFixedTime() *string {
	return s.FixedTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule) GetOffsetDayOfTime() *string {
	return s.OffsetDayOfTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule) GetOffsetUnit() *int32 {
	return s.OffsetUnit
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule) GetOffsetValue() *int32 {
	return s.OffsetValue
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule) SetAnchor(v int32) *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule {
	s.Anchor = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule) SetFixedTime(v string) *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule {
	s.FixedTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule) SetOffsetDayOfTime(v string) *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule {
	s.OffsetDayOfTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule) SetOffsetUnit(v int32) *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule {
	s.OffsetUnit = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule) SetOffsetValue(v int32) *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule {
	s.OffsetValue = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEffectTimePointRule) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsUseRuleEntryTimePeriods struct {
	// example:
	//
	// 08:00
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 17:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRuleEntryTimePeriods) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRuleEntryTimePeriods) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryTimePeriods) GetBeginTime() *string {
	return s.BeginTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryTimePeriods) GetEndTime() *string {
	return s.EndTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryTimePeriods) SetBeginTime(v string) *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryTimePeriods {
	s.BeginTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryTimePeriods) SetEndTime(v string) *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryTimePeriods {
	s.EndTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryTimePeriods) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsUseRuleEntryWithVoucherRule struct {
	// example:
	//
	// 凭二维码入园
	VoucherRemark *string `json:"VoucherRemark,omitempty" xml:"VoucherRemark,omitempty"`
	// example:
	//
	// [1]
	VoucherTypes []*int32 `json:"VoucherTypes,omitempty" xml:"VoucherTypes,omitempty" type:"Repeated"`
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRuleEntryWithVoucherRule) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRuleEntryWithVoucherRule) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryWithVoucherRule) GetVoucherRemark() *string {
	return s.VoucherRemark
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryWithVoucherRule) GetVoucherTypes() []*int32 {
	return s.VoucherTypes
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryWithVoucherRule) SetVoucherRemark(v string) *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryWithVoucherRule {
	s.VoucherRemark = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryWithVoucherRule) SetVoucherTypes(v []*int32) *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryWithVoucherRule {
	s.VoucherTypes = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleEntryWithVoucherRule) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule struct {
	// example:
	//
	// 景区游客中心
	PickupsAddress *string `json:"PickupsAddress,omitempty" xml:"PickupsAddress,omitempty"`
	// example:
	//
	// 凭短信取票
	VoucherRemark *string `json:"VoucherRemark,omitempty" xml:"VoucherRemark,omitempty"`
	// example:
	//
	// [1]
	VoucherTypes []*int32 `json:"VoucherTypes,omitempty" xml:"VoucherTypes,omitempty" type:"Repeated"`
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule) GetPickupsAddress() *string {
	return s.PickupsAddress
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule) GetVoucherRemark() *string {
	return s.VoucherRemark
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule) GetVoucherTypes() []*int32 {
	return s.VoucherTypes
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule) SetPickupsAddress(v string) *TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule {
	s.PickupsAddress = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule) SetVoucherRemark(v string) *TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule {
	s.VoucherRemark = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule) SetVoucherTypes(v []*int32) *TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule {
	s.VoucherTypes = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRulePickupsRule) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule struct {
	// example:
	//
	// [1, 2, 3, 4, 5, 6, 7]
	AvailableWeeks []*int32                                                                     `json:"AvailableWeeks,omitempty" xml:"AvailableWeeks,omitempty" type:"Repeated"`
	From           *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom `json:"From,omitempty" xml:"From,omitempty" type:"Struct"`
	To             *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo   `json:"To,omitempty" xml:"To,omitempty" type:"Struct"`
	// example:
	//
	// ["2026-01-01"]
	UnavailableDates []*string `json:"UnavailableDates,omitempty" xml:"UnavailableDates,omitempty" type:"Repeated"`
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule) GetAvailableWeeks() []*int32 {
	return s.AvailableWeeks
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule) GetFrom() *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom {
	return s.From
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule) GetTo() *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo {
	return s.To
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule) GetUnavailableDates() []*string {
	return s.UnavailableDates
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule) SetAvailableWeeks(v []*int32) *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule {
	s.AvailableWeeks = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule) SetFrom(v *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom) *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule {
	s.From = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule) SetTo(v *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo) *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule {
	s.To = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule) SetUnavailableDates(v []*string) *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule {
	s.UnavailableDates = v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRule) Validate() error {
	if s.From != nil {
		if err := s.From.Validate(); err != nil {
			return err
		}
	}
	if s.To != nil {
		if err := s.To.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom struct {
	// example:
	//
	// 1
	Anchor *int32 `json:"Anchor,omitempty" xml:"Anchor,omitempty"`
	// example:
	//
	// 2026-01-01
	FixedTime *string `json:"FixedTime,omitempty" xml:"FixedTime,omitempty"`
	// example:
	//
	// 18:00
	OffsetDayOfTime *string `json:"OffsetDayOfTime,omitempty" xml:"OffsetDayOfTime,omitempty"`
	// example:
	//
	// 1
	OffsetUnit *int32 `json:"OffsetUnit,omitempty" xml:"OffsetUnit,omitempty"`
	// example:
	//
	// 1
	OffsetValue *int32 `json:"OffsetValue,omitempty" xml:"OffsetValue,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom) GetAnchor() *int32 {
	return s.Anchor
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom) GetFixedTime() *string {
	return s.FixedTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom) GetOffsetDayOfTime() *string {
	return s.OffsetDayOfTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom) GetOffsetUnit() *int32 {
	return s.OffsetUnit
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom) GetOffsetValue() *int32 {
	return s.OffsetValue
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom) SetAnchor(v int32) *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom {
	s.Anchor = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom) SetFixedTime(v string) *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom {
	s.FixedTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom) SetOffsetDayOfTime(v string) *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom {
	s.OffsetDayOfTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom) SetOffsetUnit(v int32) *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom {
	s.OffsetUnit = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom) SetOffsetValue(v int32) *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom {
	s.OffsetValue = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleFrom) Validate() error {
	return dara.Validate(s)
}

type TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo struct {
	// example:
	//
	// 1
	Anchor *int32 `json:"Anchor,omitempty" xml:"Anchor,omitempty"`
	// example:
	//
	// 2026-01-01
	FixedTime *string `json:"FixedTime,omitempty" xml:"FixedTime,omitempty"`
	// example:
	//
	// 18:00
	OffsetDayOfTime *string `json:"OffsetDayOfTime,omitempty" xml:"OffsetDayOfTime,omitempty"`
	// example:
	//
	// 1
	OffsetUnit *int32 `json:"OffsetUnit,omitempty" xml:"OffsetUnit,omitempty"`
	// example:
	//
	// 1
	OffsetValue *int32 `json:"OffsetValue,omitempty" xml:"OffsetValue,omitempty"`
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo) GetAnchor() *int32 {
	return s.Anchor
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo) GetFixedTime() *string {
	return s.FixedTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo) GetOffsetDayOfTime() *string {
	return s.OffsetDayOfTime
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo) GetOffsetUnit() *int32 {
	return s.OffsetUnit
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo) GetOffsetValue() *int32 {
	return s.OffsetValue
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo) SetAnchor(v int32) *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo {
	s.Anchor = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo) SetFixedTime(v string) *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo {
	s.FixedTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo) SetOffsetDayOfTime(v string) *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo {
	s.OffsetDayOfTime = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo) SetOffsetUnit(v int32) *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo {
	s.OffsetUnit = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo) SetOffsetValue(v int32) *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo {
	s.OffsetValue = &v
	return s
}

func (s *TicketPageQueryProductResponseBodyDataProductsUseRuleValidityPeriodRuleTo) Validate() error {
	return dara.Validate(s)
}
