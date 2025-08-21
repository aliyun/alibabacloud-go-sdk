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
	Ip    *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Range *int64  `json:"Range,omitempty" xml:"Range,omitempty"`
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
