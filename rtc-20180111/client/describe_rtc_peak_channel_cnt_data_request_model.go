// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcPeakChannelCntDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeRtcPeakChannelCntDataRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeRtcPeakChannelCntDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeRtcPeakChannelCntDataRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeRtcPeakChannelCntDataRequest
	GetOwnerId() *int64
	SetServiceArea(v string) *DescribeRtcPeakChannelCntDataRequest
	GetServiceArea() *string
	SetStartTime(v string) *DescribeRtcPeakChannelCntDataRequest
	GetStartTime() *string
}

type DescribeRtcPeakChannelCntDataRequest struct {
	// The application ID. Only a single ID can be specified.
	//
	// By default, all application IDs are queried.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time for the data query. Specify the time in the UTC format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// The end time must be later than the start time.
	//
	// example:
	//
	// 2018-01-29T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity for the data query, in seconds.
	//
	// Valid values: **3600*	- (1 hour) or **86400*	- (1 day). Default value: **3600**.
	//
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The service area. CN: China (default value).
	//
	// example:
	//
	// CN
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	// The start time for the data query. Specify the time in the UTC format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2018-01-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRtcPeakChannelCntDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcPeakChannelCntDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakChannelCntDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeRtcPeakChannelCntDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRtcPeakChannelCntDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeRtcPeakChannelCntDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRtcPeakChannelCntDataRequest) GetServiceArea() *string {
	return s.ServiceArea
}

func (s *DescribeRtcPeakChannelCntDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetAppId(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetEndTime(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetInterval(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetOwnerId(v int64) *DescribeRtcPeakChannelCntDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetServiceArea(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetStartTime(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) Validate() error {
	return dara.Validate(s)
}
