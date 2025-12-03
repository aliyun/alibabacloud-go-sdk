// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoadBalancerPayTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyLoadBalancerPayTypeRequest
	GetAutoPay() *bool
	SetDuration(v int32) *ModifyLoadBalancerPayTypeRequest
	GetDuration() *int32
	SetLoadBalancerId(v string) *ModifyLoadBalancerPayTypeRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *ModifyLoadBalancerPayTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyLoadBalancerPayTypeRequest
	GetOwnerId() *int64
	SetPayType(v string) *ModifyLoadBalancerPayTypeRequest
	GetPayType() *string
	SetPricingCycle(v string) *ModifyLoadBalancerPayTypeRequest
	GetPricingCycle() *string
	SetRegionId(v string) *ModifyLoadBalancerPayTypeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyLoadBalancerPayTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyLoadBalancerPayTypeRequest
	GetResourceOwnerId() *int64
}

type ModifyLoadBalancerPayTypeRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false*	- (default): no
	//
	// >  This parameter is valid only when the `PayType` parameter is set to **PrePay**. This parameter is valid only for pay-as-you-go instances.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The subscription duration.
	//
	// 	- If **PricingCycle*	- is set to **month**, the valid values are **1*	- to **9**.
	//
	// 	- If **PricingCycle*	- is set to **year**, the valid values are **1*	- to **3**.
	//
	// >  This parameter is valid only when the **PayType*	- parameter is set to **PrePay**. This parameter is valid only for pay-as-you-go instances.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1b6c719dfa08ex*****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the CLB instance. Valid values:
	//
	// 	- **PayOnDemand*	- (default): pay-as-you-go
	//
	// To change the billing method of a pay-as-you-go CLB instance to subscription, you must set the parameter to **PrePay**. In addition, the previous billing method of the CLB instance must be **PayOnDemand**.
	//
	// example:
	//
	// PrePay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing cycle.
	//
	// Valid values: **year*	- and **month**.
	//
	// >  This parameter is valid only when the **PayType*	- parameter is set to **PrePay**. This parameter is valid only for pay-as-you-go instances.
	//
	// example:
	//
	// month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The ID of the region where the CLB instance is deployed.
	//
	// You can query the region ID from the [Regions and zones](https://help.aliyun.com/document_detail/40654.html) list or by calling the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation.
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

func (s ModifyLoadBalancerPayTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoadBalancerPayTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerPayTypeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyLoadBalancerPayTypeRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *ModifyLoadBalancerPayTypeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ModifyLoadBalancerPayTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyLoadBalancerPayTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLoadBalancerPayTypeRequest) GetPayType() *string {
	return s.PayType
}

func (s *ModifyLoadBalancerPayTypeRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *ModifyLoadBalancerPayTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLoadBalancerPayTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyLoadBalancerPayTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyLoadBalancerPayTypeRequest) SetAutoPay(v bool) *ModifyLoadBalancerPayTypeRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetDuration(v int32) *ModifyLoadBalancerPayTypeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetLoadBalancerId(v string) *ModifyLoadBalancerPayTypeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetOwnerAccount(v string) *ModifyLoadBalancerPayTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetOwnerId(v int64) *ModifyLoadBalancerPayTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetPayType(v string) *ModifyLoadBalancerPayTypeRequest {
	s.PayType = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetPricingCycle(v string) *ModifyLoadBalancerPayTypeRequest {
	s.PricingCycle = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetRegionId(v string) *ModifyLoadBalancerPayTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetResourceOwnerAccount(v string) *ModifyLoadBalancerPayTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetResourceOwnerId(v int64) *ModifyLoadBalancerPayTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) Validate() error {
	return dara.Validate(s)
}
