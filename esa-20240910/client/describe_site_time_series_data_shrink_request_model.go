// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteTimeSeriesDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSiteTimeSeriesDataShrinkRequest
	GetEndTime() *string
	SetFieldsShrink(v string) *DescribeSiteTimeSeriesDataShrinkRequest
	GetFieldsShrink() *string
	SetInterval(v string) *DescribeSiteTimeSeriesDataShrinkRequest
	GetInterval() *string
	SetSiteId(v string) *DescribeSiteTimeSeriesDataShrinkRequest
	GetSiteId() *string
	SetStartTime(v string) *DescribeSiteTimeSeriesDataShrinkRequest
	GetStartTime() *string
}

type DescribeSiteTimeSeriesDataShrinkRequest struct {
	// The end time for obtaining data.
	//
	// The date format follows ISO8601 notation and uses UTC+0 time, in the format yyyy-MM-ddTHH:mm:ssZ.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2023-04-09T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Query metrics.
	//
	// This parameter is required.
	FieldsShrink *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// The time granularity for querying data, in seconds.
	//
	// Depending on the maximum time span of a single query, this parameter supports values of 60 (1 minute), 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For details, see the **Supported Query Time Granularities**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Site ID. Obtain the site ID by calling the [ListSites](~~ListSites~~) interface.
	//
	// If this parameter is empty, user-level data will be queried.
	//
	// example:
	//
	// 1150376036*****
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The start time for obtaining data.
	//
	// The date format follows ISO8601 notation and uses UTC+0 time, in the format yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2023-04-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSiteTimeSeriesDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteTimeSeriesDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteTimeSeriesDataShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteTimeSeriesDataShrinkRequest) GetFieldsShrink() *string {
	return s.FieldsShrink
}

func (s *DescribeSiteTimeSeriesDataShrinkRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeSiteTimeSeriesDataShrinkRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *DescribeSiteTimeSeriesDataShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteTimeSeriesDataShrinkRequest) SetEndTime(v string) *DescribeSiteTimeSeriesDataShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataShrinkRequest) SetFieldsShrink(v string) *DescribeSiteTimeSeriesDataShrinkRequest {
	s.FieldsShrink = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataShrinkRequest) SetInterval(v string) *DescribeSiteTimeSeriesDataShrinkRequest {
	s.Interval = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataShrinkRequest) SetSiteId(v string) *DescribeSiteTimeSeriesDataShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataShrinkRequest) SetStartTime(v string) *DescribeSiteTimeSeriesDataShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
