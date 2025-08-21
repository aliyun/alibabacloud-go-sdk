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
	// The end of the time range to query. The DDoS attack events occur before **EndTime*	- are queried. This value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1640966399
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the DDoS attack events you want to query. Valid values:
	//
	// 	- **web-cc**: resource exhaustion attacks
	//
	// 	- **cc**: connection flood attacks
	//
	// 	- **defense**: DDoS attacks that trigger traffic scrubbing
	//
	// 	- **blackhole**: DDoS attacks that trigger blackhole filtering
	//
	// If you want to query multiple types of DDoS attack events, separate them with commas (,).
	//
	// If you do not configure this parameter, DDoS attack events of all types are queried.
	//
	// example:
	//
	// defense
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. The DDoS attack events occur after **StartTime*	- are queried. This value is a UNIX timestamp. Unit: seconds.
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
