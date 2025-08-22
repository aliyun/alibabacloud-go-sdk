// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosAllEventListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDdosAllEventListRequest
	GetEndTime() *string
	SetEventType(v string) *DescribeDdosAllEventListRequest
	GetEventType() *string
	SetPageNumber(v int32) *DescribeDdosAllEventListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDdosAllEventListRequest
	GetPageSize() *int32
	SetStartTime(v string) *DescribeDdosAllEventListRequest
	GetStartTime() *string
}

type DescribeDdosAllEventListRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The end time must be later than the start time. The maximum time range is 31 days.
	//
	// example:
	//
	// 2023-04-25T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the DDoS attack event to be queried. Valid values:
	//
	//  	- **web-cc**: resource exhaustion attacks
	//
	//  	- **cc**: connection flood attacks
	//
	//  	- **traffic**: volumetric attacks
	//
	// If you do not configure this parameter, DDoS attack events of all types are queried.
	//
	// example:
	//
	// web-cc
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The page number. Default value: 1. Value range: 1 to 10,000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Valid values: 5, 10, and 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-03-27T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDdosAllEventListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosAllEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosAllEventListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDdosAllEventListRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeDdosAllEventListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDdosAllEventListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDdosAllEventListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDdosAllEventListRequest) SetEndTime(v string) *DescribeDdosAllEventListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosAllEventListRequest) SetEventType(v string) *DescribeDdosAllEventListRequest {
	s.EventType = &v
	return s
}

func (s *DescribeDdosAllEventListRequest) SetPageNumber(v int32) *DescribeDdosAllEventListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDdosAllEventListRequest) SetPageSize(v int32) *DescribeDdosAllEventListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDdosAllEventListRequest) SetStartTime(v string) *DescribeDdosAllEventListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosAllEventListRequest) Validate() error {
	return dara.Validate(s)
}
