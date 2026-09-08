// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutEventsResponseBody
	GetCode() *string
	SetData(v *PutEventsResponseBodyData) *PutEventsResponseBody
	GetData() *PutEventsResponseBodyData
	SetMessage(v string) *PutEventsResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutEventsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutEventsResponseBody
	GetSuccess() *bool
}

type PutEventsResponseBody struct {
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *PutEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// EventBusNotExist
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier that Alibaba Cloud generated for the request.
	//
	// example:
	//
	// 2BC1857D-E633-5E79-B2C2-43EF5F7730D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values: true: The operation was successful. false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutEventsResponseBody) GoString() string {
	return s.String()
}

func (s *PutEventsResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutEventsResponseBody) GetData() *PutEventsResponseBodyData {
	return s.Data
}

func (s *PutEventsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutEventsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutEventsResponseBody) SetCode(v string) *PutEventsResponseBody {
	s.Code = &v
	return s
}

func (s *PutEventsResponseBody) SetData(v *PutEventsResponseBodyData) *PutEventsResponseBody {
	s.Data = v
	return s
}

func (s *PutEventsResponseBody) SetMessage(v string) *PutEventsResponseBody {
	s.Message = &v
	return s
}

func (s *PutEventsResponseBody) SetRequestId(v string) *PutEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutEventsResponseBody) SetSuccess(v bool) *PutEventsResponseBody {
	s.Success = &v
	return s
}

func (s *PutEventsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutEventsResponseBodyData struct {
	// The collection of event sending results.
	EntryList []*PutEventsResponseBodyDataEntryList `json:"EntryList,omitempty" xml:"EntryList,omitempty" type:"Repeated"`
	// The number of events that failed to be sent.
	//
	// example:
	//
	// 2
	FailedEntryCount *int32 `json:"FailedEntryCount,omitempty" xml:"FailedEntryCount,omitempty"`
}

func (s PutEventsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PutEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *PutEventsResponseBodyData) GetEntryList() []*PutEventsResponseBodyDataEntryList {
	return s.EntryList
}

func (s *PutEventsResponseBodyData) GetFailedEntryCount() *int32 {
	return s.FailedEntryCount
}

func (s *PutEventsResponseBodyData) SetEntryList(v []*PutEventsResponseBodyDataEntryList) *PutEventsResponseBodyData {
	s.EntryList = v
	return s
}

func (s *PutEventsResponseBodyData) SetFailedEntryCount(v int32) *PutEventsResponseBodyData {
	s.FailedEntryCount = &v
	return s
}

func (s *PutEventsResponseBodyData) Validate() error {
	if s.EntryList != nil {
		for _, item := range s.EntryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutEventsResponseBodyDataEntryList struct {
	// The error code.
	//
	// example:
	//
	// Success indicates success. Other values indicate exceptions
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The detailed error description.
	//
	// example:
	//
	// triggerPicture failed
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 4c8b7500-2aea-4f5a-b7dd-9d9dd986c07d
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The trace ID, which is used to query the exact call information.
	//
	// example:
	//
	// 4E17C677F5357FB23D1A7FF964CD1999
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s PutEventsResponseBodyDataEntryList) String() string {
	return dara.Prettify(s)
}

func (s PutEventsResponseBodyDataEntryList) GoString() string {
	return s.String()
}

func (s *PutEventsResponseBodyDataEntryList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PutEventsResponseBodyDataEntryList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *PutEventsResponseBodyDataEntryList) GetEventId() *string {
	return s.EventId
}

func (s *PutEventsResponseBodyDataEntryList) GetTraceId() *string {
	return s.TraceId
}

func (s *PutEventsResponseBodyDataEntryList) SetErrorCode(v string) *PutEventsResponseBodyDataEntryList {
	s.ErrorCode = &v
	return s
}

func (s *PutEventsResponseBodyDataEntryList) SetErrorMessage(v string) *PutEventsResponseBodyDataEntryList {
	s.ErrorMessage = &v
	return s
}

func (s *PutEventsResponseBodyDataEntryList) SetEventId(v string) *PutEventsResponseBodyDataEntryList {
	s.EventId = &v
	return s
}

func (s *PutEventsResponseBodyDataEntryList) SetTraceId(v string) *PutEventsResponseBodyDataEntryList {
	s.TraceId = &v
	return s
}

func (s *PutEventsResponseBodyDataEntryList) Validate() error {
	return dara.Validate(s)
}
