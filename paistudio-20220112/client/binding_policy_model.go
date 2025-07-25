// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindingPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeNodes(v []*string) *BindingPolicy
	GetExcludeNodes() []*string
	SetIncludeNodes(v []*string) *BindingPolicy
	GetIncludeNodes() []*string
	SetNodeSpecCount(v int64) *BindingPolicy
	GetNodeSpecCount() *int64
}

type BindingPolicy struct {
	ExcludeNodes []*string `json:"ExcludeNodes,omitempty" xml:"ExcludeNodes,omitempty" type:"Repeated"`
	IncludeNodes []*string `json:"IncludeNodes,omitempty" xml:"IncludeNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	NodeSpecCount *int64 `json:"NodeSpecCount,omitempty" xml:"NodeSpecCount,omitempty"`
}

func (s BindingPolicy) String() string {
	return dara.Prettify(s)
}

func (s BindingPolicy) GoString() string {
	return s.String()
}

func (s *BindingPolicy) GetExcludeNodes() []*string {
	return s.ExcludeNodes
}

func (s *BindingPolicy) GetIncludeNodes() []*string {
	return s.IncludeNodes
}

func (s *BindingPolicy) GetNodeSpecCount() *int64 {
	return s.NodeSpecCount
}

func (s *BindingPolicy) SetExcludeNodes(v []*string) *BindingPolicy {
	s.ExcludeNodes = v
	return s
}

func (s *BindingPolicy) SetIncludeNodes(v []*string) *BindingPolicy {
	s.IncludeNodes = v
	return s
}

func (s *BindingPolicy) SetNodeSpecCount(v int64) *BindingPolicy {
	s.NodeSpecCount = &v
	return s
}

func (s *BindingPolicy) Validate() error {
	return dara.Validate(s)
}
