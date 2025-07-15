// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewNetworkPackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewNetworkPackagesRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *RenewNetworkPackagesRequest
	GetAutoRenew() *bool
	SetNetworkPackageId(v []*string) *RenewNetworkPackagesRequest
	GetNetworkPackageId() []*string
	SetPeriod(v int32) *RenewNetworkPackagesRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewNetworkPackagesRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *RenewNetworkPackagesRequest
	GetPromotionId() *string
	SetRegionId(v string) *RenewNetworkPackagesRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *RenewNetworkPackagesRequest
	GetResellerOwnerUid() *int64
}

type RenewNetworkPackagesRequest struct {
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
	// true
	AutoPay   *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The IDs of premium bandwidth plans. You can specify up to 100 IDs.
	//
	// This parameter is required.
	NetworkPackageId []*string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" type:"Repeated"`
	// The subscription duration if you specify subscription as the new billing method for the cloud desktop. The unit of the value is specified by the `PeriodUnit` parameter. This parameter takes effect only when the `ChargeType` parameter is set to `PrePaid`.
	//
	// 	- If the `PeriodUnit` parameter is set to `Week`, the valid value of the Period parameter is 1.
	//
	// 	- If the `PeriodUnit` parameter is set to `Month`, the valid values of the Period parameter are 1, 2, 3, and 6.
	//
	// 	- If the `PeriodUnit` parameter is set to `Year`, the valid values of the Period parameter are 1, 2, 3, 4, and 5.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal duration specified by the Period parameter. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 500038360030606
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

func (s RenewNetworkPackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewNetworkPackagesRequest) GoString() string {
	return s.String()
}

func (s *RenewNetworkPackagesRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewNetworkPackagesRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewNetworkPackagesRequest) GetNetworkPackageId() []*string {
	return s.NetworkPackageId
}

func (s *RenewNetworkPackagesRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewNetworkPackagesRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewNetworkPackagesRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *RenewNetworkPackagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RenewNetworkPackagesRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *RenewNetworkPackagesRequest) SetAutoPay(v bool) *RenewNetworkPackagesRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetAutoRenew(v bool) *RenewNetworkPackagesRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetNetworkPackageId(v []*string) *RenewNetworkPackagesRequest {
	s.NetworkPackageId = v
	return s
}

func (s *RenewNetworkPackagesRequest) SetPeriod(v int32) *RenewNetworkPackagesRequest {
	s.Period = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetPeriodUnit(v string) *RenewNetworkPackagesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetPromotionId(v string) *RenewNetworkPackagesRequest {
	s.PromotionId = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetRegionId(v string) *RenewNetworkPackagesRequest {
	s.RegionId = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetResellerOwnerUid(v int64) *RenewNetworkPackagesRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *RenewNetworkPackagesRequest) Validate() error {
	return dara.Validate(s)
}
