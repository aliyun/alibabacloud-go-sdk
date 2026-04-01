// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkPeerConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeNetworkPeerConnectionsRequest
	GetEnsRegionId() *string
	SetInstanceId(v string) *DescribeNetworkPeerConnectionsRequest
	GetInstanceId() *string
	SetName(v string) *DescribeNetworkPeerConnectionsRequest
	GetName() *string
	SetNetworkIds(v []*string) *DescribeNetworkPeerConnectionsRequest
	GetNetworkIds() []*string
	SetPageNumber(v int32) *DescribeNetworkPeerConnectionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNetworkPeerConnectionsRequest
	GetPageSize() *int32
}

type DescribeNetworkPeerConnectionsRequest struct {
	// example:
	//
	// cn-xian-unicom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// example:
	//
	// pcc-5inkeimcipxk26yqtzm4q****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test0
	Name       *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkIds []*string `json:"NetworkIds,omitempty" xml:"NetworkIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeNetworkPeerConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkPeerConnectionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPeerConnectionsRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeNetworkPeerConnectionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkPeerConnectionsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeNetworkPeerConnectionsRequest) GetNetworkIds() []*string {
	return s.NetworkIds
}

func (s *DescribeNetworkPeerConnectionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNetworkPeerConnectionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNetworkPeerConnectionsRequest) SetEnsRegionId(v string) *DescribeNetworkPeerConnectionsRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsRequest) SetInstanceId(v string) *DescribeNetworkPeerConnectionsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsRequest) SetName(v string) *DescribeNetworkPeerConnectionsRequest {
	s.Name = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsRequest) SetNetworkIds(v []*string) *DescribeNetworkPeerConnectionsRequest {
	s.NetworkIds = v
	return s
}

func (s *DescribeNetworkPeerConnectionsRequest) SetPageNumber(v int32) *DescribeNetworkPeerConnectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsRequest) SetPageSize(v int32) *DescribeNetworkPeerConnectionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsRequest) Validate() error {
	return dara.Validate(s)
}
