// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*DescribeClusterEventsResponseBodyEvents) *DescribeClusterEventsResponseBody
	GetEvents() []*DescribeClusterEventsResponseBodyEvents
	SetPageInfo(v *DescribeClusterEventsResponseBodyPageInfo) *DescribeClusterEventsResponseBody
	GetPageInfo() *DescribeClusterEventsResponseBodyPageInfo
}

type DescribeClusterEventsResponseBody struct {
	// The list of events.
	Events []*DescribeClusterEventsResponseBodyEvents `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeClusterEventsResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
}

func (s DescribeClusterEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterEventsResponseBody) GetEvents() []*DescribeClusterEventsResponseBodyEvents {
	return s.Events
}

func (s *DescribeClusterEventsResponseBody) GetPageInfo() *DescribeClusterEventsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeClusterEventsResponseBody) SetEvents(v []*DescribeClusterEventsResponseBodyEvents) *DescribeClusterEventsResponseBody {
	s.Events = v
	return s
}

func (s *DescribeClusterEventsResponseBody) SetPageInfo(v *DescribeClusterEventsResponseBodyPageInfo) *DescribeClusterEventsResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeClusterEventsResponseBody) Validate() error {
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

type DescribeClusterEventsResponseBodyEvents struct {
	// The ID of the cluster.
	//
	// example:
	//
	// c82e6987e2961451182edacd74faf****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The description of the event.
	Data *DescribeClusterEventsResponseBodyEventsData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// cluster
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The subject related to the event.
	//
	// example:
	//
	// npdd89dc2b76c04f14b06774883b******
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// The time when the event started.
	//
	// example:
	//
	// 2020-12-01T17:31:00Z
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// The type of event. Valid values:
	//
	// example:
	//
	// nodepool_update
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeClusterEventsResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeClusterEventsResponseBodyEvents) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterEventsResponseBodyEvents) GetData() *DescribeClusterEventsResponseBodyEventsData {
	return s.Data
}

func (s *DescribeClusterEventsResponseBodyEvents) GetEventId() *string {
	return s.EventId
}

func (s *DescribeClusterEventsResponseBodyEvents) GetSource() *string {
	return s.Source
}

func (s *DescribeClusterEventsResponseBodyEvents) GetSubject() *string {
	return s.Subject
}

func (s *DescribeClusterEventsResponseBodyEvents) GetTime() *string {
	return s.Time
}

func (s *DescribeClusterEventsResponseBodyEvents) GetType() *string {
	return s.Type
}

func (s *DescribeClusterEventsResponseBodyEvents) SetClusterId(v string) *DescribeClusterEventsResponseBodyEvents {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterEventsResponseBodyEvents) SetData(v *DescribeClusterEventsResponseBodyEventsData) *DescribeClusterEventsResponseBodyEvents {
	s.Data = v
	return s
}

func (s *DescribeClusterEventsResponseBodyEvents) SetEventId(v string) *DescribeClusterEventsResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *DescribeClusterEventsResponseBodyEvents) SetSource(v string) *DescribeClusterEventsResponseBodyEvents {
	s.Source = &v
	return s
}

func (s *DescribeClusterEventsResponseBodyEvents) SetSubject(v string) *DescribeClusterEventsResponseBodyEvents {
	s.Subject = &v
	return s
}

func (s *DescribeClusterEventsResponseBodyEvents) SetTime(v string) *DescribeClusterEventsResponseBodyEvents {
	s.Time = &v
	return s
}

func (s *DescribeClusterEventsResponseBodyEvents) SetType(v string) *DescribeClusterEventsResponseBodyEvents {
	s.Type = &v
	return s
}

func (s *DescribeClusterEventsResponseBodyEvents) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterEventsResponseBodyEventsData struct {
	// The severity level of the event.
	//
	// Valid values:
	//
	// 	- warning
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- error
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- info
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// info
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The details of the event.
	//
	// example:
	//
	// Start to upgrade NodePool nodePool/npdd89dc2b76c04f14b06774883b******
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The status of the event.
	//
	// example:
	//
	// Started
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s DescribeClusterEventsResponseBodyEventsData) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterEventsResponseBodyEventsData) GoString() string {
	return s.String()
}

func (s *DescribeClusterEventsResponseBodyEventsData) GetLevel() *string {
	return s.Level
}

func (s *DescribeClusterEventsResponseBodyEventsData) GetMessage() *string {
	return s.Message
}

func (s *DescribeClusterEventsResponseBodyEventsData) GetReason() *string {
	return s.Reason
}

func (s *DescribeClusterEventsResponseBodyEventsData) SetLevel(v string) *DescribeClusterEventsResponseBodyEventsData {
	s.Level = &v
	return s
}

func (s *DescribeClusterEventsResponseBodyEventsData) SetMessage(v string) *DescribeClusterEventsResponseBodyEventsData {
	s.Message = &v
	return s
}

func (s *DescribeClusterEventsResponseBodyEventsData) SetReason(v string) *DescribeClusterEventsResponseBodyEventsData {
	s.Reason = &v
	return s
}

func (s *DescribeClusterEventsResponseBodyEventsData) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterEventsResponseBodyPageInfo struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 50.
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

func (s DescribeClusterEventsResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterEventsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterEventsResponseBodyPageInfo) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeClusterEventsResponseBodyPageInfo) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeClusterEventsResponseBodyPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeClusterEventsResponseBodyPageInfo) SetPageNumber(v int64) *DescribeClusterEventsResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterEventsResponseBodyPageInfo) SetPageSize(v int64) *DescribeClusterEventsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterEventsResponseBodyPageInfo) SetTotalCount(v int64) *DescribeClusterEventsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeClusterEventsResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
