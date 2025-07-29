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
	SetPageInfo(v *DescribeEventsForRegionResponseBodyPageInfo) *DescribeEventsForRegionResponseBody
	GetPageInfo() *DescribeEventsForRegionResponseBodyPageInfo
}

type DescribeEventsForRegionResponseBody struct {
	// The events.
	Events []*DescribeEventsForRegionResponseBodyEvents `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// The pagination details.
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

func (s *DescribeEventsForRegionResponseBody) GetPageInfo() *DescribeEventsForRegionResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeEventsForRegionResponseBody) SetEvents(v []*DescribeEventsForRegionResponseBodyEvents) *DescribeEventsForRegionResponseBody {
	s.Events = v
	return s
}

func (s *DescribeEventsForRegionResponseBody) SetPageInfo(v *DescribeEventsForRegionResponseBodyPageInfo) *DescribeEventsForRegionResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeEventsForRegionResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventsForRegionResponseBodyEvents struct {
	// The cluster ID.
	//
	// example:
	//
	// cluster-id
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The description of the event.
	Data *DescribeEventsForRegionResponseBodyEventsData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The event ID.
	//
	// example:
	//
	// A234-1234-1234
	EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
	// The event source.
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The object associated with the event.
	//
	// example:
	//
	// nodePool-id
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// The time when the event was generated.
	//
	// example:
	//
	// 2020-12-01T17:31:00Z
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
	return dara.Validate(s)
}

type DescribeEventsForRegionResponseBodyEventsData struct {
	// The severity level of the event.
	//
	// example:
	//
	// info
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The details of the event.
	//
	// example:
	//
	// Start to upgrade NodePool nodePool/nodePool-A
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The event status.
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
	// The number of pages.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of records on each page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of entries returned.
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
