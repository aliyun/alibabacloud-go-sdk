// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticDbClusterIds(v string) *DescribeAgenticDBClustersRequest
	GetAgenticDbClusterIds() *string
	SetDBClusterDescription(v string) *DescribeAgenticDBClustersRequest
	GetDBClusterDescription() *string
	SetDBClusterIds(v string) *DescribeAgenticDBClustersRequest
	GetDBClusterIds() *string
	SetDBClusterStatus(v string) *DescribeAgenticDBClustersRequest
	GetDBClusterStatus() *string
	SetMaxResults(v int32) *DescribeAgenticDBClustersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAgenticDBClustersRequest
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeAgenticDBClustersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAgenticDBClustersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAgenticDBClustersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeAgenticDBClustersRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeAgenticDBClustersRequestTag) *DescribeAgenticDBClustersRequest
	GetTag() []*DescribeAgenticDBClustersRequestTag
}

type DescribeAgenticDBClustersRequest struct {
	// The Agentic cluster ID. Separate multiple cluster IDs with commas (,).
	//
	// example:
	//
	// pagc-**************
	AgenticDbClusterIds *string `json:"AgenticDbClusterIds,omitempty" xml:"AgenticDbClusterIds,omitempty"`
	// The cluster description. Fuzzy match is supported.
	//
	// example:
	//
	// pc-****************
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The cluster ID. Separate multiple cluster IDs with commas (,).
	//
	// example:
	//
	// pc-**************
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	// The cluster status.
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The maximum number of entries to return. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous API call. If there is no next query, do not pass this parameter.
	//
	// example:
	//
	// 212db86sca4384811e0b5e8707e******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 30, 50, and 100.
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*DescribeAgenticDBClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAgenticDBClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBClustersRequest) GetAgenticDbClusterIds() *string {
	return s.AgenticDbClusterIds
}

func (s *DescribeAgenticDBClustersRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeAgenticDBClustersRequest) GetDBClusterIds() *string {
	return s.DBClusterIds
}

func (s *DescribeAgenticDBClustersRequest) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeAgenticDBClustersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAgenticDBClustersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAgenticDBClustersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAgenticDBClustersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAgenticDBClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAgenticDBClustersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAgenticDBClustersRequest) GetTag() []*DescribeAgenticDBClustersRequestTag {
	return s.Tag
}

func (s *DescribeAgenticDBClustersRequest) SetAgenticDbClusterIds(v string) *DescribeAgenticDBClustersRequest {
	s.AgenticDbClusterIds = &v
	return s
}

func (s *DescribeAgenticDBClustersRequest) SetDBClusterDescription(v string) *DescribeAgenticDBClustersRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeAgenticDBClustersRequest) SetDBClusterIds(v string) *DescribeAgenticDBClustersRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeAgenticDBClustersRequest) SetDBClusterStatus(v string) *DescribeAgenticDBClustersRequest {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeAgenticDBClustersRequest) SetMaxResults(v int32) *DescribeAgenticDBClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAgenticDBClustersRequest) SetNextToken(v string) *DescribeAgenticDBClustersRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAgenticDBClustersRequest) SetPageNumber(v int32) *DescribeAgenticDBClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAgenticDBClustersRequest) SetPageSize(v int32) *DescribeAgenticDBClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAgenticDBClustersRequest) SetRegionId(v string) *DescribeAgenticDBClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAgenticDBClustersRequest) SetResourceGroupId(v string) *DescribeAgenticDBClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAgenticDBClustersRequest) SetTag(v []*DescribeAgenticDBClustersRequestTag) *DescribeAgenticDBClustersRequest {
	s.Tag = v
	return s
}

func (s *DescribeAgenticDBClustersRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAgenticDBClustersRequestTag struct {
	// The tag key. You can filter the cluster list by tag. You can specify up to 20 tag pairs. The number n for each tag pair must be unique and must be a consecutive integer starting from 1. The value of Tag.n.Key corresponds to Tag.n.Value.
	//
	// example:
	//
	// type
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value that corresponds to the tag key.
	//
	// example:
	//
	// testValueData
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAgenticDBClustersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBClustersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBClustersRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeAgenticDBClustersRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeAgenticDBClustersRequestTag) SetKey(v string) *DescribeAgenticDBClustersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeAgenticDBClustersRequestTag) SetValue(v string) *DescribeAgenticDBClustersRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeAgenticDBClustersRequestTag) Validate() error {
	return dara.Validate(s)
}
