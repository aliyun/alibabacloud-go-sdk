// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceCoverageTotalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeResourceCoverageTotalRequest
	GetBillOwnerId() *int64
	SetEndPeriod(v string) *DescribeResourceCoverageTotalRequest
	GetEndPeriod() *string
	SetPeriodType(v string) *DescribeResourceCoverageTotalRequest
	GetPeriodType() *string
	SetResourceType(v string) *DescribeResourceCoverageTotalRequest
	GetResourceType() *string
	SetStartPeriod(v string) *DescribeResourceCoverageTotalRequest
	GetStartPeriod() *string
}

type DescribeResourceCoverageTotalRequest struct {
	// The ID of the account for which you want to query total coverage data. If you do not set this parameter, the data of the current Alibaba Cloud account and its RAM users is queried. To query the data of a RAM user, specify the ID of the RAM user.
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
	// The time granularity at which total coverage data is queried. Valid values: MONTH, DAY, and HOUR.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The type of deduction plans whose total coverage data is queried. Valid values: RI and SCU.
	//
	// This parameter is required.
	//
	// example:
	//
	// RI
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The beginning of the time range to query. The beginning is included in the time range. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01 00:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
}

func (s DescribeResourceCoverageTotalRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceCoverageTotalRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceCoverageTotalRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeResourceCoverageTotalRequest) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeResourceCoverageTotalRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeResourceCoverageTotalRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeResourceCoverageTotalRequest) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeResourceCoverageTotalRequest) SetBillOwnerId(v int64) *DescribeResourceCoverageTotalRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeResourceCoverageTotalRequest) SetEndPeriod(v string) *DescribeResourceCoverageTotalRequest {
	s.EndPeriod = &v
	return s
}

func (s *DescribeResourceCoverageTotalRequest) SetPeriodType(v string) *DescribeResourceCoverageTotalRequest {
	s.PeriodType = &v
	return s
}

func (s *DescribeResourceCoverageTotalRequest) SetResourceType(v string) *DescribeResourceCoverageTotalRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeResourceCoverageTotalRequest) SetStartPeriod(v string) *DescribeResourceCoverageTotalRequest {
	s.StartPeriod = &v
	return s
}

func (s *DescribeResourceCoverageTotalRequest) Validate() error {
	return dara.Validate(s)
}
