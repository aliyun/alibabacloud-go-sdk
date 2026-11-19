// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListEventsResponseBody) *ListEventsResponse
	GetBody() *ListEventsResponseBody
}

type ListEventsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponse) GoString() string {
	return s.String()
}

func (s *ListEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEventsResponse) GetBody() *ListEventsResponseBody {
	return s.Body
}

func (s *ListEventsResponse) SetHeaders(v map[string]*string) *ListEventsResponse {
	s.Headers = v
	return s
}

func (s *ListEventsResponse) SetStatusCode(v int32) *ListEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventsResponse) SetBody(v *ListEventsResponseBody) *ListEventsResponse {
	s.Body = v
	return s
}

func (s *ListEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
