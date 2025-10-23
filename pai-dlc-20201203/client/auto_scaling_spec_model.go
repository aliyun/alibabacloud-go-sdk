// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutoScalingSpec interface {
	dara.Model
	String() string
	GoString() string
	SetMaxReplicas(v int32) *AutoScalingSpec
	GetMaxReplicas() *int32
	SetMinReplicas(v int32) *AutoScalingSpec
	GetMinReplicas() *int32
	SetScalingStrategy(v string) *AutoScalingSpec
	GetScalingStrategy() *string
}

type AutoScalingSpec struct {
	MaxReplicas     *int32  `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	MinReplicas     *int32  `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	ScalingStrategy *string `json:"ScalingStrategy,omitempty" xml:"ScalingStrategy,omitempty"`
}

func (s AutoScalingSpec) String() string {
	return dara.Prettify(s)
}

func (s AutoScalingSpec) GoString() string {
	return s.String()
}

func (s *AutoScalingSpec) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *AutoScalingSpec) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *AutoScalingSpec) GetScalingStrategy() *string {
	return s.ScalingStrategy
}

func (s *AutoScalingSpec) SetMaxReplicas(v int32) *AutoScalingSpec {
	s.MaxReplicas = &v
	return s
}

func (s *AutoScalingSpec) SetMinReplicas(v int32) *AutoScalingSpec {
	s.MinReplicas = &v
	return s
}

func (s *AutoScalingSpec) SetScalingStrategy(v string) *AutoScalingSpec {
	s.ScalingStrategy = &v
	return s
}

func (s *AutoScalingSpec) Validate() error {
	return dara.Validate(s)
}
