// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceUsageDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeResourceUsageDetailRequest
	GetBillOwnerId() *int64
	SetEndPeriod(v string) *DescribeResourceUsageDetailRequest
	GetEndPeriod() *string
	SetMaxResults(v int32) *DescribeResourceUsageDetailRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeResourceUsageDetailRequest
	GetNextToken() *string
	SetPeriodType(v string) *DescribeResourceUsageDetailRequest
	GetPeriodType() *string
	SetResourceType(v string) *DescribeResourceUsageDetailRequest
	GetResourceType() *string
	SetStartPeriod(v string) *DescribeResourceUsageDetailRequest
	GetStartPeriod() *string
}

type DescribeResourceUsageDetailRequest struct {
	// The ID of the account whose data you want to query. If this parameter is left empty, the data of the current account and its trusteeship linked accounts is queried. You can specify the ID of a trusteeship linked account to query only the data of that account.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The end time of the query. Format: yyyy-MM-dd HH:mm:ss. The specified time is excluded from the query. If this parameter is left empty, the current time is used.
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
	// The pagination token. Leave this parameter empty when you query the first page. The token for the next page is returned in the response. If the returned token is empty, all results have been returned.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6MTB9
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The time granularity. Valid values:
	//
	// - MONTH: Monthly.
	//
	// - DAY: Daily.
	//
	// - HOUR: Hourly.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The type of the resource plan. Valid values:
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
	// The start time of the query. Format: yyyy-MM-dd HH:mm:ss. The specified time is included in the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01 00:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
}

func (s DescribeResourceUsageDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceUsageDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageDetailRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeResourceUsageDetailRequest) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeResourceUsageDetailRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeResourceUsageDetailRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeResourceUsageDetailRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeResourceUsageDetailRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeResourceUsageDetailRequest) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeResourceUsageDetailRequest) SetBillOwnerId(v int64) *DescribeResourceUsageDetailRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeResourceUsageDetailRequest) SetEndPeriod(v string) *DescribeResourceUsageDetailRequest {
	s.EndPeriod = &v
	return s
}

func (s *DescribeResourceUsageDetailRequest) SetMaxResults(v int32) *DescribeResourceUsageDetailRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeResourceUsageDetailRequest) SetNextToken(v string) *DescribeResourceUsageDetailRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeResourceUsageDetailRequest) SetPeriodType(v string) *DescribeResourceUsageDetailRequest {
	s.PeriodType = &v
	return s
}

func (s *DescribeResourceUsageDetailRequest) SetResourceType(v string) *DescribeResourceUsageDetailRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeResourceUsageDetailRequest) SetStartPeriod(v string) *DescribeResourceUsageDetailRequest {
	s.StartPeriod = &v
	return s
}

func (s *DescribeResourceUsageDetailRequest) Validate() error {
	return dara.Validate(s)
}
