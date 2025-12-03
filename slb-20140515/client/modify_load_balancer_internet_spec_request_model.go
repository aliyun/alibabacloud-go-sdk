// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoadBalancerInternetSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyLoadBalancerInternetSpecRequest
	GetAutoPay() *bool
	SetBandwidth(v int32) *ModifyLoadBalancerInternetSpecRequest
	GetBandwidth() *int32
	SetInternetChargeType(v string) *ModifyLoadBalancerInternetSpecRequest
	GetInternetChargeType() *string
	SetLoadBalancerId(v string) *ModifyLoadBalancerInternetSpecRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *ModifyLoadBalancerInternetSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyLoadBalancerInternetSpecRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyLoadBalancerInternetSpecRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyLoadBalancerInternetSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyLoadBalancerInternetSpecRequest
	GetResourceOwnerId() *int64
}

type ModifyLoadBalancerInternetSpecRequest struct {
	// Specifies whether to automatically pay the subscription fee of the Internet-facing CLB instance. Valid values:
	//
	// 	- **true**: enables automatic payments. This is the default value.
	//
	// 	- **false**: disables automatic payment. You must complete the payment in Order Center.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The maximum bandwidth of the Internet-facing CLB instance that uses the pay-by-bandwidth metering method. Unit: Mbit/s.
	//
	// Valid values: **1 to 5000**. The maximum bandwidth varies based on the region where the CLB instance is created.****
	//
	// >  You do not need to specify this parameter if you set **InternetChargeType*	- to **paybytraffic*	- (pay-by-data-transfer).
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The metering method of the Internet-facing CLB instance. Valid values:
	//
	// 	- **paybybandwidth**: pay-by-bandwidth
	//
	// 	- **paybytraffic**: pay-by-data-transfer
	//
	// example:
	//
	// paybytraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1b6c719dfa08ex******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the CLB instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyLoadBalancerInternetSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoadBalancerInternetSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInternetSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyLoadBalancerInternetSpecRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ModifyLoadBalancerInternetSpecRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *ModifyLoadBalancerInternetSpecRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ModifyLoadBalancerInternetSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyLoadBalancerInternetSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLoadBalancerInternetSpecRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLoadBalancerInternetSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyLoadBalancerInternetSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetAutoPay(v bool) *ModifyLoadBalancerInternetSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetBandwidth(v int32) *ModifyLoadBalancerInternetSpecRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetInternetChargeType(v string) *ModifyLoadBalancerInternetSpecRequest {
	s.InternetChargeType = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetLoadBalancerId(v string) *ModifyLoadBalancerInternetSpecRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetOwnerAccount(v string) *ModifyLoadBalancerInternetSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetOwnerId(v int64) *ModifyLoadBalancerInternetSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetRegionId(v string) *ModifyLoadBalancerInternetSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetResourceOwnerAccount(v string) *ModifyLoadBalancerInternetSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetResourceOwnerId(v int64) *ModifyLoadBalancerInternetSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) Validate() error {
	return dara.Validate(s)
}
