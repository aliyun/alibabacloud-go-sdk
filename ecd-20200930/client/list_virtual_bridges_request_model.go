// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualBridgesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBridgeId(v []*string) *ListVirtualBridgesRequest
	GetBridgeId() []*string
	SetMaxResults(v int32) *ListVirtualBridgesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVirtualBridgesRequest
	GetNextToken() *string
	SetOfficeSiteId(v string) *ListVirtualBridgesRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *ListVirtualBridgesRequest
	GetRegionId() *string
}

type ListVirtualBridgesRequest struct {
	// The list of virtual bridge IDs.
	BridgeId []*string `json:"BridgeId,omitempty" xml:"BridgeId,omitempty" type:"Repeated"`
	// The maximum number of entries to return. Valid values: 1 to 500.
	//
	// Default value: 500.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to start the next query. An empty NextToken indicates that no more results are available.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6mnFXZiT7NdvGNgkInJ****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The office network ID.
	//
	// > The `DirectoryId` parameter is deprecated. Use this parameter instead.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the list of regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListVirtualBridgesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualBridgesRequest) GoString() string {
	return s.String()
}

func (s *ListVirtualBridgesRequest) GetBridgeId() []*string {
	return s.BridgeId
}

func (s *ListVirtualBridgesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVirtualBridgesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVirtualBridgesRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ListVirtualBridgesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVirtualBridgesRequest) SetBridgeId(v []*string) *ListVirtualBridgesRequest {
	s.BridgeId = v
	return s
}

func (s *ListVirtualBridgesRequest) SetMaxResults(v int32) *ListVirtualBridgesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVirtualBridgesRequest) SetNextToken(v string) *ListVirtualBridgesRequest {
	s.NextToken = &v
	return s
}

func (s *ListVirtualBridgesRequest) SetOfficeSiteId(v string) *ListVirtualBridgesRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ListVirtualBridgesRequest) SetRegionId(v string) *ListVirtualBridgesRequest {
	s.RegionId = &v
	return s
}

func (s *ListVirtualBridgesRequest) Validate() error {
	return dara.Validate(s)
}
