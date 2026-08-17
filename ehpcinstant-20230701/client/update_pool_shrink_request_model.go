// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePoolShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPoolName(v string) *UpdatePoolShrinkRequest
	GetPoolName() *string
	SetPriority(v int32) *UpdatePoolShrinkRequest
	GetPriority() *int32
	SetResourceLimitsShrink(v string) *UpdatePoolShrinkRequest
	GetResourceLimitsShrink() *string
	SetSchedulingPolicyId(v string) *UpdatePoolShrinkRequest
	GetSchedulingPolicyId() *string
}

type UpdatePoolShrinkRequest struct {
	// The name of the resource pool.
	//
	// - The name can be up to 15 characters long.
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
	// - Valid values: 1 to 99. The default value is 1, which specifies the lowest priority.
	//
	// - Jobs in a higher-priority resource pool are scheduled before pending jobs in lower-priority pools. A resource pool\\"s priority overrides a job\\"s priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The limits on the resources that a user can use concurrently in the resource pool.
	ResourceLimitsShrink *string `json:"ResourceLimits,omitempty" xml:"ResourceLimits,omitempty"`
	// The ID of the scheduling policy.
	//
	// example:
	//
	// policy-xxxx
	SchedulingPolicyId *string `json:"SchedulingPolicyId,omitempty" xml:"SchedulingPolicyId,omitempty"`
}

func (s UpdatePoolShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePoolShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePoolShrinkRequest) GetPoolName() *string {
	return s.PoolName
}

func (s *UpdatePoolShrinkRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdatePoolShrinkRequest) GetResourceLimitsShrink() *string {
	return s.ResourceLimitsShrink
}

func (s *UpdatePoolShrinkRequest) GetSchedulingPolicyId() *string {
	return s.SchedulingPolicyId
}

func (s *UpdatePoolShrinkRequest) SetPoolName(v string) *UpdatePoolShrinkRequest {
	s.PoolName = &v
	return s
}

func (s *UpdatePoolShrinkRequest) SetPriority(v int32) *UpdatePoolShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdatePoolShrinkRequest) SetResourceLimitsShrink(v string) *UpdatePoolShrinkRequest {
	s.ResourceLimitsShrink = &v
	return s
}

func (s *UpdatePoolShrinkRequest) SetSchedulingPolicyId(v string) *UpdatePoolShrinkRequest {
	s.SchedulingPolicyId = &v
	return s
}

func (s *UpdatePoolShrinkRequest) Validate() error {
	return dara.Validate(s)
}
