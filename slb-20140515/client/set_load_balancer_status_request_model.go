// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *SetLoadBalancerStatusRequest
	GetLoadBalancerId() *string
	SetLoadBalancerStatus(v string) *SetLoadBalancerStatusRequest
	GetLoadBalancerStatus() *string
	SetOwnerAccount(v string) *SetLoadBalancerStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetLoadBalancerStatusRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetLoadBalancerStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetLoadBalancerStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetLoadBalancerStatusRequest
	GetResourceOwnerId() *int64
}

type SetLoadBalancerStatusRequest struct {
	// The ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1b6c719dfa08e******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The state of the CLB instance. Valid values: **active*	- and **inactive**.
	//
	// 	- **active*	- (default)
	//
	//     If a CLB instance is in the **active*	- state, listeners of the CLB instance can forward traffic based on forwarding rules.
	//
	//     By default, newly created CLB instances are in the **active*	- state.
	//
	// 	- **inactive**
	//
	//     If a CLB instance is in the **inactive*	- state, listeners of the CLB instance do not forward traffic.
	//
	// >  If all listeners of a CLB instance are deleted, the CLB instance automatically switches to the **inactive*	- state.
	//
	// This parameter is required.
	//
	// example:
	//
	// active
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the CLB instance is deployed.
	//
	// You can query region IDs from the [Regions and zones](https://help.aliyun.com/document_detail/40654.html) list or by calling the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetLoadBalancerStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerStatusRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerStatusRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerStatusRequest) GetLoadBalancerStatus() *string {
	return s.LoadBalancerStatus
}

func (s *SetLoadBalancerStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetLoadBalancerStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLoadBalancerStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetLoadBalancerStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetLoadBalancerStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetLoadBalancerStatusRequest) SetLoadBalancerId(v string) *SetLoadBalancerStatusRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetLoadBalancerStatus(v string) *SetLoadBalancerStatusRequest {
	s.LoadBalancerStatus = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetOwnerAccount(v string) *SetLoadBalancerStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetOwnerId(v int64) *SetLoadBalancerStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetRegionId(v string) *SetLoadBalancerStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetResourceOwnerId(v int64) *SetLoadBalancerStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) Validate() error {
	return dara.Validate(s)
}
