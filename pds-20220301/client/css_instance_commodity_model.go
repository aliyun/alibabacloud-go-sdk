// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCssInstanceCommodity interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v int64) *CssInstanceCommodity
	GetActivityId() *int64
	SetAliyunProduceCode(v string) *CssInstanceCommodity
	GetAliyunProduceCode() *string
	SetChargeType(v string) *CssInstanceCommodity
	GetChargeType() *string
	SetCommodityCode(v string) *CssInstanceCommodity
	GetCommodityCode() *string
	SetComponents(v []*CssInstanceComponent) *CssInstanceCommodity
	GetComponents() []*CssInstanceComponent
	SetDuration(v int64) *CssInstanceCommodity
	GetDuration() *int64
	SetInstanceId(v string) *CssInstanceCommodity
	GetInstanceId() *string
	SetIsFree(v bool) *CssInstanceCommodity
	GetIsFree() *bool
	SetIsPrePayPostCharge(v bool) *CssInstanceCommodity
	GetIsPrePayPostCharge() *bool
	SetIsRenewChange(v bool) *CssInstanceCommodity
	GetIsRenewChange() *bool
	SetIsSyncToSubscription(v bool) *CssInstanceCommodity
	GetIsSyncToSubscription() *bool
	SetOrderParams(v map[string]*string) *CssInstanceCommodity
	GetOrderParams() map[string]*string
	SetOrderType(v string) *CssInstanceCommodity
	GetOrderType() *string
	SetPlanItemId(v int64) *CssInstanceCommodity
	GetPlanItemId() *int64
	SetPricingCycle(v string) *CssInstanceCommodity
	GetPricingCycle() *string
	SetQuantity(v int64) *CssInstanceCommodity
	GetQuantity() *int64
	SetRedeemNoList(v []*string) *CssInstanceCommodity
	GetRedeemNoList() []*string
	SetRedeemOrderType(v string) *CssInstanceCommodity
	GetRedeemOrderType() *string
	SetRefundSpecCode(v string) *CssInstanceCommodity
	GetRefundSpecCode() *string
	SetSpecCode(v string) *CssInstanceCommodity
	GetSpecCode() *string
	SetSpecUpgradeOriginSpecCodes(v []*string) *CssInstanceCommodity
	GetSpecUpgradeOriginSpecCodes() []*string
	SetSpecifyStartDate(v int64) *CssInstanceCommodity
	GetSpecifyStartDate() *int64
	SetUpgradeInquireFinancialValue(v bool) *CssInstanceCommodity
	GetUpgradeInquireFinancialValue() *bool
}

type CssInstanceCommodity struct {
	ActivityId                   *int64                  `json:"activityId,omitempty" xml:"activityId,omitempty"`
	AliyunProduceCode            *string                 `json:"aliyunProduceCode,omitempty" xml:"aliyunProduceCode,omitempty"`
	ChargeType                   *string                 `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	CommodityCode                *string                 `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	Components                   []*CssInstanceComponent `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	Duration                     *int64                  `json:"duration,omitempty" xml:"duration,omitempty"`
	InstanceId                   *string                 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsFree                       *bool                   `json:"isFree,omitempty" xml:"isFree,omitempty"`
	IsPrePayPostCharge           *bool                   `json:"isPrePayPostCharge,omitempty" xml:"isPrePayPostCharge,omitempty"`
	IsRenewChange                *bool                   `json:"isRenewChange,omitempty" xml:"isRenewChange,omitempty"`
	IsSyncToSubscription         *bool                   `json:"isSyncToSubscription,omitempty" xml:"isSyncToSubscription,omitempty"`
	OrderParams                  map[string]*string      `json:"orderParams,omitempty" xml:"orderParams,omitempty"`
	OrderType                    *string                 `json:"orderType,omitempty" xml:"orderType,omitempty"`
	PlanItemId                   *int64                  `json:"planItemId,omitempty" xml:"planItemId,omitempty"`
	PricingCycle                 *string                 `json:"pricingCycle,omitempty" xml:"pricingCycle,omitempty"`
	Quantity                     *int64                  `json:"quantity,omitempty" xml:"quantity,omitempty"`
	RedeemNoList                 []*string               `json:"redeemNoList,omitempty" xml:"redeemNoList,omitempty" type:"Repeated"`
	RedeemOrderType              *string                 `json:"redeemOrderType,omitempty" xml:"redeemOrderType,omitempty"`
	RefundSpecCode               *string                 `json:"refundSpecCode,omitempty" xml:"refundSpecCode,omitempty"`
	SpecCode                     *string                 `json:"specCode,omitempty" xml:"specCode,omitempty"`
	SpecUpgradeOriginSpecCodes   []*string               `json:"specUpgradeOriginSpecCodes,omitempty" xml:"specUpgradeOriginSpecCodes,omitempty" type:"Repeated"`
	SpecifyStartDate             *int64                  `json:"specifyStartDate,omitempty" xml:"specifyStartDate,omitempty"`
	UpgradeInquireFinancialValue *bool                   `json:"upgradeInquireFinancialValue,omitempty" xml:"upgradeInquireFinancialValue,omitempty"`
}

func (s CssInstanceCommodity) String() string {
	return dara.Prettify(s)
}

func (s CssInstanceCommodity) GoString() string {
	return s.String()
}

func (s *CssInstanceCommodity) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *CssInstanceCommodity) GetAliyunProduceCode() *string {
	return s.AliyunProduceCode
}

func (s *CssInstanceCommodity) GetChargeType() *string {
	return s.ChargeType
}

func (s *CssInstanceCommodity) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *CssInstanceCommodity) GetComponents() []*CssInstanceComponent {
	return s.Components
}

func (s *CssInstanceCommodity) GetDuration() *int64 {
	return s.Duration
}

func (s *CssInstanceCommodity) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CssInstanceCommodity) GetIsFree() *bool {
	return s.IsFree
}

func (s *CssInstanceCommodity) GetIsPrePayPostCharge() *bool {
	return s.IsPrePayPostCharge
}

func (s *CssInstanceCommodity) GetIsRenewChange() *bool {
	return s.IsRenewChange
}

func (s *CssInstanceCommodity) GetIsSyncToSubscription() *bool {
	return s.IsSyncToSubscription
}

func (s *CssInstanceCommodity) GetOrderParams() map[string]*string {
	return s.OrderParams
}

func (s *CssInstanceCommodity) GetOrderType() *string {
	return s.OrderType
}

func (s *CssInstanceCommodity) GetPlanItemId() *int64 {
	return s.PlanItemId
}

func (s *CssInstanceCommodity) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CssInstanceCommodity) GetQuantity() *int64 {
	return s.Quantity
}

func (s *CssInstanceCommodity) GetRedeemNoList() []*string {
	return s.RedeemNoList
}

func (s *CssInstanceCommodity) GetRedeemOrderType() *string {
	return s.RedeemOrderType
}

func (s *CssInstanceCommodity) GetRefundSpecCode() *string {
	return s.RefundSpecCode
}

func (s *CssInstanceCommodity) GetSpecCode() *string {
	return s.SpecCode
}

func (s *CssInstanceCommodity) GetSpecUpgradeOriginSpecCodes() []*string {
	return s.SpecUpgradeOriginSpecCodes
}

func (s *CssInstanceCommodity) GetSpecifyStartDate() *int64 {
	return s.SpecifyStartDate
}

func (s *CssInstanceCommodity) GetUpgradeInquireFinancialValue() *bool {
	return s.UpgradeInquireFinancialValue
}

func (s *CssInstanceCommodity) SetActivityId(v int64) *CssInstanceCommodity {
	s.ActivityId = &v
	return s
}

func (s *CssInstanceCommodity) SetAliyunProduceCode(v string) *CssInstanceCommodity {
	s.AliyunProduceCode = &v
	return s
}

func (s *CssInstanceCommodity) SetChargeType(v string) *CssInstanceCommodity {
	s.ChargeType = &v
	return s
}

func (s *CssInstanceCommodity) SetCommodityCode(v string) *CssInstanceCommodity {
	s.CommodityCode = &v
	return s
}

func (s *CssInstanceCommodity) SetComponents(v []*CssInstanceComponent) *CssInstanceCommodity {
	s.Components = v
	return s
}

func (s *CssInstanceCommodity) SetDuration(v int64) *CssInstanceCommodity {
	s.Duration = &v
	return s
}

func (s *CssInstanceCommodity) SetInstanceId(v string) *CssInstanceCommodity {
	s.InstanceId = &v
	return s
}

func (s *CssInstanceCommodity) SetIsFree(v bool) *CssInstanceCommodity {
	s.IsFree = &v
	return s
}

func (s *CssInstanceCommodity) SetIsPrePayPostCharge(v bool) *CssInstanceCommodity {
	s.IsPrePayPostCharge = &v
	return s
}

func (s *CssInstanceCommodity) SetIsRenewChange(v bool) *CssInstanceCommodity {
	s.IsRenewChange = &v
	return s
}

func (s *CssInstanceCommodity) SetIsSyncToSubscription(v bool) *CssInstanceCommodity {
	s.IsSyncToSubscription = &v
	return s
}

func (s *CssInstanceCommodity) SetOrderParams(v map[string]*string) *CssInstanceCommodity {
	s.OrderParams = v
	return s
}

func (s *CssInstanceCommodity) SetOrderType(v string) *CssInstanceCommodity {
	s.OrderType = &v
	return s
}

func (s *CssInstanceCommodity) SetPlanItemId(v int64) *CssInstanceCommodity {
	s.PlanItemId = &v
	return s
}

func (s *CssInstanceCommodity) SetPricingCycle(v string) *CssInstanceCommodity {
	s.PricingCycle = &v
	return s
}

func (s *CssInstanceCommodity) SetQuantity(v int64) *CssInstanceCommodity {
	s.Quantity = &v
	return s
}

func (s *CssInstanceCommodity) SetRedeemNoList(v []*string) *CssInstanceCommodity {
	s.RedeemNoList = v
	return s
}

func (s *CssInstanceCommodity) SetRedeemOrderType(v string) *CssInstanceCommodity {
	s.RedeemOrderType = &v
	return s
}

func (s *CssInstanceCommodity) SetRefundSpecCode(v string) *CssInstanceCommodity {
	s.RefundSpecCode = &v
	return s
}

func (s *CssInstanceCommodity) SetSpecCode(v string) *CssInstanceCommodity {
	s.SpecCode = &v
	return s
}

func (s *CssInstanceCommodity) SetSpecUpgradeOriginSpecCodes(v []*string) *CssInstanceCommodity {
	s.SpecUpgradeOriginSpecCodes = v
	return s
}

func (s *CssInstanceCommodity) SetSpecifyStartDate(v int64) *CssInstanceCommodity {
	s.SpecifyStartDate = &v
	return s
}

func (s *CssInstanceCommodity) SetUpgradeInquireFinancialValue(v bool) *CssInstanceCommodity {
	s.UpgradeInquireFinancialValue = &v
	return s
}

func (s *CssInstanceCommodity) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
