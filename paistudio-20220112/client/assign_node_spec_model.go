// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignNodeSpec interface {
	dara.Model
	String() string
	GoString() string
	SetAntiAffinityNodeNames(v string) *AssignNodeSpec
	GetAntiAffinityNodeNames() *string
	SetEnableAssignNode(v bool) *AssignNodeSpec
	GetEnableAssignNode() *bool
	SetNodeNames(v string) *AssignNodeSpec
	GetNodeNames() *string
}

type AssignNodeSpec struct {
	AntiAffinityNodeNames *string `json:"AntiAffinityNodeNames,omitempty" xml:"AntiAffinityNodeNames,omitempty"`
	EnableAssignNode      *bool   `json:"EnableAssignNode,omitempty" xml:"EnableAssignNode,omitempty"`
	NodeNames             *string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty"`
}

func (s AssignNodeSpec) String() string {
	return dara.Prettify(s)
}

func (s AssignNodeSpec) GoString() string {
	return s.String()
}

func (s *AssignNodeSpec) GetAntiAffinityNodeNames() *string {
	return s.AntiAffinityNodeNames
}

func (s *AssignNodeSpec) GetEnableAssignNode() *bool {
	return s.EnableAssignNode
}

func (s *AssignNodeSpec) GetNodeNames() *string {
	return s.NodeNames
}

func (s *AssignNodeSpec) SetAntiAffinityNodeNames(v string) *AssignNodeSpec {
	s.AntiAffinityNodeNames = &v
	return s
}

func (s *AssignNodeSpec) SetEnableAssignNode(v bool) *AssignNodeSpec {
	s.EnableAssignNode = &v
	return s
}

func (s *AssignNodeSpec) SetNodeNames(v string) *AssignNodeSpec {
	s.NodeNames = &v
	return s
}

func (s *AssignNodeSpec) Validate() error {
	return dara.Validate(s)
}
