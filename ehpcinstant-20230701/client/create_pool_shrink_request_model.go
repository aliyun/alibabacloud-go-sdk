// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePoolShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPoolName(v string) *CreatePoolShrinkRequest
	GetPoolName() *string
	SetPriority(v int32) *CreatePoolShrinkRequest
	GetPriority() *int32
	SetResourceLimitsShrink(v string) *CreatePoolShrinkRequest
	GetResourceLimitsShrink() *string
	SetSchedulingPolicyId(v string) *CreatePoolShrinkRequest
	GetSchedulingPolicyId() *string
}

type CreatePoolShrinkRequest struct {
	// The resource pool name.
	//
	// - The name can be up to 15 characters in length.
	//
	// - The name can contain digits, uppercase letters, lowercase letters, underscores (_), and periods (.).
	//
	// This parameter is required.
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The priority of the resource pool.
	//
	// - Valid values: 1 to 99. Default value: 1, which indicates the lowest priority.
	//
	// - Jobs submitted to a resource pool with a higher priority value are scheduled before pending jobs in a resource pool with a lower priority value. The priority of a resource pool takes precedence over the priority of a job.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The resource quota limits for concurrent usage allowed for a user within a resource pool.
	ResourceLimitsShrink *string `json:"ResourceLimits,omitempty" xml:"ResourceLimits,omitempty"`
	// The scheduling policy.
	//
	// example:
	//
	// policy-xxx
	SchedulingPolicyId *string `json:"SchedulingPolicyId,omitempty" xml:"SchedulingPolicyId,omitempty"`
}

func (s CreatePoolShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePoolShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePoolShrinkRequest) GetPoolName() *string {
	return s.PoolName
}

func (s *CreatePoolShrinkRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreatePoolShrinkRequest) GetResourceLimitsShrink() *string {
	return s.ResourceLimitsShrink
}

func (s *CreatePoolShrinkRequest) GetSchedulingPolicyId() *string {
	return s.SchedulingPolicyId
}

func (s *CreatePoolShrinkRequest) SetPoolName(v string) *CreatePoolShrinkRequest {
	s.PoolName = &v
	return s
}

func (s *CreatePoolShrinkRequest) SetPriority(v int32) *CreatePoolShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreatePoolShrinkRequest) SetResourceLimitsShrink(v string) *CreatePoolShrinkRequest {
	s.ResourceLimitsShrink = &v
	return s
}

func (s *CreatePoolShrinkRequest) SetSchedulingPolicyId(v string) *CreatePoolShrinkRequest {
	s.SchedulingPolicyId = &v
	return s
}

func (s *CreatePoolShrinkRequest) Validate() error {
	return dara.Validate(s)
}
