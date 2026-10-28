// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHyperNodeSchedulingConfig interface {
	dara.Model
	String() string
	GoString() string
	SetMinAvailable(v int32) *HyperNodeSchedulingConfig
	GetMinAvailable() *int32
	SetQualityPolicy(v string) *HyperNodeSchedulingConfig
	GetQualityPolicy() *string
}

type HyperNodeSchedulingConfig struct {
	// The minimum number of available instances required when the role participates in HyperNode scheduling. The value must be greater than 0.
	//
	// example:
	//
	// 1
	MinAvailable *int32 `json:"MinAvailable,omitempty" xml:"MinAvailable,omitempty"`
	// The quality policy. Currently, only AlwaysRequire is supported.
	//
	// example:
	//
	// AlwaysRequire
	QualityPolicy *string `json:"QualityPolicy,omitempty" xml:"QualityPolicy,omitempty"`
}

func (s HyperNodeSchedulingConfig) String() string {
	return dara.Prettify(s)
}

func (s HyperNodeSchedulingConfig) GoString() string {
	return s.String()
}

func (s *HyperNodeSchedulingConfig) GetMinAvailable() *int32 {
	return s.MinAvailable
}

func (s *HyperNodeSchedulingConfig) GetQualityPolicy() *string {
	return s.QualityPolicy
}

func (s *HyperNodeSchedulingConfig) SetMinAvailable(v int32) *HyperNodeSchedulingConfig {
	s.MinAvailable = &v
	return s
}

func (s *HyperNodeSchedulingConfig) SetQualityPolicy(v string) *HyperNodeSchedulingConfig {
	s.QualityPolicy = &v
	return s
}

func (s *HyperNodeSchedulingConfig) Validate() error {
	return dara.Validate(s)
}
