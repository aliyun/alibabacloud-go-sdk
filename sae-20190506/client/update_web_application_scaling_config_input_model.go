// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebApplicationScalingConfigInput interface {
	dara.Model
	String() string
	GoString() string
	SetMaximumInstanceCount(v int64) *UpdateWebApplicationScalingConfigInput
	GetMaximumInstanceCount() *int64
	SetMinimumInstanceCount(v int64) *UpdateWebApplicationScalingConfigInput
	GetMinimumInstanceCount() *int64
}

type UpdateWebApplicationScalingConfigInput struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaximumInstanceCount *int64 `json:"MaximumInstanceCount,omitempty" xml:"MaximumInstanceCount,omitempty"`
	// This parameter is required.
	MinimumInstanceCount *int64 `json:"MinimumInstanceCount,omitempty" xml:"MinimumInstanceCount,omitempty"`
}

func (s UpdateWebApplicationScalingConfigInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebApplicationScalingConfigInput) GoString() string {
	return s.String()
}

func (s *UpdateWebApplicationScalingConfigInput) GetMaximumInstanceCount() *int64 {
	return s.MaximumInstanceCount
}

func (s *UpdateWebApplicationScalingConfigInput) GetMinimumInstanceCount() *int64 {
	return s.MinimumInstanceCount
}

func (s *UpdateWebApplicationScalingConfigInput) SetMaximumInstanceCount(v int64) *UpdateWebApplicationScalingConfigInput {
	s.MaximumInstanceCount = &v
	return s
}

func (s *UpdateWebApplicationScalingConfigInput) SetMinimumInstanceCount(v int64) *UpdateWebApplicationScalingConfigInput {
	s.MinimumInstanceCount = &v
	return s
}

func (s *UpdateWebApplicationScalingConfigInput) Validate() error {
	return dara.Validate(s)
}
