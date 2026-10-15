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
	// The ID of the resource ownership account. The resource ownership account is the account that actually uses the resource.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The end of the query time range. Format: yyyy-MM-dd HH:mm:ss. The query excludes this time. If this parameter is left empty, the current time is used.
	//
	// example:
	//
	// 2021-01-02 00:00:00
	EndPeriod *string `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	// The time granularity. Valid values:
	//
	// - MONTH: month
	//
	// - DAY: day
	//
	// - HOUR: hour
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The resource type. Valid values:
	//
	// - RI: reserved instance
	//
	// - SCU: storage capacity unit
	//
	// This parameter is required.
	//
	// example:
	//
	// RI
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The start of the query time range. Format: yyyy-MM-dd HH:mm:ss. The query includes this time.
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
