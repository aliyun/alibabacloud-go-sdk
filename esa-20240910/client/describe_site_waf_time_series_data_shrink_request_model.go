// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteWafTimeSeriesDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSiteWafTimeSeriesDataShrinkRequest
	GetEndTime() *string
	SetFieldsShrink(v string) *DescribeSiteWafTimeSeriesDataShrinkRequest
	GetFieldsShrink() *string
	SetInterval(v string) *DescribeSiteWafTimeSeriesDataShrinkRequest
	GetInterval() *string
	SetSiteId(v string) *DescribeSiteWafTimeSeriesDataShrinkRequest
	GetSiteId() *string
	SetStartTime(v string) *DescribeSiteWafTimeSeriesDataShrinkRequest
	GetStartTime() *string
}

type DescribeSiteWafTimeSeriesDataShrinkRequest struct {
	// The end time for the data query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2023-04-09T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The query metrics.
	//
	// This parameter is required.
	FieldsShrink *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// The time granularity of the queried data, in seconds.
	//
	// Based on the maximum time span of a single query, this parameter supports the following values: 60 (1 minute), 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see the **supported query time granularity*	- section above.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// If this parameter is left empty, user-level data is queried.
	//
	// example:
	//
	// 11089268296****
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The start time for the data query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// example:
	//
	// 2023-04-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSiteWafTimeSeriesDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteWafTimeSeriesDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteWafTimeSeriesDataShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteWafTimeSeriesDataShrinkRequest) GetFieldsShrink() *string {
	return s.FieldsShrink
}

func (s *DescribeSiteWafTimeSeriesDataShrinkRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeSiteWafTimeSeriesDataShrinkRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *DescribeSiteWafTimeSeriesDataShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteWafTimeSeriesDataShrinkRequest) SetEndTime(v string) *DescribeSiteWafTimeSeriesDataShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataShrinkRequest) SetFieldsShrink(v string) *DescribeSiteWafTimeSeriesDataShrinkRequest {
	s.FieldsShrink = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataShrinkRequest) SetInterval(v string) *DescribeSiteWafTimeSeriesDataShrinkRequest {
	s.Interval = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataShrinkRequest) SetSiteId(v string) *DescribeSiteWafTimeSeriesDataShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataShrinkRequest) SetStartTime(v string) *DescribeSiteWafTimeSeriesDataShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
