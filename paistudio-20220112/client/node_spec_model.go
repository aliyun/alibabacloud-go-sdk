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
	SetType(v string) *NodeSpec
	GetType() *string
}

type NodeSpec struct {
	BindingPolicy *BindingPolicy `json:"BindingPolicy,omitempty" xml:"BindingPolicy,omitempty"`
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
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

func (s *NodeSpec) SetType(v string) *NodeSpec {
	s.Type = &v
	return s
}

func (s *NodeSpec) Validate() error {
	return dara.Validate(s)
}
