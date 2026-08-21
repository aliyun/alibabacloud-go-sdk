// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TicketQueryProductResponseBodyData) *TicketQueryProductResponseBody
	GetData() *TicketQueryProductResponseBodyData
	SetErrorCode(v string) *TicketQueryProductResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *TicketQueryProductResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *TicketQueryProductResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketQueryProductResponseBody
	GetSuccess() *bool
}

type TicketQueryProductResponseBody struct {
	Data *TicketQueryProductResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s TicketQueryProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBody) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBody) GetData() *TicketQueryProductResponseBodyData {
	return s.Data
}

func (s *TicketQueryProductResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TicketQueryProductResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TicketQueryProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketQueryProductResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketQueryProductResponseBody) SetData(v *TicketQueryProductResponseBodyData) *TicketQueryProductResponseBody {
	s.Data = v
	return s
}

func (s *TicketQueryProductResponseBody) SetErrorCode(v string) *TicketQueryProductResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketQueryProductResponseBody) SetErrorMsg(v string) *TicketQueryProductResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketQueryProductResponseBody) SetRequestId(v string) *TicketQueryProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketQueryProductResponseBody) SetSuccess(v bool) *TicketQueryProductResponseBody {
	s.Success = &v
	return s
}

func (s *TicketQueryProductResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketQueryProductResponseBodyData struct {
	Product *TicketQueryProductResponseBodyDataProduct `json:"Product,omitempty" xml:"Product,omitempty" type:"Struct"`
}

func (s TicketQueryProductResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyData) GetProduct() *TicketQueryProductResponseBodyDataProduct {
	return s.Product
}

func (s *TicketQueryProductResponseBodyData) SetProduct(v *TicketQueryProductResponseBodyDataProduct) *TicketQueryProductResponseBodyData {
	s.Product = v
	return s
}

func (s *TicketQueryProductResponseBodyData) Validate() error {
	if s.Product != nil {
		if err := s.Product.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketQueryProductResponseBodyDataProduct struct {
	// example:
	//
	// 2
	BookingType *int32                                            `json:"BookingType,omitempty" xml:"BookingType,omitempty"`
	BuyRule     *TicketQueryProductResponseBodyDataProductBuyRule `json:"BuyRule,omitempty" xml:"BuyRule,omitempty" type:"Struct"`
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
	ProductName *string                                              `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	RefundRule  *TicketQueryProductResponseBodyDataProductRefundRule `json:"RefundRule,omitempty" xml:"RefundRule,omitempty" type:"Struct"`
	Region      *TicketQueryProductResponseBodyDataProductRegion     `json:"Region,omitempty" xml:"Region,omitempty" type:"Struct"`
	// example:
	//
	// 123456
	ScenicId *int64                                            `json:"ScenicId,omitempty" xml:"ScenicId,omitempty"`
	Session  *TicketQueryProductResponseBodyDataProductSession `json:"Session,omitempty" xml:"Session,omitempty" type:"Struct"`
	// example:
	//
	// 1
	SettlePriceCalculateType *int32                                        `json:"SettlePriceCalculateType,omitempty" xml:"SettlePriceCalculateType,omitempty"`
	Spu                      *TicketQueryProductResponseBodyDataProductSpu `json:"Spu,omitempty" xml:"Spu,omitempty" type:"Struct"`
	// example:
	//
	// 飞猪景区乐园旗舰店
	SupplierName *string                                              `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	TicketKind   *TicketQueryProductResponseBodyDataProductTicketKind `json:"TicketKind,omitempty" xml:"TicketKind,omitempty" type:"Struct"`
	UseRule      *TicketQueryProductResponseBodyDataProductUseRule    `json:"UseRule,omitempty" xml:"UseRule,omitempty" type:"Struct"`
}

func (s TicketQueryProductResponseBodyDataProduct) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProduct) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProduct) GetBookingType() *int32 {
	return s.BookingType
}

func (s *TicketQueryProductResponseBodyDataProduct) GetBuyRule() *TicketQueryProductResponseBodyDataProductBuyRule {
	return s.BuyRule
}

func (s *TicketQueryProductResponseBodyDataProduct) GetCostIncludeRemark() *string {
	return s.CostIncludeRemark
}

func (s *TicketQueryProductResponseBodyDataProduct) GetDeliverGuaranteeMinutes() *int32 {
	return s.DeliverGuaranteeMinutes
}

func (s *TicketQueryProductResponseBodyDataProduct) GetImages() []*string {
	return s.Images
}

func (s *TicketQueryProductResponseBodyDataProduct) GetInvoiceIssuerType() *int32 {
	return s.InvoiceIssuerType
}

func (s *TicketQueryProductResponseBodyDataProduct) GetPaymentLimitMinutes() *int32 {
	return s.PaymentLimitMinutes
}

func (s *TicketQueryProductResponseBodyDataProduct) GetProductId() *string {
	return s.ProductId
}

func (s *TicketQueryProductResponseBodyDataProduct) GetProductName() *string {
	return s.ProductName
}

func (s *TicketQueryProductResponseBodyDataProduct) GetRefundRule() *TicketQueryProductResponseBodyDataProductRefundRule {
	return s.RefundRule
}

func (s *TicketQueryProductResponseBodyDataProduct) GetRegion() *TicketQueryProductResponseBodyDataProductRegion {
	return s.Region
}

func (s *TicketQueryProductResponseBodyDataProduct) GetScenicId() *int64 {
	return s.ScenicId
}

func (s *TicketQueryProductResponseBodyDataProduct) GetSession() *TicketQueryProductResponseBodyDataProductSession {
	return s.Session
}

func (s *TicketQueryProductResponseBodyDataProduct) GetSettlePriceCalculateType() *int32 {
	return s.SettlePriceCalculateType
}

func (s *TicketQueryProductResponseBodyDataProduct) GetSpu() *TicketQueryProductResponseBodyDataProductSpu {
	return s.Spu
}

func (s *TicketQueryProductResponseBodyDataProduct) GetSupplierName() *string {
	return s.SupplierName
}

func (s *TicketQueryProductResponseBodyDataProduct) GetTicketKind() *TicketQueryProductResponseBodyDataProductTicketKind {
	return s.TicketKind
}

func (s *TicketQueryProductResponseBodyDataProduct) GetUseRule() *TicketQueryProductResponseBodyDataProductUseRule {
	return s.UseRule
}

func (s *TicketQueryProductResponseBodyDataProduct) SetBookingType(v int32) *TicketQueryProductResponseBodyDataProduct {
	s.BookingType = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetBuyRule(v *TicketQueryProductResponseBodyDataProductBuyRule) *TicketQueryProductResponseBodyDataProduct {
	s.BuyRule = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetCostIncludeRemark(v string) *TicketQueryProductResponseBodyDataProduct {
	s.CostIncludeRemark = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetDeliverGuaranteeMinutes(v int32) *TicketQueryProductResponseBodyDataProduct {
	s.DeliverGuaranteeMinutes = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetImages(v []*string) *TicketQueryProductResponseBodyDataProduct {
	s.Images = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetInvoiceIssuerType(v int32) *TicketQueryProductResponseBodyDataProduct {
	s.InvoiceIssuerType = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetPaymentLimitMinutes(v int32) *TicketQueryProductResponseBodyDataProduct {
	s.PaymentLimitMinutes = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetProductId(v string) *TicketQueryProductResponseBodyDataProduct {
	s.ProductId = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetProductName(v string) *TicketQueryProductResponseBodyDataProduct {
	s.ProductName = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetRefundRule(v *TicketQueryProductResponseBodyDataProductRefundRule) *TicketQueryProductResponseBodyDataProduct {
	s.RefundRule = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetRegion(v *TicketQueryProductResponseBodyDataProductRegion) *TicketQueryProductResponseBodyDataProduct {
	s.Region = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetScenicId(v int64) *TicketQueryProductResponseBodyDataProduct {
	s.ScenicId = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetSession(v *TicketQueryProductResponseBodyDataProductSession) *TicketQueryProductResponseBodyDataProduct {
	s.Session = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetSettlePriceCalculateType(v int32) *TicketQueryProductResponseBodyDataProduct {
	s.SettlePriceCalculateType = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetSpu(v *TicketQueryProductResponseBodyDataProductSpu) *TicketQueryProductResponseBodyDataProduct {
	s.Spu = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetSupplierName(v string) *TicketQueryProductResponseBodyDataProduct {
	s.SupplierName = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetTicketKind(v *TicketQueryProductResponseBodyDataProductTicketKind) *TicketQueryProductResponseBodyDataProduct {
	s.TicketKind = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) SetUseRule(v *TicketQueryProductResponseBodyDataProductUseRule) *TicketQueryProductResponseBodyDataProduct {
	s.UseRule = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProduct) Validate() error {
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

type TicketQueryProductResponseBodyDataProductBuyRule struct {
	AheadBuyTimePointRule           *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule             `json:"AheadBuyTimePointRule,omitempty" xml:"AheadBuyTimePointRule,omitempty" type:"Struct"`
	ContactRule                     *TicketQueryProductResponseBodyDataProductBuyRuleContactRule                       `json:"ContactRule,omitempty" xml:"ContactRule,omitempty" type:"Struct"`
	CrossOrderBuyQuantityLimitRules []*TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules `json:"CrossOrderBuyQuantityLimitRules,omitempty" xml:"CrossOrderBuyQuantityLimitRules,omitempty" type:"Repeated"`
	PerOrderBuyQuantityLimitRule    *TicketQueryProductResponseBodyDataProductBuyRulePerOrderBuyQuantityLimitRule      `json:"PerOrderBuyQuantityLimitRule,omitempty" xml:"PerOrderBuyQuantityLimitRule,omitempty" type:"Struct"`
	TravelerRule                    *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule                      `json:"TravelerRule,omitempty" xml:"TravelerRule,omitempty" type:"Struct"`
}

func (s TicketQueryProductResponseBodyDataProductBuyRule) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductBuyRule) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductBuyRule) GetAheadBuyTimePointRule() *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule {
	return s.AheadBuyTimePointRule
}

func (s *TicketQueryProductResponseBodyDataProductBuyRule) GetContactRule() *TicketQueryProductResponseBodyDataProductBuyRuleContactRule {
	return s.ContactRule
}

func (s *TicketQueryProductResponseBodyDataProductBuyRule) GetCrossOrderBuyQuantityLimitRules() []*TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules {
	return s.CrossOrderBuyQuantityLimitRules
}

func (s *TicketQueryProductResponseBodyDataProductBuyRule) GetPerOrderBuyQuantityLimitRule() *TicketQueryProductResponseBodyDataProductBuyRulePerOrderBuyQuantityLimitRule {
	return s.PerOrderBuyQuantityLimitRule
}

func (s *TicketQueryProductResponseBodyDataProductBuyRule) GetTravelerRule() *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule {
	return s.TravelerRule
}

func (s *TicketQueryProductResponseBodyDataProductBuyRule) SetAheadBuyTimePointRule(v *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule) *TicketQueryProductResponseBodyDataProductBuyRule {
	s.AheadBuyTimePointRule = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRule) SetContactRule(v *TicketQueryProductResponseBodyDataProductBuyRuleContactRule) *TicketQueryProductResponseBodyDataProductBuyRule {
	s.ContactRule = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRule) SetCrossOrderBuyQuantityLimitRules(v []*TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) *TicketQueryProductResponseBodyDataProductBuyRule {
	s.CrossOrderBuyQuantityLimitRules = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRule) SetPerOrderBuyQuantityLimitRule(v *TicketQueryProductResponseBodyDataProductBuyRulePerOrderBuyQuantityLimitRule) *TicketQueryProductResponseBodyDataProductBuyRule {
	s.PerOrderBuyQuantityLimitRule = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRule) SetTravelerRule(v *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) *TicketQueryProductResponseBodyDataProductBuyRule {
	s.TravelerRule = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRule) Validate() error {
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

type TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule struct {
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

func (s TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule) GetAnchor() *int32 {
	return s.Anchor
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule) GetFixedTime() *string {
	return s.FixedTime
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule) GetOffsetDayOfTime() *string {
	return s.OffsetDayOfTime
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule) GetOffsetUnit() *int32 {
	return s.OffsetUnit
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule) GetOffsetValue() *int32 {
	return s.OffsetValue
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule) SetAnchor(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule {
	s.Anchor = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule) SetFixedTime(v string) *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule {
	s.FixedTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule) SetOffsetDayOfTime(v string) *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule {
	s.OffsetDayOfTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule) SetOffsetUnit(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule {
	s.OffsetUnit = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule) SetOffsetValue(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule {
	s.OffsetValue = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleAheadBuyTimePointRule) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductBuyRuleContactRule struct {
	ContactFieldRule *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule `json:"ContactFieldRule,omitempty" xml:"ContactFieldRule,omitempty" type:"Struct"`
}

func (s TicketQueryProductResponseBodyDataProductBuyRuleContactRule) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductBuyRuleContactRule) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRule) GetContactFieldRule() *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule {
	return s.ContactFieldRule
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRule) SetContactFieldRule(v *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) *TicketQueryProductResponseBodyDataProductBuyRuleContactRule {
	s.ContactFieldRule = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRule) Validate() error {
	if s.ContactFieldRule != nil {
		if err := s.ContactFieldRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule struct {
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

func (s TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) GetCertificate() *bool {
	return s.Certificate
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) GetCertificateTypes() []*int32 {
	return s.CertificateTypes
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) GetDialingCode() *bool {
	return s.DialingCode
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) GetEmail() *bool {
	return s.Email
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) GetFirstName() *bool {
	return s.FirstName
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) GetLastName() *bool {
	return s.LastName
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) GetMobile() *bool {
	return s.Mobile
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) GetName() *bool {
	return s.Name
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) SetCertificate(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule {
	s.Certificate = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) SetCertificateTypes(v []*int32) *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule {
	s.CertificateTypes = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) SetDialingCode(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule {
	s.DialingCode = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) SetEmail(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule {
	s.Email = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) SetFirstName(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule {
	s.FirstName = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) SetLastName(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule {
	s.LastName = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) SetMobile(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule {
	s.Mobile = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) SetName(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule {
	s.Name = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleContactRuleContactFieldRule) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules struct {
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

func (s TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) GetLimitDayType() *int32 {
	return s.LimitDayType
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) GetLimitDays() *int32 {
	return s.LimitDays
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) GetLimitPeriod() *int32 {
	return s.LimitPeriod
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) GetLimitQuantityType() *int32 {
	return s.LimitQuantityType
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) GetLimitType() *int32 {
	return s.LimitType
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) GetMaxBuyQuantity() *int32 {
	return s.MaxBuyQuantity
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) SetLimitDayType(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules {
	s.LimitDayType = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) SetLimitDays(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules {
	s.LimitDays = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) SetLimitPeriod(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules {
	s.LimitPeriod = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) SetLimitQuantityType(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules {
	s.LimitQuantityType = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) SetLimitType(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules {
	s.LimitType = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) SetMaxBuyQuantity(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules {
	s.MaxBuyQuantity = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleCrossOrderBuyQuantityLimitRules) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductBuyRulePerOrderBuyQuantityLimitRule struct {
	// example:
	//
	// 10
	MaxBuyQuantity *int32 `json:"MaxBuyQuantity,omitempty" xml:"MaxBuyQuantity,omitempty"`
	// example:
	//
	// 1
	MinBuyQuantity *int32 `json:"MinBuyQuantity,omitempty" xml:"MinBuyQuantity,omitempty"`
}

func (s TicketQueryProductResponseBodyDataProductBuyRulePerOrderBuyQuantityLimitRule) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductBuyRulePerOrderBuyQuantityLimitRule) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductBuyRulePerOrderBuyQuantityLimitRule) GetMaxBuyQuantity() *int32 {
	return s.MaxBuyQuantity
}

func (s *TicketQueryProductResponseBodyDataProductBuyRulePerOrderBuyQuantityLimitRule) GetMinBuyQuantity() *int32 {
	return s.MinBuyQuantity
}

func (s *TicketQueryProductResponseBodyDataProductBuyRulePerOrderBuyQuantityLimitRule) SetMaxBuyQuantity(v int32) *TicketQueryProductResponseBodyDataProductBuyRulePerOrderBuyQuantityLimitRule {
	s.MaxBuyQuantity = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRulePerOrderBuyQuantityLimitRule) SetMinBuyQuantity(v int32) *TicketQueryProductResponseBodyDataProductBuyRulePerOrderBuyQuantityLimitRule {
	s.MinBuyQuantity = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRulePerOrderBuyQuantityLimitRule) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule struct {
	CrowdLimitRules     []*TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules     `json:"CrowdLimitRules,omitempty" xml:"CrowdLimitRules,omitempty" type:"Repeated"`
	CrowdQuantityLimits []*TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdQuantityLimits `json:"CrowdQuantityLimits,omitempty" xml:"CrowdQuantityLimits,omitempty" type:"Repeated"`
	// example:
	//
	// true
	NeedFillTraveler  *bool                                                                          `json:"NeedFillTraveler,omitempty" xml:"NeedFillTraveler,omitempty"`
	TravelerFieldRule *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule `json:"TravelerFieldRule,omitempty" xml:"TravelerFieldRule,omitempty" type:"Struct"`
	// example:
	//
	// 2
	TravelerFillDimension *int32 `json:"TravelerFillDimension,omitempty" xml:"TravelerFillDimension,omitempty"`
	// example:
	//
	// 3
	TravelerQuantity *int32 `json:"TravelerQuantity,omitempty" xml:"TravelerQuantity,omitempty"`
}

func (s TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) GetCrowdLimitRules() []*TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules {
	return s.CrowdLimitRules
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) GetCrowdQuantityLimits() []*TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdQuantityLimits {
	return s.CrowdQuantityLimits
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) GetNeedFillTraveler() *bool {
	return s.NeedFillTraveler
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) GetTravelerFieldRule() *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule {
	return s.TravelerFieldRule
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) GetTravelerFillDimension() *int32 {
	return s.TravelerFillDimension
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) GetTravelerQuantity() *int32 {
	return s.TravelerQuantity
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) SetCrowdLimitRules(v []*TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule {
	s.CrowdLimitRules = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) SetCrowdQuantityLimits(v []*TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdQuantityLimits) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule {
	s.CrowdQuantityLimits = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) SetNeedFillTraveler(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule {
	s.NeedFillTraveler = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) SetTravelerFieldRule(v *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule {
	s.TravelerFieldRule = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) SetTravelerFillDimension(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule {
	s.TravelerFillDimension = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) SetTravelerQuantity(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule {
	s.TravelerQuantity = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRule) Validate() error {
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

type TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules struct {
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

func (s TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules) GetAgeBaseTimeType() *int32 {
	return s.AgeBaseTimeType
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules) GetAgeCalculateType() *int32 {
	return s.AgeCalculateType
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules) GetAgeMax() *int32 {
	return s.AgeMax
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules) GetAgeMin() *int32 {
	return s.AgeMin
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules) GetName() *string {
	return s.Name
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules) SetAgeBaseTimeType(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules {
	s.AgeBaseTimeType = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules) SetAgeCalculateType(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules {
	s.AgeCalculateType = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules) SetAgeMax(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules {
	s.AgeMax = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules) SetAgeMin(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules {
	s.AgeMin = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules) SetName(v string) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules {
	s.Name = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdLimitRules) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdQuantityLimits struct {
	// example:
	//
	// 成人票
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdQuantityLimits) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdQuantityLimits) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdQuantityLimits) GetName() *string {
	return s.Name
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdQuantityLimits) GetQuantity() *int32 {
	return s.Quantity
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdQuantityLimits) SetName(v string) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdQuantityLimits {
	s.Name = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdQuantityLimits) SetQuantity(v int32) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdQuantityLimits {
	s.Quantity = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleCrowdQuantityLimits) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule struct {
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

func (s TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) GetBirthday() *bool {
	return s.Birthday
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) GetCertificate() *bool {
	return s.Certificate
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) GetCertificateTypes() []*int32 {
	return s.CertificateTypes
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) GetDialingCode() *bool {
	return s.DialingCode
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) GetEmail() *bool {
	return s.Email
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) GetFirstName() *bool {
	return s.FirstName
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) GetGender() *bool {
	return s.Gender
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) GetLastName() *bool {
	return s.LastName
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) GetMobile() *bool {
	return s.Mobile
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) GetName() *bool {
	return s.Name
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) GetNationality() *bool {
	return s.Nationality
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) SetBirthday(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule {
	s.Birthday = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) SetCertificate(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule {
	s.Certificate = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) SetCertificateTypes(v []*int32) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule {
	s.CertificateTypes = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) SetDialingCode(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule {
	s.DialingCode = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) SetEmail(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule {
	s.Email = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) SetFirstName(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule {
	s.FirstName = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) SetGender(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule {
	s.Gender = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) SetLastName(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule {
	s.LastName = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) SetMobile(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule {
	s.Mobile = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) SetName(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule {
	s.Name = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) SetNationality(v bool) *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule {
	s.Nationality = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductBuyRuleTravelerRuleTravelerFieldRule) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductRefundRule struct {
	RefundStageRules []*TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules `json:"RefundStageRules,omitempty" xml:"RefundStageRules,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	RefundType *int32 `json:"RefundType,omitempty" xml:"RefundType,omitempty"`
}

func (s TicketQueryProductResponseBodyDataProductRefundRule) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductRefundRule) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductRefundRule) GetRefundStageRules() []*TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules {
	return s.RefundStageRules
}

func (s *TicketQueryProductResponseBodyDataProductRefundRule) GetRefundType() *int32 {
	return s.RefundType
}

func (s *TicketQueryProductResponseBodyDataProductRefundRule) SetRefundStageRules(v []*TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules) *TicketQueryProductResponseBodyDataProductRefundRule {
	s.RefundStageRules = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRule) SetRefundType(v int32) *TicketQueryProductResponseBodyDataProductRefundRule {
	s.RefundType = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRule) Validate() error {
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

type TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules struct {
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
	FeeType *int32                                                                   `json:"FeeType,omitempty" xml:"FeeType,omitempty"`
	From    *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom `json:"From,omitempty" xml:"From,omitempty" type:"Struct"`
	To      *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo   `json:"To,omitempty" xml:"To,omitempty" type:"Struct"`
}

func (s TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules) GetFee() *float64 {
	return s.Fee
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules) GetFeeBase() *int32 {
	return s.FeeBase
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules) GetFeeType() *int32 {
	return s.FeeType
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules) GetFrom() *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom {
	return s.From
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules) GetTo() *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo {
	return s.To
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules) SetFee(v float64) *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules {
	s.Fee = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules) SetFeeBase(v int32) *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules {
	s.FeeBase = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules) SetFeeType(v int32) *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules {
	s.FeeType = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules) SetFrom(v *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom) *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules {
	s.From = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules) SetTo(v *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo) *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules {
	s.To = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRules) Validate() error {
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

type TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom struct {
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

func (s TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom) GetAnchor() *int32 {
	return s.Anchor
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom) GetFixedTime() *string {
	return s.FixedTime
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom) GetOffsetDayOfTime() *string {
	return s.OffsetDayOfTime
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom) GetOffsetUnit() *int32 {
	return s.OffsetUnit
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom) GetOffsetValue() *int32 {
	return s.OffsetValue
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom) SetAnchor(v int32) *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom {
	s.Anchor = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom) SetFixedTime(v string) *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom {
	s.FixedTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom) SetOffsetDayOfTime(v string) *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom {
	s.OffsetDayOfTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom) SetOffsetUnit(v int32) *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom {
	s.OffsetUnit = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom) SetOffsetValue(v int32) *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom {
	s.OffsetValue = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesFrom) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo struct {
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

func (s TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo) GetAnchor() *int32 {
	return s.Anchor
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo) GetFixedTime() *string {
	return s.FixedTime
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo) GetOffsetDayOfTime() *string {
	return s.OffsetDayOfTime
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo) GetOffsetUnit() *int32 {
	return s.OffsetUnit
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo) GetOffsetValue() *int32 {
	return s.OffsetValue
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo) SetAnchor(v int32) *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo {
	s.Anchor = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo) SetFixedTime(v string) *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo {
	s.FixedTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo) SetOffsetDayOfTime(v string) *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo {
	s.OffsetDayOfTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo) SetOffsetUnit(v int32) *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo {
	s.OffsetUnit = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo) SetOffsetValue(v int32) *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo {
	s.OffsetValue = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRefundRuleRefundStageRulesTo) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductRegion struct {
	// example:
	//
	// 普通区
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s TicketQueryProductResponseBodyDataProductRegion) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductRegion) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductRegion) GetName() *string {
	return s.Name
}

func (s *TicketQueryProductResponseBodyDataProductRegion) SetName(v string) *TicketQueryProductResponseBodyDataProductRegion {
	s.Name = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductRegion) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductSession struct {
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

func (s TicketQueryProductResponseBodyDataProductSession) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductSession) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductSession) GetSessionEndTime() *string {
	return s.SessionEndTime
}

func (s *TicketQueryProductResponseBodyDataProductSession) GetSessionName() *string {
	return s.SessionName
}

func (s *TicketQueryProductResponseBodyDataProductSession) GetSessionStartTime() *string {
	return s.SessionStartTime
}

func (s *TicketQueryProductResponseBodyDataProductSession) SetSessionEndTime(v string) *TicketQueryProductResponseBodyDataProductSession {
	s.SessionEndTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductSession) SetSessionName(v string) *TicketQueryProductResponseBodyDataProductSession {
	s.SessionName = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductSession) SetSessionStartTime(v string) *TicketQueryProductResponseBodyDataProductSession {
	s.SessionStartTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductSession) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductSpu struct {
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

func (s TicketQueryProductResponseBodyDataProductSpu) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductSpu) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductSpu) GetPrimaryTypeName() *string {
	return s.PrimaryTypeName
}

func (s *TicketQueryProductResponseBodyDataProductSpu) GetReserveDetail() *string {
	return s.ReserveDetail
}

func (s *TicketQueryProductResponseBodyDataProductSpu) GetReserveTitle() *string {
	return s.ReserveTitle
}

func (s *TicketQueryProductResponseBodyDataProductSpu) GetSecondaryTypeName() *string {
	return s.SecondaryTypeName
}

func (s *TicketQueryProductResponseBodyDataProductSpu) GetSpuId() *int64 {
	return s.SpuId
}

func (s *TicketQueryProductResponseBodyDataProductSpu) GetSpuName() *string {
	return s.SpuName
}

func (s *TicketQueryProductResponseBodyDataProductSpu) SetPrimaryTypeName(v string) *TicketQueryProductResponseBodyDataProductSpu {
	s.PrimaryTypeName = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductSpu) SetReserveDetail(v string) *TicketQueryProductResponseBodyDataProductSpu {
	s.ReserveDetail = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductSpu) SetReserveTitle(v string) *TicketQueryProductResponseBodyDataProductSpu {
	s.ReserveTitle = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductSpu) SetSecondaryTypeName(v string) *TicketQueryProductResponseBodyDataProductSpu {
	s.SecondaryTypeName = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductSpu) SetSpuId(v int64) *TicketQueryProductResponseBodyDataProductSpu {
	s.SpuId = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductSpu) SetSpuName(v string) *TicketQueryProductResponseBodyDataProductSpu {
	s.SpuName = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductSpu) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductTicketKind struct {
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

func (s TicketQueryProductResponseBodyDataProductTicketKind) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductTicketKind) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductTicketKind) GetDescription() *string {
	return s.Description
}

func (s *TicketQueryProductResponseBodyDataProductTicketKind) GetId() *int64 {
	return s.Id
}

func (s *TicketQueryProductResponseBodyDataProductTicketKind) GetName() *string {
	return s.Name
}

func (s *TicketQueryProductResponseBodyDataProductTicketKind) SetDescription(v string) *TicketQueryProductResponseBodyDataProductTicketKind {
	s.Description = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductTicketKind) SetId(v int64) *TicketQueryProductResponseBodyDataProductTicketKind {
	s.Id = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductTicketKind) SetName(v string) *TicketQueryProductResponseBodyDataProductTicketKind {
	s.Name = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductTicketKind) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductUseRule struct {
	EffectTimePointRule *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule `json:"EffectTimePointRule,omitempty" xml:"EffectTimePointRule,omitempty" type:"Struct"`
	// example:
	//
	// 景区正门入口
	EntryAddress *string `json:"EntryAddress,omitempty" xml:"EntryAddress,omitempty"`
	// example:
	//
	// 请于入园前2小时至香港迪士尼度假区售票处旁人工服务站进行取票预约，服务时间上午 10:00 至 晚上 8:30
	EntryRemark      *string                                                             `json:"EntryRemark,omitempty" xml:"EntryRemark,omitempty"`
	EntryTimePeriods []*TicketQueryProductResponseBodyDataProductUseRuleEntryTimePeriods `json:"EntryTimePeriods,omitempty" xml:"EntryTimePeriods,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	EntryType            *int32                                                                `json:"EntryType,omitempty" xml:"EntryType,omitempty"`
	EntryWithVoucherRule *TicketQueryProductResponseBodyDataProductUseRuleEntryWithVoucherRule `json:"EntryWithVoucherRule,omitempty" xml:"EntryWithVoucherRule,omitempty" type:"Struct"`
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
	OtherNote          *string                                                             `json:"OtherNote,omitempty" xml:"OtherNote,omitempty"`
	PickupsRule        *TicketQueryProductResponseBodyDataProductUseRulePickupsRule        `json:"PickupsRule,omitempty" xml:"PickupsRule,omitempty" type:"Struct"`
	ValidityPeriodRule *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule `json:"ValidityPeriodRule,omitempty" xml:"ValidityPeriodRule,omitempty" type:"Struct"`
}

func (s TicketQueryProductResponseBodyDataProductUseRule) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductUseRule) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) GetEffectTimePointRule() *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule {
	return s.EffectTimePointRule
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) GetEntryAddress() *string {
	return s.EntryAddress
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) GetEntryRemark() *string {
	return s.EntryRemark
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) GetEntryTimePeriods() []*TicketQueryProductResponseBodyDataProductUseRuleEntryTimePeriods {
	return s.EntryTimePeriods
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) GetEntryType() *int32 {
	return s.EntryType
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) GetEntryWithVoucherRule() *TicketQueryProductResponseBodyDataProductUseRuleEntryWithVoucherRule {
	return s.EntryWithVoucherRule
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) GetNeedAssemble() *bool {
	return s.NeedAssemble
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) GetNeedPrebook() *bool {
	return s.NeedPrebook
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) GetOtherNote() *string {
	return s.OtherNote
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) GetPickupsRule() *TicketQueryProductResponseBodyDataProductUseRulePickupsRule {
	return s.PickupsRule
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) GetValidityPeriodRule() *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule {
	return s.ValidityPeriodRule
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) SetEffectTimePointRule(v *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule) *TicketQueryProductResponseBodyDataProductUseRule {
	s.EffectTimePointRule = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) SetEntryAddress(v string) *TicketQueryProductResponseBodyDataProductUseRule {
	s.EntryAddress = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) SetEntryRemark(v string) *TicketQueryProductResponseBodyDataProductUseRule {
	s.EntryRemark = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) SetEntryTimePeriods(v []*TicketQueryProductResponseBodyDataProductUseRuleEntryTimePeriods) *TicketQueryProductResponseBodyDataProductUseRule {
	s.EntryTimePeriods = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) SetEntryType(v int32) *TicketQueryProductResponseBodyDataProductUseRule {
	s.EntryType = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) SetEntryWithVoucherRule(v *TicketQueryProductResponseBodyDataProductUseRuleEntryWithVoucherRule) *TicketQueryProductResponseBodyDataProductUseRule {
	s.EntryWithVoucherRule = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) SetNeedAssemble(v bool) *TicketQueryProductResponseBodyDataProductUseRule {
	s.NeedAssemble = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) SetNeedPrebook(v bool) *TicketQueryProductResponseBodyDataProductUseRule {
	s.NeedPrebook = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) SetOtherNote(v string) *TicketQueryProductResponseBodyDataProductUseRule {
	s.OtherNote = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) SetPickupsRule(v *TicketQueryProductResponseBodyDataProductUseRulePickupsRule) *TicketQueryProductResponseBodyDataProductUseRule {
	s.PickupsRule = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) SetValidityPeriodRule(v *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule) *TicketQueryProductResponseBodyDataProductUseRule {
	s.ValidityPeriodRule = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRule) Validate() error {
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

type TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule struct {
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

func (s TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule) GetAnchor() *int32 {
	return s.Anchor
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule) GetFixedTime() *string {
	return s.FixedTime
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule) GetOffsetDayOfTime() *string {
	return s.OffsetDayOfTime
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule) GetOffsetUnit() *int32 {
	return s.OffsetUnit
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule) GetOffsetValue() *int32 {
	return s.OffsetValue
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule) SetAnchor(v int32) *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule {
	s.Anchor = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule) SetFixedTime(v string) *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule {
	s.FixedTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule) SetOffsetDayOfTime(v string) *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule {
	s.OffsetDayOfTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule) SetOffsetUnit(v int32) *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule {
	s.OffsetUnit = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule) SetOffsetValue(v int32) *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule {
	s.OffsetValue = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEffectTimePointRule) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductUseRuleEntryTimePeriods struct {
	// example:
	//
	// 08:00
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 17:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s TicketQueryProductResponseBodyDataProductUseRuleEntryTimePeriods) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductUseRuleEntryTimePeriods) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEntryTimePeriods) GetBeginTime() *string {
	return s.BeginTime
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEntryTimePeriods) GetEndTime() *string {
	return s.EndTime
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEntryTimePeriods) SetBeginTime(v string) *TicketQueryProductResponseBodyDataProductUseRuleEntryTimePeriods {
	s.BeginTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEntryTimePeriods) SetEndTime(v string) *TicketQueryProductResponseBodyDataProductUseRuleEntryTimePeriods {
	s.EndTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEntryTimePeriods) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductUseRuleEntryWithVoucherRule struct {
	// example:
	//
	// 凭二维码入园
	VoucherRemark *string `json:"VoucherRemark,omitempty" xml:"VoucherRemark,omitempty"`
	// example:
	//
	// [1]
	VoucherTypes []*int32 `json:"VoucherTypes,omitempty" xml:"VoucherTypes,omitempty" type:"Repeated"`
}

func (s TicketQueryProductResponseBodyDataProductUseRuleEntryWithVoucherRule) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductUseRuleEntryWithVoucherRule) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEntryWithVoucherRule) GetVoucherRemark() *string {
	return s.VoucherRemark
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEntryWithVoucherRule) GetVoucherTypes() []*int32 {
	return s.VoucherTypes
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEntryWithVoucherRule) SetVoucherRemark(v string) *TicketQueryProductResponseBodyDataProductUseRuleEntryWithVoucherRule {
	s.VoucherRemark = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEntryWithVoucherRule) SetVoucherTypes(v []*int32) *TicketQueryProductResponseBodyDataProductUseRuleEntryWithVoucherRule {
	s.VoucherTypes = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleEntryWithVoucherRule) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductUseRulePickupsRule struct {
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

func (s TicketQueryProductResponseBodyDataProductUseRulePickupsRule) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductUseRulePickupsRule) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductUseRulePickupsRule) GetPickupsAddress() *string {
	return s.PickupsAddress
}

func (s *TicketQueryProductResponseBodyDataProductUseRulePickupsRule) GetVoucherRemark() *string {
	return s.VoucherRemark
}

func (s *TicketQueryProductResponseBodyDataProductUseRulePickupsRule) GetVoucherTypes() []*int32 {
	return s.VoucherTypes
}

func (s *TicketQueryProductResponseBodyDataProductUseRulePickupsRule) SetPickupsAddress(v string) *TicketQueryProductResponseBodyDataProductUseRulePickupsRule {
	s.PickupsAddress = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRulePickupsRule) SetVoucherRemark(v string) *TicketQueryProductResponseBodyDataProductUseRulePickupsRule {
	s.VoucherRemark = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRulePickupsRule) SetVoucherTypes(v []*int32) *TicketQueryProductResponseBodyDataProductUseRulePickupsRule {
	s.VoucherTypes = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRulePickupsRule) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule struct {
	// example:
	//
	// [1, 2, 3, 4, 5, 6, 7]
	AvailableWeeks []*int32                                                                `json:"AvailableWeeks,omitempty" xml:"AvailableWeeks,omitempty" type:"Repeated"`
	From           *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom `json:"From,omitempty" xml:"From,omitempty" type:"Struct"`
	To             *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo   `json:"To,omitempty" xml:"To,omitempty" type:"Struct"`
	// example:
	//
	// ["2026-01-01"]
	UnavailableDates []*string `json:"UnavailableDates,omitempty" xml:"UnavailableDates,omitempty" type:"Repeated"`
}

func (s TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule) GetAvailableWeeks() []*int32 {
	return s.AvailableWeeks
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule) GetFrom() *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom {
	return s.From
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule) GetTo() *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo {
	return s.To
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule) GetUnavailableDates() []*string {
	return s.UnavailableDates
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule) SetAvailableWeeks(v []*int32) *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule {
	s.AvailableWeeks = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule) SetFrom(v *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom) *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule {
	s.From = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule) SetTo(v *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo) *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule {
	s.To = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule) SetUnavailableDates(v []*string) *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule {
	s.UnavailableDates = v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRule) Validate() error {
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

type TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom struct {
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

func (s TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom) GetAnchor() *int32 {
	return s.Anchor
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom) GetFixedTime() *string {
	return s.FixedTime
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom) GetOffsetDayOfTime() *string {
	return s.OffsetDayOfTime
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom) GetOffsetUnit() *int32 {
	return s.OffsetUnit
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom) GetOffsetValue() *int32 {
	return s.OffsetValue
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom) SetAnchor(v int32) *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom {
	s.Anchor = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom) SetFixedTime(v string) *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom {
	s.FixedTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom) SetOffsetDayOfTime(v string) *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom {
	s.OffsetDayOfTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom) SetOffsetUnit(v int32) *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom {
	s.OffsetUnit = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom) SetOffsetValue(v int32) *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom {
	s.OffsetValue = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleFrom) Validate() error {
	return dara.Validate(s)
}

type TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo struct {
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

func (s TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo) GetAnchor() *int32 {
	return s.Anchor
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo) GetFixedTime() *string {
	return s.FixedTime
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo) GetOffsetDayOfTime() *string {
	return s.OffsetDayOfTime
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo) GetOffsetUnit() *int32 {
	return s.OffsetUnit
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo) GetOffsetValue() *int32 {
	return s.OffsetValue
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo) SetAnchor(v int32) *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo {
	s.Anchor = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo) SetFixedTime(v string) *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo {
	s.FixedTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo) SetOffsetDayOfTime(v string) *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo {
	s.OffsetDayOfTime = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo) SetOffsetUnit(v int32) *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo {
	s.OffsetUnit = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo) SetOffsetValue(v int32) *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo {
	s.OffsetValue = &v
	return s
}

func (s *TicketQueryProductResponseBodyDataProductUseRuleValidityPeriodRuleTo) Validate() error {
	return dara.Validate(s)
}
