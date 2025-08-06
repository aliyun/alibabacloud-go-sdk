// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRealtimeDeliveryAccRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeVodRealtimeDeliveryAccRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodRealtimeDeliveryAccRequest
	GetInterval() *string
	SetLogStore(v string) *DescribeVodRealtimeDeliveryAccRequest
	GetLogStore() *string
	SetOwnerId(v int64) *DescribeVodRealtimeDeliveryAccRequest
	GetOwnerId() *int64
	SetProject(v string) *DescribeVodRealtimeDeliveryAccRequest
	GetProject() *string
	SetStartTime(v string) *DescribeVodRealtimeDeliveryAccRequest
	GetStartTime() *string
}

type DescribeVodRealtimeDeliveryAccRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The end time must be later than the start time. The interval between the start time and the end time cannot exceed a year.
	//
	// example:
	//
	// 2016-06-30T13:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the data entries. Unit: seconds. Valid values: 300, 3600, and 86400.
	//
	// The default time granularity varies based on the interval between the start time and end time that you specify. If you set the interval to a value within (0,3] days, the default time granularity is 300 seconds. If you set the interval to a value within (3,30] days, the default time granularity is 3,600 seconds. If you set the interval to a value greater than 30 days, the default time granularity is 86,400 seconds. The default value is used if you specify an invalid value for this parameter or leave this parameter empty.
	//
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Logstore to which log entries are delivered.
	//
	// example:
	//
	// vod-log-store
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the Log Service project that is used for real-time log delivery.
	//
	// example:
	//
	// vod-logs
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2016-10-20T04:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodRealtimeDeliveryAccRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRealtimeDeliveryAccRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodRealtimeDeliveryAccRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodRealtimeDeliveryAccRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodRealtimeDeliveryAccRequest) GetLogStore() *string {
	return s.LogStore
}

func (s *DescribeVodRealtimeDeliveryAccRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodRealtimeDeliveryAccRequest) GetProject() *string {
	return s.Project
}

func (s *DescribeVodRealtimeDeliveryAccRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodRealtimeDeliveryAccRequest) SetEndTime(v string) *DescribeVodRealtimeDeliveryAccRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodRealtimeDeliveryAccRequest) SetInterval(v string) *DescribeVodRealtimeDeliveryAccRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodRealtimeDeliveryAccRequest) SetLogStore(v string) *DescribeVodRealtimeDeliveryAccRequest {
	s.LogStore = &v
	return s
}

func (s *DescribeVodRealtimeDeliveryAccRequest) SetOwnerId(v int64) *DescribeVodRealtimeDeliveryAccRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodRealtimeDeliveryAccRequest) SetProject(v string) *DescribeVodRealtimeDeliveryAccRequest {
	s.Project = &v
	return s
}

func (s *DescribeVodRealtimeDeliveryAccRequest) SetStartTime(v string) *DescribeVodRealtimeDeliveryAccRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodRealtimeDeliveryAccRequest) Validate() error {
	return dara.Validate(s)
}
