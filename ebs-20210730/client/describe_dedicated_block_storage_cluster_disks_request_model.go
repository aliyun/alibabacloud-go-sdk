// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedBlockStorageClusterDisksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbscId(v string) *DescribeDedicatedBlockStorageClusterDisksRequest
	GetDbscId() *string
	SetMaxResults(v int64) *DescribeDedicatedBlockStorageClusterDisksRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeDedicatedBlockStorageClusterDisksRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeDedicatedBlockStorageClusterDisksRequest
	GetRegionId() *string
}

type DescribeDedicatedBlockStorageClusterDisksRequest struct {
	// The dedicated block storage cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbsc-cn-od43bf****
	DbscId *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	// The maximum number of entries per page for a paged query. Maximum value: 500.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous call. You do not need to set this parameter for the first request.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the dedicated block storage cluster. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDedicatedBlockStorageClusterDisksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) GetDbscId() *string {
	return s.DbscId
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) SetDbscId(v string) *DescribeDedicatedBlockStorageClusterDisksRequest {
	s.DbscId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) SetMaxResults(v int64) *DescribeDedicatedBlockStorageClusterDisksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) SetNextToken(v string) *DescribeDedicatedBlockStorageClusterDisksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) SetRegionId(v string) *DescribeDedicatedBlockStorageClusterDisksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) Validate() error {
	return dara.Validate(s)
}
