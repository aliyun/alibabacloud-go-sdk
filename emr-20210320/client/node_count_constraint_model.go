// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeCountConstraint interface {
	dara.Model
	String() string
	GoString() string
	SetMax(v int32) *NodeCountConstraint
	GetMax() *int32
	SetMin(v int32) *NodeCountConstraint
	GetMin() *int32
	SetType(v string) *NodeCountConstraint
	GetType() *string
	SetValues(v []*int32) *NodeCountConstraint
	GetValues() []*int32
}

type NodeCountConstraint struct {
	// example:
	//
	// 100
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// example:
	//
	// 1
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// range
	Type   *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Values []*int32 `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s NodeCountConstraint) String() string {
	return dara.Prettify(s)
}

func (s NodeCountConstraint) GoString() string {
	return s.String()
}

func (s *NodeCountConstraint) GetMax() *int32 {
	return s.Max
}

func (s *NodeCountConstraint) GetMin() *int32 {
	return s.Min
}

func (s *NodeCountConstraint) GetType() *string {
	return s.Type
}

func (s *NodeCountConstraint) GetValues() []*int32 {
	return s.Values
}

func (s *NodeCountConstraint) SetMax(v int32) *NodeCountConstraint {
	s.Max = &v
	return s
}

func (s *NodeCountConstraint) SetMin(v int32) *NodeCountConstraint {
	s.Min = &v
	return s
}

func (s *NodeCountConstraint) SetType(v string) *NodeCountConstraint {
	s.Type = &v
	return s
}

func (s *NodeCountConstraint) SetValues(v []*int32) *NodeCountConstraint {
	s.Values = v
	return s
}

func (s *NodeCountConstraint) Validate() error {
	return dara.Validate(s)
}
