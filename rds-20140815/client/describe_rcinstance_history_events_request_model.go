// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceHistoryEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventPublishTime(v *DescribeRCInstanceHistoryEventsRequestEventPublishTime) *DescribeRCInstanceHistoryEventsRequest
	GetEventPublishTime() *DescribeRCInstanceHistoryEventsRequestEventPublishTime
	SetNotBefore(v *DescribeRCInstanceHistoryEventsRequestNotBefore) *DescribeRCInstanceHistoryEventsRequest
	GetNotBefore() *DescribeRCInstanceHistoryEventsRequestNotBefore
	SetEventCycleStatus(v string) *DescribeRCInstanceHistoryEventsRequest
	GetEventCycleStatus() *string
	SetEventId(v []*string) *DescribeRCInstanceHistoryEventsRequest
	GetEventId() []*string
	SetEventType(v string) *DescribeRCInstanceHistoryEventsRequest
	GetEventType() *string
	SetImpactLevel(v string) *DescribeRCInstanceHistoryEventsRequest
	GetImpactLevel() *string
	SetInstanceEventCycleStatus(v []*string) *DescribeRCInstanceHistoryEventsRequest
	GetInstanceEventCycleStatus() []*string
	SetInstanceEventType(v []*string) *DescribeRCInstanceHistoryEventsRequest
	GetInstanceEventType() []*string
	SetInstanceId(v string) *DescribeRCInstanceHistoryEventsRequest
	GetInstanceId() *string
	SetMaxResults(v string) *DescribeRCInstanceHistoryEventsRequest
	GetMaxResults() *string
	SetPageNumber(v string) *DescribeRCInstanceHistoryEventsRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeRCInstanceHistoryEventsRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeRCInstanceHistoryEventsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeRCInstanceHistoryEventsRequest
	GetResourceGroupId() *string
	SetResourceId(v []*string) *DescribeRCInstanceHistoryEventsRequest
	GetResourceId() []*string
	SetTag(v []*DescribeRCInstanceHistoryEventsRequestTag) *DescribeRCInstanceHistoryEventsRequest
	GetTag() []*DescribeRCInstanceHistoryEventsRequestTag
}

type DescribeRCInstanceHistoryEventsRequest struct {
	EventPublishTime         *DescribeRCInstanceHistoryEventsRequestEventPublishTime `json:"EventPublishTime,omitempty" xml:"EventPublishTime,omitempty" type:"Struct"`
	NotBefore                *DescribeRCInstanceHistoryEventsRequestNotBefore        `json:"NotBefore,omitempty" xml:"NotBefore,omitempty" type:"Struct"`
	EventCycleStatus         *string                                                 `json:"EventCycleStatus,omitempty" xml:"EventCycleStatus,omitempty"`
	EventId                  []*string                                               `json:"EventId,omitempty" xml:"EventId,omitempty" type:"Repeated"`
	EventType                *string                                                 `json:"EventType,omitempty" xml:"EventType,omitempty"`
	ImpactLevel              *string                                                 `json:"ImpactLevel,omitempty" xml:"ImpactLevel,omitempty"`
	InstanceEventCycleStatus []*string                                               `json:"InstanceEventCycleStatus,omitempty" xml:"InstanceEventCycleStatus,omitempty" type:"Repeated"`
	InstanceEventType        []*string                                               `json:"InstanceEventType,omitempty" xml:"InstanceEventType,omitempty" type:"Repeated"`
	InstanceId               *string                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults               *string                                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	PageNumber               *string                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                 *string                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId        *string                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                                      `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      []*string                                    `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag             []*DescribeRCInstanceHistoryEventsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeRCInstanceHistoryEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceHistoryEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetEventPublishTime() *DescribeRCInstanceHistoryEventsRequestEventPublishTime {
	return s.EventPublishTime
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetNotBefore() *DescribeRCInstanceHistoryEventsRequestNotBefore {
	return s.NotBefore
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetEventCycleStatus() *string {
	return s.EventCycleStatus
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetEventId() []*string {
	return s.EventId
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetImpactLevel() *string {
	return s.ImpactLevel
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetInstanceEventCycleStatus() []*string {
	return s.InstanceEventCycleStatus
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetInstanceEventType() []*string {
	return s.InstanceEventType
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetTag() []*DescribeRCInstanceHistoryEventsRequestTag {
	return s.Tag
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetEventPublishTime(v *DescribeRCInstanceHistoryEventsRequestEventPublishTime) *DescribeRCInstanceHistoryEventsRequest {
	s.EventPublishTime = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetNotBefore(v *DescribeRCInstanceHistoryEventsRequestNotBefore) *DescribeRCInstanceHistoryEventsRequest {
	s.NotBefore = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetEventCycleStatus(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.EventCycleStatus = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetEventId(v []*string) *DescribeRCInstanceHistoryEventsRequest {
	s.EventId = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetEventType(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.EventType = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetImpactLevel(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.ImpactLevel = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetInstanceEventCycleStatus(v []*string) *DescribeRCInstanceHistoryEventsRequest {
	s.InstanceEventCycleStatus = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetInstanceEventType(v []*string) *DescribeRCInstanceHistoryEventsRequest {
	s.InstanceEventType = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetInstanceId(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetMaxResults(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetPageNumber(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetPageSize(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetRegionId(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetResourceGroupId(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetResourceId(v []*string) *DescribeRCInstanceHistoryEventsRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetTag(v []*DescribeRCInstanceHistoryEventsRequestTag) *DescribeRCInstanceHistoryEventsRequest {
	s.Tag = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) Validate() error {
	if s.EventPublishTime != nil {
		if err := s.EventPublishTime.Validate(); err != nil {
			return err
		}
	}
	if s.NotBefore != nil {
		if err := s.NotBefore.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCInstanceHistoryEventsRequestEventPublishTime struct {
	End   *string `json:"End,omitempty" xml:"End,omitempty"`
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeRCInstanceHistoryEventsRequestEventPublishTime) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceHistoryEventsRequestEventPublishTime) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceHistoryEventsRequestEventPublishTime) GetEnd() *string {
	return s.End
}

func (s *DescribeRCInstanceHistoryEventsRequestEventPublishTime) GetStart() *string {
	return s.Start
}

func (s *DescribeRCInstanceHistoryEventsRequestEventPublishTime) SetEnd(v string) *DescribeRCInstanceHistoryEventsRequestEventPublishTime {
	s.End = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequestEventPublishTime) SetStart(v string) *DescribeRCInstanceHistoryEventsRequestEventPublishTime {
	s.Start = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequestEventPublishTime) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstanceHistoryEventsRequestNotBefore struct {
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// 2017-11-30T06:32:31Z
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeRCInstanceHistoryEventsRequestNotBefore) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceHistoryEventsRequestNotBefore) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceHistoryEventsRequestNotBefore) GetEnd() *string {
	return s.End
}

func (s *DescribeRCInstanceHistoryEventsRequestNotBefore) GetStart() *string {
	return s.Start
}

func (s *DescribeRCInstanceHistoryEventsRequestNotBefore) SetEnd(v string) *DescribeRCInstanceHistoryEventsRequestNotBefore {
	s.End = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequestNotBefore) SetStart(v string) *DescribeRCInstanceHistoryEventsRequestNotBefore {
	s.Start = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequestNotBefore) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstanceHistoryEventsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRCInstanceHistoryEventsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceHistoryEventsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceHistoryEventsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeRCInstanceHistoryEventsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeRCInstanceHistoryEventsRequestTag) SetKey(v string) *DescribeRCInstanceHistoryEventsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequestTag) SetValue(v string) *DescribeRCInstanceHistoryEventsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequestTag) Validate() error {
	return dara.Validate(s)
}
