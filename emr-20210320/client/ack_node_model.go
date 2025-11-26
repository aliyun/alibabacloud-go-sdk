// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAckNode interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *AckNode
	GetNodeId() *string
	SetNodeSelector(v *AckNodeSelector) *AckNode
	GetNodeSelector() *AckNodeSelector
}

type AckNode struct {
	// 节点ID。
	//
	// example:
	//
	// i-bp1cudc25w2bfwl5****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// ACK节点选择器。
	NodeSelector *AckNodeSelector `json:"NodeSelector,omitempty" xml:"NodeSelector,omitempty"`
}

func (s AckNode) String() string {
	return dara.Prettify(s)
}

func (s AckNode) GoString() string {
	return s.String()
}

func (s *AckNode) GetNodeId() *string {
	return s.NodeId
}

func (s *AckNode) GetNodeSelector() *AckNodeSelector {
	return s.NodeSelector
}

func (s *AckNode) SetNodeId(v string) *AckNode {
	s.NodeId = &v
	return s
}

func (s *AckNode) SetNodeSelector(v *AckNodeSelector) *AckNode {
	s.NodeSelector = v
	return s
}

func (s *AckNode) Validate() error {
	if s.NodeSelector != nil {
		if err := s.NodeSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}
