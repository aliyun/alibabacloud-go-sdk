// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertEventsResponseBody) *ListAlertEventsResponse
	GetBody() *ListAlertEventsResponseBody
}

type ListAlertEventsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponse) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertEventsResponse) GetBody() *ListAlertEventsResponseBody {
	return s.Body
}

func (s *ListAlertEventsResponse) SetHeaders(v map[string]*string) *ListAlertEventsResponse {
	s.Headers = v
	return s
}

func (s *ListAlertEventsResponse) SetStatusCode(v int32) *ListAlertEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertEventsResponse) SetBody(v *ListAlertEventsResponseBody) *ListAlertEventsResponse {
	s.Body = v
	return s
}

func (s *ListAlertEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
