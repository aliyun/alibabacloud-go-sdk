// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosAllEventListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttackEvents(v []*DescribeDDosAllEventListResponseBodyAttackEvents) *DescribeDDosAllEventListResponseBody
	GetAttackEvents() []*DescribeDDosAllEventListResponseBodyAttackEvents
	SetRequestId(v string) *DescribeDDosAllEventListResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeDDosAllEventListResponseBody
	GetTotal() *int64
}

type DescribeDDosAllEventListResponseBody struct {
	// The DDoS attack events.
	AttackEvents []*DescribeDDosAllEventListResponseBodyAttackEvents `json:"AttackEvents,omitempty" xml:"AttackEvents,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25D83ED5-28CB-5683-9CF7-AECE521F3005
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of DDoS attack events.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDDosAllEventListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosAllEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDosAllEventListResponseBody) GetAttackEvents() []*DescribeDDosAllEventListResponseBodyAttackEvents {
	return s.AttackEvents
}

func (s *DescribeDDosAllEventListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDosAllEventListResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeDDosAllEventListResponseBody) SetAttackEvents(v []*DescribeDDosAllEventListResponseBodyAttackEvents) *DescribeDDosAllEventListResponseBody {
	s.AttackEvents = v
	return s
}

func (s *DescribeDDosAllEventListResponseBody) SetRequestId(v string) *DescribeDDosAllEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBody) SetTotal(v int64) *DescribeDDosAllEventListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBody) Validate() error {
	if s.AttackEvents != nil {
		for _, item := range s.AttackEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDDosAllEventListResponseBodyAttackEvents struct {
	// The source location or region from which the attack was initiated. Valid values:
	//
	// 	- **cn**: Chinese mainland
	//
	// 	- **alb-cn-hongkong-gf-2**: China (Hongkong)
	//
	// 	- **alb-us-west-1-gf-2**: US (Silicon Valley)
	//
	// 	- **alb-ap-northeast-1-gf-1**: Japan (Tokyo)
	//
	// 	- **alb-ap-southeast-gf-1**: Singapore
	//
	// 	- **alb-eu-central-1-gf-1**: Germany (Frankfurt)
	//
	// 	- **alb-eu-central-1-gf-1*	- or **selb-eu-west-1-gf-1a**: UK (London)
	//
	// 	- **alb-us-east-gf-1**: US (Virginia)
	//
	// 	- **CT-yundi**: China (Hongkong) This value is returned only for Anti-DDoS Premium instances of the Secure Chinese Mainland Acceleration (Sec-CMA) mitigation plan.
	//
	// example:
	//
	// cn
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The time when the DDoS attack stopped. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1634546030
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the DDoS attack event. Valid values:
	//
	// 	- **web-cc**: resource exhaustion attacks
	//
	// 	- **cc**: connection flood attacks
	//
	// 	- **defense**: DDoS attacks that trigger traffic scrubbing
	//
	// 	- **blackhole**: DDoS attacks that trigger blackhole filtering
	//
	// example:
	//
	// cc
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The attacked object. The attacked object varies based on the attack event type. The following list describes different attacked objects of different attack event types:
	//
	// 	- If **EventType*	- is set to **web-cc**, the value of this parameter indicates the domain name of the attacked website.
	//
	// 	- If **EventType*	- is set to **cc**, the value of this parameter indicates the IP address of the attacked Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// 	- If **EventType*	- is set to **defense*	- or **blackhole**, the value of this parameter indicates the IP address of the attacked Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// example:
	//
	// 203.107.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The peak bandwidth of the attack traffic. Unit: Mbit/s.
	//
	// example:
	//
	// 101899
	Mbps *int64 `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
	// The attacked port.
	//
	// > If **EventType*	- is set to **web-cc**, this parameter is not returned.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The peak packet forwarding rate of attack traffic. Unit: packets per second (pps).
	//
	// example:
	//
	// 9664270
	Pps *int64 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The time when the DDoS attack started. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1634543764
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDosAllEventListResponseBodyAttackEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosAllEventListResponseBodyAttackEvents) GoString() string {
	return s.String()
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) GetArea() *string {
	return s.Area
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) GetEventType() *string {
	return s.EventType
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) GetIp() *string {
	return s.Ip
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) GetMbps() *int64 {
	return s.Mbps
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) GetPort() *string {
	return s.Port
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) GetPps() *int64 {
	return s.Pps
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) SetArea(v string) *DescribeDDosAllEventListResponseBodyAttackEvents {
	s.Area = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) SetEndTime(v int64) *DescribeDDosAllEventListResponseBodyAttackEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) SetEventType(v string) *DescribeDDosAllEventListResponseBodyAttackEvents {
	s.EventType = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) SetIp(v string) *DescribeDDosAllEventListResponseBodyAttackEvents {
	s.Ip = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) SetMbps(v int64) *DescribeDDosAllEventListResponseBodyAttackEvents {
	s.Mbps = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) SetPort(v string) *DescribeDDosAllEventListResponseBodyAttackEvents {
	s.Port = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) SetPps(v int64) *DescribeDDosAllEventListResponseBodyAttackEvents {
	s.Pps = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) SetStartTime(v int64) *DescribeDDosAllEventListResponseBodyAttackEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) Validate() error {
	return dara.Validate(s)
}
