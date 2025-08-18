// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSL7QpsListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDDoSL7QpsListRequest
	GetEndTime() *string
	SetInterval(v int32) *DescribeDDoSL7QpsListRequest
	GetInterval() *int32
	SetRecordId(v int64) *DescribeDDoSL7QpsListRequest
	GetRecordId() *int64
	SetSiteId(v int64) *DescribeDDoSL7QpsListRequest
	GetSiteId() *int64
	SetStartTime(v string) *DescribeDDoSL7QpsListRequest
	GetStartTime() *string
}

type DescribeDDoSL7QpsListRequest struct {
	// The end time of the query.
	//
	// The date format follows ISO8601 notation and uses UTC+0, formatted as yyyy-MM-ddTHH:mm:ssZ. The maximum span between the start and end times is 31 days.
	//
	// If this parameter is not set, the current time will be used as the end time of the query.
	//
	// example:
	//
	// 2023-04-19T19:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the queried data, in seconds.
	//
	// Depending on the maximum time span of a single query, this parameter supports values of 60 (1 minute), 300 (5 minutes), 1800 (half an hour), and 3600 (1 hour).
	//
	// This parameter is required.
	//
	// example:
	//
	// 300
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Record ID, which can be obtained by calling the [ListRecords](~~ListRecords~~) interface.
	//
	// example:
	//
	// 86510927836942****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The start time of the query.
	//
	// The date format follows ISO8601 notation and uses UTC+0, formatted as yyyy-MM-ddTHH:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-19T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSL7QpsListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSL7QpsListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSL7QpsListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDDoSL7QpsListRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeDDoSL7QpsListRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *DescribeDDoSL7QpsListRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeDDoSL7QpsListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDDoSL7QpsListRequest) SetEndTime(v string) *DescribeDDoSL7QpsListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSL7QpsListRequest) SetInterval(v int32) *DescribeDDoSL7QpsListRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDDoSL7QpsListRequest) SetRecordId(v int64) *DescribeDDoSL7QpsListRequest {
	s.RecordId = &v
	return s
}

func (s *DescribeDDoSL7QpsListRequest) SetSiteId(v int64) *DescribeDDoSL7QpsListRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeDDoSL7QpsListRequest) SetStartTime(v string) *DescribeDDoSL7QpsListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSL7QpsListRequest) Validate() error {
	return dara.Validate(s)
}
