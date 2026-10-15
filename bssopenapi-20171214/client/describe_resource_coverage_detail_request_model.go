// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceCoverageDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeResourceCoverageDetailRequest
	GetBillOwnerId() *int64
	SetEndPeriod(v string) *DescribeResourceCoverageDetailRequest
	GetEndPeriod() *string
	SetMaxResults(v int32) *DescribeResourceCoverageDetailRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeResourceCoverageDetailRequest
	GetNextToken() *string
	SetPeriodType(v string) *DescribeResourceCoverageDetailRequest
	GetPeriodType() *string
	SetResourceType(v string) *DescribeResourceCoverageDetailRequest
	GetResourceType() *string
	SetStartPeriod(v string) *DescribeResourceCoverageDetailRequest
	GetStartPeriod() *string
}

type DescribeResourceCoverageDetailRequest struct {
	// The ID of the account to which the resource belongs. The resource ownership account is the account that actually uses the resource.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The end time of the query period. Format: yyyy-MM-dd HH:mm:ss. The specified time is excluded from the query. If this parameter is left empty, the current time is used as the end time.
	//
	// example:
	//
	// 2021-01-02 00:00:00
	EndPeriod *string `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	// The maximum number of entries to return in a single request. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 200
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Leave this parameter empty when you query the first page. The token for the next page is returned in the response. If the returned token is empty, all results have been retrieved.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6MTB9
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The time granularity of the query. Valid values:
	//
	// - MONTH: month.
	//
	// - DAY: day.
	//
	// - HOUR: hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The type of the resource. Valid values:
	//
	// - RI: reserved instance.
	//
	// - SCU: storage capacity unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// RI
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The start time of the query period. Format: yyyy-MM-dd HH:mm:ss. The specified time is included in the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01 00:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
}

func (s DescribeResourceCoverageDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceCoverageDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceCoverageDetailRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeResourceCoverageDetailRequest) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeResourceCoverageDetailRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeResourceCoverageDetailRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeResourceCoverageDetailRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeResourceCoverageDetailRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeResourceCoverageDetailRequest) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeResourceCoverageDetailRequest) SetBillOwnerId(v int64) *DescribeResourceCoverageDetailRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeResourceCoverageDetailRequest) SetEndPeriod(v string) *DescribeResourceCoverageDetailRequest {
	s.EndPeriod = &v
	return s
}

func (s *DescribeResourceCoverageDetailRequest) SetMaxResults(v int32) *DescribeResourceCoverageDetailRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeResourceCoverageDetailRequest) SetNextToken(v string) *DescribeResourceCoverageDetailRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeResourceCoverageDetailRequest) SetPeriodType(v string) *DescribeResourceCoverageDetailRequest {
	s.PeriodType = &v
	return s
}

func (s *DescribeResourceCoverageDetailRequest) SetResourceType(v string) *DescribeResourceCoverageDetailRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeResourceCoverageDetailRequest) SetStartPeriod(v string) *DescribeResourceCoverageDetailRequest {
	s.StartPeriod = &v
	return s
}

func (s *DescribeResourceCoverageDetailRequest) Validate() error {
	return dara.Validate(s)
}
