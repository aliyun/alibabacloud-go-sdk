// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomEventCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeCustomEventCountRequest
	GetEndTime() *string
	SetEventId(v string) *DescribeCustomEventCountRequest
	GetEventId() *string
	SetGroupId(v string) *DescribeCustomEventCountRequest
	GetGroupId() *string
	SetName(v string) *DescribeCustomEventCountRequest
	GetName() *string
	SetRegionId(v string) *DescribeCustomEventCountRequest
	GetRegionId() *string
	SetSearchKeywords(v string) *DescribeCustomEventCountRequest
	GetSearchKeywords() *string
	SetStartTime(v string) *DescribeCustomEventCountRequest
	GetStartTime() *string
}

type DescribeCustomEventCountRequest struct {
	// The end of the time range to query.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1552220485596
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 123
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 12345
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The event name.
	//
	// example:
	//
	// BABEL_BUY
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// 1552209685596
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCustomEventCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEventCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventCountRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCustomEventCountRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeCustomEventCountRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeCustomEventCountRequest) GetName() *string {
	return s.Name
}

func (s *DescribeCustomEventCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomEventCountRequest) GetSearchKeywords() *string {
	return s.SearchKeywords
}

func (s *DescribeCustomEventCountRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCustomEventCountRequest) SetEndTime(v string) *DescribeCustomEventCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCustomEventCountRequest) SetEventId(v string) *DescribeCustomEventCountRequest {
	s.EventId = &v
	return s
}

func (s *DescribeCustomEventCountRequest) SetGroupId(v string) *DescribeCustomEventCountRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeCustomEventCountRequest) SetName(v string) *DescribeCustomEventCountRequest {
	s.Name = &v
	return s
}

func (s *DescribeCustomEventCountRequest) SetRegionId(v string) *DescribeCustomEventCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomEventCountRequest) SetSearchKeywords(v string) *DescribeCustomEventCountRequest {
	s.SearchKeywords = &v
	return s
}

func (s *DescribeCustomEventCountRequest) SetStartTime(v string) *DescribeCustomEventCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCustomEventCountRequest) Validate() error {
	return dara.Validate(s)
}
