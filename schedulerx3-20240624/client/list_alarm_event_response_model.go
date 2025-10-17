// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlarmEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlarmEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlarmEventResponse
	GetStatusCode() *int32
	SetBody(v *ListAlarmEventResponseBody) *ListAlarmEventResponse
	GetBody() *ListAlarmEventResponseBody
}

type ListAlarmEventResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlarmEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlarmEventResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmEventResponse) GoString() string {
	return s.String()
}

func (s *ListAlarmEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlarmEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlarmEventResponse) GetBody() *ListAlarmEventResponseBody {
	return s.Body
}

func (s *ListAlarmEventResponse) SetHeaders(v map[string]*string) *ListAlarmEventResponse {
	s.Headers = v
	return s
}

func (s *ListAlarmEventResponse) SetStatusCode(v int32) *ListAlarmEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlarmEventResponse) SetBody(v *ListAlarmEventResponseBody) *ListAlarmEventResponse {
	s.Body = v
	return s
}

func (s *ListAlarmEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
