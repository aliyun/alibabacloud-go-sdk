// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutonomousNotifyEventsInRangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetAutonomousNotifyEventsInRangeRequest
	GetEndTime() *string
	SetEventContext(v string) *GetAutonomousNotifyEventsInRangeRequest
	GetEventContext() *string
	SetInstanceId(v string) *GetAutonomousNotifyEventsInRangeRequest
	GetInstanceId() *string
	SetLevel(v string) *GetAutonomousNotifyEventsInRangeRequest
	GetLevel() *string
	SetMinLevel(v string) *GetAutonomousNotifyEventsInRangeRequest
	GetMinLevel() *string
	SetNodeId(v string) *GetAutonomousNotifyEventsInRangeRequest
	GetNodeId() *string
	SetPageOffset(v string) *GetAutonomousNotifyEventsInRangeRequest
	GetPageOffset() *string
	SetPageSize(v string) *GetAutonomousNotifyEventsInRangeRequest
	GetPageSize() *string
	SetStartTime(v string) *GetAutonomousNotifyEventsInRangeRequest
	GetStartTime() *string
	SetContext(v string) *GetAutonomousNotifyEventsInRangeRequest
	GetContext() *string
}

type GetAutonomousNotifyEventsInRangeRequest struct {
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1568265711221
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	EventContext *string `json:"EventContext,omitempty" xml:"EventContext,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-18ff4a195d****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The urgency level of the events. If you specify this parameter, the MinLevel parameter does not take effect. Valid values:
	//
	// 	- **Notice**: events for which the system sends notifications.
	//
	// 	- **Optimization**: events that need to be optimized.
	//
	// 	- **Warn**: events for which the system sends warnings.
	//
	// 	- **Critical**: critical events.
	//
	// example:
	//
	// Warn
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The minimum urgency level of the events. Valid values:
	//
	// 	- **Notice**: events for which the system sends notifications.
	//
	// 	- **Optimization**: events that need to be optimized.
	//
	// 	- **Warn**: events for which the system sends warnings.
	//
	// 	- **Critical**: critical events.
	//
	// example:
	//
	// Notice
	MinLevel *string `json:"MinLevel,omitempty" xml:"MinLevel,omitempty"`
	// The ID of the node in a PolarDB for MySQL cluster. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the node ID returned by the DBNodeId response parameter.
	//
	// >  You must specify the node ID if your database instance is a PolarDB for MySQL cluster.
	//
	// example:
	//
	// r-x****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The page number. The value must be a positive integer. Default value: 1.
	//
	// example:
	//
	// 1
	PageOffset *string `json:"PageOffset,omitempty" xml:"PageOffset,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1568269711000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	Context *string `json:"__context,omitempty" xml:"__context,omitempty"`
}

func (s GetAutonomousNotifyEventsInRangeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutonomousNotifyEventsInRangeRequest) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventsInRangeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetAutonomousNotifyEventsInRangeRequest) GetEventContext() *string {
	return s.EventContext
}

func (s *GetAutonomousNotifyEventsInRangeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAutonomousNotifyEventsInRangeRequest) GetLevel() *string {
	return s.Level
}

func (s *GetAutonomousNotifyEventsInRangeRequest) GetMinLevel() *string {
	return s.MinLevel
}

func (s *GetAutonomousNotifyEventsInRangeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetAutonomousNotifyEventsInRangeRequest) GetPageOffset() *string {
	return s.PageOffset
}

func (s *GetAutonomousNotifyEventsInRangeRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetAutonomousNotifyEventsInRangeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAutonomousNotifyEventsInRangeRequest) GetContext() *string {
	return s.Context
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetEndTime(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.EndTime = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetEventContext(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.EventContext = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetInstanceId(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetLevel(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.Level = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetMinLevel(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.MinLevel = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetNodeId(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.NodeId = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetPageOffset(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.PageOffset = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetPageSize(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.PageSize = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetStartTime(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.StartTime = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetContext(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.Context = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) Validate() error {
	return dara.Validate(s)
}
