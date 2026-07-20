// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUrlObservationDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientPlatform(v string) *DescribeUrlObservationDataRequest
	GetClientPlatform() *string
	SetEndTime(v string) *DescribeUrlObservationDataRequest
	GetEndTime() *string
	SetMetric(v string) *DescribeUrlObservationDataRequest
	GetMetric() *string
	SetSiteId(v string) *DescribeUrlObservationDataRequest
	GetSiteId() *string
	SetStartTime(v string) *DescribeUrlObservationDataRequest
	GetStartTime() *string
	SetUrl(v string) *DescribeUrlObservationDataRequest
	GetUrl() *string
}

type DescribeUrlObservationDataRequest struct {
	// The device platform. If this parameter is left empty, data for all platforms is queried.
	//
	// - PC
	//
	// - Mobile
	//
	// example:
	//
	// PC
	ClientPlatform *string `json:"ClientPlatform,omitempty" xml:"ClientPlatform,omitempty"`
	// The end time for the data query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2023-04-19T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The metric to query.
	//
	// 	- TTFB: Measures the time from when a resource request is initiated to when the first byte of the response begins to arrive.
	//
	// 	- FCP: Measures the time from when the page starts loading to when any part of the page content is rendered on the screen.
	//
	// 	- LCP: Reports the render time of the largest image or text block visible within the viewport.
	//
	// 	- CLS: A metric that measures the largest burst of layout shift scores for every unexpected layout shift that occurs throughout the entire lifecycle of a page.
	//
	// 	- INP: Measures the responsiveness of a page, specifically how long it takes for the page to visibly respond to user input.
	//
	// 	- FID: Measures the time from when a user first interacts with a page to when the browser is actually able to begin processing event handlers in response to that interaction.
	//
	// example:
	//
	// TTFB
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54362329990032
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The start time for the data query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2023-04-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The URL of the web page to monitor.
	//
	// example:
	//
	// example.com/test
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeUrlObservationDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUrlObservationDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUrlObservationDataRequest) GetClientPlatform() *string {
	return s.ClientPlatform
}

func (s *DescribeUrlObservationDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeUrlObservationDataRequest) GetMetric() *string {
	return s.Metric
}

func (s *DescribeUrlObservationDataRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *DescribeUrlObservationDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUrlObservationDataRequest) GetUrl() *string {
	return s.Url
}

func (s *DescribeUrlObservationDataRequest) SetClientPlatform(v string) *DescribeUrlObservationDataRequest {
	s.ClientPlatform = &v
	return s
}

func (s *DescribeUrlObservationDataRequest) SetEndTime(v string) *DescribeUrlObservationDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUrlObservationDataRequest) SetMetric(v string) *DescribeUrlObservationDataRequest {
	s.Metric = &v
	return s
}

func (s *DescribeUrlObservationDataRequest) SetSiteId(v string) *DescribeUrlObservationDataRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeUrlObservationDataRequest) SetStartTime(v string) *DescribeUrlObservationDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUrlObservationDataRequest) SetUrl(v string) *DescribeUrlObservationDataRequest {
	s.Url = &v
	return s
}

func (s *DescribeUrlObservationDataRequest) Validate() error {
	return dara.Validate(s)
}
