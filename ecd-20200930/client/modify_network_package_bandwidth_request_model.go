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
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The maximum bandwidth of the premium bandwidth plan, in Mbit/s.
	//
	// - For subscription premium bandwidth, the valid range is 2 to 1000.
	//
	// - For pay-as-you-go premium bandwidth with pay-by-traffic billing, the valid range is 2 to 200.
	//
	// - For pay-as-you-go premium bandwidth with pay-by-bandwidth billing, the valid range is 2 to 1000.
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
	// The region ID. You can call the [](t2167755.xdita#)operation to query the list of regions supported by Elastic Desktop Service (EDS).
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
