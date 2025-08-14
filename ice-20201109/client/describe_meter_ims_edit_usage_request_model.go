// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterImsEditUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTs(v int64) *DescribeMeterImsEditUsageRequest
	GetEndTs() *int64
	SetInterval(v int64) *DescribeMeterImsEditUsageRequest
	GetInterval() *int64
	SetRegion(v string) *DescribeMeterImsEditUsageRequest
	GetRegion() *string
	SetStartTs(v int64) *DescribeMeterImsEditUsageRequest
	GetStartTs() *int64
}

type DescribeMeterImsEditUsageRequest struct {
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
	// 86400
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter does not take effect. By default, the usage data of all regions is returned.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The beginning of the time range to query. The value is a 10-digit timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1654403036
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterImsEditUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsEditUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsEditUsageRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribeMeterImsEditUsageRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeMeterImsEditUsageRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeMeterImsEditUsageRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeMeterImsEditUsageRequest) SetEndTs(v int64) *DescribeMeterImsEditUsageRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterImsEditUsageRequest) SetInterval(v int64) *DescribeMeterImsEditUsageRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMeterImsEditUsageRequest) SetRegion(v string) *DescribeMeterImsEditUsageRequest {
	s.Region = &v
	return s
}

func (s *DescribeMeterImsEditUsageRequest) SetStartTs(v int64) *DescribeMeterImsEditUsageRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeMeterImsEditUsageRequest) Validate() error {
	return dara.Validate(s)
}
