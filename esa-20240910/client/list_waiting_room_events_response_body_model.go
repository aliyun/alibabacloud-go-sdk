// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaitingRoomEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWaitingRoomEventsResponseBody
	GetRequestId() *string
	SetWaitingRoomEvents(v []*ListWaitingRoomEventsResponseBodyWaitingRoomEvents) *ListWaitingRoomEventsResponseBody
	GetWaitingRoomEvents() []*ListWaitingRoomEventsResponseBodyWaitingRoomEvents
}

type ListWaitingRoomEventsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// f3c3700a-4c0f-4a24-b576-fd7dbf9e7c55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of waiting room events, which contains the specific information of each event.
	WaitingRoomEvents []*ListWaitingRoomEventsResponseBodyWaitingRoomEvents `json:"WaitingRoomEvents,omitempty" xml:"WaitingRoomEvents,omitempty" type:"Repeated"`
}

func (s ListWaitingRoomEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingRoomEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWaitingRoomEventsResponseBody) GetWaitingRoomEvents() []*ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	return s.WaitingRoomEvents
}

func (s *ListWaitingRoomEventsResponseBody) SetRequestId(v string) *ListWaitingRoomEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBody) SetWaitingRoomEvents(v []*ListWaitingRoomEventsResponseBodyWaitingRoomEvents) *ListWaitingRoomEventsResponseBody {
	s.WaitingRoomEvents = v
	return s
}

func (s *ListWaitingRoomEventsResponseBody) Validate() error {
	if s.WaitingRoomEvents != nil {
		for _, item := range s.WaitingRoomEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWaitingRoomEventsResponseBodyWaitingRoomEvents struct {
	// The custom content of the waiting room page. This parameter is required when the waiting room type is custom. The content must be URL-encoded.
	//
	// example:
	//
	// html-yets-maqi1111
	CustomPageHtml *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	// The detailed description of the event.
	//
	// example:
	//
	// terraform-example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether session renewal is disabled. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// off
	DisableSessionRenewalEnable *string `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	// The enabled status of the event. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The end time of the event. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1719814497
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to enable JSON responses. If this feature is enabled and the Accept request header contains "application/json", JSON data is returned. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// off
	JsonResponseEnable *string `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	// The language of the waiting room page. This parameter is required when the waiting room type is default. Valid values:
	//
	// - **enus**: English.
	//
	// - **zhcn**: Simplified Chinese.
	//
	// - **zhhk**: Traditional Chinese.
	//
	// example:
	//
	// zhcn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The event name, which is a custom event description.
	//
	// example:
	//
	// test event
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of new users per minute.
	//
	// example:
	//
	// 11
	NewUsersPerMinute *string `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	// Specifies whether to enable pre-queuing. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// on
	PreQueueEnable *string `json:"PreQueueEnable,omitempty" xml:"PreQueueEnable,omitempty"`
	// The start time of pre-queuing. The value is a UNIX timestamp. This parameter is valid only when pre-queuing is enabled.
	//
	// example:
	//
	// 1719814097
	PreQueueStartTime *string `json:"PreQueueStartTime,omitempty" xml:"PreQueueStartTime,omitempty"`
	// The queuing method. Valid values:
	//
	// - **random**: Random.
	//
	// - **fifo**: First-in, first-out (FIFO).
	//
	// - **passthrough**: Passthrough.
	//
	// - **reject-all**: Reject all.
	//
	// example:
	//
	// fifo
	QueuingMethod *string `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	// The waiting room status code. Valid values:
	//
	// - **200**
	//
	// - **202**
	//
	// - **429**
	//
	// example:
	//
	// 200
	QueuingStatusCode *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	// Specifies whether to enable random pre-queuing. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// on
	RandomPreQueueEnable *string `json:"RandomPreQueueEnable,omitempty" xml:"RandomPreQueueEnable,omitempty"`
	// The session duration, in minutes.
	//
	// example:
	//
	// 3
	SessionDuration *string `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// The start time of the event. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1719814398
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of active users.
	//
	// example:
	//
	// 22
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// The event ID, which uniquely identifies a waiting room event.
	//
	// example:
	//
	// 89677721098****
	WaitingRoomEventId *int64 `json:"WaitingRoomEventId,omitempty" xml:"WaitingRoomEventId,omitempty"`
	// The ID of the associated waiting room.
	//
	// example:
	//
	// 5c938a045c9ca46607163d34966****
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
	// The waiting room type. Valid values:
	//
	// - **default**: Default type.
	//
	// - **custom**: Custom type.
	//
	// example:
	//
	// custom
	WaitingRoomType *string `json:"WaitingRoomType,omitempty" xml:"WaitingRoomType,omitempty"`
}

func (s ListWaitingRoomEventsResponseBodyWaitingRoomEvents) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetCustomPageHtml() *string {
	return s.CustomPageHtml
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetDescription() *string {
	return s.Description
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetDisableSessionRenewalEnable() *string {
	return s.DisableSessionRenewalEnable
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetEnable() *string {
	return s.Enable
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetEndTime() *string {
	return s.EndTime
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetJsonResponseEnable() *string {
	return s.JsonResponseEnable
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetLanguage() *string {
	return s.Language
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetName() *string {
	return s.Name
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetNewUsersPerMinute() *string {
	return s.NewUsersPerMinute
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetPreQueueEnable() *string {
	return s.PreQueueEnable
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetPreQueueStartTime() *string {
	return s.PreQueueStartTime
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetQueuingMethod() *string {
	return s.QueuingMethod
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetQueuingStatusCode() *string {
	return s.QueuingStatusCode
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetRandomPreQueueEnable() *string {
	return s.RandomPreQueueEnable
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetSessionDuration() *string {
	return s.SessionDuration
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetStartTime() *string {
	return s.StartTime
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetTotalActiveUsers() *string {
	return s.TotalActiveUsers
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetWaitingRoomEventId() *int64 {
	return s.WaitingRoomEventId
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetWaitingRoomId() *string {
	return s.WaitingRoomId
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GetWaitingRoomType() *string {
	return s.WaitingRoomType
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetCustomPageHtml(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.CustomPageHtml = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetDescription(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.Description = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetDisableSessionRenewalEnable(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetEnable(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.Enable = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetEndTime(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.EndTime = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetJsonResponseEnable(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.JsonResponseEnable = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetLanguage(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.Language = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetName(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.Name = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetNewUsersPerMinute(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.NewUsersPerMinute = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetPreQueueEnable(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.PreQueueEnable = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetPreQueueStartTime(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.PreQueueStartTime = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetQueuingMethod(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.QueuingMethod = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetQueuingStatusCode(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.QueuingStatusCode = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetRandomPreQueueEnable(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.RandomPreQueueEnable = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetSessionDuration(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.SessionDuration = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetStartTime(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.StartTime = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetTotalActiveUsers(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.TotalActiveUsers = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetWaitingRoomEventId(v int64) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.WaitingRoomEventId = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetWaitingRoomId(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.WaitingRoomId = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetWaitingRoomType(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.WaitingRoomType = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) Validate() error {
	return dara.Validate(s)
}
