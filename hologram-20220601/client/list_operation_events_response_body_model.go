// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListOperationEventsResponseBodyData) *ListOperationEventsResponseBody
	GetData() []*ListOperationEventsResponseBodyData
	SetPageNumber(v int64) *ListOperationEventsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListOperationEventsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListOperationEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListOperationEventsResponseBody
	GetTotalCount() *int64
}

type ListOperationEventsResponseBody struct {
	Data []*ListOperationEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D3AE84AB-0873-5FC7-A4C4-8CF869D2FA70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 120
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOperationEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperationEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationEventsResponseBody) GetData() []*ListOperationEventsResponseBodyData {
	return s.Data
}

func (s *ListOperationEventsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListOperationEventsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOperationEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOperationEventsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOperationEventsResponseBody) SetData(v []*ListOperationEventsResponseBodyData) *ListOperationEventsResponseBody {
	s.Data = v
	return s
}

func (s *ListOperationEventsResponseBody) SetPageNumber(v int64) *ListOperationEventsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListOperationEventsResponseBody) SetPageSize(v int64) *ListOperationEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOperationEventsResponseBody) SetRequestId(v string) *ListOperationEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationEventsResponseBody) SetTotalCount(v int64) *ListOperationEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOperationEventsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOperationEventsResponseBodyData struct {
	// example:
	//
	// true
	Cancelable *bool `json:"Cancelable,omitempty" xml:"Cancelable,omitempty"`
	// example:
	//
	// true
	ChangeScheduleTime *bool `json:"ChangeScheduleTime,omitempty" xml:"ChangeScheduleTime,omitempty"`
	// example:
	//
	// {}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// HOT_UPGRADE
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// UPGRADE_ON_CUSTOMERS_BEHALF
	EventType         *string                                                 `json:"EventType,omitempty" xml:"EventType,omitempty"`
	FollowerInstances []*ListOperationEventsResponseBodyDataFollowerInstances `json:"FollowerInstances,omitempty" xml:"FollowerInstances,omitempty" type:"Repeated"`
	// Id
	//
	// example:
	//
	// 1826503661244379138
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// haha
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 21:00-22:00
	MaintainWindow *string `json:"MaintainWindow,omitempty" xml:"MaintainWindow,omitempty"`
	// example:
	//
	// hgpost-cn-sxfsdfsd
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// ScheduleTime
	//
	// example:
	//
	// 2024-08-22 15:49:28
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// State
	//
	// example:
	//
	// queued
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListOperationEventsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListOperationEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOperationEventsResponseBodyData) GetCancelable() *bool {
	return s.Cancelable
}

func (s *ListOperationEventsResponseBodyData) GetChangeScheduleTime() *bool {
	return s.ChangeScheduleTime
}

func (s *ListOperationEventsResponseBodyData) GetDetail() *string {
	return s.Detail
}

func (s *ListOperationEventsResponseBodyData) GetEventName() *string {
	return s.EventName
}

func (s *ListOperationEventsResponseBodyData) GetEventType() *string {
	return s.EventType
}

func (s *ListOperationEventsResponseBodyData) GetFollowerInstances() []*ListOperationEventsResponseBodyDataFollowerInstances {
	return s.FollowerInstances
}

func (s *ListOperationEventsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListOperationEventsResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListOperationEventsResponseBodyData) GetMaintainWindow() *string {
	return s.MaintainWindow
}

func (s *ListOperationEventsResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListOperationEventsResponseBodyData) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *ListOperationEventsResponseBodyData) GetState() *string {
	return s.State
}

func (s *ListOperationEventsResponseBodyData) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListOperationEventsResponseBodyData) SetCancelable(v bool) *ListOperationEventsResponseBodyData {
	s.Cancelable = &v
	return s
}

func (s *ListOperationEventsResponseBodyData) SetChangeScheduleTime(v bool) *ListOperationEventsResponseBodyData {
	s.ChangeScheduleTime = &v
	return s
}

func (s *ListOperationEventsResponseBodyData) SetDetail(v string) *ListOperationEventsResponseBodyData {
	s.Detail = &v
	return s
}

func (s *ListOperationEventsResponseBodyData) SetEventName(v string) *ListOperationEventsResponseBodyData {
	s.EventName = &v
	return s
}

func (s *ListOperationEventsResponseBodyData) SetEventType(v string) *ListOperationEventsResponseBodyData {
	s.EventType = &v
	return s
}

func (s *ListOperationEventsResponseBodyData) SetFollowerInstances(v []*ListOperationEventsResponseBodyDataFollowerInstances) *ListOperationEventsResponseBodyData {
	s.FollowerInstances = v
	return s
}

func (s *ListOperationEventsResponseBodyData) SetId(v string) *ListOperationEventsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListOperationEventsResponseBodyData) SetInstanceName(v string) *ListOperationEventsResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListOperationEventsResponseBodyData) SetMaintainWindow(v string) *ListOperationEventsResponseBodyData {
	s.MaintainWindow = &v
	return s
}

func (s *ListOperationEventsResponseBodyData) SetResourceId(v string) *ListOperationEventsResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *ListOperationEventsResponseBodyData) SetScheduleTime(v string) *ListOperationEventsResponseBodyData {
	s.ScheduleTime = &v
	return s
}

func (s *ListOperationEventsResponseBodyData) SetState(v string) *ListOperationEventsResponseBodyData {
	s.State = &v
	return s
}

func (s *ListOperationEventsResponseBodyData) SetZoneId(v string) *ListOperationEventsResponseBodyData {
	s.ZoneId = &v
	return s
}

func (s *ListOperationEventsResponseBodyData) Validate() error {
	if s.FollowerInstances != nil {
		for _, item := range s.FollowerInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOperationEventsResponseBodyDataFollowerInstances struct {
	// example:
	//
	// hgxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// hi
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s ListOperationEventsResponseBodyDataFollowerInstances) String() string {
	return dara.Prettify(s)
}

func (s ListOperationEventsResponseBodyDataFollowerInstances) GoString() string {
	return s.String()
}

func (s *ListOperationEventsResponseBodyDataFollowerInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOperationEventsResponseBodyDataFollowerInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListOperationEventsResponseBodyDataFollowerInstances) SetInstanceId(v string) *ListOperationEventsResponseBodyDataFollowerInstances {
	s.InstanceId = &v
	return s
}

func (s *ListOperationEventsResponseBodyDataFollowerInstances) SetInstanceName(v string) *ListOperationEventsResponseBodyDataFollowerInstances {
	s.InstanceName = &v
	return s
}

func (s *ListOperationEventsResponseBodyDataFollowerInstances) Validate() error {
	return dara.Validate(s)
}
