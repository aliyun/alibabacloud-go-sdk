// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosAllEventListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeDDosAllEventListRequest
	GetEndTime() *int64
	SetEventType(v string) *DescribeDDosAllEventListRequest
	GetEventType() *string
	SetIp(v string) *DescribeDDosAllEventListRequest
	GetIp() *string
	SetPageNumber(v int32) *DescribeDDosAllEventListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDDosAllEventListRequest
	GetPageSize() *int32
	SetStartTime(v int64) *DescribeDDosAllEventListRequest
	GetStartTime() *int64
}

type DescribeDDosAllEventListRequest struct {
	// The end of the time range to query. DDoS attack events that occurred before **EndTime*	- are queried. Specify a UNIX timestamp. Unit: seconds.
	//
	// > EndTime can specify only events within the last three months.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1640966399
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of DDoS attack event to query. Valid values:
	//
	// - **web-cc**: web resource exhaustion attack.
	//
	// - **cc**: connection flood attack.
	//
	// - **defense**: volumetric attack (scrubbing event).
	//
	// - **blackhole**: volumetric attack (blackhole filtering event).
	//
	// You can specify multiple types. Separate them with commas (,).
	//
	// If you do not set this parameter, attack events of all types are queried.
	//
	// example:
	//
	// defense
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The IP address of the Anti-DDoS Pro or Anti-DDoS Premium instance to query.
	//
	// example:
	//
	// 192.168.1.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The page number to return in a paging query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of attack events to return on each page in a paging query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. DDoS attack events that occurred after **StartTime*	- are queried. Specify a UNIX timestamp. Unit: seconds.
	//
	// > StartTime can specify only events within the last three months.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1609430400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDosAllEventListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosAllEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDosAllEventListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDDosAllEventListRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeDDosAllEventListRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeDDosAllEventListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDDosAllEventListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDDosAllEventListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDDosAllEventListRequest) SetEndTime(v int64) *DescribeDDosAllEventListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDosAllEventListRequest) SetEventType(v string) *DescribeDDosAllEventListRequest {
	s.EventType = &v
	return s
}

func (s *DescribeDDosAllEventListRequest) SetIp(v string) *DescribeDDosAllEventListRequest {
	s.Ip = &v
	return s
}

func (s *DescribeDDosAllEventListRequest) SetPageNumber(v int32) *DescribeDDosAllEventListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDDosAllEventListRequest) SetPageSize(v int32) *DescribeDDosAllEventListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDDosAllEventListRequest) SetStartTime(v int64) *DescribeDDosAllEventListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDosAllEventListRequest) Validate() error {
	return dara.Validate(s)
}
