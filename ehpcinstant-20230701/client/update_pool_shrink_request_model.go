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
}

type UpdatePoolShrinkRequest struct {
	// The name of the resource pool.
	//
	// 	- The value can be up to 15 characters in length.
	//
	// 	- It can contain digits, uppercase letters, lowercase letters, underscores (_), and dots (.).
	//
	// This parameter is required.
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The priority of the resource pool.
	//
	// 	- You can set a priority in the range of 1 to 99. The default value is 1, which is the lowest priority.
	//
	// 	- Jobs submitted to a resource pool with a higher priority level value will be scheduled before pending jobs in a resource pool with a lower priority level value, and the priority level of the resource pool takes precedence over the priority of the job.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The quota of resources that users are allowed to concurrently use in a resource pool.
	ResourceLimitsShrink *string `json:"ResourceLimits,omitempty" xml:"ResourceLimits,omitempty"`
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

func (s *UpdatePoolShrinkRequest) Validate() error {
	return dara.Validate(s)
}
