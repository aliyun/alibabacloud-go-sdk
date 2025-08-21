// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDestinationPortEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventType(v string) *DescribeDestinationPortEventRequest
	GetEventType() *string
	SetIp(v string) *DescribeDestinationPortEventRequest
	GetIp() *string
	SetRange(v int64) *DescribeDestinationPortEventRequest
	GetRange() *int64
	SetRegion(v string) *DescribeDestinationPortEventRequest
	GetRegion() *string
	SetStartTime(v int64) *DescribeDestinationPortEventRequest
	GetStartTime() *int64
}

type DescribeDestinationPortEventRequest struct {
	// The type of the attack event that you want to query. Valid values:
	//
	// 	- **defense**: attack events that trigger traffic scrubbing.
	//
	// 	- **blackhole**: attack events that trigger blackhole filtering.
	//
	// This parameter is required.
	//
	// example:
	//
	// defense
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The IP address of the attacker.
	//
	// This parameter is required.
	//
	// example:
	//
	// 203.107.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The number of destination ports to return. The ports are sorted in descending order of the number of received request packets. By default, the first **10*	- ports are returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Range *int64 `json:"Range,omitempty" xml:"Range,omitempty"`
	// The region in which your service is deployed. Valid values:
	//
	// 	- **cn**: a region in the Chinese mainland.
	//
	// 	- **cn-hongkong**: a region outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// >  This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1720059000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDestinationPortEventRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDestinationPortEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeDestinationPortEventRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeDestinationPortEventRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeDestinationPortEventRequest) GetRange() *int64 {
	return s.Range
}

func (s *DescribeDestinationPortEventRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeDestinationPortEventRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDestinationPortEventRequest) SetEventType(v string) *DescribeDestinationPortEventRequest {
	s.EventType = &v
	return s
}

func (s *DescribeDestinationPortEventRequest) SetIp(v string) *DescribeDestinationPortEventRequest {
	s.Ip = &v
	return s
}

func (s *DescribeDestinationPortEventRequest) SetRange(v int64) *DescribeDestinationPortEventRequest {
	s.Range = &v
	return s
}

func (s *DescribeDestinationPortEventRequest) SetRegion(v string) *DescribeDestinationPortEventRequest {
	s.Region = &v
	return s
}

func (s *DescribeDestinationPortEventRequest) SetStartTime(v int64) *DescribeDestinationPortEventRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDestinationPortEventRequest) Validate() error {
	return dara.Validate(s)
}
