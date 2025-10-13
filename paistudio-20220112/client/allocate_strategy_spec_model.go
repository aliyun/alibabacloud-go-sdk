// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateStrategySpec interface {
	dara.Model
	String() string
	GoString() string
	SetNodeSpecs(v []*NodeSpec) *AllocateStrategySpec
	GetNodeSpecs() []*NodeSpec
}

type AllocateStrategySpec struct {
	NodeSpecs []*NodeSpec `json:"NodeSpecs,omitempty" xml:"NodeSpecs,omitempty" type:"Repeated"`
}

func (s AllocateStrategySpec) String() string {
	return dara.Prettify(s)
}

func (s AllocateStrategySpec) GoString() string {
	return s.String()
}

func (s *AllocateStrategySpec) GetNodeSpecs() []*NodeSpec {
	return s.NodeSpecs
}

func (s *AllocateStrategySpec) SetNodeSpecs(v []*NodeSpec) *AllocateStrategySpec {
	s.NodeSpecs = v
	return s
}

func (s *AllocateStrategySpec) Validate() error {
	if s.NodeSpecs != nil {
		for _, item := range s.NodeSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
