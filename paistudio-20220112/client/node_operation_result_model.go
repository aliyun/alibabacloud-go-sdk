// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeOperationResult interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *NodeOperationResult
	GetMessage() *string
	SetNodeName(v string) *NodeOperationResult
	GetNodeName() *string
	SetStatus(v string) *NodeOperationResult
	GetStatus() *string
}

type NodeOperationResult struct {
	// The node operation status message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The node name.
	//
	// example:
	//
	// trains930928remn-master-0
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The node operation status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s NodeOperationResult) String() string {
	return dara.Prettify(s)
}

func (s NodeOperationResult) GoString() string {
	return s.String()
}

func (s *NodeOperationResult) GetMessage() *string {
	return s.Message
}

func (s *NodeOperationResult) GetNodeName() *string {
	return s.NodeName
}

func (s *NodeOperationResult) GetStatus() *string {
	return s.Status
}

func (s *NodeOperationResult) SetMessage(v string) *NodeOperationResult {
	s.Message = &v
	return s
}

func (s *NodeOperationResult) SetNodeName(v string) *NodeOperationResult {
	s.NodeName = &v
	return s
}

func (s *NodeOperationResult) SetStatus(v string) *NodeOperationResult {
	s.Status = &v
	return s
}

func (s *NodeOperationResult) Validate() error {
	return dara.Validate(s)
}
