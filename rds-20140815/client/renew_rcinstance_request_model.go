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
	// 	- **true**: enables the feature. You must make sure that your account balance is sufficient.
	//
	// 	- **false**: disables the feature. An unpaid order is generated.
	//
	// >  Default value: true. If your account balance is insufficient, you can set AutoPay to false to generate an unpaid order. Then, you can log on to the ApsaraDB RDS console to complete the payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to use a coupon. Default value: false. Valid values:
	//
	// 	- **true**: uses a coupon.
	//
	// 	- **false**: does not use a coupon.
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOC****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The commodity code of the instance.
	//
	// Default value: **rds_customprepaid_public_intl**.
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
	// The billing method of the instance. Set the value to **PrePaid**, which indicates the subscription billing method.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Specifies whether the instance is a subscription instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
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
	// The resources.
	//
	// example:
	//
	// buy
	Resource             *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The unit of the renewal period specified by the **UsedTime*	- parameter. Valid values:
	//
	// 	- **1**: year
	//
	// 	- **2*	- (default): month
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	TimeType *string `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The subscription duration of the instance. Valid values:
	//
	// 	- If you set the **TimeType*	- parameter to **1**, the value of the UsedTime parameter ranges from **1 to 5**. Unit: year.
	//
	// 	- If you set the **TimeType*	- parameter to **2**, the value of the UsedTime parameter ranges from **1 to 11**. Unit: month.
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
