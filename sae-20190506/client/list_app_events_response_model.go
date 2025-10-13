// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppEventsResponseBody) *ListAppEventsResponse
	GetBody() *ListAppEventsResponseBody
}

type ListAppEventsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppEventsResponse) GoString() string {
	return s.String()
}

func (s *ListAppEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppEventsResponse) GetBody() *ListAppEventsResponseBody {
	return s.Body
}

func (s *ListAppEventsResponse) SetHeaders(v map[string]*string) *ListAppEventsResponse {
	s.Headers = v
	return s
}

func (s *ListAppEventsResponse) SetStatusCode(v int32) *ListAppEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppEventsResponse) SetBody(v *ListAppEventsResponseBody) *ListAppEventsResponse {
	s.Body = v
	return s
}

func (s *ListAppEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
