// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebScalingConfig interface {
	dara.Model
	String() string
	GoString() string
	SetMaximumInstanceCount(v int64) *WebScalingConfig
	GetMaximumInstanceCount() *int64
	SetMinimumInstanceCount(v int64) *WebScalingConfig
	GetMinimumInstanceCount() *int64
}

type WebScalingConfig struct {
	// example:
	//
	// 10
	MaximumInstanceCount *int64 `json:"MaximumInstanceCount,omitempty" xml:"MaximumInstanceCount,omitempty"`
	MinimumInstanceCount *int64 `json:"MinimumInstanceCount,omitempty" xml:"MinimumInstanceCount,omitempty"`
}

func (s WebScalingConfig) String() string {
	return dara.Prettify(s)
}

func (s WebScalingConfig) GoString() string {
	return s.String()
}

func (s *WebScalingConfig) GetMaximumInstanceCount() *int64 {
	return s.MaximumInstanceCount
}

func (s *WebScalingConfig) GetMinimumInstanceCount() *int64 {
	return s.MinimumInstanceCount
}

func (s *WebScalingConfig) SetMaximumInstanceCount(v int64) *WebScalingConfig {
	s.MaximumInstanceCount = &v
	return s
}

func (s *WebScalingConfig) SetMinimumInstanceCount(v int64) *WebScalingConfig {
	s.MinimumInstanceCount = &v
	return s
}

func (s *WebScalingConfig) Validate() error {
	return dara.Validate(s)
}
