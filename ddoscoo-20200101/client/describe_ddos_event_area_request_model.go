// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosEventAreaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventType(v string) *DescribeDDosEventAreaRequest
	GetEventType() *string
	SetIp(v string) *DescribeDDosEventAreaRequest
	GetIp() *string
	SetRange(v int64) *DescribeDDosEventAreaRequest
	GetRange() *int64
	SetStartTime(v int64) *DescribeDDosEventAreaRequest
	GetStartTime() *int64
}

type DescribeDDosEventAreaRequest struct {
	// The event type of the attack event to query. Valid values:
	//
	// - **defense**: traffic scrubbing events for volumetric attacks.
	//
	// - **blackhole**: blackhole filtering events for volumetric attacks.
	//
	// This parameter is required.
	//
	// example:
	//
	// defense
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The attacked IP Address of Anti-DDoS Proxy instance of Anti-DDoS Pro and Anti-DDoS Premium.
	//
	// This parameter is required.
	//
	// example:
	//
	// 203.XX.XX.199
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The number of data entries to return. Unit: entries.
	//
	// example:
	//
	// 100
	Range *int64 `json:"Range,omitempty" xml:"Range,omitempty"`
	// The start timestamp of the event to query. Unit: seconds.
	//
	// > You can call [DescribeDDosAllEventList](https://help.aliyun.com/document_detail/188604.html) to query the start time of all events.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1598948471
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDosEventAreaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventAreaRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAreaRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeDDosEventAreaRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeDDosEventAreaRequest) GetRange() *int64 {
	return s.Range
}

func (s *DescribeDDosEventAreaRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDDosEventAreaRequest) SetEventType(v string) *DescribeDDosEventAreaRequest {
	s.EventType = &v
	return s
}

func (s *DescribeDDosEventAreaRequest) SetIp(v string) *DescribeDDosEventAreaRequest {
	s.Ip = &v
	return s
}

func (s *DescribeDDosEventAreaRequest) SetRange(v int64) *DescribeDDosEventAreaRequest {
	s.Range = &v
	return s
}

func (s *DescribeDDosEventAreaRequest) SetStartTime(v int64) *DescribeDDosEventAreaRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDosEventAreaRequest) Validate() error {
	return dara.Validate(s)
}
