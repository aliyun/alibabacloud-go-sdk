// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutoScalingSpec interface {
	dara.Model
	String() string
	GoString() string
	SetScalingStrategy(v string) *AutoScalingSpec
	GetScalingStrategy() *string
}

type AutoScalingSpec struct {
	ScalingStrategy *string `json:"ScalingStrategy,omitempty" xml:"ScalingStrategy,omitempty"`
}

func (s AutoScalingSpec) String() string {
	return dara.Prettify(s)
}

func (s AutoScalingSpec) GoString() string {
	return s.String()
}

func (s *AutoScalingSpec) GetScalingStrategy() *string {
	return s.ScalingStrategy
}

func (s *AutoScalingSpec) SetScalingStrategy(v string) *AutoScalingSpec {
	s.ScalingStrategy = &v
	return s
}

func (s *AutoScalingSpec) Validate() error {
	return dara.Validate(s)
}
