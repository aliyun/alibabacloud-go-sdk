// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyRCInstanceChargeTypeRequest
	GetAutoPay() *bool
	SetAutoRenew(v string) *ModifyRCInstanceChargeTypeRequest
	GetAutoRenew() *string
	SetAutoUseCoupon(v bool) *ModifyRCInstanceChargeTypeRequest
	GetAutoUseCoupon() *bool
	SetBusinessInfo(v string) *ModifyRCInstanceChargeTypeRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *ModifyRCInstanceChargeTypeRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyRCInstanceChargeTypeRequest
	GetDryRun() *bool
	SetIncludeDataDisks(v bool) *ModifyRCInstanceChargeTypeRequest
	GetIncludeDataDisks() *bool
	SetInstanceChargeType(v string) *ModifyRCInstanceChargeTypeRequest
	GetInstanceChargeType() *string
	SetInstanceId(v string) *ModifyRCInstanceChargeTypeRequest
	GetInstanceId() *string
	SetInstanceIds(v string) *ModifyRCInstanceChargeTypeRequest
	GetInstanceIds() *string
	SetPayType(v string) *ModifyRCInstanceChargeTypeRequest
	GetPayType() *string
	SetPeriod(v string) *ModifyRCInstanceChargeTypeRequest
	GetPeriod() *string
	SetPromotionCode(v string) *ModifyRCInstanceChargeTypeRequest
	GetPromotionCode() *string
	SetRegionId(v string) *ModifyRCInstanceChargeTypeRequest
	GetRegionId() *string
	SetUsedTime(v int32) *ModifyRCInstanceChargeTypeRequest
	GetUsedTime() *int32
}

type ModifyRCInstanceChargeTypeRequest struct {
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable the auto-renewal feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > 	- This parameter is valid only when you change the billing method from pay-as-you-go to subscription.
	//
	// > 	- All strings except **true*	- are considered **false**.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The additional business information about the instance.
	//
	// example:
	//
	// None
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The custom client token that is used to ensure the idempotence of the request.
	//
	// > The value can contain ASCII characters and can be up to 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOC****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	IncludeDataDisks *bool `json:"IncludeDataDisks,omitempty" xml:"IncludeDataDisks,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The ID of the instance or disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// rc-dh2jf9n6j4s14926****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The new billing method of the instance. Valid values:
	//
	// 	- **Prepaid**: subscription.
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The renewal cycle of the instance. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// > This parameter must be specified if you set the PayType parameter to **Prepaid**.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 72802442****
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The subscription duration of the instance.
	//
	// 	- If you set the **Period*	- parameter to **Year**, the value of the **UsedTime*	- parameter ranges from **1*	- to **5**.
	//
	// 	- If the **Period*	- parameter is set to **Month**, the value of the **UsedTime*	- parameter ranges from **1*	- to **11**.
	//
	// > If you set the **PayType*	- parameter to **Prepaid**, you must specify this parameter.
	//
	// example:
	//
	// 2
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s ModifyRCInstanceChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceChargeTypeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyRCInstanceChargeTypeRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *ModifyRCInstanceChargeTypeRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *ModifyRCInstanceChargeTypeRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *ModifyRCInstanceChargeTypeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyRCInstanceChargeTypeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyRCInstanceChargeTypeRequest) GetIncludeDataDisks() *bool {
	return s.IncludeDataDisks
}

func (s *ModifyRCInstanceChargeTypeRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ModifyRCInstanceChargeTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyRCInstanceChargeTypeRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ModifyRCInstanceChargeTypeRequest) GetPayType() *string {
	return s.PayType
}

func (s *ModifyRCInstanceChargeTypeRequest) GetPeriod() *string {
	return s.Period
}

func (s *ModifyRCInstanceChargeTypeRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyRCInstanceChargeTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCInstanceChargeTypeRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *ModifyRCInstanceChargeTypeRequest) SetAutoPay(v bool) *ModifyRCInstanceChargeTypeRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeRequest) SetAutoRenew(v string) *ModifyRCInstanceChargeTypeRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeRequest) SetAutoUseCoupon(v bool) *ModifyRCInstanceChargeTypeRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeRequest) SetBusinessInfo(v string) *ModifyRCInstanceChargeTypeRequest {
	s.BusinessInfo = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeRequest) SetClientToken(v string) *ModifyRCInstanceChargeTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeRequest) SetDryRun(v bool) *ModifyRCInstanceChargeTypeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeRequest) SetIncludeDataDisks(v bool) *ModifyRCInstanceChargeTypeRequest {
	s.IncludeDataDisks = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeRequest) SetInstanceChargeType(v string) *ModifyRCInstanceChargeTypeRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeRequest) SetInstanceId(v string) *ModifyRCInstanceChargeTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeRequest) SetInstanceIds(v string) *ModifyRCInstanceChargeTypeRequest {
	s.InstanceIds = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeRequest) SetPayType(v string) *ModifyRCInstanceChargeTypeRequest {
	s.PayType = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeRequest) SetPeriod(v string) *ModifyRCInstanceChargeTypeRequest {
	s.Period = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeRequest) SetPromotionCode(v string) *ModifyRCInstanceChargeTypeRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeRequest) SetRegionId(v string) *ModifyRCInstanceChargeTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeRequest) SetUsedTime(v int32) *ModifyRCInstanceChargeTypeRequest {
	s.UsedTime = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
