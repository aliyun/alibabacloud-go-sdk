// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomEventHistogramRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeCustomEventHistogramRequest
	GetEndTime() *string
	SetEventId(v string) *DescribeCustomEventHistogramRequest
	GetEventId() *string
	SetGroupId(v string) *DescribeCustomEventHistogramRequest
	GetGroupId() *string
	SetLevel(v string) *DescribeCustomEventHistogramRequest
	GetLevel() *string
	SetName(v string) *DescribeCustomEventHistogramRequest
	GetName() *string
	SetRegionId(v string) *DescribeCustomEventHistogramRequest
	GetRegionId() *string
	SetSearchKeywords(v string) *DescribeCustomEventHistogramRequest
	GetSearchKeywords() *string
	SetStartTime(v string) *DescribeCustomEventHistogramRequest
	GetStartTime() *string
}

type DescribeCustomEventHistogramRequest struct {
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
	// The severity level of the event. Valid values:
	//
	// 	- CRITICAL
	//
	// 	- WARN
	//
	// 	- INFO
	//
	// example:
	//
	// CRITICAL
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The event name.
	//
	// example:
	//
	// BucketIngressBandwidth
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The keywords that are used to search for the event.
	//
	// 	- If you need to query the custom event whose content contains a and b, set the value to "a and b".
	//
	// 	- If you need to query the custom event whose content contains a or b, set the value to "a or b".
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

func (s DescribeCustomEventHistogramRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEventHistogramRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventHistogramRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCustomEventHistogramRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeCustomEventHistogramRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeCustomEventHistogramRequest) GetLevel() *string {
	return s.Level
}

func (s *DescribeCustomEventHistogramRequest) GetName() *string {
	return s.Name
}

func (s *DescribeCustomEventHistogramRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomEventHistogramRequest) GetSearchKeywords() *string {
	return s.SearchKeywords
}

func (s *DescribeCustomEventHistogramRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCustomEventHistogramRequest) SetEndTime(v string) *DescribeCustomEventHistogramRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCustomEventHistogramRequest) SetEventId(v string) *DescribeCustomEventHistogramRequest {
	s.EventId = &v
	return s
}

func (s *DescribeCustomEventHistogramRequest) SetGroupId(v string) *DescribeCustomEventHistogramRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeCustomEventHistogramRequest) SetLevel(v string) *DescribeCustomEventHistogramRequest {
	s.Level = &v
	return s
}

func (s *DescribeCustomEventHistogramRequest) SetName(v string) *DescribeCustomEventHistogramRequest {
	s.Name = &v
	return s
}

func (s *DescribeCustomEventHistogramRequest) SetRegionId(v string) *DescribeCustomEventHistogramRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomEventHistogramRequest) SetSearchKeywords(v string) *DescribeCustomEventHistogramRequest {
	s.SearchKeywords = &v
	return s
}

func (s *DescribeCustomEventHistogramRequest) SetStartTime(v string) *DescribeCustomEventHistogramRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCustomEventHistogramRequest) Validate() error {
	return dara.Validate(s)
}
