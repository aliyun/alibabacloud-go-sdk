// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisTrafficRankingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeNisTrafficRankingRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeNisTrafficRankingRequest
	GetNextToken() *string
	SetNisTrafficRankingId(v string) *DescribeNisTrafficRankingRequest
	GetNisTrafficRankingId() *string
}

type DescribeNisTrafficRankingRequest struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the value of NextToken returned by the previous API call.
	//
	// example:
	//
	// 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the traffic ranking analysis result. Set this parameter to the value returned by the StartNisTrafficRanking operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// task-6462a7b4c4a54b****
	NisTrafficRankingId *string `json:"NisTrafficRankingId,omitempty" xml:"NisTrafficRankingId,omitempty"`
}

func (s DescribeNisTrafficRankingRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisTrafficRankingRequest) GoString() string {
	return s.String()
}

func (s *DescribeNisTrafficRankingRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNisTrafficRankingRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNisTrafficRankingRequest) GetNisTrafficRankingId() *string {
	return s.NisTrafficRankingId
}

func (s *DescribeNisTrafficRankingRequest) SetMaxResults(v int32) *DescribeNisTrafficRankingRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNisTrafficRankingRequest) SetNextToken(v string) *DescribeNisTrafficRankingRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeNisTrafficRankingRequest) SetNisTrafficRankingId(v string) *DescribeNisTrafficRankingRequest {
	s.NisTrafficRankingId = &v
	return s
}

func (s *DescribeNisTrafficRankingRequest) Validate() error {
	return dara.Validate(s)
}
