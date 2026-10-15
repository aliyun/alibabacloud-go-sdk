// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceUsageTotalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeResourceUsageTotalRequest
	GetBillOwnerId() *int64
	SetEndPeriod(v string) *DescribeResourceUsageTotalRequest
	GetEndPeriod() *string
	SetPeriodType(v string) *DescribeResourceUsageTotalRequest
	GetPeriodType() *string
	SetResourceType(v string) *DescribeResourceUsageTotalRequest
	GetResourceType() *string
	SetStartPeriod(v string) *DescribeResourceUsageTotalRequest
	GetStartPeriod() *string
}

type DescribeResourceUsageTotalRequest struct {
	// The ID of the resource ownership account. The resource ownership account is the account that actually uses the resource. Only one account ID can be specified.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The end time. Format: yyyy-MM-dd HH:mm:ss. The query excludes data at this time. If this parameter is left empty, the query range extends to the current time. This value specifies the end time of the query range.
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
	// The resource plan type. Valid values:
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
	// The start time. Format: yyyy-MM-dd HH:mm:ss. The query includes data at this time. This value specifies the start time of the query range.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01 00:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
}

func (s DescribeResourceUsageTotalRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceUsageTotalRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageTotalRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeResourceUsageTotalRequest) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeResourceUsageTotalRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeResourceUsageTotalRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeResourceUsageTotalRequest) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeResourceUsageTotalRequest) SetBillOwnerId(v int64) *DescribeResourceUsageTotalRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeResourceUsageTotalRequest) SetEndPeriod(v string) *DescribeResourceUsageTotalRequest {
	s.EndPeriod = &v
	return s
}

func (s *DescribeResourceUsageTotalRequest) SetPeriodType(v string) *DescribeResourceUsageTotalRequest {
	s.PeriodType = &v
	return s
}

func (s *DescribeResourceUsageTotalRequest) SetResourceType(v string) *DescribeResourceUsageTotalRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeResourceUsageTotalRequest) SetStartPeriod(v string) *DescribeResourceUsageTotalRequest {
	s.StartPeriod = &v
	return s
}

func (s *DescribeResourceUsageTotalRequest) Validate() error {
	return dara.Validate(s)
}
