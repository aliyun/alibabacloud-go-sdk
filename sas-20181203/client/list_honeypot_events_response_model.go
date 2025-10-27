// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHoneypotEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHoneypotEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListHoneypotEventsResponseBody) *ListHoneypotEventsResponse
	GetBody() *ListHoneypotEventsResponseBody
}

type ListHoneypotEventsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHoneypotEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHoneypotEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotEventsResponse) GoString() string {
	return s.String()
}

func (s *ListHoneypotEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHoneypotEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHoneypotEventsResponse) GetBody() *ListHoneypotEventsResponseBody {
	return s.Body
}

func (s *ListHoneypotEventsResponse) SetHeaders(v map[string]*string) *ListHoneypotEventsResponse {
	s.Headers = v
	return s
}

func (s *ListHoneypotEventsResponse) SetStatusCode(v int32) *ListHoneypotEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHoneypotEventsResponse) SetBody(v *ListHoneypotEventsResponseBody) *ListHoneypotEventsResponse {
	s.Body = v
	return s
}

func (s *ListHoneypotEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
