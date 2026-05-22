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
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	FieldsShrink *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	Interval     *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	SiteId       *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
