// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iComponentCapacityConstraint interface {
	dara.Model
	String() string
	GoString() string
	SetComponentType(v string) *ComponentCapacityConstraint
	GetComponentType() *string
	SetMaxCapacity(v int32) *ComponentCapacityConstraint
	GetMaxCapacity() *int32
	SetMinCapacity(v int32) *ComponentCapacityConstraint
	GetMinCapacity() *int32
}

type ComponentCapacityConstraint struct {
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	MaxCapacity   *int32  `json:"maxCapacity,omitempty" xml:"maxCapacity,omitempty"`
	MinCapacity   *int32  `json:"minCapacity,omitempty" xml:"minCapacity,omitempty"`
}

func (s ComponentCapacityConstraint) String() string {
	return dara.Prettify(s)
}

func (s ComponentCapacityConstraint) GoString() string {
	return s.String()
}

func (s *ComponentCapacityConstraint) GetComponentType() *string {
	return s.ComponentType
}

func (s *ComponentCapacityConstraint) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *ComponentCapacityConstraint) GetMinCapacity() *int32 {
	return s.MinCapacity
}

func (s *ComponentCapacityConstraint) SetComponentType(v string) *ComponentCapacityConstraint {
	s.ComponentType = &v
	return s
}

func (s *ComponentCapacityConstraint) SetMaxCapacity(v int32) *ComponentCapacityConstraint {
	s.MaxCapacity = &v
	return s
}

func (s *ComponentCapacityConstraint) SetMinCapacity(v int32) *ComponentCapacityConstraint {
	s.MinCapacity = &v
	return s
}

func (s *ComponentCapacityConstraint) Validate() error {
	return dara.Validate(s)
}
