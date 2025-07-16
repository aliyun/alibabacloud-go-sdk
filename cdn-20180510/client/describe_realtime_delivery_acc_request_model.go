// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRealtimeDeliveryAccRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeRealtimeDeliveryAccRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeRealtimeDeliveryAccRequest
	GetInterval() *string
	SetLogStore(v string) *DescribeRealtimeDeliveryAccRequest
	GetLogStore() *string
	SetProject(v string) *DescribeRealtimeDeliveryAccRequest
	GetProject() *string
	SetStartTime(v string) *DescribeRealtimeDeliveryAccRequest
	GetStartTime() *string
}

type DescribeRealtimeDeliveryAccRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The end time must be later than the start time.
	//
	// example:
	//
	// 2016-10-20T05:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the data entries. Unit: seconds. The value varies based on the values of the **StartTime*	- and **EndTime*	- parameters. Valid values:
	//
	// 	- If the time span between StartTime and EndTime is less than 3 days, valid values are **300**, **3600**, and **86400**. Default value: **300**.
	//
	// 	- If the time span between StartTime and EndTime is greater than or equal to 3 days and less than 31 days, valid values are **3600*	- and **86400**. Default value: **3600**.
	//
	// 	- If the time span between StartTime and EndTime is 31 days or longer, the valid value is **86400**. Default value: **86400**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Logstore that stores log data. If you do leave this parameter empty, real-time log deliveries of all Logstores are queried.
	//
	// example:
	//
	// LogStore
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The name of the Log Service project that is used for real-time log delivery. If you do leave this parameter empty, real-time log deliveries of all projects are queried.
	//
	// example:
	//
	// Project
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2016-10-20T04:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRealtimeDeliveryAccRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRealtimeDeliveryAccRequest) GoString() string {
	return s.String()
}

func (s *DescribeRealtimeDeliveryAccRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRealtimeDeliveryAccRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeRealtimeDeliveryAccRequest) GetLogStore() *string {
	return s.LogStore
}

func (s *DescribeRealtimeDeliveryAccRequest) GetProject() *string {
	return s.Project
}

func (s *DescribeRealtimeDeliveryAccRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRealtimeDeliveryAccRequest) SetEndTime(v string) *DescribeRealtimeDeliveryAccRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccRequest) SetInterval(v string) *DescribeRealtimeDeliveryAccRequest {
	s.Interval = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccRequest) SetLogStore(v string) *DescribeRealtimeDeliveryAccRequest {
	s.LogStore = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccRequest) SetProject(v string) *DescribeRealtimeDeliveryAccRequest {
	s.Project = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccRequest) SetStartTime(v string) *DescribeRealtimeDeliveryAccRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccRequest) Validate() error {
	return dara.Validate(s)
}
