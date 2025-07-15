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
}

type CreatePoolShrinkRequest struct {
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

func (s *CreatePoolShrinkRequest) Validate() error {
	return dara.Validate(s)
}
