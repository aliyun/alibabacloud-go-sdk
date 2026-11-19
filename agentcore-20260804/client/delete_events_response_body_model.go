// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteEventsResponseBody
	GetCode() *string
	SetData(v *DeleteEventsResponseBodyData) *DeleteEventsResponseBody
	GetData() *DeleteEventsResponseBodyData
	SetHttpStatusCode(v int32) *DeleteEventsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteEventsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEventsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteEventsResponseBody
	GetSuccess() *bool
}

type DeleteEventsResponseBody struct {
	// The business result code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The operation result.
	Data *DeleteEventsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The description of the request processing result.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which is used for troubleshooting.
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful. For batch operations, you must still check the result of each item.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteEventsResponseBody) GetData() *DeleteEventsResponseBodyData {
	return s.Data
}

func (s *DeleteEventsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteEventsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteEventsResponseBody) SetCode(v string) *DeleteEventsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventsResponseBody) SetData(v *DeleteEventsResponseBodyData) *DeleteEventsResponseBody {
	s.Data = v
	return s
}

func (s *DeleteEventsResponseBody) SetHttpStatusCode(v int32) *DeleteEventsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteEventsResponseBody) SetMessage(v string) *DeleteEventsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEventsResponseBody) SetRequestId(v string) *DeleteEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventsResponseBody) SetSuccess(v bool) *DeleteEventsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEventsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteEventsResponseBodyData struct {
	// The itemized deletion results.
	//
	// This parameter is required.
	Results []*DeleteEventsResponseBodyDataResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s DeleteEventsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteEventsResponseBodyData) GetResults() []*DeleteEventsResponseBodyDataResults {
	return s.Results
}

func (s *DeleteEventsResponseBodyData) SetResults(v []*DeleteEventsResponseBodyDataResults) *DeleteEventsResponseBodyData {
	s.Results = v
	return s
}

func (s *DeleteEventsResponseBodyData) Validate() error {
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

type DeleteEventsResponseBodyDataResults struct {
	// The error for the failed item.
	Error *DeleteEventsResponseBodyDataResultsError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The ID of the event requested to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// evt-2f5f9c5a8d0e4c3b
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// The status. Valid values: DELETED, NOT_FOUND, and FAILED.
	//
	// This parameter is required.
	//
	// example:
	//
	// DELETED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeleteEventsResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventsResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *DeleteEventsResponseBodyDataResults) GetError() *DeleteEventsResponseBodyDataResultsError {
	return s.Error
}

func (s *DeleteEventsResponseBodyDataResults) GetEventId() *string {
	return s.EventId
}

func (s *DeleteEventsResponseBodyDataResults) GetStatus() *string {
	return s.Status
}

func (s *DeleteEventsResponseBodyDataResults) SetError(v *DeleteEventsResponseBodyDataResultsError) *DeleteEventsResponseBodyDataResults {
	s.Error = v
	return s
}

func (s *DeleteEventsResponseBodyDataResults) SetEventId(v string) *DeleteEventsResponseBodyDataResults {
	s.EventId = &v
	return s
}

func (s *DeleteEventsResponseBodyDataResults) SetStatus(v string) *DeleteEventsResponseBodyDataResults {
	s.Status = &v
	return s
}

func (s *DeleteEventsResponseBodyDataResults) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteEventsResponseBodyDataResultsError struct {
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

func (s DeleteEventsResponseBodyDataResultsError) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventsResponseBodyDataResultsError) GoString() string {
	return s.String()
}

func (s *DeleteEventsResponseBodyDataResultsError) GetCode() *string {
	return s.Code
}

func (s *DeleteEventsResponseBodyDataResultsError) GetMessage() *string {
	return s.Message
}

func (s *DeleteEventsResponseBodyDataResultsError) SetCode(v string) *DeleteEventsResponseBodyDataResultsError {
	s.Code = &v
	return s
}

func (s *DeleteEventsResponseBodyDataResultsError) SetMessage(v string) *DeleteEventsResponseBodyDataResultsError {
	s.Message = &v
	return s
}

func (s *DeleteEventsResponseBodyDataResultsError) Validate() error {
	return dara.Validate(s)
}
