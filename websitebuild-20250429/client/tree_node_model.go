// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTreeNode interface {
	dara.Model
	String() string
	GoString() string
	SetChildren(v []*TreeNode) *TreeNode
	GetChildren() []*TreeNode
	SetFinalStepNo(v int32) *TreeNode
	GetFinalStepNo() *int32
	SetFinishTime(v int64) *TreeNode
	GetFinishTime() *int64
	SetIsContainerNode(v bool) *TreeNode
	GetIsContainerNode() *bool
	SetNodeId(v string) *TreeNode
	GetNodeId() *string
	SetNodeName(v string) *TreeNode
	GetNodeName() *string
	SetNodeStatus(v string) *TreeNode
	GetNodeStatus() *string
	SetOperatorRole(v string) *TreeNode
	GetOperatorRole() *string
	SetParentNodeId(v string) *TreeNode
	GetParentNodeId() *string
	SetStepNo(v int32) *TreeNode
	GetStepNo() *int32
}

type TreeNode struct {
	Children        []*TreeNode `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	FinalStepNo     *int32      `json:"FinalStepNo,omitempty" xml:"FinalStepNo,omitempty"`
	FinishTime      *int64      `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	IsContainerNode *bool       `json:"IsContainerNode,omitempty" xml:"IsContainerNode,omitempty"`
	NodeId          *string     `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName        *string     `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeStatus      *string     `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	OperatorRole    *string     `json:"OperatorRole,omitempty" xml:"OperatorRole,omitempty"`
	ParentNodeId    *string     `json:"ParentNodeId,omitempty" xml:"ParentNodeId,omitempty"`
	StepNo          *int32      `json:"StepNo,omitempty" xml:"StepNo,omitempty"`
}

func (s TreeNode) String() string {
	return dara.Prettify(s)
}

func (s TreeNode) GoString() string {
	return s.String()
}

func (s *TreeNode) GetChildren() []*TreeNode {
	return s.Children
}

func (s *TreeNode) GetFinalStepNo() *int32 {
	return s.FinalStepNo
}

func (s *TreeNode) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *TreeNode) GetIsContainerNode() *bool {
	return s.IsContainerNode
}

func (s *TreeNode) GetNodeId() *string {
	return s.NodeId
}

func (s *TreeNode) GetNodeName() *string {
	return s.NodeName
}

func (s *TreeNode) GetNodeStatus() *string {
	return s.NodeStatus
}

func (s *TreeNode) GetOperatorRole() *string {
	return s.OperatorRole
}

func (s *TreeNode) GetParentNodeId() *string {
	return s.ParentNodeId
}

func (s *TreeNode) GetStepNo() *int32 {
	return s.StepNo
}

func (s *TreeNode) SetChildren(v []*TreeNode) *TreeNode {
	s.Children = v
	return s
}

func (s *TreeNode) SetFinalStepNo(v int32) *TreeNode {
	s.FinalStepNo = &v
	return s
}

func (s *TreeNode) SetFinishTime(v int64) *TreeNode {
	s.FinishTime = &v
	return s
}

func (s *TreeNode) SetIsContainerNode(v bool) *TreeNode {
	s.IsContainerNode = &v
	return s
}

func (s *TreeNode) SetNodeId(v string) *TreeNode {
	s.NodeId = &v
	return s
}

func (s *TreeNode) SetNodeName(v string) *TreeNode {
	s.NodeName = &v
	return s
}

func (s *TreeNode) SetNodeStatus(v string) *TreeNode {
	s.NodeStatus = &v
	return s
}

func (s *TreeNode) SetOperatorRole(v string) *TreeNode {
	s.OperatorRole = &v
	return s
}

func (s *TreeNode) SetParentNodeId(v string) *TreeNode {
	s.ParentNodeId = &v
	return s
}

func (s *TreeNode) SetStepNo(v int32) *TreeNode {
	s.StepNo = &v
	return s
}

func (s *TreeNode) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
