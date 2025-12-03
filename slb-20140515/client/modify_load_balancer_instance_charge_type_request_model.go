// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoadBalancerInstanceChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *ModifyLoadBalancerInstanceChargeTypeRequest
	GetBandwidth() *int32
	SetInstanceChargeType(v string) *ModifyLoadBalancerInstanceChargeTypeRequest
	GetInstanceChargeType() *string
	SetInternetChargeType(v string) *ModifyLoadBalancerInstanceChargeTypeRequest
	GetInternetChargeType() *string
	SetLoadBalancerId(v string) *ModifyLoadBalancerInstanceChargeTypeRequest
	GetLoadBalancerId() *string
	SetLoadBalancerSpec(v string) *ModifyLoadBalancerInstanceChargeTypeRequest
	GetLoadBalancerSpec() *string
	SetOwnerAccount(v string) *ModifyLoadBalancerInstanceChargeTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyLoadBalancerInstanceChargeTypeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyLoadBalancerInstanceChargeTypeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyLoadBalancerInstanceChargeTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyLoadBalancerInstanceChargeTypeRequest
	GetResourceOwnerId() *int64
}

type ModifyLoadBalancerInstanceChargeTypeRequest struct {
	// The maximum bandwidth of the Internet-facing CLB instance that is billed on a pay-by-bandwidth basis.
	//
	// You do not need to set this parameter. The metering method of Internet data transfer for pay-by-LCU instances supports only pay-by-traffic.
	//
	// example:
	//
	// 5
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The metering method of the instance after the change.
	//
	// Valid value: **PayByCLCU**. Only pay-by-LCU is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// PayByCLCU
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of Internet data transfer after the change.
	//
	// Valid value: **paybytraffic**.
	//
	// > 	- If the value of the **InstanceChargeType*	- parameter is set to **PayByCLCU**, only pay-by-data-transfer is supported.
	//
	// >	- When you change the metering method, the new metering method takes effect at 00:00:00 the next day.
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
	// lb-bp1b3jus5hpenznuu****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The specification of the CLB instance.
	//
	// You do not need to set this parameter. For pay-as-you-go CLB instances, you can only change the metering method from pay-by-specification to pay-by-LCU. You cannot change the metering method from pay-by-LCU to pay-by-specification.
	//
	// example:
	//
	// slb.s1.small
	LoadBalancerSpec *string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the CLB instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyLoadBalancerInstanceChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoadBalancerInstanceChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) GetLoadBalancerSpec() *string {
	return s.LoadBalancerSpec
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetBandwidth(v int32) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetInstanceChargeType(v string) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetInternetChargeType(v string) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.InternetChargeType = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetLoadBalancerId(v string) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetLoadBalancerSpec(v string) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.LoadBalancerSpec = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetOwnerAccount(v string) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetOwnerId(v int64) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetRegionId(v string) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetResourceOwnerAccount(v string) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetResourceOwnerId(v int64) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
