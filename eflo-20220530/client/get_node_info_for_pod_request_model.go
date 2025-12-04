// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeInfoForPodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *GetNodeInfoForPodRequest
	GetNodeId() *string
	SetRegionId(v string) *GetNodeInfoForPodRequest
	GetRegionId() *string
}

type GetNodeInfoForPodRequest struct {
	// The ID of the node for this operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// node-be70****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetNodeInfoForPodRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeInfoForPodRequest) GoString() string {
	return s.String()
}

func (s *GetNodeInfoForPodRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNodeInfoForPodRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNodeInfoForPodRequest) SetNodeId(v string) *GetNodeInfoForPodRequest {
	s.NodeId = &v
	return s
}

func (s *GetNodeInfoForPodRequest) SetRegionId(v string) *GetNodeInfoForPodRequest {
	s.RegionId = &v
	return s
}

func (s *GetNodeInfoForPodRequest) Validate() error {
	return dara.Validate(s)
}
