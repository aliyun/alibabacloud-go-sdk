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
	// The ID of the account for which you want to query coverage details. If you do not set this parameter, the data of the current Alibaba Cloud account and its RAM users is queried. To query the data of a RAM user, specify the ID of the RAM user.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The end of the time range to query. The end is excluded from the time range. If you do not set this parameter, the end time is the current time. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2021-01-02 00:00:00
	EndPeriod *string `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	// The maximum number of entries to return. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 200
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. You do not need to set this parameter if you query coverage details within a specific time range for the first time. The response returns a token that you can use to query coverage details that are displayed on the next page. If a null value is returned for the NextToken parameter, no more coverage details can be queried.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6MTB9
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The time granularity at which coverage details are queried. Valid values: MONTH, DAY, and HOUR.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The type of deduction plans whose coverage details are queried. Valid values: RI and SCU.
	//
	// This parameter is required.
	//
	// example:
	//
	// RI
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The beginning of the time range to query.
	//
	// The beginning is included in the time range. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
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
