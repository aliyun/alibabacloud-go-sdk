// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewRCInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewRCInstanceRequest
	GetAutoPay() *bool
	SetAutoRenew(v string) *RenewRCInstanceRequest
	GetAutoRenew() *string
	SetAutoUseCoupon(v bool) *RenewRCInstanceRequest
	GetAutoUseCoupon() *bool
	SetBusinessInfo(v string) *RenewRCInstanceRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *RenewRCInstanceRequest
	GetClientToken() *string
	SetCommodityCode(v string) *RenewRCInstanceRequest
	GetCommodityCode() *string
	SetInstanceId(v string) *RenewRCInstanceRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *RenewRCInstanceRequest
	GetOwnerId() *int64
	SetPayType(v string) *RenewRCInstanceRequest
	GetPayType() *string
	SetPeriodAlign(v bool) *RenewRCInstanceRequest
	GetPeriodAlign() *bool
	SetPromotionCode(v string) *RenewRCInstanceRequest
	GetPromotionCode() *string
	SetRegionId(v string) *RenewRCInstanceRequest
	GetRegionId() *string
	SetResource(v string) *RenewRCInstanceRequest
	GetResource() *string
	SetResourceOwnerAccount(v string) *RenewRCInstanceRequest
	GetResourceOwnerAccount() *string
	SetTimeType(v string) *RenewRCInstanceRequest
	GetTimeType() *string
	SetUsedTime(v string) *RenewRCInstanceRequest
	GetUsedTime() *string
}

type RenewRCInstanceRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **true**: Automatic payment is enabled. Make sure that your account balance is sufficient.
	//
	// - **false**: Only an order is generated. No payment is made.
	//
	//
	//
	//
	// > Default value: true. If your payment method has insufficient balance, set AutoPay to false. In this case, an unpaid order is generated. You can log on to the ApsaraDB RDS console to pay for the order.
	//
	// >
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**: Auto-renewal is enabled.
	//
	// 	- **false*	- (default): Auto-renewal is disabled.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to use coupons. Valid values:
	//
	// 	- **true*	- (default): Coupons are used.
	//
	// 	- **false**: Coupons are not used.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The additional information about the order.
	//
	// example:
	//
	// {\\"promotion_input_param\\":\\"{\\\\\\"promotionFilter\\\\\\":{},\\\\\\"promotionOptionCode\\\\\\":\\\\\\"youhui_quan\\\\\\"}\\"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOC****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The commodity code.
	//
	// <props="china">Default value: **rds_customprepaid_public_cn**.
	//
	//
	//
	// <props="intl">Default value: **rds_customprepaid_public_intl**.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds_customprepaid_public_**
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The ID of the RDS Custom instance.
	//
	// example:
	//
	// rc-dh2jf9n6j4s14926****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the target instance. Only **Prepaid*	- (upfront, subscription) is supported.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Specifies whether to use annual subscription. Valid values:
	//
	// - **true**: Annual subscription is used.
	//
	// - **false*	- (default): Annual subscription is not used.
	//
	// example:
	//
	// true
	PeriodAlign *bool `json:"PeriodAlign,omitempty" xml:"PeriodAlign,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 72329885****
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource.
	//
	// example:
	//
	// buy
	Resource             *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The unit of the renewal duration specified by the **UsedTime*	- parameter. Valid values:
	//
	// - **1**: year
	//
	// - **2*	- (default): month
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	TimeType *string `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The subscription duration. Valid values:
	//
	// 	- If **TimeType*	- is set to **1*	- (year), the valid values of UsedTime are **1 to 5**.
	//
	// 	- If **TimeType*	- is set to **2*	- (month), the valid values of UsedTime are **1 to 11**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s RenewRCInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewRCInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewRCInstanceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewRCInstanceRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *RenewRCInstanceRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *RenewRCInstanceRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *RenewRCInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewRCInstanceRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *RenewRCInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewRCInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RenewRCInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *RenewRCInstanceRequest) GetPeriodAlign() *bool {
	return s.PeriodAlign
}

func (s *RenewRCInstanceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *RenewRCInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RenewRCInstanceRequest) GetResource() *string {
	return s.Resource
}

func (s *RenewRCInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RenewRCInstanceRequest) GetTimeType() *string {
	return s.TimeType
}

func (s *RenewRCInstanceRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *RenewRCInstanceRequest) SetAutoPay(v bool) *RenewRCInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewRCInstanceRequest) SetAutoRenew(v string) *RenewRCInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewRCInstanceRequest) SetAutoUseCoupon(v bool) *RenewRCInstanceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *RenewRCInstanceRequest) SetBusinessInfo(v string) *RenewRCInstanceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *RenewRCInstanceRequest) SetClientToken(v string) *RenewRCInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewRCInstanceRequest) SetCommodityCode(v string) *RenewRCInstanceRequest {
	s.CommodityCode = &v
	return s
}

func (s *RenewRCInstanceRequest) SetInstanceId(v string) *RenewRCInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewRCInstanceRequest) SetOwnerId(v int64) *RenewRCInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewRCInstanceRequest) SetPayType(v string) *RenewRCInstanceRequest {
	s.PayType = &v
	return s
}

func (s *RenewRCInstanceRequest) SetPeriodAlign(v bool) *RenewRCInstanceRequest {
	s.PeriodAlign = &v
	return s
}

func (s *RenewRCInstanceRequest) SetPromotionCode(v string) *RenewRCInstanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *RenewRCInstanceRequest) SetRegionId(v string) *RenewRCInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RenewRCInstanceRequest) SetResource(v string) *RenewRCInstanceRequest {
	s.Resource = &v
	return s
}

func (s *RenewRCInstanceRequest) SetResourceOwnerAccount(v string) *RenewRCInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewRCInstanceRequest) SetTimeType(v string) *RenewRCInstanceRequest {
	s.TimeType = &v
	return s
}

func (s *RenewRCInstanceRequest) SetUsedTime(v string) *RenewRCInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *RenewRCInstanceRequest) Validate() error {
	return dara.Validate(s)
}
