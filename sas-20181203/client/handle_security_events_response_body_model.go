// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleSecurityEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHandleSecurityEventsResponse(v *HandleSecurityEventsResponseBodyHandleSecurityEventsResponse) *HandleSecurityEventsResponseBody
	GetHandleSecurityEventsResponse() *HandleSecurityEventsResponseBodyHandleSecurityEventsResponse
	SetRequestId(v string) *HandleSecurityEventsResponseBody
	GetRequestId() *string
}

type HandleSecurityEventsResponseBody struct {
	// The handling result of the alert events.
	HandleSecurityEventsResponse *HandleSecurityEventsResponseBodyHandleSecurityEventsResponse `json:"HandleSecurityEventsResponse,omitempty" xml:"HandleSecurityEventsResponse,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// FF0020B9-999F-5DE2-985F-DB282BDA5311
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s HandleSecurityEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HandleSecurityEventsResponseBody) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsResponseBody) GetHandleSecurityEventsResponse() *HandleSecurityEventsResponseBodyHandleSecurityEventsResponse {
	return s.HandleSecurityEventsResponse
}

func (s *HandleSecurityEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HandleSecurityEventsResponseBody) SetHandleSecurityEventsResponse(v *HandleSecurityEventsResponseBodyHandleSecurityEventsResponse) *HandleSecurityEventsResponseBody {
	s.HandleSecurityEventsResponse = v
	return s
}

func (s *HandleSecurityEventsResponseBody) SetRequestId(v string) *HandleSecurityEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *HandleSecurityEventsResponseBody) Validate() error {
	if s.HandleSecurityEventsResponse != nil {
		if err := s.HandleSecurityEventsResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HandleSecurityEventsResponseBodyHandleSecurityEventsResponse struct {
	// The ID of the task to handle the alert events.
	//
	// example:
	//
	// 15411
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s HandleSecurityEventsResponseBodyHandleSecurityEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s HandleSecurityEventsResponseBodyHandleSecurityEventsResponse) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsResponseBodyHandleSecurityEventsResponse) GetTaskId() *int64 {
	return s.TaskId
}

func (s *HandleSecurityEventsResponseBodyHandleSecurityEventsResponse) SetTaskId(v int64) *HandleSecurityEventsResponseBodyHandleSecurityEventsResponse {
	s.TaskId = &v
	return s
}

func (s *HandleSecurityEventsResponseBodyHandleSecurityEventsResponse) Validate() error {
	return dara.Validate(s)
}
