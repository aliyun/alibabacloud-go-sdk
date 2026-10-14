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
	// The list of attack events.
	AttackEvents []*DescribeDDosAllEventListResponseBodyAttackEvents `json:"AttackEvents,omitempty" xml:"AttackEvents,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 25D83ED5-28CB-5683-9CF7-AECE521F3005
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of attack events returned.
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
	// The source region of the attack. Valid values:
	//
	// - **cn**: the Chinese mainland.
	//
	// - **alb-cn-hongkong-gf-2**: Hong Kong (China).
	//
	// - **alb-us-west-1-gf-2**: US (Silicon Valley).
	//
	// - **alb-ap-northeast-1-gf-1**: Japan (Tokyo).
	//
	// - **alb-ap-southeast-gf-1**: Singapore.
	//
	// - **alb-eu-central-1-gf-1**: Germany (Frankfurt).
	//
	// - **alb-eu-west-1-gf-1*	- and **selb-eu-west-1-gf-1a**: UK (London).
	//
	// - **alb-us-east-gf-1**: US (Virginia).
	//
	// - **CT-yundi**: Hong Kong (China). This value applies only to Anti-DDoS Pro and Anti-DDoS Premium secure acceleration line instances.
	//
	// example:
	//
	// cn
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The end time of the attack. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1634546030
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the DDoS attack event. Valid values:
	//
	// - **web-cc**: web resource exhaustion attack.
	//
	// - **cc**: connection flood attack.
	//
	// - **defense**: volumetric attack (scrubbing event).
	//
	// - **blackhole**: volumetric attack (blackhole filtering event).
	//
	// example:
	//
	// cc
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The attacked object. The attacked object varies based on the attack event type:
	//
	// - Web resource exhaustion attack (**EventType*	- is **web-cc**): this parameter indicates the attacked website domain name.
	//
	// - Connection flood attack (**EventType*	- is **cc**): this parameter indicates the attacked IP address of the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// - Volumetric attack (**EventType*	- is **defense*	- or **blackhole**): this parameter indicates the attacked IP address of the Anti-DDoS Pro or Anti-DDoS Premium instance.
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
	// The attacked port number.
	//
	// > This parameter is not returned for web resource exhaustion attacks (**EventType*	- is **web-cc**).
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The peak packet forwarding rate of the attack traffic. Unit: pps.
	//
	// example:
	//
	// 9664270
	Pps *int64 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The start time of the attack. The value is a UNIX timestamp. Unit: seconds.
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
