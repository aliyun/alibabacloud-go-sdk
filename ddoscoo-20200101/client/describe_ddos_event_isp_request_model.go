// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosEventIspRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventType(v string) *DescribeDDosEventIspRequest
	GetEventType() *string
	SetIp(v string) *DescribeDDosEventIspRequest
	GetIp() *string
	SetRange(v int64) *DescribeDDosEventIspRequest
	GetRange() *int64
	SetStartTime(v int64) *DescribeDDosEventIspRequest
	GetStartTime() *int64
}

type DescribeDDosEventIspRequest struct {
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
	// The attacked IP Address of Anti-DDoS Proxy instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 203.***.***.199
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The number of data entries to return.
	//
	// example:
	//
	// 100
	Range *int64 `json:"Range,omitempty" xml:"Range,omitempty"`
	// The start timestamp of the event to query. Unit: seconds.
	//
	// > You can call [DescribeDDosAllEventList](https://help.aliyun.com/document_detail/473646.html) to query the start time of all events.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1598948471
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDosEventIspRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventIspRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventIspRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeDDosEventIspRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeDDosEventIspRequest) GetRange() *int64 {
	return s.Range
}

func (s *DescribeDDosEventIspRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDDosEventIspRequest) SetEventType(v string) *DescribeDDosEventIspRequest {
	s.EventType = &v
	return s
}

func (s *DescribeDDosEventIspRequest) SetIp(v string) *DescribeDDosEventIspRequest {
	s.Ip = &v
	return s
}

func (s *DescribeDDosEventIspRequest) SetRange(v int64) *DescribeDDosEventIspRequest {
	s.Range = &v
	return s
}

func (s *DescribeDDosEventIspRequest) SetStartTime(v int64) *DescribeDDosEventIspRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDosEventIspRequest) Validate() error {
	return dara.Validate(s)
}
