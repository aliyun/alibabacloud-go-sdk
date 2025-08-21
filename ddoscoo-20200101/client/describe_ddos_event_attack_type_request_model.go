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
