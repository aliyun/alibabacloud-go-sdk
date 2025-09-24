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
	// The ID of the account whose data you want to query. If you do not specify this parameter, the data of the current account and its linked accounts is queried. To query the data of a linked account, specify the ID of the linked account. You can specify only one account ID.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-dd HH:mm:ss format. The specified time is excluded from the time range. If you do not specify this parameter, this parameter is set to the current time.
	//
	// example:
	//
	// 2021-01-02 00:00:00
	EndPeriod *string `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	// The time granularity at which the data is queried. Valid values: MONTH, DAY, and HOUR.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The type of the resource plan. Valid values: RI and SCU.
	//
	// This parameter is required.
	//
	// example:
	//
	// RI
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-dd HH:mm:ss format. The specified time is included in the time range.
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
