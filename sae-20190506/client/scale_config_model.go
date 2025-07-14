// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAlwaysAllocateCPU(v bool) *ScaleConfig
	GetAlwaysAllocateCPU() *bool
	SetMaximumInstanceCount(v int64) *ScaleConfig
	GetMaximumInstanceCount() *int64
	SetMinimumInstanceCount(v int64) *ScaleConfig
	GetMinimumInstanceCount() *int64
	SetRequestId(v string) *ScaleConfig
	GetRequestId() *string
}

type ScaleConfig struct {
	AlwaysAllocateCPU    *bool   `json:"alwaysAllocateCPU,omitempty" xml:"alwaysAllocateCPU,omitempty"`
	MaximumInstanceCount *int64  `json:"maximumInstanceCount,omitempty" xml:"maximumInstanceCount,omitempty"`
	MinimumInstanceCount *int64  `json:"minimumInstanceCount,omitempty" xml:"minimumInstanceCount,omitempty"`
	RequestId            *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ScaleConfig) String() string {
	return dara.Prettify(s)
}

func (s ScaleConfig) GoString() string {
	return s.String()
}

func (s *ScaleConfig) GetAlwaysAllocateCPU() *bool {
	return s.AlwaysAllocateCPU
}

func (s *ScaleConfig) GetMaximumInstanceCount() *int64 {
	return s.MaximumInstanceCount
}

func (s *ScaleConfig) GetMinimumInstanceCount() *int64 {
	return s.MinimumInstanceCount
}

func (s *ScaleConfig) GetRequestId() *string {
	return s.RequestId
}

func (s *ScaleConfig) SetAlwaysAllocateCPU(v bool) *ScaleConfig {
	s.AlwaysAllocateCPU = &v
	return s
}

func (s *ScaleConfig) SetMaximumInstanceCount(v int64) *ScaleConfig {
	s.MaximumInstanceCount = &v
	return s
}

func (s *ScaleConfig) SetMinimumInstanceCount(v int64) *ScaleConfig {
	s.MinimumInstanceCount = &v
	return s
}

func (s *ScaleConfig) SetRequestId(v string) *ScaleConfig {
	s.RequestId = &v
	return s
}

func (s *ScaleConfig) Validate() error {
	return dara.Validate(s)
}
