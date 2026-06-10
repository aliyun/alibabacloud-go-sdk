// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*DescribeEventsResponseBodyEvents) *DescribeEventsResponseBody
	GetEvents() []*DescribeEventsResponseBodyEvents
	SetNextToken(v string) *DescribeEventsResponseBody
	GetNextToken() *string
	SetPageInfo(v *DescribeEventsResponseBodyPageInfo) *DescribeEventsResponseBody
	GetPageInfo() *DescribeEventsResponseBodyPageInfo
}

type DescribeEventsResponseBody struct {
	// A list of events.
	Events []*DescribeEventsResponseBodyEvents `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// The token to retrieve the next page of results. If this parameter is absent from the response, all results have been returned.
	//
	// example:
	//
	// 5c0a1c0f91c14c6****
	NextToken *string `json:"next_token,omitempty" xml:"next_token,omitempty"`
	// The pagination information.
	PageInfo *DescribeEventsResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
}

func (s DescribeEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBody) GetEvents() []*DescribeEventsResponseBodyEvents {
	return s.Events
}

func (s *DescribeEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeEventsResponseBody) GetPageInfo() *DescribeEventsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeEventsResponseBody) SetEvents(v []*DescribeEventsResponseBodyEvents) *DescribeEventsResponseBody {
	s.Events = v
	return s
}

func (s *DescribeEventsResponseBody) SetNextToken(v string) *DescribeEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeEventsResponseBody) SetPageInfo(v *DescribeEventsResponseBodyPageInfo) *DescribeEventsResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeEventsResponseBody) Validate() error {
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

type DescribeEventsResponseBodyEvents struct {
	// The cluster ID.
	//
	// example:
	//
	// cf62854ac2130470897be7a27ed1f****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The event details.
	Data *DescribeEventsResponseBodyEventsData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The event ID.
	//
	// example:
	//
	// e-dba703c8-953b-40d8-82e8-cb713590****
	EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
	// The source of the event.
	//
	// example:
	//
	// task
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The object that the event is about.
	//
	// example:
	//
	// np6a5c86f4ecae436f8f4a3dc034a7****
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// The event timestamp.
	//
	// example:
	//
	// 2025-04-23T20:48:01+08:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// The event type. Valid values:
	//
	// - `cluster_create`: A cluster is created.
	//
	// - `cluster_scaleout`: A cluster is scaled out.
	//
	// - `cluster_attach`: An existing node is added.
	//
	// - `cluster_delete`: A cluster is deleted.
	//
	// - `cluster_upgrade`: A cluster is upgraded.
	//
	// - `cluster_migrate`: A cluster is migrated.
	//
	// - `cluster_node_delete`: A node is removed.
	//
	// - `cluster_node_drain`: A node is drained.
	//
	// - `cluster_modify`: A cluster is modified.
	//
	// - `cluster_configuration_modify`: The control plane configuration of a cluster is modified.
	//
	// - `cluster_addon_install`: An add-on is installed.
	//
	// - `cluster_addon_upgrade`: An add-on is upgraded.
	//
	// - `cluster_addon_uninstall`: An add-on is uninstalled.
	//
	// - `runtime_upgrade`: The runtime is upgraded.
	//
	// - `nodepool_upgrade`: A node pool is upgraded.
	//
	// - `nodepool_update`: A node pool is updated.
	//
	// example:
	//
	// nodepool_upgrade
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeEventsResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyEvents) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeEventsResponseBodyEvents) GetData() *DescribeEventsResponseBodyEventsData {
	return s.Data
}

func (s *DescribeEventsResponseBodyEvents) GetEventId() *string {
	return s.EventId
}

func (s *DescribeEventsResponseBodyEvents) GetSource() *string {
	return s.Source
}

func (s *DescribeEventsResponseBodyEvents) GetSubject() *string {
	return s.Subject
}

func (s *DescribeEventsResponseBodyEvents) GetTime() *string {
	return s.Time
}

func (s *DescribeEventsResponseBodyEvents) GetType() *string {
	return s.Type
}

func (s *DescribeEventsResponseBodyEvents) SetClusterId(v string) *DescribeEventsResponseBodyEvents {
	s.ClusterId = &v
	return s
}

func (s *DescribeEventsResponseBodyEvents) SetData(v *DescribeEventsResponseBodyEventsData) *DescribeEventsResponseBodyEvents {
	s.Data = v
	return s
}

func (s *DescribeEventsResponseBodyEvents) SetEventId(v string) *DescribeEventsResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *DescribeEventsResponseBodyEvents) SetSource(v string) *DescribeEventsResponseBodyEvents {
	s.Source = &v
	return s
}

func (s *DescribeEventsResponseBodyEvents) SetSubject(v string) *DescribeEventsResponseBodyEvents {
	s.Subject = &v
	return s
}

func (s *DescribeEventsResponseBodyEvents) SetTime(v string) *DescribeEventsResponseBodyEvents {
	s.Time = &v
	return s
}

func (s *DescribeEventsResponseBodyEvents) SetType(v string) *DescribeEventsResponseBodyEvents {
	s.Type = &v
	return s
}

func (s *DescribeEventsResponseBodyEvents) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEventsResponseBodyEventsData struct {
	// The severity level of the event. Valid values:
	//
	// - info: An informational message.
	//
	// - warning: A warning.
	//
	// - error: An error.
	//
	// example:
	//
	// info
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// A human-readable description of the event.
	//
	// example:
	//
	// Start to upgrade NodePool nodePool/nodePool-A
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// A brief, machine-readable string that describes the reason for the event.
	//
	// example:
	//
	// Started
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s DescribeEventsResponseBodyEventsData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponseBodyEventsData) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyEventsData) GetLevel() *string {
	return s.Level
}

func (s *DescribeEventsResponseBodyEventsData) GetMessage() *string {
	return s.Message
}

func (s *DescribeEventsResponseBodyEventsData) GetReason() *string {
	return s.Reason
}

func (s *DescribeEventsResponseBodyEventsData) SetLevel(v string) *DescribeEventsResponseBodyEventsData {
	s.Level = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsData) SetMessage(v string) *DescribeEventsResponseBodyEventsData {
	s.Message = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsData) SetReason(v string) *DescribeEventsResponseBodyEventsData {
	s.Reason = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsData) Validate() error {
	return dara.Validate(s)
}

type DescribeEventsResponseBodyPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of entries that match the query.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s DescribeEventsResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyPageInfo) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeEventsResponseBodyPageInfo) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeEventsResponseBodyPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeEventsResponseBodyPageInfo) SetPageNumber(v int64) *DescribeEventsResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventsResponseBodyPageInfo) SetPageSize(v int64) *DescribeEventsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsResponseBodyPageInfo) SetTotalCount(v int64) *DescribeEventsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeEventsResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
