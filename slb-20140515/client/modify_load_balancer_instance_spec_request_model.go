// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoadBalancerInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyLoadBalancerInstanceSpecRequest
	GetAutoPay() *bool
	SetLoadBalancerId(v string) *ModifyLoadBalancerInstanceSpecRequest
	GetLoadBalancerId() *string
	SetLoadBalancerSpec(v string) *ModifyLoadBalancerInstanceSpecRequest
	GetLoadBalancerSpec() *string
	SetOwnerAccount(v string) *ModifyLoadBalancerInstanceSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyLoadBalancerInstanceSpecRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyLoadBalancerInstanceSpecRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyLoadBalancerInstanceSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyLoadBalancerInstanceSpecRequest
	GetResourceOwnerId() *int64
}

type ModifyLoadBalancerInstanceSpecRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true**: automatically completes the payment.
	//
	// 	- **false*	- (default): If you select this option, you must complete the payment in the Order Center.
	//
	// > This parameter takes effect only for subscription instances.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1b6c719df*********
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The specification of the CLB instance. Valid values:
	//
	// 	- **slb.s1.small**
	//
	// 	- **slb.s2.small**
	//
	// 	- **slb.s2.medium**
	//
	// 	- **slb.s3.small**
	//
	// 	- **slb.s3.medium**
	//
	// 	- **slb.s3.large**
	//
	// The specifications available vary by region. For more information about the specifications, see [High-performance CLB instance](https://help.aliyun.com/document_detail/85931.html).
	//
	// > When you switch a shared-resource CLB instance to a high-performance CLB instance, your service may be interrupted for 10 to 30 seconds. We recommend that you modify the specification during off-peak hours or use Alibaba Cloud DNS to schedule your workloads to another CLB instance before you modify the specification.
	//
	// This parameter is required.
	//
	// example:
	//
	// slb.s2.small
	LoadBalancerSpec *string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the CLB instance.
	//
	// You can query the region ID from the [Regions and zones](https://help.aliyun.com/document_detail/40654.html) list or by calling the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyLoadBalancerInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoadBalancerInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInstanceSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyLoadBalancerInstanceSpecRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ModifyLoadBalancerInstanceSpecRequest) GetLoadBalancerSpec() *string {
	return s.LoadBalancerSpec
}

func (s *ModifyLoadBalancerInstanceSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyLoadBalancerInstanceSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLoadBalancerInstanceSpecRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLoadBalancerInstanceSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyLoadBalancerInstanceSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetAutoPay(v bool) *ModifyLoadBalancerInstanceSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetLoadBalancerId(v string) *ModifyLoadBalancerInstanceSpecRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetLoadBalancerSpec(v string) *ModifyLoadBalancerInstanceSpecRequest {
	s.LoadBalancerSpec = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetOwnerAccount(v string) *ModifyLoadBalancerInstanceSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetOwnerId(v int64) *ModifyLoadBalancerInstanceSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetRegionId(v string) *ModifyLoadBalancerInstanceSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetResourceOwnerAccount(v string) *ModifyLoadBalancerInstanceSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetResourceOwnerId(v int64) *ModifyLoadBalancerInstanceSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecRequest) Validate() error {
	return dara.Validate(s)
}
