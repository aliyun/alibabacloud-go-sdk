// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*DescribeDDoSEventsResponseBodyEvents) *DescribeDDoSEventsResponseBody
	GetEvents() []*DescribeDDoSEventsResponseBodyEvents
	SetRequestId(v string) *DescribeDDoSEventsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeDDoSEventsResponseBody
	GetTotal() *int64
}

type DescribeDDoSEventsResponseBody struct {
	Events []*DescribeDDoSEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DescribeDDoSEventsResponseBody) GetEvents() []*DescribeDDoSEventsResponseBodyEvents {
	return s.Events
}

func (s *DescribeDDoSEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDoSEventsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeDDoSEventsResponseBody) SetEvents(v []*DescribeDDoSEventsResponseBodyEvents) *DescribeDDoSEventsResponseBody {
	s.Events = v
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
	return dara.Validate(s)
}

type DescribeDDoSEventsResponseBodyEvents struct {
	// example:
	//
	// 3289457398
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 12
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 3289457324
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// blackhole_start
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDDoSEventsResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsResponseBodyEvents) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDDoSEventsResponseBodyEvents) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeDDoSEventsResponseBodyEvents) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDDoSEventsResponseBodyEvents) GetStatus() *string {
	return s.Status
}

func (s *DescribeDDoSEventsResponseBodyEvents) SetEndTime(v int64) *DescribeDDoSEventsResponseBodyEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyEvents) SetInterval(v int32) *DescribeDDoSEventsResponseBodyEvents {
	s.Interval = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyEvents) SetStartTime(v int64) *DescribeDDoSEventsResponseBodyEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyEvents) SetStatus(v string) *DescribeDDoSEventsResponseBodyEvents {
	s.Status = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyEvents) Validate() error {
	return dara.Validate(s)
}
