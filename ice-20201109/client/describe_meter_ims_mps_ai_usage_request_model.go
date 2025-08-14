// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterImsMpsAiUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTs(v int64) *DescribeMeterImsMpsAiUsageRequest
	GetEndTs() *int64
	SetInterval(v int64) *DescribeMeterImsMpsAiUsageRequest
	GetInterval() *int64
	SetRegion(v string) *DescribeMeterImsMpsAiUsageRequest
	GetRegion() *string
	SetStartTs(v int64) *DescribeMeterImsMpsAiUsageRequest
	GetStartTs() *int64
}

type DescribeMeterImsMpsAiUsageRequest struct {
	// The end of the time range to query. The value is a 10-digit timestamp. The maximum query range is 31 days. The duration between StartTs and EndTs cannot exceed 31 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656995036
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// The time granularity of the query. Valid values: 3600 (hour) and 86400 (day).
	//
	// This parameter is required.
	//
	// example:
	//
	// 86400
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter does not take effect. By default, the usage data of all regions is returned.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The beginning of the time range to query. The value is a 10-digit timestamp. You can query data within the last 90 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1654403036
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterImsMpsAiUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsMpsAiUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsMpsAiUsageRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribeMeterImsMpsAiUsageRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeMeterImsMpsAiUsageRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeMeterImsMpsAiUsageRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeMeterImsMpsAiUsageRequest) SetEndTs(v int64) *DescribeMeterImsMpsAiUsageRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterImsMpsAiUsageRequest) SetInterval(v int64) *DescribeMeterImsMpsAiUsageRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMeterImsMpsAiUsageRequest) SetRegion(v string) *DescribeMeterImsMpsAiUsageRequest {
	s.Region = &v
	return s
}

func (s *DescribeMeterImsMpsAiUsageRequest) SetStartTs(v int64) *DescribeMeterImsMpsAiUsageRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeMeterImsMpsAiUsageRequest) Validate() error {
	return dara.Validate(s)
}
