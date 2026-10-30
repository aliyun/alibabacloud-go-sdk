// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*DescribeDdosEventResponseBodyEvents) *DescribeDdosEventResponseBody
	GetEvents() []*DescribeDdosEventResponseBodyEvents
	SetRequestId(v string) *DescribeDdosEventResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeDdosEventResponseBody
	GetTotal() *int64
}

type DescribeDdosEventResponseBody struct {
	// The details of the DDoS attack events.
	Events []*DescribeDdosEventResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// F3B6C3F9-6B21-519D-B976-A1E14166F909
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of DDoS attack events returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDdosEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponseBody) GetEvents() []*DescribeDdosEventResponseBodyEvents {
	return s.Events
}

func (s *DescribeDdosEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDdosEventResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeDdosEventResponseBody) SetEvents(v []*DescribeDdosEventResponseBodyEvents) *DescribeDdosEventResponseBody {
	s.Events = v
	return s
}

func (s *DescribeDdosEventResponseBody) SetRequestId(v string) *DescribeDdosEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosEventResponseBody) SetTotal(v int64) *DescribeDdosEventResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDdosEventResponseBody) Validate() error {
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDdosEventResponseBodyEvents struct {
	// The end time of the attack. The time is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1637554335
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The attacked IP address.
	//
	// example:
	//
	// 47.89.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The traffic volume at the start of the attack. Unit: Mbit/s.
	//
	// example:
	//
	// 0
	Mbps *int32 `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
	// The packet rate at the start of the attack. Unit: pps.
	//
	// example:
	//
	// 456
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The end time of the attack. The time is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1637554034
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The current status of the attack event. Valid values:
	//
	// - **hole_begin**: Indicates that the attacked IP address is in the blackhole filtering status.
	//
	// - **hole_end**: Indicates that the attacked IP address is removed from the blackhole filtering status.
	//
	// - **defense_begin**: Indicates that the system scrubs traffic.
	//
	// - **defense_end**: Indicates that the attack ends and the system scrubs traffic no more.
	//
	// example:
	//
	// defense_end
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDdosEventResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosEventResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponseBodyEvents) GetEndTime() *int32 {
	return s.EndTime
}

func (s *DescribeDdosEventResponseBodyEvents) GetIp() *string {
	return s.Ip
}

func (s *DescribeDdosEventResponseBodyEvents) GetMbps() *int32 {
	return s.Mbps
}

func (s *DescribeDdosEventResponseBodyEvents) GetPps() *int32 {
	return s.Pps
}

func (s *DescribeDdosEventResponseBodyEvents) GetStartTime() *int32 {
	return s.StartTime
}

func (s *DescribeDdosEventResponseBodyEvents) GetStatus() *string {
	return s.Status
}

func (s *DescribeDdosEventResponseBodyEvents) SetEndTime(v int32) *DescribeDdosEventResponseBodyEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetIp(v string) *DescribeDdosEventResponseBodyEvents {
	s.Ip = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetMbps(v int32) *DescribeDdosEventResponseBodyEvents {
	s.Mbps = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetPps(v int32) *DescribeDdosEventResponseBodyEvents {
	s.Pps = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetStartTime(v int32) *DescribeDdosEventResponseBodyEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetStatus(v string) *DescribeDdosEventResponseBodyEvents {
	s.Status = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) Validate() error {
	return dara.Validate(s)
}
