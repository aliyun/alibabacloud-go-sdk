// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeSpec interface {
	dara.Model
	String() string
	GoString() string
	SetBindingPolicy(v *BindingPolicy) *NodeSpec
	GetBindingPolicy() *BindingPolicy
	SetCount(v int64) *NodeSpec
	GetCount() *int64
	SetHyperType(v string) *NodeSpec
	GetHyperType() *string
	SetType(v string) *NodeSpec
	GetType() *string
}

type NodeSpec struct {
	// The resource node binding policy.
	BindingPolicy *BindingPolicy `json:"BindingPolicy,omitempty" xml:"BindingPolicy,omitempty"`
	// The number of resource node specifications.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The supernode specifications in use. When you specify a specification, you cannot specify both the node specifications and the supernode at the same time.
	//
	// example:
	//
	// efg3.GN9A.ch72
	HyperType *string `json:"HyperType,omitempty" xml:"HyperType,omitempty"`
	// The resource node specifications type.
	//
	// example:
	//
	// ecs.g6.4xlarge
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s NodeSpec) String() string {
	return dara.Prettify(s)
}

func (s NodeSpec) GoString() string {
	return s.String()
}

func (s *NodeSpec) GetBindingPolicy() *BindingPolicy {
	return s.BindingPolicy
}

func (s *NodeSpec) GetCount() *int64 {
	return s.Count
}

func (s *NodeSpec) GetHyperType() *string {
	return s.HyperType
}

func (s *NodeSpec) GetType() *string {
	return s.Type
}

func (s *NodeSpec) SetBindingPolicy(v *BindingPolicy) *NodeSpec {
	s.BindingPolicy = v
	return s
}

func (s *NodeSpec) SetCount(v int64) *NodeSpec {
	s.Count = &v
	return s
}

func (s *NodeSpec) SetHyperType(v string) *NodeSpec {
	s.HyperType = &v
	return s
}

func (s *NodeSpec) SetType(v string) *NodeSpec {
	s.Type = &v
	return s
}

func (s *NodeSpec) Validate() error {
	if s.BindingPolicy != nil {
		if err := s.BindingPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
