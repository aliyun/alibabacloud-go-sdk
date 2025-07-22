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
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 2018-01-29T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// CN
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
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
