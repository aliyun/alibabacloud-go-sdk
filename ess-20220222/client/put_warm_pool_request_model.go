// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutWarmPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapacityOptions(v *PutWarmPoolRequestCapacityOptions) *PutWarmPoolRequest
	GetCapacityOptions() *PutWarmPoolRequestCapacityOptions
	SetClientToken(v string) *PutWarmPoolRequest
	GetClientToken() *string
	SetInstanceReusePolicy(v *PutWarmPoolRequestInstanceReusePolicy) *PutWarmPoolRequest
	GetInstanceReusePolicy() *PutWarmPoolRequestInstanceReusePolicy
	SetPoolState(v string) *PutWarmPoolRequest
	GetPoolState() *string
	SetRegionId(v string) *PutWarmPoolRequest
	GetRegionId() *string
	SetScalingGroupId(v string) *PutWarmPoolRequest
	GetScalingGroupId() *string
}

type PutWarmPoolRequest struct {
	// The capacity options for the warm pool.
	//
	// example:
	//
	// 0
	CapacityOptions *PutWarmPoolRequestCapacityOptions `json:"CapacityOptions,omitempty" xml:"CapacityOptions,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must ensure that it is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance reclaim policy.
	InstanceReusePolicy *PutWarmPoolRequestInstanceReusePolicy `json:"InstanceReusePolicy,omitempty" xml:"InstanceReusePolicy,omitempty" type:"Struct"`
	// The state of instances in the warm pool. Valid values:
	//
	// - Running: The instances are in the running state. This state allows faster scale-out, but you are charged for the full cost of the instances.
	//
	// - Stopped: The instances are in the economical mode. You are charged only for resources such as cloud disks, elastic IP addresses (EIPs), and snapshots of the instances. However, instance startup may fail during scale-out due to insufficient inventory.
	//
	// Default value: Stopped.
	//
	// example:
	//
	// Running
	PoolState *string `json:"PoolState,omitempty" xml:"PoolState,omitempty"`
	// The region ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1fo0dbtsbmqa9h****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s PutWarmPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s PutWarmPoolRequest) GoString() string {
	return s.String()
}

func (s *PutWarmPoolRequest) GetCapacityOptions() *PutWarmPoolRequestCapacityOptions {
	return s.CapacityOptions
}

func (s *PutWarmPoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PutWarmPoolRequest) GetInstanceReusePolicy() *PutWarmPoolRequestInstanceReusePolicy {
	return s.InstanceReusePolicy
}

func (s *PutWarmPoolRequest) GetPoolState() *string {
	return s.PoolState
}

func (s *PutWarmPoolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PutWarmPoolRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *PutWarmPoolRequest) SetCapacityOptions(v *PutWarmPoolRequestCapacityOptions) *PutWarmPoolRequest {
	s.CapacityOptions = v
	return s
}

func (s *PutWarmPoolRequest) SetClientToken(v string) *PutWarmPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *PutWarmPoolRequest) SetInstanceReusePolicy(v *PutWarmPoolRequestInstanceReusePolicy) *PutWarmPoolRequest {
	s.InstanceReusePolicy = v
	return s
}

func (s *PutWarmPoolRequest) SetPoolState(v string) *PutWarmPoolRequest {
	s.PoolState = &v
	return s
}

func (s *PutWarmPoolRequest) SetRegionId(v string) *PutWarmPoolRequest {
	s.RegionId = &v
	return s
}

func (s *PutWarmPoolRequest) SetScalingGroupId(v string) *PutWarmPoolRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *PutWarmPoolRequest) Validate() error {
	if s.CapacityOptions != nil {
		if err := s.CapacityOptions.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceReusePolicy != nil {
		if err := s.InstanceReusePolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutWarmPoolRequestCapacityOptions struct {
	// The expected number of instances to be warmed. The system maintains the sum of the current number of instances in the scaling group and the number of instances in the warm pool at this value. Instances that exceed the current number of instances in the scaling group are scaled out into the warm pool in advance. When traffic peaks occur, instances can be quickly added from the warm pool, avoiding the wait time caused by creating instances on demand.
	//
	// Default value: the maximum number of instances in the scaling group, which is specified by `MaxSize`.
	//
	// > The value of `DesiredGroupPreparedCapacity` must be less than or equal to the maximum number of instances in the scaling group, which is specified by `MaxSize`.
	//
	// example:
	//
	// 20
	DesiredGroupPreparedCapacity *int32 `json:"DesiredGroupPreparedCapacity,omitempty" xml:"DesiredGroupPreparedCapacity,omitempty"`
	// The minimum number of instances in the warm pool. The system always keeps at least this number of warmed instances in the warm pool to ensure that instances can be quickly added from the warm pool during scale-out.
	//
	// Default value: 0.
	//
	// > The value of `MinCapacity` must be less than or equal to the value of `DesiredGroupPreparedCapacity`.
	//
	// example:
	//
	// 0
	MinCapacity *int32 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
}

func (s PutWarmPoolRequestCapacityOptions) String() string {
	return dara.Prettify(s)
}

func (s PutWarmPoolRequestCapacityOptions) GoString() string {
	return s.String()
}

func (s *PutWarmPoolRequestCapacityOptions) GetDesiredGroupPreparedCapacity() *int32 {
	return s.DesiredGroupPreparedCapacity
}

func (s *PutWarmPoolRequestCapacityOptions) GetMinCapacity() *int32 {
	return s.MinCapacity
}

func (s *PutWarmPoolRequestCapacityOptions) SetDesiredGroupPreparedCapacity(v int32) *PutWarmPoolRequestCapacityOptions {
	s.DesiredGroupPreparedCapacity = &v
	return s
}

func (s *PutWarmPoolRequestCapacityOptions) SetMinCapacity(v int32) *PutWarmPoolRequestCapacityOptions {
	s.MinCapacity = &v
	return s
}

func (s *PutWarmPoolRequestCapacityOptions) Validate() error {
	return dara.Validate(s)
}

type PutWarmPoolRequestInstanceReusePolicy struct {
	// Specifies whether instances are reclaimed to the warm pool when the scaling group scales in. Valid values:
	//
	// - true: Instances are reclaimed to the warm pool.
	//
	// - false: Instances are not reclaimed to the warm pool.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	ReuseOnScaleIn *bool `json:"ReuseOnScaleIn,omitempty" xml:"ReuseOnScaleIn,omitempty"`
}

func (s PutWarmPoolRequestInstanceReusePolicy) String() string {
	return dara.Prettify(s)
}

func (s PutWarmPoolRequestInstanceReusePolicy) GoString() string {
	return s.String()
}

func (s *PutWarmPoolRequestInstanceReusePolicy) GetReuseOnScaleIn() *bool {
	return s.ReuseOnScaleIn
}

func (s *PutWarmPoolRequestInstanceReusePolicy) SetReuseOnScaleIn(v bool) *PutWarmPoolRequestInstanceReusePolicy {
	s.ReuseOnScaleIn = &v
	return s
}

func (s *PutWarmPoolRequestInstanceReusePolicy) Validate() error {
	return dara.Validate(s)
}
