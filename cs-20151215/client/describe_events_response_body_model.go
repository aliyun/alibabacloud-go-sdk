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
	SetPageInfo(v *DescribeEventsResponseBodyPageInfo) *DescribeEventsResponseBody
	GetPageInfo() *DescribeEventsResponseBodyPageInfo
}

type DescribeEventsResponseBody struct {
	// The details of the events.
	Events []*DescribeEventsResponseBodyEvents `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
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

func (s *DescribeEventsResponseBody) GetPageInfo() *DescribeEventsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeEventsResponseBody) SetEvents(v []*DescribeEventsResponseBodyEvents) *DescribeEventsResponseBody {
	s.Events = v
	return s
}

func (s *DescribeEventsResponseBody) SetPageInfo(v *DescribeEventsResponseBodyPageInfo) *DescribeEventsResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeEventsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventsResponseBodyEvents struct {
	// The ID of the cluster.
	//
	// example:
	//
	// cluster-id
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The description of the event.
	Data *DescribeEventsResponseBodyEventsData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The event ID.
	//
	// example:
	//
	// A234-1234-1234
	EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
	// The source of the event.
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The subject of the event.
	//
	// example:
	//
	// nodePool-id
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// The time when the event started.
	//
	// example:
	//
	// 2022-11-23T20:48:01+08:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// The event type. Valid values:
	//
	// 	- `cluster_create`: cluster creation.
	//
	// 	- `cluster_scaleout`: cluster scale-out.
	//
	// 	- `cluster_attach`: node addition.
	//
	// 	- `cluster_delete`: cluster deletion.
	//
	// 	- `cluster_upgrade`: cluster upgrades.
	//
	// 	- `cluster_migrate`: cluster migration.
	//
	// 	- `cluster_node_delete`: node removal.
	//
	// 	- `cluster_node_drain`: node draining.
	//
	// 	- `cluster_modify`: cluster modifications.
	//
	// 	- `cluster_configuration_modify`: modifications of control plane configurations.
	//
	// 	- `cluster_addon_install`: component installation.
	//
	// 	- `cluster_addon_upgrade`: component updates.
	//
	// 	- `cluster_addon_uninstall`: component uninstallation.
	//
	// 	- `runtime_upgrade`: runtime updates.
	//
	// 	- `nodepool_upgrade`: node pool upgrades.
	//
	// 	- `nodepool_update`: node pool updates.
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
	return dara.Validate(s)
}

type DescribeEventsResponseBodyEventsData struct {
	// The severity level of the event. Valid values:
	//
	// 	- info
	//
	// 	- warning
	//
	// 	- error
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
	// The status of the event.
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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page.
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
