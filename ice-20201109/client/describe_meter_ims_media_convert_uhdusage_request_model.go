// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterImsMediaConvertUHDUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTs(v int64) *DescribeMeterImsMediaConvertUHDUsageRequest
	GetEndTs() *int64
	SetInterval(v string) *DescribeMeterImsMediaConvertUHDUsageRequest
	GetInterval() *string
	SetRegionId(v string) *DescribeMeterImsMediaConvertUHDUsageRequest
	GetRegionId() *string
	SetStartTs(v int64) *DescribeMeterImsMediaConvertUHDUsageRequest
	GetStartTs() *int64
}

type DescribeMeterImsMediaConvertUHDUsageRequest struct {
	// The end of the time range to query. The value is a 10-digit timestamp.
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
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter does not take effect. By default, the usage data of all regions is returned.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. The value is a 10-digit timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1654403036
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterImsMediaConvertUHDUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsMediaConvertUHDUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsMediaConvertUHDUsageRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribeMeterImsMediaConvertUHDUsageRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeMeterImsMediaConvertUHDUsageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMeterImsMediaConvertUHDUsageRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeMeterImsMediaConvertUHDUsageRequest) SetEndTs(v int64) *DescribeMeterImsMediaConvertUHDUsageRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterImsMediaConvertUHDUsageRequest) SetInterval(v string) *DescribeMeterImsMediaConvertUHDUsageRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMeterImsMediaConvertUHDUsageRequest) SetRegionId(v string) *DescribeMeterImsMediaConvertUHDUsageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMeterImsMediaConvertUHDUsageRequest) SetStartTs(v int64) *DescribeMeterImsMediaConvertUHDUsageRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeMeterImsMediaConvertUHDUsageRequest) Validate() error {
	return dara.Validate(s)
}
