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
	// This parameter is required.
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// example:
	//
	// 1
	Priority             *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
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
