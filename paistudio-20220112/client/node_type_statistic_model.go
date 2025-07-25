// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeTypeStatistic interface {
	dara.Model
	String() string
	GoString() string
	SetCanBeBoundCount(v int32) *NodeTypeStatistic
	GetCanBeBoundCount() *int32
	SetNodeType(v string) *NodeTypeStatistic
	GetNodeType() *string
	SetTotalCount(v int32) *NodeTypeStatistic
	GetTotalCount() *int32
}

type NodeTypeStatistic struct {
	// example:
	//
	// 4
	CanBeBoundCount *int32 `json:"CanBeBoundCount,omitempty" xml:"CanBeBoundCount,omitempty"`
	// example:
	//
	// ecs.g6.4xlarge
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s NodeTypeStatistic) String() string {
	return dara.Prettify(s)
}

func (s NodeTypeStatistic) GoString() string {
	return s.String()
}

func (s *NodeTypeStatistic) GetCanBeBoundCount() *int32 {
	return s.CanBeBoundCount
}

func (s *NodeTypeStatistic) GetNodeType() *string {
	return s.NodeType
}

func (s *NodeTypeStatistic) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *NodeTypeStatistic) SetCanBeBoundCount(v int32) *NodeTypeStatistic {
	s.CanBeBoundCount = &v
	return s
}

func (s *NodeTypeStatistic) SetNodeType(v string) *NodeTypeStatistic {
	s.NodeType = &v
	return s
}

func (s *NodeTypeStatistic) SetTotalCount(v int32) *NodeTypeStatistic {
	s.TotalCount = &v
	return s
}

func (s *NodeTypeStatistic) Validate() error {
	return dara.Validate(s)
}
