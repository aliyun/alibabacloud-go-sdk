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
	// The ID of the account whose data you want to query. If you do not specify this parameter, the data of the current Alibaba Cloud account and its Resource Access Management (RAM) users is queried. To query the data of a RAM user, specify the ID of the RAM user.
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
	// The maximum number of entries to return. Default value: 20. The maximum value is 300.
	//
	// example:
	//
	// 200
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6MTB9
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The time granularity at which usage details are queried. Valid values: MONTH, DAY, and HOUR.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The type of deduction plan whose usage details are queried. Valid values: RI and SCU.
	//
	// This parameter is required.
	//
	// example:
	//
	// RI
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The beginning of the time range to query. The beginning is included in the time range. Specify the time in the yyyy-MM-dd HH:mm:ss format.
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
