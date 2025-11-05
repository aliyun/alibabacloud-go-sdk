// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePairDrillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrillId(v string) *DescribePairDrillsRequest
	GetDrillId() *string
	SetMaxResults(v int64) *DescribePairDrillsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribePairDrillsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *DescribePairDrillsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePairDrillsRequest
	GetPageSize() *int32
	SetPairId(v string) *DescribePairDrillsRequest
	GetPairId() *string
	SetRegionId(v string) *DescribePairDrillsRequest
	GetRegionId() *string
}

type DescribePairDrillsRequest struct {
	// The ID of the drill.
	//
	// example:
	//
	// drill-xxxx
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The maximum number of entries to be returned. You can use this parameter together with NextToken.
	//
	// Valid values: 1 to 500.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Set the value to the NextToken value returned in the previous call to the DescribeDiskReplicaPairs operation. Leave this parameter empty the first time you call this operation. When you specify NextToken, the PageSize and PageNumber request parameters do not take effect and the TotalCount response parameter is invalid.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the replication pair. You can call the [DescribeDiskReplicaPairs](https://help.aliyun.com/document_detail/354206.html) operation to query a list of asynchronous replication pairs, including replication pair IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pair-xxxx
	PairId *string `json:"PairId,omitempty" xml:"PairId,omitempty"`
	// The region ID of the primary or secondary disk in the async replication pair. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) operation to query the most recent list of regions in which async replication is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePairDrillsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePairDrillsRequest) GoString() string {
	return s.String()
}

func (s *DescribePairDrillsRequest) GetDrillId() *string {
	return s.DrillId
}

func (s *DescribePairDrillsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribePairDrillsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePairDrillsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePairDrillsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePairDrillsRequest) GetPairId() *string {
	return s.PairId
}

func (s *DescribePairDrillsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePairDrillsRequest) SetDrillId(v string) *DescribePairDrillsRequest {
	s.DrillId = &v
	return s
}

func (s *DescribePairDrillsRequest) SetMaxResults(v int64) *DescribePairDrillsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePairDrillsRequest) SetNextToken(v string) *DescribePairDrillsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePairDrillsRequest) SetPageNumber(v int32) *DescribePairDrillsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePairDrillsRequest) SetPageSize(v int32) *DescribePairDrillsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePairDrillsRequest) SetPairId(v string) *DescribePairDrillsRequest {
	s.PairId = &v
	return s
}

func (s *DescribePairDrillsRequest) SetRegionId(v string) *DescribePairDrillsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePairDrillsRequest) Validate() error {
	return dara.Validate(s)
}
