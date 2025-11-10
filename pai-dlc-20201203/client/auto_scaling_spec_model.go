// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutoScalingSpec interface {
	dara.Model
	String() string
	GoString() string
	SetAutoscalingMetricSpec(v *AutoscalingMetricSpec) *AutoScalingSpec
	GetAutoscalingMetricSpec() *AutoscalingMetricSpec
	SetMaxReplicas(v int32) *AutoScalingSpec
	GetMaxReplicas() *int32
	SetMinReplicas(v int32) *AutoScalingSpec
	GetMinReplicas() *int32
	SetPodsToDelete(v []*string) *AutoScalingSpec
	GetPodsToDelete() []*string
	SetScalingStrategy(v string) *AutoScalingSpec
	GetScalingStrategy() *string
}

type AutoScalingSpec struct {
	AutoscalingMetricSpec *AutoscalingMetricSpec `json:"AutoscalingMetricSpec,omitempty" xml:"AutoscalingMetricSpec,omitempty"`
	MaxReplicas           *int32                 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	MinReplicas           *int32                 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	PodsToDelete          []*string              `json:"PodsToDelete,omitempty" xml:"PodsToDelete,omitempty" type:"Repeated"`
	ScalingStrategy       *string                `json:"ScalingStrategy,omitempty" xml:"ScalingStrategy,omitempty"`
}

func (s AutoScalingSpec) String() string {
	return dara.Prettify(s)
}

func (s AutoScalingSpec) GoString() string {
	return s.String()
}

func (s *AutoScalingSpec) GetAutoscalingMetricSpec() *AutoscalingMetricSpec {
	return s.AutoscalingMetricSpec
}

func (s *AutoScalingSpec) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *AutoScalingSpec) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *AutoScalingSpec) GetPodsToDelete() []*string {
	return s.PodsToDelete
}

func (s *AutoScalingSpec) GetScalingStrategy() *string {
	return s.ScalingStrategy
}

func (s *AutoScalingSpec) SetAutoscalingMetricSpec(v *AutoscalingMetricSpec) *AutoScalingSpec {
	s.AutoscalingMetricSpec = v
	return s
}

func (s *AutoScalingSpec) SetMaxReplicas(v int32) *AutoScalingSpec {
	s.MaxReplicas = &v
	return s
}

func (s *AutoScalingSpec) SetMinReplicas(v int32) *AutoScalingSpec {
	s.MinReplicas = &v
	return s
}

func (s *AutoScalingSpec) SetPodsToDelete(v []*string) *AutoScalingSpec {
	s.PodsToDelete = v
	return s
}

func (s *AutoScalingSpec) SetScalingStrategy(v string) *AutoScalingSpec {
	s.ScalingStrategy = &v
	return s
}

func (s *AutoScalingSpec) Validate() error {
	if s.AutoscalingMetricSpec != nil {
		if err := s.AutoscalingMetricSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}
