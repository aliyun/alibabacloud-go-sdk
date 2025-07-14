// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationScaleConfigInput interface {
	dara.Model
	String() string
	GoString() string
	SetAlwaysAllocateCPU(v bool) *UpdateApplicationScaleConfigInput
	GetAlwaysAllocateCPU() *bool
	SetMaximumInstanceCount(v int64) *UpdateApplicationScaleConfigInput
	GetMaximumInstanceCount() *int64
	SetMinimumInstanceCount(v int64) *UpdateApplicationScaleConfigInput
	GetMinimumInstanceCount() *int64
}

type UpdateApplicationScaleConfigInput struct {
	AlwaysAllocateCPU    *bool  `json:"alwaysAllocateCPU,omitempty" xml:"alwaysAllocateCPU,omitempty"`
	MaximumInstanceCount *int64 `json:"maximumInstanceCount,omitempty" xml:"maximumInstanceCount,omitempty"`
	MinimumInstanceCount *int64 `json:"minimumInstanceCount,omitempty" xml:"minimumInstanceCount,omitempty"`
}

func (s UpdateApplicationScaleConfigInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScaleConfigInput) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScaleConfigInput) GetAlwaysAllocateCPU() *bool {
	return s.AlwaysAllocateCPU
}

func (s *UpdateApplicationScaleConfigInput) GetMaximumInstanceCount() *int64 {
	return s.MaximumInstanceCount
}

func (s *UpdateApplicationScaleConfigInput) GetMinimumInstanceCount() *int64 {
	return s.MinimumInstanceCount
}

func (s *UpdateApplicationScaleConfigInput) SetAlwaysAllocateCPU(v bool) *UpdateApplicationScaleConfigInput {
	s.AlwaysAllocateCPU = &v
	return s
}

func (s *UpdateApplicationScaleConfigInput) SetMaximumInstanceCount(v int64) *UpdateApplicationScaleConfigInput {
	s.MaximumInstanceCount = &v
	return s
}

func (s *UpdateApplicationScaleConfigInput) SetMinimumInstanceCount(v int64) *UpdateApplicationScaleConfigInput {
	s.MinimumInstanceCount = &v
	return s
}

func (s *UpdateApplicationScaleConfigInput) Validate() error {
	return dara.Validate(s)
}
