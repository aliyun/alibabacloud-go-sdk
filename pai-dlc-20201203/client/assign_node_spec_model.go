// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignNodeSpec interface {
	dara.Model
	String() string
	GoString() string
	SetAntiAffinityHyperNodes(v []*HyperNodeSpec) *AssignNodeSpec
	GetAntiAffinityHyperNodes() []*HyperNodeSpec
	SetAntiAffinityNodeNames(v string) *AssignNodeSpec
	GetAntiAffinityNodeNames() *string
	SetEnableAssignNode(v bool) *AssignNodeSpec
	GetEnableAssignNode() *bool
	SetHyperNodes(v []*HyperNodeSpec) *AssignNodeSpec
	GetHyperNodes() []*HyperNodeSpec
	SetNodeNames(v string) *AssignNodeSpec
	GetNodeNames() *string
}

type AssignNodeSpec struct {
	AntiAffinityHyperNodes []*HyperNodeSpec `json:"AntiAffinityHyperNodes,omitempty" xml:"AntiAffinityHyperNodes,omitempty" type:"Repeated"`
	// example:
	//
	// lingjxxxxxxxx
	AntiAffinityNodeNames *string `json:"AntiAffinityNodeNames,omitempty" xml:"AntiAffinityNodeNames,omitempty"`
	// example:
	//
	// true
	EnableAssignNode *bool            `json:"EnableAssignNode,omitempty" xml:"EnableAssignNode,omitempty"`
	HyperNodes       []*HyperNodeSpec `json:"HyperNodes,omitempty" xml:"HyperNodes,omitempty" type:"Repeated"`
	// example:
	//
	// lingjxxxxxxxx
	NodeNames *string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty"`
}

func (s AssignNodeSpec) String() string {
	return dara.Prettify(s)
}

func (s AssignNodeSpec) GoString() string {
	return s.String()
}

func (s *AssignNodeSpec) GetAntiAffinityHyperNodes() []*HyperNodeSpec {
	return s.AntiAffinityHyperNodes
}

func (s *AssignNodeSpec) GetAntiAffinityNodeNames() *string {
	return s.AntiAffinityNodeNames
}

func (s *AssignNodeSpec) GetEnableAssignNode() *bool {
	return s.EnableAssignNode
}

func (s *AssignNodeSpec) GetHyperNodes() []*HyperNodeSpec {
	return s.HyperNodes
}

func (s *AssignNodeSpec) GetNodeNames() *string {
	return s.NodeNames
}

func (s *AssignNodeSpec) SetAntiAffinityHyperNodes(v []*HyperNodeSpec) *AssignNodeSpec {
	s.AntiAffinityHyperNodes = v
	return s
}

func (s *AssignNodeSpec) SetAntiAffinityNodeNames(v string) *AssignNodeSpec {
	s.AntiAffinityNodeNames = &v
	return s
}

func (s *AssignNodeSpec) SetEnableAssignNode(v bool) *AssignNodeSpec {
	s.EnableAssignNode = &v
	return s
}

func (s *AssignNodeSpec) SetHyperNodes(v []*HyperNodeSpec) *AssignNodeSpec {
	s.HyperNodes = v
	return s
}

func (s *AssignNodeSpec) SetNodeNames(v string) *AssignNodeSpec {
	s.NodeNames = &v
	return s
}

func (s *AssignNodeSpec) Validate() error {
	if s.AntiAffinityHyperNodes != nil {
		for _, item := range s.AntiAffinityHyperNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HyperNodes != nil {
		for _, item := range s.HyperNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
