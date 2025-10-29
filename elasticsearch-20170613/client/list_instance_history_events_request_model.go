// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceHistoryEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*ListInstanceHistoryEventsRequestBody) *ListInstanceHistoryEventsRequest
	GetBody() []*ListInstanceHistoryEventsRequestBody
	SetEventCreateEndTime(v string) *ListInstanceHistoryEventsRequest
	GetEventCreateEndTime() *string
	SetEventCreateStartTime(v string) *ListInstanceHistoryEventsRequest
	GetEventCreateStartTime() *string
	SetEventCycleStatus(v []*string) *ListInstanceHistoryEventsRequest
	GetEventCycleStatus() []*string
	SetEventExecuteEndTime(v string) *ListInstanceHistoryEventsRequest
	GetEventExecuteEndTime() *string
	SetEventExecuteStartTime(v string) *ListInstanceHistoryEventsRequest
	GetEventExecuteStartTime() *string
	SetEventFinashEndTime(v string) *ListInstanceHistoryEventsRequest
	GetEventFinashEndTime() *string
	SetEventFinashStartTime(v string) *ListInstanceHistoryEventsRequest
	GetEventFinashStartTime() *string
	SetEventLevel(v []*string) *ListInstanceHistoryEventsRequest
	GetEventLevel() []*string
	SetEventType(v []*string) *ListInstanceHistoryEventsRequest
	GetEventType() []*string
	SetInstanceId(v string) *ListInstanceHistoryEventsRequest
	GetInstanceId() *string
	SetNodeIP(v string) *ListInstanceHistoryEventsRequest
	GetNodeIP() *string
	SetPage(v int32) *ListInstanceHistoryEventsRequest
	GetPage() *int32
	SetSize(v int32) *ListInstanceHistoryEventsRequest
	GetSize() *int32
}

type ListInstanceHistoryEventsRequest struct {
	Body []*ListInstanceHistoryEventsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// example:
	//
	// 1645596516000
	EventCreateEndTime *string `json:"eventCreateEndTime,omitempty" xml:"eventCreateEndTime,omitempty"`
	// example:
	//
	// 1645596516000
	EventCreateStartTime *string   `json:"eventCreateStartTime,omitempty" xml:"eventCreateStartTime,omitempty"`
	EventCycleStatus     []*string `json:"eventCycleStatus,omitempty" xml:"eventCycleStatus,omitempty" type:"Repeated"`
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
	EventFinashStartTime *string   `json:"eventFinashStartTime,omitempty" xml:"eventFinashStartTime,omitempty"`
	EventLevel           []*string `json:"eventLevel,omitempty" xml:"eventLevel,omitempty" type:"Repeated"`
	EventType            []*string `json:"eventType,omitempty" xml:"eventType,omitempty" type:"Repeated"`
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

func (s ListInstanceHistoryEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHistoryEventsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceHistoryEventsRequest) GetBody() []*ListInstanceHistoryEventsRequestBody {
	return s.Body
}

func (s *ListInstanceHistoryEventsRequest) GetEventCreateEndTime() *string {
	return s.EventCreateEndTime
}

func (s *ListInstanceHistoryEventsRequest) GetEventCreateStartTime() *string {
	return s.EventCreateStartTime
}

func (s *ListInstanceHistoryEventsRequest) GetEventCycleStatus() []*string {
	return s.EventCycleStatus
}

func (s *ListInstanceHistoryEventsRequest) GetEventExecuteEndTime() *string {
	return s.EventExecuteEndTime
}

func (s *ListInstanceHistoryEventsRequest) GetEventExecuteStartTime() *string {
	return s.EventExecuteStartTime
}

func (s *ListInstanceHistoryEventsRequest) GetEventFinashEndTime() *string {
	return s.EventFinashEndTime
}

func (s *ListInstanceHistoryEventsRequest) GetEventFinashStartTime() *string {
	return s.EventFinashStartTime
}

func (s *ListInstanceHistoryEventsRequest) GetEventLevel() []*string {
	return s.EventLevel
}

func (s *ListInstanceHistoryEventsRequest) GetEventType() []*string {
	return s.EventType
}

func (s *ListInstanceHistoryEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceHistoryEventsRequest) GetNodeIP() *string {
	return s.NodeIP
}

func (s *ListInstanceHistoryEventsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListInstanceHistoryEventsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListInstanceHistoryEventsRequest) SetBody(v []*ListInstanceHistoryEventsRequestBody) *ListInstanceHistoryEventsRequest {
	s.Body = v
	return s
}

func (s *ListInstanceHistoryEventsRequest) SetEventCreateEndTime(v string) *ListInstanceHistoryEventsRequest {
	s.EventCreateEndTime = &v
	return s
}

func (s *ListInstanceHistoryEventsRequest) SetEventCreateStartTime(v string) *ListInstanceHistoryEventsRequest {
	s.EventCreateStartTime = &v
	return s
}

func (s *ListInstanceHistoryEventsRequest) SetEventCycleStatus(v []*string) *ListInstanceHistoryEventsRequest {
	s.EventCycleStatus = v
	return s
}

func (s *ListInstanceHistoryEventsRequest) SetEventExecuteEndTime(v string) *ListInstanceHistoryEventsRequest {
	s.EventExecuteEndTime = &v
	return s
}

func (s *ListInstanceHistoryEventsRequest) SetEventExecuteStartTime(v string) *ListInstanceHistoryEventsRequest {
	s.EventExecuteStartTime = &v
	return s
}

func (s *ListInstanceHistoryEventsRequest) SetEventFinashEndTime(v string) *ListInstanceHistoryEventsRequest {
	s.EventFinashEndTime = &v
	return s
}

func (s *ListInstanceHistoryEventsRequest) SetEventFinashStartTime(v string) *ListInstanceHistoryEventsRequest {
	s.EventFinashStartTime = &v
	return s
}

func (s *ListInstanceHistoryEventsRequest) SetEventLevel(v []*string) *ListInstanceHistoryEventsRequest {
	s.EventLevel = v
	return s
}

func (s *ListInstanceHistoryEventsRequest) SetEventType(v []*string) *ListInstanceHistoryEventsRequest {
	s.EventType = v
	return s
}

func (s *ListInstanceHistoryEventsRequest) SetInstanceId(v string) *ListInstanceHistoryEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceHistoryEventsRequest) SetNodeIP(v string) *ListInstanceHistoryEventsRequest {
	s.NodeIP = &v
	return s
}

func (s *ListInstanceHistoryEventsRequest) SetPage(v int32) *ListInstanceHistoryEventsRequest {
	s.Page = &v
	return s
}

func (s *ListInstanceHistoryEventsRequest) SetSize(v int32) *ListInstanceHistoryEventsRequest {
	s.Size = &v
	return s
}

func (s *ListInstanceHistoryEventsRequest) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceHistoryEventsRequestBody struct {
	// example:
	//
	// true
	Desc *bool `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// event_time
	SortField *string `json:"sortField,omitempty" xml:"sortField,omitempty"`
}

func (s ListInstanceHistoryEventsRequestBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHistoryEventsRequestBody) GoString() string {
	return s.String()
}

func (s *ListInstanceHistoryEventsRequestBody) GetDesc() *bool {
	return s.Desc
}

func (s *ListInstanceHistoryEventsRequestBody) GetSortField() *string {
	return s.SortField
}

func (s *ListInstanceHistoryEventsRequestBody) SetDesc(v bool) *ListInstanceHistoryEventsRequestBody {
	s.Desc = &v
	return s
}

func (s *ListInstanceHistoryEventsRequestBody) SetSortField(v string) *ListInstanceHistoryEventsRequestBody {
	s.SortField = &v
	return s
}

func (s *ListInstanceHistoryEventsRequestBody) Validate() error {
	return dara.Validate(s)
}
