// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcDurationDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeRtcDurationDataRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeRtcDurationDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeRtcDurationDataRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeRtcDurationDataRequest
	GetOwnerId() *int64
	SetServiceArea(v string) *DescribeRtcDurationDataRequest
	GetServiceArea() *string
	SetStartTime(v string) *DescribeRtcDurationDataRequest
	GetStartTime() *string
}

type DescribeRtcDurationDataRequest struct {
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 2020-02-04T07:00:00Z
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
	// 2020-02-04T05:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRtcDurationDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcDurationDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcDurationDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeRtcDurationDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRtcDurationDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeRtcDurationDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRtcDurationDataRequest) GetServiceArea() *string {
	return s.ServiceArea
}

func (s *DescribeRtcDurationDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRtcDurationDataRequest) SetAppId(v string) *DescribeRtcDurationDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetEndTime(v string) *DescribeRtcDurationDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetInterval(v string) *DescribeRtcDurationDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetOwnerId(v int64) *DescribeRtcDurationDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetServiceArea(v string) *DescribeRtcDurationDataRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetStartTime(v string) *DescribeRtcDurationDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) Validate() error {
	return dara.Validate(s)
}
