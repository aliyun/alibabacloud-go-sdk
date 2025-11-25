// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDDoSEvents(v []*DescribeDDoSEventsResponseBodyDDoSEvents) *DescribeDDoSEventsResponseBody
	GetDDoSEvents() []*DescribeDDoSEventsResponseBodyDDoSEvents
	SetRequestId(v string) *DescribeDDoSEventsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeDDoSEventsResponseBody
	GetTotal() *int64
}

type DescribeDDoSEventsResponseBody struct {
	// The DDoS attack events.
	DDoSEvents []*DescribeDDoSEventsResponseBodyDDoSEvents `json:"DDoSEvents,omitempty" xml:"DDoSEvents,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0CA72AF5-1795-4350-8C77-50A448A2F334
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned attack events.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDDoSEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsResponseBody) GetDDoSEvents() []*DescribeDDoSEventsResponseBodyDDoSEvents {
	return s.DDoSEvents
}

func (s *DescribeDDoSEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDoSEventsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeDDoSEventsResponseBody) SetDDoSEvents(v []*DescribeDDoSEventsResponseBodyDDoSEvents) *DescribeDDoSEventsResponseBody {
	s.DDoSEvents = v
	return s
}

func (s *DescribeDDoSEventsResponseBody) SetRequestId(v string) *DescribeDDoSEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSEventsResponseBody) SetTotal(v int64) *DescribeDDoSEventsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDDoSEventsResponseBody) Validate() error {
	if s.DDoSEvents != nil {
		for _, item := range s.DDoSEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDDoSEventsResponseBodyDDoSEvents struct {
	// The bandwidth of attack traffic. Unit: bit/s.
	//
	// example:
	//
	// 0
	Bps *int64 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The time when the DDoS attack stopped. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1583933330
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the attack event. Valid values:
	//
	// 	- **defense**: traffic scrubbing events
	//
	// 	- **blackhole**: blackhole filtering events
	//
	// example:
	//
	// blackhole
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The attacked IP address.
	//
	// example:
	//
	// 203.***.***.132
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The attacked port.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The packet forwarding rate of attack traffic. Unit: packets per second (pps).
	//
	// example:
	//
	// 0
	Pps *int64 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The region from which the attack was launched. Valid values:
	//
	// 	- **cn**: a region in the Chinese mainland
	//
	// 	- **alb-ap-northeast-1-gf-x**: Japan (Tokyo)
	//
	// 	- **alb-ap-southeast-gf-x**: Singapore
	//
	// 	- **alb-cn-hongkong-gf-x**: Hong Kong (China)
	//
	// 	- **alb-eu-central-1-gf-x**: Germany (Frankfurt)
	//
	// 	- **alb-us-west-1-gf-x**: US (Silicon Valley)
	//
	// > The values except **cn*	- are returned only when **RegionId*	- is set to **ap-southeast-1**.
	//
	// example:
	//
	// cn
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The time when the DDoS attack started. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1583933277
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSEventsResponseBodyDDoSEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSEventsResponseBodyDDoSEvents) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) GetBps() *int64 {
	return s.Bps
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) GetEventType() *string {
	return s.EventType
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) GetIp() *string {
	return s.Ip
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) GetPort() *string {
	return s.Port
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) GetPps() *int64 {
	return s.Pps
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) GetRegion() *string {
	return s.Region
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetBps(v int64) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.Bps = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetEndTime(v int64) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetEventType(v string) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.EventType = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetIp(v string) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.Ip = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetPort(v string) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.Port = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetPps(v int64) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.Pps = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetRegion(v string) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.Region = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetStartTime(v int64) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) Validate() error {
	return dara.Validate(s)
}
