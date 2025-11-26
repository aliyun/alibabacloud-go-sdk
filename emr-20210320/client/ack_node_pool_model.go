// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAckNodePool interface {
	dara.Model
	String() string
	GoString() string
	SetNodePoolId(v string) *AckNodePool
	GetNodePoolId() *string
	SetNodeSelector(v *AckNodeSelector) *AckNodePool
	GetNodeSelector() *AckNodeSelector
}

type AckNodePool struct {
	// 节点池ID。
	//
	// example:
	//
	// npe76629caa1b14a73bf3e47c6d481****
	NodePoolId *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	// ACK节点选择器。
	NodeSelector *AckNodeSelector `json:"NodeSelector,omitempty" xml:"NodeSelector,omitempty"`
}

func (s AckNodePool) String() string {
	return dara.Prettify(s)
}

func (s AckNodePool) GoString() string {
	return s.String()
}

func (s *AckNodePool) GetNodePoolId() *string {
	return s.NodePoolId
}

func (s *AckNodePool) GetNodeSelector() *AckNodeSelector {
	return s.NodeSelector
}

func (s *AckNodePool) SetNodePoolId(v string) *AckNodePool {
	s.NodePoolId = &v
	return s
}

func (s *AckNodePool) SetNodeSelector(v *AckNodeSelector) *AckNodePool {
	s.NodeSelector = v
	return s
}

func (s *AckNodePool) Validate() error {
	if s.NodeSelector != nil {
		if err := s.NodeSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}
