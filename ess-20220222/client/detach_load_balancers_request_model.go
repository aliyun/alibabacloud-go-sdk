// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachLoadBalancersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsync(v bool) *DetachLoadBalancersRequest
	GetAsync() *bool
	SetClientToken(v string) *DetachLoadBalancersRequest
	GetClientToken() *string
	SetForceDetach(v bool) *DetachLoadBalancersRequest
	GetForceDetach() *bool
	SetLoadBalancers(v []*string) *DetachLoadBalancersRequest
	GetLoadBalancers() []*string
	SetOwnerId(v int64) *DetachLoadBalancersRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DetachLoadBalancersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DetachLoadBalancersRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *DetachLoadBalancersRequest
	GetScalingGroupId() *string
}

type DetachLoadBalancersRequest struct {
	// Specifies whether to detach the CLB instance from the scaling group in an asynchronous manner. If you detach the CLB instance from the scaling group in an asynchronous manner, the call is successful only after all operations are successful. If a specific operation fails, the call fails. We recommend that you set this parameter to true.
	//
	// Valid values:
	//
	// 	- true: detaches the CLB instance from the scaling group in an asynchronous manner. In this case, the ID of the scaling activity is returned.
	//
	// 	- false: does not detach the CLB instance from the scaling group in an asynchronous manner.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to remove Elastic Compute Service (ECS) instances in the scaling group from the backend server groups of the load balancer. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceDetach *bool `json:"ForceDetach,omitempty" xml:"ForceDetach,omitempty"`
	// The IDs of the CLB instances. You can specify up to five instance IDs.
	//
	// This parameter is required.
	LoadBalancers []*string `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Repeated"`
	OwnerId       *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1ffogfdauy0jw0****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DetachLoadBalancersRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *DetachLoadBalancersRequest) GetAsync() *bool {
	return s.Async
}

func (s *DetachLoadBalancersRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetachLoadBalancersRequest) GetForceDetach() *bool {
	return s.ForceDetach
}

func (s *DetachLoadBalancersRequest) GetLoadBalancers() []*string {
	return s.LoadBalancers
}

func (s *DetachLoadBalancersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachLoadBalancersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachLoadBalancersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachLoadBalancersRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DetachLoadBalancersRequest) SetAsync(v bool) *DetachLoadBalancersRequest {
	s.Async = &v
	return s
}

func (s *DetachLoadBalancersRequest) SetClientToken(v string) *DetachLoadBalancersRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachLoadBalancersRequest) SetForceDetach(v bool) *DetachLoadBalancersRequest {
	s.ForceDetach = &v
	return s
}

func (s *DetachLoadBalancersRequest) SetLoadBalancers(v []*string) *DetachLoadBalancersRequest {
	s.LoadBalancers = v
	return s
}

func (s *DetachLoadBalancersRequest) SetOwnerId(v int64) *DetachLoadBalancersRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachLoadBalancersRequest) SetRegionId(v string) *DetachLoadBalancersRequest {
	s.RegionId = &v
	return s
}

func (s *DetachLoadBalancersRequest) SetResourceOwnerAccount(v string) *DetachLoadBalancersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachLoadBalancersRequest) SetScalingGroupId(v string) *DetachLoadBalancersRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DetachLoadBalancersRequest) Validate() error {
	return dara.Validate(s)
}
