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
	// example:
	//
	// PC
	ClientPlatform *string `json:"ClientPlatform,omitempty" xml:"ClientPlatform,omitempty"`
	// example:
	//
	// 2023-04-19T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// TTFB
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 54362329990032
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 2023-04-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
