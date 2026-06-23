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
	// The event details.
	Events []*DescribeEventsResponseBodyEvents `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// The query token. This value is the next_token value returned by the previous API call.
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
	// The event description.
	Data *DescribeEventsResponseBodyEventsData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The event ID.
	//
	// example:
	//
	// e-dba703c8-953b-40d8-82e8-cb713590****
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
	// np6a5c86f4ecae436f8f4a3dc034a7****
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// The event start time.
	//
	// example:
	//
	// 2025-04-23T20:48:01+08:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// The event type. Valid values:
	//
	// - `cluster_create`: creates a cluster.
	//
	// - `cluster_scaleout`: scales out a cluster.
	//
	// - `cluster_attach`: adds existing nodes.
	//
	// - `cluster_delete`: deletes a cluster.
	//
	// - `cluster_upgrade`: upgrades a cluster.
	//
	// - `cluster_migrate`: migrates a cluster.
	//
	// - `cluster_node_delete`: removes nodes.
	//
	// - `cluster_node_drain`: drains nodes.
	//
	// - `cluster_modify`: modifies a cluster.
	//
	// - `cluster_configuration_modify`: modifies cluster management configurations.
	//
	// - `cluster_addon_install`: installs a component.
	//
	// - `cluster_addon_upgrade`: upgrades a component.
	//
	// - `cluster_addon_uninstall`: uninstalls a component.
	//
	// - `runtime_upgrade`: upgrades the runtime.
	//
	// - `nodepool_upgrade`: upgrades a node pool.
	//
	// - `nodepool_update`: updates a node pool.
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
	// The event level. Valid values:
	//
	// - info: informational.
	//
	// - warning: warning.
	//
	// - error: error.
	//
	// example:
	//
	// info
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The event details.
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
	// The page number for the paged query.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The maximum number of results per page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of results.
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
