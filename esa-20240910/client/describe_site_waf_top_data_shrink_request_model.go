// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteWafTopDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSiteWafTopDataShrinkRequest
	GetEndTime() *string
	SetFieldsShrink(v string) *DescribeSiteWafTopDataShrinkRequest
	GetFieldsShrink() *string
	SetInterval(v string) *DescribeSiteWafTopDataShrinkRequest
	GetInterval() *string
	SetLimit(v string) *DescribeSiteWafTopDataShrinkRequest
	GetLimit() *string
	SetSiteId(v string) *DescribeSiteWafTopDataShrinkRequest
	GetSiteId() *string
	SetStartTime(v string) *DescribeSiteWafTopDataShrinkRequest
	GetStartTime() *string
}

type DescribeSiteWafTopDataShrinkRequest struct {
	// The end of the time range to query.
	//
	// Specify the time in ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
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
	// The time granularity for querying data. Unit: seconds.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The number of top data entries to query.
	//
	// example:
	//
	// 5
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	//
	// If this parameter is left empty, user-level data is queried.
	//
	// example:
	//
	// 1150376036*****
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// example:
	//
	// 2023-04-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSiteWafTopDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteWafTopDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteWafTopDataShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteWafTopDataShrinkRequest) GetFieldsShrink() *string {
	return s.FieldsShrink
}

func (s *DescribeSiteWafTopDataShrinkRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeSiteWafTopDataShrinkRequest) GetLimit() *string {
	return s.Limit
}

func (s *DescribeSiteWafTopDataShrinkRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *DescribeSiteWafTopDataShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteWafTopDataShrinkRequest) SetEndTime(v string) *DescribeSiteWafTopDataShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteWafTopDataShrinkRequest) SetFieldsShrink(v string) *DescribeSiteWafTopDataShrinkRequest {
	s.FieldsShrink = &v
	return s
}

func (s *DescribeSiteWafTopDataShrinkRequest) SetInterval(v string) *DescribeSiteWafTopDataShrinkRequest {
	s.Interval = &v
	return s
}

func (s *DescribeSiteWafTopDataShrinkRequest) SetLimit(v string) *DescribeSiteWafTopDataShrinkRequest {
	s.Limit = &v
	return s
}

func (s *DescribeSiteWafTopDataShrinkRequest) SetSiteId(v string) *DescribeSiteWafTopDataShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeSiteWafTopDataShrinkRequest) SetStartTime(v string) *DescribeSiteWafTopDataShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteWafTopDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
