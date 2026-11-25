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
	// Specifies whether to use asynchronous calls when adding load balancing instances. Asynchronous calls ensure the transactionality of operations, meaning that either all operations succeed or none of the operations take effect if any operation fails. Use asynchronous calls. Valid values:
	//
	// - true: asynchronous call. The request returns the ID of the scaling activity.
	//
	// - false: synchronous call.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// Ensures the idempotence of the request.
	//
	// Generate a parameter value from your client to ensure that the value is unique across different requests. The value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to add all existing instances in the scaling group as backend servers of the load balancing instances. Valid values:
	//
	//
	//
	// - true: adds all existing instances. If this parameter is set to `true`, when the load balancing instances are added to the scaling group, the system adds all existing instances in the scaling group as backend servers of the load balancing instances.
	//
	//   > If a load balancing instance is already added to the scaling group and you want to add all instances in the scaling group as backend servers of the load balancing instance, you can add the load balancing instance to the scaling group again and set ForceAttach to true.
	//
	// - false: does not add existing instances. If this parameter is set to false, when the load balancing instances are added to the scaling group, the system does not add the existing instances in the scaling group as backend servers of the load balancing instances.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceAttach *bool `json:"ForceAttach,omitempty" xml:"ForceAttach,omitempty"`
	// The configurations of the load balancing CLB (formerly SLB) instances.
	LoadBalancerConfigs []*AttachLoadBalancersRequestLoadBalancerConfigs `json:"LoadBalancerConfigs,omitempty" xml:"LoadBalancerConfigs,omitempty" type:"Repeated"`
	// The IDs of the load balancing CLB (formerly SLB) instances to add.
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
	if s.LoadBalancerConfigs != nil {
		for _, item := range s.LoadBalancerConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachLoadBalancersRequestLoadBalancerConfigs struct {
	// The ID of the load balancing CLB (formerly SLB) instance.
	//
	// example:
	//
	// 147b46d767c-cn-qingdao-cm5****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The weight of the ECS instance or ECI instance as a backend server after Auto Scaling adds the instance to the SLB server group. A higher weight indicates that the ECS instance or ECI instance is allocated more access requests. If the weight is 0, the ECS instance or ECI instance does not receive access requests.
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
