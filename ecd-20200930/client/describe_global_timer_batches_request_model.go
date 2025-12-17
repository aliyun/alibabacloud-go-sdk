// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalTimerBatchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DescribeGlobalTimerBatchesRequest
	GetGroupId() *string
	SetMaxResults(v string) *DescribeGlobalTimerBatchesRequest
	GetMaxResults() *string
	SetNextToken(v string) *DescribeGlobalTimerBatchesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeGlobalTimerBatchesRequest
	GetRegionId() *string
	SetSearchRegionId(v string) *DescribeGlobalTimerBatchesRequest
	GetSearchRegionId() *string
	SetTimerType(v string) *DescribeGlobalTimerBatchesRequest
	GetTimerType() *string
}

type DescribeGlobalTimerBatchesRequest struct {
	// example:
	//
	// ccg-i1ruuudp92qpj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 20
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-hangzhou
	SearchRegionId *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
	// example:
	//
	// 1
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s DescribeGlobalTimerBatchesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalTimerBatchesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalTimerBatchesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeGlobalTimerBatchesRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *DescribeGlobalTimerBatchesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeGlobalTimerBatchesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalTimerBatchesRequest) GetSearchRegionId() *string {
	return s.SearchRegionId
}

func (s *DescribeGlobalTimerBatchesRequest) GetTimerType() *string {
	return s.TimerType
}

func (s *DescribeGlobalTimerBatchesRequest) SetGroupId(v string) *DescribeGlobalTimerBatchesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeGlobalTimerBatchesRequest) SetMaxResults(v string) *DescribeGlobalTimerBatchesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeGlobalTimerBatchesRequest) SetNextToken(v string) *DescribeGlobalTimerBatchesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeGlobalTimerBatchesRequest) SetRegionId(v string) *DescribeGlobalTimerBatchesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalTimerBatchesRequest) SetSearchRegionId(v string) *DescribeGlobalTimerBatchesRequest {
	s.SearchRegionId = &v
	return s
}

func (s *DescribeGlobalTimerBatchesRequest) SetTimerType(v string) *DescribeGlobalTimerBatchesRequest {
	s.TimerType = &v
	return s
}

func (s *DescribeGlobalTimerBatchesRequest) Validate() error {
	return dara.Validate(s)
}
