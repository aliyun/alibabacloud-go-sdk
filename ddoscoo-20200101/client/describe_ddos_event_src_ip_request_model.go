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
	// The type of the attack event that you want to query. Valid values:
	//
	// 	- **defense**: attack events that trigger traffic scrubbing
	//
	// 	- **blackhole**: attack events that trigger blackhole filtering
	//
	// This parameter is required.
	//
	// example:
	//
	// defense
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The IP address of the attacked Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 203.***.***.199
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The number of source IP addresses that you want to return. The source IP addresses are returned in descending order of attack traffic. By default, the top **five*	- source IP addresses are returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Range *int64 `json:"Range,omitempty" xml:"Range,omitempty"`
	// The UNIX timestamp when the query starts. Unit: seconds.
	//
	// > You can call the [DescribeDDosAllEventList](https://help.aliyun.com/document_detail/188604.html) operation to query the beginning time of all attack events.
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
