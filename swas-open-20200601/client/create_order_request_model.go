// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessInfo(v *CreateOrderRequestBusinessInfo) *CreateOrderRequest
	GetBusinessInfo() *CreateOrderRequestBusinessInfo
	SetCommodity(v *CreateOrderRequestCommodity) *CreateOrderRequest
	GetCommodity() *CreateOrderRequestCommodity
	SetCouponNo(v string) *CreateOrderRequest
	GetCouponNo() *string
	SetFromApp(v string) *CreateOrderRequest
	GetFromApp() *string
	SetOrderType(v string) *CreateOrderRequest
	GetOrderType() *string
	SetRegionId(v string) *CreateOrderRequest
	GetRegionId() *string
}

type CreateOrderRequest struct {
	// The additional parameter for business information.
	BusinessInfo *CreateOrderRequestBusinessInfo `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty" type:"Struct"`
	// The order configuration information.
	//
	// This parameter is required.
	Commodity *CreateOrderRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The coupon code. Default value: `youhuiquan_promotion_option_id_for_blank`.
	//
	// example:
	//
	// 772339720881
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The request source. Valid values:
	//
	// - **OpenApi**: The request is sent from OpenAPI.
	//
	// - **mongo_buy**: The request is sent from the console.
	//
	// example:
	//
	// Buy
	FromApp *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	// The order type. Valid values:
	//
	// - Buy
	//
	// Default value: Buy.
	//
	// example:
	//
	// Buy
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) GetBusinessInfo() *CreateOrderRequestBusinessInfo {
	return s.BusinessInfo
}

func (s *CreateOrderRequest) GetCommodity() *CreateOrderRequestCommodity {
	return s.Commodity
}

func (s *CreateOrderRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *CreateOrderRequest) GetFromApp() *string {
	return s.FromApp
}

func (s *CreateOrderRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *CreateOrderRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOrderRequest) SetBusinessInfo(v *CreateOrderRequestBusinessInfo) *CreateOrderRequest {
	s.BusinessInfo = v
	return s
}

func (s *CreateOrderRequest) SetCommodity(v *CreateOrderRequestCommodity) *CreateOrderRequest {
	s.Commodity = v
	return s
}

func (s *CreateOrderRequest) SetCouponNo(v string) *CreateOrderRequest {
	s.CouponNo = &v
	return s
}

func (s *CreateOrderRequest) SetFromApp(v string) *CreateOrderRequest {
	s.FromApp = &v
	return s
}

func (s *CreateOrderRequest) SetOrderType(v string) *CreateOrderRequest {
	s.OrderType = &v
	return s
}

func (s *CreateOrderRequest) SetRegionId(v string) *CreateOrderRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOrderRequest) Validate() error {
	if s.BusinessInfo != nil {
		if err := s.BusinessInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Commodity != nil {
		if err := s.Commodity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrderRequestBusinessInfo struct {
	// The language environment. Default value: Chinese.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The order source.
	//
	// example:
	//
	// -
	OrderSource *string `json:"OrderSource,omitempty" xml:"OrderSource,omitempty"`
	// The promotion option information.
	PromotionOptions *CreateOrderRequestBusinessInfoPromotionOptions `json:"PromotionOptions,omitempty" xml:"PromotionOptions,omitempty" type:"Struct"`
	// Specifies whether to use a voucher. Valid values:
	//
	// - **true**: A voucher is used.
	//
	// - **false**: No voucher is used.
	//
	// Default value: **false**.
	//
	// example:
	//
	// false
	UseCoupon *bool `json:"UseCoupon,omitempty" xml:"UseCoupon,omitempty"`
	// The WeChat product type.
	WeChatProType *string `json:"WeChatProType,omitempty" xml:"WeChatProType,omitempty"`
	// The channel cookie value.
	//
	// example:
	//
	// IaEvKFpwjO1lMRzGWyJdT9KWob7ruIK%2F2zTcdPRkQ%2BIDp5GIJSQT%2FL3qUUcjJ3KQfud3NrMUlEqLaJWWIVwJgw%3D%3D
	LxChannelCookieValue *string `json:"lx_channel_cookie_value,omitempty" xml:"lx_channel_cookie_value,omitempty"`
}

func (s CreateOrderRequestBusinessInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequestBusinessInfo) GoString() string {
	return s.String()
}

func (s *CreateOrderRequestBusinessInfo) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *CreateOrderRequestBusinessInfo) GetOrderSource() *string {
	return s.OrderSource
}

func (s *CreateOrderRequestBusinessInfo) GetPromotionOptions() *CreateOrderRequestBusinessInfoPromotionOptions {
	return s.PromotionOptions
}

func (s *CreateOrderRequestBusinessInfo) GetUseCoupon() *bool {
	return s.UseCoupon
}

func (s *CreateOrderRequestBusinessInfo) GetWeChatProType() *string {
	return s.WeChatProType
}

func (s *CreateOrderRequestBusinessInfo) GetLxChannelCookieValue() *string {
	return s.LxChannelCookieValue
}

func (s *CreateOrderRequestBusinessInfo) SetAliyunLang(v string) *CreateOrderRequestBusinessInfo {
	s.AliyunLang = &v
	return s
}

func (s *CreateOrderRequestBusinessInfo) SetOrderSource(v string) *CreateOrderRequestBusinessInfo {
	s.OrderSource = &v
	return s
}

func (s *CreateOrderRequestBusinessInfo) SetPromotionOptions(v *CreateOrderRequestBusinessInfoPromotionOptions) *CreateOrderRequestBusinessInfo {
	s.PromotionOptions = v
	return s
}

func (s *CreateOrderRequestBusinessInfo) SetUseCoupon(v bool) *CreateOrderRequestBusinessInfo {
	s.UseCoupon = &v
	return s
}

func (s *CreateOrderRequestBusinessInfo) SetWeChatProType(v string) *CreateOrderRequestBusinessInfo {
	s.WeChatProType = &v
	return s
}

func (s *CreateOrderRequestBusinessInfo) SetLxChannelCookieValue(v string) *CreateOrderRequestBusinessInfo {
	s.LxChannelCookieValue = &v
	return s
}

func (s *CreateOrderRequestBusinessInfo) Validate() error {
	if s.PromotionOptions != nil {
		if err := s.PromotionOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrderRequestBusinessInfoPromotionOptions struct {
	// The activity ID.
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// The promotion filter.
	PromotionFilter *CreateOrderRequestBusinessInfoPromotionOptionsPromotionFilter `json:"PromotionFilter,omitempty" xml:"PromotionFilter,omitempty" type:"Struct"`
	// The promotion option code.
	//
	// example:
	//
	// youhui_quan
	PromotionOptionCode *string `json:"PromotionOptionCode,omitempty" xml:"PromotionOptionCode,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// 802681580959
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s CreateOrderRequestBusinessInfoPromotionOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequestBusinessInfoPromotionOptions) GoString() string {
	return s.String()
}

func (s *CreateOrderRequestBusinessInfoPromotionOptions) GetActivityId() *string {
	return s.ActivityId
}

func (s *CreateOrderRequestBusinessInfoPromotionOptions) GetPromotionFilter() *CreateOrderRequestBusinessInfoPromotionOptionsPromotionFilter {
	return s.PromotionFilter
}

func (s *CreateOrderRequestBusinessInfoPromotionOptions) GetPromotionOptionCode() *string {
	return s.PromotionOptionCode
}

func (s *CreateOrderRequestBusinessInfoPromotionOptions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *CreateOrderRequestBusinessInfoPromotionOptions) SetActivityId(v string) *CreateOrderRequestBusinessInfoPromotionOptions {
	s.ActivityId = &v
	return s
}

func (s *CreateOrderRequestBusinessInfoPromotionOptions) SetPromotionFilter(v *CreateOrderRequestBusinessInfoPromotionOptionsPromotionFilter) *CreateOrderRequestBusinessInfoPromotionOptions {
	s.PromotionFilter = v
	return s
}

func (s *CreateOrderRequestBusinessInfoPromotionOptions) SetPromotionOptionCode(v string) *CreateOrderRequestBusinessInfoPromotionOptions {
	s.PromotionOptionCode = &v
	return s
}

func (s *CreateOrderRequestBusinessInfoPromotionOptions) SetPromotionOptionNo(v string) *CreateOrderRequestBusinessInfoPromotionOptions {
	s.PromotionOptionNo = &v
	return s
}

func (s *CreateOrderRequestBusinessInfoPromotionOptions) Validate() error {
	if s.PromotionFilter != nil {
		if err := s.PromotionFilter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrderRequestBusinessInfoPromotionOptionsPromotionFilter struct {
	// Specifies whether to use a voucher. Valid values:
	//
	// - **true**: A voucher is used.
	//
	// - **false**: No voucher is used.
	//
	// Default value: **false**.
	//
	// example:
	//
	// False
	UseCoupon *bool `json:"UseCoupon,omitempty" xml:"UseCoupon,omitempty"`
}

func (s CreateOrderRequestBusinessInfoPromotionOptionsPromotionFilter) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequestBusinessInfoPromotionOptionsPromotionFilter) GoString() string {
	return s.String()
}

func (s *CreateOrderRequestBusinessInfoPromotionOptionsPromotionFilter) GetUseCoupon() *bool {
	return s.UseCoupon
}

func (s *CreateOrderRequestBusinessInfoPromotionOptionsPromotionFilter) SetUseCoupon(v bool) *CreateOrderRequestBusinessInfoPromotionOptionsPromotionFilter {
	s.UseCoupon = &v
	return s
}

func (s *CreateOrderRequestBusinessInfoPromotionOptionsPromotionFilter) Validate() error {
	return dara.Validate(s)
}

type CreateOrderRequestCommodity struct {
	// The number of Simple Application Server instances to create. Valid values: 1 to 20. Default value: 1.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether automatic payment is enabled for the order. Valid values:
	//
	// - **true**: Automatic payment.
	//
	// - **false**: No automatic payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal upon expiration. Valid values:
	//
	// - true: Auto-renewal is enabled.
	//
	// - false: Auto-renewal is disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period. Valid values:
	//
	// - If PeriodUnit is set to Year: 1, 2, and 3.
	//
	// - If PeriodUnit is set to Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The resource name.
	//
	// example:
	//
	// Server
	CommodityType *string `json:"CommodityType,omitempty" xml:"CommodityType,omitempty"`
	// The size of the mounted data cloud disk. Unit: GB. Valid values: 0 to 16380. The value must be a multiple of 20. Default value: 0.
	//
	// example:
	//
	// 0
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The image ID. You can call the [ListImages](https://help.aliyun.com/document_detail/189313.html) operation to query available image IDs in a specified region.
	//
	// example:
	//
	// 3e8eea0c00a04ff0b34a3ad8f1df455a
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The IDs of the Simple Application Server instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The billing method. Valid values:
	//
	// - Prepaid: subscription.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription period. Valid values: 1 to 20. Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period. Valid values:
	//
	// - Month: monthly billing.
	//
	// - Year: yearly billing.
	//
	// Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The plan ID. You can call the [ListPlans](https://help.aliyun.com/document_detail/189314.html) operation to query all plans in a specified region.
	//
	// example:
	//
	// swas.s.c2m1s30b1.linux
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s CreateOrderRequestCommodity) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequestCommodity) GoString() string {
	return s.String()
}

func (s *CreateOrderRequestCommodity) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateOrderRequestCommodity) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateOrderRequestCommodity) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateOrderRequestCommodity) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateOrderRequestCommodity) GetCommodityType() *string {
	return s.CommodityType
}

func (s *CreateOrderRequestCommodity) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *CreateOrderRequestCommodity) GetImageId() *string {
	return s.ImageId
}

func (s *CreateOrderRequestCommodity) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateOrderRequestCommodity) GetPayType() *string {
	return s.PayType
}

func (s *CreateOrderRequestCommodity) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateOrderRequestCommodity) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateOrderRequestCommodity) GetPlanId() *string {
	return s.PlanId
}

func (s *CreateOrderRequestCommodity) SetAmount(v int32) *CreateOrderRequestCommodity {
	s.Amount = &v
	return s
}

func (s *CreateOrderRequestCommodity) SetAutoPay(v bool) *CreateOrderRequestCommodity {
	s.AutoPay = &v
	return s
}

func (s *CreateOrderRequestCommodity) SetAutoRenew(v bool) *CreateOrderRequestCommodity {
	s.AutoRenew = &v
	return s
}

func (s *CreateOrderRequestCommodity) SetAutoRenewPeriod(v int32) *CreateOrderRequestCommodity {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateOrderRequestCommodity) SetCommodityType(v string) *CreateOrderRequestCommodity {
	s.CommodityType = &v
	return s
}

func (s *CreateOrderRequestCommodity) SetDataDiskSize(v int32) *CreateOrderRequestCommodity {
	s.DataDiskSize = &v
	return s
}

func (s *CreateOrderRequestCommodity) SetImageId(v string) *CreateOrderRequestCommodity {
	s.ImageId = &v
	return s
}

func (s *CreateOrderRequestCommodity) SetInstanceIds(v []*string) *CreateOrderRequestCommodity {
	s.InstanceIds = v
	return s
}

func (s *CreateOrderRequestCommodity) SetPayType(v string) *CreateOrderRequestCommodity {
	s.PayType = &v
	return s
}

func (s *CreateOrderRequestCommodity) SetPeriod(v int32) *CreateOrderRequestCommodity {
	s.Period = &v
	return s
}

func (s *CreateOrderRequestCommodity) SetPeriodUnit(v string) *CreateOrderRequestCommodity {
	s.PeriodUnit = &v
	return s
}

func (s *CreateOrderRequestCommodity) SetPlanId(v string) *CreateOrderRequestCommodity {
	s.PlanId = &v
	return s
}

func (s *CreateOrderRequestCommodity) Validate() error {
	return dara.Validate(s)
}
