// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosEventSrcIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventType(v string) *DescribeDDosEventSrcIpRequest
	GetEventType() *string
	SetIp(v string) *DescribeDDosEventSrcIpRequest
	GetIp() *string
	SetRange(v int64) *DescribeDDosEventSrcIpRequest
	GetRange() *int64
	SetStartTime(v int64) *DescribeDDosEventSrcIpRequest
	GetStartTime() *int64
}

type DescribeDDosEventSrcIpRequest struct {
	// The type of the attack event to query. Valid values:
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
	// The attacked IP Address of Anti-DDoS Proxy instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 203.***.***.199
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The number of attack source IP addresses to return. The IP addresses are sorted by attack traffic in descending order. Default value: **5**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
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

func (s DescribeDDosEventSrcIpRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventSrcIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventSrcIpRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeDDosEventSrcIpRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeDDosEventSrcIpRequest) GetRange() *int64 {
	return s.Range
}

func (s *DescribeDDosEventSrcIpRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDDosEventSrcIpRequest) SetEventType(v string) *DescribeDDosEventSrcIpRequest {
	s.EventType = &v
	return s
}

func (s *DescribeDDosEventSrcIpRequest) SetIp(v string) *DescribeDDosEventSrcIpRequest {
	s.Ip = &v
	return s
}

func (s *DescribeDDosEventSrcIpRequest) SetRange(v int64) *DescribeDDosEventSrcIpRequest {
	s.Range = &v
	return s
}

func (s *DescribeDDosEventSrcIpRequest) SetStartTime(v int64) *DescribeDDosEventSrcIpRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDosEventSrcIpRequest) Validate() error {
	return dara.Validate(s)
}
