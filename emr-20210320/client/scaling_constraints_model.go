// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingConstraints interface {
	dara.Model
	String() string
	GoString() string
	SetMaxCapacity(v int32) *ScalingConstraints
	GetMaxCapacity() *int32
	SetMinCapacity(v int32) *ScalingConstraints
	GetMinCapacity() *int32
}

type ScalingConstraints struct {
	// 最大值。
	//
	// example:
	//
	// 2000
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// 最小值。
	//
	// example:
	//
	// 0
	MinCapacity *int32 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
}

func (s ScalingConstraints) String() string {
	return dara.Prettify(s)
}

func (s ScalingConstraints) GoString() string {
	return s.String()
}

func (s *ScalingConstraints) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *ScalingConstraints) GetMinCapacity() *int32 {
	return s.MinCapacity
}

func (s *ScalingConstraints) SetMaxCapacity(v int32) *ScalingConstraints {
	s.MaxCapacity = &v
	return s
}

func (s *ScalingConstraints) SetMinCapacity(v int32) *ScalingConstraints {
	s.MinCapacity = &v
	return s
}

func (s *ScalingConstraints) Validate() error {
	return dara.Validate(s)
}
