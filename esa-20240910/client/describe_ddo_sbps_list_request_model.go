// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSBpsListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverage(v string) *DescribeDDoSBpsListRequest
	GetCoverage() *string
	SetEndTime(v string) *DescribeDDoSBpsListRequest
	GetEndTime() *string
	SetSiteId(v int64) *DescribeDDoSBpsListRequest
	GetSiteId() *int64
	SetStartTime(v string) *DescribeDDoSBpsListRequest
	GetStartTime() *string
}

type DescribeDDoSBpsListRequest struct {
	// Protection area, defaulting to global if not filled. When specified, the values are as follows:
	//
	// - domestic: Mainland China.
	//
	// - overseas: Global (excluding Mainland China).
	//
	// - global: Global.
	//
	// example:
	//
	// global
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The end time for fetching data. In ISO8601 format, using UTC+0, formatted as: yyyy-MM-ddTHH:mm:ssZ.
	//
	// The end time must be later than the start time, and the span between start and end times should not exceed 31 days.
	//
	// example:
	//
	// 2023-05-18T06:19:42Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 70966210986912
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The start time for fetching data, in ISO8601 format, using UTC+0, formatted as: yyyy-MM-ddTHH:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-14T17:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSBpsListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSBpsListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSBpsListRequest) GetCoverage() *string {
	return s.Coverage
}

func (s *DescribeDDoSBpsListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDDoSBpsListRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeDDoSBpsListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDDoSBpsListRequest) SetCoverage(v string) *DescribeDDoSBpsListRequest {
	s.Coverage = &v
	return s
}

func (s *DescribeDDoSBpsListRequest) SetEndTime(v string) *DescribeDDoSBpsListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSBpsListRequest) SetSiteId(v int64) *DescribeDDoSBpsListRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeDDoSBpsListRequest) SetStartTime(v string) *DescribeDDoSBpsListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSBpsListRequest) Validate() error {
	return dara.Validate(s)
}
