// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkPackageBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyNetworkPackageBandwidthRequest
	GetAutoPay() *bool
	SetBandwidth(v int32) *ModifyNetworkPackageBandwidthRequest
	GetBandwidth() *int32
	SetNetworkPackageId(v string) *ModifyNetworkPackageBandwidthRequest
	GetNetworkPackageId() *string
	SetPromotionId(v string) *ModifyNetworkPackageBandwidthRequest
	GetPromotionId() *string
	SetRegionId(v string) *ModifyNetworkPackageBandwidthRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *ModifyNetworkPackageBandwidthRequest
	GetResellerOwnerUid() *int64
}

type ModifyNetworkPackageBandwidthRequest struct {
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
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The maximum bandwidth of the premium bandwidth plan, in Mbit/s. Valid range: The allowed range depends on the billing method:
	//
	// 	- Subscription: 2 to 1000
	//
	// 	- Pay-as-you-go, by data transfer (PayByTraffic): 2 to 200
	//
	// 	- Pay-as-you-go, by fixed bandwidth (PayByBandwidth): 2 to 1000
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the premium bandwidth plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// np-cxj99qb8d34vo****
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 500033080110596
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

func (s ModifyNetworkPackageBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkPackageBandwidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageBandwidthRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyNetworkPackageBandwidthRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ModifyNetworkPackageBandwidthRequest) GetNetworkPackageId() *string {
	return s.NetworkPackageId
}

func (s *ModifyNetworkPackageBandwidthRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *ModifyNetworkPackageBandwidthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNetworkPackageBandwidthRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *ModifyNetworkPackageBandwidthRequest) SetAutoPay(v bool) *ModifyNetworkPackageBandwidthRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyNetworkPackageBandwidthRequest) SetBandwidth(v int32) *ModifyNetworkPackageBandwidthRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyNetworkPackageBandwidthRequest) SetNetworkPackageId(v string) *ModifyNetworkPackageBandwidthRequest {
	s.NetworkPackageId = &v
	return s
}

func (s *ModifyNetworkPackageBandwidthRequest) SetPromotionId(v string) *ModifyNetworkPackageBandwidthRequest {
	s.PromotionId = &v
	return s
}

func (s *ModifyNetworkPackageBandwidthRequest) SetRegionId(v string) *ModifyNetworkPackageBandwidthRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNetworkPackageBandwidthRequest) SetResellerOwnerUid(v int64) *ModifyNetworkPackageBandwidthRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *ModifyNetworkPackageBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
