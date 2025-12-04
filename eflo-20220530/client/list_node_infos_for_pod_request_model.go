// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeInfosForPodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListNodeInfosForPodRequest
	GetClusterId() *string
	SetNodeId(v string) *ListNodeInfosForPodRequest
	GetNodeId() *string
	SetRegionId(v string) *ListNodeInfosForPodRequest
	GetRegionId() *string
	SetZoneId(v string) *ListNodeInfosForPodRequest
	GetZoneId() *string
}

type ListNodeInfosForPodRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// cluster-***
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the node for this operation.
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
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNodeInfosForPodRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodeInfosForPodRequest) GoString() string {
	return s.String()
}

func (s *ListNodeInfosForPodRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListNodeInfosForPodRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ListNodeInfosForPodRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNodeInfosForPodRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListNodeInfosForPodRequest) SetClusterId(v string) *ListNodeInfosForPodRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodeInfosForPodRequest) SetNodeId(v string) *ListNodeInfosForPodRequest {
	s.NodeId = &v
	return s
}

func (s *ListNodeInfosForPodRequest) SetRegionId(v string) *ListNodeInfosForPodRequest {
	s.RegionId = &v
	return s
}

func (s *ListNodeInfosForPodRequest) SetZoneId(v string) *ListNodeInfosForPodRequest {
	s.ZoneId = &v
	return s
}

func (s *ListNodeInfosForPodRequest) Validate() error {
	return dara.Validate(s)
}
