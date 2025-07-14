// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAppEventsResponseBody
	GetCode() *string
	SetData(v *ListAppEventsResponseBodyData) *ListAppEventsResponseBody
	GetData() *ListAppEventsResponseBodyData
	SetErrorCode(v string) *ListAppEventsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListAppEventsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAppEventsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAppEventsResponseBody
	GetSuccess() *bool
}

type ListAppEventsResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The events.
	Data *ListAppEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the call failed. Take note of the following rules:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4D805CA-926D-41B1-8E63-7AD0C1ED****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the events that occurred in the application were queried. Valid values:
	//
	// 	- **true**: The events were queried.
	//
	// 	- **false**: The events failed to be queried.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAppEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppEventsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAppEventsResponseBody) GetData() *ListAppEventsResponseBodyData {
	return s.Data
}

func (s *ListAppEventsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAppEventsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAppEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppEventsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAppEventsResponseBody) SetCode(v string) *ListAppEventsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppEventsResponseBody) SetData(v *ListAppEventsResponseBodyData) *ListAppEventsResponseBody {
	s.Data = v
	return s
}

func (s *ListAppEventsResponseBody) SetErrorCode(v string) *ListAppEventsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAppEventsResponseBody) SetMessage(v string) *ListAppEventsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppEventsResponseBody) SetRequestId(v string) *ListAppEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppEventsResponseBody) SetSuccess(v bool) *ListAppEventsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAppEventsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAppEventsResponseBodyData struct {
	// The events.
	AppEventEntity []*ListAppEventsResponseBodyDataAppEventEntity `json:"AppEventEntity,omitempty" xml:"AppEventEntity,omitempty" type:"Repeated"`
	// The number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of events that occurred in an application.
	//
	// example:
	//
	// 20
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListAppEventsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAppEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppEventsResponseBodyData) GetAppEventEntity() []*ListAppEventsResponseBodyDataAppEventEntity {
	return s.AppEventEntity
}

func (s *ListAppEventsResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAppEventsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppEventsResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListAppEventsResponseBodyData) SetAppEventEntity(v []*ListAppEventsResponseBodyDataAppEventEntity) *ListAppEventsResponseBodyData {
	s.AppEventEntity = v
	return s
}

func (s *ListAppEventsResponseBodyData) SetCurrentPage(v int32) *ListAppEventsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListAppEventsResponseBodyData) SetPageSize(v int32) *ListAppEventsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAppEventsResponseBodyData) SetTotalSize(v int32) *ListAppEventsResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListAppEventsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListAppEventsResponseBodyDataAppEventEntity struct {
	CauseAnalysis *string `json:"CauseAnalysis,omitempty" xml:"CauseAnalysis,omitempty"`
	// The type of the event. Valid values:
	//
	// example:
	//
	// Normal
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The timestamp of the first occurrence of the event.
	//
	// example:
	//
	// 2020-02-19T05:01:28Z
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	// The timestamp of the last occurrence of the event.
	//
	// example:
	//
	// 2020-02-19T05:01:28Z
	LastTimestamp *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The information about the event.
	//
	// example:
	//
	// Created container
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The type of the object.
	//
	// example:
	//
	// Pod
	ObjectKind *string `json:"ObjectKind,omitempty" xml:"ObjectKind,omitempty"`
	// The name of the object.
	//
	// example:
	//
	// errew-b86bf540-b4dc-47d8-a42f-b4997c14bd8f-5595cbddd6-2****
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// The cause of the event.
	//
	// example:
	//
	// Created
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ListAppEventsResponseBodyDataAppEventEntity) String() string {
	return dara.Prettify(s)
}

func (s ListAppEventsResponseBodyDataAppEventEntity) GoString() string {
	return s.String()
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) GetCauseAnalysis() *string {
	return s.CauseAnalysis
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) GetEventType() *string {
	return s.EventType
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) GetFirstTimestamp() *string {
	return s.FirstTimestamp
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) GetLastTimestamp() *string {
	return s.LastTimestamp
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) GetMessage() *string {
	return s.Message
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) GetObjectKind() *string {
	return s.ObjectKind
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) GetObjectName() *string {
	return s.ObjectName
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) GetReason() *string {
	return s.Reason
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) SetCauseAnalysis(v string) *ListAppEventsResponseBodyDataAppEventEntity {
	s.CauseAnalysis = &v
	return s
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) SetEventType(v string) *ListAppEventsResponseBodyDataAppEventEntity {
	s.EventType = &v
	return s
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) SetFirstTimestamp(v string) *ListAppEventsResponseBodyDataAppEventEntity {
	s.FirstTimestamp = &v
	return s
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) SetLastTimestamp(v string) *ListAppEventsResponseBodyDataAppEventEntity {
	s.LastTimestamp = &v
	return s
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) SetMessage(v string) *ListAppEventsResponseBodyDataAppEventEntity {
	s.Message = &v
	return s
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) SetObjectKind(v string) *ListAppEventsResponseBodyDataAppEventEntity {
	s.ObjectKind = &v
	return s
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) SetObjectName(v string) *ListAppEventsResponseBodyDataAppEventEntity {
	s.ObjectName = &v
	return s
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) SetReason(v string) *ListAppEventsResponseBodyDataAppEventEntity {
	s.Reason = &v
	return s
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) Validate() error {
	return dara.Validate(s)
}
