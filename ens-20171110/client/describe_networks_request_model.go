// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeNetworksRequest
	GetEnsRegionId() *string
	SetEnsRegionIds(v []*string) *DescribeNetworksRequest
	GetEnsRegionIds() []*string
	SetNetworkId(v string) *DescribeNetworksRequest
	GetNetworkId() *string
	SetNetworkIds(v []*string) *DescribeNetworksRequest
	GetNetworkIds() []*string
	SetNetworkName(v string) *DescribeNetworksRequest
	GetNetworkName() *string
	SetPageNumber(v int32) *DescribeNetworksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNetworksRequest
	GetPageSize() *int32
}

type DescribeNetworksRequest struct {
	// The ID of the edge node.
	//
	// example:
	//
	// cn-beijing-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The IDs of edge nodes. You can specify 1 to 100 IDs.
	EnsRegionIds []*string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty" type:"Repeated"`
	// The ID of the network.
	//
	// example:
	//
	// n-5***
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The IDs of VPCs You can specify 1 to 100 IDs.
	NetworkIds []*string `json:"NetworkIds,omitempty" xml:"NetworkIds,omitempty" type:"Repeated"`
	// The name of the network.
	//
	// example:
	//
	// example
	NetworkName *string `json:"NetworkName,omitempty" xml:"NetworkName,omitempty"`
	// The page number of the returned page. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeNetworksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworksRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworksRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeNetworksRequest) GetEnsRegionIds() []*string {
	return s.EnsRegionIds
}

func (s *DescribeNetworksRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeNetworksRequest) GetNetworkIds() []*string {
	return s.NetworkIds
}

func (s *DescribeNetworksRequest) GetNetworkName() *string {
	return s.NetworkName
}

func (s *DescribeNetworksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNetworksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNetworksRequest) SetEnsRegionId(v string) *DescribeNetworksRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNetworksRequest) SetEnsRegionIds(v []*string) *DescribeNetworksRequest {
	s.EnsRegionIds = v
	return s
}

func (s *DescribeNetworksRequest) SetNetworkId(v string) *DescribeNetworksRequest {
	s.NetworkId = &v
	return s
}

func (s *DescribeNetworksRequest) SetNetworkIds(v []*string) *DescribeNetworksRequest {
	s.NetworkIds = v
	return s
}

func (s *DescribeNetworksRequest) SetNetworkName(v string) *DescribeNetworksRequest {
	s.NetworkName = &v
	return s
}

func (s *DescribeNetworksRequest) SetPageNumber(v int32) *DescribeNetworksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworksRequest) SetPageSize(v int32) *DescribeNetworksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworksRequest) Validate() error {
	return dara.Validate(s)
}
