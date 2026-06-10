// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsForRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*DescribeEventsForRegionResponseBodyEvents) *DescribeEventsForRegionResponseBody
	GetEvents() []*DescribeEventsForRegionResponseBodyEvents
	SetNextToken(v string) *DescribeEventsForRegionResponseBody
	GetNextToken() *string
	SetPageInfo(v *DescribeEventsForRegionResponseBodyPageInfo) *DescribeEventsForRegionResponseBody
	GetPageInfo() *DescribeEventsForRegionResponseBodyPageInfo
}

type DescribeEventsForRegionResponseBody struct {
	// A list of events.
	Events []*DescribeEventsForRegionResponseBodyEvents `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results. If this parameter is empty, all results have been returned.
	NextToken *string `json:"next_token,omitempty" xml:"next_token,omitempty"`
	// The pagination information.
	PageInfo *DescribeEventsForRegionResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
}

func (s DescribeEventsForRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsForRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventsForRegionResponseBody) GetEvents() []*DescribeEventsForRegionResponseBodyEvents {
	return s.Events
}

func (s *DescribeEventsForRegionResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeEventsForRegionResponseBody) GetPageInfo() *DescribeEventsForRegionResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeEventsForRegionResponseBody) SetEvents(v []*DescribeEventsForRegionResponseBodyEvents) *DescribeEventsForRegionResponseBody {
	s.Events = v
	return s
}

func (s *DescribeEventsForRegionResponseBody) SetNextToken(v string) *DescribeEventsForRegionResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeEventsForRegionResponseBody) SetPageInfo(v *DescribeEventsForRegionResponseBodyPageInfo) *DescribeEventsForRegionResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeEventsForRegionResponseBody) Validate() error {
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEventsForRegionResponseBodyEvents struct {
	// The cluster ID.
	//
	// example:
	//
	// c82e6987e2961451182edacd74faf****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The event data.
	Data *DescribeEventsForRegionResponseBodyEventsData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The event ID.
	//
	// example:
	//
	// e-9ad04f72-8ee7-46bf-a02c-e4a06b39****
	EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
	// The event source.
	//
	// example:
	//
	// task
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The event subject.
	//
	// example:
	//
	// npdd89dc2b76c04f14b06774883b******
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// The time when the event occurred.
	//
	// example:
	//
	// 2025-05-14T10:00:56+08:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// The event type.
	//
	// example:
	//
	// nodePool_upgrade
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeEventsForRegionResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsForRegionResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeEventsForRegionResponseBodyEvents) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeEventsForRegionResponseBodyEvents) GetData() *DescribeEventsForRegionResponseBodyEventsData {
	return s.Data
}

func (s *DescribeEventsForRegionResponseBodyEvents) GetEventId() *string {
	return s.EventId
}

func (s *DescribeEventsForRegionResponseBodyEvents) GetSource() *string {
	return s.Source
}

func (s *DescribeEventsForRegionResponseBodyEvents) GetSubject() *string {
	return s.Subject
}

func (s *DescribeEventsForRegionResponseBodyEvents) GetTime() *string {
	return s.Time
}

func (s *DescribeEventsForRegionResponseBodyEvents) GetType() *string {
	return s.Type
}

func (s *DescribeEventsForRegionResponseBodyEvents) SetClusterId(v string) *DescribeEventsForRegionResponseBodyEvents {
	s.ClusterId = &v
	return s
}

func (s *DescribeEventsForRegionResponseBodyEvents) SetData(v *DescribeEventsForRegionResponseBodyEventsData) *DescribeEventsForRegionResponseBodyEvents {
	s.Data = v
	return s
}

func (s *DescribeEventsForRegionResponseBodyEvents) SetEventId(v string) *DescribeEventsForRegionResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *DescribeEventsForRegionResponseBodyEvents) SetSource(v string) *DescribeEventsForRegionResponseBodyEvents {
	s.Source = &v
	return s
}

func (s *DescribeEventsForRegionResponseBodyEvents) SetSubject(v string) *DescribeEventsForRegionResponseBodyEvents {
	s.Subject = &v
	return s
}

func (s *DescribeEventsForRegionResponseBodyEvents) SetTime(v string) *DescribeEventsForRegionResponseBodyEvents {
	s.Time = &v
	return s
}

func (s *DescribeEventsForRegionResponseBodyEvents) SetType(v string) *DescribeEventsForRegionResponseBodyEvents {
	s.Type = &v
	return s
}

func (s *DescribeEventsForRegionResponseBodyEvents) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEventsForRegionResponseBodyEventsData struct {
	// The event level.
	//
	// example:
	//
	// info
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The event message.
	//
	// example:
	//
	// Start to upgrade NodePool nodePool/nodePool-A
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The reason for the event.
	//
	// example:
	//
	// Started
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s DescribeEventsForRegionResponseBodyEventsData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsForRegionResponseBodyEventsData) GoString() string {
	return s.String()
}

func (s *DescribeEventsForRegionResponseBodyEventsData) GetLevel() *string {
	return s.Level
}

func (s *DescribeEventsForRegionResponseBodyEventsData) GetMessage() *string {
	return s.Message
}

func (s *DescribeEventsForRegionResponseBodyEventsData) GetReason() *string {
	return s.Reason
}

func (s *DescribeEventsForRegionResponseBodyEventsData) SetLevel(v string) *DescribeEventsForRegionResponseBodyEventsData {
	s.Level = &v
	return s
}

func (s *DescribeEventsForRegionResponseBodyEventsData) SetMessage(v string) *DescribeEventsForRegionResponseBodyEventsData {
	s.Message = &v
	return s
}

func (s *DescribeEventsForRegionResponseBodyEventsData) SetReason(v string) *DescribeEventsForRegionResponseBodyEventsData {
	s.Reason = &v
	return s
}

func (s *DescribeEventsForRegionResponseBodyEventsData) Validate() error {
	return dara.Validate(s)
}

type DescribeEventsForRegionResponseBodyPageInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s DescribeEventsForRegionResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsForRegionResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeEventsForRegionResponseBodyPageInfo) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeEventsForRegionResponseBodyPageInfo) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeEventsForRegionResponseBodyPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeEventsForRegionResponseBodyPageInfo) SetPageNumber(v int64) *DescribeEventsForRegionResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventsForRegionResponseBodyPageInfo) SetPageSize(v int64) *DescribeEventsForRegionResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsForRegionResponseBodyPageInfo) SetTotalCount(v int64) *DescribeEventsForRegionResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeEventsForRegionResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
