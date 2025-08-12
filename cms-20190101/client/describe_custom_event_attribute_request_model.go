// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomEventAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeCustomEventAttributeRequest
	GetEndTime() *string
	SetEventId(v string) *DescribeCustomEventAttributeRequest
	GetEventId() *string
	SetGroupId(v string) *DescribeCustomEventAttributeRequest
	GetGroupId() *string
	SetName(v string) *DescribeCustomEventAttributeRequest
	GetName() *string
	SetPageNumber(v int32) *DescribeCustomEventAttributeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCustomEventAttributeRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCustomEventAttributeRequest
	GetRegionId() *string
	SetSearchKeywords(v string) *DescribeCustomEventAttributeRequest
	GetSearchKeywords() *string
	SetStartTime(v string) *DescribeCustomEventAttributeRequest
	GetStartTime() *string
}

type DescribeCustomEventAttributeRequest struct {
	// The end of the time range to query.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1552227965971
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 123****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 123****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The event name.
	//
	// example:
	//
	// test123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The keywords that are used to search for the event.
	//
	// 	- If you need to query the custom event whose content contains a and b, set the value to a and b.
	//
	// 	- If you need to query the custom event whose content contains a or b, set the value to a or b.
	//
	// example:
	//
	// cms
	SearchKeywords *string `json:"SearchKeywords,omitempty" xml:"SearchKeywords,omitempty"`
	// The beginning of the time range to query.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1552224365971
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCustomEventAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEventAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventAttributeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCustomEventAttributeRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeCustomEventAttributeRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeCustomEventAttributeRequest) GetName() *string {
	return s.Name
}

func (s *DescribeCustomEventAttributeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCustomEventAttributeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCustomEventAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomEventAttributeRequest) GetSearchKeywords() *string {
	return s.SearchKeywords
}

func (s *DescribeCustomEventAttributeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCustomEventAttributeRequest) SetEndTime(v string) *DescribeCustomEventAttributeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) SetEventId(v string) *DescribeCustomEventAttributeRequest {
	s.EventId = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) SetGroupId(v string) *DescribeCustomEventAttributeRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) SetName(v string) *DescribeCustomEventAttributeRequest {
	s.Name = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) SetPageNumber(v int32) *DescribeCustomEventAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) SetPageSize(v int32) *DescribeCustomEventAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) SetRegionId(v string) *DescribeCustomEventAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) SetSearchKeywords(v string) *DescribeCustomEventAttributeRequest {
	s.SearchKeywords = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) SetStartTime(v string) *DescribeCustomEventAttributeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) Validate() error {
	return dara.Validate(s)
}
