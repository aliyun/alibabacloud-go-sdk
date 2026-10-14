// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosEventAttackTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventType(v string) *DescribeDDosEventAttackTypeRequest
	GetEventType() *string
	SetIp(v string) *DescribeDDosEventAttackTypeRequest
	GetIp() *string
	SetStartTime(v int64) *DescribeDDosEventAttackTypeRequest
	GetStartTime() *int64
}

type DescribeDDosEventAttackTypeRequest struct {
	// The type of the attack event to query. Valid values:
	//
	// - **defense**: volumetric attack scrubbing events.
	//
	// - **blackhole**: volumetric attack blackhole filtering events.
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

func (s DescribeDDosEventAttackTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventAttackTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAttackTypeRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeDDosEventAttackTypeRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeDDosEventAttackTypeRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDDosEventAttackTypeRequest) SetEventType(v string) *DescribeDDosEventAttackTypeRequest {
	s.EventType = &v
	return s
}

func (s *DescribeDDosEventAttackTypeRequest) SetIp(v string) *DescribeDDosEventAttackTypeRequest {
	s.Ip = &v
	return s
}

func (s *DescribeDDosEventAttackTypeRequest) SetStartTime(v int64) *DescribeDDosEventAttackTypeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDosEventAttackTypeRequest) Validate() error {
	return dara.Validate(s)
}
