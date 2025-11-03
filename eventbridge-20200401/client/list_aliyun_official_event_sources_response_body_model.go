// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliyunOfficialEventSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAliyunOfficialEventSourcesResponseBody
	GetCode() *string
	SetData(v *ListAliyunOfficialEventSourcesResponseBodyData) *ListAliyunOfficialEventSourcesResponseBody
	GetData() *ListAliyunOfficialEventSourcesResponseBodyData
	SetMessage(v string) *ListAliyunOfficialEventSourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAliyunOfficialEventSourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAliyunOfficialEventSourcesResponseBody
	GetSuccess() *bool
}

type ListAliyunOfficialEventSourcesResponseBody struct {
	// The response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For a list of error codes, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ListAliyunOfficialEventSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// InvalidArgument
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5f80e9b3-98d5-4f51-8412-c758818a03e4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. If the operation is successful, the value true is returned.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAliyunOfficialEventSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAliyunOfficialEventSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAliyunOfficialEventSourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAliyunOfficialEventSourcesResponseBody) GetData() *ListAliyunOfficialEventSourcesResponseBodyData {
	return s.Data
}

func (s *ListAliyunOfficialEventSourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAliyunOfficialEventSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAliyunOfficialEventSourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAliyunOfficialEventSourcesResponseBody) SetCode(v string) *ListAliyunOfficialEventSourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBody) SetData(v *ListAliyunOfficialEventSourcesResponseBodyData) *ListAliyunOfficialEventSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBody) SetMessage(v string) *ListAliyunOfficialEventSourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBody) SetRequestId(v string) *ListAliyunOfficialEventSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBody) SetSuccess(v bool) *ListAliyunOfficialEventSourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAliyunOfficialEventSourcesResponseBodyData struct {
	// The event sources.
	EventSourceList []*ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList `json:"EventSourceList,omitempty" xml:"EventSourceList,omitempty" type:"Repeated"`
}

func (s ListAliyunOfficialEventSourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAliyunOfficialEventSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAliyunOfficialEventSourcesResponseBodyData) GetEventSourceList() []*ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	return s.EventSourceList
}

func (s *ListAliyunOfficialEventSourcesResponseBodyData) SetEventSourceList(v []*ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) *ListAliyunOfficialEventSourcesResponseBodyData {
	s.EventSourceList = v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyData) Validate() error {
	if s.EventSourceList != nil {
		for _, item := range s.EventSourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList struct {
	// The Alibaba Cloud Resource Name (ARN) of the event bus.
	//
	// example:
	//
	// acs:eventbridge:cn-hangzhou:SYSTEM:eventsource/acs.aliyuncvc
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The time when the event source was created. Unit: milliseconds.
	//
	// example:
	//
	// 1607071602000
	Ctime *float32 `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	// The description of the event source.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event source to which the event type belongs.
	//
	// example:
	//
	// acs.aliyuncvc
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event types.
	EventTypes []*ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	// The full name of the event source.
	//
	// example:
	//
	// E-MapReduce
	FullName *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	// The name of the event source.
	//
	// example:
	//
	// acs.aliyuncvc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the event source. Valid value: Activated.
	//
	// example:
	//
	// Activated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the event source.
	//
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) String() string {
	return dara.Prettify(s)
}

func (s ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) GoString() string {
	return s.String()
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) GetArn() *string {
	return s.Arn
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) GetCtime() *float32 {
	return s.Ctime
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) GetDescription() *string {
	return s.Description
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) GetEventBusName() *string {
	return s.EventBusName
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) GetEventTypes() []*ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes {
	return s.EventTypes
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) GetFullName() *string {
	return s.FullName
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) GetName() *string {
	return s.Name
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) GetStatus() *string {
	return s.Status
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) GetType() *string {
	return s.Type
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetArn(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.Arn = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetCtime(v float32) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.Ctime = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetDescription(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.Description = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetEventBusName(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.EventBusName = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetEventTypes(v []*ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.EventTypes = v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetFullName(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.FullName = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetName(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.Name = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetStatus(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.Status = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetType(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.Type = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) Validate() error {
	if s.EventTypes != nil {
		for _, item := range s.EventTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes struct {
	// The name of the event source.
	//
	// example:
	//
	// name
	EventSourceName *string `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
	// The name of the group to which the event type belongs.
	//
	// example:
	//
	// aliyuncvc:MeetingEvent
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The full name of the event type.
	//
	// example:
	//
	// aliyuncvc:MeetingEvent:MeetingStateEvent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The short name of the event type.
	//
	// example:
	//
	// MeetingStateEvent
	ShortName *string `json:"ShortName,omitempty" xml:"ShortName,omitempty"`
}

func (s ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) String() string {
	return dara.Prettify(s)
}

func (s ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) GoString() string {
	return s.String()
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) GetEventSourceName() *string {
	return s.EventSourceName
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) GetGroupName() *string {
	return s.GroupName
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) GetName() *string {
	return s.Name
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) GetShortName() *string {
	return s.ShortName
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) SetEventSourceName(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes {
	s.EventSourceName = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) SetGroupName(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes {
	s.GroupName = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) SetName(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes {
	s.Name = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) SetShortName(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes {
	s.ShortName = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) Validate() error {
	return dara.Validate(s)
}
