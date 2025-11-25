// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainAttackEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainAttackEvents(v []*DescribeDomainAttackEventsResponseBodyDomainAttackEvents) *DescribeDomainAttackEventsResponseBody
	GetDomainAttackEvents() []*DescribeDomainAttackEventsResponseBodyDomainAttackEvents
	SetRequestId(v string) *DescribeDomainAttackEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDomainAttackEventsResponseBody
	GetTotalCount() *int64
}

type DescribeDomainAttackEventsResponseBody struct {
	// An array that consists of the details of the DDoS attack event.
	DomainAttackEvents []*DescribeDomainAttackEventsResponseBodyDomainAttackEvents `json:"DomainAttackEvents,omitempty" xml:"DomainAttackEvents,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned DDoS attack events.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDomainAttackEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAttackEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventsResponseBody) GetDomainAttackEvents() []*DescribeDomainAttackEventsResponseBodyDomainAttackEvents {
	return s.DomainAttackEvents
}

func (s *DescribeDomainAttackEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainAttackEventsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDomainAttackEventsResponseBody) SetDomainAttackEvents(v []*DescribeDomainAttackEventsResponseBodyDomainAttackEvents) *DescribeDomainAttackEventsResponseBody {
	s.DomainAttackEvents = v
	return s
}

func (s *DescribeDomainAttackEventsResponseBody) SetRequestId(v string) *DescribeDomainAttackEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBody) SetTotalCount(v int64) *DescribeDomainAttackEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBody) Validate() error {
	if s.DomainAttackEvents != nil {
		for _, item := range s.DomainAttackEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainAttackEventsResponseBodyDomainAttackEvents struct {
	// The attacked domain name.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The time when the DDoS attack stopped. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1560320160
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The peak attack QPS.
	//
	// example:
	//
	// 1000
	MaxQps *int64 `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
	// The time when the DDoS attack started. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1560312900
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainAttackEventsResponseBodyDomainAttackEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAttackEventsResponseBodyDomainAttackEvents) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventsResponseBodyDomainAttackEvents) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainAttackEventsResponseBodyDomainAttackEvents) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainAttackEventsResponseBodyDomainAttackEvents) GetMaxQps() *int64 {
	return s.MaxQps
}

func (s *DescribeDomainAttackEventsResponseBodyDomainAttackEvents) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainAttackEventsResponseBodyDomainAttackEvents) SetDomain(v string) *DescribeDomainAttackEventsResponseBodyDomainAttackEvents {
	s.Domain = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBodyDomainAttackEvents) SetEndTime(v int64) *DescribeDomainAttackEventsResponseBodyDomainAttackEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBodyDomainAttackEvents) SetMaxQps(v int64) *DescribeDomainAttackEventsResponseBodyDomainAttackEvents {
	s.MaxQps = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBodyDomainAttackEvents) SetStartTime(v int64) *DescribeDomainAttackEventsResponseBodyDomainAttackEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBodyDomainAttackEvents) Validate() error {
	return dara.Validate(s)
}
