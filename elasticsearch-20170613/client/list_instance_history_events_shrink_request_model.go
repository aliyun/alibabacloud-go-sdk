// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceHistoryEventsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*ListInstanceHistoryEventsShrinkRequestBody) *ListInstanceHistoryEventsShrinkRequest
	GetBody() []*ListInstanceHistoryEventsShrinkRequestBody
	SetEventCreateEndTime(v string) *ListInstanceHistoryEventsShrinkRequest
	GetEventCreateEndTime() *string
	SetEventCreateStartTime(v string) *ListInstanceHistoryEventsShrinkRequest
	GetEventCreateStartTime() *string
	SetEventCycleStatusShrink(v string) *ListInstanceHistoryEventsShrinkRequest
	GetEventCycleStatusShrink() *string
	SetEventExecuteEndTime(v string) *ListInstanceHistoryEventsShrinkRequest
	GetEventExecuteEndTime() *string
	SetEventExecuteStartTime(v string) *ListInstanceHistoryEventsShrinkRequest
	GetEventExecuteStartTime() *string
	SetEventFinashEndTime(v string) *ListInstanceHistoryEventsShrinkRequest
	GetEventFinashEndTime() *string
	SetEventFinashStartTime(v string) *ListInstanceHistoryEventsShrinkRequest
	GetEventFinashStartTime() *string
	SetEventLevelShrink(v string) *ListInstanceHistoryEventsShrinkRequest
	GetEventLevelShrink() *string
	SetEventTypeShrink(v string) *ListInstanceHistoryEventsShrinkRequest
	GetEventTypeShrink() *string
	SetInstanceId(v string) *ListInstanceHistoryEventsShrinkRequest
	GetInstanceId() *string
	SetNodeIP(v string) *ListInstanceHistoryEventsShrinkRequest
	GetNodeIP() *string
	SetPage(v int32) *ListInstanceHistoryEventsShrinkRequest
	GetPage() *int32
	SetSize(v int32) *ListInstanceHistoryEventsShrinkRequest
	GetSize() *int32
}

type ListInstanceHistoryEventsShrinkRequest struct {
	Body []*ListInstanceHistoryEventsShrinkRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// example:
	//
	// 1645596516000
	EventCreateEndTime *string `json:"eventCreateEndTime,omitempty" xml:"eventCreateEndTime,omitempty"`
	// example:
	//
	// 1645596516000
	EventCreateStartTime   *string `json:"eventCreateStartTime,omitempty" xml:"eventCreateStartTime,omitempty"`
	EventCycleStatusShrink *string `json:"eventCycleStatus,omitempty" xml:"eventCycleStatus,omitempty"`
	// example:
	//
	// 1645596516000
	EventExecuteEndTime *string `json:"eventExecuteEndTime,omitempty" xml:"eventExecuteEndTime,omitempty"`
	// example:
	//
	// 1645596516000
	EventExecuteStartTime *string `json:"eventExecuteStartTime,omitempty" xml:"eventExecuteStartTime,omitempty"`
	// example:
	//
	// 1645596516000
	EventFinashEndTime *string `json:"eventFinashEndTime,omitempty" xml:"eventFinashEndTime,omitempty"`
	// example:
	//
	// 1645596516000
	EventFinashStartTime *string `json:"eventFinashStartTime,omitempty" xml:"eventFinashStartTime,omitempty"`
	EventLevelShrink     *string `json:"eventLevel,omitempty" xml:"eventLevel,omitempty"`
	EventTypeShrink      *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// example:
	//
	// es-cn-2r42l7a740005****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 10.1.xx.xx
	NodeIP *string `json:"nodeIP,omitempty" xml:"nodeIP,omitempty"`
	// example:
	//
	// 0
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListInstanceHistoryEventsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHistoryEventsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceHistoryEventsShrinkRequest) GetBody() []*ListInstanceHistoryEventsShrinkRequestBody {
	return s.Body
}

func (s *ListInstanceHistoryEventsShrinkRequest) GetEventCreateEndTime() *string {
	return s.EventCreateEndTime
}

func (s *ListInstanceHistoryEventsShrinkRequest) GetEventCreateStartTime() *string {
	return s.EventCreateStartTime
}

func (s *ListInstanceHistoryEventsShrinkRequest) GetEventCycleStatusShrink() *string {
	return s.EventCycleStatusShrink
}

func (s *ListInstanceHistoryEventsShrinkRequest) GetEventExecuteEndTime() *string {
	return s.EventExecuteEndTime
}

func (s *ListInstanceHistoryEventsShrinkRequest) GetEventExecuteStartTime() *string {
	return s.EventExecuteStartTime
}

func (s *ListInstanceHistoryEventsShrinkRequest) GetEventFinashEndTime() *string {
	return s.EventFinashEndTime
}

func (s *ListInstanceHistoryEventsShrinkRequest) GetEventFinashStartTime() *string {
	return s.EventFinashStartTime
}

func (s *ListInstanceHistoryEventsShrinkRequest) GetEventLevelShrink() *string {
	return s.EventLevelShrink
}

func (s *ListInstanceHistoryEventsShrinkRequest) GetEventTypeShrink() *string {
	return s.EventTypeShrink
}

func (s *ListInstanceHistoryEventsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceHistoryEventsShrinkRequest) GetNodeIP() *string {
	return s.NodeIP
}

func (s *ListInstanceHistoryEventsShrinkRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListInstanceHistoryEventsShrinkRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListInstanceHistoryEventsShrinkRequest) SetBody(v []*ListInstanceHistoryEventsShrinkRequestBody) *ListInstanceHistoryEventsShrinkRequest {
	s.Body = v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequest) SetEventCreateEndTime(v string) *ListInstanceHistoryEventsShrinkRequest {
	s.EventCreateEndTime = &v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequest) SetEventCreateStartTime(v string) *ListInstanceHistoryEventsShrinkRequest {
	s.EventCreateStartTime = &v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequest) SetEventCycleStatusShrink(v string) *ListInstanceHistoryEventsShrinkRequest {
	s.EventCycleStatusShrink = &v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequest) SetEventExecuteEndTime(v string) *ListInstanceHistoryEventsShrinkRequest {
	s.EventExecuteEndTime = &v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequest) SetEventExecuteStartTime(v string) *ListInstanceHistoryEventsShrinkRequest {
	s.EventExecuteStartTime = &v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequest) SetEventFinashEndTime(v string) *ListInstanceHistoryEventsShrinkRequest {
	s.EventFinashEndTime = &v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequest) SetEventFinashStartTime(v string) *ListInstanceHistoryEventsShrinkRequest {
	s.EventFinashStartTime = &v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequest) SetEventLevelShrink(v string) *ListInstanceHistoryEventsShrinkRequest {
	s.EventLevelShrink = &v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequest) SetEventTypeShrink(v string) *ListInstanceHistoryEventsShrinkRequest {
	s.EventTypeShrink = &v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequest) SetInstanceId(v string) *ListInstanceHistoryEventsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequest) SetNodeIP(v string) *ListInstanceHistoryEventsShrinkRequest {
	s.NodeIP = &v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequest) SetPage(v int32) *ListInstanceHistoryEventsShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequest) SetSize(v int32) *ListInstanceHistoryEventsShrinkRequest {
	s.Size = &v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type ListInstanceHistoryEventsShrinkRequestBody struct {
	// example:
	//
	// true
	Desc *bool `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// event_time
	SortField *string `json:"sortField,omitempty" xml:"sortField,omitempty"`
}

func (s ListInstanceHistoryEventsShrinkRequestBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHistoryEventsShrinkRequestBody) GoString() string {
	return s.String()
}

func (s *ListInstanceHistoryEventsShrinkRequestBody) GetDesc() *bool {
	return s.Desc
}

func (s *ListInstanceHistoryEventsShrinkRequestBody) GetSortField() *string {
	return s.SortField
}

func (s *ListInstanceHistoryEventsShrinkRequestBody) SetDesc(v bool) *ListInstanceHistoryEventsShrinkRequestBody {
	s.Desc = &v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequestBody) SetSortField(v string) *ListInstanceHistoryEventsShrinkRequestBody {
	s.SortField = &v
	return s
}

func (s *ListInstanceHistoryEventsShrinkRequestBody) Validate() error {
	return dara.Validate(s)
}
