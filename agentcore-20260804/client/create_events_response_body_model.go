// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateEventsResponseBody
	GetCode() *string
	SetData(v *CreateEventsResponseBodyData) *CreateEventsResponseBody
	GetData() *CreateEventsResponseBodyData
	SetHttpStatusCode(v int32) *CreateEventsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateEventsResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateEventsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateEventsResponseBody
	GetSuccess() *bool
}

type CreateEventsResponseBody struct {
	// The business result code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The operation result.
	Data *CreateEventsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request processing description.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID used for troubleshooting.
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful. For batch operations, you must still check the result of each item.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateEventsResponseBody) GetData() *CreateEventsResponseBodyData {
	return s.Data
}

func (s *CreateEventsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateEventsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEventsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateEventsResponseBody) SetCode(v string) *CreateEventsResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEventsResponseBody) SetData(v *CreateEventsResponseBodyData) *CreateEventsResponseBody {
	s.Data = v
	return s
}

func (s *CreateEventsResponseBody) SetHttpStatusCode(v int32) *CreateEventsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateEventsResponseBody) SetMessage(v string) *CreateEventsResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEventsResponseBody) SetRequestId(v string) *CreateEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventsResponseBody) SetSuccess(v bool) *CreateEventsResponseBody {
	s.Success = &v
	return s
}

func (s *CreateEventsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEventsResponseBodyData struct {
	// The write result for each item. A successful top-level response does not indicate that every item was successful.
	//
	// This parameter is required.
	Results []*CreateEventsResponseBodyDataResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s CreateEventsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEventsResponseBodyData) GetResults() []*CreateEventsResponseBodyDataResults {
	return s.Results
}

func (s *CreateEventsResponseBodyData) SetResults(v []*CreateEventsResponseBodyDataResults) *CreateEventsResponseBodyData {
	s.Results = v
	return s
}

func (s *CreateEventsResponseBodyData) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEventsResponseBodyDataResults struct {
	// The error for the failed item.
	Error *CreateEventsResponseBodyDataResultsError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The event ID generated after the event is successfully saved.
	//
	// example:
	//
	// evt-2f5f9c5a8d0e4c3b
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// The time when the event reached the server after being successfully saved, in Unix milliseconds.
	//
	// example:
	//
	// 1758614401000
	ReceivedAt *int64 `json:"receivedAt,omitempty" xml:"receivedAt,omitempty"`
	// The status. Valid values: STORED and FAILED.
	//
	// This parameter is required.
	//
	// example:
	//
	// STORED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateEventsResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s CreateEventsResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *CreateEventsResponseBodyDataResults) GetError() *CreateEventsResponseBodyDataResultsError {
	return s.Error
}

func (s *CreateEventsResponseBodyDataResults) GetEventId() *string {
	return s.EventId
}

func (s *CreateEventsResponseBodyDataResults) GetReceivedAt() *int64 {
	return s.ReceivedAt
}

func (s *CreateEventsResponseBodyDataResults) GetStatus() *string {
	return s.Status
}

func (s *CreateEventsResponseBodyDataResults) SetError(v *CreateEventsResponseBodyDataResultsError) *CreateEventsResponseBodyDataResults {
	s.Error = v
	return s
}

func (s *CreateEventsResponseBodyDataResults) SetEventId(v string) *CreateEventsResponseBodyDataResults {
	s.EventId = &v
	return s
}

func (s *CreateEventsResponseBodyDataResults) SetReceivedAt(v int64) *CreateEventsResponseBodyDataResults {
	s.ReceivedAt = &v
	return s
}

func (s *CreateEventsResponseBodyDataResults) SetStatus(v string) *CreateEventsResponseBodyDataResults {
	s.Status = &v
	return s
}

func (s *CreateEventsResponseBodyDataResults) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEventsResponseBodyDataResultsError struct {
	// The error code.
	//
	// This parameter is required.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description.
	//
	// This parameter is required.
	//
	// example:
	//
	// The event content format is incorrect
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateEventsResponseBodyDataResultsError) String() string {
	return dara.Prettify(s)
}

func (s CreateEventsResponseBodyDataResultsError) GoString() string {
	return s.String()
}

func (s *CreateEventsResponseBodyDataResultsError) GetCode() *string {
	return s.Code
}

func (s *CreateEventsResponseBodyDataResultsError) GetMessage() *string {
	return s.Message
}

func (s *CreateEventsResponseBodyDataResultsError) SetCode(v string) *CreateEventsResponseBodyDataResultsError {
	s.Code = &v
	return s
}

func (s *CreateEventsResponseBodyDataResultsError) SetMessage(v string) *CreateEventsResponseBodyDataResultsError {
	s.Message = &v
	return s
}

func (s *CreateEventsResponseBodyDataResultsError) Validate() error {
	return dara.Validate(s)
}
