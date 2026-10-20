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
	// The end time of the query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The time span between the start time and the end time cannot exceed 31 days. If you do not specify this parameter, the current time is used as the end time.
	//
	// example:
	//
	// 2023-04-19T19:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the queried data. Unit: seconds. Valid values: 60 (1 minute), 300 (5 minutes), 1800 (30 minutes), and 3600 (1 hour). The supported values vary based on the maximum time span of a single query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The record ID. You can call the [ListRecords](~~ListRecords~~) operation to obtain the record ID.
	//
	// example:
	//
	// 86510927836942****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The start time of the query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
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
