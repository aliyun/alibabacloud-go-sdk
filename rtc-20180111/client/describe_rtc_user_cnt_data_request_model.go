// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcUserCntDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeRtcUserCntDataRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeRtcUserCntDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeRtcUserCntDataRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeRtcUserCntDataRequest
	GetOwnerId() *int64
	SetServiceArea(v string) *DescribeRtcUserCntDataRequest
	GetServiceArea() *string
	SetStartTime(v string) *DescribeRtcUserCntDataRequest
	GetStartTime() *string
}

type DescribeRtcUserCntDataRequest struct {
	// The application ID. You can specify only one ID.
	//
	// By default, all application IDs are queried.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time of the queried data. The time is displayed in UTC. The format is yyyy-MM-ddTHH:mm:ssZ.
	//
	// The end time must be later than the start time.
	//
	// example:
	//
	// 2018-01-29T01:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity at which the data is queried. Unit: seconds.
	//
	// Valid values: **3600*	- (one hour) or **86400*	- (one day). Default value: **3600**.
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
	// The start time of the queried data. The time is displayed in UTC. The format is yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2018-01-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRtcUserCntDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcUserCntDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserCntDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeRtcUserCntDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRtcUserCntDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeRtcUserCntDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRtcUserCntDataRequest) GetServiceArea() *string {
	return s.ServiceArea
}

func (s *DescribeRtcUserCntDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRtcUserCntDataRequest) SetAppId(v string) *DescribeRtcUserCntDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetEndTime(v string) *DescribeRtcUserCntDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetInterval(v string) *DescribeRtcUserCntDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetOwnerId(v int64) *DescribeRtcUserCntDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetServiceArea(v string) *DescribeRtcUserCntDataRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetStartTime(v string) *DescribeRtcUserCntDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) Validate() error {
	return dara.Validate(s)
}
