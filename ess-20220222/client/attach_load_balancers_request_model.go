// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachLoadBalancersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsync(v bool) *AttachLoadBalancersRequest
	GetAsync() *bool
	SetClientToken(v string) *AttachLoadBalancersRequest
	GetClientToken() *string
	SetForceAttach(v bool) *AttachLoadBalancersRequest
	GetForceAttach() *bool
	SetLoadBalancerConfigs(v []*AttachLoadBalancersRequestLoadBalancerConfigs) *AttachLoadBalancersRequest
	GetLoadBalancerConfigs() []*AttachLoadBalancersRequestLoadBalancerConfigs
	SetLoadBalancers(v []*string) *AttachLoadBalancersRequest
	GetLoadBalancers() []*string
	SetOwnerId(v int64) *AttachLoadBalancersRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AttachLoadBalancersRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *AttachLoadBalancersRequest
	GetScalingGroupId() *string
}

type AttachLoadBalancersRequest struct {
	// Specifies whether to attach the CLB instance to the scaling group in an asynchronous manner. If you attach the CLB instance from the scaling group in an asynchronous manner, the call is successful only after all operations are successful. If a specific operation fails, the call fails. We recommend that you set this parameter to true. Valid values:
	//
	// 	- true: attaches the CLB instance to the scaling group in an asynchronous manner. In this case, the ID of the scaling activity is returned.
	//
	// 	- false: does not attach the CLB instance to the scaling group in an asynchronous manner.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to add the existing instances in the scaling group as backend servers of the load balancer. Valid values:
	//
	// 	- true: If you set this parameter to `true`, the attachment of the load balancer entails the addition of the existing instances in the scaling group to the backend server groups of the load balancer.
	//
	//     **
	//
	//     **Note*	- If a load balancer is currently attached to your scaling group, and you only want to add the instances in your scaling group to the backend server groups of the load balancer, you can call this operation again and set ForceAttach request to true.
	//
	// 	- false: If you set this parameter to false, the attachment of the load balancer does not entail the addition of the existing instances in the scaling group to the backend server groups of the load balancer.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceAttach *bool `json:"ForceAttach,omitempty" xml:"ForceAttach,omitempty"`
	// The configurations of the classic load balancer (CLB, formerly known as SLB) instance.
	LoadBalancerConfigs []*AttachLoadBalancersRequestLoadBalancerConfigs `json:"LoadBalancerConfigs,omitempty" xml:"LoadBalancerConfigs,omitempty" type:"Repeated"`
	// The IDs of the load balancers that you want to attach to the scaling group.
	LoadBalancers        []*string `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1avr6ensitts3w****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s AttachLoadBalancersRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *AttachLoadBalancersRequest) GetAsync() *bool {
	return s.Async
}

func (s *AttachLoadBalancersRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AttachLoadBalancersRequest) GetForceAttach() *bool {
	return s.ForceAttach
}

func (s *AttachLoadBalancersRequest) GetLoadBalancerConfigs() []*AttachLoadBalancersRequestLoadBalancerConfigs {
	return s.LoadBalancerConfigs
}

func (s *AttachLoadBalancersRequest) GetLoadBalancers() []*string {
	return s.LoadBalancers
}

func (s *AttachLoadBalancersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachLoadBalancersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachLoadBalancersRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *AttachLoadBalancersRequest) SetAsync(v bool) *AttachLoadBalancersRequest {
	s.Async = &v
	return s
}

func (s *AttachLoadBalancersRequest) SetClientToken(v string) *AttachLoadBalancersRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachLoadBalancersRequest) SetForceAttach(v bool) *AttachLoadBalancersRequest {
	s.ForceAttach = &v
	return s
}

func (s *AttachLoadBalancersRequest) SetLoadBalancerConfigs(v []*AttachLoadBalancersRequestLoadBalancerConfigs) *AttachLoadBalancersRequest {
	s.LoadBalancerConfigs = v
	return s
}

func (s *AttachLoadBalancersRequest) SetLoadBalancers(v []*string) *AttachLoadBalancersRequest {
	s.LoadBalancers = v
	return s
}

func (s *AttachLoadBalancersRequest) SetOwnerId(v int64) *AttachLoadBalancersRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachLoadBalancersRequest) SetResourceOwnerAccount(v string) *AttachLoadBalancersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachLoadBalancersRequest) SetScalingGroupId(v string) *AttachLoadBalancersRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *AttachLoadBalancersRequest) Validate() error {
	return dara.Validate(s)
}

type AttachLoadBalancersRequestLoadBalancerConfigs struct {
	// The ID of the CLB instance.
	//
	// example:
	//
	// 147b46d767c-cn-qingdao-cm5****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The weight of an Elastic Compute Service (ECS) instance or elastic container instance as a backend sever of the CLB instance. If an instance has a higher weight, more access traffic is routed to the instance. If an instance has zero weight, no access traffic is routed to the instance.
	//
	// Valid values: 0 to 100.
	//
	// example:
	//
	// 10
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AttachLoadBalancersRequestLoadBalancerConfigs) String() string {
	return dara.Prettify(s)
}

func (s AttachLoadBalancersRequestLoadBalancerConfigs) GoString() string {
	return s.String()
}

func (s *AttachLoadBalancersRequestLoadBalancerConfigs) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *AttachLoadBalancersRequestLoadBalancerConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *AttachLoadBalancersRequestLoadBalancerConfigs) SetLoadBalancerId(v string) *AttachLoadBalancersRequestLoadBalancerConfigs {
	s.LoadBalancerId = &v
	return s
}

func (s *AttachLoadBalancersRequestLoadBalancerConfigs) SetWeight(v int32) *AttachLoadBalancersRequestLoadBalancerConfigs {
	s.Weight = &v
	return s
}

func (s *AttachLoadBalancersRequestLoadBalancerConfigs) Validate() error {
	return dara.Validate(s)
}
