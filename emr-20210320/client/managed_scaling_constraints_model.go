// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManagedScalingConstraints interface {
	dara.Model
	String() string
	GoString() string
	SetMaxCapacity(v int32) *ManagedScalingConstraints
	GetMaxCapacity() *int32
	SetMaxOnDemandCapacity(v int32) *ManagedScalingConstraints
	GetMaxOnDemandCapacity() *int32
	SetMinCapacity(v int32) *ManagedScalingConstraints
	GetMinCapacity() *int32
}

type ManagedScalingConstraints struct {
	// 最大值。
	//
	// example:
	//
	// 2000
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// 最大按量节点数量
	//
	// example:
	//
	// 0
	MaxOnDemandCapacity *int32 `json:"MaxOnDemandCapacity,omitempty" xml:"MaxOnDemandCapacity,omitempty"`
	// 最小值。
	//
	// example:
	//
	// 0
	MinCapacity *int32 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
}

func (s ManagedScalingConstraints) String() string {
	return dara.Prettify(s)
}

func (s ManagedScalingConstraints) GoString() string {
	return s.String()
}

func (s *ManagedScalingConstraints) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *ManagedScalingConstraints) GetMaxOnDemandCapacity() *int32 {
	return s.MaxOnDemandCapacity
}

func (s *ManagedScalingConstraints) GetMinCapacity() *int32 {
	return s.MinCapacity
}

func (s *ManagedScalingConstraints) SetMaxCapacity(v int32) *ManagedScalingConstraints {
	s.MaxCapacity = &v
	return s
}

func (s *ManagedScalingConstraints) SetMaxOnDemandCapacity(v int32) *ManagedScalingConstraints {
	s.MaxOnDemandCapacity = &v
	return s
}

func (s *ManagedScalingConstraints) SetMinCapacity(v int32) *ManagedScalingConstraints {
	s.MinCapacity = &v
	return s
}

func (s *ManagedScalingConstraints) Validate() error {
	return dara.Validate(s)
}
