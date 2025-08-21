// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuggestShrinkableNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *GetSuggestShrinkableNodesRequest
	GetCount() *int32
	SetIgnoreStatus(v bool) *GetSuggestShrinkableNodesRequest
	GetIgnoreStatus() *bool
	SetNodeType(v string) *GetSuggestShrinkableNodesRequest
	GetNodeType() *string
}

type GetSuggestShrinkableNodesRequest struct {
	// The number of nodes that you want to remove.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// Specifies whether to ignore the instance status. Default value: false.
	//
	// example:
	//
	// false
	IgnoreStatus *bool `json:"ignoreStatus,omitempty" xml:"ignoreStatus,omitempty"`
	// The type of removing nodes. WORKER indicates hot node and WORKER_WARM indicates warm node.
	//
	// This parameter is required.
	//
	// example:
	//
	// WORKER
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
}

func (s GetSuggestShrinkableNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSuggestShrinkableNodesRequest) GoString() string {
	return s.String()
}

func (s *GetSuggestShrinkableNodesRequest) GetCount() *int32 {
	return s.Count
}

func (s *GetSuggestShrinkableNodesRequest) GetIgnoreStatus() *bool {
	return s.IgnoreStatus
}

func (s *GetSuggestShrinkableNodesRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *GetSuggestShrinkableNodesRequest) SetCount(v int32) *GetSuggestShrinkableNodesRequest {
	s.Count = &v
	return s
}

func (s *GetSuggestShrinkableNodesRequest) SetIgnoreStatus(v bool) *GetSuggestShrinkableNodesRequest {
	s.IgnoreStatus = &v
	return s
}

func (s *GetSuggestShrinkableNodesRequest) SetNodeType(v string) *GetSuggestShrinkableNodesRequest {
	s.NodeType = &v
	return s
}

func (s *GetSuggestShrinkableNodesRequest) Validate() error {
	return dara.Validate(s)
}
