// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduleEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScheduleEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScheduleEventResponse
	GetStatusCode() *int32
	SetBody(v *ListScheduleEventResponseBody) *ListScheduleEventResponse
	GetBody() *ListScheduleEventResponseBody
}

type ListScheduleEventResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduleEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduleEventResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleEventResponse) GoString() string {
	return s.String()
}

func (s *ListScheduleEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScheduleEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScheduleEventResponse) GetBody() *ListScheduleEventResponseBody {
	return s.Body
}

func (s *ListScheduleEventResponse) SetHeaders(v map[string]*string) *ListScheduleEventResponse {
	s.Headers = v
	return s
}

func (s *ListScheduleEventResponse) SetStatusCode(v int32) *ListScheduleEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduleEventResponse) SetBody(v *ListScheduleEventResponseBody) *ListScheduleEventResponse {
	s.Body = v
	return s
}

func (s *ListScheduleEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
