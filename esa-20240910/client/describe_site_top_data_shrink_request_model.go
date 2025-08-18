// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteTopDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSiteTopDataShrinkRequest
	GetEndTime() *string
	SetFieldsShrink(v string) *DescribeSiteTopDataShrinkRequest
	GetFieldsShrink() *string
	SetInterval(v string) *DescribeSiteTopDataShrinkRequest
	GetInterval() *string
	SetLimit(v string) *DescribeSiteTopDataShrinkRequest
	GetLimit() *string
	SetSiteId(v string) *DescribeSiteTopDataShrinkRequest
	GetSiteId() *string
	SetStartTime(v string) *DescribeSiteTopDataShrinkRequest
	GetStartTime() *string
}

type DescribeSiteTopDataShrinkRequest struct {
	// example:
	//
	// 2023-04-09T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	FieldsShrink *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Limit    *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// 1150376036*****
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 2023-04-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSiteTopDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteTopDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteTopDataShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteTopDataShrinkRequest) GetFieldsShrink() *string {
	return s.FieldsShrink
}

func (s *DescribeSiteTopDataShrinkRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeSiteTopDataShrinkRequest) GetLimit() *string {
	return s.Limit
}

func (s *DescribeSiteTopDataShrinkRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *DescribeSiteTopDataShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteTopDataShrinkRequest) SetEndTime(v string) *DescribeSiteTopDataShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteTopDataShrinkRequest) SetFieldsShrink(v string) *DescribeSiteTopDataShrinkRequest {
	s.FieldsShrink = &v
	return s
}

func (s *DescribeSiteTopDataShrinkRequest) SetInterval(v string) *DescribeSiteTopDataShrinkRequest {
	s.Interval = &v
	return s
}

func (s *DescribeSiteTopDataShrinkRequest) SetLimit(v string) *DescribeSiteTopDataShrinkRequest {
	s.Limit = &v
	return s
}

func (s *DescribeSiteTopDataShrinkRequest) SetSiteId(v string) *DescribeSiteTopDataShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeSiteTopDataShrinkRequest) SetStartTime(v string) *DescribeSiteTopDataShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteTopDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
