// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewDesktopsRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *RenewDesktopsRequest
	GetAutoRenew() *bool
	SetDesktopId(v []*string) *RenewDesktopsRequest
	GetDesktopId() []*string
	SetPeriod(v int32) *RenewDesktopsRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewDesktopsRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *RenewDesktopsRequest
	GetPromotionId() *string
	SetRegionId(v string) *RenewDesktopsRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *RenewDesktopsRequest
	GetResellerOwnerUid() *int64
	SetResourceType(v string) *RenewDesktopsRequest
	GetResourceType() *string
}

type RenewDesktopsRequest struct {
	// Specifies whether to enable the auto-payment feature.
	//
	// Valid values:
	//
	// 	- true (default): enables the auto-payment feature. Make sure that your account balance is sufficient. Otherwise, an abnormal order is generated.
	//
	// 	- false: disables the auto-payment feature. In this case, an order is generated but you need to complete the payment. You can log on to the EDS console and complete the payment based on the order ID on the Orders page.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable the auto-renewal feature.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The cloud computer IDs. You can only renew monthly subscription cloud computers.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-ia2zw38bi6cm7****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The renewal duration. Valid values of this parameter are determined by the value of the `PeriodUnit` parameter.
	//
	// 	- Valid values if you set the `PeriodUnit` parameter to `Month`: 1, 2, 3, and 6
	//
	// 	- Valid values if you set the `PeriodUnit` parameter to `Year`: 1, 2, 3, 4, 5, and 6
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal duration specified by the `Period` parameter.
	//
	// Default value: Month. Valid values:
	//
	// 	- Month
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Year
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the promotional activity.
	//
	// example:
	//
	// 500030980150146
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions supported by Elastic Desktop Service (EDS).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// >  This field is not available for public use.
	//
	// example:
	//
	// null
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s RenewDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RenewDesktopsRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewDesktopsRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *RenewDesktopsRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewDesktopsRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewDesktopsRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *RenewDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RenewDesktopsRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *RenewDesktopsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *RenewDesktopsRequest) SetAutoPay(v bool) *RenewDesktopsRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewDesktopsRequest) SetAutoRenew(v bool) *RenewDesktopsRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewDesktopsRequest) SetDesktopId(v []*string) *RenewDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *RenewDesktopsRequest) SetPeriod(v int32) *RenewDesktopsRequest {
	s.Period = &v
	return s
}

func (s *RenewDesktopsRequest) SetPeriodUnit(v string) *RenewDesktopsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewDesktopsRequest) SetPromotionId(v string) *RenewDesktopsRequest {
	s.PromotionId = &v
	return s
}

func (s *RenewDesktopsRequest) SetRegionId(v string) *RenewDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RenewDesktopsRequest) SetResellerOwnerUid(v int64) *RenewDesktopsRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *RenewDesktopsRequest) SetResourceType(v string) *RenewDesktopsRequest {
	s.ResourceType = &v
	return s
}

func (s *RenewDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
