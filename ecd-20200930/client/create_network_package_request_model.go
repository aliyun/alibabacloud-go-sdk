// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateNetworkPackageRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateNetworkPackageRequest
	GetAutoRenew() *bool
	SetBandwidth(v int32) *CreateNetworkPackageRequest
	GetBandwidth() *int32
	SetInternetChargeType(v string) *CreateNetworkPackageRequest
	GetInternetChargeType() *string
	SetOfficeSiteId(v string) *CreateNetworkPackageRequest
	GetOfficeSiteId() *string
	SetPayType(v string) *CreateNetworkPackageRequest
	GetPayType() *string
	SetPeriod(v int32) *CreateNetworkPackageRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateNetworkPackageRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *CreateNetworkPackageRequest
	GetPromotionId() *string
	SetRegionId(v string) *CreateNetworkPackageRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *CreateNetworkPackageRequest
	GetResellerOwnerUid() *int64
}

type CreateNetworkPackageRequest struct {
	// Specifies whether to enable the automatic payment feature.
	//
	// Valid values:
	//
	// 	- true (default): enables the auto-payment feature.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     Make sure that your account has sufficient balance. Otherwise, no order is generated.
	//
	//     <!-- -->
	//
	// 	- false: disables the auto-payment feature. In this case, an order is generated but you need to make the payment manually.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     To make the payment, log on to the Elastic Desktop Service console, go to the Orders page, and find the order based on the order ID.
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the premium bandwidth plan.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The bandwidth provided by the premium bandwidth plan. Unit: Mbit/s.
	//
	// 	- Valid values if the premium bandwidth plan is a subscription plan: 2 to 1000.
	//
	// 	- Valid values if the premium bandwidth plan is a pay-as-you-go plan that charges by data transfer (PayByTraffic): 2 to 200.
	//
	// 	- Valid values if the premium bandwidth plan is a pay-as-you-go plan that charges by fixed bandwidth (PayByBandwidth): 2 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The charge type of the premium bandwidth plan.
	//
	// 	- Valid value when the `PayType` parameter is set to `PrePaid`:
	//
	//     	- PayByBandwidth: charges by fixed bandwidth.
	//
	// 	- Valid values when the `PayType` parameter is set to `PostPaid`:
	//
	//     	- PayByTraffic: charges by data transfer.
	//
	//     	- PayByBandwidth: charges by fixed bandwidth.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The billing method of the premium bandwidth plan.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go
	//
	// 	- PrePaid: subscription
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription duration of the premium bandwidth plan. This parameter takes effect and is required only when the `PayType` parameter is set to `PrePaid`. The valid values of this parameter vary based on the `PeriodUnit` value.
	//
	// 	- Valid value when the `PeriodUnit` parameter is set to `Week`: 1
	//
	// 	- Valid values when the `PeriodUnit` parameter is set to `Month`: 1, 2, 3, and 6
	//
	// 	- Valid values when the `PeriodUnit` parameter is set to `Year`: 1, 2, and 3
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration of the premium bandwidth plan. This parameter takes effect and is required only when the `PayType` parameter is set to `PrePaid`.
	//
	// Valid values:
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
	// 	- Week
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
	// The ID of the sales promotion.
	//
	// example:
	//
	// 23141
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
}

func (s CreateNetworkPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkPackageRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateNetworkPackageRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateNetworkPackageRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateNetworkPackageRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateNetworkPackageRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateNetworkPackageRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateNetworkPackageRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateNetworkPackageRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateNetworkPackageRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateNetworkPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNetworkPackageRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *CreateNetworkPackageRequest) SetAutoPay(v bool) *CreateNetworkPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetAutoRenew(v bool) *CreateNetworkPackageRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetBandwidth(v int32) *CreateNetworkPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetInternetChargeType(v string) *CreateNetworkPackageRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetOfficeSiteId(v string) *CreateNetworkPackageRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetPayType(v string) *CreateNetworkPackageRequest {
	s.PayType = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetPeriod(v int32) *CreateNetworkPackageRequest {
	s.Period = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetPeriodUnit(v string) *CreateNetworkPackageRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetPromotionId(v string) *CreateNetworkPackageRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetRegionId(v string) *CreateNetworkPackageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetResellerOwnerUid(v int64) *CreateNetworkPackageRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *CreateNetworkPackageRequest) Validate() error {
	return dara.Validate(s)
}
