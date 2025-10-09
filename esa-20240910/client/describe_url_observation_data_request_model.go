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
	// The platform of the device. If the parameter is left empty, all devices are queried.
	//
	// 	- PC
	//
	// 	- Mobile
	//
	// example:
	//
	// PC
	ClientPlatform *string `json:"ClientPlatform,omitempty" xml:"ClientPlatform,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The time must be in UTC.
	//
	// example:
	//
	// 2023-04-19T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The metric data that is detected.
	//
	// 	- TTFB: Measures the time between when a resource initiates a request and when the first byte of the response starts to arrive.
	//
	// 	- FCP: Measures the time between when the page is loaded and when any part of the page\\"s content is rendered on the screen.
	//
	// 	- LCP: Reports the rendering time of the largest image or text block visible in the viewport.
	//
	// 	- CLS: A metric that measures the maximum layout mutation score for every unexpected layout change that occurs throughout the life of the page.
	//
	// 	- INP: Measures the responsiveness of the page, or how long it takes for the page to respond to user input in a visible way.
	//
	// 	- FID: Measures the time between when the user first interacts with the page and when the browser is actually able to start processing the event handler in response to that interaction.
	//
	// example:
	//
	// TTFB
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54362329990032
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The beginning of the time range to query.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
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
