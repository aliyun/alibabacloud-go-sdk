// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReplicaGroupDrillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrillId(v string) *DescribeReplicaGroupDrillsRequest
	GetDrillId() *string
	SetGroupId(v string) *DescribeReplicaGroupDrillsRequest
	GetGroupId() *string
	SetMaxResults(v int32) *DescribeReplicaGroupDrillsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeReplicaGroupDrillsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeReplicaGroupDrillsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeReplicaGroupDrillsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeReplicaGroupDrillsRequest
	GetRegionId() *string
}

type DescribeReplicaGroupDrillsRequest struct {
	// The ID of the drill.
	//
	// example:
	//
	// pg-drill-xxxx
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The ID of the replication pair-consistent group. You can call the [DescribeDiskReplicaGroups](https://help.aliyun.com/document_detail/426614.html) operation to query a list of async replication pair-consistent groups, including group IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The maximum number of entries to be returned. You can use this parameter together with NextToken.
	//
	// Valid values: 1 to 500.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken. When you specify NextToken, the PageSize and PageNumber request parameters do not take effect and the TotalCount response parameter is invalid.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 5
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the primary or secondary disk in the async replication pair-consistent group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) operation to query the most recent list of regions in which async replication is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeReplicaGroupDrillsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicaGroupDrillsRequest) GoString() string {
	return s.String()
}

func (s *DescribeReplicaGroupDrillsRequest) GetDrillId() *string {
	return s.DrillId
}

func (s *DescribeReplicaGroupDrillsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeReplicaGroupDrillsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeReplicaGroupDrillsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeReplicaGroupDrillsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeReplicaGroupDrillsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeReplicaGroupDrillsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeReplicaGroupDrillsRequest) SetDrillId(v string) *DescribeReplicaGroupDrillsRequest {
	s.DrillId = &v
	return s
}

func (s *DescribeReplicaGroupDrillsRequest) SetGroupId(v string) *DescribeReplicaGroupDrillsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeReplicaGroupDrillsRequest) SetMaxResults(v int32) *DescribeReplicaGroupDrillsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeReplicaGroupDrillsRequest) SetNextToken(v string) *DescribeReplicaGroupDrillsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeReplicaGroupDrillsRequest) SetPageNumber(v int32) *DescribeReplicaGroupDrillsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeReplicaGroupDrillsRequest) SetPageSize(v int32) *DescribeReplicaGroupDrillsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeReplicaGroupDrillsRequest) SetRegionId(v string) *DescribeReplicaGroupDrillsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeReplicaGroupDrillsRequest) Validate() error {
	return dara.Validate(s)
}
