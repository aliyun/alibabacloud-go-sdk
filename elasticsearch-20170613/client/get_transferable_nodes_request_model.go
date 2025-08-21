// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransferableNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *GetTransferableNodesRequest
	GetCount() *int32
	SetNodeType(v string) *GetTransferableNodesRequest
	GetNodeType() *string
}

type GetTransferableNodesRequest struct {
	// The number of nodes to be migrated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The type of nodes.**WORKER**represents a hot node,**WORKER_WARM*	- represents a warm node.
	//
	// This parameter is required.
	//
	// example:
	//
	// WORKER
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
}

func (s GetTransferableNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTransferableNodesRequest) GoString() string {
	return s.String()
}

func (s *GetTransferableNodesRequest) GetCount() *int32 {
	return s.Count
}

func (s *GetTransferableNodesRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *GetTransferableNodesRequest) SetCount(v int32) *GetTransferableNodesRequest {
	s.Count = &v
	return s
}

func (s *GetTransferableNodesRequest) SetNodeType(v string) *GetTransferableNodesRequest {
	s.NodeType = &v
	return s
}

func (s *GetTransferableNodesRequest) Validate() error {
	return dara.Validate(s)
}
